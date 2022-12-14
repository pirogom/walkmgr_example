package main

import (
	"fmt"
	"testing"

	"github.com/pirogom/walk"
	"github.com/pirogom/walkmgr"
)

// 메뉴 관련

func TestMenu1(t *testing.T) {
	wm := walkmgr.NewWin("메뉴 테스트1", 640, 480)

	//
	m1 := walkmgr.NewMenu("메인메뉴1")
	m1.AddAction("메뉴1", func() { fmt.Println("액숀1") })
	m1.AddAction("메뉴2", func() { fmt.Println("액숀2") })
	m1.AddAction("메뉴3", func() { fmt.Println("액숀3") })
	m1.AddAction("메뉴4", func() { fmt.Println("액숀4") })

	wm.AddMenu(m1)

	wm.Start()
}

func TestMenu2(t *testing.T) {
	wm := walkmgr.NewWin("메뉴 테스트1", 640, 480)

	//
	m1 := walkmgr.NewMenu("메인메뉴1")
	m1.AddAction("메뉴1", func() { fmt.Println("액숀1") })
	m1.AddAction("메뉴2", func() { fmt.Println("액숀2") })
	m1.AddAction("메뉴3", func() { fmt.Println("액숀3") })
	m1.AddAction("메뉴4", func() { fmt.Println("액숀4") })

	wm.AddMenu(m1)

	m2 := walkmgr.NewMenu("메인메뉴1")
	m2.AddAction("메뉴1", func() { fmt.Println("액숀1") })
	m2.AddAction("메뉴2", func() { fmt.Println("액숀2") })
	m2.AddAction("메뉴3", func() { fmt.Println("액숀3") })
	m2.AddAction("메뉴4", func() { fmt.Println("액숀4") })

	wm.AddMenu(m2)

	wm.Start()
}

func TestMenu3(t *testing.T) {
	wm := walkmgr.NewWin("메뉴 테스트1", 640, 480)

	//
	m1 := walkmgr.NewMenu("메인메뉴1")
	m1.AddAction("메뉴1", func() { fmt.Println("액숀1") })
	m1.AddAction("메뉴2", func() { fmt.Println("액숀2") })
	m1.AddSeparator()
	m1.AddAction("메뉴3", func() { fmt.Println("액숀3") })
	m1.AddAction("메뉴4", func() { fmt.Println("액숀4") })

	wm.AddMenu(m1)

	wm.Start()
}

func TestMenu4(t *testing.T) {
	wm := walkmgr.NewWin("메뉴 테스트1", 640, 480)

	//
	m1 := walkmgr.NewMenu("메인메뉴1")
	m1.AddAction("메뉴1", func() { fmt.Println("액숀1") })
	m1.AddAction("메뉴2", func() { fmt.Println("액숀2") })
	m1.AddSeparator()
	m1.AddAction("메뉴3", func() { fmt.Println("액숀3") })
	m1.AddAction("메뉴4", func() { fmt.Println("액숀4") })
	m1.AddSeparator()
	c1 := m1.AddCheck("체크메뉴1", false, func() {})
	c2 := m1.AddCheck("체크메뉴2", true, func() {})

	m1.AddAction("체크토글", func() {
		if c1.Checked() {
			c1.SetChecked(false)
		} else {
			c1.SetChecked(true)
		}
		if c2.Checked() {
			c2.SetChecked(false)
		} else {
			c2.SetChecked(true)
		}
	})

	wm.AddMenu(m1)

	wm.Start()
}

func TestMenu5(t *testing.T) {
	wm := walkmgr.NewWin("메뉴 테스트1", 640, 480)

	//
	m1 := walkmgr.NewMenu("메인메뉴1")

	s1 := walkmgr.NewMenu("과일")
	s1.AddAction("사과", func() {})
	s1.AddAction("포토", func() {})
	s1.AddAction("배", func() {})
	s1.AddAction("바나나", func() {})
	m1.AddMenu(s1)

	s2 := walkmgr.NewMenu("한식")
	s2.AddAction("된장찌개", func() {})
	s2.AddAction("김치찌개", func() {})
	s2.AddAction("순두부", func() {})
	s2.AddAction("고추장찌개", func() {})
	s2.AddSeparator()
	s2.AddAction("갈비찜", func() {})
	s2.AddAction("불고기", func() {})
	s2.AddAction("곱창전골", func() {})
	m1.AddMenu(s2)

	s3 := walkmgr.NewMenu("양식")
	ss1 := walkmgr.NewMenu("유럽")
	sss1 := walkmgr.NewMenu("이탈리아")
	sss1.AddAction("파스타", func() {})
	sss1.AddAction("피자", func() {})
	ss1.AddMenu(sss1)
	s3.AddMenu(ss1)
	m1.AddMenu(s3)

	wm.AddMenu(m1)

	wm.Start()
}

func TestMenuMgr(t *testing.T) {
	wm := walkmgr.NewWin("메뉴 테스트", 640, 480)

	//
	m1 := walkmgr.NewMenu("메뉴1")
	m1.AddAction("액숀1", func() { fmt.Println("액숀1") })
	m1.AddAction("액숀2", func() { fmt.Println("액숀2") })
	m1.AddAction("액숀3", func() { fmt.Println("액숀3") })
	m1.AddAction("액숀4", func() { fmt.Println("액숀4") })
	m1.AddSeparator()
	var ca1, ca2 *walk.Action
	ca1 = m1.AddCheck("체크액션1", true, func() {
		fmt.Println("체크", ca1.Checked())
	})
	ca2 = m1.AddCheck("체크액션2", false, func() {
		fmt.Println("체크", ca2.Checked())
	})
	m1.AddSeparator()
	m1.AddAction("액숀5", func() { fmt.Println("액숀5") })
	m1.AddAction("액숀6", func() { fmt.Println("액숀6") })

	s1 := walkmgr.NewMenu("서브메뉴1")
	s1.AddAction("서브액숀1", func() {})
	s1.AddAction("서브액숀2", func() {})
	s1.AddAction("서브액숀3", func() {})
	s1.AddAction("서브액숀4", func() {})
	s1.AddAction("서브액숀5", func() {})
	s1.AddAction("서브액숀6", func() {})
	s1.AddAction("서브액숀7", func() {})

	ss1 := walkmgr.NewMenu("서브서브메뉴1")
	ss1.AddAction("서브서브액숀1", func() {})
	ss1.AddAction("서브서브액숀2", func() {})
	ss1.AddAction("서브서브액숀3", func() {})
	ss1.AddAction("서브서브액숀4", func() {})
	ss1.AddAction("서브서브액숀5", func() {})
	ss1.AddAction("서브서브액숀6", func() {})
	s1.AddMenu(ss1)

	m1.AddMenu(s1)
	//

	wm.AddMenu(m1)

	wm.Start()
}
