package topojson

import (
	"testing"

	"github.com/paulmach/orb"
	geojson "github.com/paulmach/orb/geojson"
)

func TestSimplify(t *testing.T) {
	poly := geojson.NewFeature(orb.Polygon{
		orb.Ring{
			{0, 0}, {0, 1}, {0.5, 1.1}, {1, 1}, {1, 0}, {0, 0},
		},
	})
	poly.ID = "poly"

	fc := geojson.NewFeatureCollection()
	fc.Append(poly)
	t.Run("Reducing a triangle, 0.0501", func(tt *testing.T) {
		tt.Parallel()
		topo := NewTopology(fc, &TopologyOptions{
			Simplify: 0.0501, // 1 x 0.1 / 2 = 0.05
		})
		if len(topo.Arcs[0]) != 5 {
			tt.Error("failed to simplify, Arcs must be [0 0] [0 1] [1 1] [1 0] [0 0], actual:", topo.Arcs[0])
		}
	})
	t.Run("Reducing a triangle, 0.0500", func(tt *testing.T) {
		tt.Parallel()
		topo := NewTopology(fc, &TopologyOptions{
			Simplify: 0.0500, // 1 x 0.1 / 2 = 0.05
		})
		if len(topo.Arcs[0]) != 6 {
			tt.Error("failed to simplify, Arcs must be [0 0] [0 1] [0.5 1.1] [1 1] [1 0] [0 0], actual:", topo.Arcs[0])
		}
	})
}
