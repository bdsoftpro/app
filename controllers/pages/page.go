package pages

import (
	"fmt"
	"time"
	"log"
	"encoding/json"
	"net/http"
	"strconv"
	"html/template"
	"github.com/bdsoftpro/app/views"
	"github.com/bdsoftpro/app/system"
	"github.com/bdsoftpro/app/config"
	"github.com/bdsoftpro/app/models/pages"
	"github.com/bdsoftpro/app/models/users"
	"github.com/bdsoftpro/app/models/users/roles"
	"github.com/bdsoftpro/app/models/users/profiles"
	"github.com/bdsoftpro/app/middleware"
	
	"github.com/bdsoftpro/websocket"
)
const (
	layout = "2006-01-02T15:04:05.999999999Z07:00"
)
var (
	db, e = config.GetSqlDB()
	pdb pgmodels.DB = pgmodels.DB{Db:db}
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)
// Home Methods Defination
func Home(w http.ResponseWriter, r *http.Request) {
	auth, err := Session.AuthCheck(r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
	}
	p := struct {
		BaseUrl string
		Title string
		Body  string
		Session Session.Auth
	}{
		BaseUrl:	r.Host,
		Title:		"DashBoad Page",
		Body:		"You can login by click <strong>Login</strong> button!",
		Session:	auth,
	}
	pageTpl, _ := resources.Asset("templates/pages/dashboard.html")
	t := template.Must(template.New("dashboard.html").Parse(string(pageTpl)))
	if err := t.Execute(w, p); err != nil {
		panic(err)
	}
}
// GetSession Method Defination
func GetSession(w http.ResponseWriter, r *http.Request){
	auth, err := Session.AuthCheck(r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
	}
	pageTpl, _ := resources.Asset("templates/pages/session.html")
	t := template.Must(template.New("session.html").Parse(string(pageTpl)))
	if err := t.Execute(w, auth); err != nil {
		panic(err)
	}
}
// Page Methods Defination
func Page(w http.ResponseWriter, r *http.Request) {
	pid, _ := strconv.ParseInt(System.GetField(r, 0), 10, 64)
	page, err := pdb.Find(pid)
	if err != nil {
		pageTpl, _ := resources.Asset("templates/pages/notfound.html")
		t := template.Must(template.New("page.html").Parse(string(pageTpl)))
		item := struct {
			Name int64
		}{
			Name: pid,
		}
		if err := t.Execute(w, item); err != nil {
			panic(err)
		}
	} else {
		pageTpl, _ := resources.Asset("templates/pages/page.html")
		t := template.Must(template.New("page.html").Parse(string(pageTpl)))
		if err := t.Execute(w, page); err != nil {
			panic(err)
		}
	}
}

// UserCheck Method Defination
func UserCheck(w http.ResponseWriter, r *http.Request) {
	usr := usrmodels.User{
		Uname:	r.FormValue("uname"),
		Passwd:	r.FormValue("pass"),
	}
	udb := usrmodels.DB{
		Db: db,
	}
	user, err := udb.Login(usr)
	if err != nil {
		panic(err)
		pageTpl, _ := resources.Asset("templates/users/notfound.html")
		t := template.Must(template.New("user.html").Parse(string(pageTpl)))
		item := struct {
			Uname string
		}{
			Uname: usr.Uname,
		}
		if err := t.Execute(w, item); err != nil {
			panic(err)
		}
	} else {
		/*
		profiles, _ := profmodels.Find(user.Id)
		kvs := map[string]string{}
		for _, profile := range profiles {
			kvs[profile.Name] = profile.Detail
		}
		*/
		
		prb := profmodels.DB{
			Db: db,
		}
		pic, _ := prb.Value(user.Id, "profpic")
		rdb := rolemodels.DB{
			Db: db,
		}
		role, _ := rdb.Find(user.Id)
		
		sess := Session.Auth{
			Id:		user.Id,
			Fname:	user.Fname,
			Lname:	user.Lname,
			Gender:	user.Gender,
			Profpic:	pic,
			RoleId:		role.Id,
			RoleName:	role.Name,
			Status:	user.Status,
		}
		jsonBytes, _ :=json.Marshal(sess)
		cookie := http.Cookie{
			Name: "session",
			Value: System.Encrypt(string(jsonBytes)),
			Path: "/",
			Expires: time.Now().AddDate(0, 0, 1),
			HttpOnly: true,
			MaxAge: 1*24*60*60,
		}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

// Register Methods Defination
func Register(w http.ResponseWriter, r *http.Request) {
	p := struct {
		BaseUrl string
		Title string
		Body  string
	}{
		BaseUrl: r.Host,
		Title:	"TestPage",
		Body:	"We have created a fictional band website. Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
	}
	pageTpl, err := resources.Asset("templates/pages/register.html")
	if err != nil {
		pageTpl, _ = resources.Asset("templates/pages/notsource.html")
	}
	t := template.Must(template.New("register.html").Parse(string(pageTpl)))
	if err := t.Execute(w, p); err != nil {
		panic(err)
	}
}

// RegisterSave Methods Defination
func RegisterSave(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "register\n")
}

// Login Methods Defination
func Login(w http.ResponseWriter, r *http.Request) {
	if _, err := Session.AuthCheck(r); err == nil {
		http.Redirect(w, r, "/dashboard", http.StatusFound)
	}
	p := struct {
		BaseUrl string
		Title string
		Body  string
	}{
		BaseUrl: r.Host,
		Title:	"Login Page",
		Body:	"You can login by click <strong>Login</strong> button!",
	}
	Tpl, _ := resources.Asset("templates/pages/login.html")
	t := template.Must(template.New("login.html").Parse(string(Tpl)))
	if err := t.Execute(w, p); err != nil {
		panic(err)
	}
}

// Logout Methods Defination
func Logout(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		Expires:  time.Now(),
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)
	w.WriteHeader(http.StatusOK)
}

// Websoket Methods Defination
func Websoket(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
    ws, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println(err)
    }
    log.Println("Client Connected")
	for {
        messageType, p, err := ws.ReadMessage()
        if err != nil {
            log.Println(err)
            return
        }
        fmt.Println(string(p))
        if err := ws.WriteMessage(messageType, p); err != nil {
            log.Println(err)
            return
        }
    }
}
// PageCategory Methods Defination
func PageCategory(w http.ResponseWriter, r *http.Request) {
	page := System.GetField(r, 0)
	cat := System.GetField(r, 1)
	fmt.Fprintf(w, "Page is %s with %s Category\n", page, cat)
}
