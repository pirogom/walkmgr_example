package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/pirogom/walkmgr"
)

// 리스트 컨트롤, 미완성 .. 개발중
func TestListControl(t *testing.T) {
	wm := walkmgr.NewWin("리스트 컨트롤 테스트", 640, 480)

	lc := walkmgr.NewListControl(wm, &walkmgr.ListControlCfg{MultiSelect: true, CheckBox: true})

	lc.AddColumn("이름", 300, walkmgr.LISTCTRL_ALIGN_LEFT, false)
	lc.AddColumn("나이", 50, walkmgr.LISTCTRL_ALIGN_LEFT, false)
	lc.AddColumn("성별", 50, walkmgr.LISTCTRL_ALIGN_LEFT, false)
	lc.AddColumn("상태", 50, walkmgr.LISTCTRL_ALIGN_LEFT, false)
	lc.Create()

	wm.HComposite()
	wm.PushButton("값추가", func() {
		idx := lc.AddItemData(fmt.Sprintf("피로곰-%d", time.Now().Unix()))
		lc.SetItemData(idx, 1, "나이")
		lc.SetItemData(idx, 2, "성별")
		lc.SetItemData(idx, 3, "상태")
	})
	wm.End()

	wm.Start()
}

func TestListControlOrdering(t *testing.T) {
	wm := walkmgr.NewWin("리스트 컨트롤 테스트", 640, 480)

	lc := walkmgr.NewListControl(wm, &walkmgr.ListControlCfg{MultiSelect: true, CheckBox: true})

	lc.AddColumn("이름", 300, walkmgr.LISTCTRL_ALIGN_LEFT, true)
	lc.AddColumn("나이", 50, walkmgr.LISTCTRL_ALIGN_LEFT, true)
	lc.AddColumn("레벨", 50, walkmgr.LISTCTRL_ALIGN_LEFT, true)
	lc.Create()

	wm.Starting(func() {
		for i := 0; i < 10; i++ {
			nr := lc.AddItemData(fmt.Sprintf("피로곰-%d", i))
			lc.SetItemData(nr, 1, fmt.Sprintf("%d", i))
			lc.SetItemData(nr, 2, fmt.Sprintf("%d", 10-i))
		}
	})

	wm.Start()
}

func TestListControlItemDblClick(t *testing.T) {
	wm := walkmgr.NewWin("리스트 컨트롤 테스트", 640, 480)

	lc := walkmgr.NewListControl(wm, &walkmgr.ListControlCfg{MultiSelect: true, CheckBox: true})

	lc.AddColumn("이름", 300, walkmgr.LISTCTRL_ALIGN_LEFT, false)
	lc.AddColumn("나이", 50, walkmgr.LISTCTRL_ALIGN_LEFT, false)
	lc.AddColumn("레벨", 50, walkmgr.LISTCTRL_ALIGN_LEFT, false)
	lc.Create()

	// 더블클릭한 아이템의 정보를 출력
	lc.ItemDoubleClicked(func(idx int) {
		fmt.Println(lc.GetItemData(idx, 0))
		fmt.Println(lc.GetItemData(idx, 1))
		fmt.Println(lc.GetItemData(idx, 2))
	})

	wm.Starting(func() {
		for i := 0; i < 10; i++ {
			nr := lc.AddItemData(fmt.Sprintf("피로곰-%d", i))
			lc.SetItemData(nr, 1, fmt.Sprintf("%d", i))
			lc.SetItemData(nr, 2, fmt.Sprintf("%d", 10-i))
		}
	})

	wm.Start()
}
