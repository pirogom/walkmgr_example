package main

import (
	"testing"

	"github.com/pirogom/walkmgr"
)

func TestDialog(t *testing.T) {
	wm := walkmgr.NewWin("DialogTest", 640, 480)

	wm.PushButton("다이얼로그", func() {
		dlg := walkmgr.NewDialog(wm.Window(), "TestDLG", 300, 300, nil)

		dlg.PushButton("테스트", func() {})
		dlg.PushButton("테스트1", func() {})
		dlg.PushButton("테스트2", func() { dlg.CloseDLG(0) })

		dlg.StartDLG()
	})

	wm.Start()
}
