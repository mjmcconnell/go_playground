package products

func (a *App) initializeRoutes() {

	a.BaseApp.Router.HandleFunc("/product", a.createView).Methods("POST")
	a.BaseApp.Router.HandleFunc("/products", a.listView).Methods("GET")
	a.BaseApp.Router.HandleFunc("/product/{id:[0-9]+}", a.readView).Methods("GET")
	a.BaseApp.Router.HandleFunc("/product/{id:[0-9]+}", a.updateView).Methods("PUT")
	a.BaseApp.Router.HandleFunc("/product/{id:[0-9]+}", a.deleteView).Methods("DELETE")
}
