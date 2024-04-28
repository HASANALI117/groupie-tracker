package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func fetchData(urlPath string, data interface{}) error {
	client := &http.Client{}
	request, err := http.NewRequest("GET", "https://groupietrackers.herokuapp.com/api/"+urlPath, nil)
	if err != nil {
		return err
	}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Error making request:", err)
		return err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return err
	}

	err = json.Unmarshal(body, data)
	if err != nil {
		fmt.Println("Error unmarshalling response:", err)
		return err
	}
	return nil
}

func fetchAndHandleData(path string, data interface{}, errorMsg string, w http.ResponseWriter) error {
    err := fetchData(path, data)
    if err != nil {
        log.Printf("Error fetching %s: %v", errorMsg, err) // Log the error
        InternalServerError(w, fmt.Sprintf("Error fetching %s", errorMsg))
        return err
    }
    return nil
}