package links

import (
	log "github.com/sirupsen/logrus"
	"github.com/sivaprasadreddy/devzone-api-golang/models"
	"time"
)

type linkService struct {
	repo LinkRepository
}

func NewLinkService(repo LinkRepository) *linkService {
	return &linkService{repo: repo}
}

func (b *linkService) GetLinks() ([]models.Link, error) {
	return b.repo.GetLinks()
}

func (b *linkService) GetLinkById(linkId int) (models.Link, error) {
	return b.repo.GetLinkById(linkId)
}

func (b *linkService) CreateLink(createLink CreateLinkModel) (models.Link, error) {
	err := createLink.Validate()
	if err != nil {
		log.Error(err)
		return models.Link{}, err
	}

	//TODO; Get Login User Id
	user := models.User{Id: 1}
	var tags []models.Tag = nil //TODO convert string[] to Tag entities
	link := models.Link{
		Title:       createLink.Title,
		Url:         createLink.Url,
		Tags:        tags,
		CreatedBy:   user,
		CreatedDate: time.Time{},
	}
	return b.repo.CreateLink(link)
}

func (b *linkService) UpdateLink(link models.Link) (models.Link, error) {
	return b.repo.UpdateLink(link)
}

func (b *linkService) DeleteLink(linkId int) error {
	return b.repo.DeleteLink(linkId)
}
