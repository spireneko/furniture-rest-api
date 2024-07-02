package service

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type Service struct {
}

type CreateRequest struct {
	Name       string `json:"name"`
	Fabricator string `json:"fabricator"`
	Height     uint32 `json:"height"`
	Width      uint32 `json:"width"`
	Length     uint32 `json:"length"`
}

func (s *Service) Create(w http.ResponseWriter, r *http.Request) {
	req := new(CreateRequest)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		responseError(w, http.StatusBadRequest, err)
		return
	}
	r.Body.Close()

	response(w, http.StatusCreated, nil)
}

type GetResponse struct {
	ID         uint32 `json:"id"`
	Name       string `json:"name"`
	Fabricator string `json:"fabricator"`
	Height     uint32 `json:"height"`
	Width      uint32 `json:"width"`
	Length     uint32 `json:"length"`
}

func (s *Service) Get(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		responseError(w, http.StatusBadRequest, err)
		return
	}

	_ = id

	response(w, http.StatusOK, nil)
}

func (s *Service) GetAll(w http.ResponseWriter, r *http.Request) {

	response(w, http.StatusOK, nil)
}

func (s *Service) Update(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		responseError(w, http.StatusBadRequest, err)
		return
	}

	req := new(CreateRequest)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		responseError(w, http.StatusBadRequest, err)
		return
	}
	r.Body.Close()

	_ = id

	response(w, http.StatusNoContent, nil)
}

func (s *Service) Patch(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		responseError(w, http.StatusBadRequest, err)
		return
	}

	req := new(CreateRequest)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		responseError(w, http.StatusBadRequest, err)
		return
	}
	r.Body.Close()

	_ = id

	response(w, http.StatusNoContent, nil)
}

func (s *Service) Delete(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		responseError(w, http.StatusBadRequest, err)
		return
	}

	_ = id

	response(w, http.StatusNoContent, nil)
}

func response(w http.ResponseWriter, code int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if data != nil {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			log.Println(err)
		}
	}
}

func responseError(w http.ResponseWriter, code int, err error) {
	response(w, code, map[string]string{"error": err.Error()})
}
