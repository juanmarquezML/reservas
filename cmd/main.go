package main

import (
	"learngo/api_only_go_task/pkg/router"
	"learngo/api_only_go_task/pkg/service"
	"learngo/api_only_go_task/pkg/storage"
	"log"
	"net/http"
)

func main() {
	//init storage
	userStorage := storage.NewUserStorage()
	customerStorage := storage.NewCustomerStorage()
	// init service
	userService := service.NewUserService(userStorage)
	customerService := service.NewCustomerService(customerStorage)

	// init routers
	userRouter := router.NewUserRouter(userService)
	customerRouter := router.NewCustomerRouter(customerService)

	mux := http.NewServeMux()

	mux.HandleFunc("/customer", customerRouter.HandlerCustomer)
	mux.HandleFunc("/user", userRouter.HandleUsers)

	log.Println("Server starting")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}
