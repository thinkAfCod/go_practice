package main

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"log"
	"strings"
)

var isSpecialMode = walk.NewMutableCondition()

type MyMainWindow struct {
	*walk.MainWindow
}
type FileDropTextEdit struct {
	TextEdit
	OnDropFiles walk.DropFilesEventHandler
}

func main() {
	MustRegisterCondition("isSpecialMode", isSpecialMode)

	mw := new(MyMainWindow)

	var filePath1 *walk.TextEdit
	var filePath2 *walk.TextEdit

	if err := (MainWindow{
		AssignTo: &mw.MainWindow,
		Title:    "Walk Actions Example",
		MinSize:  Size{320, 240},
		Layout:   VBox{},
		OnDropFiles: func(files []string) {
			log.Println("OnDropFiles event has been published")
		},
		Children: []Widget{
			TextEdit{
				AssignTo: &filePath1,
				Text:     "",
			},
			TextEdit{
				AssignTo: &filePath2,
				Text:     "",
			},
		},
	}.Create()); err != nil {
		log.Fatal(err)
	}

	filePath1.DropFiles().Attach(func(files []string) {
		filePath1.SetText(strings.Join(files, "\r\n"))
	})
	filePath2.DropFiles().Attach(func(files []string) {
		filePath2.SetText(filePath2.Text() + strings.Join(files, "\r\n"))
	})

	mw.Run()
}
