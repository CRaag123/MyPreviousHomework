package main

import "fmt"

func main() {
	cmd, ir := checkAge(19)
	fmt.Println(cmd)
	fmt.Println(ir)
}
func checkAge(age int) (string, bool) {
	switch {
	case age < 18:
		return "к сожалению вы не подходите к нам на работу", false
	case age > 65:
		return "мы не можем принять вас на работу из-за возростных ограничений", false
	case age >= 18 && age < 65:
		return "мы вас берем", true
	default:
		return "значение переменной не определено", false
	}
}
