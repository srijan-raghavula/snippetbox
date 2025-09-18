package main

import (
	"github.com/srijan-raghavula/snippetbox/internal/models"
)

type templateData struct {
	Snippet  models.Snippet
	Snippets []models.Snippet
}
