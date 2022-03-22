package entity

type Order struct {
	Id          string
	PersonId    string
	Products    []ProductSnapshot
	Status      string
	Description string
}

func NewOrder(personId string) {

}
