package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func Calculator(w http.ResponseWriter, r *http.Request) {
	// TODO: implement a calculator
	parts := strings.Split(r.URL.Path, "/")
	if len(parts)!=4{
		fmt.Fprintf(w, "Error!")
		return
	}
	// fmt.Print(parts)
	operator := parts[1]
	n1, err1:= strconv.Atoi(parts[2])
	n2, err2:= strconv.Atoi(parts[3])
	num1 := int64(n1)
	num2 := int64(n2)
	if err1 != nil || err2 != nil {
		fmt.Fprintf(w, "Error!")
		return
	}
	var result int64
	var reminder int64
	switch operator {
		case "add":
			result = int64(num1 + num2)
			fmt.Fprintf(w, "%d + %d = %d", num1, num2, result)
			return
		case "sub":
			result = int64(num1 - num2)
			fmt.Fprintf(w, "%d - %d = %d", num1, num2, result)
			return
		case "mul":
			result = int64(num1 * num2)
			fmt.Fprintf(w, "%d * %d = %d", num1, num2, result)
			return
		case "div":{
			if num2 == 0 {
				fmt.Fprintf(w, "Error!")
				return
			}
			result = int64(num1) / int64(num2)
			if (num1 % num2) != 0 {
				reminder = num1 - (result*num2)
				fmt.Fprintf(w, "%d / %d = %d, reminder = %d", num1, num2, result, reminder)
				return
			}else {
				fmt.Fprintf(w, "%d / %d = %d", num1, num2, result)
				return
			}
		}
		default:
			fmt.Fprintf(w, "Error!")
			return
	}	
	
}

func main() {
	http.HandleFunc("/", Calculator)
	log.Fatal(http.ListenAndServe(":8083", nil))
}
