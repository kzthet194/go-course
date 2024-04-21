package main

import (
	"testing"

	"github.com/go-chi/chi"
	"github.com/kzthet194/go-course/internal/config"
)

func TestRou(t *testing.T) {
	var app config.AppConfig

	mux := routes(&app)

	switch v := mux.(type) {
	case *chi.Mux:
		// do nothing
	default:
		t.Errorf("type is not *chi.Mux, but is %T", v)
	}
}
