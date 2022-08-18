package routes

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"https://github.com/zarif007/GO-CRUD/config"
	"https://github.com/zarif007/GO-CRUD/internal/handlers/health"
	"https://github.com/zarif007/GO-CRUD/internal/handlers/product"
)

type Router struct {
	config *Config
	router *chi.Mux
}

func NewRouter() *Router {
	return &Router{
		config: NewConfig().SetTimeout(serviceConfig.GetConfig().Timeout),
		router: chi.NewRouter(),
	}
}

func (r *Router) SetRouters(repository adapter.Interface) *chi.Mux {
	r.setConfigRouters()
	r.RouterHealth(repository)
	r.RouterProduct(repository)

	return r.router
}

func (r *Router) setConfigRouters() {
	r.EnableCORS()
	e.EnableLogger()
	r.EnableRequestID()
	r.EnableTimeout()
	r.EnableRecover()
	r.EnableRealIP()
}

func (r *Router) RouterHealth(repository adapter.Interface) {

}

func RouterProduct() {

}

func EnableTimeout() {

}

func EnableLogger() {

}

func EnableCORS() {

}

func EnableRecover() {

}

func EnableRequestID() {

}

func EnableRealIP() {

}
