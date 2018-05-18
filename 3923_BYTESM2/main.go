/*
   Created by Smart Mek on May 18, 2018 2:20 PM
*/

package main

import (
	"fmt"
)

func main() {
	//start := time.Now()
	var T, h, w, m int
	//var line []string
	tiles := make([][]int, 0)
	tilesCopy := make([][]int, 0)

	//file, err := os.Open("SPOJ/3923_BYTESM2/bigdata.txt")
	//if err != nil {
	//	log.Fatal(err)
	//}

	//defer file.Close()

	//scanner := bufio.NewScanner(file)
	//if scanner.Scan() {
	//	T, _ = strconv.Atoi(scanner.Text())
	//}

	fmt.Scan(&T)
	for t := 0; t < T; t++ {
		fmt.Scan(&h)
		fmt.Scan(&w)

		//if scanner.Scan() {
		//	line = strings.Split(scanner.Text(), " ")
		//	h, _ = strconv.Atoi(line[0])
		//	w, _ = strconv.Atoi(line[1])
		//}

		tiles = make([][]int, h)
		tilesCopy = make([][]int, h)

		for i := 0; i < h; i++ {
			//if scanner.Scan() {
			//	line = strings.Split(scanner.Text(), " ")
			//}
			for j := 0; j < w; j++ {
				fmt.Scan(&m)
				//m, _ = strconv.Atoi(line[j])
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
					//fmt.Println("tilesCopy", tilesCopy)
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

		//fmt.Println("max:", max, tiles)
		//fmt.Println(tilesCopy)
		for i := 0; i < w; i++ {
			if tiles[h-1][i] > max {
				max = tiles[h-1][i]
			}
		}
		fmt.Println(max)
		//fmt.Println("time taken", time.Since(start).Seconds())

	}

}
