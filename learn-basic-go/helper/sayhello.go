//IN 1 PACKAGE CANNOT CONTAIN SAME FUNCTIONS NAME OR REDECLARE
//TO EXPORT SOMETHING MUST BE INITIALIZED WITH CAPITAL LETTER

package helper

import "fmt"

func SayHello(name string) {
	fmt.Println("Hello", name)
}

// CANNOT BE EXPORTED TO ANOTHER
func sayHi(name string) {
	fmt.Println("Hi", name)
}
