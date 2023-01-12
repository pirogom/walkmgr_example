package main

import (
	"fmt"
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
	item := treeView.AddItem("하하하하", icon1)
	item2 := item.AddItem("히히히", icon2)
	item2.AddItem("ㅋㅋㅋㅋ", icon1)
	treeView.Update()

	wm.Start()
}

func TestTreeView2(t *testing.T) {
	var addIdx int

	icon1, _ := walk.NewIconFromFile("./icon/icon1.ico")
	icon2, _ := walk.NewIconFromFile("./icon/icon2.ico")

	wm := walkmgr.NewWin("트리뷰 예제2(추가)", 640, 480)

	treeView := wm.NewTreeView()
	item := treeView.AddItem("하하하하", icon1)
	item2 := item.AddItem("히히히", icon2)
	item2.AddItem("ㅋㅋㅋㅋ", icon1)
	treeView.Update()

	wm.PushButton("추가", func() {
		addIdx++
		treeView.AddItem(fmt.Sprintf("추가됨%d", addIdx), icon1)
		treeView.Update()
	})

	wm.Start()
}
