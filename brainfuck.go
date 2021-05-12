package gobf

import "fmt"

func BrainFuck(code string) {
	cnt := 0
	ptr := 0
	var ascii [30000]int
	for {
		if cnt == len(code) {
			break
		}
		if string(code[cnt]) == "+" {
			ascii[ptr] += 1
		} else if string(code[cnt]) == "-" {
			ascii[ptr] -= 1
		} else if string(code[cnt]) == ">" {
			ptr += 1
		} else if string(code[cnt]) == "<" {
			ptr -= 1
		} else if string(code[cnt]) == "[" {
			if ascii[ptr] == 0 {
				cnt1 := 1
				for cnt1 != 0 {
					cnt++
					if string(code[cnt]) == "[" {
						cnt1++
					} else if string(code[cnt]) == "]" {
						cnt1--
					} else {
						continue
					}
				}
			}
		} else if string(code[cnt]) == "]" {
			if ascii[ptr] != 0 {
				cnt2 := 1
				for cnt2 != 0 {
					cnt--
					if string(code[cnt]) == "]" {
						cnt2++
					} else if string(code[cnt]) == "[" {
						cnt2--
					} else {
						continue
					}
				}
			}
		} else if string(code[cnt]) == "." {
			fmt.Printf("%c", ascii[ptr])
		}
		cnt++
	}
}
