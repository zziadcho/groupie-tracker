package utils

type Artists struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creation_date"`
	FirstAlbum   string   `json:"first_album"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concert_dates"`
	Relations    string   `json:"relations"`
}

type Locations struct {
	Id        int    `json:"id"`
	Locations string `json:"locations"`
	Dates     string `json:"dates"`
}

type Dates struct {
	Id    int `json:"id"`
	Dates int `json:"dates"`
}

type Relations struct {
	Id             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type LocationsIndex struct {
	Locations []Locations `json:"index"`
}

type DatesIndex struct {
	Dates []Dates `json:"index"`
}

type RelationsIndex struct {
	Relations []Relations `json:"index"`
}