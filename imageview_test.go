package main

import (
	"testing"

	"github.com/pirogom/walk"
	"github.com/pirogom/walkmgr"
)

// 이미지 뷰 관련

func TestImageView(t *testing.T) {
	img := walkmgr.LoadImage(".\\test_data\\gopher.png")

	if img == nil {
		return
	}

	defer func(img *walk.Image) {
		if img != nil {
			(*img).Dispose()
			img = nil
		}
	}(img)

	wm := walkmgr.NewWin("이미지뷰", 640, 480)

	iv := wm.ImageView(walkmgr.IV_ZOOM)
	iv.SetImage(*img)

	wm.Start()
}
