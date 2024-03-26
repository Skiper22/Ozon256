package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)

	for test := 0; test < t; test++ {
		var n int
		fmt.Scan(&n)

		var temperature int
		fmt.Scan(&temperature)

		// Проверка, что начальная температура в пределах [15, 30]
		if temperature < 15 || temperature > 30 {
			fmt.Println("-1")
			continue
		}

		for i := 0; i < n; i++ {
			var ai int
			var op string
			fmt.Scan(&ai, &op)

			// Проверка, что требование сотрудника в пределах [15, 30]
			if ai < 15 || ai > 30 {
				fmt.Println("-1")
				return
			}

			// Проверка, что текущая температура соответствует требованию сотрудника
			if op == ">=" {
				if ai > temperature {
					temperature = ai
				}
			} else if op == "<=" {
				if ai < temperature {
					temperature = ai
				}
			} else {
				fmt.Println("-1")
				return
			}
		}

		// Выводим температуру после прихода каждого сотрудника
		fmt.Println(temperature)
	}
}
