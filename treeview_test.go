package main

import (
	"testing"

	"github.com/pirogom/walkmgr"
)

func TestTreeView(t *testing.T) {

	wm := walkmgr.NewWin("트리뷰 예제", 640, 480)

	treeView := wm.NewTreeView("./icon/icon1.ico", "./icon/icon2.ico")

	item := treeView.AddItem("하하하하", nil)
	item.AddChildren("히히히", nil)

	treeView.Create()

	wm.Start()
}
