package service

import (
	"log"
	"net/http"
)

func HandleRequests() {
	http.HandleFunc("/createUserPromotion", createUserPromotion)
	http.HandleFunc("/readUserPromotion", readUserPromotion)
	http.HandleFunc("/updateUserPromotion", updateUserPromotion)
	http.HandleFunc("/deleteUserPromotion", deleteUserPromotion)
	http.HandleFunc("/createCompanyPromotion", createCompanyPromotion)
	http.HandleFunc("/readCompanyPromotion", readCompanyPromotion)
	http.HandleFunc("/updateCompanyPromotion", updateCompanyPromotion)
	http.HandleFunc("/deleteCompanyPromotion", deleteCompanyPromotion)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
