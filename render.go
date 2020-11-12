package leetcode

import (
	"bytes"
	"github.com/goccy/go-graphviz"
	"github.com/goccy/go-graphviz/cgraph"
	"os"
)

func render(f func(graph *cgraph.Graph) error) error {
	g := graphviz.New()
	graph, err := g.Graph()
	if err != nil {
		return err
	}
	if err := f(graph); err != nil {
		return err
	}

	var buf bytes.Buffer
	if err := g.Render(graph, graphviz.PNG, &buf); err != nil {
		return err
	}
	path := os.TempDir() + `/graph.png`
	if err := g.RenderFilename(graph, graphviz.PNG, path); err != nil {
		return err
	}
	openbrowser(`file://` + path)
	return nil
}
