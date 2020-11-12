package leetcode

import (
	"bytes"
	"github.com/goccy/go-graphviz"
	"github.com/goccy/go-graphviz/cgraph"
	"os"
)

func render(f func(graph *cgraph.Graph)) {
	g := graphviz.New()
	graph, err := g.Graph()
	if err != nil {
		panic(err)
	}
	f(graph)

	var buf bytes.Buffer
	if err := g.Render(graph, graphviz.PNG, &buf); err != nil {
		panic(err)
	}
	path := os.TempDir() + `/data.png`
	if err := g.RenderFilename(graph, graphviz.PNG, path); err != nil {
		panic(err)
	}
	openbrowser(`file://` + path)
}
