package routes

import (
	"github.com/gorilla/mux"
	"goblog/app/http/controllers"
	"net/http"
)

// RegisterWebRoutes 注册网页相关路由
func RegisterWebRoutes(r *mux.Router) {
	// 静态页面
	pagesController := new(controllers.PagesController)
	r.HandleFunc("/", pagesController.Home).Methods("GET").Name("home")
	r.HandleFunc("/about", pagesController.About).Methods("GET").Name("about")
	r.NotFoundHandler = http.HandlerFunc(pagesController.NotFound)
}
