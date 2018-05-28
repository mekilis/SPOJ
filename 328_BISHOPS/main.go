/*
   Created by Smart Mek on May 28, 2018 7:52 PM
*/

package main

import (
	"math/big"
	"fmt"
	"bufio"
	"os"
)

func main() {
	i := big.NewInt(0)
	two := big.NewInt(2)

	var str string
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		str = scanner.Text()
		if str == "" {
			break
		}
		i.SetString(str, 10)

		if str != "0" && str != "1" {
			i.Mul(i, two)
			i.Sub(i, two)
		}

		fmt.Println(i)
	}
}
