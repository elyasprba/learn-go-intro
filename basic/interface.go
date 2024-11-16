package basic

import "fmt"

type HasName interface {
	GetName() string
	GetAge() int
	GetHobby() string
}

type Person struct {
	Name  string
	Age   int
	Hobby string
}

func (p *Person) GetName() string {
	return p.Name
}

func (p *Person) GetAge() int {
	return p.Age
}

func (p *Person) GetHobby() string {
	return p.Hobby
}

func Interface(hasName HasName) {
	fmt.Println("Hello", hasName.GetName()+" you are", hasName.GetAge(), "and your hobby is", hasName.GetHobby())
}
