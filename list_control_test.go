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

	lc := walkmgr.NewListControl(wm,
		walkmgr.LISTCONTROL_CHECKBOX,
		walkmgr.LISTCONTROL_MULTI_SELECT)

	lc.AddColumn("이름", 300)
	lc.AddColumn("나이", 50)
	lc.AddColumn("성별", 50)
	lc.AddColumn("상태", 50)
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
