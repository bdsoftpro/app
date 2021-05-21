package reports

import (
	"fmt"
	"encoding/json"
	"net/http"
	"github.com/bdsoftpro/app/middleware"
)
type data struct {
	Product string `json:"product"`
	Sku string `json:"sku"`
	ProductType string `json:"product_type"`
	Variation string `json:"variation"`
	ProductVariation string `json:"product_variation"`
	Location string `json:"location"`
	MfgDate string `json:"mfg_date"`
	ExpDate string `json:"exp_date"`
	Unit string `json:"unit"`
	StockLeft string `json:"stock_left"`
	RefNo string `json:"ref_no"`
	TransactionId int64 `json:"transaction_id"`
	PurchaseLineId int64 `json:"purchase_line_id"`
	LotNumber int64 `json:"lot_number"`
	Edit string `json:"edit"`
	Name string `json:"name"`
}
// StockExpiry Method Defination
func StockExpiry(w http.ResponseWriter, r *http.Request) {
	_, err := Session.AuthCheck(r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
	}
	datas := []data{
		{
			Product:	"BR-11 Rice",
			Sku:		"0001",
			ProductType:	"single",
			Variation:	"DUMMY",
			ProductVariation:	"DUMMY",
			Location:	"Unlimited Buy",
			MfgDate:	"2019-03-07",
			ExpDate:	"2020-12-07",
			Unit:	"KG",
			StockLeft:	"<span data-is_quantity=\"true\" class=\"display_currency stock_left\" data-currency_symbol=false data-orig-value=\"5.0000\" data-unit=\"KG\" >5.0000</span> KG",
			RefNo:	"<button type=\"button\" data-href=\"http://pos.com/purchases/6\" class=\"btn btn-link btn-modal\" data-container=\".view_modal\"  >2352234234</button>",
			TransactionId:	6,
			PurchaseLineId:	4,
			LotNumber:	0,
			Edit:	"<button type=\"button\" class=\"btn btn-primary btn-xs stock_expiry_edit_btn\" data-transaction_id=\"6\" data-purchase_line_id=\"4\"> <i class=\"fa fa-edit\"></i> Edit</button>",
			Name:	"BR-11 Rice",
		},
	}
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