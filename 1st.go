// 4-1	3-2  2-3	1-4

package main
import "fmt"

func main(){
	var count int
	var arr [1000]int
	var one_p, two_p,three_p, four_p int
	fmt.Scan(&count)
	value := 0
	for i:=0; i < count; i++{
		for j:=0; j < 10; j++{
			fmt.Scan(&value)
			switch value{
			case 1:
				one_p++
			case 2:
				two_p++
			case 3:
				three_p++
			case 4:
				four_p++
			default:
			}
		}
		if one_p == 4 && two_p == 3 && three_p == 2 && four_p == 1{
			arr[i] = 1
		}
		one_p = 0
		two_p = 0
		three_p = 0
		four_p = 0
	}

	for i := 0; i < count; i++{
		if arr[i] == 1{
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}

	}
}
