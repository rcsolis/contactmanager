package contact

import (
	"errors"

	"github.com/rcsolis/contactmanager/internal/validation"
)

type Contact struct {
	Name        string
	Email       string
	PhoneNumber string
}

var ContactDB map[string]Contact

func init() {
	ContactDB = make(map[string]Contact)
}

func NewContact(id string, name string, email string, phoneNumber string) error {
	if validation.ValidateData(id, name, email, phoneNumber) {
		if _, exists := ContactDB[id]; exists {
			return &contactError{Err: errors.New("NewContactError"), Reason: "Contact already exists"}
		}
		ContactDB[id] = Contact{Name: name, Email: email, PhoneNumber: phoneNumber}
		return nil
	}
	return &contactError{Err: errors.New("NewContactError"), Reason: "Invalid data"}
}

func UpdateContact(id string, name string, email string, phoneNumber string) error {
	if validation.ValidateData(id, name, email, phoneNumber) {
		if _, exists := ContactDB[id]; exists {
			c := ContactDB[id]
			c.Name = name
			c.Email = email
			c.PhoneNumber = phoneNumber
			ContactDB[id] = c
			return nil
		}
		return &contactError{Err: errors.New("UpdateContactError"), Reason: "Contact not found"}
	}

	return &contactError{
		Err:    errors.New("UpdateContactError"),
		Reason: "Invalid data to update contact",
	}
}

func GetContact(id string) (*Contact, error) {
	contact, ok := ContactDB[id]
	if !ok {
		return nil, &contactError{Err: errors.New("GetContactError"), Reason: "Contact not found"}
	}
	return &contact, nil
}

func ListAllContacts() []Contact {
	var contacts []Contact
	for _, contact := range ContactDB {
		contacts = append(contacts, contact)
	}
	return contacts
}

func DeleteContact(id string) error {
	if _, ok := ContactDB[id]; !ok {
		return &contactError{Err: errors.New("DeleteContactError"), Reason: "Contact not found"}
	}
	delete(ContactDB, id)
	return nil
}
