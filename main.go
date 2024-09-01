package main

import (
	"fmt"
	"log"
	"net/http"

	Groupie_tracker "groupie_tracker/handlers"
)

func main() {
	port := ":8080"
	http.HandleFunc("/", Groupie_tracker.GetDataFromJson)
	http.HandleFunc("/Artist/{id}", Groupie_tracker.HandlerShowRelation)
	http.HandleFunc("/geoMap", Groupie_tracker.GeoMap)
	http.HandleFunc("/filters/", Groupie_tracker.Filters)
	http.HandleFunc("/styles/", Groupie_tracker.HandleStyle)
	fmt.Printf("http://localhost%s", port)
	fmt.Println()
	log.Fatal(http.ListenAndServe(port, nil))
}
