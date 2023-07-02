package main

import (
	"os"
	"fmt"
	"log"
	"bytes"
	"strconv"
	"strings"
	"io/ioutil"
	"encoding/json"
	"net/http"

	batchLog "github.com/observation-system/observation-backend/log"
	"github.com/observation-system/observation-backend/domain"
	"github.com/observation-system/observation-backend/infrastructure"
	"github.com/observation-system/observation-backend/cmd/controller"
	"github.com/observation-system/observation-backend/api/response"
)

type Spot struct {
	SpotKey string `json:"spot_key"`
	Url     string `json:"url"`
}

type SpotsResponse struct {
	SpotKey string `json:"spot_key"`
	Count   int    `json:"count"`
}

func main() {
	file := batchLog.GeneratBatchLog()
	log.SetOutput(file)

	spotController := controller.NewSpotController(infrastructure.NewSqlHandler())
	dayController := controller.NewDayController(infrastructure.NewSqlHandler())

	data, err := spotController.Spots()
	if err != nil {
		log.Println(response.NewError(err))
	}
	
	var spotResponse []SpotsResponse
	body := callDetectionApi(data)
	json.Unmarshal(body, &spotResponse)

	for i := 0; i < len(spotResponse); i++ {
		spotKey := spotResponse[i].SpotKey
		spotsCount := strconv.Itoa(spotResponse[i].Count)
		data, err := spotController.Spot(spotKey)
		if err != nil {
			log.Println(response.NewError(err))
		}

		spotsCountDay := data[0].CountDay
		spotsCountDayArray := strings.Split(spotsCountDay, ",")

		if len(spotsCountDayArray) >= 24 {
			dayController.UpdateDayRecord(spotKey, spotsCountDay)
			spotsCountDayArray[0] = spotsCount
			spotsCountDayArray = spotsCountDayArray[:0]
		}

		if spotsCountDay == "None" {
			spotsCountDayArray[0] = spotsCount
		} else {
			spotsCountDayArray = append(spotsCountDayArray, spotsCount)
		}

		spotDomain := domain.Spot{}
		spotDomain.Count = spotsCount
		spotsCountDayString := strings.Join(spotsCountDayArray, ",")
		spotDomain.CountDay = spotsCountDayString

		spotController.UpdateSpot(spotKey, spotDomain)
	}
}

func callDetectionApi(data domain.Spots) []byte {
	var spots []Spot

	for i := 0; i < len(data); i++ {
		json := Spot{
			SpotKey: data[i].SpotKey,
			Url: data[i].Url,
		}
		spots = append(spots, json)
	}

	host := os.Getenv("DETECTION_API_HOST")
	url := fmt.Sprintf("%s/batch_detection", host)
	headers := map[string]string{"Content-Type": "application/json"}
	jsonStr, _ := json.Marshal(spots)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(response.NewError(err))
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	
	return body
}
