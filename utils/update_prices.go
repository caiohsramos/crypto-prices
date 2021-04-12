package utils

import (
	"fmt"
	"go-crypto-prices/components"
	"go-crypto-prices/services"

	"fyne.io/fyne/v2/widget"
)

func UpdatePrices(service *services.CMCService, cards ...*widget.Card) {
	if len(cards) == 0 {
		return
	}
	var symbols []string
	for _, card := range cards {
		symbols = append(symbols, card.Subtitle)
		components.SetCardContent(card, "...")
	}

	quotes, err := service.GetQuotes(symbols)
	if err != nil {
		fmt.Println("Error requesting quotes\n", err)
		return
	}

	for _, card := range cards {
		for _, quote := range quotes {
			if quote.Symbol == card.Subtitle {
				price := fmt.Sprintf("%.2f", quote.Quote.Usd.Price)
				components.SetCardContent(card, "USD "+price)
			}
		}
	}

}
