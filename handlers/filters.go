package Groupie_tracker

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func Filters(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		HandleErrors(w, errors.BadRequest, errors.DescriptionBadRequest, http.StatusBadRequest)
		return
	}

	if r.URL.Path != "/filters/" {
		HandleErrors(w, errors.BadRequest, errors.DescriptionBadRequest, http.StatusBadRequest)
		return
	}

	fromYear := r.FormValue("from-year")
	toYear := r.FormValue("to-year")
	fromAlbum := r.FormValue("from-first-album")
	toAlbum := r.FormValue("to-first-album")
	LocationFilteer := r.FormValue("Location-Filter")
	members := r.Form["members"]
	if len(members) == 0 {
		members = []string{"1", "2", "3", "4", "5", "6", "7", "8"}
	}
	

	artisData, err := fromToYear(fromYear, toYear)
	if err != nil {
		HandleErrors(w, errors.InternalError, errors.DescriptionInternalError, http.StatusInternalServerError)
		return
	}
	artisData, err = Members(members, artisData)
	if err != nil {
		HandleErrors(w, errors.InternalError, errors.DescriptionInternalError, http.StatusInternalServerError)
		return
	}

	artisData, err = firstAlbum(fromAlbum, toAlbum, artisData)
	if err != nil {
		HandleErrors(w, errors.InternalError, errors.DescriptionInternalError, http.StatusInternalServerError)
		return
	}

	artisData, err = LocationFilter(LocationFilteer, artisData)
	if err != nil {
		HandleErrors(w, errors.InternalError, errors.DescriptionInternalError, http.StatusInternalServerError)
		return
	}

	// allLocation, err := locationFilter()
	// if err != nil {
	// 	HandleErrors(w, errors.BadRequest, errors.DescriptionBadRequest, http.StatusBadRequest)
	// 	return
	// }
	// artisData[0].LocationFilters = allLocation

	var buf bytes.Buffer
	if err := tmpl.ExecuteTemplate(&buf, "index.html", artisData); err != nil {
		HandleErrors(w, errors.InternalError, errors.DescriptionInternalError, http.StatusInternalServerError)
		return
	}

	if _, err := buf.WriteTo(w); err != nil {
		fmt.Printf("Error writing response: %v\n", err)
		HandleErrors(w, errors.InternalError, errors.DescriptionInternalError, http.StatusInternalServerError)
		return
	}
}

func firstAlbum(fromAlbum, toAlbum string, artisData []JsonData) ([]JsonData, error) {
	var result []JsonData

	fromAlbum = strings.ReplaceAll(fromAlbum, "-", "")
	toAlbum = strings.ReplaceAll(toAlbum, "-", "")

	fromAlbumInt, err := strconv.Atoi(fromAlbum)
	if err != nil {
		return nil, fmt.Errorf("invalid 'from' year: %v", err)
	}

	toAlbumInt, err := strconv.Atoi(toAlbum)
	if err != nil {
		return nil, fmt.Errorf("invalid 'from' year: %v", err)
	}

	for _, data := range artisData {

		t, _ := time.Parse("02-01-2006", data.FirstAlbum)
		outputDate := t.Format("2006-01-02")

		dataFirstAlbum := strings.ReplaceAll(outputDate, "-", "")
		dataFirstAlbumInt, err := strconv.Atoi(dataFirstAlbum)
		if err != nil {
			return nil, fmt.Errorf("invalid 'from' year: %v", err)
		}

		if dataFirstAlbumInt >= fromAlbumInt && dataFirstAlbumInt <= toAlbumInt {
			result = append(result, data)
		}
	}

	return result, nil
}

func fromToYear(fromYear, toYear string) ([]JsonData, error) {
	fromYearInt, err := strconv.Atoi(fromYear)
	if err != nil {
		return nil, fmt.Errorf("invalid 'from' year: %v", err)
	}

	toYearInt, err := strconv.Atoi(toYear)
	if err != nil {
		return nil, fmt.Errorf("invalid 'to' year: %v", err)
	}

	// if fromYearInt > toYearInt {
	// 	return nil, fmt.Errorf("'from' year cannot be greater than 'to' year")
	// }

	artisData, err := GetArtistsDataStruct()
	if err != nil {
		return nil, fmt.Errorf("failed to get artists data: %v", err)
	}

	var result []JsonData
	for _, data := range artisData {
		if data.CreationDate >= fromYearInt && data.CreationDate <= toYearInt {
			result = append(result, data)
		}
	}

	return result, nil
}

func Members(members []string, artisData []JsonData) ([]JsonData, error) {
	var res []int
	for _, memner := range members {
		inte, err := strconv.Atoi(memner)
		if err != nil {
			return nil, fmt.Errorf("invalid 'from' year: %v", err)
		}
		res = append(res, inte)
	}
	var result []JsonData
	for _, data := range artisData {
		for _, checkbox := range res {
			if len(data.Members) == checkbox {
				result = append(result, data)
			}
		}
	}

	return result, nil
}

func LocationFilter(LocationFilter string, artisData []JsonData) ([]JsonData, error) {
	var location Location
	var result []JsonData
	for _, data := range artisData {

		resp, err := http.Get(data.Locations)
		if err != nil {
			return nil, fmt.Errorf("invalid location: %v", err)
		}
		err = json.NewDecoder(resp.Body).Decode(&location)
		if err != nil {
			return nil, fmt.Errorf("no results the error is: %s", err)
		}
		// fmt.Println(location.Location)
		for _, f := range location.Location {
			if LocationFilter == f {
				result = append(result, data)
			}
		}
	}

	return result, nil
}
