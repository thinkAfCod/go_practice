package main

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"log"
	"strings"
)

var isSpecialMode = walk.NewMutableCondition()

func main() {
	MustRegisterCondition("isSpecialMode", isSpecialMode)

	var mw *walk.MainWindow
	var tabWidget *walk.TabWidget

	if err := (MainWindow{
		AssignTo: &mw,
		Title:    "MyFirstApp",
		MinSize:  Size{320, 240},
		Layout:   VBox{},
		OnDropFiles: func(files []string) {
			log.Println("OnDropFiles event has been published")
		},
		Children: []Widget{
			TabWidget{
				AssignTo: &tabWidget,
				Pages: []TabPage{
					NewFileToolTabPage(),
				},
			},
		},
	}.Create()); err != nil {
		log.Fatal(err)
	}

	mw.Run()
}

func NewFileToolTabPage() TabPage {
	var tab *walk.TabPage
	var fileDialogButton *walk.PushButton
	var fileMoveButton *walk.PushButton
	var fileDropText *walk.TextEdit
	var filePath string
	return TabPage{
		AssignTo: &tab,
		Title:    "文件",
		Layout:   VBox{},
		Children: []Widget{
			GroupBox{
				Title:  "移除中间层文件夹(保留最内层文件夹)",
				Layout: Grid{Rows: 1},
				Children: []Widget{
					CustTextEdit{
						AssignTo: &fileDropText,
						OnTextChanged: func() {
							filePath = fileDropText.Text()
						},
						OnDropFiles: func(files []string) {
							fileDropText.SetText(strings.Join(files, ";"))
						},
					},
					PushButton{
						AssignTo: &fileDialogButton,
						Text:     "...",
						MinSize:  Size{Height: 1},
						OnClicked: func() {
							log.Println("...")
						},
					},
					PushButton{
						AssignTo: &fileMoveButton,
						Text:     "移除",
						OnClicked: func() {
							log.Println("移除")
						},
					},
				},
			},
			VSpacer{},
		},
	}
}
