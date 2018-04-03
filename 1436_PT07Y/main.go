/*
   Created by Smart Mek on Apr 03, 2018 11:12 AM
*/

package main

import "fmt"

var adj [][]int
var visited []bool
var recursion []bool

func main() {
	var N, M, u, v int
	var cyclic bool

	fmt.Scan(&N)
	fmt.Scan(&M)

	adj = make([][]int, N + 1)
	visited = make([]bool, N + 1)
	recursion = make([]bool, N + 1)

	for m := 0; m < M; m++ {
		fmt.Scan(&u)
		fmt.Scan(&v)

		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	fmt.Println("before", adj, visited)

	// iterate through the nodes
	for n := 1; n <= N; n++ {
		fmt.Println(n, visited)
		if isCyclic(n) {
			cyclic = true
			break
		}
	}

	fmt.Println("after", adj, visited, cyclic)

	if cyclic {
		fmt.Println("graph is cyclic")
	} else {
		fmt.Println("graph isn't cyclic")
	}
}

// redundant part
func dfs(n int) {
	visited[n] = true

	for i := 0; i < len(adj[n]); i++ {
		w := adj[n][i]

		if !visited[w] {
			dfs(w)
		}
	}
}

func isCyclic(n int) bool {
	if !visited[n] {
		visited[n] = true
		recursion[n] = true

		for i := 0; i < len(adj[n]); i++ {
			if !visited[adj[n][i]] && isCyclic(adj[n][i]) || recursion[adj[n][i]]{
				return true
			}
		}
	}
	recursion[n] = false
	return false
}
