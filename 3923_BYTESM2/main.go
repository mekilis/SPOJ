/*
   Created by Smart Mek on May 18, 2018 2:20 PM
*/

package main

import (
	"fmt"
)

func main() {
	var T, h, w, m int
	tiles := make([][]int, 0)
	tilesCopy := make([][]int, 0)

	fmt.Scan(&T)
	for t := 0; t < T; t++ {
		fmt.Scan(&h)
		fmt.Scan(&w)

		tiles = make([][]int, h)
		tilesCopy = make([][]int, h)

		for i := 0; i < h; i++ {
			for j := 0; j < w; j++ {
				tiles[i] = append(tiles[i], m)
				tilesCopy[i] = append(tilesCopy[i], m)
			}
		}

		max := 0
		var left, right, down int
		for i := 0; i < h-1; i++ {
			for j := 0; j < w; j++ {

				// constant direction
				down = tilesCopy[i][j] + tilesCopy[i+1][j]
				if down > tiles[i+1][j] {
					tiles[i+1][j] = down
				}

				if j > 0 && j < w-1 {
					// in between
					left = tilesCopy[i][j] + tilesCopy[i+1][j-1]
					right = tilesCopy[i][j] + tilesCopy[i+1][j+1]

					// check
					if left > tiles[i+1][j-1] {
						tiles[i+1][j-1] = left
					}
					if right > tiles[i+1][j+1] {
						tiles[i+1][j+1] = right
					}

				} else if j == 0 {
					// leftmost edge
					right = tilesCopy[i][j] + tilesCopy[i+1][j+1]

					if right > tiles[i+1][j+1] {
						tiles[i+1][j+1] = right
					}

				} else {
					left = tilesCopy[i][j] + tilesCopy[i+1][j-1]
					if left > tiles[i+1][j-1] {
						tiles[i+1][j-1] = left
					}
				}
			}
			// update tilesCopy row
			tilesCopy[i+1] = tiles[i+1] // this works :)
			//break
		}

		for i := 0; i < w; i++ {
			if tiles[h-1][i] > max {
				max = tiles[h-1][i]
			}//fmt.Println("time taken", time.Since(start).Seconds())
		}

		fmt.Println(max)
	}

}
