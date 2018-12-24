package group

import (
	"github.com/vvanpo/titian/contact"
)

type Group struct {
	name        string
	description string
	contacts    map[*contact.Contact]struct{}
}

// Name returns the name of the group.
func (g Group) Name() string {
	return g.name
}

func (g *Group) SetName(name string) error {
	g.name = name
	return nil
}

// Description returns an optional description of the group.
func (g Group) Description() string {
	return g.description
}

func (g *Group) SetDescription(description string) error {
	g.description = description
	return nil
}

// Members returns an unordered list of all contacts belonging to the group.
func (g Group) Members() []*contact.Contact {
	contacts := make([]*contact.Contact, 0, len(g.contacts))

	for contact := range g.contacts {
		contacts = append(contacts, contact)
	}

	return contacts
}

func (g *Group) AddContact(contact *contact.Contact) {
	g.contacts[contact] = struct{}{}
}

func (g *Group) RemoveContact(contact *contact.Contact) {
	delete(g.contacts, contact)
}
