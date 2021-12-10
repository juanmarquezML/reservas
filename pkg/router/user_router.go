package router

import (
	"encoding/json"
	"fmt"
	"learngo/api_only_go_task/cmd/utils"
	"learngo/api_only_go_task/pkg/model"
	"learngo/api_only_go_task/pkg/service"
	"net/http"
)

type UserRouter interface {
	HandleUsers(w http.ResponseWriter, req *http.Request)
}

type userRouter struct {
	service service.UserService
}

func NewUserRouter(s service.UserService) *userRouter {
	return &userRouter{s}
}

func (ro *userRouter) HandleUsers(w http.ResponseWriter, req *http.Request) {
	var err error
	resp := make([]byte, 0)
	var u *model.User
	w.Header().Add("Content-Type", "application/json")
	switch req.Method {
	case "GET":
		// -> BUSCAMOS PARA ENCONTRAR EL USUARIO
		if email := req.URL.Query().Get("email"); email != "" {
			user, err := ro.service.GetUser(email)
			if err != nil {
				utils.ShowMessage(w, fmt.Sprint(err))

			} else {
				resp, err = json.Marshal(user)
				_, err = w.Write(resp)
			}
		} else {
			utils.ShowMessage(w, "email_obligatory")
			return
		}

	case "POST":
		// -> BUSCAMOS PARA INSERTAR EL REGISTRO
		err = json.NewDecoder(req.Body).Decode(&u)
		if err != nil {
			utils.ShowMessage(w, "error_body")
			return
		}
		//validar  los datos que vengan sean valido y los correspondientes
		//err = utils.Validation(u)
		err = utils.Validation(u)
		if err != nil {
			utils.ShowMessage(w, fmt.Sprintf(err.Error()))
			return
		}

		//reqBodyBytes := new(bytes.Buffer)
		err = ro.service.CreateUser(u)
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
