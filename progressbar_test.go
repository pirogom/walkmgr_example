package main

import (
	"testing"
	"time"

	"github.com/pirogom/walkmgr"
)

// 프로그래스바 테스트
func TestProgressBar(t *testing.T) {
	wm := walkmgr.NewWin("Progressbar 테스트", 640, 480)

	pb := wm.ProgressBar(0, 100, 0)

	wm.Starting(func() {
		go func() {
			for {
				currVal := pb.Value()
				wm.Sync(func() {
					pb.SetValue(currVal + 1)
				})
				time.Sleep(250 * time.Millisecond)

				if currVal == pb.MaxValue() {
					return
				}
			}
		}()
	})

	wm.Start()
}
