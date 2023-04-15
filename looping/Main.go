package main

import (
	"fmt"
)

func main () {
	// similar syntax to JS
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	// can add second index
	// avoid manipulating the index counter
	for i, j := 0, 0; i < 5;  i, j = i+1, j+2 {
		fmt.Print(i)
		fmt.Print(j)
	}

	i := 2
	// can declare index outside of loop
	// typically scoped within loop
	for ; i < 10; i ++{
		fmt.Print(i)
	}
	
	j := 0
	// while loops -> remove incrementor 
	for j < 10 {
		fmt.Print(j)
		j++
	}

	k := 0
	// infinite loop
	for {
		fmt.Println(k)
		k++
		if k == 11 {
			// break loop
			break
		}
	}

	// nested loop
	Loop: // label to break out of nested loop
		for i := 1; i <= 3;  i++ {
			for j := 1; j <=3; j++ {
				fmt.Print(i * j)
				if i * j >=3 {
					break Loop
				}
			}
		}

	slice := []int{1,2,3}
	// can recplace k with _ if not needed
	for k, v := range slice { 
		// loop through slice with key and value
		fmt.Println(k,v)
	}
}
