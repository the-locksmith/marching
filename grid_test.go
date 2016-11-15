package marching

import (
	"bytes"
	"image/png"
	"io/ioutil"
	"testing"
)

var (
	testAValues = []float64{
		1, 1, 1, 1, 1,
		1, 2, 3, 2, 1,
		1, 3, 4, 3, 1,
		1, 2, 3, 2, 1,
		1, 1, 1, 1, 1,
	}
	testAWidth          = 5
	testAHeight         = 5
	testALevel  float64 = 2
	testACases          = []Case{
		13, 12, 12, 14,
		9, 0, 0, 6,
		9, 0, 0, 6,
		11, 3, 3, 7,
	}
)

func TestGrid(t *testing.T) {
	grid := NewGrid(testAValues, testAWidth, testAHeight, testALevel)
	if len(grid.Cells) != len(testACases) {
		t.Fatalf("expected %v, got %v", len(testACases), len(grid.Cells))
	}
	if false {
		for i := 0; i < len(grid.Cells); i++ {
			if testACases[i] != grid.Cells[i].Case {
				t.Fatalf("expected %v, got %v for #%d", testACases[i], grid.Cells[i].Case, i)

			}
		}
	}
	img := grid.Image(1000, 1000, &ImageOptions{Marks: false})
	var buf bytes.Buffer
	if err := png.Encode(&buf, img); err != nil {
		t.Fatal(err)
	}
	ioutil.WriteFile("testgrid.png", buf.Bytes(), 0600)
}
