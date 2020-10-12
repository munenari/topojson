package topojson

import (
	"github.com/paulmach/orb"
	"github.com/paulmach/orb/simplify"
)

func (t *Topology) simplify() {
	t.deletedArcs = make(map[int]bool)
	t.shiftArcs = make(map[int]int)

	if t.opts.Simplify == 0 {
		for i := range t.Arcs {
			t.deletedArcs[i] = false
			t.shiftArcs[i] = 0
		}
		return
	}

	newArcs := make([][][]float64, 0)
	s := simplify.VisvalingamThreshold(t.opts.Simplify)
	for i, arc := range t.Arcs {
		ls := make([]orb.Point, len(arc))
		for i, a := range arc {
			ls[i] = orb.Point{a[1], a[0]}
		}
		ls = s.LineString(ls)
		newArc := make([][]float64, len(ls))
		for j, p := range ls {
			newArc[j] = []float64{p[1], p[0]}
		}

		if i == 0 {
			t.shiftArcs[i] = 0
		} else {
			t.shiftArcs[i] = t.shiftArcs[i-1]
		}

		remove := len(newArc) <= 2 && pointEquals(newArc[0], newArc[1])
		if remove {
			// Zero-length arc, remove it!
			t.deletedArcs[i] = true
			t.shiftArcs[i]++
		} else {
			newArcs = append(newArcs, newArc)
		}
	}
	t.Arcs = newArcs
}
