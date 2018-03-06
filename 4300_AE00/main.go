/* 
    Created by Smart Mek on 3/6/18 4:38 PM
*/

package main

import (
	"fmt"
	"math"
)

func main() {

	var N float64
	fmt.Scan(&N)

	sq := math.Floor(math.Sqrt(N))

	sum := 0.0
	for i := 1.0; i <= sq; i++ {

		sum += math.Floor(float64(N) / i) - (i - 1.0)
	}

	fmt.Println(int64(sum))
}
