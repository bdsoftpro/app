package users
import (
	"fmt"
	"regexp"
	"strconv"
	"net/http"
	"html/template"
	"time"
	
	"github.com/bdsoftpro/app/system"
	"github.com/bdsoftpro/app/config"
	"github.com/bdsoftpro/app/models/users"
	"github.com/bdsoftpro/app/views"
)
/*
index()			--  view all	-- get
create()		--	create view	-- get
store()			--  save		-- post
show($id)		--	single view	-- get
edit($id)		--  edit view	-- get
update($id)		--  update		-- post
destroy($id)	--  delete		-- post, get

+===============================================================================+
| Verb		| URI							| Action	| Route Name			|
+===========+===============================+===========+=======================+
| GET		| /users						| index		| users.index			|
+-----------+-------------------------------+-----------+-----------------------+
| GET		| /users/create					| create	| users.create			|
+-------------------------------------------------------------------------------+
| POST		| /users						| store		| users.store			|
+-------------------------------------------------------------------------------+
| GET		| /users/{user}					| show		| users.show			|
+-------------------------------------------------------------------------------+
| GET		| /users/{user}/edit			| edit		| users.edit			|
+-------------------------------------------------------------------------------+
| PUT/PATCH	| /users/{user}					| update	| users.update			|
+-------------------------------------------------------------------------------+
| DELETE	| /users/{user}					| destroy	| users.destroy			|
+-------------------------------------------------------------------------------+
+-------------------------------------------------------------------------------+
| GET		| /users/{user}/comments		| index		| users.comments.index	|
+-------------------------------------------------------------------------------+
| GET		| /users/{user}/comments/create	| create	| users.comments.create	|
+-------------------------------------------------------------------------------+
| POST		| /users/{user}/comments		| store		| users.comments.store	|
+-------------------------------------------------------------------------------+
| GET		| /comments/{comment}			| show		| comments.show			|
+-------------------------------------------------------------------------------+
| GET		| /comments/{comment}/edit		| edit		| comments.edit			|
+-------------------------------------------------------------------------------+
| PUT/PATCH	| /comments/{comment}			| update	| comments.update		|
+-------------------------------------------------------------------------------+
| DELETE	| /comments/{comment}			| destroy	| comments.destroy		|
+===============================================================================+
*/
const (
	layout = "2006-01-02T15:04:05.999999999Z07:00"
)
var (
	udb usrmodels.DB = func()(usrmodels.DB){
		db, _ := config.GetSqlDB()
		return usrmodels.DB{Db:db}
	}()
)
// Index Method Defination
func Index(w http.ResponseWriter, r *http.Request) {
	users, err := udb.FindAll()
	if err != nil {
		pageTpl, _ := resources.Asset("templates/users/notfound.html")
		t := template.Must(template.New("user.html").Parse(string(pageTpl)))
		item := struct {
			Name string
		}{
			Name: "User",
		}
		if err := t.Execute(w, item); err != nil {
			panic(err)
		}
	} else {
		/*
		header, _ := resources.Asset("templates/users/header.html")
		t, _ := template.New("header").Parse(header)
		footer, _ := resources.Asset("templates/users/footer.html")
		t.New("footer").Parse(footer)
		layout, _ := resources.Asset("templates/users/layout.html")
		t.New("layout").Parse(layout)
		t.ExecuteTemplate(w, "layout", users)
		*/
		pageTpl, _ := resources.Asset("templates/users/index.html")
		t := template.Must(template.New("index.html").Parse(string(pageTpl)))
		if err := t.Execute(w, users); err != nil {
			panic(err)
		}
	}
}

// Create Method Defination
func Create(w http.ResponseWriter, r *http.Request) {
	p := struct {
		Title string
		Body  string
	}{
		Title:	"TestPage",
		Body:	"We have created a fictional band website. Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
	}
	pageTpl, err := resources.Asset("templates/pages/create.html")
	if err != nil {
		pageTpl, _ = resources.Asset("templates/pages/notsource.html")
	}
	t := template.Must(template.New("create.html").Parse(string(pageTpl)))
	if err := t.Execute(w, p); err != nil {
		panic(err)
	}
}

// Store Method Defination
func Store(w http.ResponseWriter, r *http.Request) {
	usr := usrmodels.User{
		Fname:	r.FormValue("fname"),
		Lname:	r.FormValue("lname"),
		Uname:	r.FormValue("uname"),
		Passwd:	System.Encrypt(r.FormValue("pass")),
		Gender:	r.FormValue("gender"),
		Email:	r.FormValue("email"),
		Status:	1,
		Created_at: time.Now().UTC().Format(layout),
	}
	err := udb.Create(&usr)
	if err != nil {
		dpKey := regexp.MustCompile(`^Error [0-9]+: Duplicate entry '([^/]+)' for key '([^/]+)'.*$`).FindStringSubmatch(err.Error())
		pageTpl, _ := resources.Asset("templates/users/userexist.html")
		t := template.Must(template.New("user.html").Parse(string(pageTpl)))
		item := struct {
			Field string
			Value string
		}{
			Field:	dpKey[1:][1],
			Value:	dpKey[1:][0],
		}
		if err := t.Execute(w, item); err != nil {
			panic(err)
		}
	} else {
		http.Redirect(w, r, "/users", http.StatusFound)
	}
}

// Show Method Defination
func Show(w http.ResponseWriter, r *http.Request) {
	pid, _ := strconv.ParseInt(System.GetField(r, 0), 10, 64)
	user, err := udb.Find(pid)
	if err != nil {
		pageTpl, _ := resources.Asset("templates/users/notfound.html")
		t := template.Must(template.New("user.html").Parse(string(pageTpl)))
		item := struct {
			Name int64
		}{
			Name: pid,
		}
		if err := t.Execute(w, item); err != nil {
			panic(err)
		}
	} else {
		pageTpl, _ := resources.Asset("templates/users/user.html")
		t := template.Must(template.New("user.html").Parse(string(pageTpl)))
		if err := t.Execute(w, user); err != nil {
			panic(err)
		}
	}
}

// Edit Method Defination
func Edit(w http.ResponseWriter, r *http.Request) {
	pid, _ := strconv.ParseInt(System.GetField(r, 0), 10, 64)
	user, err := udb.Find(pid)
	if err != nil {
		pageTpl, _ := resources.Asset("templates/users/notfound.html")
		t := template.Must(template.New("user.html").Parse(string(pageTpl)))
		item := struct {
			Name int64
		}{
			Name: pid,
		}
		if err := t.Execute(w, item); err != nil {
			panic(err)
		}
	} else {
		pageTpl, _ := resources.Asset("templates/users/edit.html")
		t := template.Must(template.New("edit.html").Parse(string(pageTpl)))
		
		type param struct {
			Id int64
		}
		
		data := struct {
			User	usrmodels.User
			Param	param
			Title	string
		}{
			User:	user,
			Param:	param{
				Id:	pid,
			},
			Title:	"User Edit",
		}
		if err := t.Execute(w, data); err != nil {
			panic(err)
		}
	}
}

// Update Method Defination
func Update(w http.ResponseWriter, r *http.Request) {
	pid, _ := strconv.ParseInt(System.GetField(r, 0), 10, 64)
	stus, _ := strconv.ParseInt(r.FormValue("status"), 10, 8)
	usr := usrmodels.User{
		Id:		pid,
		Fname:	r.FormValue("fname"),
		Lname:	r.FormValue("lname"),
		Email:	r.FormValue("email"),
		Status:	int8(stus),
		Updated_at: time.Now().UTC().Format(layout),
	}
	_, err := udb.Update(usr)
	if err != nil {
		pageTpl, _ := resources.Asset("templates/users/notfound.html")
		t := template.Must(template.New("user.html").Parse(string(pageTpl)))
		item := struct {
			Name int64
		}{
			Name: pid,
		}
		if err := t.Execute(w, item); err != nil {
			panic(err)
		}
	} else {
		http.Redirect(w, r, fmt.Sprintf("/users/%d", pid), http.StatusFound)
	}
}

// Destroy Method Defination
func Destroy(w http.ResponseWriter, r *http.Request) {
	pid, _ := strconv.ParseInt(System.GetField(r, 0), 10, 64)
	user, err := udb.Find(pid)
	if err != nil {
		pageTpl, _ := resources.Asset("templates/users/notfound.html")
		t := template.Must(template.New("user.html").Parse(string(pageTpl)))
		item := struct {
			Name int64
		}{
			Name: pid,
		}
		if err := t.Execute(w, item); err != nil {
			panic(err)
		}
	} else {
		pageTpl, _ := resources.Asset("templates/users/user.html")
		t := template.Must(template.New("user.html").Parse(string(pageTpl)))
		if err := t.Execute(w, user); err != nil {
			panic(err)
		}
	}
}