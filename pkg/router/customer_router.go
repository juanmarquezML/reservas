package router

import (
	"encoding/json"
	"fmt"
	"learngo/api_only_go_task/cmd/utils"
	"learngo/api_only_go_task/pkg/model"
	"learngo/api_only_go_task/pkg/service"
	"net/http"
)

type CustomerRouter interface {
	HandlerCustumer(w http.ResponseWriter, req *http.Request)
}

type customerRouter struct {
	service service.CustomerService
}

func NewCustomerRouter(s service.CustomerService) *customerRouter {
	return &customerRouter{s}
}

func (csr *customerRouter) HandlerCustomer(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var cst *model.Customer
	var err error
	resp := make([]byte, 0)
	switch req.Method {
	case "GET":
		if cuit := req.URL.Query().Get("cuit"); cuit != "" {
			if ctsData, err := csr.service.GetCustomer(cuit); err == nil {
				resp, err = json.Marshal(ctsData)
				_, err = w.Write(resp)
			} else {
				utils.ShowMessage(w, fmt.Sprint(err))
			}
		} else {
			utils.ShowMessage(w, "cuit_obligatory")
			return
		}

	case "POST":
		// -> BUSCAMOS PARA INSERTAR EL REGISTRO
		err = json.NewDecoder(req.Body).Decode(&cst)
		if err != nil {
			utils.ShowMessage(w, "error_body")
			return
		}
		//validar  los datos que vengan sean valido y los correspondientes
		err = utils.Validation(cst)
		if err != nil {
			utils.ShowMessage(w, fmt.Sprintf(err.Error()))
			return
		}
		err = csr.service.CreateCustomer(cst)
		if err != nil {
			utils.ShowMessage(w, "error_post")
			return
		}

		utils.ShowMessage(w, "post_succefully")

	default:
		http.Error(w, "Method not enabled", http.StatusInternalServerError)
		return
	}

}
