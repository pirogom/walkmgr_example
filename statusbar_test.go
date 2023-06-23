package main

import (
	"testing"

	"github.com/pirogom/walk"
	"github.com/pirogom/walkmgr"
)

// Status Bar 관련
func TestStatusBar_Basic(t *testing.T) {
	wm := walkmgr.NewWin("상태바(StatusBar)", 640, 480)

	wm.AddStatusItem("아이템1", 100, nil, nil)
	wm.AddStatusItem("아이템2", 100, nil, nil)
	wm.AddStatusItem("아이템3", 100, nil, nil)
	wm.AddStatusItem("아이템4", 100, nil, nil)
	wm.AddStatusItem("아이템5", 100, nil, nil)

	// 창크기 변경시 상태바 다시 그리기
	wm.Window().SizeChanged().Attach(func() {
		wm.Window().StatusBar().Invalidate()
	})

	wm.StartForeground()
}

func TestStatusBar_Icon(t *testing.T) {

	icon1, _ := walk.NewIconFromFile("icon\\icon1.ico")
	icon2, _ := walk.NewIconFromFile("icon\\icon2.ico")

	wm := walkmgr.NewWin("상태바+아이콘(StatusBar)", 640, 480)

	wm.AddStatusItem("아이템1", 100, icon1, nil)
	wm.AddStatusItem("아이템2", 100, icon2, nil)
	wm.AddStatusItem("아이템3", 100, icon1, nil)
	wm.AddStatusItem("아이템4", 100, icon2, nil)
	wm.AddStatusItem("아이템5", 100, icon1, nil)

	// 창크기 변경시 상태바 다시 그리기
	wm.Window().SizeChanged().Attach(func() {
		wm.Window().StatusBar().Invalidate()
	})

	wm.StartForeground()
}

func TestStatusBar_ClickEvent(t *testing.T) {
	icon1, _ := walk.NewIconFromFile("icon\\icon1.ico")
	icon2, _ := walk.NewIconFromFile("icon\\icon2.ico")

	wm := walkmgr.NewWin("상태바+아이콘+클릭(StatusBar)", 640, 480)

	var sb1, sb2 *walk.StatusBarItem

	sb1 = wm.AddStatusItem("아이템1", 200, icon1, func() {
		if sb1.Text() == "아이템1" {
			sb1.SetIcon(icon2)
			sb1.SetText("아이템2")
		} else {
			sb1.SetIcon(icon1)
			sb1.SetText("아이템1")
		}
	})
	sb2 = wm.AddStatusItem("아이템2", 200, icon2, func() {
		if sb2.Text() == "아이템2" {
			sb2.SetIcon(icon1)
			sb2.SetText("아이템1")
		} else {
			sb2.SetIcon(icon2)
			sb2.SetText("아이템2")
		}
	})

	// 창크기 변경시 상태바 다시 그리기
	wm.Window().SizeChanged().Attach(func() {
		wm.Window().StatusBar().Invalidate()
	})

	wm.StartForeground()
}
