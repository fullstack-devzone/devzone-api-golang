package links

import (
	log "github.com/sirupsen/logrus"
	"github.com/sivaprasadreddy/devzone-api-golang/helpers"
	"github.com/sivaprasadreddy/devzone-api-golang/models"
	"net/http"
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
