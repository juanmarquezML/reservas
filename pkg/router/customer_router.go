package router

import (
	"encoding/json"
	"fmt"
	"learngo/api_only_go_task/cmd/utils"
	"learngo/api_only_go_task/pkg/model"
	"learngo/api_only_go_task/pkg/service"
	"net/http"
)

type CustumerRouter interface {
	HandlerCustumer(w http.ResponseWriter, req *http.Request)
}

type custumerRouter struct {
	service service.CustumerService
}

func NewCostumerRouter(s service.CustumerService) *custumerRouter {
	return &custumerRouter{s}
}

func (csr *custumerRouter) HandlerCustumer(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":

	case "POST":
		var cst *model.Customer
		var err error
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
		err = csr.service.CreateCustumer(cst)
		if err != nil {
			utils.ShowMessage(w, "error_post")
			return
		}

		utils.ShowMessage(w, "post_succefully")

	default:
		http.Error(w, "Method not enabled", http.StatusInternalServerError)
		return
	}

	fmt.Println("aca andamos")

}
