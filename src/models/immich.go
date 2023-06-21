package models

type StructImmich struct {
	APIKey string
	URL    string
}

var myuser StructImmich

func Setuser(url string, apikey string) {
	myuser.URL = url
	myuser.APIKey = apikey

}

func Getbaseurl() string {
	return myuser.URL
}

func GetApiKey() string {
	return myuser.APIKey
}
