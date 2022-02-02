package input

import (
	"bufio"
	"fmt"
	"os"
)

func ToArray(filePath string) []string {
	var list []string
	file, err := os.Open(filePath)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
		list = append(list, scanner.Text())

	}

	return list

}
