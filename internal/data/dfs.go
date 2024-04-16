package data

import "slices"

type DfsResult struct {
	Depth   []int
	Parents []int
	End     int
	Start   int
}

func Dfs(g Graph) DfsResult {
	max_depth := 0
	max_depth_node := 0

	var dfs1 func(int, int, int)
	dfs1 = func(node int, parent int, depth int) {
		if depth > max_depth {
			max_depth = depth
			max_depth_node = node
		}
		for _, neighbor := range g.Edges[node] {
			if neighbor != parent {
				dfs1(neighbor, node, depth+1)
			}
		}
	}
	dfs1(0, -1, 0)

	r := DfsResult{
		Depth:   make([]int, g.Nodes),
		Parents: make([]int, g.Nodes),
		Start:   max_depth_node,
	}

	var dfs2 func(int, int, int)
	dfs2 = func(node int, parent int, depth int) {
		if depth > max_depth {
			max_depth = depth
			max_depth_node = node
		}
		r.Depth[node] = depth
		r.Parents[node] = parent
		for _, neighbor := range g.Edges[node] {
			if neighbor != parent {
				dfs2(neighbor, node, depth+1)
			}
		}
	}

	max_depth = 0
	dfs2(max_depth_node, -1, 0)

	r.End = max_depth_node
	return r
}

func Path(d DfsResult) []int {
	path := []int{d.End}
	for path[len(path)-1] != d.Start {
		path = append(path, d.Parents[path[len(path)-1]])
	}
	slices.Reverse(path)
	return path
}
