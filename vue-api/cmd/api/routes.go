package main

import (
	"net/http"
	"time"
	"vue-api/internal/data"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

// Routes generates our routes and attaches them to handlers, using the chi router
// Note that we return type http.Handler, and not *chi.Mux, since chi.Mux satisfies
// The interface requirements for http.Hander, it makes sense to returns the type
// That is part of the standard library.
func (app *application) routes() http.Handler {

	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)

	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	mux.Post("/users/login", app.Login)
	mux.Post("/users/logout", app.Logout)

	mux.Post("/books", app.AllBooks)
	mux.Get("/books", app.AllBooks)
	mux.Get("/books/{slug}", app.OneBook)

	mux.Post("/validate-token", app.ValidateToken)

	// All of this routes in the block below are prefixed whit /admin, and also require that the user have a valid token provided in the request
	mux.Route("/admin", func(mux chi.Router) {
		mux.Use(app.AuthTokenMiddleware)

		// Admin user routes
		mux.Post("/users", app.AllUsers)
		mux.Post("/users/save", app.EditUser)
		mux.Post("/users/get/{id}", app.GetUser)
		mux.Post("/users/delete", app.DeleteUser)
		mux.Post("/log-user-out/{id}", app.LogUserOutAndSetInactive)

		// Admin book routes
		mux.Post("/authors/all", app.AuthorsAll)
		mux.Post("/books/save", app.EditBook)
		mux.Post("/books/delete", app.DeleteBook)
		mux.Post("/books/{id}", app.BookByID)
	})

	// ---------------------------------------------------------------------------------------------------------------------------------------------
	// Auxiliar functions:
	mux.Get("/users/add", func(w http.ResponseWriter, r *http.Request) {
		var u = data.User{
			Email:     "you@there.com",
			FirstName: "You",
			LastName:  "There",
			Password:  "password",
		}

		app.infoLog.Println("Adding user...")

		id, err := app.models.User.Insert(u)
		if err != nil {
			app.errorLog.Println(err)
			app.errorJSON(w, err, http.StatusForbidden)
			return
		}

		app.infoLog.Println("Got back id of", id)
		newUser, _ := app.models.User.GetOne(id)
		app.writeJSON(w, http.StatusOK, newUser)
	})

	mux.Get("/test-generate-token", func(w http.ResponseWriter, r *http.Request) {
		token, err := app.models.User.Token.GenerateToken(2, 60*time.Minute)
		if err != nil {
			app.errorLog.Println(err)
			return
		}

		token.Email = "gabriel@example.com"
		token.CreatedAt = time.Now()
		token.UpdatedAt = time.Now()

		payload := jsonResponse{
			Error:   false,
			Message: "Success",
			Data:    token,
		}

		app.writeJSON(w, http.StatusOK, payload)
	})

	mux.Get("/test-save-token", func(w http.ResponseWriter, r *http.Request) {
		token, err := app.models.User.Token.GenerateToken(2, 60*time.Minute)
		if err != nil {
			app.errorLog.Println(err)
			return
		}

		user, err := app.models.User.GetOne(2)
		if err != nil {
			app.errorLog.Println(err)
			return
		}

		token.UserID = user.ID
		token.CreatedAt = time.Now()
		token.UpdatedAt = time.Now()

		err = token.Insert(*token, *user)
		if err != nil {
			app.errorLog.Println(err)
			return
		}

		payload := jsonResponse{
			Error:   false,
			Message: "Success",
			Data:    token,
		}

		app.writeJSON(w, http.StatusOK, payload)
	})

	mux.Get("/test-validate-token", func(w http.ResponseWriter, r *http.Request) {
		tokenToValidate := r.URL.Query().Get("token")
		valid, err := app.models.Token.ValidToken(tokenToValidate)
		if err != nil {
			app.errorJSON(w, err)
			return
		}

		var payload jsonResponse
		payload.Error = false
		payload.Data = valid

		app.writeJSON(w, http.StatusOK, payload)

	})

	// Static files
	filesServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", filesServer))

	return mux
}
