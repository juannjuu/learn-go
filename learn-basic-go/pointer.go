//POINTER IS PASSING DATA BY REFERENCE
//TO USE VARIABLE AS POINTER WITH * AND REFERENCE WITH &

package main

import "fmt"

type Address struct {
	City, Province string
}

func ChangeCityToBandung(address *Address) {
	address.City = "Bandung"
}

func main() {
	var address1 Address = Address{"Bekasi", "West Java"}
	var address2 *Address = &address1
	var address3 *Address = &address1

	address2.City = "Depok"

	*address2 = Address{"Cirebon", "West Java"}

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

	var address4 *Address = new(Address)
	// address4.City = "Jakarta"
	// address4.Province = "DKI Jakarta"
	fmt.Println(address4)

	var alamat = Address{
		City:     "",
		Province: "West Java",
	}
	var alamatPointer *Address = &alamat
	ChangeCityToBandung(alamatPointer)
	fmt.Println(alamatPointer)
}
