// Created by: haokai
// Created on: May 2021

// This program displays, "address_Input"

package main

import "fmt"

func main() {
	// This function does addition
	var streetName string
	var streetNumber int
	// input
	fmt.Println("This program gets a user's street number and name.")
	fmt.Println()
	fmt.Print("Enter your street Number: ")
	fmt.Scanln(&streetNumber)
	fmt.Print("Enter your street Name: ")
	fmt.Scanln(&streetName)
	// output
	fmt.Println("Your street address is: ", streetNumber, " ", streetName, ".")
	fmt.Println("\nDone.")
	}