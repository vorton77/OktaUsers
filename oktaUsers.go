package oktaUsers

import (
	"fmt"
	"strings"
	"net/http"
	"io/ioutil"
	"time"
	"encoding/json"
)

type DataLinks struct{
	Suspend struct{
		Href string `json:"href"`
		Method string `json:"method"`
	}
	ResetPassword struct{
		Href string `json:"href"`
		Method string `json:"method"`
	}
	ExpirePassword struct{
		Href string `json:"href"`
		Method string `json:"method"`
	}
	Self struct{
		Href string `json:"href"`
	}
	ChangeRecoveryQuestion struct{
		Href string `json:"href"`
		Method string `json:"method"`
	}
	Deactivate struct{
		Href string `json:"href"`
		Method string `json:"method"`
	}
	ChangePassword struct{
		Href string `json:"href"`
		Method string `json:"method"`
	}
}

type ProfileData struct{
	Login string `json:"login"`
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
	MobilePhone string `json:"mobilePhone"`
	Email string `json:"email"`
	SecondEmail string `json:"secondEmail"`
}

type CredentialData struct{
	Password struct {}
	Provider struct{
		Type string `json:"type"`
		Name string `json:"name"`
	}
}

type RegResponse struct {
	ID        string `json:"id"`
	Status    string `json:"status"`
	Created   time.Time `json:"created"`
	Activated time.Time `json:"activated"`
	StatusChanged time.Time `json:"statusChanged"`
	LastLogin time.Time `json:"lastLogin"`
	LastUpdated time.Time `json:"lastUpdated"`
	PasswordChanged time.Time `json:"passwordChanged"`
	Profile ProfileData `json:"profile"`
	Credentials CredentialData `json:"credentials"`
	Links DataLinks `json:"_links"`
}

func CreateUserNoCreds( oktaOrg string,
			apiKey string,
			firstName string,
			lastName string,
			email string,
			login string,
			activate string ) (RegResponse) {

	fmt.Println("Enter createUserNoCreds():")

	key := "SSWS " + apiKey

	fmt.Println(key)

	url := "https://" + oktaOrg + "/api/v1/users?activate=" + activate

	payload := strings.NewReader("{\n  \"profile\": {\n    \"firstName\": \"" + firstName + "\",\n    \"lastName\": \"" + lastName + "\",\n    \"email\": \"" + email + "\",\n    \"login\": \"" + login + "\"\n  }\n}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("authorization", key)
	req.Header.Add("cache-control", "no-cache")

	fmt.Println("Before POST:")
	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

	// declare a variable of type RegResponse
	var regResponse RegResponse

	// unmarshal the []byte body which represents the json returned in the response

	json.Unmarshal(body, &regResponse)

	fmt.Println("Before return:")
	return regResponse
}

func CreateUserWithCreds(oktaOrg string,
			 apiKey string,
			 firstName string,
			 lastName string,
			 email string,
			 login string,
			 password string,
			 activate string ) (RegResponse) {

	fmt.Println("Enter createUserWithCreds(): using new struck")

	key := "SSWS " + apiKey

	fmt.Println(key)

	url := "https://" + oktaOrg + "/api/v1/users?activate=" + activate

	payload := strings.NewReader("{\n  \"profile\": {\n    \"firstName\": \"" + firstName + "\",\n    \"lastName\": \"" + lastName + "\",\n    \"email\": \"" + email + "\",\n    \"login\": \"" + login + "\"\n  },\n  \"credentials\": {\n    \"password\" : { \"value\": \"" + password + "\" }\n  }\n}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("authorization", key)
	req.Header.Add("cache-control", "no-cache")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

	// declare a variable of type RegResponse
	var regResponse RegResponse

	// unmarshal the []byte body which represents the json returned in the response

	json.Unmarshal(body, &regResponse)

	fmt.Println("Before return:")
	return regResponse
}





