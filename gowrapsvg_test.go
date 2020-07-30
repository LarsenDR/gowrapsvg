package gowrapsvg

import (
	"testing"
)

func TestSVGProlog(t *testing.T) {
	want := "GoWrapSVG a tool to generate SVG 2-D Graphs, By David R. Larsen, 2020"
	if got := SVGProlog(); got != want {
		t.Errorf("SVGProlog() = %q,\n want = %q", got, want)
	}
}

func TestXMLstart(t *testing.T) {
	want := `<?xml version="1.0" encoding="UTF-8" standalone="no"?>`
	if got := XMLStart(); got != want {
		t.Errorf("XMLStart() = %q,\n want = %q", got, want)
	}
}

func TestSVGEnd(t *testing.T) {
	want := `</svg>`
	if got := SVGEnd(); got != want {
		t.Errorf("SVGEnd() = %q,\n want = %q", got, want)
	}
}

func TestSVGStart(t *testing.T) {
	var tests = []struct {
		id   string
		x    int
		y    int
		want string
	}{
		{"t1", 100, 100, `<svg id="t1" viewbox="0 0 1000 1000"  width="100" height="100"  xmlns="http://www.w3.org/2000/svg">`},
		{"t2", 500, 500, `<svg id="t2" viewbox="0 0 1000 1000"  width="500" height="500"  xmlns="http://www.w3.org/2000/svg">`},
		{"t3", 1000, 1000, `<svg id="t3" viewbox="0 0 1000 1000"  width="1000" height="1000"  xmlns="http://www.w3.org/2000/svg">`},
	}

	for _, tt := range tests {
		ans := SVGStart(tt.id, tt.x, tt.y)
		if ans != tt.want {
			t.Errorf("Got: %q \n Want: %q", ans, tt.want)
		}
	}

}
