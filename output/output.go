package output

import (
	"fmt"
	"os"
)

func Output(dest string, valid []string) {
	if len(valid) == 0 {
		fmt.Println("None of the URLs are valid")
	}
	file, err := os.Create(dest)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	for i := 0; i < len(valid); i++ {
		file.Write([]byte(valid[i] + "\n"))
	}
}
