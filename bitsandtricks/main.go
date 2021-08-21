package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

const (
	// this is the file permissions const
	// read this to know more -> https://phoenixnap.com/kb/linux-file-permissions
	filePermission = 0644
	// file name itself here
	// it is a const 'cause we must not change that in any case (otherwise we will list the saved data)
	fileName = "user_data.bin"
)

var (
	// it is a maid data part we will work with
	// 8 booleans in one small uint8
	data uint8

	// ErrNotExist describe the error we use in a case file do not exist
	// It is very similar to what do we have in the "os" package and just an example how can we go
	// with custom errors:
	ErrNotExist = errors.New("file does not exist")
)

func main() {
	// greetings!
	fmt.Println("hello, dear user (^_^)")
	readedData, err := readData(fileName)
	if errors.Is(ErrNotExist, err) {
		fmt.Println("make an input please:")
		if err := updateData(); err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Println("found saved state!")
		data = readedData
	}

	fmt.Println("the data is:", data)
	for {
		var operation string
		_, err := fmt.Scanln(&operation)
		if err != nil {
			log.Fatal(err)
		}

		switch operation {
		case "update":
			if err := updateData(); err != nil {
				log.Fatal(err)
			}

			if err := saveData(data); err != nil {
				log.Fatal(err)
			}

		case "show binary representation":
		// TODO: show the value in bin format
		case "bitwise AND":
		// TODO: ask user for another value and execute & on that
		case "bitwise OR":
		// TODO: ask user for another value and execute | on that
		case "bitwise XOR":
		// TODO: ask user for another value and execute ^ on that
		case "clear":
		// TODO: delete user's file
		case "exit":
			saveData(data)
			fmt.Println("buy buy!")
			os.Exit(0)
		}

	}
}

// here we do update a global variable "data".
// not really a good way - but better keep updating global state in one function,
// not in main()
func updateData() error {
	value, err := askUserInput()
	if err != nil {
		return err
	}

	data = value
	return nil
}

func askUserInput() (uint8, error) {
	var theSeed string
	_, err := fmt.Scanln(&theSeed)
	if err != nil {
		return 0, err
	}

	seed, err := strconv.ParseUint(theSeed, 10, 8)
	if err != nil {
		return 0, err
	}

	return uint8(seed), nil
}

func saveData(inputData uint8) error {
	return ioutil.WriteFile(fileName, []byte{inputData}, filePermission)
}

func readData(file string) (uint8, error) {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return 0, ErrNotExist
	}

	info, err := ioutil.ReadFile(fileName)
	if err != nil {
		return 0, err
	}

	return info[0], nil
}
