package main

import (
	"log"
	"testing"

	"github.com/pirogom/walk"
	"github.com/pirogom/walkmgr"
)

var icon1 *walk.Icon

type SampleTreeItem struct {
	name     string
	parent   *SampleTreeItem
	children []*SampleTreeItem
}

func NewSampleTreeItem(name string, parent *SampleTreeItem) *SampleTreeItem {
	nd := SampleTreeItem{}
	nd.name = name
	nd.parent = parent
	return &nd
}

func (d *SampleTreeItem) Text() string {
	return d.name
}

func (d *SampleTreeItem) Parent() walk.TreeItem {
	if d.parent == nil {
		return nil
	}

	return d.parent
}

func (d *SampleTreeItem) ChildCount() int {
	if d.children == nil {
		if err := d.ResetChildren(); err != nil {
			log.Print(err)
		}
	}

	return len(d.children)
}

func (d *SampleTreeItem) ChildAt(index int) walk.TreeItem {
	return d.children[index]
}

func (d *SampleTreeItem) Image() interface{} {
	return icon1
	//return nil
}

func (d *SampleTreeItem) ResetChildren() error {
	d.children = nil
	return nil
}

type SampleTreeModel struct {
	walk.TreeModelBase
	roots []*SampleTreeItem
}

func (*SampleTreeModel) LazyPopulation() bool {
	return true
}

func (m *SampleTreeModel) RootCount() int {
	return len(m.roots)
}

func (m *SampleTreeModel) RootAt(index int) walk.TreeItem {
	return m.roots[index]
}

func TestTreeView(t *testing.T) {
	icon1, _ = walk.NewIconFromFile("./icon/icon1.ico")

	wm := walkmgr.NewWin("트리뷰 예제", 640, 480)

	tv, _ := walk.NewTreeView(wm.Parent())

	tv.CurrentItemChanged().Attach(func() {

	})

	tvm := new(SampleTreeModel)

	tvi := new(SampleTreeItem)
	tvi.name = "Haha"

	tvi2 := new(SampleTreeItem)
	tvi2.name = "hoho"
	tvi2.parent = tvi
	tvi.children = append(tvi.children, tvi2)

	tvi3 := new(SampleTreeItem)
	tvi3.name = "hehe"
	tvi3.parent = tvi2
	tvi2.children = append(tvi2.children, tvi3)

	tvi4 := new(SampleTreeItem)
	tvi4.name = "hihi"
	tvi4.parent = tvi2
	tvi2.children = append(tvi2.children, tvi4)

	tvm.roots = append(tvm.roots, tvi)

	tv.SetModel(tvm)
	tv.SetExpanded(tvi, true)
	tv.SetExpanded(tvi2, true)

	wm.Append(tv)
	wm.Start()
}
