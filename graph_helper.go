package leetcode

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/goccy/go-graphviz"
	"github.com/goccy/go-graphviz/cgraph"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strconv"
)

func ParseGraphInput(in string) ([][]int, error) {
	out := [][]int{}
	if err := json.Unmarshal([]byte(in), &out); err != nil {
		return nil, err
	}
	return out, nil
}

func RenderGraph(in string) error {
	graph, err := ParseGraphInput(in)
	if err != nil {
		return err
	}
	return renderGraphLocal(graph)
}

func renderGraphLocal(input [][]int) error {
	g := graphviz.New()
	graph, err := g.Graph()
	if err != nil {
		log.Fatal(err)
	}
	nodes := make([]*cgraph.Node, 0, len(input))
	for i := 0; i < len(input); i++ {
		n, err := graph.CreateNode(strconv.Itoa(i))
		if err != nil {
			return err
		}
		nodes = append(nodes, n)
	}
	for k, l1 := range input {
		for _, l2 := range l1 {
			s := nodes[k]
			t := nodes[l2]
			if _, err := graph.CreateEdge(``, s, t); err != nil {
				return err
			}
		}
	}
	var buf bytes.Buffer
	if err := g.Render(graph, graphviz.PNG, &buf); err != nil {
		log.Fatal(err)
	}
	path := os.TempDir() + `/graph.png`
	if err := g.RenderFilename(graph, graphviz.PNG, path); err != nil {
		log.Fatal(err)
	}
	openbrowser(`file://` + path)
	return nil
}

func openbrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}
}
