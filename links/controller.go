package links

import (
	"encoding/json"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/sivaprasadreddy/devzone-api-golang/helpers"
	"github.com/sivaprasadreddy/devzone-api-golang/models"
	"net/http"
	"strconv"
	"time"
)

type LinkController struct {
	service *linkService
}

func NewLinkController(service *linkService) *LinkController {
	return &LinkController{service}
}

func (b *LinkController) GetAll(w http.ResponseWriter, r *http.Request) {
	log.Info("Fetching all links")
	links, err := b.service.GetLinks()
	if err != nil {
		log.Errorf("Error while fetching links")
		helpers.RespondWithError(w, http.StatusInternalServerError, "Unable to fetch links")
		return
	}
	if links == nil {
		links = []models.Link{}
	}
	helpers.RespondWithJSON(w, http.StatusOK, links)
}

func (b *LinkController) GetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	log.Infof("Fetching link by id %d", id)
	link, err := b.service.GetLinkById(id)
	if err != nil {
		log.Errorf("Error while fetching link by id")
		helpers.RespondWithError(w, http.StatusInternalServerError, "Unable to fetch link by id")
		return
	}
	helpers.RespondWithJSON(w, http.StatusOK, link)
}

func (b *LinkController) Create(w http.ResponseWriter, r *http.Request) {
	log.Info("create link")
	contentType := r.Header.Get("Content-Type")
	if contentType != "" && contentType != "application/json" {
		//http.Error(w, "Content-Type header is not application/json", http.StatusUnsupportedMediaType)
		helpers.RespondWithError(w, http.StatusUnsupportedMediaType, "Content-Type header is not application/json")
		return
	}
	var link models.Link
	err := json.NewDecoder(r.Body).Decode(&link)
	if err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, "Unable to parse request body. Error: "+err.Error())
		return
	}
	//TODO; Get Login User Id
	user := models.User{Id: 1}

	link.CreatedBy = user
	link.CreatedDate = time.Now()
	link, err = b.service.CreateLink(link)
	if err != nil {
		log.Errorf("Error while create link %v", err)
		helpers.RespondWithError(w, http.StatusInternalServerError, "Unable to create link")
		return
	}
	helpers.RespondWithJSON(w, http.StatusCreated, link)
}

func (b *LinkController) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	log.Infof("update link id=%d", id)
	var link models.Link
	err := json.NewDecoder(r.Body).Decode(&link)
	if err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, "Unable to parse request body. Error: "+err.Error())
		return
	}
	link.Id = id
	link.UpdatedDate = time.Now()
	link, err = b.service.UpdateLink(link)
	if err != nil {
		log.Errorf("Error while update link")
		helpers.RespondWithError(w, http.StatusInternalServerError, "Unable to update link")
		return
	}
	link, _ = b.service.GetLinkById(id)
	helpers.RespondWithJSON(w, http.StatusOK, link)
}

func (b *LinkController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	log.Infof("delete link with id=%d", id)
	err := b.service.DeleteLink(id)
	if err != nil {
		log.Errorf("Error while deleting link")
		helpers.RespondWithError(w, http.StatusInternalServerError, "Unable to delete link")
		return
	}
	helpers.RespondWithJSON(w, http.StatusOK, nil)
}
