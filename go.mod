module github.com/paulstuart/geom

require (
	github.com/ctessum/geom v0.2.12
	github.com/ctessum/polyclip-go v1.1.0
	github.com/jonas-p/go-shp v0.1.2-0.20190401125246-9fd306ae10a6
	github.com/paulmach/osm v0.1.1
	golang.org/x/sync v0.0.0-20200625203802-6e8e738ad208
	gonum.org/v1/gonum v0.9.3
	gonum.org/v1/plot v0.9.0
)

go 1.13

replace github.com/ctessum/geom => ./

replace github.com/ctessum/geom/op => ./op

replace github.com/ctessum/geom/proj => ./proj

replace github.com/ctessum/geom/encoding/geojson => ./encoding/geojson

replace github.com/ctessum/geom/encoding/hex => ./encoding/hex

replace github.com/ctessum/geom/encoding/wkb => ./encoding/wkb

replace github.com/ctessum/geom/index/rtree => ./index/rtree
