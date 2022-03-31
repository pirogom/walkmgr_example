package main

import (
	"testing"

	"github.com/pirogom/walk"
	"github.com/pirogom/walkmgr"
)

// 메뉴 관련

func TestMenu(t *testing.T) {
	wm := walkmgr.NewWin("메뉴 테스트", 640, 480)

	menubar, _ := walk.NewMenu()
	//

	menu1 := walk.NewMenuAction(menubar)
	menu1.SetText("메뉴1")

	act1 := walk.NewAction()
	act1.SetText("서브메뉴1")

	menubar.Actions().Add(act1)

	act2 := walk.NewAction()
	act2.SetText("서브메뉴2")

	menubar.Actions().Add(act2)

	sep1 := walk.NewSeparatorAction()
	menubar.Actions().Add(sep1)

	act3 := walk.NewAction()
	act3.SetText("서브메뉴3")

	menubar.Actions().Add(act3)

	//
	submenubar, _ := walk.NewMenu()
	submenu1 := walk.NewMenuAction(submenubar)
	submenu1.SetText("섭섭메뉴1")

	menubar.Actions().Add(submenu1)
	//

	//
	wm.Window().Menu().Actions().Add(menu1)

	wm.Start()
}
