package main

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/udayp11/chirpy/internal/auth"
)

func (cfg *apiConfig) handlerDeleteChirps(w http.ResponseWriter, r *http.Request) {

	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't find JWT", err)
		return
	}
	userID, err := auth.ValidateJWT(token, cfg.jwtSecret)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't validate JWT", err)
		return
	}

	chirpIDString := r.PathValue("chirpID")
	chirpID, err := uuid.Parse(chirpIDString)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid chirp ID", err)
		return
	}
	dbChirp, err := cfg.db.GetChirp(r.Context(), chirpID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Couldn't get chirp", err)
		return
	}
	if userID == dbChirp.UserID {
		err := cfg.db.DeleteChirp(r.Context(), chirpID)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Couldn't Delete chirp", err)
			return
		}
		w.WriteHeader(http.StatusNoContent)

		return
	}

	respondWithError(w, http.StatusForbidden, "You cannot delete this chirp", err)

}
