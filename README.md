# Go learning notebook

## * Pointers 
Pointers are simply memory adresses holding references to the values 


The reciver of this function is the pointer that points to instance of a person in the memory

```go
func (pointToPerson *Person) UpdateName(newName string) {
	(*pointToPerson).Name = newName
}
```

---
```go 
func (variableName *Person) 
``` 
This mean that we are working with a pointer to a person. This is also description of a type. In this case is a pointer to person struct type

---
```go
(*pointToPerson) 
```
This mean we want to manipulate the value that pointer is pointing to

// This update function won't work because we don't have access to the memory with the object - we need to use pointers
```go 
func (person Person) UpdateName(newName string) {
	person.Name = newName
}
```
---

### Value Types - Use pointers to change these things in a function 
- int 
- float 
- string
- bool 
- structs

### Reference types - Don't worry about pointers with these. When passing around Go will get the reference and copy the values but they're still referencing the same underlying like value structure in memory. 
- slices: Fun fact is when we make the slice of something than go internally is creating two separate data structures. The first is what we refer to as slice and the slice is a data structure that has three elements inside. Pointer to array head, capacity and lenght. Pointer will point to the underlying array. 
- maps
- channels
- pointers
- functions

## Maps

Maps in go are just simply dictionaries. Both keys and values are staticly typed.
```go 
	//   			key	  value
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}

	var colors2 map[string]string 
	colors2["white"] = "#fffff"

	var menu map[string]float64
	menu = map[string]float64{
    "eggs": 1.75,
    "bacon": 3.22,
    "sausage": 1.89,
	}

	colors3 := make(map[string]string)
	colors3["white"] = "#fffff"

	delete(colors, "red")
``` 
Iterating over maps and slices like so:
```go
		// key, value
	for color, hex := range colors 
	{
		// code
	}
```

---
### Interfaces

If you are a type in this program with a function called "XYZ" and you return type "string - example type" then you are now a type of Interface. Example:

```go 
type example interface {
	XYZ() string 
}


type obj struct {}

func (obj) XYZ() string {
	//...
}

```
