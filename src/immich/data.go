package immich

import (
	"encoding/json"
	"fmt"
	"immich-exporter/src/models"
	"io/ioutil"
	"log"
	"net/http"
	"sync"

	"github.com/prometheus/client_golang/prometheus"
)

var wg sync.WaitGroup

func Allrequests(r *prometheus.Registry) {

	wg.Add(1)
	go ServerVersion(r)
	wg.Add(1)
	go Analyze(r)
	wg.Wait()
}

func Analyze(r *prometheus.Registry) {
	defer wg.Done()
	allusers := make(chan func() (*models.StructAllUsers, error))

	serverinfo := make(chan func() (*models.StructServerInfo, error))
	wg.Add(1)
	go GetAllUsers(allusers)
	res1, err := (<-allusers)()
	wg.Add(1)
	go ServerInfo(serverinfo)

	res2, err2 := (<-serverinfo)()

	if err != nil && err2 != nil {
	} else {
		SendBackMessagePreference(res2, res1, r)
	}
	close(serverinfo)
	close(allusers)
}

func GetAllUsers(c chan func() (*models.StructAllUsers, error)) {
	defer wg.Done()
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

			result := new(models.StructAllUsers)
			if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to go struct pointer
				log.Println("Can not unmarshal JSON")
			}

			c <- (func() (*models.StructAllUsers, error) { return result, nil })

		}
	}

}

func ServerVersion(r *prometheus.Registry) {
	defer wg.Done()
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

			var result models.StructServerVersion
			if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to go struct pointer
				log.Println("Can not unmarshal JSON")
			}

			SendBackMessageserverVersion(&result, r)

		}
	}

}

func ServerInfo(c chan func() (*models.StructServerInfo, error)) {
	defer wg.Done()
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

			result := new(models.StructServerInfo)
			if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to go struct pointer
				log.Println("Can not unmarshal JSON")
			}
			c <- (func() (*models.StructServerInfo, error) { return result, nil })

		}
	}

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
