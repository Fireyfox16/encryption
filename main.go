package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

var T1 = [95]rune{}
var T2 = [95]rune{}
var T3 = [95]rune{}
var T4 = [95]rune{}

func main() {
	fmt.Println("Would you like to encrypt or decrypt a message?")
	fmt.Println("1. Encrypt")
	fmt.Println("2. Decrypt")
	fmt.Println("3. Exit")
	var choice int
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		Encrypt()
	case 2:
		Decrypt()
	case 3:
		fmt.Println("Exiting...")
	default:
		fmt.Println("Invalid choice")
	}
}

func Encrypt() {
	fmt.Println("Enter the message to encrypt:")
	reader := bufio.NewReader(os.Stdin)
	message, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Invalid message")
	}
	message = strings.TrimSpace(message)

	a := rand.Intn(72) + 1
	b := rand.Intn(72) + 1
	c := rand.Intn(72) + 1

	encryptedMessage := Cipher(a, b, c, message)

	fmt.Println("Encrypted message:", encryptedMessage)
}

func Decrypt() {
	fmt.Println("Enter the message to decrypt:")
	reader := bufio.NewReader(os.Stdin)
	message, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Invalid message")
	}
	message = strings.TrimSpace(message)

	decryptedMessage := Decipher(message)

	fmt.Println("Decrypted message:", decryptedMessage)
}

func Tables(a int) {
	count := a
	for i := 0; i < 95; i++ {
		T1[i] = ' ' + rune(count%95)
		if T1[i] < ' ' {
			T1[i] += ' '
		}
		count++
	}

	count = a + a
	for i := 0; i < 95; i++ {
		T2[i] = ' ' + rune(count%95)
		if T2[i] < ' ' {
			T2[i] += ' '
		}
		count++
	}

	count = a * a
	for i := 0; i < 95; i++ {
		T3[i] = ' ' + rune(count%95)
		if T3[i] < ' ' {
			T3[i] += ' '
		}
		count++
	}

	count = (a + a) * a
	for i := 0; i < 95; i++ {
		T4[i] = ' ' + rune(count%95)
		if T4[i] < ' ' {
			T4[i] += ' '
		}
		count++
	}
}

func Cipher(a, b, c int, message string) string {
	d := (a * b) + (a * c)
	j := d % 4

	Tables(a)

	for i := 0; i < 4; i++ {
		message = processTable(b, c, message, j)
		j = (j + 1) % 4
	}

	encryptedMessage := fmt.Sprintf("%02d%02d%02d%s", a, b, c, message)

	return encryptedMessage
}

func Decipher(encryptedMessage string) string {
	a, _ := strconv.Atoi(encryptedMessage[:2])
	b, _ := strconv.Atoi(encryptedMessage[2:4])
	c, _ := strconv.Atoi(encryptedMessage[4:6])

	message := encryptedMessage[6:]

	d := (a * b) + (a * c)
	j := d % 4
	j--

	Tables(a)

	for i := 0; i < 4; i++ {
		if j < 0 {
			j = 3
		}
		message = reverseProcessTable(b, c, message, j)
		j--
	}

	return message
}

func processTable(b, c int, message string, tableIndex int) string {
	var table [95]rune

	switch tableIndex {
	case 0:
		table = T1
	case 1:
		table = T2
	case 2:
		table = T3
	case 3:
		table = T4
	}

	encryptedMessage := ""

	for _, char := range message {
		index := ((int(char) - 32 + b + c) + 95) % 95

		encryptedMessage += string(table[index])

		b = (b + 1) % 95
	}

	return encryptedMessage
}

func reverseProcessTable(b, c int, encryptedMessage string, tableIndex int) string {
	var table [95]rune

	switch tableIndex {
	case 0:
		table = T1
	case 1:
		table = T2
	case 2:
		table = T3
	case 3:
		table = T4
	}

	decryptedMessage := ""

	for _, char := range encryptedMessage {
		index := -1
		for i, tableChar := range table {
			if tableChar == char {
				index = i
				break
			}
		}
		if index == -1 {
			decryptedMessage += string(char)
		} else {
			index = (index - b - c) % 95
			if index < 0 {
				index += 95
			}
			decryptedMessage += string(rune(index + 32))
		}

		b = (b + 1) % 95
	}

	return decryptedMessage
}
