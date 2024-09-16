package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/yahn1ukov/personal-blog/internal/dto"
	"github.com/yahn1ukov/personal-blog/internal/repository"
	"github.com/yahn1ukov/personal-blog/internal/service"
	"github.com/yahn1ukov/personal-blog/pkg/respond"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Handler struct {
	service service.Service
}

func New(service service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(1 << 20); err != nil {
		respond.Error(w, http.StatusBadRequest, err)
		return
	}

	var input dto.CreateInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		respond.Error(w, http.StatusInternalServerError, err)
		return
	}

	if err := h.service.Create(r.Context(), &input); err != nil {
		if errors.Is(err, service.ErrTitleRequired) || errors.Is(err, service.ErrContentRequired) {
			respond.Error(w, http.StatusBadRequest, err)
			return
		}

		respond.Error(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	blogs, err := h.service.GetAll(r.Context())
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, err)
		return
	}

	respond.JSON(w, http.StatusOK, blogs)
}

func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, err)
		return
	}

	blog, err := h.service.GetByID(r.Context(), objectID)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			respond.Error(w, http.StatusNotFound, err)
			return
		}

		respond.Error(w, http.StatusInternalServerError, err)
		return
	}

	respond.JSON(w, http.StatusOK, blog)
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, err)
		return
	}

	if err = r.ParseMultipartForm(1 << 20); err != nil {
		respond.Error(w, http.StatusBadRequest, err)
		return
	}

	var input dto.UpdateInput
	if err = json.NewDecoder(r.Body).Decode(&input); err != nil {
		respond.Error(w, http.StatusInternalServerError, err)
		return
	}

	if err = h.service.Update(r.Context(), objectID, &input); err != nil {
		if errors.Is(err, service.ErrNoFieldsUpdate) {
			respond.Error(w, http.StatusBadRequest, err)
			return
		}

		if errors.Is(err, repository.ErrNotFound) {
			respond.Error(w, http.StatusNotFound, err)
			return
		}

		respond.Error(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, err)
		return
	}

	if err = h.service.Delete(r.Context(), objectID); err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			respond.Error(w, http.StatusNotFound, err)
			return
		}

		respond.Error(w, http.StatusNotFound, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
