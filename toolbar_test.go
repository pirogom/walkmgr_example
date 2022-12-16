package main

import (
	"fmt"
	"testing"

	"github.com/pirogom/walk"
	"github.com/pirogom/walkmgr"
)

func TestToolbar1(t *testing.T) {
	icon1, _ := walk.NewIconFromFile("icon\\icon1.ico")

	wm := walkmgr.NewWin("툴바 테스트", 640, 480)

	tb, _ := walk.NewToolBar(wm.Window())

	act1 := walk.NewAction()
	act1.SetText("하하하하")
	act1.Triggered().Attach(func() {
		fmt.Println("하하하하")
	})
	act1.SetImage(icon1)
	tb.Actions().Add(act1)

	wm.Window().SetToolBar(tb)
	wm.Start()
}
