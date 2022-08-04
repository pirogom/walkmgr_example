package main

import (
	"fmt"
	"testing"

	"github.com/pirogom/walkmgr"
)

// 라디오 버튼 관련

func TestRadio(t *testing.T) {
	wm := walkmgr.NewWin("RADIO BUTTON", 640, 480)

	wm.HGroupBox("라디오 버튼 그룹")
	radio1 := wm.RadioButton("라디오1", int(10))
	radio2 := wm.RadioButton("라디오2", int(20))
	radio3 := wm.RadioButton("라디오3", int(30))
	wm.End()

	_ = radio1
	_ = radio2
	_ = radio3

	wm.PushButton("확인", func() {
		checkedBtn := radio1.Group().CheckedButton()

		if checkedBtn != nil {
			fmt.Print("체크된 버튼: ")
			fmt.Print(checkedBtn.Text())
			fmt.Print(" 값:")
			fmt.Println(checkedBtn.Value())
		}
	})

	wm.Start()
}
