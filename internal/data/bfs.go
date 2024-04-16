package data

import "slices"

type BFSRes struct {
	Distances []int
	Parents   []int
}

func BFS(graph *Graph, start int) BFSRes {
	r := BFSRes{
		Distances: make([]int, graph.Nodes),
		Parents:   make([]int, graph.Nodes),
	}
	visited := make([]bool, graph.Nodes)
	queue := []int{start}
	r.Distances[start] = 0
	r.Parents[start] = -1

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		visited[node] = true
		for _, edge := range graph.Edges[node] {
			if !visited[edge] {
				r.Distances[edge] = r.Distances[node] + 1
				r.Parents[edge] = node

				queue = append(queue, edge)
			}
		}
	}

	return r
}

func PathTo(start, end int, bfs BFSRes) []int {
	path := []int{end}
	for path[len(path)-1] != start {
		path = append(path, bfs.Parents[path[len(path)-1]])
	}
	slices.Reverse(path)
	return path
}
