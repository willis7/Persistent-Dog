package common

import (
	"gopkg.in/mgo.v2"
	"time"
	"log"
)

var session *mgo.Session

func GetSession() *mgo.Session {
	if session == nil {
		var err error
		session, err = mgo.DialWithInfo(&mgo.DialInfo{
			Addrs:    []string{AppConfig.MongoDBHost},
			Username: AppConfig.DBUser,
			Password: AppConfig.DBPwd,
			Timeout:  60 * time.Second,
		})
		if err != nil {
			log.Fatalf("[GetSession]: %s\n", err)
		}
	}
	return session
}

func createDbSession() {
	var err error
	session, err = mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    []string{AppConfig.MongoDBHost},
		Username: AppConfig.DBUser,
		Password: AppConfig.DBPwd,
		Timeout:  60 * time.Second,
	})
	if err != nil {
		log.Fatalf("[createDbSession]: %s\n", err)
	}
}

func addIndexes() {
	var err error
	releaseIndex := mgo.Index{
		Key: []string{"name"},
		Unique: false,
		Background: true,
		Sparse: true,
	}
	// Add indexes into MongoDB
	session := GetSession().Copy()
	defer session.Close()

	releaseCol := session.DB(AppConfig.Database).C("releases")
	err = releaseCol.EnsureIndex(releaseIndex)
	if err != nil {
		log.Fatalf("[addIndexes]: %s\n", err)
	}
}
