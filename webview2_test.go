package main

import (
	"os"
	"testing"

	"github.com/pirogom/walkmgr"
)

func TestWebView2Example(t *testing.T) {
	wm := walkmgr.NewWin("WebView2 Test1", 800, 600)

	var wv *walkmgr.WebView2
	wv = wm.WebView2(os.TempDir(), func() {
		wv.Navigate("https://modu-print.com")
	}, func() {
		walkmgr.MsgBox("WebView2 런타임을 찾을 수 없습니다.")
	})

	wm.StartForeground()
}

func TestWebView2Example2(t *testing.T) {
	wm := walkmgr.NewWin("WebView2 Test2", 800, 600)

	wm.VGroupBox("웹뷰 그룹1")

	var wv1 *walkmgr.WebView2
	wv1 = wm.WebView2(os.TempDir(), func() {
		wv1.Navigate("https://modu-print.com")
	}, func() {
		walkmgr.MsgBox("WebView2 런타임을 찾을 수 없습니다.")
	})

	var wv2 *walkmgr.WebView2
	wv2 = wm.WebView2(os.TempDir(), func() {
		wv2.Navigate("https://teus.tistory.com")
	}, func() {
		walkmgr.MsgBox("WebView2 런타임을 찾을 수 없습니다.")
	})

	wm.End()

	wm.HGroupBox("웹뷰 그룹2")
	var wv3 *walkmgr.WebView2
	wv3 = wm.WebView2(os.TempDir(), func() {
		wv3.Navigate("https://google.com")
	}, func() {
		walkmgr.MsgBox("WebView2 런타임을 찾을 수 없습니다.")
	})

	var wv4 *walkmgr.WebView2
	wv4 = wm.WebView2(os.TempDir(), func() {
		wv4.Navigate("https://naver.com")
	}, func() {
		walkmgr.MsgBox("WebView2 런타임을 찾을 수 없습니다.")
	})

	wm.End()

	wm.StartForeground()
}
