package main

type stack []string

func (s *stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *stack) push(value string) {
	*s = append(*s, value)
}

func (s *stack) pop() (string, bool) {
	if s.isEmpty() {
		return "", false
	}
	value := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return value, true
}

func main() {
	deps := map[string][]string{
		"tensorflow": {"nvcc", "gpu", "linux"},
		"nvcc":       {"linux"},
		"linux":      {"core"},
		"mylib":      {"tensorflow"},
		"mylib2":     {"requests"},
	}
	stack := stack{}
	color := make(map[string]string)
	res := make([]string, 0)
	for key := range deps {
		stack.push(key)
		color[key] = "white"
		for _, v := range deps[key] {
			color[v] = "white"
		}
	}
	for !stack.isEmpty() {
		key, _ := stack.pop()
		if color[key] == "white" {
			stack.push(key)
			color[key] = "gray"
			data, ok := deps[key]
			if ok && len(data) > 0 {
				for _, v := range deps[key] {
					if color[v] == "white" {
						stack.push(v)
					}
				}
			}
		} else if color[key] == "gray" {
			res = append(res, key)
			color[key] = "black"
		}
	}
	// godump.Dump(res)
}
