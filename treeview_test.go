package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/pirogom/walk"
	"github.com/pirogom/walkmgr"
)

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
		newItem := treeView.AddItem(fmt.Sprintf("추가됨%d", itemCnt), icon2)
		treeView.UpdateItems()
		treeView.SetCurrentItem(newItem)
	})

	wm.PushButton("서브 추가", func() {

		currItem := treeView.CurrentItem()

		if currItem == nil {
			walkmgr.MsgBox("트리뷰 아이템을 선택하세요.")
			return
		}

		itemCnt++
		newItem := currItem.AddItem(fmt.Sprintf("추가된 서브아이템%d", itemCnt), icon2)
		treeView.UpdateItems()
		treeView.SetCurrentItem(newItem)
		treeView.SetExpanded(currItem, true)

	})

	wm.Start()
}

func TestTreeView3(t *testing.T) {
	icon1, _ := walk.NewIconFromFile("./icon/icon1.ico")
	icon2, _ := walk.NewIconFromFile("./icon/icon2.ico")

	wm := walkmgr.NewWin("트리뷰 예제", 640, 480)

	treeView := wm.NewTreeView()
	treeView.AddItem("1", icon1)
	treeView.AddItem("2", icon2)
	treeView.AddItem("3", icon1)
	treeView.UpdateItems()

	wm.PushButton("선택 아이템 이름 변경", func() {
		ci := treeView.CurrentItem()

		if ci != nil {
			ci.SetText(fmt.Sprintf("변경된 이름 - %d", time.Now().Unix()))
			treeView.UpdateItems()
			treeView.SetCurrentItem(ci)
		}
	})

	wm.Start()
}

func TestTreeView4(t *testing.T) {
	icon1, _ := walk.NewIconFromFile("./icon/icon1.ico")
	icon2, _ := walk.NewIconFromFile("./icon/icon2.ico")

	wm := walkmgr.NewWin("트리뷰 예제", 640, 480)

	treeView := wm.NewTreeView()
	treeView.AddItem("1", icon1)
	treeView.AddItem("2", icon2)
	treeView.AddItem("3", icon1)
	treeView.AddItem("4", icon1)
	treeView.AddItem("5", icon1)
	treeView.AddItem("6", icon1)
	treeView.AddItem("7", icon1)
	treeView.AddItem("8", icon1)
	treeView.UpdateItems()

	treeView.Tv().CurrentItemChanged().Attach(func() {
		ci := treeView.CurrentItem()
		if ci != nil {
			walkmgr.MsgBox(ci.Text(), wm.Window())
		}
	})

	wm.Start()
}

func TestTreeView5(t *testing.T) {
	icon1, _ := walk.NewIconFromFile("./icon/icon1.ico")
	icon2, _ := walk.NewIconFromFile("./icon/icon2.ico")

	wm := walkmgr.NewWin("트리뷰 예제", 640, 480)

	treeView := wm.NewTreeView()
	item1 := treeView.AddItem("1", icon1)
	data1 := "This is data number 1"
	item1.SetData(&data1)
	item2 := treeView.AddItem("2", icon2)
	data2 := "This is data number 2"
	item2.SetData(&data2)
	item3 := treeView.AddItem("3", icon1)
	data3 := "This is data number 3"
	item3.SetData(&data3)
	treeView.UpdateItems()

	treeView.Tv().CurrentItemChanged().Attach(func() {
		ci := treeView.CurrentItem()
		if ci != nil {
			data := ci.Data().(*string)
			if data != nil {
				walkmgr.MsgBox("Item Data : "+*data, wm.Window())

				fi := treeView.FindItem(data)

				if fi != nil {
					walkmgr.MsgBox("Find item from data : " + (*fi).Text())
				}
			}
		}
	})

	wm.Start()
}
