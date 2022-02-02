package main

import (
	"flag"
	"fmt"

	"github.com/parkerdouglass/go-bust/enumeration"
	"github.com/parkerdouglass/go-bust/input"
	"github.com/parkerdouglass/go-bust/output"
	"github.com/parkerdouglass/go-bust/structs"
)

func main() {
	var url string
	var wordlist string
	var dest string
	flag.StringVar(&url, "u", "", "The URL which you wish to enumerate through")
	flag.StringVar(&wordlist, "w", "", "The wordlist which you would like to use")
	flag.StringVar(&dest, "d", "", "The location and name of the file that you would like to output the working results to")
	flag.Parse()
	fmt.Println(input.ToArray("./test.txt"))

	bash := structs.Bash{
		WordListArray: input.ToArray(wordlist),
		URL:           url,
	}

	valid := enumeration.Enumerate(bash)
	output.Output(dest, valid)

}
