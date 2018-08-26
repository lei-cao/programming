// Copyright 2018 The Algoman Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Draw Tree algorithm: https://llimllib.github.io/pymag-trees/

package shapes

import (
	"github.com/lei-cao/programming/code/algoman/pkg/defaults"
	"github.com/lei-cao/programming/code/v2/algorithms/sorting/basicsort"
	"math"
	"image/color"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

func NewDrawTree(values []int) (root *DrawTree, drawTreeSlice *DrawTreeSlice) {
	heap := new(basicsort.HeapSort)
	drawTreeSlice = new(DrawTreeSlice)
	root = NewDrawTreeNode(nil, 0, 1, -1)
	dts := make([]*DrawTree, len(values))
	if len(values) == 0 {
		return
	}
	for k, v := range values {
		// root
		if k == 0 {
			root = NewDrawTreeNode(nil, 0, 1, v)
			dts[k] = root
		}
		parent := dts[k]
		// left child
		if heap.ILeftChild(k) < len(values) {
			dt := NewDrawTreeNode(parent, parent.y+1, 1, values[heap.ILeftChild(k)])
			dts[heap.ILeftChild(k)] = dt
			parent.children = append(parent.children, dt)
		}
		// right child
		if heap.IRightChild(k) < len(values) {
			dt := NewDrawTreeNode(parent, parent.y+1, 2, values[heap.IRightChild(k)])
			dts[heap.IRightChild(k)] = dt
			parent.children = append(parent.children, dt)
		}
	}
	drawTreeSlice.dts = dts
	drawTreeSlice.tree = root
	depth := math.Floor(math.Log2(float64(len(values))))
	drawTreeSlice.rs = NewRectSlice(values, 0, float64(depth+1)*rh)
	root = buchheim(root)
	return
}

func NewDrawTreeNode(parent *DrawTree, depth int, number int, value int) *DrawTree {
	t := new(DrawTree)
	t.V = value
	t.Image = NewCircle(value)
	t.x = -1
	t.y = depth
	t.tree = t // ?
	t.children = []*DrawTree{}
	t.parent = parent
	t.ancestor = t
	t.change = 0
	t.shift = 0
	t.number = number

	t.color = defaults.BarColor

	return t
}

type DrawTree struct {
	Image           *Circle
	V               int
	x               int
	y               int
	tree            *DrawTree
	children        []*DrawTree
	parent          *DrawTree
	thread          *DrawTree
	offset          int
	ancestor        *DrawTree
	change          int
	shift           int
	leftMostSibling *DrawTree
	number          int
	mod             int

	DestX int
	DestY int
	color color.Color
}

var (
	rw    = float64(defaults.CircleRadius) * 2.8
	rh    = float64(defaults.CircleRadius) * 2.8
	halfR = float64(defaults.CircleRadius)
)

func (t *DrawTree) Update(progress float64) {
	dx, dy := progress*float64(t.DestX-t.x), progress*float64(t.DestY-t.y)
	t.x += int(dx)
	t.y += int(dy)
	if progress == 1 {
		t.x = t.DestX
		t.y = t.DestY
	}
}

func (t *DrawTree) Draw(img *ebiten.Image, depth int) {
	if t.Image == nil {
		return
	}
	offsetX := float64(t.x) * rw
	offsetY := float64(depth) * rh

	for _, c := range t.children {
		offsetX1 := float64(c.x) * rw
		offsetY1 := float64(depth+1) * rh

		ebitenutil.DrawLine(img, offsetX+halfR, offsetY+halfR, offsetX1+halfR , offsetY1+halfR, defaults.ColorC)

		c.Draw(img, depth+1)
	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(offsetX, offsetY)
	//op.GeoM.Scale(0.8, 0.8)
	img.DrawImage(t.Image.Img, op)

}

func (t *DrawTree) Depth() int {
	return t.y
}

func (t *DrawTree) left() *DrawTree {
	if t.thread != nil {
		return t.thread
	}
	if len(t.children) > 0 {
		return t.children[0]
	}
	return nil
}

func (t *DrawTree) right() *DrawTree {
	if t.thread != nil {
		return t.thread
	}
	if len(t.children) > 0 {
		return t.children[len(t.children)-1]
	}
	return nil
}

func (t *DrawTree) leftBrother() *DrawTree {
	var n *DrawTree = nil
	if t.parent != nil {
		for _, v := range t.parent.children {
			if v == t {
				return n
			} else {
				n = v
			}
		}
	}
	return n
}

func (t *DrawTree) getLeftMostSibling() *DrawTree {
	if t.leftMostSibling == nil && t.parent != nil && t != t.parent.children[0] {
		t.leftMostSibling = t.parent.children[0]
	}
	return t.leftMostSibling
}

func buchheim(tree *DrawTree) *DrawTree {
	dt := firstWalk(tree, 1.0)
	min := secondWalk(dt, 0, 0, -1)
	if min < 0 {
		thirdWalk(dt, -min)
	}

	return dt
}

func thirdWalk(tree *DrawTree, n int) {
	tree.x += n
	for _, c := range tree.children {
		thirdWalk(c, n)
	}
}

func firstWalk(v *DrawTree, distance int) *DrawTree {
	if len(v.children) == 0 {
		if v.getLeftMostSibling() != nil {
			v.x = v.leftBrother().x + distance
		} else {
			v.x = 0
		}
	} else {
		defaultAncestor := v.children[0]
		for _, w := range v.children {
			firstWalk(w, 1)
			defaultAncestor = apportion(w, defaultAncestor, distance)
		}
		executeShifts(v)
		midPoint := (v.children[0].x + v.children[len(v.children)-1].x) / 2

		w := v.leftBrother()
		if w != nil {
			v.x = w.x + distance
			v.mod = v.x - midPoint
		} else {
			v.x = midPoint
		}
	}
	return v
}

func apportion(v *DrawTree, defaultAncestor *DrawTree, distance int) *DrawTree {
	w := v.leftBrother()
	if w != nil {
		vir := v
		vor := v
		vil := w
		vol := v.getLeftMostSibling()
		sir := v.mod
		sor := v.mod
		sil := vil.mod
		sol := vol.mod
		for vil.right() != nil && vir.left() != nil {
			vil = vil.right()
			vir = vir.left()
			vol = vol.left()
			vor = vor.right()
			vor.ancestor = v
			shift := (vil.x + sil) - (vir.x + sir) + distance
			if shift > 0 {
				a := ancestor(vil, v, defaultAncestor)
				moveSubtree(a, v, shift)
				sir = sir + shift
				sor = sor + shift
			}
			sil += vil.mod
			sir += vir.mod
			sol += vol.mod
			sor += vor.mod
		}
		if vil.right() != nil && vor.right() == nil {
			vor.thread = vil.right()
			vor.mod += sil - sor
		} else {
			if vir.left() != nil && vol.left() == nil {
				vol.thread = vir.left()
				vol.mod += sir - sol
			}
			defaultAncestor = v
		}
	}
	return defaultAncestor
}

func moveSubtree(wl *DrawTree, wr *DrawTree, shift int) {
	subtrees := wr.number - wl.number
	wr.change -= shift / subtrees
	wr.shift += shift
	wl.change += shift / subtrees
	wr.x += shift
	wr.mod += shift
}

func executeShifts(v *DrawTree) {
	change := 0
	shift := 0
	for i := len(v.children) - 1; i >= 0; i-- {
		w := v.children[i]
		w.x += shift
		w.mod += shift
		change += w.change
		shift += w.shift + change

	}
}

func ancestor(vil *DrawTree, v *DrawTree, defaultAncestor *DrawTree) *DrawTree {
	for _, c := range v.parent.children {
		if vil.ancestor == c {
			return vil.ancestor
		}
	}
	return defaultAncestor
}

func secondWalk(v *DrawTree, m int, depth int, min int) int {
	v.x += m
	v.y = depth

	if min == -1 || v.x < min {
		min = v.x
	}
	for _, w := range v.children {
		min = secondWalk(w, m+v.mod, depth+1, min)
	}
	return min
}
