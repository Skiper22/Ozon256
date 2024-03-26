package main

import (
	"fmt"
	"strconv"
	"strings"
)

func conv_tostr(arr []int) string {
	var result []string
	start := arr[0]
	end := arr[0]

	for i := 1; i < len(arr); i++ {
		if arr[i] == -1 {
			continue
		}

		if arr[i] == end+1 {
			end = arr[i]
		} else {
			if start == end {
				result = append(result, strconv.Itoa(start))
			} else {
				result = append(result, fmt.Sprintf("%d-%d", start, end))
			}
			start = arr[i]
			end = arr[i]
		}
	}

	if start == end {
		result = append(result, strconv.Itoa(start))
	} else {
		result = append(result, fmt.Sprintf("%d-%d", start, end))
	}

	return strings.Join(result, ",")
}

func parse_numbers(str string) []int{
	nums := strings.Split(input, ",")

	var result []int
	for _, num := range nums {
		if strings.Contains(num, "-") {
			rp := strings.Split(num, "-")
			start, err := strconv.Atoi(rp[0])
			end, err := strconv.Atoi(rp[1])

			for i := start; i <= end; i++ {
				result = append(result, i)
			}
		} else {
			num, err := strconv.Atoi(num)
			result = append(result, num)
		}
	}

	return result
}

func main() {
	var count int
	size := [100]int{}
	arr := [100][100]int{}
	str := [100]string{}
	fmt.Scan(&count)
	for i := 0; i < count; i++ {
		fmt.Scan(&size[i])
		fmt.Scan(&str[i])

	}
	for i:=0; i < count; i++{
		arr[i] = parse_numbers(str[i])
	}
	for i:=0; i < count; i++{
		for j:=0; j < size[i]; j++{
			if j+1 == arr[i][j]{
				arr[i][j]=-1
			}

		}
	}
	for i:=0; i < count; i++{
		conv_tostr(arr[i])
	}

}
