package main

import (
	"fmt"
	"os"

	"pickup/helper"
)

func main() {
	data := helper.InputData(os.Stdin)
	fmt.Println("--------------------")
	fmt.Println(helper.PickRandom(data))
}
