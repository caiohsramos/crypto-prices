package components

import (
	"fyne.io/fyne/v2/widget"
)

func NewRefreshButton(callback func()) *widget.Button {
	return widget.NewButton("Refresh", callback)
}
