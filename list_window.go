package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

// summonerName : 소환사 이름
func newListWindow() {
	listWindow := timerApp.NewWindow("타이머 선택")
	listWindow.Resize(fyne.NewSize(listWindowWidth, listWindowHeight))
	listWindow.SetFixedSize(true) // 창 크기 조절 막기

	// TODO : 소환사 이름으로 Riot API에서 현재 게임 정보 받아오기
	// TODO : 강타를 든 경우에는 정글 타이머 띄우기

	// Update
	// TODO
	// go clock.animate(listWindow.Canvas(), listWindow)

	// Set search content
	content := getListContent()

	// Run
	listWindow.SetContent(content)
	listWindow.Show()
}

// getListContent returns layout
func getListContent() *widget.Box {
	// TODO : 정글 타이머, 스펠 타이머 열 수 있는 버튼

	// layout 2 : 정글 타이머, 스펠 타이머 선택

	nameLabel := widget.NewLabel("소환사 : " + summonerName)

	jungleButton := widget.NewButton("정글", func() {
		// 입력한 이름을 parameter로 넘김
		// newJungleWindow()
	})
	spellButton := widget.NewButton("스펠", func() {
		// 입력한 이름을 parameter로 넘김
		// newSpellWindow()
	})

	return widget.NewVBox(nameLabel, jungleButton, spellButton)
}
