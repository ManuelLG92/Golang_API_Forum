package handlers

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// Init get id by URL
func getIdByUrl(r *http.Request)  int {
	vars := mux.Vars(r)
	paramId, errConv := strconv.Atoi(vars["id"])
	if errConv != nil {
		fmt.Println("The ID couldn't been catch")
		return 0
	}
	return paramId
}
// End get id by URL
