package titian

import (
	"testing"
)

// Tests that a newly-created contact is in the contact list and has a non-zero
// identifier.
func TestNewContact(t *testing.T) {
	var zeroId id
	contact, err := stubContact()

	if err != nil {
		t.Fatalf("NewContact method returned error: %v", err)
	}

	if contact.id == zeroId {
		t.Fatal("Contact does not have initialized ID")
	}

	if contact.app.contacts.fetch(contact.id) != contact {
		t.Errorf("Cannot fetch contact ID: %v", contact.id)
	}
}

//
func TestDeleteContact(t *testing.T) {
	contact, _ := stubContact()
	id := contact.id
	contacts := contact.app.contacts
	contacts.Delete(contact)

	if contacts.fetch(id) != nil {
		t.Error("Failed to delete contact")
	}
}
