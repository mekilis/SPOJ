/*
   Created by Smart Mek on May 18, 2018 2:20 PM
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"log"
)

func main() {
	var T, h, w, m int
	var line []string
	tiles := make([][]int, 0)
	tilesCopy := make([][]int, 0)

	file, err := os.Open("SPOJ/3923_BYTESM2/data.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		T, _ = strconv.Atoi(scanner.Text())
	}

	//fmt.Scan(&T)
	for t := 0; t < T; t++ {
		//fmt.Scan(&h)
		//fmt.Scan(&w)

		if scanner.Scan() {
			line = strings.Split(scanner.Text(), " ")
			h, _ = strconv.Atoi(line[0])
			w, _ = strconv.Atoi(line[1])
		}


		tiles = make([][]int, h)
		tilesCopy = make([][]int, h)

		for i := 0; i < h; i++ {
			if scanner.Scan() {
				line = strings.Split(scanner.Text(), " ")
			}
			for j := 0; j < w; j++ {
				//fmt.Scan(&m)
				m, _ = strconv.Atoi(line[j])
				tiles[i] = append(tiles[i], m)
				tilesCopy[i] = append(tilesCopy[i], m)
			}
		}

		max := 0
		var left, right, down int
		for i := 0; i < h - 1; i++ {
			for j := 0; j < w; j++ {

				//// downwards
				//temp := tiles[i][j] + tiles[i+1][j]
				//if temp > max {
				//	max = temp
				//}
				//
				//// left diagonal
				//if j > 0 {
				//	temp = tiles[i][j] + tiles[i+1][j-1]
				//	if temp > max {
				//		max = temp
				//	}
				//}
				//
				//// check right diagonal
				//if j < w - 1 {
				//	temp = tiles[i][j] + tiles[i+1][j+1]
				//	if temp > max {
				//		max = temp
				//	}
				//}

				if j > 0 && j < w -1 {
					// in between
					//fmt.Println("tilesCopy", tilesCopy)
					left = tilesCopy[i][j] + tilesCopy[i+1][j-1]
					right = tilesCopy[i][j] + tilesCopy[i+1][j+1]
					down = tilesCopy[i][j] + tilesCopy[i+1][j]

					// check
					if left > tiles[i+1][j-1] {
						tiles[i+1][j-1] = left
					}
					if right > tiles[i+1][j+1] {
						tiles[i+1][j+1] = right
					}
					if down > tiles[i+1][j] {
						tiles[i+1][j] = down
					}
				} else if j == 0 {
					right = tilesCopy[i][j] + tilesCopy[i+1][j+1]
					down = tilesCopy[i][j] + tilesCopy[i+1][j]

					if right > tiles[i+1][j+1] {
						tiles[i+1][j+1] = right
					}
					if down > tiles[i+1][j] {
						tiles[i+1][j] = down
					}

				} else {
					left = tilesCopy[i][j] + tilesCopy[i+1][j-1]
					down = tilesCopy[i][j] + tilesCopy[i+1][j]

					if left > tiles[i+1][j-1] {
						tiles[i+1][j-1] = left
					}

					if down > tiles[i+1][j] {
						tiles[i+1][j] = down
					}
				}
			}
			break
		}

		fmt.Println("max:", max, tiles)
		//fmt.Println(tilesCopy)

	}

}
