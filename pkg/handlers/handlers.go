package handlers

import (
	"fmt"
	"log"
	"net/http"

	"generative-web/pkg/templates"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Status(dbPool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		conn, err := dbPool.Acquire(r.Context())
		if err != nil {
			http.Error(w, "Failed to acquire database connection", http.StatusInternalServerError)
			return
		}
		defer conn.Release()

		fmt.Fprintln(w, "Server is running...")
		// check database connection
		err = conn.Ping(r.Context())
		if err != nil {
			fmt.Fprintln(w, "Database connection failed:", err)
		} else {
			fmt.Fprintln(w, "Database connection successful")
		}
	}
}

func Ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Pong !")
	log.Println("Pong !")
}

func Welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the generative web AI, go to /docs for more information")
}

func LoadTemplate(w http.ResponseWriter, r *http.Request) {
	// get template name from request
	name := r.URL.Query().Get("name")
	template := templates.LoadTemplate(name)

	fmt.Fprintf(w, "Template loaded successfully: %v\n\n", template.Name)

	fmt.Fprintf(w, "Template: \n%v", template)

}
