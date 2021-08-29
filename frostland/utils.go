package frostland

import (
	"errors"
	"log"

	uuid "github.com/satori/go.uuid"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var mongoAddress, mongoDb, mongoCol string

type Record struct {
	UUID    string `bson:"uuid"`
	UID     string `bson:"uid"`
	Premium bool   `bson:"premium"`
}

func XCreateDataBaseDial() (*mgo.Session, *mgo.Database, *mgo.Collection, error) {
	mongo, err := mgo.Dial(mongoAddress)

	if err != nil {
		return nil, nil, nil, err
	}

	db := mongo.DB(mongoDb)
	c := db.C(mongoCol)

	return mongo, db, c, nil
}

func MCreateUser(uid string, premium bool) (int, string, string) {
	mongo, _, c, err := XCreateDataBaseDial()

	defer mongo.Close()

	if err != nil {
		log.Panic(err)
		return 5001, err.Error(), ""
	}

	if ICheckUIDIfExists(uid) {
		return 5002, "Document already exists", ""
	}

	u := uuid.NewV4().String()
	for ICheckUUIDIfExists(u) {
		u = uuid.NewV4().String()
	}

	data := Record{
		UUID:    u,
		UID:     uid,
		Premium: premium,
	}

	err = c.Insert(&data)

	if err != nil {
		return 5003, err.Error(), ""
	}

	return 0, "OK", u
}

func MQueryUser(uid string) (uuid string, err error) {
	mongo, _, c, err := XCreateDataBaseDial()

	defer mongo.Close()

	if err != nil {
		log.Panic(err)
		return "", err
	}

	if !ICheckUIDIfExists(uid) {
		return "", errors.New("No user found!")
	}

	result := Record{}
	err = c.Find(bson.M{"uid": uid}).One(&result)

	if err != nil {
		return "", err
	}
	return result.UUID, nil
}

func MQueryUUID(uuid string) {

}

func MImportUser(uid string, premium bool, uuid string) (int, string) {
	mongo, _, c, err := XCreateDataBaseDial()

	defer mongo.Close()

	if err != nil {
		log.Panic(err)
		return 5001, err.Error()
	}

	if ICheckUIDIfExists(uid) {
		return 5002, "Document already exists"
	}

	data := Record{
		UUID:    uuid,
		UID:     uid,
		Premium: premium,
	}

	err = c.Insert(&data)

	if err != nil {
		log.Panic(err)
		return 5003, err.Error()
	}

	return 0, "OK"
}

func ICheckUIDIfExists(uid string) bool {
	mongo, _, c, err := XCreateDataBaseDial()

	defer mongo.Close()

	if err != nil {
		log.Panic(err)
		return false
	}

	count, err := c.Find(bson.M{"uid": uid}).Count()

	if err != nil {
		log.Panic(err)
		return false
	}

	return count > 0
}

func ICheckUUIDIfExists(uid string) bool {
	mongo, _, c, err := XCreateDataBaseDial()

	defer mongo.Close()

	if err != nil {
		log.Panic(err)
		return false
	}

	count, err := c.Find(bson.M{"uuid": uid}).Count()

	if err != nil {
		log.Panic(err)
		return false
	}

	return count > 0
}

func IUpdateConfig(jmongoAddress, jmongoDb, jmongoCol string) {
	mongoAddress = jmongoAddress
	mongoDb = jmongoDb
	mongoCol = jmongoCol
}
