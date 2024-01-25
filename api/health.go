package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
	e "myapp.com/errors"
)

func Health(w http.ResponseWriter, r *http.Request) {
	log.Info().Msg("Health Check...")
	if r.Method != http.MethodGet {
		e.ContentTypeNotAllowed(w, r)
	}

	body, err := json.Marshal(map[string]string{"time": time.Now().String()})
	if err != nil {
		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}
