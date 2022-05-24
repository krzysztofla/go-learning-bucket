package models

type Contact struct {
	City   string
	Adress string
}

type Person struct {
	Name    string
	Surname string
	Contact Contact
}

func NewPerson(name string, surname string, contact Contact) *Person {
	return &Person{
		Name:    name,
		Surname: surname,
		Contact: contact,
	}
}

func (person *Person) UpdateName(newName string) {
	(*person).Name = newName
}
