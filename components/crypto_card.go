package components

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

func NewEthCard() *widget.Card {
	ethCard := widget.NewCard("Ethereum", "ETH", nil)
	SetCardContent(ethCard, "...")
	ethImage := canvas.NewImageFromFile("assets/eth.png")
	ethCard.SetImage(ethImage)
	ethImage.FillMode = canvas.ImageFillContain

	return ethCard
}

func NewBtcCard() *widget.Card {
	btcCard := widget.NewCard("Bitcoin", "BTC", nil)
	SetCardContent(btcCard, "...")
	btcImage := canvas.NewImageFromFile("assets/btc.png")
	btcCard.SetImage(btcImage)
	btcImage.FillMode = canvas.ImageFillContain

	return btcCard
}

func SetCardContent(card *widget.Card, text string) {
	content := canvas.NewText(text, color.White)
	content.Alignment = fyne.TextAlignCenter
	content.TextSize = 32.0

	card.SetContent(content)
}
