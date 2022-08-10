package ex3

import "fmt"

// print the unicode point of all capital letters in the alphabet, three times each.

func ex2() {
	for i := 65; i <= 90; i++ {
		fmt.Println(i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t%#U", i)
		}
	}
}
