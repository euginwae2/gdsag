package main

import (
	"crypto/md5"
	"crypto/sha256"
	"fmt"
)

func main() {
	name1 := "Richard"
	name2 := "Richards"

	mdshash := md5.Sum([]byte(name1))
	sha256hash := sha256.Sum256([]byte(name1))
	fmt.Println("   MD5: ", mdshash)
	fmt.Println("SHA256: ", sha256hash)

	mdshash = md5.Sum([]byte(name2))
	sha256hash = sha256.Sum256([]byte(name2))
	fmt.Println("   MD5: ", mdshash)
	fmt.Println("SHA256: ", sha256hash)
}
