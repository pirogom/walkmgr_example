package main

import (
	"testing"

	"github.com/pirogom/walkmgr"
)

func TestDefToolBar(t *testing.T) {
	wm := walkmgr.NewWin("툴바 테스트(기본)", 640, 480)

	wm.NewToolBar()
	wm.AddTool("김", "icon\\icon1.ico", func() {})
	wm.AddTool("수한무", "icon\\icon2.ico", func() {})
	wm.ToolBarEnd()

	wm.Start()
}

func TestTextToolBar(t *testing.T) {
	wm := walkmgr.NewWin("툴바 테스트(기본 텍스트)", 640, 480)

	wm.NewToolBarWithText()
	wm.AddTool("김", "icon\\icon1.ico", func() {})
	wm.AddTool("수한무", "icon\\icon2.ico", func() {})
	wm.ToolBarEnd()

	wm.Start()
}

func TestTextToolBar2(t *testing.T) {
	wm := walkmgr.NewWin("툴바 테스트(하단 텍스트)", 640, 480)

	wm.NewToolBarWithText2()
	wm.AddTool("김", "icon\\icon1.ico", func() {})
	wm.AddTool("수한무", "icon\\icon2.ico", func() {})
	wm.ToolBarEnd()

	wm.Start()
}

func TestToolbarWithLine(t *testing.T) {
	wm := walkmgr.NewWin("툴바 테스트(구분선)", 640, 480)

	wm.NewToolBar()
	wm.AddTool("김", "icon\\icon1.ico", func() {})
	wm.AddTool("수한무", "icon\\icon2.ico", func() {})
	wm.AddToolSeparator()
	wm.AddTool("거북이", "icon\\icon1.ico", func() {})
	wm.AddTool("두루미", "icon\\icon2.ico", func() {})
	wm.AddToolSeparator()
	wm.AddTool("삼천갑자", "icon\\icon1.ico", func() {})
	wm.AddTool("동방삭", "icon\\icon2.ico", func() {})
	wm.AddTool("치치카포", "icon\\icon1.ico", func() {})
	wm.AddToolSeparator()
	wm.AddTool("사리사리", "icon\\icon2.ico", func() {})
	wm.AddTool("센타", "icon\\icon1.ico", func() {})
	wm.AddTool("워리워리", "icon\\icon2.ico", func() {})
	wm.AddToolSeparator()
	wm.AddTool("세브리깡", "icon\\icon1.ico", func() {})
	wm.AddTool("무두셀라", "icon\\icon2.ico", func() {})
	wm.ToolBarEnd()

	wm.Start()
}

func TestToolbarWithLineAndText(t *testing.T) {
	wm := walkmgr.NewWin("툴바 테스트(구분선+텍스트)", 640, 480)

	wm.NewToolBarWithText()
	wm.AddTool("김", "icon\\icon1.ico", func() {})
	wm.AddTool("수한무", "icon\\icon2.ico", func() {})
	wm.AddToolSeparator()
	wm.AddTool("거북이", "icon\\icon1.ico", func() {})
	wm.AddTool("두루미", "icon\\icon2.ico", func() {})
	wm.AddToolSeparator()
	wm.AddTool("삼천갑자", "icon\\icon1.ico", func() {})
	wm.AddTool("동방삭", "icon\\icon2.ico", func() {})
	wm.AddTool("치치카포", "icon\\icon1.ico", func() {})
	wm.AddToolSeparator()
	wm.AddTool("사리사리", "icon\\icon2.ico", func() {})
	wm.AddTool("센타", "icon\\icon1.ico", func() {})
	wm.AddTool("워리워리", "icon\\icon2.ico", func() {})
	wm.AddToolSeparator()
	wm.AddTool("세브리깡", "icon\\icon1.ico", func() {})
	wm.AddTool("무두셀라", "icon\\icon2.ico", func() {})
	wm.ToolBarEnd()

	wm.Start()
}

func TestToolbarWithLineAndText2(t *testing.T) {
	wm := walkmgr.NewWin("툴바 테스트(구분선+텍스트)", 640, 480)

	wm.NewToolBarWithText2()
	wm.AddTool("김", "icon\\icon1.ico", func() {})
	wm.AddTool("수한무", "icon\\icon2.ico", func() {})
	wm.AddToolSeparator()
	wm.AddTool("거북이", "icon\\icon1.ico", func() {})
	wm.AddTool("두루미", "icon\\icon2.ico", func() {})
	wm.AddToolSeparator()
	wm.AddTool("삼천갑자", "icon\\icon1.ico", func() {})
	wm.AddTool("동방삭", "icon\\icon2.ico", func() {})
	wm.AddTool("치치카포", "icon\\icon1.ico", func() {})
	wm.AddToolSeparator()
	wm.AddTool("사리사리", "icon\\icon2.ico", func() {})
	wm.AddTool("센타", "icon\\icon1.ico", func() {})
	wm.AddTool("워리워리", "icon\\icon2.ico", func() {})
	wm.AddToolSeparator()
	wm.AddTool("세브리깡", "icon\\icon1.ico", func() {})
	wm.AddTool("무두셀라", "icon\\icon2.ico", func() {})
	wm.ToolBarEnd()

	wm.Start()
}
