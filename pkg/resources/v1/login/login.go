package login

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"server/pkg/resources/v1/crypto"

	"github.com/go-chi/chi/v5"
)

type PassResponse struct {
	HashedPass string
}

func LoginRouter() chi.Router {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome from login!"))
	})

	r.Get("/hashtest", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		textParam := r.URL.Query().Get("text")
		hashedText := hex.EncodeToString(crypto.HashText(textParam))
		response := PassResponse{
			HashedPass: string(hashedText),
		}

		data, _ := json.Marshal(response)

		fmt.Fprintf(w, string(data))
	})

	return r
}
