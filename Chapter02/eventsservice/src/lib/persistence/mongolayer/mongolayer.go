package mongolayer

import (
	"Chapter02/eventsservice/src/lib/persistence"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	DB = "myevents"
	USERS = "users"
	EVENTS = "events"
)

type MongoDBLayer struct {
	session *mgo.Session
}

func NewMongoDBLayer(connection string) (persistence.DatabaseHandler, error) {
	s, err := mgo.Dial(connection)
	return &MongoDBLayer{
		session: s,
	}, err
}

func (mgoLayer *MongoDBLayer) AddEvent(e persistence.Event) ([]byte, error) {
	s := mgoLayer.getFreshSession()
	defer s.Close()

	if !e.ID.Valid() {
		e.ID = bson.NewObjectId()
	}

	if !e.Location.ID.Valid() {
		e.Location.ID = bson.NewObjectId()
	}

	return []byte(e.ID), s.DB(DB).C(EVENTS).Insert(e)
	
	}

	

