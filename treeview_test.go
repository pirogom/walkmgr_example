package main

import (
	"testing"

	"github.com/pirogom/walk"
	"github.com/pirogom/walkmgr"
)

// under construction

func TestTreeView(t *testing.T) {
	icon1, _ := walk.NewIconFromFile("./icon/icon1.ico")
	icon2, _ := walk.NewIconFromFile("./icon/icon2.ico")

	wm := walkmgr.NewWin("트리뷰 예제", 640, 480)

	treeView := wm.NewTreeView()
	item := treeView.AddItem("1", icon1)
	item2 := item.AddItem("1-1", icon2)
	item2.AddItem("1-1-1", icon1)
	item2.AddItem("1-1-2", icon2)
	item2.AddItem("1-1-3", icon1)
	item2.AddItem("1-1-4", icon2)
	item2.AddItem("1-1-5", icon1)
	item2.AddItem("1-1-6", icon2)

	item3 := item.AddItem("1-2", icon2)
	item3.AddItem("1-2-1", icon1)
	item3.AddItem("1-2-2", icon2)
	item3.AddItem("1-2-3", icon1)
	item3.AddItem("1-2-4", icon2)
	item3.AddItem("1-2-5", icon1)
	item3.AddItem("1-2-6", icon2)

	item4 := treeView.AddItem("2", icon1)
	item4.AddItem("2-1", icon2)

	treeView.Create()

	wm.Start()
}
