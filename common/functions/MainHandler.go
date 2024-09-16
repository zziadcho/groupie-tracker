package functions

import (
	//"01/groupie-tracker/common/utils"
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
	"path/filepath"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	response, err := http.Get("https://groupietrackers.herokuapp.com/api")
	if err != nil {
		log.Fatal(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(w, err.Error(), http.StatusInternalServerError)
		return
	}

	/////////////////////////////////////////////////////////////////////

	var responseObject map[string]interface{}
	if err := json.Unmarshal(responseData, &responseObject); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// for key, value := range responseObject {
	
	// }

	tmpl := template.Must(template.ParseFiles(filepath.Join("./common/static/index.html")))
	if err := tmpl.Execute(w, responseObject); err != nil {
		http.Error(w, "error executing the template", http.StatusInternalServerError)
	}

	// response, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	// defer response.Body.Close()

	// responseData, err := io.ReadAll(response.Body)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// var responseObject []Artists
	// if err := json.Unmarshal(responseData, &responseObject); err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// tmpl := template.Must(template.ParseFiles(filepath.Join("./common/static/index.html")))
	// if err := tmpl.Execute(w, responseObject); err != nil {
	// 	http.Error(w, "error executing template", http.StatusInternalServerError)
	// }
}
