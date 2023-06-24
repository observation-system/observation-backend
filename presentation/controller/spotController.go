package controller

import (
	"fmt"
	"time"
	"strings"
	"strconv"
	"io/ioutil"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	
	"github.com/observation-system/observation-backend/lib"
	"github.com/observation-system/observation-backend/domain"
	"github.com/observation-system/observation-backend/usecase"
	"github.com/observation-system/observation-backend/presentation/database"
	"github.com/observation-system/observation-backend/presentation/response"
)

type SpotController struct {
	Interactor usecase.SpotInteractor
}

func NewSpotController(sqlHandler database.SqlHandler) *SpotController {
	return &SpotController{
		Interactor: usecase.SpotInteractor{
			SpotRepository: &database.SpotRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

// Register
// @Summary     地点登録
// @tags        Spot
// @Accept      json
// @Produce     json
// @Param       name     body string true "地点名"
// @Param       url      body string true "URL"
// @Param       address  body string true "住所"
// @Param       user_key path int    true "ユーザーキー"
// @Success     200  {object} response.SpotRegister
// @Failure     500  {array}  response.Error
// @Router      /user/:userKey/spot_register [post]
func (controller *SpotController) Register(c echo.Context) (err error) {
	u := domain.Spot{}
	c.Bind(&u)

	userKey := c.Param("userKey")
	spotKey, err := lib.GenerateKey()
	if err != nil {
		return c.JSON(500, response.NewError(err))
	}

	u.UserKey = userKey
	u.SpotKey = spotKey

	currentTime := time.Now()
	hour := currentTime.Hour()
	zeroString := strings.Repeat(",0", hour - 1)
	u.CountDay = "0" + zeroString
	u.Count = "0"

	latitude, longitude, err := callGeocodingApi(u.Address)
	if err != nil {
		return c.JSON(500, response.NewError(err))
	}

	u.Latitude = latitude
	u.Longitude = longitude

	spot, err := controller.Interactor.Add(u)
	if err != nil {
		return c.JSON(500, response.NewError(err))
	}

	return c.JSON(200, response.ToSpotRegister(spot))
}

func callGeocodingApi(address string) (latitude string, longitude string, err error) {
	url := fmt.Sprintf("https://msearch.gsi.go.jp/address-search/AddressSearch?q=%s", address)
	request, _ := http.NewRequest("GET", url, nil)

	client := &http.Client{}
	res, err := client.Do(request)
	if err != nil {
		return latitude, longitude, fmt.Errorf("Failed to get")
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return latitude, longitude, fmt.Errorf("Failed to get")
	}

	var features []response.Feature
	json.Unmarshal(body, &features)

	for _, feature := range features {
		latitude = strconv.FormatFloat(feature.Geometry.Coordinates[1], 'f', -1, 64)
		longitude = strconv.FormatFloat(feature.Geometry.Coordinates[0], 'f', -1, 64)
	}
	
	return latitude, longitude, nil
}
