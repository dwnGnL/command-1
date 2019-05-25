package main

import (
	"crypto/sha512"
	eb64 "encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	// "time"
)


func checkHash(password string, hashStr string) bool {
	// now := time.Now()
	// timeFormat := now.Format("02.2006")
	hash := sha512.New()
	io.WriteString(hash, "root:"+password+":23.2019")
	
	return hashStr == eb64.RawStdEncoding.EncodeToString(hash.Sum(nil))
}

func readPasswordHash() string {
	file, err := ioutil.ReadFile("./pass.txt")

	if err != nil {
		fmt.Println("couldn't read file: " + err.Error())
		return ""
	}

	return string(file)

}

