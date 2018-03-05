/* 
    Created by Smart Mek on 3/4/18 11:05 PM
*/

package main

import "fmt"

func main() {

	var t int
	var exp string

	operators := make(map[string]int32)
	operators["+"] = 2
	operators["-"] = 2
	operators["*"] = 3
	operators["/"] = 3
	operators["^"] = 4

	fmt.Scan(&t)
	for tt := 0; tt < t; tt++ {

		//assuming input has no spaces, hence no need to trim
		fmt.Scanln(&exp)

		stack := make([]string, 0)
		word := make([]string, 0)

		// reading all tokens
		for _, c:= range exp {

			// reading next token
			// check if it's not an operator
			if c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' || c > '0' && c <= '9' {

				// print out
				//fmt.Println("pushing", string(c), "to output")
				word = append(word, string(c))
				continue

			} else if c == '(' {

				// push to stack
				//fmt.Println("pushing ( to stack")
				stack = append(stack, string(c))
				continue

			} else if c == ')' {

				//fmt.Println(") found keep popping...")

				// keep popping from the stack until token is '('
				for top := len(stack) - 1; top >= 0; top-- {

					if stack[top] == string('(') {

						//fmt.Println("removing ( from stack")
						stack = append(stack[:top])
						//fmt.Println("after removing (", stack, top)
						break

					} else {

						//fmt.Println("before (", word, stack, top)

						word = append(word, stack[top])
						stack = append(stack[:top])

						//fmt.Println("after (", word, stack, top)
					}
				}

			} else {

				// assumed to be other operators
				//fmt.Println("pushing operator", string(c))

				// while there is another operator at the top of the stack with a higher precedence
				top := len(stack) - 1

				if top >= 0 {

					for operators[stack[top]] > c && top >= 0 {

						//fmt.Println("top at index", top)

						// pop this other operator
						word = append(word, stack[top])
						stack = append(stack[:top])
						top--
					}
				}

				// push current operator to the top of the stack
				stack = append(stack, string(c))

			}

		} //(1+2)*(3/4)^(5+6) //((a+t)*((b+(a+c))^(c+d)))\

		// pop out all chars in word
		for _, c := range word {
			fmt.Print(c)
		}

		fmt.Println()
	}

	//fmt.Println(t, exp)
}
