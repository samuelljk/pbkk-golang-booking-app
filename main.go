package main

import "fmt"

func main(){
	var conferenceName = "Go-Conference"
	const conferenceTickets = 50
	var soldTickets uint = 30
	
	fmt.Println("✨Welcome to", conferenceName, "application!✨")
	fmt.Println("Get your tickets here!", "(Available:", conferenceTickets-soldTickets, "Tickets)")

	var identity [][2]string

	fmt.Printf("\n")
	fmt.Println("📄Registration Site!📄")
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	fmt.Printf("Enter the number of tickets you want to purchase: ")
	fmt.Scan(&userTickets)
	for i := uint(1); i <= userTickets; i++ {
		fmt.Printf("[%d]\n", i)
		fmt.Printf("Enter your name\t: ")
		fmt.Scan(&firstName)
		fmt.Printf("Enter your last name\t: ")
		fmt.Scan(&lastName)
		fmt.Printf("Enter your email\t: ")
		fmt.Scan(&email)
		identity = append(identity, [2]string{firstName + " " + lastName, email})
		// identity[i][0] = firstName + " " + lastName
		// identity[i][1] = email
	}
	soldTickets += userTickets
	
	fmt.Printf("\n[%s ORDER DETAILS]\n", conferenceName)
	fmt.Printf("Tickets\t: %d\n", userTickets)
	// for i := uint(1); i<= userTickets; i++
	for _, entry := range identity{
		fmt.Printf("Name\t: %s\nEmail\t: %s\n", entry[0], entry[1])
	}
}