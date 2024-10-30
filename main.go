package main

import (
	"fmt"
	"strings"
)

func main(){
	var conferenceName = "Go-Conference"
	const conferenceTickets = 5
	var soldTickets uint = 0
	var remainingTickets uint = conferenceTickets - soldTickets
	
	// Welcome Message
	fmt.Printf("\n")
	fmt.Println("✨Welcome to", conferenceName, "application!✨")

	var identities [][2]string // [name, email]
	fmt.Printf("\n")

	// Registration System
	for remainingTickets > 0{
		fmt.Println("Get your tickets here!", "(Available:", conferenceTickets-soldTickets, "Tickets)")
		fmt.Printf("\n")
		fmt.Println("📄Registration Site!📄")
		var firstName, lastName, email string
		var userTickets uint
		fmt.Printf("Enter the number of tickets you want to purchase: ")
		fmt.Scan(&userTickets) // Abs() for safety, fix later
		isValidTickets := userTickets > 0 && userTickets <= remainingTickets

		if !isValidTickets{ 
			fmt.Printf("Sorry, we only have %d tickets left!\n", remainingTickets)
		} else{
			temporaryIdentities := [][2]string{}
			var isValidEmail, isValidName bool
			for i := uint(1); i <= userTickets; i++ {
				fmt.Printf("[%d]\n", i)
				fmt.Printf("Enter your name\t: ")
				fmt.Scan(&firstName)
				
				fmt.Printf("Enter your last name\t: ")
				fmt.Scan(&lastName)
				
				fmt.Printf("Enter your email\t: ")
				fmt.Scan(&email)
				
				isValidName = len(firstName) > 2 && len(lastName) > 2
				isValidEmail = strings.Contains(email, "@") && strings.Contains(email, ".")
				
				temporaryIdentities = append(temporaryIdentities, [2]string{firstName + " " + lastName, email})
			}
			if isValidName && isValidEmail{
				for _, tempIdentity := range temporaryIdentities{
					identities = append(identities, [2]string{tempIdentity[0], tempIdentity[1]})
				}
				soldTickets += userTickets
				remainingTickets = conferenceTickets - soldTickets
				
				fmt.Printf("\n[%s ORDER DETAILS]\n", conferenceName)
				fmt.Printf("Tickets\t: %d\n", userTickets)
				for _, entry := range identities{
					fmt.Printf("Name\t: %s\nEmail\t: %s\n", entry[0], entry[1])
				}
		
				fmt.Printf("\n[ADMIN ONLY]")
				firstNames := []string{}
				for _, identity := range identities{
					var names = strings.Fields(identity[0])
					firstNames = append(firstNames, names[0])
				}
				fmt.Printf("The first names of the attendees are: %v\n", firstNames)
			} else{
				fmt.Printf("\n")
				fmt.Println("⚠️Invalid name or email, please try again.⚠️")
				fmt.Printf("\n")
			}
		}

		ticketsSoldOut := remainingTickets == 0
		if ticketsSoldOut{
			fmt.Println("All tickets are sold out! Thank you for your interest!")
			break
		}
	}
}