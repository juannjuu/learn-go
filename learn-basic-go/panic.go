//NOTES
//PANIC IS SIMILAR WITH BREAK TO END A FUNCTION

package main

import "fmt"

func endApp(){
	fmt.Println("Aplikasi Selesai")
}

func runApp(error bool){
	defer endApp()
	if error {
		panic("Aplikasi Error")
	}
	fmt.Println("Aplikasi Berjalan")
}

func main() {
	runApp(false)
	runApp(true)
}