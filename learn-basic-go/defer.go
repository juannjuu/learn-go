//NOTES
//DEFER USE TO EXECUTE AN FUNCTION ALTHOUGH THERE IS AN ERROR BEFORE

package main

import "fmt"

func logging(){
	fmt.Println("Selesai memanggil function")
}

func runApplication(value int){
	defer logging()
	fmt.Println("Run application")
	result := 10/value
	fmt.Println(result)
}

func main() {
	runApplication(0)
}