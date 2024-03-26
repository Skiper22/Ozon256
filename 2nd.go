// Январе, Марте, Мае, Июле, Августе, Октябре и Декабре

package main
import "fmt"

func isLeapYear(year int) int{
	if ((year % 400 == 0) || (year % 100 != 0) && (year % 4 == 0)){
    	return 1
	} else {
    	return 0
    }
}

func main(){
	var count int
	var arr [1000]int
	var day, mon,year int
	fmt.Scan(&count)
	for i:=0; i < count; i++{
		fmt.Scan(&day, &mon, &year)
		if mon == 1 || mon == 3 || mon == 5 || mon == 7 || mon == 8 || mon == 7 || mon == 10 || mon == 12{
			arr[i] = 1
		} else if (day < 31 && mon != 2){
			arr[i] = 1
		} else if mon == 2{
			if day < 29{
				arr[i] = 1
			} else if day == 29 && isLeapYear(year) == 1{
				arr[i] = 1
			}
		}
	}
	for i := 0; i < count; i++{
		if arr[i] == 1{
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}

	}
}
