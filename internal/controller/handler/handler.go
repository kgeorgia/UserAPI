package handler

import (
	"log"
	"net/http"
	"os"
	"refactoring/internal/model"
	"refactoring/internal/repository"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type Handler struct {
	l    *log.Logger
	repo repository.Storage
}

func New(filename string) (*Handler, error) {
	l := log.New(os.Stderr, "Error\t", log.Ldate|log.Ltime)
	r, err := repository.New(filename)
	if err != nil {
		return nil, err
	}

	return &Handler{
		l:    l,
		repo: r,
	}, nil
}

func (h *Handler) SearchUsers(w http.ResponseWriter, r *http.Request) {
	u := h.repo.SearchUsers()
	render.JSON(w, r, u)
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	request := model.CreateUserRequest{}

	if err := render.Bind(r, &request); err != nil {
		h.l.Printf("%s", err.Error())
		_ = render.Render(w, r, model.ErrInvalidRequest(err))
		return
	}

	id, err := h.repo.CreateUser(request)
	if err != nil {
		h.l.Printf("Create User: %s", err.Error())
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, map[string]interface{}{
		"user_id": id,
	})
}

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	user, err := h.repo.GetByID(id)
	if err != nil {
		h.l.Printf("repo: %s", err.Error())
		_ = render.Render(w, r, model.ErrInvalidRequest(err))
		return
	}

	render.JSON(w, r, user)
}

func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	request := model.UpdateUserRequest{}

	if err := render.Bind(r, &request); err != nil {
		h.l.Printf("%s", err.Error())
		_ = render.Render(w, r, model.ErrInvalidRequest(err))
		return
	}

	id := chi.URLParam(r, "id")

	err := h.repo.UpdateUser(id, request)
	if err != nil {
		h.l.Printf("update user: %s", err.Error())
		_ = render.Render(w, r, model.ErrInvalidRequest(err))
		return
	}

	render.Status(r, http.StatusNoContent)
}

func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := h.repo.DeleteUser(id)
	if err != nil {
		h.l.Printf("delete user: %s", err.Error())
		_ = render.Render(w, r, model.ErrInvalidRequest(err))
		return
	}

	render.Status(r, http.StatusNoContent)
}
