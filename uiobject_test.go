package main

import (
	"testing"

	"github.com/pirogom/walkmgr"
)

// 버튼, 체크박스등 ui 요소들 관련

func TestUI(t *testing.T) {
	wm := walkmgr.NewWin("UI테스트", 640, 1080)

	// 버튼
	wm.PushButton("버튼", func() {

	})

	// 체크박스
	wm.CheckBox("체크박스(체크)", true, func() {

	})
	wm.CheckBox("체크박스(체크X)", false, func() {

	})

	// 드롭다운
	wm.DropDownBox([]string{"하나", "둘", "셋", "넷"})
	wm.DropDownBox([]string{"다섯", "여섯", "일곱", "여덟"}, 2)

	// 에디트 박스
	wm.NumberEdit()
	wm.NumberEdit(100)

	wm.LineEdit()
	wm.LineEdit("LineEdit 수정가능")

	wm.LineStatic()
	wm.LineStatic("LineStatic 수정불가")

	wm.TextEdit()
	wm.TextEdit("TextEdit 수정가능")

	wm.TextStatic()
	wm.TextStatic("TextStatic 수정불가")

	wm.TextArea()
	wm.TextArea("TextArea 수정가능")

	wm.TextAreaStatic()
	wm.TextAreaStatic("TextAreaStatic 수정불가")

	wm.Label("라벨 - 기본")
	wm.Label("라벨 - LEFT", walkmgr.ALIGN_LEFT)
	wm.Label("라벨 - CENTER", walkmgr.ALIGN_CENTER)
	wm.Label("라벨 - RIGHT", walkmgr.ALIGN_RIGHT)

	wm.Slider(0, 100, 50)

	wm.Start()
}
