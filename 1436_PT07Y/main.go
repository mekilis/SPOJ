/*
   Created by Smart Mek on Apr 03, 2018 11:12 AM
*/

package main

import "fmt"

var adj [][]int
var visited []bool

func main() {
	var N, M, u, v int
	var cyclic bool

	fmt.Scan(&N)
	fmt.Scan(&M)

	// one-indexed
	adj = make([][]int, N + 1)
	visited = make([]bool, N + 1)

	for m := 0; m < M; m++ {
		fmt.Scan(&u)
		fmt.Scan(&v)

		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	// check if it's possible
	if M != N - 1 || N == 1 {
		cyclic = true
		goto fin
	}

	///fmt.Println("before", adj, visited)

	// reuse u and v
	for u = 1; u <= N; u++ {
		if !visited[u] {
			if isCyclic(u, -1) {
				cyclic = true
				break
			}
		}
	}

	//fmt.Println("after", adj, visited)
	fin:
		if cyclic {
			fmt.Println("NO")
		} else {
			fmt.Println("YES")
		}
}

func isCyclic(v int, parent int) bool {

	// mark this node
	//fmt.Println("marking node", v, "with parent", parent)
	visited[v] = true

	// iterate through all adjacent nodes
	for i := 0; i < len(adj[v]); i++ {

		//fmt.Println("working on edge", v, ",", adj[v][i])
		// check if node has been visited, if not, dig deeper
		if !visited[adj[v][i]] {
			if isCyclic(adj[v][i], v) {
				//fmt.Println("a cycle for", v)
				return true
			}
			// check if node is not same as parent
		} else if adj[v][i] != parent {
			//fmt.Println("node", adj[v][i], "has been visited and is not a parent. a cycle exists")
			return true
		}

		//fmt.Println("node", adj[v][i], "has already been visited and is a parent")

	}

	//fmt.Println("no cycle for", v)
	return false
}
