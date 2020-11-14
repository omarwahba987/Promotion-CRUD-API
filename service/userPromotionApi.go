package service

import (
	"crudPromotion/entities"
	"fmt"
	"net/http"
)

func createUserPromotion(w http.ResponseWriter, r *http.Request) {

	userPromotion := entities.NewUserPromotion(r.FormValue("Promotion"))
	fmt.Fprintf(w, "done")
	userPromotion.CreatePormotion()

}
func readUserPromotion(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "User Promotion data")

}
func updateUserPromotion(w http.ResponseWriter, r *http.Request) {

	userPromotion := entities.NewUserPromotion(r.FormValue("Promotion"))
	fmt.Fprintf(w, "done")
	userPromotion.UptadePormotion()

}
func deleteUserPromotion(w http.ResponseWriter, r *http.Request) {

	userPromotion := entities.NewUserPromotion(r.FormValue("Promotion"))
	fmt.Fprintf(w, "done")
	userPromotion.DeletePormotion()

}
