package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

func newSearchWindow() {
	searchWindow := timerApp.NewWindow("질리언 타이머")
	searchWindow.Resize(fyne.NewSize(searchWindowWidth, searchWindowHeight))
	searchWindow.SetFixedSize(true) // 창 크기 조절 막기

	// 닫을 때 작동
	// searchWindow.SetOnClosed(func() {
	// 	// TODO : 검색 history 저장
	// })

	// Update
	// TODO
	// go clock.animate(searchWindow.Canvas(), searchWindow)

	// Set search content
	content := getSearchContent()

	// Run
	searchWindow.SetContent(content)
	searchWindow.ShowAndRun()
}

// getSearchContent returns layout
func getSearchContent() *widget.Box {

	// layout 1 : 소환사 이름
	searchLabel := widget.NewLabel("소환사 이름")

	nameEntry := newEnterEntry()
	nameEntry.SetPlaceHolder("소환사명을 입력해주세요")

	startButton := widget.NewButton("시작", func() {
		// 입력한 이름을 parameter로 넘김
		summonerName = nameEntry.Text
		newListWindow()
	})
	startButton.Resize(fyne.NewSize(200, 100))

	searchBox := widget.NewVBox(
		searchLabel,
		nameEntry,
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
		tmpName := name // button function bind를 위해 for loop 안에 변수 추가
		tmpButton := widget.NewButton(name, func() {
			summonerName = tmpName
			newListWindow()
		})
		historyContainer.AddObject(tmpButton)
	}

	historyBox := widget.NewVBox(
		historyLabel,
		historyContainer,
	)

	return widget.NewVBox(searchBox, historyBox)
}
