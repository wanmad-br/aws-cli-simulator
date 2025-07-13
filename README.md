# AWS CLI Simulator

This project simulates a simplified version of the AWS CLI, specifically focusing on the EC2 service. The primary command implemented is `aws ec2 create`, which simulates the creation of an EC2 instance.

## Project Structure

```
aws-cli-simulator
├── cmd
│   ├── ec2
│   │   └── create.go
│   └── main.go
├── internal
│   ├── ec2
│   │   └── instance.go
├── go.mod
└── README.md
```

## Getting Started

### Prerequisites

- Go 1.16 or later

### Building the Application

To build the application, navigate to the project directory and run:

```
go build -o aws-cli-simulator ./cmd/main.go
```

### Running the Application

You can run the application with the following command:

```
./aws-cli-simulator ec2 create
```

### Example Usage

To simulate the creation of an EC2 instance, use the command:

```
./aws-cli-simulator ec2 create
```

This will output a message indicating that an EC2 instance has been created.

## Contributing

Feel free to submit issues or pull requests to enhance the functionality of this simulator. 

## License

This project is licensed under the MIT License.