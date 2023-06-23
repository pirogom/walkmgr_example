package main

import (
	"fmt"
	"testing"

	"github.com/pirogom/walkmgr"
	"github.com/pirogom/win"
)

func TestDialog1(t *testing.T) {
	wm := walkmgr.NewWin("다이얼로그 예제1", 640, 480)

	// 레이아웃 지정 없음 ( 기본 LAYOUT_VERT )
	wm.PushButton("기본", func() {
		dlg := walkmgr.NewDialog(wm.Window(), "기본", 300, 300, nil)

		dlg.Label("라벨1")
		dlg.Label("라벨2")
		dlg.Label("라벨3")

		dlg.StartDLG()
	})

	// 세로 배치
	wm.PushButton("LAYOUT_VERT", func() {
		dlg := walkmgr.NewDialog(wm.Window(), "LAYOUT_VERT", 300, 300, nil, walkmgr.LAYOUT_VERT)

		dlg.Label("라벨1")
		dlg.Label("라벨2")
		dlg.Label("라벨3")

		dlg.StartDLG()
	})

	// 가로 배치
	wm.PushButton("LAYOUT_HORI", func() {
		dlg := walkmgr.NewDialog(wm.Window(), "LAYOUT_VERT", 300, 300, nil, walkmgr.LAYOUT_HORI)

		dlg.Label("라벨1")
		dlg.Label("라벨2")
		dlg.Label("라벨3")

		dlg.StartDLG()
	})

	wm.Start()
}

func TestDialog2(t *testing.T) {
	wm := walkmgr.NewWin("다이얼로그 결과", 640, 480)

	wm.PushButton("팝업", func() {
		dlg := walkmgr.NewDialog(wm.Window(), "다이얼로그 결과", 300, 300, nil)

		dlg.Label("다이얼로그 결과")
		dlg.PushButton("확인", func() {
			dlg.Dlg().Accept()
		})

		dlg.PushButton("취소", func() {
			dlg.Dlg().Cancel()
		})

		dlg.PushButton("100", func() {
			dlg.CloseDLG(100)
		})

		result := dlg.StartDLG()

		switch result {
		case 0:
			walkmgr.MsgBox("Not choice")
		case win.IDOK:
			walkmgr.MsgBox("IDOK")
		case win.IDCANCEL:
			walkmgr.MsgBox("IDCANCEL")
		default:
			walkmgr.MsgBox(fmt.Sprintf("%d", result))
		}
	})

	wm.Start()
}

func TestDialog3(t *testing.T) {
	wm := walkmgr.NewWin("DialogTest", 640, 480)

	wm.PushButton("다이얼로그", func() {

		var canClose = true

		dlg := walkmgr.NewDialog(wm.Window(), "TestDLG", 300, 300, nil)

		dlg.PushButton("창닫기 가능", func() {
			canClose = true
		})
		dlg.PushButton("창닫기 불가", func() {
			canClose = false
		})

		dlg.Starting(func() {
			walkmgr.MsgBox("창 시작 이벤트")
		})

		dlg.Closing(func() bool {

			if canClose {
				walkmgr.MsgBox("창닫기 가능한 상태입니다. 창을 닫습니다.")
			} else {
				walkmgr.MsgBox("창닫기 불가한 상태입니다. 창을 닫을 수 없습니다.")
			}

			return canClose
		})

		dlg.StartDLG()
	})

	wm.Start()
}
