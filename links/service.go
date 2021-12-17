package links

import (
	"github.com/sivaprasadreddy/devzone-api-golang/models"
)

type linkService struct {
	repo *linkRepo
}

func NewLinkService(repo *linkRepo) *linkService {
	return &linkService{repo: repo}
}

func (b *linkService) GetLinks() ([]models.Link, error) {
	return b.repo.GetLinks()
}
