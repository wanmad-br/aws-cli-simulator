package vpc

import (
	"fmt"
)

// CreateVpc simulates the creation of a VPC.
func CreateVpc(vpcName string) {
	vpcID := "vpc-12345678" // Simulated VPC ID
	cidrBlock := "10.0.0.0/16" // Simulated CIDR block

	fmt.Printf("VPC Created:\n")
	fmt.Printf("  VPC ID: %s\n", vpcID)
	fmt.Printf("  VPC Name: %s\n", vpcName)
	fmt.Printf("  CIDR Block: %s\n", cidrBlock)
}