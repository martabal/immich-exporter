package models

type User struct {
	Username    string
	Password    string
	URL         string
	accessToken string
}

var myuser User

func mask(input string) string {
	hide := ""
	for i := 0; i < len(input); i++ {
		hide += "*"
	}
	return hide
}

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
	return mask(myuser.Password)
}

func SetAccessToken(accessToken string) {
	myuser.accessToken = accessToken
}

func GetAccessToken() string {
	return myuser.accessToken
}

func GetURL() string {
	return myuser.URL
}
