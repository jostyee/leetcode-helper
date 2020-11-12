package leetcode

import (
	"encoding/json"
	"fmt"
	"github.com/goccy/go-graphviz/cgraph"
	"log"
	"os/exec"
	"runtime"
	"strconv"
)

func ParseGraphInput(in string) ([][]int, error) {
	graph := [][]int{}
	if err := json.Unmarshal([]byte(in), &graph); err != nil {
		return nil, err
	}
	return graph, nil
}

func ParseEdgesInput(in string) ([][2]int, error) {

	edges := [][2]int{}
	if err := json.Unmarshal([]byte(in), &edges); err != nil {
		return nil, err
	}
	return edges, nil
}

func RenderGraph(graph [][]int) error {
	edges := [][2]int{}
	for s, ds := range graph {
		for _, d := range ds {
			edges = append(edges, [2]int{s, d})
		}
	}
	return RenderGraphByEdges(edges)
}

func RenderGraphByEdges(edges [][2]int) error {
	return render(func(graph *cgraph.Graph) error {
		var err error
		nodes := map[int]*cgraph.Node{}
		for _, l := range edges {
			s, d := l[0], l[1]
			sn := nodes[s]
			if sn == nil {
				sn, err = graph.CreateNode(strconv.Itoa(s))
				if err != nil {
					return err
				}
				nodes[s] = sn
			}
			dn := nodes[d]
			if dn == nil {
				dn, err = graph.CreateNode(strconv.Itoa(d))
				if err != nil {
					return err
				}
				nodes[d] = dn
			}
			if _, err := graph.CreateEdge(``, sn, dn); err != nil {
				return err
			}
		}
		return nil
	})
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
