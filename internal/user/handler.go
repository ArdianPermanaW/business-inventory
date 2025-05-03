package user

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Routes(svc *Service) chi.Router {
	r := chi.NewRouter()
	r.Post("/register", registerHandler(svc))
	r.Post("/login", loginHandler(svc))
	return r
}

func registerHandler(svc *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req RegisterRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid input", http.StatusBadRequest)
			return
		}

		if err := svc.Register(r.Context(), req); err != nil {
			http.Error(w, "could not register", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("registered successfully"))
	}
}

func loginHandler(svc *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req LoginRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid input", http.StatusBadRequest)
			return
		}

		user, err := svc.AuthenticateUser(r.Context(), req)
		if err != nil {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		//return a fake token for now!
		resp := map[string]string{
			"message": "login successful",
			"user":    user.Username,
			"token":   "mock-token-here", //TODO: real JWT soon
		}
		json.NewEncoder(w).Encode(resp)
	}
}
