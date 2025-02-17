package main

import (
	"crypto/sha256"
	"encoding/hex"
	"flag"
	"fmt"
	"os"
	"strings"

	f "github.com/common-nighthawk/go-figure"
	b "github.com/cyberpsyquack/pbkdf2Decode/build"
	"golang.org/x/crypto/pbkdf2"
)

// pbkdf2Hash: calculate PBKDF2 hash
func pbkdf2Hash(password, salt []byte, iterations int, derivedKeyLen int) []byte {
	derivedHash := pbkdf2.Key(password, salt, iterations, derivedKeyLen, sha256.New)
	return derivedHash
}

// findPassword: find the password matching the target hash
func findPassword(wordlist, targetHash string, salt []byte, iterations, derivedKeyLen int) string {
	targetHashBytes, _ := hex.DecodeString(targetHash)
	content, err := os.ReadFile(wordlist)
	if err != nil {
		fmt.Println("Error while loading wordlist:", err)
		return ""
	}

	lines := strings.Split(string(content), "\n")
	for i, line := range lines {
		password := strings.TrimSpace(line)
		hashValue := pbkdf2Hash([]byte(password), salt, iterations, derivedKeyLen)
		fmt.Printf("Try with: %d: %s\n", i+1, password)
		if string(hashValue) == string(targetHashBytes) {
			fmt.Printf("\nFound password: %s\n", password)
			return password
		}
	}
	fmt.Println("Password not found.")
	return ""
}

func main() {
	pbkdf2DecodeLogo := f.NewFigure("pbkdf2Decode", "", true)
	pbkdf2DecodeLogo.Print()
	fmt.Printf("\nVersion: %s\nCreator: %s\nBuild Time: %s\nBuild User: %s\n\n", b.Version, b.Creator, b.BuildTime, b.BuildUser)
	salt := flag.String("s", "", "salt")
	targetHash := flag.String("h", "", "hash")
	wordlist := flag.String("w", "/usr/share/wordlists/rockyou.txt", "wordlist")
	var usage = func() {
		fmt.Println("Help:")
		flag.PrintDefaults()
	}
	flag.Parse()
	if len(os.Args[1:]) == 0 {
		usage()
	} else {
		saltDecoded, _ := hex.DecodeString(*salt)
		findPassword(*wordlist, *targetHash, saltDecoded, 50000, 50)
	}
}
