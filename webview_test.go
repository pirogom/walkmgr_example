package main

import (
	"testing"

	"github.com/pirogom/walkmgr"
)

// 웹뷰 관련

func TestWebView(t *testing.T) {
	wm := walkmgr.NewWin("웹뷰(Alert Disabled)", 640, 480)

	wm.WebView("https://modu-print.tistory.com")

	wm.Start()
}

func TestWebViewWithAlert(t *testing.T) {
	wm := walkmgr.NewWin("웹뷰(Alert Enabled)", 640, 480)

	wm.WebViewWithAlert("https://modu-print.tistory.com")

	wm.Start()
}
