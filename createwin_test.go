package main

import (
	"testing"

	"github.com/pirogom/walkmgr"
)

/**
* 기본 윈도 생성 옵션에 대한 예제들
**/

func TestDefaultWin(t *testing.T) {
	wm := walkmgr.NewWin("기본윈도", 640, 480)
	wm.Start()
}

func TestForegroundWin(t *testing.T) {
	wm := walkmgr.NewWin("기본윈도-최상단에 생성", 640, 480)
	wm.StartForeground()
}

func TestNoMinBox(t *testing.T) {
	wm := walkmgr.NewWin("기본윈도-최소화X", 640, 480)
	wm.DisableMinBox()

	wm.Start()
}

func TestNoMaxBox(t *testing.T) {
	wm := walkmgr.NewWin("기본윈도-최대화X", 640, 480)
	wm.DisableMaxBox()

	wm.Start()
}

func TestNoSysMenu(t *testing.T) {
	wm := walkmgr.NewWin("기본윈도-시스템메뉴X", 640, 480)
	wm.DisableSysmenu()

	wm.Start()
}

func TestNoTitle(t *testing.T) {
	wm := walkmgr.NewWin("기본윈도-타이틀바X", 640, 480)
	wm.DisableTitleBar()

	wm.Start()
}

func TestNoResize(t *testing.T) {
	wm := walkmgr.NewWin("기본윈도-사이즈변경X", 640, 480)
	wm.NoResize()
	wm.Start()
}
