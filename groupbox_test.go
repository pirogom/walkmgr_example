package main

import (
	"testing"

	"github.com/pirogom/walkmgr"
)

// 그룹박스 관련

func TestGroupBox(t *testing.T) {
	wm := walkmgr.NewWin("GroupBox(vert)", 640, 480)

	wm.GroupBox("그룹박스(vert)", walkmgr.LAYOUT_VERT)
	wm.PushButton("버튼1", func() {
	})
	wm.PushButton("버튼2", func() {
	})
	wm.End()

	wm.Start()
}

func TestGroupBox2(t *testing.T) {
	wm := walkmgr.NewWin("GroupBox(hori)", 640, 480)

	wm.GroupBox("그룹박스(hori)", walkmgr.LAYOUT_HORI)
	wm.PushButton("버튼1", func() {
	})
	wm.PushButton("버튼2", func() {
	})
	wm.End()

	wm.Start()
}
