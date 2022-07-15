package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//54404 54430
	sliceData := make([]interface{}, 0)
	sliceData = append(sliceData, "沉默小管")
	sliceData = append(sliceData, 18)
	fmt.Println(sliceData)
	bytes1, _ := json.Marshal(sliceData)
	fmt.Println(bytes1)
	stringData1 := string(bytes1)
	fmt.Println(stringData1)
}

func change(intP int) {
	intP = 100
}
