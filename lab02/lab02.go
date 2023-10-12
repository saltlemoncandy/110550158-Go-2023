package main

import "fmt"
import "strconv"


func main() {
	var n int64

	fmt.Print("Enter a number: ")
	fmt.Scanln(&n)

	result := Sum(n)
	fmt.Println(result)
}

func Sum(n int64) string {
	// TODO: Finish this function
	var x int
	var str string
	num := int(n)
	if num != 0 {
		str = "1"
		x = x+1
		for i:=2; i<=num; i++ {
			if i%7!=0 {
				x += i
				tmpstr := strconv.Itoa(i)
				str = fmt.Sprintf("%s+%s", str, tmpstr)
			}
		}
	}
	ans := strconv.Itoa(x)
	str = fmt.Sprintf("%s=%s", str, ans)
	return str
}