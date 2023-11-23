package main

import (
	"fmt"
	"math/big"
	"syscall/js"
	"strconv"
)

func CheckPrime(this js.Value, args []js.Value) interface{} {
	// TODO: Check if the number is prime
	
	numStr := js.Global().Get("value").Get("value").String()
	// str := numStr.toString()
	v, _ := strconv.ParseInt(numStr, 10, 64)
	bigNumber := big.NewInt(v)

	isPrime := bigNumber.ProbablyPrime(0)
	
	if isPrime {
		js.Global().Get("answer").Set("innerText", "It's prime")
	} else {
		js.Global().Get("answer").Set("innerText", "It's not prime")
	}
	
	return "ok"
}

func registerCallbacks() {
	// TODO: Register the function CheckPrime
	js.Global().Set("CheckPrime", js.FuncOf(CheckPrime))
}

func main() {
	fmt.Println("Golang main function executed")
	registerCallbacks()

	// Block the main thread forever
	select {}
}
