package subsonic

import (
	"encoding/json"
	"net/http"
)

type SubResponse struct {
	Status        string `json:"status"`
	Version       string `json:"version"`
	Type          string `json:"type"`
	ServerVersion string `json:"serverVersion"`
	OpenSubsonic  bool   `json:"openSubsonic"`
	Error *APIError `json:"error,omitempty"`
}

type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Envelope struct {
	Response SubResponse `json:"subsonic-response"`
}

const (
	APIVersion    = "1.16.1"
	ServerVersion = "0.1.0"
	ServerType    = "TempoServer"
)


func WriteResult(w http.ResponseWriter, status string ) {
	w.WriteHeader(http.StatusOK)
	
	response:=Envelope{
		Response: SubResponse{
			Status: status,
			Version:       APIVersion,
			Type:          ServerType,
			ServerVersion: ServerVersion,
			OpenSubsonic:  true,
		},
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
	
}

func WriteError(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	response := Envelope{
		Response: SubResponse{
			Status:        "failed",
			Version:       APIVersion,
			Type:          ServerType,
			ServerVersion: ServerVersion,
			OpenSubsonic:  true,
			Error: &APIError{
				Code:    code,
				Message: message,
			},
		},
	}

	json.NewEncoder(w).Encode(response)
}