package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)


// TODO: Create a struct to hold the data sent to the template
type pagedata struct {
    Expression string
    Result    string
}

func Calculator(w http.ResponseWriter, r *http.Request) {
	// TODO: Finish this function

	values := r.URL.Query()
	operator := values.Get("op")
	num1Str := values.Get("num1")
	num2Str := values.Get("num2")
	n1, err1:= strconv.Atoi(num1Str)
	n2, err2:= strconv.Atoi(num2Str)
	num1 := int64(n1)
	num2 := int64(n2)

	if err1 != nil || err2 != nil {
		http.ServeFile(w, r, "error.html")
		return		
    }

	var result int64
	var e string
	var re string
	switch operator {
		case "add":
			result = int64(num1 + num2)
			e = fmt.Sprintf("%d + %d", num1, num2)
			re = fmt.Sprintf("%d", result)
			data := pagedata{
				Expression: e,
				Result: re,
			}
			err := template.Must(template.ParseFiles("index.html")).Execute(w, data)
			if err != nil {
				http.ServeFile(w, r, "index.html")
				return
			}
			
		case "sub":
			result = int64(num1 - num2)
			e = fmt.Sprintf("%d - %d", num1, num2)
			re = fmt.Sprintf("%d", result)
			data := pagedata{
				Expression: e,
				Result: re,
			}
			err := template.Must(template.ParseFiles("index.html")).Execute(w, data)
			if err != nil {
				http.ServeFile(w, r, "index.html")
				return
			}
		case "mul":
			result = int64(num1 * num2)
			e = fmt.Sprintf("%d * %d", num1, num2)
			re = fmt.Sprintf("%d", result)
			data := pagedata{
				Expression: e,
				Result: re,
			}
			err := template.Must(template.ParseFiles("index.html")).Execute(w, data)
			if err != nil {
				http.ServeFile(w, r, "index.html")
				return
			}
		case "div":
			if num2 == 0 {
				http.ServeFile(w, r, "error.html")
				return				
			}
			result = int64(num1/num2)
			e = fmt.Sprintf("%d / %d", num1, num2)
			re = fmt.Sprintf("%d", result)
			data := pagedata{
				Expression: e,
				Result: re,
			}
			err := template.Must(template.ParseFiles("index.html")).Execute(w, data)
			if err != nil {
				http.ServeFile(w, r, "index.html")
				return
			}
		case "gcd":
			var tmp1 = num1
			var tmp2 = num2
			var tmp int64
			for tmp2 != 0 {				
				tmp = tmp1%tmp2
				tmp1 = tmp2
				tmp2 = tmp
			}
			result = int64(tmp1)
			e = fmt.Sprintf("GCD(%d, %d)", num1, num2)
			re = fmt.Sprintf("%d", result)
			data := pagedata{
				Expression: e,
				Result: re,
			}
			err := template.Must(template.ParseFiles("index.html")).Execute(w, data)
			if err != nil {
				http.ServeFile(w, r, "index.html")
				return
			}
		case "lcm":
			var tmp1 int64
			var tmp2 int64
			var tmp int64
			var gcd int64
			tmp1 = num1
			tmp2 = num2
			for tmp2 != 0{
				tmp = tmp1%tmp2
				tmp1 = tmp2
				tmp2 = tmp
			}
			gcd = tmp1
			result = int64(num1*num2/gcd)
			e = fmt.Sprintf("LCM(%d, %d)", num1, num2)
			re = fmt.Sprintf("%d", result)
			data := pagedata{
				Expression: e,
				Result: re,
			}
			err := template.Must(template.ParseFiles("index.html")).Execute(w, data)
			if err != nil {
				http.ServeFile(w, r, "index.html")
				return
			}
		
		default:
			http.ServeFile(w, r, "error.html")
			return
	}
}

func main() {
	http.HandleFunc("/", Calculator)
	log.Fatal(http.ListenAndServe(":8084", nil))
}
