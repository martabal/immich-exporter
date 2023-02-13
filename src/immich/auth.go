package immich

import (
	"encoding/json"
	"fmt"
	"immich-exporter/src/models"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func Auth() {

	url := models.GetURL() + "/api/auth/login"
	method := "POST"

	payload := strings.NewReader(`{
  	"email": "` + models.GetUsername() + `",
  	"password": "` + models.Getpassword() + `"}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	} else {
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Accept", "application/json")

		res, err := client.Do(req)
		if err != nil {
			log.Println(err)
			return
		} else {
			defer res.Body.Close()

			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				log.Println(err)
				return
			}

			if res.StatusCode == 400 {
				log.Fatalln("Incorrect login")
			} else {
				var result models.StructLogin
				if err := json.Unmarshal(body, &result); err != nil {
					log.Println("Can not unmarshal JSON")
				}

				models.SetAccessToken(result.AccessToken)
			}

		}

	}

}
