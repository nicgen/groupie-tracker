package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestSetupMux(t *testing.T) {
	mux := SetupMux()

	tests := []struct {
		route          string
		expectedStatus int
	}{
		{route: "/", expectedStatus: http.StatusOK},
		{route: "/about", expectedStatus: http.StatusOK},
		{route: "/force500", expectedStatus: http.StatusInternalServerError},
	}

	for _, test := range tests {
		req, err := http.NewRequest(http.MethodGet, test.route, nil)
		if err != nil {
			t.Fatalf("Could not create request: %v", err)
		}

		rr := httptest.NewRecorder()
		mux.ServeHTTP(rr, req)

		if status := rr.Code; status != test.expectedStatus {
			t.Errorf("Handler for %s returned wrong status code: got %v want %v", test.route, status, test.expectedStatus)
		}
	}
}

func TestSetupServer(t *testing.T) {
	mux := SetupMux()
	server := SetupServer(mux)

	if server.Addr != "localhost:8080" {
		t.Errorf("Expected server address 'localhost:8080', got %v", server.Addr)
	}

	if server.ReadHeaderTimeout != 10*time.Second {
		t.Errorf("Expected ReadHeaderTimeout '10s', got %v", server.ReadHeaderTimeout)
	}

	if server.WriteTimeout != 10*time.Second {
		t.Errorf("Expected WriteTimeout '10s', got %v", server.WriteTimeout)
	}

	if server.IdleTimeout != 120*time.Second {
		t.Errorf("Expected IdleTimeout '120s', got %v", server.IdleTimeout)
	}

	if server.MaxHeaderBytes != 1<<20 {
		t.Errorf("Expected MaxHeaderBytes '1048576', got %v", server.MaxHeaderBytes)
	}
}
