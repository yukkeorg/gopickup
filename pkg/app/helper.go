package app

import (
	"bufio"
	"io"
	"math/rand"
)

func InputData(reader io.Reader) []string {
	scanner := bufio.NewScanner(reader)
	var data []string
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	return data
}

func PickRandom(data []string) string {
	index := rand.Intn(len(data))
	return data[index]
}
