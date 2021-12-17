package links

import (
	"database/sql"
	"github.com/sivaprasadreddy/devzone-api-golang/models"
)

type linkRepo struct {
	db *sql.DB
}

func NewLinkRepo(db *sql.DB) *linkRepo {
	return &linkRepo{db}
}

func (b *linkRepo) GetLinks() ([]models.Link, error) {
	rows, err := b.db.Query(`SELECT id, title, url FROM links`)
	if err != nil {
		return nil, err
	}
	var links []models.Link

	defer rows.Close()
	for rows.Next() {
		var link = models.Link{}
		err = rows.Scan(&link.Id, &link.Title, &link.Url)
		if err != nil {
			return nil, err
		}
		links = append(links, link)
	}
	return links, nil
}
