package event

import (
	"fmt"
	"github.com/upper/db/v4/adapter/postgresql"
	"log"
)

type Repository interface {
	FindAll() ([]Event, error)
	FindOne(id int64) (*Event, error)
	Create(event Event) (string, error)
	Update(event Event) (string, error)
	Delete(id int64) (string, error)
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

func (r *repository) Create(event Event) (string, error) {

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

	return fmt.Sprintf("Event added, id: %d, Title: %s", res.ID(), event.Title),
		nil
}

func (r *repository) Update(event Event) (string, error) {

	sess, err := postgresql.Open(settings)
	if err != nil {
		log.Fatal("Open: ", err)
	}
	defer sess.Close()

	var updateEvent Event
	eventColl := sess.Collection("events")

	res := eventColl.Find("id", event.Id)
	err = res.One(&updateEvent)
	if err != nil {
		log.Fatal("eventCol.Update: ", err)
	}

	updateEvent.Title = event.Title
	updateEvent.ShortDesc = event.ShortDesc
	updateEvent.Desc = event.Desc
	updateEvent.Long = event.Long
	updateEvent.Lat = event.Lat

	err = res.Update(updateEvent)
	if err != nil {
		log.Fatal("eventCol.Update: ", err)
	}

	return fmt.Sprintf("Event updated, id: %d, New title: %s", updateEvent.Id, updateEvent.Title),
		nil
}

func (r repository) Delete(id int64) (string, error) {
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

	return fmt.Sprintf("Event with id: %d deleted", id),
		nil
}
