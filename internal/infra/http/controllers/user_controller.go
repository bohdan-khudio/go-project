package controllers

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/test_server/internal/domain/user"
	"net/http"
	"strconv"
)

type UserController struct {
	service *user.Service
}

func NewUserController(s *user.Service) *UserController {
	return &UserController{
		service: s,
	}
}

func (c *UserController) FindAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := (*c.service).FindAll()
		if err != nil {
			fmt.Printf("UserController.FindAll(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("UserController.FindAll(): %s", err)
			}
			return
		}

		err = success(w, users)
		if err != nil {
			fmt.Printf("UserController.FindAll(): %s", err)
		}
	}
}

func (c *UserController) FindOne() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		if err != nil {
			fmt.Printf("EventController.FindOne(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.FindOne(): %s", err)
			}
			return
		}
		user, err := (*c.service).FindOne(id)
		if err != nil {
			fmt.Printf("EventController.FindOne(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.FindOne(): %s", err)
			}
			return
		}

		err = success(w, user)
		if err != nil {
			fmt.Printf("UserController.FindOne(): %s", err)
		}
	}
}
