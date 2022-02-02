package enumeration

import (
	"fmt"
	"net/http"

	"github.com/parkerdouglass/go-bust/structs"
)

func Enumerate(bash structs.Bash) []string {
	var validURLS = []string{}
	for i := 0; i < len(bash.WordListArray); i++ {
		res, err := http.Get(bash.URL + bash.WordListArray[i])

		if err != nil {
			fmt.Println(err)
		}
		if res.StatusCode == 200 {
			validURLS = append(validURLS, bash.WordListArray[i])
		} else {
			fmt.Println(bash.WordListArray[i] + " IS NOT VALID")
		}
		defer res.Body.Close()
	}

	fmt.Println("Done!")

	return validURLS
}
