package products

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"

	"github.com/mjmcconnell/go_playground/base"
)

func (a *App) listView(w http.ResponseWriter, r *http.Request) {
	count, _ := strconv.Atoi(r.FormValue("count"))
	start, _ := strconv.Atoi(r.FormValue("start"))

	if count > 10 || count < 1 {
		count = 10
	}
	if start < 0 {
		start = 0
	}

	records, err := fetch(a.DB, start, count)
	if err != nil {
		base.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	base.RespondWithJSON(w, http.StatusOK, records)
}

func (a *App) createView(w http.ResponseWriter, r *http.Request) {
	var p product
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		base.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := p.create(a.DB); err != nil {
		base.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	base.RespondWithJSON(w, http.StatusCreated, p)
}

func (a *App) readView(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		base.RespondWithError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	p := product{ID: id}
	if err := p.read(a.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			base.RespondWithError(w, http.StatusNotFound, "Product not found")
		default:
			base.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	base.RespondWithJSON(w, http.StatusOK, p)
}

func (a *App) updateView(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		base.RespondWithError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	var p product
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		base.RespondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
		return
	}
	defer r.Body.Close()
	p.ID = id

	if err := p.update(a.DB); err != nil {
		base.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	base.RespondWithJSON(w, http.StatusOK, p)
}

func (a *App) deleteView(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		base.RespondWithError(w, http.StatusBadRequest, "Invalid Product ID")
		return
	}

	p := product{ID: id}
	if err := p.delete(a.DB); err != nil {
		base.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	base.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}
