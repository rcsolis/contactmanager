package main

import (
	"fmt"
	"log"

	"github.com/rcsolis/contactmanager/internal/contact"
)

func menuOptions() {
	fmt.Println("Select an option:")
	fmt.Println("1. Add contact")
	fmt.Println("2. Get contact")
	fmt.Println("3. List all contacts")
	fmt.Println("4. Delete contact")
	fmt.Println("5. Update Contact")
	fmt.Println("6. Exit")
}

func createNewContact() {
	var id, name, email, phoneNumber string
	log.Println("--- Create New Contact ---")
	fmt.Println("Enter id:")
	fmt.Scanln(&id)
	fmt.Println("Enter name:")
	fmt.Scanln(&name)
	fmt.Println("Enter email:")
	fmt.Scanln(&email)
	fmt.Println("Enter phone number:")
	fmt.Scanln(&phoneNumber)
	err := contact.NewContact(id, name, email, phoneNumber)
	if err != nil {
		log.Print(err)
	}
}

func getContact() {
	var id string
	log.Println("--- Get Contact ---")
	fmt.Println("Enter the id:")
	fmt.Scanln(&id)
	c, err := contact.GetContact(id)
	if err != nil {
		log.Print(err.Error())
	}
	fmt.Println("Contact details:")
	fmt.Println("Name:", c.Name)
	fmt.Println("Email:", c.Email)
	fmt.Println("Phone number:", c.PhoneNumber)
}

func listAllContacts() {
	contacts := contact.ListAllContacts()
	log.Print("--- Listing all contacts ---")
	for _, c := range contacts {
		fmt.Println("Name:", c.Name)
		fmt.Println("Email:", c.Email)
		fmt.Println("Phone number:", c.PhoneNumber)
	}
}

func deleteContact() {
	var id string
	log.Println("--- Delete Contact ---")
	fmt.Println("Enter the id:")
	fmt.Scanln(&id)
	err := contact.DeleteContact(id)
	if err != nil {
		log.Print(err.Error())
	}
	log.Print("Contact deleted")
}

func updateContact() {
	var id, name, email, phoneNumber string
	log.Println("--- Update Contact ---")
	fmt.Println("Enter the contact id to update:")
	fmt.Scanln(&id)
	fmt.Println("Enter new name:")
	fmt.Scanln(&name)
	fmt.Println("Enter new email:")
	fmt.Scanln(&email)
	fmt.Println("Enter new phone number:")
	fmt.Scanln(&phoneNumber)
	err := contact.UpdateContact(id, name, email, phoneNumber)
	if err != nil {
		log.Print(err.Error())
	}
	log.Print("Contact updated")
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Print("Recover from: ", r)
		}
	}()
	// Loop to show options menu
	var option int
	loop := true
	for loop {
		// Print options menu
		menuOptions()
		// Read user input
		fmt.Scanln(&option)
		switch option {
		case 1:
			createNewContact()
		case 2:
			getContact()
		case 3:
			listAllContacts()
		case 4:
			deleteContact()
		case 5:
			updateContact()
		case 6:
			loop = false
		}
	}
	log.Print("Exiting...")
}
