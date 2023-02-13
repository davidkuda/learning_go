package main

import (
	"flag"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	pw := flag.String("pw", "", "Password to hash")
	flag.Parse()
	hash := hashPW(*pw)
	if err := bcrypt.CompareHashAndPassword(hash, []byte(*pw)); err != nil {
		log.Fatal("Password do not match!")
	} else {
		fmt.Println("Passwords do match :)")
	}
}

func hashPW(s string) []byte {
	hash, err := bcrypt.GenerateFromPassword([]byte(s), 14)
	if err != nil {
		fmt.Println(err)
	}

	return hash
}
