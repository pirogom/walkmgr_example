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

	treeView.UpdateItems()

	wm.Start()
}

func TestTreeView2(t *testing.T) {
	var itemCnt int
	icon1, _ := walk.NewIconFromFile("./icon/icon1.ico")
	icon2, _ := walk.NewIconFromFile("./icon/icon2.ico")

	wm := walkmgr.NewWin("트리뷰 예제", 640, 480)

	treeView := wm.NewTreeView()
	treeView.AddItem("1", icon1)

	treeView.UpdateItems()

	wm.PushButton("루트 추가", func() {
		itemCnt++
		treeView.AddItem(fmt.Sprintf("추가됨%d", itemCnt), icon2)
		treeView.UpdateItems()
	})

	wm.PushButton("서브 추가", func() {

		currItem := treeView.CurrentItem()

		if currItem == nil {
			walkmgr.MsgBox("트리뷰 아이템을 선택하세요.")
			return
		}

		itemCnt++
		currItem.AddItem(fmt.Sprintf("추가된 서브아이템%d", itemCnt), icon2)
		treeView.UpdateItems()
		treeView.SetCurrentItem(currItem)
		treeView.SetExpanded(currItem, true)

	})

	wm.Start()
}
