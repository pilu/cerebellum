package main

import (
	"database/sql"
	"encoding/json"
	"github.com/pilu/cerebellum/models"
	"github.com/pilu/cerebellum/models/releasegroup"
	"github.com/pilu/traffic"
	"net/http"
)

func ReleaseGroupHandler(w traffic.ResponseWriter, r *traffic.Request) {
	artistId := r.URL.Query().Get("artist_id")
	id := r.URL.Query().Get("id")

	var ReleaseGroup *models.ReleaseGroup
	var err error

	if artistId != "" {
		ReleaseGroup, err = releasegroup.ByArtistIdAndId(artistId, id)
	} else {
		ReleaseGroup, err = releasegroup.ById(id)
	}

	if err == nil {
		json.NewEncoder(w).Encode(ReleaseGroup)
	} else if err == sql.ErrNoRows {
		ReleaseGroupNotFoundHandler(w, r)
	} else if _, ok := err.(models.InvalidUUID); ok {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		panic(err)
	}
}
