package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"aws-cli-simulator/cmd/ec2"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("AWS CLI Simulator")
	fmt.Println("Type 'exit' to quit")

	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		// Remove newline and extra spaces
		input = strings.TrimSpace(input)

		// Exit condition
		if input == "exit" {
			fmt.Println("Exiting AWS CLI Simulator...")
			break
		}

		// Split input into arguments
		args := strings.Split(input, " ")
		if len(args) < 3 {
			fmt.Println("Usage: aws ec2 create <instance-type>")
			continue
		}

		command := args[0]
		subCommand := args[1]
		action := args[2]

		if command == "aws" && subCommand == "ec2" && action == "create" {
			if len(args) < 4 {
				fmt.Println("Usage: aws ec2 create <instance-type>")
				continue
			}
			instanceType := args[3]
			ec2.CreateInstance(instanceType)
		} else {
			fmt.Println("Unknown command")
		}
	}
}