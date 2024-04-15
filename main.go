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

var T1 = [9][8]rune{}
var T2 = [9][8]rune{}
var T3 = [9][8]rune{}
var T4 = [9][8]rune{}

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

	Tables()

	a := rand.Intn(36)
	b := rand.Intn(36)
	c := rand.Intn(36)
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

	Tables()

	decryptedMessage := Decipher(message)

	fmt.Println("Decrypted message:", decryptedMessage)
}

func Tables() {
	rows := 9
	cols := 8

	count := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			T1[i][j] = '0' + rune(count) // ASCII values from '0' to 'z'
			count++
		}
	}

	// Populate T2: Right to left, top to bottom
	count = 0
	for i := 0; i < rows; i++ {
		for j := cols - 1; j >= 0; j-- {
			T2[i][j] = '0' + rune(count)
			count++
		}
	}

	// Populate T3: Right to left, bottom to top
	count = 0
	for i := rows - 1; i >= 0; i-- {
		for j := cols - 1; j >= 0; j-- {
			T3[i][j] = '0' + rune(count)
			count++
		}
	}

	// Populate T4: Left to right, bottom to top
	count = 0
	for i := rows - 1; i >= 0; i-- {
		for j := 0; j < cols; j++ {
			T4[i][j] = '0' + rune(count)
			count++
		}
	}

	/*
		// Print tables
		fmt.Println("Table 1:")
		for i := 0; i < rows; i++ {
			for j := 0; j < cols; j++ {
				fmt.Printf("%c ", T1[i][j])
			}
			fmt.Println()
		}

		fmt.Println("Table 2:")
		for i := 0; i < rows; i++ {
			for j := 0; j < cols; j++ {
				fmt.Printf("%c ", T2[i][j])
			}
			fmt.Println()
		}

		fmt.Println("Table 3:")
		for i := 0; i < rows; i++ {
			for j := 0; j < cols; j++ {
				fmt.Printf("%c ", T3[i][j])
			}
			fmt.Println()
		}

		fmt.Println("Table 4:")
		for i := 0; i < rows; i++ {
			for j := 0; j < cols; j++ {
				fmt.Printf("%c ", T4[i][j])
			}
			fmt.Println()
		}
	*/
}

func Cipher(a, b, c int, message string) string {
	a = (a-1)%36 + 1
	b = (b-1)%36 + 1
	c = (c-1)%36 + 1
	d := a + b + c
	j := d % 4

	// Pass through each table sequentially
	for i := 0; i < 4; i++ {
		message = processTable(b, c, message, j)
		j = (j + 1) % 4
	}

	// Append the values of a, b, c, and d to the encrypted message
	encryptedMessage := fmt.Sprintf("%02d%02d%02d%03d", a, b, c, d) + message

	return encryptedMessage
}

func Decipher(encryptedMessage string) string {
	// Extract values of a, b, c and d
	a, _ := strconv.Atoi(encryptedMessage[:2])
	b, _ := strconv.Atoi(encryptedMessage[2:4])
	c, _ := strconv.Atoi(encryptedMessage[4:6])
	d, _ := strconv.Atoi(encryptedMessage[6:9])

	// Ensure values are within range 1 to 36
	a = (a-1)%36 + 1
	b = (b-1)%36 + 1
	c = (c-1)%36 + 1
	d = (d-1)%108 + 1

	message := encryptedMessage[9:]
	decryptedMessage := ""

	// Calculate the value of j based on d
	j := d % 4

	// Pass through each table sequentially in reverse order
	for i := 0; i < 4; i++ {
		decryptedMessage = reverseProcessTable(b, c, message, j)
		j = (j - 1) % 4
		if j == -1 {
			j = 3
		}
	}

	return decryptedMessage
}

func processTable(b, c int, message string, tableIndex int) string {
	var table [9][8]rune

	// Choose the appropriate table
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

	// Initialize the encrypted message with the original message
	encryptedMessage := message

	// Process each character in the message
	for _, char := range message {
		// Get the row and column index for the current character
		row := (b + (int(char) - 48)) % 9
		col := (c + (int(char) - 48)) % 8

		// Append the corresponding character from the selected table to the encrypted message
		encryptedMessage += string(table[row][col])
	}

	encryptedMessage = encryptedMessage[len(message):]

	return encryptedMessage
}

func reverseProcessTable(b, c int, encryptedMessage string, tableIndex int) string {
	var table [9][8]rune

	// Choose the appropriate table
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

	// Initialize the decrypted message
	decryptedMessage := ""

	// Process each character in the message
	for _, char := range encryptedMessage {
		// Get the row and column index for the current character
		row, col := findPositionInTable(table, char, b, c)

		// Ensure row and col are within range
		row = (row + 9) % 9
		col = (col + 8) % 8

		// Append the corresponding character from the selected table to the decrypted message
		decryptedMessage += string(table[row][col])
	}

	return decryptedMessage
}

func findPositionInTable(table [9][8]rune, char rune, b, c int) (int, int) {
	// Find the character's position in the table
	for i := 0; i < 9; i++ {
		for j := 0; j < 8; j++ {
			if table[i][j] == char {
				// Calculate the original character's position in the table
				row := (i - b + 9) % 9
				col := (j - c + 8) % 8

				return row, col
			}
		}
	}

	return -1, -1
}
