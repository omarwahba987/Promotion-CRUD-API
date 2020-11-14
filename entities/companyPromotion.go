package entities

import "fmt"

type CompanyPromotion struct {
	content string
}

func (c CompanyPromotion) CreatePormotion() {
	fmt.Println("Company Promotion  created")

}

func (c CompanyPromotion) ReadPormotion() {
	fmt.Println("Company Promotion  read")

}

func (c CompanyPromotion) UptadePormotion() {
	fmt.Println("Company Promotion  update")

}

func (c CompanyPromotion) DeletePormotion() {
	fmt.Println("Company Promotion  delete")

}

func NewCompanyPromotion(content string) Promotion {
	return CompanyPromotion{content}
}
