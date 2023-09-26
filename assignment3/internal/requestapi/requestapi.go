package requestapi

import (
	"assignment_3/internal/repository/models"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func getRequest() *http.Response {
	var client *http.Client = &http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:8000/weather", nil)
	checkError(err)

	resp, err := client.Do(req)
	checkError(err)
	return resp

}

func getStringFromResp(resp *http.Response) string {
	var respBody models.Weather
	jsonByte, err := io.ReadAll(resp.Body)
	checkError(err)
	err = json.Unmarshal(jsonByte, &respBody)
	checkError(err)
	prettyJSON, err := json.MarshalIndent(respBody, "", " ")
	checkError(err)
	return string(prettyJSON)
}

func getStruct(jsonString string) *models.Weather {
	var respBody models.Weather
	jsonByte := []byte(jsonString)

	err := json.Unmarshal(jsonByte, &respBody)
	checkError(err)
	return &respBody
}

func PUTRequest() {
	var err error
	var client *http.Client = &http.Client{}
	rand.Seed(time.Now().UnixNano())

	// Generate a random integer between 1 and 100 (inclusive)
	randomNumber := rand.Intn(15) + 1
	randomNumber2 := rand.Intn(15) + 1
	requestBody := models.Weather{
		Wind:  float64(randomNumber),
		Water: float64(randomNumber2),
	}

	reqBodyByte, err := json.Marshal(requestBody)
	checkError(err)

	req, err := http.NewRequest("PUT", "http://localhost:8000/weather", bytes.NewBuffer(reqBodyByte))
	checkError(err)

	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	checkError(err)
	defer resp.Body.Close()

	resp = getRequest()
	prettyJSONSTring := getStringFromResp(resp)
	data := getStruct(prettyJSONSTring)

	fmt.Println()
	fmt.Println(prettyJSONSTring)
	fmt.Printf("water: %s\n", waterStatus(data))
	fmt.Printf("wind: %s\n", windStatus(data))

	defer resp.Body.Close()

}

func waterStatus(data *models.Weather) string {
	water := data.Water
	var status string
	if water <= 5 {
		status = "aman"
	} else if water >= 6 && water <= 8 {
		status = "siaga"
	} else if water > 8 {
		status = "bahaya"
	}
	return status
}

func windStatus(data *models.Weather) string {
	wind := data.Wind
	var status string
	if wind <= 6 {
		status = "aman"
	} else if wind >= 7 && wind <= 15 {
		status = "siaga"
	} else if wind > 15 {
		status = "bahaya"
	}
	return status

}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
