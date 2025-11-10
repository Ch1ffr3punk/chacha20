package main

import (
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"golang.org/x/crypto/chacha20"
)

func readHexFile(filename string) ([]byte, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	hexString := strings.TrimSpace(string(content))
	decoded, err := hex.DecodeString(hexString)
	if err != nil {
		return nil, fmt.Errorf("error decoding hex: %v", err)
	}

	return decoded, nil
}

func printUsage() {
	fmt.Fprintf(os.Stderr, "Usage: <keyfile> <noncefile> < infile > outfile\n")
}

func main() {
	if len(os.Args) != 3 {
		printUsage()
		os.Exit(1)
	}

	keyFile := os.Args[1]
	nonceFile := os.Args[2]

	key, err := readHexFile(keyFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading key file: %v\n", err)
		os.Exit(1)
	}

	nonce, err := readHexFile(nonceFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading nonce file: %v\n", err)
		os.Exit(1)
	}

	if len(key) != chacha20.KeySize {
		fmt.Fprintf(os.Stderr, "Error: Invalid key size. Expected %d hex bytes, got %d hex bytes\n", chacha20.KeySize, len(key))
		os.Exit(1)
	}

	if len(nonce) != chacha20.NonceSize {
		fmt.Fprintf(os.Stderr, "Error: Invalid nonce size. Expected %d hex bytes), got %d hex bytes\n", chacha20.NonceSize, len(nonce))
		os.Exit(1)
	}

	cipher, err := chacha20.NewUnauthenticatedCipher(key, nonce)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating cipher: %v\n", err)
		os.Exit(1)
	}

	buf := make([]byte, 8192) // read/write buffer
	for {
		n, err := os.Stdin.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
			os.Exit(1)
		}
		if n == 0 {
			break
		}

		cipher.XORKeyStream(buf[:n], buf[:n])
		
		if _, err := os.Stdout.Write(buf[:n]); err != nil {
			fmt.Fprintf(os.Stderr, "Error writing output: %v\n", err)
			os.Exit(1)
		}

		if err == io.EOF {
			break
		}
	}
}
