package event

import (
	"github.com/upper/db/v4/adapter/postgresql"
	"log"
)

type Repository interface {
	FindAll() ([]Event, error)
	FindOne(id int64) (*Event, error)
	Create(event Event) (*Event, error)
	Update(event Event) (*Event, error)
	Delete(id int64) (int64, error)
}

type repository struct {
	// Some internal data
}

var settings = postgresql.ConnectionURL{
	Database: `eventdb`,
	Host:     `localhost:5432`,
	User:     `postgres`,
	Password: `school22`,
}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) FindAll() ([]Event, error) {

	sess, err := postgresql.Open(settings)
	if err != nil {
		log.Fatal("Open: ", err)
	}
	defer sess.Close()

	var events []Event
	eventColl := sess.Collection("events")
	err = eventColl.Find().All(&events)
	if err != nil {
		log.Fatal("eventCol.FindAll: ", err)
	}
	return events, nil
}

func (r *repository) FindOne(id int64) (*Event, error) {

	sess, err := postgresql.Open(settings)
	if err != nil {
		log.Fatal("Open: ", err)
	}
	defer sess.Close()

	var event Event
	eventColl := sess.Collection("events")
	err = eventColl.Find("id", id).One(&event)
	if err != nil {
		log.Fatal("eventCol.FindOne: ", err)
	}

	return &event, nil

}

func (r *repository) Create(event Event) (*Event, error) {

	sess, err := postgresql.Open(settings)
	if err != nil {
		log.Fatal("Open: ", err)
	}
	defer sess.Close()

	eventColl := sess.Collection("events")
	res, e := eventColl.Insert(event)
	if e != nil {
		log.Fatal("eventCol.Insert: ", err)
	}

	newEvent, err := r.FindOne(res.ID().(int64))
	if e != nil {
		log.Fatal("eventCol.Insert: ", err)
	}

	return newEvent, nil
}

func (r *repository) Update(event Event) (*Event, error) {

	sess, err := postgresql.Open(settings)
	if err != nil {
		log.Fatal("Open: ", err)
	}
	defer sess.Close()

	eventColl := sess.Collection("events")
	err = eventColl.Find("id", event.Id).Update(event)
	if err != nil {
		log.Fatal("eventCol.Update: ", err)
	}

	updatedEvent, err := r.FindOne(event.Id)
	if err != nil {
		log.Fatal("eventCol.Insert: ", err)
	}

	return updatedEvent, nil
}

func (r repository) Delete(id int64) (int64, error) {
	sess, err := postgresql.Open(settings)
	if err != nil {
		log.Fatal("Open: ", err)
	}
	defer sess.Close()

	eventColl := sess.Collection("events")
	err = eventColl.Find("id", id).Delete()
	if err != nil {
		log.Fatal("eventCol.Delete: ", err)
	}

	return id, nil
}
