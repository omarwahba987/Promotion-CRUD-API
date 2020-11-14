package entities

import "fmt"

type UserPromotion struct {
	content string
}

func (u UserPromotion) CreatePormotion() {
	fmt.Println("user promotion created")
}

func (u UserPromotion) ReadPormotion() {
	fmt.Println("user promotion read")

}

func (u UserPromotion) UptadePormotion() {
	fmt.Println("user promotion update")

}

func (u UserPromotion) DeletePormotion() {
	fmt.Println("user promotion delete")

}

func NewUserPromotion(content string) Promotion {
	return UserPromotion{content}
}
