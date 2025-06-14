package bruteforce

import (
	"geoindexing_comparison/pkg/generator"
	"testing"

	"github.com/teadove/teasutils/utils/test_utils"

	"github.com/stretchr/testify/assert"
)

func TestUnit_PointGenerator_GeneratePoint_Ok(t *testing.T) {
	t.Parallel()

	points := generator.NewSimplerGenerator().Points(&generator.DefaultInput, 1000)
	collection := New()
	collection.FromArray(points)

	origin := points.GetRandomPoint()
	items, _ := collection.KNNTimed(origin, 900)
	assert.Len(t, items, 900)
}

func TestUnit_PointGenerator_RadiusSearch_Ok(t *testing.T) {
	t.Parallel()

	points := generator.NewSimplerGenerator().Points(&generator.DefaultInput, 1000)
	collection := New()
	collection.FromArray(points)

	points, _ = collection.BBoxTimed(points.GetRandomPoint(), points.GetRandomPoint())
	test_utils.Pprint(points)
}
