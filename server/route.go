package routers

import (
	"context"
	"net/http"
	"regexp"
	"strings"

	"github.com/bdsoftpro/app/controllers/pages"
	"github.com/bdsoftpro/app/controllers/products"
	"github.com/bdsoftpro/app/controllers/users"
	"github.com/bdsoftpro/app/controllers/home"
	"github.com/bdsoftpro/app/controllers/report"
	"github.com/bdsoftpro/app/views"
)
type route struct {
	method  string
	regex   *regexp.Regexp
	handler http.HandlerFunc
}
var routes = []route{
	newRoute("GET", "/", pages.Home),
	newRoute("GET", "/login", pages.Login),
	newRoute("POST", "/login", pages.UserCheck),
	newRoute("GET", "/register", pages.Register),
	newRoute("POST", "/register", users.Store),
	newRoute("HEAD", "/logout", pages.Logout),
	newRoute("GET", "/session", pages.GetSession),
	newRoute("GET", "/ws", pages.Websoket),
	
	
	newRoute("GET", "/home/get-totals/([0-9]{4}-[0-9]{2}-[0-9]{2})/([0-9]{4}-[0-9]{2}-[0-9]{2})", homes.Totals),
	newRoute("GET", "/home/sales-payment-dues", homes.Spd),
	newRoute("GET", "/home/purchase-payment-dues", homes.Ppd),
	newRoute("GET", "/home/product-stock-alert", homes.Psa),
	
	newRoute("GET", "/reports/stock-expiry", reports.StockExpiry),
	
	newRoute("GET", "/products", products.Product),
	newRoute("GET", "/products/([^/]+)", products.ProductItem),
	
	newRoute("GET", "/products/categories", products.ProductAll),
	newRoute("GET", "/products/category/([^/]+)", products.ProductCategory),
	/*
	newRoute("GET", "/roles", roles.Index),
	newRoute("GET", "/roles/create", roles.Create),
	newRoute("POST", "/roles", roles.Store),
	newRoute("GET", "/roles/([^/]+)/edit", roles.Edit),
	newRoute("POST", "/roles/([^/]+)", roles.Update),
	newRoute("DELETE", "/roles/([^/]+)", roles.Destroy),
	*/
	
	newRoute("GET", "/users", users.Index),
	newRoute("GET", "/users/create", users.Create),
	newRoute("POST", "/users", users.Store),
	newRoute("GET", "/users/([^/]+)", users.Show),
	newRoute("GET", "/users/([^/]+)/edit", users.Edit),
	newRoute("POST", "/users/([^/]+)", users.Update),
	newRoute("DELETE", "/users/([^/]+)", users.Destroy),
	
	//newRoute("GET", "/([^/]+)", pages.Page),
}


func newRoute(method, pattern string, handler http.HandlerFunc) route {
	return route{method, regexp.MustCompile("^" + pattern + "$"), handler}
}

func Serve(w http.ResponseWriter, r *http.Request) {
	requestArray := strings.Split(r.URL.Path[1:], "/")
	if requestArray[0] == "lib" {
		requestArray[0] = "assets"
		fileName := strings.Join(requestArray, "/")
		file, err := resources.Asset(fileName)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		if strings.HasSuffix(r.URL.Path, ".css") {
			w.Header().Add("Content-Type", "text/css")
		} else if strings.HasSuffix(r.URL.Path, ".js") {
			w.Header().Add("Content-Type", "text/javascript")
		} else if strings.HasSuffix(r.URL.Path, ".svg") {
			w.Header().Add("Content-Type", "image/svg+xml")
		}
		w.Write(file)
		return
	} else if requestArray[0] == "uploads" {
		http.FileServer(http.Dir("./public")).ServeHTTP(w, r)
		return
	}
	
	var allow []string
	for _, route := range routes {
		
		matches := route.regex.FindStringSubmatch(r.URL.Path)
		if len(matches) > 0 {
			if r.Method != route.method {
				allow = append(allow, route.method)
				continue
			}
			ctx := context.WithValue(r.Context(), "fields", matches[1:])
			route.handler(w, r.WithContext(ctx))
			return
		}
	}
	if len(allow) > 0 {
		w.Header().Set("Allow", strings.Join(allow, ", "))
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.NotFound(w, r)
}