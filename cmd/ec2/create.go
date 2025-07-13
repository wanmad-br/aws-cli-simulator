package ec2

import (
	"fmt"
)

// CreateInstance simulates the creation of an EC2 instance.
func CreateInstance(instanceType string) {
	instanceID := "i-1234567890abcdef0" // Simulated instance ID
	state := "running"                  // Simulated initial state

	fmt.Printf("EC2 Instance Created:\n")
	fmt.Printf("  Instance ID: %s\n", instanceID)
	fmt.Printf("  Instance Type: %s\n", instanceType)
	fmt.Printf("  State: %s\n", state)
}