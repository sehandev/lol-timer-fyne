package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

// getSearchContent returns layout
func getSearchContent() *widget.Box {

	// layout 1 : 소환사 이름
	searchLabel := widget.NewLabel("소환사 이름")
	startButton := widget.NewButton("시작", func() {
		newListWindow("test")
	})
	startButton.Resize(fyne.NewSize(200, 100))

	searchBox := widget.NewVBox(
		searchLabel,
		widget.NewHBox(
			layout.NewSpacer(),
			startButton,
			layout.NewSpacer(),
		),
	)

	// layout 2 : 입력 기록
	historyLabel := widget.NewLabel("최근 기록")

	nameArr := []string{"ABC", "D한글체크D", "세한", "여덟글자제한인데"}
	historyContainer := fyne.NewContainerWithLayout(layout.NewAdaptiveGridLayout(2))
	for _, name := range nameArr {
		tmpButton := widget.NewButton(name, func() {
			// TODO : 버튼마다 이름이 안되고 for문 끝났을 때의 name이 일괄 적용됨
			newListWindow(name)
		})
		historyContainer.AddObject(tmpButton)
	}

	historyBox := widget.NewVBox(
		historyLabel,
		historyContainer,
	)

	return widget.NewVBox(searchBox, historyBox)
}
