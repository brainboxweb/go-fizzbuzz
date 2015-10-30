package main

func FizzBuzz(input int) string {

	result := ""
	if input%3 == 0 {result += "Fizz"}
	if input%5 == 0 {result += "Buzz"}

	return result
}


