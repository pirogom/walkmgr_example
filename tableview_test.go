package main

import (
	"fmt"
	"testing"

	"github.com/pirogom/walk"
	"github.com/pirogom/walkmgr"
)

// 테이블뷰 관련

/**
*	testTvItem
**/
type testTvItem struct {
	Name    string
	Level   int
	Sex     int
	Class   string
	checked bool
}

/**
*	testTvModel
**/
type testTvModel struct {
	walk.TableModelBase
	items []testTvItem
}

func (m *testTvModel) RowCount() int {
	return len(m.items)
}

func (m *testTvModel) ResetRows() {
	m.items = nil
	m.PublishRowsReset()
}

/**
*	Value
**/
func (m *testTvModel) Value(row, col int) interface{} {
	item := m.items[row]

	switch col {
	case 0:
		return item.Name
	case 1:
		return item.Level
	case 2:
		if item.Sex == 0 {
			return "중성"
		} else if item.Sex == 1 {
			return "남성"
		} else if item.Sex == 2 {
			return "여성"
		} else {
			return "알수없음"
		}
	case 3:
		return item.Class
	}
	panic("unexpected col")
}

func TestTv(t *testing.T) {
	wm := walkmgr.NewWin("Table view", 640, 480)

	tm := new(testTvModel)
	th := walkmgr.NewTH("이름", 100)
	th.Add("레벨", 100).Add("성별", 100)
	th.Add("직업", 100)
	testTv := wm.TableView(tm, th.Get(), true, true)

	testTv.ItemActivated().Attach(func() {
		currIdx := testTv.CurrentIndex()

		if currIdx < 0 {
			return
		}

		fmt.Println("= 더블클릭 ========================")
		fmt.Println("선택된 아이템:", currIdx)
		fmt.Println("이름:", tm.items[currIdx].Name)
		fmt.Println("레벨:", tm.items[currIdx].Level)
		fmt.Println("성별:", tm.items[currIdx].Sex)
		fmt.Println("직업:", tm.items[currIdx].Class)
	})

	wm.Starting(func() {
		for i := 0; i < 10; i++ {
			od := testTvItem{}
			od.Name = fmt.Sprintf("사용자%02d", i)
			od.Level = i + 1
			od.Sex = (i % 2)
			od.Class = fmt.Sprintf("직업%d", i+1)
			tm.items = append(tm.items, od)
		}
		tm.PublishRowsReset()
	})

	wm.Start()
}
