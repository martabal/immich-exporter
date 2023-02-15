package models

type StructImmich struct {
	Username    string
	Password    string
	URL         string
	AccessToken string
}

var myuser StructImmich

func Getuser() (string, string) {
	return myuser.Username, myuser.Password
}

func Setuser(username string, password string, url string) {
	myuser.Username = username
	myuser.Password = password
	myuser.URL = url
}

func GetUsername() string {
	return myuser.Username
}

func Getpassword() string {
	return myuser.Password
}
func Getpasswordmasked() string {
	hide := ""
	for i := 0; i < len(myuser.Password); i++ {
		hide += "*"
	}
	return hide
}

func SetAccessToken(accessToken string) {
	myuser.AccessToken = accessToken
}

func GetAccessToken() string {
	return myuser.AccessToken
}

func GetURL() string {
	return myuser.URL
}
