package main

import (
	"testing"

	"github.com/pirogom/walkmgr"
)

func TestDialog(t *testing.T) {
	wm := walkmgr.NewWin("DialogTest", 640, 480)

	wm.PushButton("다이얼로그", func() {
		// var acceptPB, cancelPB *walk.PushButton
		// dlg, err := walk.NewDialog(wm.Window())

		// if err != nil {
		// 	return
		// }

		// dlg.SetTitle("Test DLG")
		// dlg.SetSize(walk.Size{Width: 300, Height: 200})
		// dlg.SetMinMaxSize(walk.Size{Width: 300, Height: 200}, walk.Size{Width: 300, Height: 200})
		// dlg.SetName("Test DLG")
		// dlg.SetLayout(walk.NewVBoxLayout())
		// dlg.SetDefaultButton(acceptPB)
		// dlg.SetCancelButton(cancelPB)

		// dlg.Run()

		dlg := walkmgr.NewDialog(wm.Window(), "TestDLG", 300, 300, nil)

		dlg.PushButton("테스트", func() {})
		dlg.PushButton("테스트1", func() {})
		dlg.PushButton("테스트2", func() { dlg.Close() })
		//dlg.IgnoreClosing()

		dlg.Start()
	})

	wm.Start()
}
