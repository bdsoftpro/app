package products

import (
	"fmt"
	"strconv"
	"net/http"
	"html/template"
	"github.com/bdsoftpro/app/views"
	"github.com/bdsoftpro/app/config"
	"github.com/bdsoftpro/app/models/products"
)

/*
index()			--  view all	-- get
create()		--	create view	-- get
store()			--  save		-- post
show($id)		--	single view	-- get
edit($id)		--  edit view	-- get
update($id)		--  update		-- post
destroy($id)	--  delete		-- post, get

+===========================================================================================+
| Verb		| URI									| Action	| Route Name				|
+===========+=======================================+===========+===========================+
| GET		| /products								| index		| products.index			|
+-----------+---------------------------------------+-----------+---------------------------+
| GET		| /products/create						| create	| products.create			|
+-----------+---------------------------------------+-----------+---------------------------+
| POST		| /products								| store		| products.store			|
+-----------+---------------------------------------+-----------+---------------------------+
| GET		| /products/{product}					| show		| products.show				|
+-----------+---------------------------------------+-----------+---------------------------+
| GET		| /products/{product}/edit				| edit		| products.edit				|
+-----------+---------------------------------------+-----------+---------------------------+
| PUT/PATCH	| /products/{product}					| update	| products.update			|
+-----------+---------------------------------------+-----------+---------------------------+
| DELETE	| /products/{product}					| destroy	| products.destroy			|
+-----------+---------------------------------------+-----------+---------------------------+
+-----------+---------------------------------------+---------------------------------------+
| GET		| /products/categories					| index		| products.category.index	|
+-----------+---------------------------------------+-----------+---------------------------+
| GET		| /products/category/create				| create	| products.category.create	|
+-----------+---------------------------------------+-----------+---------------------------+
| POST		| /products/category					| store		| users.comments.store		|
+-----------+---------------------------------------+-----------+---------------------------+
| GET		| /products/category/{category}			| show		| comments.show				|
+-----------+---------------------------------------+-----------+---------------------------+
| GET		| /products/category/{category}/edit	| edit		| comments.edit				|
+-----------+---------------------------------------+-----------+---------------------------+
| PUT/PATCH	| /products/category/{category}			| update	| comments.update			|
+-----------+---------------------------------------+-----------+---------------------------+
| DELETE	| /products/category/{category}			| destroy	| comments.destroy			|
+===========================================================================================+
*/
const (
	layout = "2006-01-02T15:04:05.999999999Z07:00"
)
var (
	db, e = config.GetSqlDB()
	pdb pdmodels.DB = pdmodels.DB{Db:db}
)

func GetField(r *http.Request, index int) string{
	field := r.Context().Value("fields").([]string)
	return field[index]
}

// Product Methods Defination
func Product(w http.ResponseWriter, r *http.Request) {
	page := GetField(r, 0)
	p := struct {
		Name string
	}{
		Name: page,
	}
	pageTpl, _ := resources.Asset("templates/pages/page.html")
	t := template.Must(template.New("page.html").Parse(string(pageTpl)))
	if err := t.Execute(w, p); err != nil {
		panic(err)
	}
}

// ProductItem Methods Defination
func ProductItem(w http.ResponseWriter, r *http.Request) {
	pid, _ := strconv.ParseInt(GetField(r, 0), 10, 64)
	product, err := pdb.Find(pid)
	if err != nil {
		pageTpl, _ := resources.Asset("templates/products/notfound.html")
		t := template.Must(template.New("product.html").Parse(string(pageTpl)))
		item := struct {
			Name int64
		}{
			Name: pid,
		}
		if err := t.Execute(w, item); err != nil {
			panic(err)
		}
	} else {
		pageTpl, _ := resources.Asset("templates/products/product.html")
		t := template.Must(template.New("product.html").Parse(string(pageTpl)))
		if err := t.Execute(w, product); err != nil {
			panic(err)
		}
	}
}
// ProductAll Methods Defination
func ProductAll(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Product All\n")
}

// ProductCategory Methods Defination
func ProductCategory(w http.ResponseWriter, r *http.Request) {
	page := GetField(r, 0)
	cat := GetField(r, 1)
	fmt.Fprintf(w, "Product is %s with %s Category\n", page, cat)
}