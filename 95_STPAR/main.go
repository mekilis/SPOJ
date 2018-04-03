/*
   Created by Smart Mek on Mar 29, 2018 12:44 PM
*/

package main

import "fmt"

func main() {

	var N, x, dequeLen, queueLen int
	var possible bool
	deque := make([]int, 0)
	queue := make([]int, 0)
	/*
	========| 		| ==========
						QUEUE
	========|		| ==========
			|		|
			| STACK/|
			| DEQUE |
	*/

	fmt.Scan(&N)

	for N != 0 {

		possible = true

		for n := 1; n <= N; n++ {
			fmt.Scan(&x)

			if !possible {
				// just continue
				continue
			}

			// append item to deque first
			deque = append(deque, x)
			dequeLen++

			if dequeLen > 1 {
				// at least two or more items are in the deque

				// compare entry with last item in queue if any
				if queueLen > 0 {
					if x < queue[queueLen - 1] {
						possible = false
						continue
					}
				}

				// if new entry is greater than the top, remove top
				for dequeLen > 1 && deque[dequeLen- 1] > deque[dequeLen- 2] {

					// move lower to queue
					queue = append(queue, deque[dequeLen- 2])
					queueLen++

					// update deque
					deque = append(deque[:dequeLen- 2], deque[dequeLen- 1:]...)
					dequeLen--
				}

				// check deque
				if dequeLen > 1 {
					if deque[dequeLen - 1] > deque[dequeLen - 2] {
						possible = false
						continue
					}
				}
			}

			// check queue
			if queueLen > 1 {
				if queue[queueLen - 1] < queue[queueLen - 2] {
					// invalid order
					possible = false
					continue
				}
			}
		}

		if possible {
			fmt.Println("yes")
		} else {
			fmt.Println("no")
		}

		deque = make([]int, 0)
		queue = make([]int, 0)
		dequeLen = 0
		queueLen = 0

		fmt.Scan(&N)
	}
}
