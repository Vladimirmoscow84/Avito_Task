package main

var order []string
var color map[string]string

func main() {
	deps := map[string][]string{
		"tensorflow": {"nvcc", "gpu", "linux"},
		"nvcc":       {"linux"},
		"linux":      {"core"},
		"mylib":      {"tensorflow"},
		"mylib2":     {"requests"},
	}

	initialize(deps)
	mainTopSort(deps)
	// godump.Dump(order)

}

func initialize(vert map[string][]string) {
	order = make([]string, 0)
	color = make(map[string]string)
	for key, value := range vert {
		color[key] = "white"
		for _, v := range value {
			color[v] = "white"
		}
	}
}

func topSort(v string, vert map[string][]string) {
	color[v] = "gray"
	outgoingEdges := getOutgoingEdges(v, vert)
	for _, w := range outgoingEdges {
		if color[w] == "white" {
			topSort(w, vert)
		}
	}
	color[v] = "black"
	order = append(order, v)
}

func mainTopSort(vert map[string][]string) {
	for key := range vert {
		if color[key] == "white" {
			topSort(key, vert)
		}
	}
}

func getOutgoingEdges(v string, vert map[string][]string) []string {
	return vert[v]
}
