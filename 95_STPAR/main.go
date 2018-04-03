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
			| STACK |
	*/

	fmt.Scan(&N)

	for N != 0 {

		possible = true

		for n := 1; n <= N; n++ {
			fmt.Scan(&x)

			// append item to deque first
			deque = append(deque, x)
			dequeLen++

			fmt.Println("appended to deque", deque, dequeLen)

			if dequeLen > 1 {
				// at least two or more items are in the deque
				fmt.Println("two or more items are in the deque", deque)

				// compare entry with last item in queue if any
				if queueLen > 0 {
					if x < queue[queueLen - 1] {
						fmt.Println("bad sequence for queue")
						possible = false
						break
					}
				}

				// if new entry is greater than the top, remove top
				if deque[dequeLen- 1] > deque[dequeLen- 2] {

					fmt.Println(deque[dequeLen - 1], "is greater than", deque[dequeLen - 2], "in", deque)
					// move lower to queue
					queue = append(queue, deque[dequeLen- 2])
					queueLen++
					fmt.Println("queue now has", queue)

					// update deque
					deque = append(deque[:dequeLen- 2], deque[dequeLen- 1:]...)
					dequeLen--
					fmt.Println("deque is now updated to", deque, "with length", dequeLen)
				}
			}

			// check queue
			if queueLen > 1 {
				fmt.Println("two or more items in the queue", queue, "now comparing", queue[queueLen - 1], "with", queue[queueLen - 2])
				if queue[queueLen - 1] < queue[queueLen - 2] {
					// invalid order
					fmt.Println("invalid queue sequence... breaking")
					possible = false
					break
				}

				fmt.Println("sequence is valid for queue", queue)
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

		fmt.Println("variables used", x, N, dequeLen, queueLen, deque, queue, possible)

		fmt.Scan(&N)
	}
}
