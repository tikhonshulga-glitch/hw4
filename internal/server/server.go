package server

import (
	"fmt"
	"hwproject/internal"
	"net/http"
)

type RentAPI struct {
	srv *http.Server
}

func NewServer(cfg internal.Config) *RentAPI {
	httpSrv := http.Server{
		Addr: fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
	}

	api := RentAPI{
		srv: &httpSrv,
	}

	api.configRouter()

	return &api

}

func (api *RentAPI) Run() error {
	return api.srv.ListenAndServe()
}

func (api *RentAPI) ShutDown() error {
	return nil
}

func (api *RentAPI) configRouter() {
	router := http.NewServeMux()

	router.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		api.GetTasks(w, r)

	})

	router.HandleFunc("/tasks/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		api.GetTask(w, r)
	})

	router.HandleFunc("/tasks/create", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		api.CreateTask(w, r)
	})

	router.HandleFunc("/tasks/update/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		api.UpdateTask(w, r)

	})

	router.HandleFunc("/tasks/delete/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		api.DeleteTask(w, r)
	})

	router.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		api.GetUsers(w, r)
	})

	router.HandleFunc("/users/,", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		api.GetUsers(w, r)
	})

	router.HandleFunc("/users/create", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		api.CreateUser(w, r)
	})

	router.HandleFunc("/users/update/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		api.UpdateUser(w, r)
	})

	router.HandleFunc("/users/delete/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		api.DeleteUser(w, r)
	})

	router.HandleFunc("/users/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
	})

	router.HandleFunc("/users/profile", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
	})

	router.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		w.Write([]byte("service is working correctly"))
	})

	api.srv.Handler = router

}
