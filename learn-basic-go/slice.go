package main

import "fmt"

func main() {
	var months = [...]string{
		"January",
        "February",
        "March",
        "April",
        "May",
        "June",
        "July",
        "August",
        "September",
        "October",
        "November",
        "December",
	}
	var slice1 = months[4:7] // index 4 to 7
	fmt.Println(slice1)
	fmt.Println(len(slice1)) // length of slice 1 = 3
	fmt.Println(cap(slice1)) // capacity of slice = 8 (may - dec)

	// months[5] = "Slice Diubah"
	// fmt.Println(slice1)

	// slice1[0] = "Index 0 Diubah" 
	// fmt.Println(months) // effects the array

	var slice2 = months[2:4]
	fmt.Println(slice2) // [March, April]

	var slice3 = append(slice2, "Juju")
	fmt.Println(slice3) // [March April Juju]
	slice3[1] = "Bukan April"
	fmt.Println(slice3) // [March Bukan Aprik Juju]
	
	fmt.Println(slice2) // [March Bukan April] 
	fmt.Println(months) // [January February March Bukan April May June July August September October November December]

	// if slice is full, it will makes new array
	// var slice2 = months[10:]
	// fmt.Println(slice2)

	// var slice3 = append(slice2, "Juju")
	// fmt.Println(slice3) // [November December Juju]
	// slice3[1] = "Bukan Desember"
	// fmt.Println(slice3) // [November Bukan December Juju]
	
	// fmt.Println(slice2) // [November Desember] 
	// fmt.Println(months) // [January February March April May June July August September October November December]

	//Make Slice
	newSlice := make([]string, 2, 5)
	newSlice[0] = "Juan"
	newSlice[1] = "Rahmadhika"
	fmt.Println(newSlice)
	fmt.Println(len(newSlice)) // length = 2
	fmt.Println(cap(newSlice)) // capcatity = 5

	//Copy Slice
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)
}