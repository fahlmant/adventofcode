package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"regexp"
)

func main() {

	secretKey := "bgvyzdsv"
	hashNumber := 0
	regex := regexp.MustCompile("\\A00000")

	for {
		hashNumber++
		hash := md5.New()
		io.WriteString(hash, secretKey)
		io.WriteString(hash, fmt.Sprintf("%d", hashNumber))
		result := fmt.Sprintf("%x", hash.Sum(nil))

		if regex.MatchString(result) {
			fmt.Println(result)
			fmt.Println(hashNumber)
			os.Exit(0)
		}
	}

}
