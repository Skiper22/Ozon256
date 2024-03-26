package main

import (
	"fmt"
)

func main() {
	var count int
	seq := [100][101]int{}
	fmt.Scan(&count)
	for i := 0; i < count; i++ {
		fmt.Scan(&seq[i][100])
		for j := 0; j < seq[i][100]; j++ {
			fmt.Scan(&seq[i][j])
		}
	}
	for i:=0; i < count; i++{
		ls := make([]int, 0)

		start := seq[i][0]
		length := 0

		for j := 1; j < seq[i][100]; j++ {
			if (seq[i][j]-seq[i][j-1])*length >= 0 && seq[i][j]-seq[i][j-1] == -1 {
				length += seq[i][j] - seq[i][j-1]
			} else if (seq[i][j]-seq[i][j-1])*length >= 0 && seq[i][j]-seq[i][j-1] == 1 {
				length += seq[i][j] - seq[i][j-1]
			} else {
				ls = append(ls, start)
				ls = append(ls, length)
				start = seq[i][j]
				length = 0
			}
		}

		ls = append(ls, start)
		ls = append(ls, length)
		fmt.Println(len(ls))
		for k:=0; k< len(ls); k++{
			fmt.Print(ls[k]," ")
		}
		fmt.Println()
	}
}
