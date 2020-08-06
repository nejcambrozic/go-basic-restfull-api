package rest

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/nejcambrozic/go-basic-restfull-api/pkg/health"
	"github.com/nejcambrozic/go-basic-restfull-api/pkg/models"
	"github.com/nejcambrozic/go-basic-restfull-api/pkg/users"
	"net/http"
	"strconv"
)

// Setups the routes and handler functions
// each handler function gets passed the service interface so it can call the needed method
func Handler(h health.Service, u users.Service) http.Handler {
	router := httprouter.New()

	router.GET("/health", getHealth(h))
	// TODO figure out if there is a nice way to namespace the following endpoints with httprouter
	router.GET("/v1/users", getUsers(u))
	router.POST("/v1/users", createUser(u))
	router.GET("/v1/users/:id", getUser(u))
	router.PUT("/v1/users/:id", updateUser(u))
	router.DELETE("/v1/users/:id", deleteUser(u))
	return router
}

func getHealth(h health.Service) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		// Call the service implementation to get health information
		healthMessage, status := h.GetHealth()
		_addResponseHeaders(w, status)
		json.NewEncoder(w).Encode(healthMessage)
	}
}

func createUser(u users.Service) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		// Decode User object from request body
		decoder := json.NewDecoder(r.Body)
		var newUser models.User
		err := decoder.Decode(&newUser)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		u.CreateUser(newUser)

		_addResponseHeaders(w, 200)
		json.NewEncoder(w).Encode("User created")
	}
}

func getUsers(u users.Service) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		usersResponse := u.GetUsers()

		_addResponseHeaders(w, 200)
		json.NewEncoder(w).Encode(usersResponse)
	}
}

func getUser(u users.Service) func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		// Parse string to int, because we are not using any frameworks
		userId, err := strconv.ParseInt(p.ByName("id"), 10, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		user := u.GetUserById(userId)

		_addResponseHeaders(w, 200)
		json.NewEncoder(w).Encode(user)
	}
}

func deleteUser(u users.Service) func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		// Parse string to int, because we are not using any frameworks
		userId, err := strconv.ParseInt(p.ByName("id"), 10, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		u.DeleteUserById(userId)

		_addResponseHeaders(w, 200)
		json.NewEncoder(w).Encode("User deleted")
	}
}

func updateUser(u users.Service) func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		// Parse string to int, because we are not using any frameworks
		userId, err := strconv.ParseInt(p.ByName("id"), 10, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// Decode User object from request body
		decoder := json.NewDecoder(r.Body)
		var newUser models.User
		err = decoder.Decode(&newUser)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		u.UpdateUserById(userId, newUser)

		_addResponseHeaders(w, 200)
		json.NewEncoder(w).Encode("User updated")
	}
}

func _addResponseHeaders(w http.ResponseWriter, status int) {
	// Always do this: set return type to json
	w.Header().Set("Content-Type", "application/json")
	// Return response status and message returned by the service
	w.WriteHeader(status)
}