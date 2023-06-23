package main

import (
	"testing"

	"github.com/pirogom/walkmgr"
)

func TestScrollView1(t *testing.T) {
	//wm := NewWin("", 640, 480) default layout is vertical
	wm := walkmgr.NewWin("스크롤뷰", 320, 320, walkmgr.LAYOUT_VERT)

	wm.ScrollView(true, true)

	wm.HGroupBox("그룹")
	wm.PushButton("버튼1", func() {
	})
	wm.PushButton("버튼2", func() {
	})
	wm.PushButton("버튼3", func() {
	})
	wm.PushButton("버튼4", func() {
	})
	wm.PushButton("버튼5", func() {
	})
	wm.PushButton("버튼6", func() {
	})
	wm.PushButton("버튼7", func() {
	})
	wm.PushButton("버튼8", func() {
	})
	wm.PushButton("버튼9", func() {
	})
	wm.PushButton("버튼10", func() {
	})
	wm.End() // END of HGroupBox

	wm.End()

	wm.Start()
}

func TestScrollView2(t *testing.T) {
	//wm := NewWin("", 640, 480) default layout is vertical
	wm := walkmgr.NewWin("스크롤뷰", 200, 200, walkmgr.LAYOUT_VERT)

	wm.ScrollView(true, true)

	wm.HGroupBox("그룹")
	wm.PushButton("버튼1", func() {
	})
	wm.PushButton("버튼2", func() {
	})
	wm.PushButton("버튼3", func() {
	})
	wm.PushButton("버튼4", func() {
	})
	wm.PushButton("버튼5", func() {
	})
	wm.PushButton("버튼6", func() {
	})
	wm.PushButton("버튼7", func() {
	})
	wm.PushButton("버튼8", func() {
	})
	wm.PushButton("버튼9", func() {
	})
	wm.PushButton("버튼10", func() {
	})
	wm.End() // END of HGroupBox

	wm.HGroupBox("그룹")
	wm.PushButton("버튼1", func() {
	})
	wm.PushButton("버튼2", func() {
	})
	wm.PushButton("버튼3", func() {
	})
	wm.PushButton("버튼4", func() {
	})
	wm.PushButton("버튼5", func() {
	})
	wm.PushButton("버튼6", func() {
	})
	wm.PushButton("버튼7", func() {
	})
	wm.PushButton("버튼8", func() {
	})
	wm.PushButton("버튼9", func() {
	})
	wm.PushButton("버튼10", func() {
	})
	wm.End() // END of HGroupBox

	wm.HGroupBox("그룹")
	wm.PushButton("버튼1", func() {
	})
	wm.PushButton("버튼2", func() {
	})
	wm.PushButton("버튼3", func() {
	})
	wm.PushButton("버튼4", func() {
	})
	wm.PushButton("버튼5", func() {
	})
	wm.PushButton("버튼6", func() {
	})
	wm.PushButton("버튼7", func() {
	})
	wm.PushButton("버튼8", func() {
	})
	wm.PushButton("버튼9", func() {
	})
	wm.PushButton("버튼10", func() {
	})
	wm.End() // END of HGroupBox

	wm.HGroupBox("그룹")
	wm.PushButton("버튼1", func() {
	})
	wm.PushButton("버튼2", func() {
	})
	wm.PushButton("버튼3", func() {
	})
	wm.PushButton("버튼4", func() {
	})
	wm.PushButton("버튼5", func() {
	})
	wm.PushButton("버튼6", func() {
	})
	wm.PushButton("버튼7", func() {
	})
	wm.PushButton("버튼8", func() {
	})
	wm.PushButton("버튼9", func() {
	})
	wm.PushButton("버튼10", func() {
	})
	wm.End() // END of HGroupBox

	wm.End()

	wm.Start()
}
