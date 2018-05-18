/*
   Created by Smart Mek on May 17, 2018 10:50 AM
*/

package main

import "fmt"

var adj [][]int
var dist []int
var longest, node = 0, 0

func main() {
	var N int // nodes
	fmt.Scan(&N)
	adj = make([][]int, N+1)
	dist = make([]int, N+1)
	resetDistance()

	var u, v int
	for n := 0; n < N - 1; n++ {
		fmt.Scan(&u)
		fmt.Scan(&v)

		// since graph is undirected
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}
	fmt.Println("graph", adj)
	longest = 0
	dfs(N)
	fmt.Println(longest)
}

func resetDistance() {
	for d := 0; d < len(dist); d++ {
		dist[d] = -1
	}
}

func dfs(n int) {
	// using non-recursive method
	resetDistance()

	stack := make([]int, 0)
	top := 0
	stack = append(stack, n)
	top++

	// distance from node to itself
	//dist[n] = 0

	var u, v int
	for len(stack) != 0 {
		// pop
		u = stack[top - 1]
		stack = append(stack[:top - 1])
		top--

		// catch same node here
		if dist[u] == -1 {
			fmt.Println("node", u, "has been visited. stack:", stack)
			dist[u] = 0
		}

		for i := 0; i < len(adj[u]); i++ {

			// check neighbours
			v = adj[u][i]
			// correct offset (v starts from zero, dist from one)
			if dist[v] == -1 {
				stack = append(stack, v)
				top++
				//dist[v] = dist[u] + 1
			}
		}
	}

	var maxDistance, node int
	for i := 0; i < len(dist); i++ {
		if dist[i] > maxDistance {
			maxDistance, node = dist[i], i
		}
	}

	fmt.Println("max and node", maxDistance, node)

}
//graph [[] [2 3] [1 4 5] [1] [2 6] [2 7] [4] [5]]