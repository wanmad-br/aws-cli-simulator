package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"aws-cli-simulator/cmd/ec2"
)

func main() {
	// Lista de comandos disponíveis
	availableCommands := []string{
		"acm", "acm-pca", "alexaforbusiness", "amplify",
		"apigateway", "apigatewaymanagementapi", "apigatewayv2",
		"appconfig", "appflow", "ec2", // Adicione outros serviços aqui
	}

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

		// Remove newline e espaços extras
		input = strings.TrimSpace(input)

		// Condição de saída
		if input == "exit" {
			fmt.Println("Exiting AWS CLI Simulator...")
			break
		}

		// Dividir entrada em argumentos
		args := strings.Split(input, " ")
		if len(args) < 2 || args[0] != "aws" {
			fmt.Println(`usage: aws [options] <command> <subcommand> [<subcommand> ...] [parameters]
To see help text, you can run:

  aws help
  aws <command> help
  aws <command> <subcommand> help

aws: error: too few arguments`)
			continue
		}

		// Verificar se o comando é válido
		command := args[1]
		isValidCommand := false
		for _, availableCommand := range availableCommands {
			if command == availableCommand {
				isValidCommand = true
				break
			}
		}

		if !isValidCommand {
			fmt.Printf(`usage: aws [options] <command> <subcommand> [<subcommand> ...] [parameters]

The AWS Command Line Interface is a unified tool to manage your AWS services.

To see help text, you can run:

  aws help
  aws <command> help
  aws <command> <subcommand> help

Available services:

  acm                        | acm-pca
  alexaforbusiness           | amplify
  apigateway                 | apigatewaymanagementapi
  apigatewayv2               | appconfig
  appflow                    | ec2
  (lista continua com todos os serviços disponíveis)

aws: error: argument command: Invalid choice: '%s' (choose from '%s')
`, command, strings.Join(availableCommands, "', '"))
			continue
		}

		// Caso "aws ec2" seja digitado sem subcomando
		if len(args) == 2 && command == "ec2" {
			fmt.Println(`usage: aws ec2 <command> [<args>]

The AWS EC2 service provides APIs for managing EC2 resources.

To see help text, you can run:

  aws ec2 help
  aws ec2 <command> help

aws: error: too few arguments`)
			continue
		}

		// Caso "aws ec2 create" seja digitado sem argumentos adicionais
		if len(args) < 3 {
			fmt.Println("Usage: aws ec2 create <instance-type>")
			continue
		}

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