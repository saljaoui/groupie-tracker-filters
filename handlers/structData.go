package Groupie_tracker

// structs for artist jeson data
type JsonData struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
	LocationFilters []string 
}

// struct for artist jeson data
type Artist struct {
	ID             int                 `json:"id"`
	Image          string              `json:"image"`
	Name           string              `json:"name"`
	Members        []string            `json:"members"`
	DatesLocations map[string][]string `json:"datesLocations"`
	CreationDate   int                 `json:"creationDate"`
	FirstAlbum     string              `json:"firstAlbum"`
	Date           []string            `json:"dates"`
	Location       []string
	Url            []string
}

// struct for Location jeson data
type Location struct {
	Location []string `json:"locations"`
}

// struct for Date jeson data
type Date struct {
	Date []string `json:"dates"`
}

// struct for Relation jeson data
type Relation struct {
	DatesLocations map[string][]string `json:"datesLocations"`
}

// struct for Errors
type Errors struct {
	Message     string
	Code        int
	Description string
}

type GeocodeResponse struct {
	Lat string `json:"lat"`
	Lon string `json:"lon"`
}

// struct for AllMessageErrors jeson data
type AllMessageErrors struct {
	NotFound                    string `json:"notfound"`
	BadRequest                  string `json:"badrequest"`
	InternalError               string `json:"internalerror"`
	MethodNotAllowed            string `json:"methodnotallowed"`
	DescriptionNotFound         string `json:"description_notfound"`
	DescriptionBadRequest       string `json:"description_badrequest"`
	DescriptionInternalError    string `json:"description_internalerror"`
	DescriptionMethodNotAllowed string `json:"description_methodnotallowed"`
}
