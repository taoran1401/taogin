package utils

import "fmt"

//打印map
func printMap(result map[string]interface{}) {
	for k, v := range result {
		fmt.Println(k, v)
	}
}

//检查错误并打印
func checkErr(err error, printType string) {
	if printType == "" {
		printType = "panic"
	}
	if err != nil {
		if printType == "panic" {
			panic(err)
		} else if printType == "printf" {
			fmt.Printf("error: %v\n", err)
		}
	}
}
