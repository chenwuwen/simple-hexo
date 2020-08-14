package ui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
	"simple-hexo/command"
	c "simple-hexo/const"
)

var ReviewButton *widget.Button
var DeployButton *widget.Button
var CurDirLabel *widget.Label

func CreateReviewButton(w fyne.Window) *widget.Button {

	ReviewButton = widget.NewButton(c.BUTTON_NAME_REVIEW, func() {
		result, state := command.Server(4000)
		ServerResult(result, state, w)
	})
	return ReviewButton
}

func CreateDeployButton(w fyne.Window) *widget.Button {
	DeployButton = widget.NewButton(c.BUTTON_NAME_DEPLOY, func() {
		result, state := command.Deploy()
		DeployResult(result, state, w)
	})
	return DeployButton
}

func ShowCurDir() *widget.Label {
	CurDirLabel = widget.NewLabel("Current Dir:  " + command.CurrentDir())

	return CurDirLabel
}
