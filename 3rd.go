// Январе, Марте, Мае, Июле, Августе, Октябре и Декабре
// 4 5 8 9 10 12 13 14 15 16 17 18 19

package main
import "fmt"

func numberA(s string) int{
	if (s[0] >= 'A' && s[0] <= 'Z') && (s[1] >= '0' && s[1] <= '9') && (s[2] >= '0' && s[2] <= '9') && (s[3] >= 'A' && s[3] <= 'Z') && (s[4] >= 'A' && s[4] <= 'Z'){
		return 5
	}
	//fmt.Println(s)
	return 1
}

func numberB(s string) int{
	if (s[0] >= 'A' && s[0] <= 'Z') && (s[1] >= '0' && s[1] <= '9') && (s[2] >= 'A' && s[2] <= 'Z') && (s[3] >= 'A' && s[3] <= 'Z'){
		return 4
	}
	// fmt.Println(s)
	return 1
}

func main(){
	var word [1000]string
	var size int
	fmt.Scan(&size)
	for i := 0; i < size; i++{
	 	fmt.Scan(&word[i])
	}
	for i := 0; i < size; i++{

	 	ok := 0
	 	start := 0
		stop := 4
		if len(word[i]) < 4{
			word[i] = " "
		} else {
			for ok == 0{
		 		if start + 4 <= len(word[i]) && numberB(word[i][start:stop]) == 4{
		 			stop += 4
		 			start += 4
		 		} else if start + 5 <= len(word[i]) && numberA(word[i][start:stop + 1]) == 5{
		 			stop += 5
		 			start += 5
		 		} else if start == len(word[i]){
		 			break;
		 		} else {ok = 1}
		 	}
		 	if (ok == 1){
		 		word[i] = " "
		 	}
		}

	}
	for i:= 0; i < size; i++{
		start := 0
		stop := 4
		if word[i] == " "{
			fmt.Println("-")
		} else {
			for {
				if start + 4 <= len(word[i]) && numberB(word[i][start:stop]) == 4{
					fmt.Printf("%s ", word[i][start:stop])
					stop += 4
			 		start += 4
				} else if start + 5 <= len(word[i]) && numberA(word[i][start:stop + 1]) == 5{
					fmt.Printf("%s ", word[i][start:stop+1])
					stop += 5
			 		start += 5
				} else if start == len(word[i]){
					fmt.Println()
			 			break;
			 	}
			}
		 }
	}
}
