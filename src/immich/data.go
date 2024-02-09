package immich

import (
	"encoding/json"
	"fmt"
	"immich-exp/logger"
	"immich-exp/models"
	"io"
	"strconv"

	API "immich-exp/api"
	"net/http"
	"sync"

	prom "immich-exp/prometheus"

	"github.com/prometheus/client_golang/prometheus"
)

var wg sync.WaitGroup

var (
	mutex sync.Mutex
)

type Data struct {
	URL        string
	HTTPMethod string
}

var httpGetUsers = Data{
	URL:        "/api/user?isAll=true",
	HTTPMethod: http.MethodGet,
}
var httpServerVersion = Data{
	URL:        "/api/server-info/version",
	HTTPMethod: http.MethodGet,
}
var httpStatistics = Data{
	URL:        "/api/server-info/statistics",
	HTTPMethod: http.MethodGet,
}
var httpGetJobs = Data{
	URL:        "/api/jobs",
	HTTPMethod: http.MethodGet,
}

var unmarshalError = "Can not unmarshal JSON"

func Allrequests(r *prometheus.Registry) {

	wg.Add(1)
	go ServerVersion(r)
	wg.Add(1)
	go Analyze(r)
	wg.Wait()
}

func Analyze(r *prometheus.Registry) {
	defer wg.Done()

	alljobsstatus := make(chan func() (*API.StructAllJobsStatus, error))
	allusers := make(chan func() (*API.StructAllUsers, error))
	serverinfo := make(chan func() (*API.StructServerInfo, error))
	defer func() {
		close(serverinfo)
		close(allusers)
		close(alljobsstatus)
	}()

	wg.Add(1)
	go GetAllJobsStatus(alljobsstatus)
	res1, err := (<-alljobsstatus)()
	wg.Add(1)
	go GetAllUsers(allusers)
	res2, err2 := (<-allusers)()
	wg.Add(1)
	go ServerInfo(serverinfo)

	res3, err3 := (<-serverinfo)()

	if err == nil && err2 == nil && err3 == nil {
		prom.SendBackMessagePreference(res3, res2, res1, r)
	}
}

func GetAllUsers(c chan func() (*API.StructAllUsers, error)) {
	defer wg.Done()
	resp, err := Apirequest(httpGetUsers.URL, httpGetUsers.HTTPMethod)
	if err == nil {

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		} else {

			result := new(API.StructAllUsers)
			if err := json.Unmarshal(body, &result); err != nil {
				logger.Log.Error(unmarshalError)
			}

			c <- (func() (*API.StructAllUsers, error) { return result, nil })
			return
		}
	}
	c <- (func() (*API.StructAllUsers, error) { return new(API.StructAllUsers), err })
}

func ServerVersion(r *prometheus.Registry) {
	defer wg.Done()
	resp, err := Apirequest(httpServerVersion.URL, httpServerVersion.HTTPMethod)
	if err == nil {

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		} else {

			var result API.StructServerVersion
			if err := json.Unmarshal(body, &result); err != nil {
				logger.Log.Error(unmarshalError)
			}

			prom.SendBackMessageserverVersion(&result, r)
		}
	}
}

func ServerInfo(c chan func() (*API.StructServerInfo, error)) {
	defer wg.Done()
	resp, err := Apirequest(httpStatistics.URL, httpStatistics.HTTPMethod)
	if err == nil {

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		} else {

			result := new(API.StructServerInfo)
			if err := json.Unmarshal(body, &result); err != nil {
				logger.Log.Error(unmarshalError)
			}
			c <- (func() (*API.StructServerInfo, error) { return result, nil })
			return
		}
	}
	c <- (func() (*API.StructServerInfo, error) { return new(API.StructServerInfo), err })
}

func GetAllJobsStatus(c chan func() (*API.StructAllJobsStatus, error)) {
	defer wg.Done()
	resp, err := Apirequest(httpGetJobs.URL, httpGetJobs.HTTPMethod)
	if err == nil {

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		} else {

			result := new(API.StructAllJobsStatus)
			if err := json.Unmarshal(body, &result); err != nil {
				logger.Log.Error(unmarshalError)
			}
			c <- (func() (*API.StructAllJobsStatus, error) { return result, nil })
			return
		}
	}
	c <- (func() (*API.StructAllJobsStatus, error) { return new(API.StructAllJobsStatus), err })
}

func Apirequest(uri string, method string) (*http.Response, error) {

	req, err := http.NewRequest(method, models.Getbaseurl()+uri, nil)
	if err != nil {
		panic("Error with url")
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("x-api-key", models.GetApiKey())
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		err := fmt.Errorf("Can't connect to server")
		mutex.Lock()
		if !models.GetPromptError() {
			logger.Log.Error(err.Error())
			models.SetPromptError(true)
		}
		mutex.Unlock()
		return resp, err

	}
	switch resp.StatusCode {
	case http.StatusOK:
		mutex.Lock()
		if models.GetPromptError() {
			models.SetPromptError(false)
		}
		mutex.Unlock()
		return resp, nil
	case http.StatusNotFound:
		errormessage := "Error code " + strconv.Itoa(resp.StatusCode) + " for " + models.Getbaseurl() + uri
		panic(errormessage)

	case http.StatusUnauthorized, http.StatusForbidden:
		panic("Api key unauthorized")

	default:
		err := fmt.Errorf("%d", resp.StatusCode)
		mutex.Lock()
		if !models.GetPromptError() {
			models.SetPromptError(true)
			logger.Log.Debug("Error code " + strconv.Itoa(resp.StatusCode))
		}
		mutex.Unlock()
		return resp, err
	}

}
