package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAPI_authentication(t *testing.T) {
	api := New()
	log := LogInfo{
		Usr: "root",
		Psw: "root",
	}

	payload, _ := json.Marshal(log)
	req := httptest.NewRequest(http.MethodPost, "/auth", bytes.NewBuffer(payload))
	rr := httptest.NewRecorder()
	api.router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("incorrect code: get %d, want %d", rr.Code, http.StatusOK)
	}
}
