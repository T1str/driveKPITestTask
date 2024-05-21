package handlers

import (
	"encoding/json"
	"net/http"
	"testtask/internal/entities"
	"testtask/internal/validators"
)

type Response struct {
	Message string `json:"message"`
}

func BufferHandler(w http.ResponseWriter, r *http.Request, ch chan entities.Request) {
	if r.Method != http.MethodPost {
		writeJSONResponse(w, Response{Message: "Method not allowed"}, http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		writeJSONResponse(w, Response{Message: "Unable to parse form"}, http.StatusBadRequest)
		return
	}

	request := *entities.NewRequest(
		r.FormValue("period_start"),
		r.FormValue("period_end"),
		r.FormValue("period_key"),
		r.FormValue("indicator_to_mo_id"),
		r.FormValue("indicator_to_mo_fact_id"),
		r.FormValue("value"),
		r.FormValue("fact_time"),
		r.FormValue("is_plan"),
		r.FormValue("auth_user_id"),
		r.FormValue("comment"),
	)

	select {
	case ch <- request:
		if ok := validators.ValidateRequest(request); !ok {
			writeJSONResponse(w, Response{Message: "Invalid request"}, http.StatusBadRequest)
			return
		}
		writeJSONResponse(w, Response{Message: "Thanks for the fact"}, http.StatusOK)
	default:
		writeJSONResponse(w, Response{Message: "Queue is full, please try again later"}, http.StatusServiceUnavailable)
	}
}

func writeJSONResponse(w http.ResponseWriter, response Response, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Failed to create response", http.StatusInternalServerError)
		return
	}
	w.Write(jsonResponse)
}
