package main

import (
	"testing"

	"github.com/pirogom/walkmgr"
)

// 고급(?) 레이아웃 처리?? 복합 레이아웃 구성

func TestComposite(t *testing.T) {
	wm := walkmgr.NewWin("복합 레이아웃 구성1", 640, 480, walkmgr.LAYOUT_VERT)

	wm.HComposite()
	wm.Label("이름")
	wm.LineEdit()
	wm.Label("나이")
	wm.LineEdit()
	wm.End()

	wm.VComposite()
	wm.HComposite()
	wm.Label("성별")
	wm.LineEdit()
	wm.LabelCenter("직업")
	wm.LineEdit()
	wm.End()
	wm.End()

	wm.Start()
}

func TestComposite2(t *testing.T) {
	wm := walkmgr.NewWin("복합 레이아웃 구성2", 640, 480, walkmgr.LAYOUT_VERT)

	wm.HComposite()
	wm.Label("이름")
	wm.LineEdit()
	wm.Label("나이")
	wm.LineEdit()
	wm.End()

	wm.HComposite()
	wm.Label("성별")
	wm.LineEdit()
	wm.LabelCenter("직업")
	wm.LineEdit()
	wm.End()

	wm.VSpacer()

	wm.Start()
}

func TestComposite3(t *testing.T) {
	wm := walkmgr.NewWin("복합 레이아웃 구성3", 640, 480, walkmgr.LAYOUT_HORI)

	wm.VComposite()
	wm.Label("이름")
	wm.LineEdit()
	wm.Label("나이")
	wm.LineEdit()
	wm.Label("종족")
	wm.LineEdit()
	wm.End()

	wm.VComposite()
	wm.Label("성별")
	wm.LineEdit()
	wm.LabelCenter("직업")
	wm.LineEdit()
	wm.End()

	wm.Start()
}
