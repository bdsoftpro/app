package homes

import (
	"fmt"
	"encoding/json"
	"net/http"
	"github.com/bdsoftpro/app/middleware"
)
type data struct {
	Id int64 `json:"id"`
	Number int64 `json:"number"`
	Amount int64 `json:"amount"`
}
// Totals Method Defination
func Totals(w http.ResponseWriter, r *http.Request) {
	_, err := Session.AuthCheck(r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
	}
	// start := System.GetField(r, 0)
	// end := System.GetField(r, 1)
	p := struct{
			Tpit	int64	`json:"total_purchase_inc_tax"`
			Tpet	int64	`json:"total_purchase_exc_tax"`
			Pd	int64	`json:"purchase_due"`
			Tsc	int64	`json:"total_shipping_charges"`
			Tp	int64	`json:"total_purchase"`
			Id	int64	`json:"invoice_due"`
			Te	int64	`json:"total_expense"`
		}{
			Tpit:	0,
			Tpet:	0,
			Pd:		0,
			Tsc:	0,
			Tp:		0,
			Id:		0,
			Te:		0,
		}
		
	b, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Fprint(w, string(b))
}
// Spd Method Defination
func Spd(w http.ResponseWriter, r *http.Request) {
	_, err := Session.AuthCheck(r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
	}
	datas := []data{}
	p := struct{
			Draw			int64		`json:"draw"`
			RecordsTotal	int64		`json:"recordsTotal"`
			RecordsFiltered	int64		`json:"recordsFiltered"`
			Data	[]data	`json:"data"`
		}{
			Draw:	1,
			RecordsTotal:	0,
			RecordsFiltered:	0,
			Data:	datas,
		}
		
	b, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Fprint(w, string(b))
}

// Ppd Method Defination
func Ppd(w http.ResponseWriter, r *http.Request) {
	_, err := Session.AuthCheck(r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
	}
	datas := []data{}
	p := struct{
			Draw			int64		`json:"draw"`
			RecordsTotal	int64		`json:"recordsTotal"`
			RecordsFiltered	int64		`json:"recordsFiltered"`
			Data	[]data	`json:"data"`
		}{
			Draw:	1,
			RecordsTotal:	0,
			RecordsFiltered:	0,
			Data:	datas,
		}
		
	b, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Fprint(w, string(b))
}

// Psa Method Defination
func Psa(w http.ResponseWriter, r *http.Request) {
	_, err := Session.AuthCheck(r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
	}
	datas := []data{}
	p := struct{
			Draw			int64		`json:"draw"`
			RecordsTotal	int64		`json:"recordsTotal"`
			RecordsFiltered	int64		`json:"recordsFiltered"`
			Data	[]data	`json:"data"`
		}{
			Draw:	1,
			RecordsTotal:	0,
			RecordsFiltered:	0,
			Data:	datas,
		}
		
	b, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Fprint(w, string(b))
}