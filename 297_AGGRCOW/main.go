/*
   Created by Smart Mek on Apr 16, 2018 1:00 PM
*/

package main

import (
	"fmt"
	"sort"
)

// would define sort.Interface on this
type MyRange []uint64

func (m MyRange) Len() int           { return len(m) }
func (m MyRange) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
func (m MyRange) Less(i, j int) bool { return m[i] < m[j] }

/*func maximumProgramValue(n int32) int64 {
 */ /*
 * Write your code here.
 */ /*
	reader := bufio.NewReader(os.Stdin)
	var sum, max, num int32
	var code string
	fmt.Println("about to start iter")
	for i := int32(0); i < n; i++ {
		fmt.Println("scanning")
		fmt.Fscan(reader, &code)
		fmt.Scan(&num)
		if code == "add" {
			sum += num
		} else if code == "set" {
			sum = num
		}
		if sum > max {
			max = sum
		}
	}
	return int64(max)
}*/

func main() {
	/*

	type m struct {
		Id int
	}

	permMap := make(map[string][]int)

	permMap["hello"] = []int{1, 2, 3}

	if _, ok := permMap["hello"]; ok {
		fmt.Println("Exists")
	}

	if _, ok := permMap["hellox"]; ok {
		fmt.Println("Exists")
	} else {
		fmt.Println("Not exists")
	}
	return
	*/

	/** HACKERRANK W37

	ave := (93.443 + 93.445) / 2.0
	baseAveFloat := ave + 0.001

	baseAve := strconv.FormatFloat(baseAveFloat, 'f', 2, 64)
	baseAveFloat, _ = strconv.ParseFloat(baseAve, 64)

	baseAve = strconv.FormatFloat(baseAveFloat, 'f', 2, 64)
	fmt.Println(baseAve)

	return
	**/

	var N, C, T uint64
	fmt.Scan(&T)

	for t := uint64(0); t < T; t++ {
		fmt.Scan(&N)
		fmt.Scan(&C)

		stalls := make(MyRange, 0)
		var x uint64
		for n := uint64(0); n < N; n++ {
			fmt.Scan(&x)
			stalls = append(stalls, x)
		}

		// sort array first
		sort.Sort(stalls)

		fmt.Println("stalls:", stalls)

		p := stalls[N - 1] - stalls[0]

		fmt.Println(binarySearch(stalls, N, p))
	}
}

func binarySearch(A MyRange, n, x uint64) uint64 {
	var p, q, r uint64
	p, r = 0, n-1

	for p <= r {
		q = (p + r) / 2

		if A[q] == x {
			return q
		} else if A[q] > x {
			r = q - 1
		} else {
			p = q + 1
		}
	}

	return 0
}
