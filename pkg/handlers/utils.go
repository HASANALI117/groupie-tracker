package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
