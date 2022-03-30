package main

import (
	"testing"

	"github.com/pirogom/walkmgr"
)

func TestHSplit(t *testing.T) {
	wm := walkmgr.NewWin("HSplit", 640, 480)

	wm.Split(walkmgr.LAYOUT_HORI)
	wm.PushButton("버튼1", func() {
	})
	wm.PushButton("버튼2", func() {
	})
	wm.End()

	wm.Start()
}

func TestVSplit(t *testing.T) {
	wm := walkmgr.NewWin("VSplit", 640, 480)

	wm.Split(walkmgr.LAYOUT_VERT)
	wm.PushButton("버튼1", func() {
	})
	wm.PushButton("버튼2", func() {
	})
	wm.End()

	wm.Start()
}
