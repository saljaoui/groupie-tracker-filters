package Groupie_tracker

import (
	"bytes"
	"fmt"
	"net/http"
	"strconv"
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

	// if err := r.ParseForm(); err != nil {
	// 	handleError(w, "Failed to parse form", http.StatusBadRequest)
	// 	return
	// }

	fromYear := r.FormValue("from-year")
	toYear := r.FormValue("to-year")
	members := r.Form["members"]

	artisData, err := fromToYear(fromYear, toYear)
	if err != nil {
		HandleErrors(w, errors.InternalError, errors.DescriptionInternalError, http.StatusInternalServerError)
		return
	}
	artisData , err = Members(members, artisData)
	if err != nil {
		HandleErrors(w, errors.InternalError, errors.DescriptionInternalError, http.StatusInternalServerError)
		return
	}




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
