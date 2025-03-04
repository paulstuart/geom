package geom

import (
	"reflect"

	"github.com/ctessum/polyclip-go"
)

// A Path is a series of connected points.
type Path []Point

// Len returns the number of Points in the receiver.
func (p Path) Len() int {
	return len(p)
}

// XY returns the coordinates of point i.
func (p Path) XY(i int) (x, y float64) {
	return p[i].X, p[i].Y
}

// A Polygon is a series of closed rings. The inner rings should be nested
// inside of the outer ring.
type Polygon []Path

// Bounds gives the rectangular extents of the polygon.
func (p Polygon) Bounds() *Bounds {
	b := NewBounds()
	b.extendPointss(p)
	return b
}

// Polygons returns []{p} to fulfill the Polygonal interface.
func (p Polygon) Polygons() []Polygon {
	return []Polygon{p}
}

// Intersection returns the area(s) shared by p and p2.
func (p Polygon) Intersection(p2 Polygonal) Polygonal {
	return p.op(p2, polyclip.INTERSECTION)
}

// Union returns the combination of p and p2.
func (p Polygon) Union(p2 Polygonal) Polygonal {
	return p.op(p2, polyclip.UNION)
}

// XOr returns the area(s) occupied by either p or p2 but not both.
func (p Polygon) XOr(p2 Polygonal) Polygonal {
	return p.op(p2, polyclip.XOR)
}

// Difference subtracts p2 from p.
func (p Polygon) Difference(p2 Polygonal) Polygonal {
	return p.op(p2, polyclip.DIFFERENCE)
}

func (p Polygon) op(p2 Polygonal, op polyclip.Op) Polygon {
	pp := p.toPolyClip()
	var pp2 polyclip.Polygon
	for _, pp2x := range p2.Polygons() {
		pp2 = append(pp2, pp2x.toPolyClip()...)
	}
	return polyClipToPolygon(pp.Construct(op, pp2))
}

func (p Polygon) toPolyClip() polyclip.Polygon {
	o := make(polyclip.Polygon, len(p))
	for i, r := range p {
		o[i] = make(polyclip.Contour, len(r))
		for j, pp := range r {
			o[i][j] = polyclip.Point(pp)
		}
	}
	return o
}

func polyClipToPolygon(p polyclip.Polygon) Polygon {
	pp := make(Polygon, len(p))
	for i, r := range p {
		pp[i] = make([]Point, len(r)+1)
		for j, ppp := range r {
			pp[i][j] = Point(ppp)
		}
		// Close the ring as per OGC standard.
		pp[i][len(r)] = pp[i][0]
	}
	return pp
}

// Len returns the number of points in the receiver.
func (p Polygon) Len() int {
	var i int
	for _, r := range p {
		i += len(r)
	}
	return i
}

// Points returns an iterator for the points in the receiver.
func (p Polygon) Points() func() Point {
	var i, j int
	return func() Point {
		if i == len(p[j]) {
			j++
			i = 0
		}
		i++
		return p[j][i-1]
	}
}

// Within calculates whether p is within poly.
func (p Polygon) Within(poly Polygonal) WithinStatus {
	if reflect.DeepEqual(p, poly) {
		return OnEdge
	}
	for _, r := range p {
		for _, pt := range r {
			if PointInPolygonal(pt, poly) == Outside {
				return Outside
			}
		}
	}
	return Inside
}
