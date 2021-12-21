package links

import (
	"github.com/sivaprasadreddy/devzone-api-golang/models"
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

func (b *linkService) CreateLink(link models.Link) (models.Link, error) {
	return b.repo.CreateLink(link)
}

func (b *linkService) UpdateLink(link models.Link) (models.Link, error) {
	return b.repo.UpdateLink(link)
}

func (b *linkService) DeleteLink(linkId int) error {
	return b.repo.DeleteLink(linkId)
}
