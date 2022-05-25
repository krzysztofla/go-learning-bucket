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

func NewPerson(name string, surname string, city string, adress string) *Person {
	return &Person{
		Name:    name,
		Surname: surname,
		Contact: Contact{
			City:   city,
			Adress: adress,
		},
	}
}

// the reciver of this function is the pointer that points to instance of a person in the memory
func (pointToPerson *Person) UpdateName(newName string) {
	(*pointToPerson).Name = newName
}

// func (variableName *Person) - this mean that we are working with a pointer to a person. This is also description of a type. In this case is a pointer to person struct type
// (*pointToPerson) - this mean we want to manipulate the value that pointer is pointing to

// This update function won't work because we don't have access to the memory with the object - we need to use pointers
// func (person Person) UpdateName(newName string) {
// 	person.Name = newName
// }
