package immich

import (
	"encoding/json"
	"fmt"
	"immich-exporter/src/models"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
)

func Allrequests(r *prometheus.Registry) {

	serverversion(r)

	Analyze(r)
}

func Analyze(r *prometheus.Registry) {
	allusers, err := GetAllUsers()
	users, err2 := users()
	if err != nil && err2 != nil {
	} else {
		Sendbackmessagepreference(users, allusers, r)
	}

}

func GetAllUsers() (*models.AllUsers, error) {
	resp, err := Apirequest("/api/user?isAll=true", "GET")
	if err != nil {
		if err.Error() == "403" {
			log.Println("Cookie changed, try to reconnect ...")
			Auth()
		} else {
			if models.GetPromptError() == false {
				log.Println("Error : ", err)
			}
		}
	} else {
		if models.GetPromptError() == true {
			models.SetPromptError(false)
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		} else {

			var result models.AllUsers
			if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to go struct pointer
				log.Println("Can not unmarshal JSON")
			}

			return &result, nil

		}
	}
	return &models.AllUsers{}, err
}

func serverversion(r *prometheus.Registry) {
	resp, err := Apirequest("/api/server-info/version", "GET")
	if err != nil {
		if err.Error() == "403" {
			log.Println("Cookie changed, try to reconnect ...")
			Auth()
		} else {
			if models.GetPromptError() == false {
				log.Println("Error : ", err)
			}
		}
	} else {
		if models.GetPromptError() == true {
			models.SetPromptError(false)
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		} else {

			var result models.ServerVersion
			if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to go struct pointer
				log.Println("Can not unmarshal JSON")
			}

			Sendbackmessageserverversion(&result, r)

		}
	}

}

func users() (*models.Users, error) {
	resp, err := Apirequest("/api/server-info/stats", "GET")
	if err != nil {
		if err.Error() == "403" {
			log.Println("Cookie changed, try to reconnect ...")
			Auth()
		} else {
			if models.GetPromptError() == false {
				log.Println("Error : ", err)
			}
		}

	} else {
		if models.GetPromptError() == true {
			models.SetPromptError(false)
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		} else {

			var result models.Users
			if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to go struct pointer
				log.Println("Can not unmarshal JSON")
			}

			return &result, nil

		}
	}
	return &models.Users{}, err
}

func Apirequest(uri string, method string) (*http.Response, error) {

	req, err := http.NewRequest(method, models.GetURL()+uri, nil)
	if err != nil {
		log.Fatalln("Error with url")
	}

	req.AddCookie(&http.Cookie{Name: "immich_access_token", Value: models.GetAccessToken()})
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		err := fmt.Errorf("Can't connect to server")
		if models.GetPromptError() == false {
			log.Println(err.Error())
			models.SetPromptError(true)
		}

		return resp, err

	} else {
		models.SetPromptError(false)
		if resp.StatusCode == 200 {

			return resp, nil

		} else {
			err := fmt.Errorf("%d", resp.StatusCode)
			if models.GetPromptError() == false {
				models.SetPromptError(true)

				log.Println("Error code", err.Error())

			}
			return resp, err

		}

	}

}
