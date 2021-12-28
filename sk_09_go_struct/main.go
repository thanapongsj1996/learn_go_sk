package main

import "fmt"

type Book struct {
	name   string
	author Author
}

type Author struct {
	name string
	age  int
}

// Method Pointer Receiver
func (b *Book) getBookNameWithHello() string {
	return fmt.Sprintf("%s hello", b.name)
}
func (a *Author) getAuthorNameWithAge() string {
	return fmt.Sprintf("%s is %d years old", a.name, a.age)
}
func (a *Author) changeNamePointer(newName string) {
	a.name = newName
}

// Method Normal Receiver
func (a Author) changeNameNormal(newName string) {
	a.name = newName
}

func main() {
	theLord := Book{
		name: "The Lord of The Rings",
		author: Author{
			name: "Thanapong Somjai",
			age:  25,
		},
	}

	fmt.Println(theLord.name)        // The Lord of The Rings
	fmt.Println(theLord.author.name) // Thanapong Somjai
	fmt.Println(theLord.author.age)  // 25

	fmt.Println(theLord.getBookNameWithHello())        // The Lord of The Rings hello
	fmt.Println(theLord.author.getAuthorNameWithAge()) // Thanapong Somjai is 25 years old

	prayuth := Author{
		name: "Prayuth",
		age:  60,
	}
	prayuth.changeNameNormal("Prawit")
	fmt.Println(prayuth.name) // Prayuth

	prayuth.changeNamePointer("Prawong")
	fmt.Println(prayuth.name) // Prawong
}
