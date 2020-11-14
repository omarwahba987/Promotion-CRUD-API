package service

import (
	"crudPromotion/entities"
	"fmt"
	"net/http"
)

func createCompanyPromotion(w http.ResponseWriter, r *http.Request){

	companyPromotion := entities.NewCompanyPromotion(r.FormValue("Promotion"))
	fmt.Fprintf(w, "done")
	companyPromotion.CreatePormotion()

}
func readCompanyPromotion(w http.ResponseWriter, r *http.Request){


	fmt.Fprintf(w, "Company Promotion data")

}
func updateCompanyPromotion(w http.ResponseWriter, r *http.Request){

	companyPromotion := entities.NewCompanyPromotion(r.FormValue("Promotion"))
	fmt.Fprintf(w, "done")
	companyPromotion.UptadePormotion()

}
func deleteCompanyPromotion(w http.ResponseWriter, r *http.Request){

	companyPromotion := entities.NewCompanyPromotion(r.FormValue("Promotion"))
	fmt.Fprintf(w, "done")
	companyPromotion.DeletePormotion()

}