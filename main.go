package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"path/filepath"
	"strings"

	"groupie-tracker/source"
)

type Info struct {
	ID           int      `json:"id"`
	ImageUrl     string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

type ExtraInfo struct {
	ID       int                 `json:"id"`
	Relation map[string][]string `json:"datesLocations"`
}

func getResponse(url string) io.ReadCloser {
	resp, getErr := http.Get(url)
	source.CheckError(getErr)
	if resp.StatusCode != http.StatusOK {
		source.CheckError(fmt.Errorf("bad status: %s", resp.Status))
	}
	return resp.Body
}

func fetchData() {
	artRespBody := getResponse(apiUrl + "/artists")
	defer artRespBody.Close()
	source.CheckError(json.NewDecoder(artRespBody).Decode(&Infos))

	extra, relerr := io.ReadAll(getResponse(apiUrl + "/relation"))
	source.CheckError(relerr)
	source.CheckError(json.NewDecoder(strings.NewReader(string(extra[9:]))).Decode(&ExtraInfos))
}

const apiUrl string = "https://groupietrackers.herokuapp.com/api"

var (
	Infos      []Info
	ExtraInfos []ExtraInfo
)

func renderTemplate(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "hgy 404", http.StatusNotFound)
		return
	}

	// Combine Info and ExtraInfo into a single structure for each artist
	type ArtistData struct {
		Info      Info      `json:"info"`
		ExtraInfo ExtraInfo `json:"extraInfo"`
	}

	// Prepare a slice of combined artist data
	var combinedData []ArtistData
	for i := range Infos {
		combinedData = append(combinedData, ArtistData{
			Info:      Infos[i],
			ExtraInfo: ExtraInfos[i],
		})
	}

	// Parse the template and execute it with combinedData
	tmpl := template.Must(template.ParseFiles(filepath.Join("./templates/index.html")))
	if err := tmpl.Execute(w, combinedData); err != nil {
		http.Error(w, "error executing template", http.StatusInternalServerError)
	}
}

func main() {
	fetchData()

	http.HandleFunc("/", renderTemplate)

	// Start the server on port 8080
	fmt.Println("Server starting on port 8080...")
	http.ListenAndServe(":8080", nil)
}
