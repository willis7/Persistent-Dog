package data

import (
	"github.com/willis7/Persistent-Dog/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ReleaseRepository interface {
	GetAll() [] models.Release
	Create(*models.Release) error
	Validate(*models.Release) error
	GetById() models.Release
}

type ReleaseRepoDB struct {
	C *mgo.Collection
}

func (r *ReleaseRepoDB) GetReleases() []models.Release {

	return
}

func (r *ReleaseRepoDB) Create(r *models.Release) error {

	return nil
}

func (r *ReleaseRepoDB) GetById(id string) (release models.Release, err error) {
	err = r.C.FindId(bson.ObjectIdHex(id)).One(&release)
	return
}
