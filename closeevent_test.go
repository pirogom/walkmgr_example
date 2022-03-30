package main

import (
	"testing"

	"github.com/pirogom/walkmgr"
)

// 창 닫힘 이벤트 관련

func TestIgnoreClose(t *testing.T) {
	wm := walkmgr.NewWin("닫지 못하는 창", 640, 480)

	wm.PushButton("강제닫음", func() {
		wm.Close()
	})
	wm.IgnoreClosing()

	wm.Start()
}

func TestCloseCheck(t *testing.T) {
	wm := walkmgr.NewWin("창 닫음 이벤트", 640, 480)

	wm.Closing(func() bool {
		return walkmgr.Confirm("닫을까요?")
	})

	wm.Start()
}
