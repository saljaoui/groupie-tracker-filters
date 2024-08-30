package Groupie_tracker

import (
	"fmt"
	"net/http"
)


func Filters(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r.URL.Path)

    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    err := r.ParseForm()
    if err != nil {
        http.Error(w, "Error parsing form", http.StatusBadRequest)
        return
    }

    fromYear := r.FormValue("from-year")
    toYear := r.FormValue("to-year")
    members := r.Form["members"]

    fmt.Println("From Year:", fromYear)
    fmt.Println("To Year:", toYear)
    fmt.Println("Members:", members)

    // Process the data as needed...

    fmt.Fprintf(w, "Filters applied successfully")
}