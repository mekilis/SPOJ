/*
   Created by Smart Mek on Apr 27, 2018 4:47 PM
*/

package main

import "fmt"

func main() {
	var n, x int

	fmt.Scan(&n)

	for n != 0 {
		lookup := make(map[int]int)
		lookupList := make([]int, 0)

		for i := 1; i <= n; i++ {
			fmt.Scan(&x)
			lookup[x] = i
			lookupList = append(lookupList, x)
		}
		//fmt.Println(lookup)
		//fmt.Println(lookupList)

		ambiguous := true

		for i := 1; i <= n; i++ {
			if lookupList[i - 1] != lookup[i] {
				ambiguous = false
				break
			}
		}

		if ambiguous {
			fmt.Println("ambiguous")
		} else {
			fmt.Println("not ambiguous")
		}

		fmt.Scan(&n)
	}
}

/*
ith no is the position of i in the perm
4
1 4 3 2
------------
1 2 3 4
-------------
1 4 3 2 == same
------------

amb
5
2 3 4 5 1
-------------
1 2 3 4 5
-------------
5 1 2 3 4
-------------
non amb
1
1
-------------
amb
0
*/