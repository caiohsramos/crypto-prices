package main

import (
	"go-crypto-prices/components"
	"go-crypto-prices/services"
	"go-crypto-prices/utils"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func main() {
	CMCService := services.NewCMCService("API-KEY")
	cryptoPrices := app.New()
	mainWindow := cryptoPrices.NewWindow("Main")

	ethCard := components.NewEthCard()
	btcCard := components.NewBtcCard()
	refreshButton := components.NewRefreshButton(func() { utils.UpdatePrices(CMCService, ethCard, btcCard) })

	cryptoCards := container.New(layout.NewGridLayout(2), ethCard, btcCard)
	vBox := container.New(layout.NewVBoxLayout(), cryptoCards, layout.NewSpacer(), refreshButton)
	mainWindow.SetContent(vBox)
	mainWindow.Resize(fyne.NewSize(400, 400))

	// First update
	refreshButton.OnTapped()

	mainWindow.ShowAndRun()
}
