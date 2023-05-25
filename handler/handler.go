package handler

import (
	"somename/repository"
)

type Handler struct {
	r *repository.Repository
}

func GetHandler(repo *repository.Repository) *Handler {
	return &Handler{r: repo}
}
