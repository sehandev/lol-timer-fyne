package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func errCheck(e error) {
	if e != nil {
		panic(e)
	}
}

const (
	searchWindowWidth  int = 300
	searchWindowHeight int = 400

	listWindowWidth  int = 300
	listWindowHeight int = 300

	spellWindowWidth  int = 500
	spellWindowHeight int = 1000
)

var timerApp fyne.App

func main() {

	// New fyne app
	timerApp = app.New()

	// Set fyne theme
	timerApp.Settings().SetTheme(&myTheme{})

	// Window 1 : Search - user name & history
	newSearchWindow()

	// clock := &clockLayout{}

	// // Set content
	// content := clock.render()

	// 기록 넣어둘 폴더와 파일 생성
	// dirPath := "C:\\Users\\KSD\\.lol_timer"
	// 폴더 있는지 확인 후 없으면 생성
	// if _, err := os.Stat(dirPath); os.IsNotExist(err) {
	// 	os.Mkdir(dirPath, 0700)
	// }

	// newpath := filepath.Join(dirPath, ".history")
	// testString := []byte("Ready to go!\n")

	// err := ioutil.WriteFile(newpath, testString, 0700)
	// errCheck(err)
}

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

func newListWindow(summonerName string) {
	listWindow := timerApp.NewWindow("질리언 타이머")
	listWindow.Resize(fyne.NewSize(listWindowWidth, listWindowHeight))
	listWindow.SetFixedSize(true) // 창 크기 조절 막기

	// Update
	// TODO
	// go clock.animate(listWindow.Canvas(), listWindow)

	// Set search content
	// content := getListContent()
	content := widget.NewLabel(summonerName)

	// Run
	listWindow.SetContent(content)
	listWindow.Show()
}
