package links

import (
	"database/sql"
	log "github.com/sirupsen/logrus"
	"github.com/sivaprasadreddy/devzone-api-golang/models"
)

type linkRepo struct {
	db *sql.DB
}

func NewLinkRepo(db *sql.DB) *linkRepo {
	return &linkRepo{db}
}

func (b *linkRepo) GetLinks() ([]models.Link, error) {
	rows, err := b.db.Query(`SELECT id, title, url, created_at FROM links`)
	if err != nil {
		return nil, err
	}
	var links []models.Link

	defer rows.Close()
	for rows.Next() {
		var link = models.Link{}
		err = rows.Scan(&link.Id, &link.Title, &link.Url, &link.CreatedDate)
		if err != nil {
			return nil, err
		}
		links = append(links, link)
	}
	return links, nil
}

func (b *linkRepo) GetLinkById(linkId int) (models.Link, error) {
	log.Infof("Fetching link with id=%d", linkId)
	var link = models.Link{CreatedBy: models.User{}}
	b.db.QueryRow(`select id, title, url, created_by, created_at, updated_at FROM links where id=$1`, linkId).Scan(
		&link.Id, &link.Title, &link.Url, &link.CreatedBy.Id, &link.CreatedDate, &link.UpdatedDate)

	return link, nil
}

func (b *linkRepo) CreateLink(link models.Link) (models.Link, error) {
	var lastInsertID int
	err := b.db.QueryRow("insert into links(title, url, created_by, created_at) values($1, $2, $3,$4) RETURNING id",
		link.Title, link.Url, link.CreatedBy.Id, link.CreatedDate).Scan(&lastInsertID)
	if err != nil {
		log.Errorf("Error while inserting link row: %v", err)
		return models.Link{}, err
	}
	link.Id = lastInsertID
	return link, nil
}

func (b *linkRepo) UpdateLink(link models.Link) (models.Link, error) {
	_, err := b.db.Exec("update links set title = $1, url=$2, updated_at=$3 where id=$4",
		link.Title, link.Url, link.UpdatedDate, link.Id)
	if err != nil {
		return models.Link{}, err
	}
	return link, nil
}

func (b *linkRepo) DeleteLink(linkId int) error {
	deleteStmt := `delete from links where id=$1`
	_, err := b.db.Exec(deleteStmt, linkId)
	return err
}
