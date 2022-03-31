package main

import (
	"fmt"
	"testing"

	"github.com/pirogom/walk"
	"github.com/pirogom/walkmgr"
)

// 메뉴 관련

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
