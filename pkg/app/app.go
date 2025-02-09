package app

import (
	"fmt"
	"os"
)

func App() {
	data := InputData(os.Stdin)
	fmt.Println("--------------------")
	fmt.Println(PickRandom(data))
}
