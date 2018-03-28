/*
   Created by Smart Mek on 3/26/18 1:53 PM
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var T, augend, addend, sum int
	scanner := bufio.NewScanner(os.Stdin)
	var line string
	var err1, err2 error

	fmt.Scanln(&T)

	for t := 0; t < T; t++ {

		if scanner.Scan() {
			line = scanner.Text()

			for line == "" {
				scanner.Scan()
				line = scanner.Text()
			}

			p := strings.Split(line, " ")

			augend, err1 = strconv.Atoi(p[0])
			addend, err2 = strconv.Atoi(p[2])
			sum, _ = strconv.Atoi(p[4])

			if err1 != nil {
				augend = sum - addend
			} else if err2 != nil {
				addend = sum - augend
			} else {
				sum = addend + augend
			}

			fmt.Println(augend, p[1], addend, p[3], sum)
		}
	}
}
