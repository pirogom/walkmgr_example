package main

import (
	"testing"

	"github.com/pirogom/walkmgr"
)

// 기본 윈도 레이아웃 관련

func TestVertLayout(t *testing.T) {
	//wm := NewWin("", 640, 480) default layout is vertical
	wm := walkmgr.NewWin("LAYOUT_VERT", 640, 480, walkmgr.LAYOUT_VERT)

	wm.PushButton("버튼1", func() {
	})
	wm.PushButton("버튼2", func() {
	})

	wm.Start()
}

func TestHoriLayout(t *testing.T) {
	wm := walkmgr.NewWin("LAYOUT_HORI", 640, 480, walkmgr.LAYOUT_HORI)

	wm.PushButton("버튼1", func() {
	})
	wm.PushButton("버튼2", func() {
	})

	wm.Start()
}
