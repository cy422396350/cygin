package models

import "github.com/cy422396350/cygin/database"

type HandlePerson interface {
	GetPerson() (person Person, err error)
}

type Person struct {
	Id       int
	Name     string
	Password string
}

func (p *Person) GetPerson() (person Person, err error) {
	err = database.SqlDB.QueryRow("SELECT id, name, password FROM person WHERE id=?", p.Id).Scan(
		&person.Id, &person.Name, &person.Password,
	)
	return
}
