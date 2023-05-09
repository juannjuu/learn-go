// NOTES
// LAST PARAM IN FUNCTION CAN BE VARAGS
// VARAGS CAN BE MORE THAN 1 VALUE
// IF WE USE ARRAY AS PARAM, WE HAVE TO DECLARE AN ARRAY BUT NO NEED IN VARAGS

package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0
	for _, value := range numbers {
		total += value
	}
	return total
}

func main() {
	total := sumAll(10, 20, 10, 30)
	fmt.Println(total)

	//USE SLICE AS PARAM
	slice := []int{10, 20, 30, 40}
	total = sumAll(slice...)
	fmt.Println(total)
}
