package main

import "fmt"

func main() {
	/*intP := new(int)
	*intP = 10*/

	intP := 10
	change(intP)
	fmt.Printf("%v -- %p\n", intP)
}

func change(intP int) {
	intP = 100
}
