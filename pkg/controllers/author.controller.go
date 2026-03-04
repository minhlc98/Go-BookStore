package controllers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/minhlc98/bookstore/pkg/models"
	"github.com/minhlc98/bookstore/pkg/repo"
	"github.com/minhlc98/bookstore/pkg/utils"
	"gorm.io/gorm"
)

type AuthorController struct {
	repo *repo.AuthorRepo
}

func NewAuthorController(r *repo.AuthorRepo) *AuthorController {
	return &AuthorController{repo: r}
}

func (c *AuthorController) CreateAuthor(w http.ResponseWriter, r *http.Request) {
	var author models.Author
	utils.ParseBody(r, &author)
	if err := c.repo.Create(&author); err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(author)
}

func (c *AuthorController) GetAuthor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	author, err := c.repo.GetByID(vars["id"])
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "Author Not Found", http.StatusNotFound)
			return
		}
		panic(err)
	}
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(author)
}

func (c *AuthorController) GetAllAuthor(w http.ResponseWriter, r *http.Request) {
	authors, err := c.repo.List()
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(authors)
}

func (c *AuthorController) DeleteAuthorById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	author, errGetAuthor := c.repo.GetByID(vars["id"])
	if errGetAuthor != nil {
		if errors.Is(errGetAuthor, gorm.ErrRecordNotFound) {
			http.Error(w, "Author Not Found", http.StatusNotFound)
			return
		}
		panic(errGetAuthor)
	}
	errDeleteAuthor := c.repo.Delete(author)
	if errDeleteAuthor != nil {
		if errors.Is(errDeleteAuthor, gorm.ErrRecordNotFound) {
			http.Error(w, "Author Not Found", http.StatusNotFound)
			return
		}
		panic(errDeleteAuthor)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(author)
}

func (c *AuthorController) UpdateAuthorById(w http.ResponseWriter, r *http.Request) {
	var author models.Author
	utils.ParseBody(r, &author)
	vars := mux.Vars(r)
	authorDetail, err := c.repo.GetByID(vars["id"])
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "Author Not Found", http.StatusNotFound)
			return
		}
		panic(err)
	}
	authorDetail.Name = author.Name
	authorDetail.Bio = author.Bio

	if errUpdate := c.repo.Update(authorDetail); errUpdate != nil {
		if errors.Is(errUpdate, gorm.ErrRecordNotFound) {
			http.Error(w, "Author Not Found", http.StatusNotFound)
			return
		}
		panic(errUpdate)
	}
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(authorDetail)
}
