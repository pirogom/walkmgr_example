package main

import (
	"testing"

	"github.com/pirogom/walkmgr"
)

func TestLinkLabel(t *testing.T) {
	wm := walkmgr.NewWin("Link Label", 640, 480)

	wm.LinkLabel(`<a href="https://modu-print.tistory.com">피로곰's 모두의 프린터</a>`, walkmgr.ALIGN_LEFT)
	wm.LinkLabel(`<a href="https://modu-print.tistory.com/category/%EB%AA%A8%EB%91%90%EC%9D%98%20PDF">피로곰's 모두의 PDF</a>`, walkmgr.ALIGN_LEFT)

	wm.Start()
}

func TestLinkLabel2(t *testing.T) {
	wm := walkmgr.NewWin("Link Label", 640, 480)

	wm.LinkLabel(`모두의 프린터를 다운로드 받으시려면 <a href="https://modu-print.tistory.com">여기</a>를 클릭하십시요.`)
	wm.LinkLabel(`피로곰이 만든 PDF편집기 '모두의 PDF'를 다운로드 받으시려면 <a href="https://modu-print.tistory.com/category/%EB%AA%A8%EB%91%90%EC%9D%98%20PDF">여기</a>를 클릭하십시요.`)

	wm.LinkLabel(`모두의 프린터를 다운로드 받으시려면 <a href="https://modu-print.tistory.com">여기</a>를 클릭하십시요.`, walkmgr.ALIGN_LEFT)
	wm.LinkLabel(`피로곰이 만든 PDF편집기 '모두의 PDF'를 다운로드 받으시려면 <a href="https://modu-print.tistory.com/category/%EB%AA%A8%EB%91%90%EC%9D%98%20PDF">여기</a>를 클릭하십시요.`, walkmgr.ALIGN_LEFT)

	wm.LinkLabel(`모두의 프린터를 다운로드 받으시려면 <a href="https://modu-print.tistory.com">여기</a>를 클릭하십시요.`, walkmgr.ALIGN_RIGHT)
	wm.LinkLabel(`피로곰이 만든 PDF편집기 '모두의 PDF'를 다운로드 받으시려면 <a href="https://modu-print.tistory.com/category/%EB%AA%A8%EB%91%90%EC%9D%98%20PDF">여기</a>를 클릭하십시요.`, walkmgr.ALIGN_RIGHT)

	wm.Start()
}

func TestLinkLabel3(t *testing.T) {
	wm := walkmgr.NewWin("Link Label", 640, 480)

	wm.LinkLabel(`모두의 프린터를 다운로드 받으시려면 <a href="https://modu-print.tistory.com">여기</a>를 클릭하십시요.`)
	wm.LinkLabel(`피로곰이 만든 PDF편집기 '모두의 PDF'를 다운로드 받으시려면 <a href="https://modu-print.tistory.com/category/%EB%AA%A8%EB%91%90%EC%9D%98%20PDF">여기</a>를 클릭하십시요.`)

	wm.LinkLabel(`모두의 프린터를 다운로드 받으시려면 <a href="https://modu-print.tistory.com">여기</a>를 클릭하십시요.`, walkmgr.ALIGN_LEFT)
	wm.LinkLabel(`피로곰이 만든 PDF편집기 '모두의 PDF'를 다운로드 받으시려면 <a href="https://modu-print.tistory.com/category/%EB%AA%A8%EB%91%90%EC%9D%98%20PDF">여기</a>를 클릭하십시요.`, walkmgr.ALIGN_LEFT)

	wm.LinkLabel(`모두의 프린터를 다운로드 받으시려면 <a href="https://modu-print.tistory.com">여기</a>를 클릭하십시요.`, walkmgr.ALIGN_RIGHT)
	wm.LinkLabel(`피로곰이 만든 PDF편집기 '모두의 PDF'를 다운로드 받으시려면 <a href="https://modu-print.tistory.com/category/%EB%AA%A8%EB%91%90%EC%9D%98%20PDF">여기</a>를 클릭하십시요.`, walkmgr.ALIGN_RIGHT)

	wm.VSpacer()

	wm.Start()
}
