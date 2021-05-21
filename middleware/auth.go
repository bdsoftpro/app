package Session

import (
	"errors"
	"encoding/json"
	"net/http"
	"github.com/bdsoftpro/app/system"
	"unsafe"
)

// Auth Format
type Auth struct {
	Id		int64	`json:"id"`
	Fname	string	`json:"fname"`
	Lname	string	`json:"lname"`
	Gender	string	`json:"gender"`
	Profpic	string	`json:"profpic"`
	RoleId	int64	`json:"roleid"`
	RoleName	string	`json:"rolename"`
	Status	int8	`json:"status"`
}

// AuthCheck Method Defination
func AuthCheck(r *http.Request) (Auth, error){
	cookie, err := r.Cookie("session")
	if err != nil {
		return Auth{}, err
	}
	var auth Auth
	json.Unmarshal([]byte(System.Decrypt(cookie.Value)), &auth)
	if unsafe.Sizeof(auth) == 0 {
		return Auth{}, errors.New("It has no session")
	}
	return auth, nil
}
