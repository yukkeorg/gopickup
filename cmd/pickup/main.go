package main

import (
	"fmt"
	"os"

	"pickup/lib/helper"
)

func main() {
	data := helper.InputData(os.Stdin)
	fmt.Println("--------------------")
	fmt.Println(helper.PickRandom(data))
}
