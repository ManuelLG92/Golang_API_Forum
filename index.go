package main

import (
	"fmt"
	"net/http"
)



func Index(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "You are in golang app!");
}
