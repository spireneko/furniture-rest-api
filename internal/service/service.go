package service

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/spireneko/furniture-rest-api/internal/model"
	"github.com/spireneko/furniture-rest-api/internal/repository"
)

type Service struct {
	JSONDB repository.JSONDB
}

func NewService(path string) Service {
	return Service{
		JSONDB: repository.NewJSONDB(path),
	}
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

	newFurniture := model.Furniture{
		Name:       req.Name,
		Fabricator: req.Fabricator,
		Height:     req.Height,
		Width:      req.Width,
		Length:     req.Length,
	}

	if err := repository.Create(&newFurniture, &s.JSONDB); err != nil {
		responseError(w, http.StatusInternalServerError, err)
		log.Printf("Error while adding data to db:%s", err)
		return
	}

	response(w, http.StatusCreated, nil)
}

func (s *Service) Get(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		responseError(w, http.StatusBadRequest, err)
		return
	}

	furniture := repository.Get(uint32(id), &s.JSONDB)
	if furniture == nil {
		response(w, http.StatusNoContent, nil)
	}

	response(w, http.StatusOK, furniture)
}

func (s *Service) GetAll(w http.ResponseWriter, r *http.Request) {
	arr := &s.JSONDB.FurnitureJSON.FurnitureArray
	if len(*arr) > 0 {
		response(w, http.StatusOK, *arr)
		return
	}

	response(w, http.StatusNoContent, nil)
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

	newFurniture := model.Furniture{
		Name:       req.Name,
		Fabricator: req.Fabricator,
		Height:     req.Height,
		Width:      req.Width,
		Length:     req.Length,
	}

	if err := repository.Update(uint32(id), &s.JSONDB, &newFurniture); err != nil {
		responseError(w, http.StatusNoContent, err)
	}

	response(w, http.StatusOK, nil)
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

	newFurniture := model.Furniture{
		Name:       req.Name,
		Fabricator: req.Fabricator,
		Height:     req.Height,
		Width:      req.Width,
		Length:     req.Length,
	}

	if err := repository.Patch(uint32(id), &s.JSONDB, &newFurniture); err != nil {
		responseError(w, http.StatusNoContent, err)
	}

	response(w, http.StatusOK, nil)
}

func (s *Service) Delete(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		responseError(w, http.StatusBadRequest, err)
		return
	}

	if err := repository.Delete(uint32(id), &s.JSONDB); err != nil {
		responseError(w, http.StatusNoContent, err)
	}

	response(w, http.StatusOK, nil)
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
