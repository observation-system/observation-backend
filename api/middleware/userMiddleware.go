package middleware

import (
	"os"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"net/http"
	
	"github.com/labstack/echo/v4"
	"github.com/observation-system/observation-backend/api/response"
)

func UserMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
		baseToken := c.Request().Header.Get("Authorization")
		userCheckToken, err := callAuthApi(baseToken)
		if err != nil {
			return fmt.Errorf("Invalid token")
		}

		userKey := c.Param("userKey")
		if userCheckToken.UserKey != userKey {
			return fmt.Errorf("Invalid token")
		}

		return next(c)
    }
}

func callAuthApi(baseToken string) (userCheckToken response.UserCheckToken, err error) {
	host := os.Getenv("AUTH_API_HOST")
	url := fmt.Sprintf("%s/user/user_check", host)
	request, _ := http.NewRequest("GET", url, nil)
    request.Header.Set("Authorization", baseToken)

	client := &http.Client{}
	res, err := client.Do(request)
	if err != nil {
		return userCheckToken, fmt.Errorf("Invalid token")
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return userCheckToken, fmt.Errorf("Invalid token")
	}

	json.Unmarshal(body, &userCheckToken)
	
	return userCheckToken, nil
}
