package internal_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gi4nks/cesar/internal"
	"github.com/gorilla/mux"
)

// setupRequest function is common utility used to setup the http integration tests infrastructure.
func setupRequest(method, url string) (*httptest.ResponseRecorder, *http.Request) {
	handler := internal.NewHandler()

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		panic(err)
	}

	rr := httptest.NewRecorder()

	r := mux.NewRouter()
	r.HandleFunc("/convert/{start}/{end}", handler.ConvertHandler).Methods("GET")
	r.HandleFunc("/convert/{start}", handler.ConvertHandler).Methods("GET")

	r.ServeHTTP(rr, req)

	return rr, req
}

// TestHandler_ConvertHandlerIntegration_1_5 tests the converter Rest Api from 1 to 5.
// It should simply work without any kind of failure.
func TestHandler_ConvertHandlerIntegration_1_5(t *testing.T) {
	rr, _ := setupRequest("GET", "/convert/1/5")

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: cnv %v expected %v",
			status, http.StatusOK)
	}

	// Check the response body
	expected := `{"1":"I","2":"II","3":"III","4":"IV","5":"V"}`
	current := strings.TrimSpace(rr.Body.String())

	if current != expected {
		t.Errorf("handler returned unexpected body: cnv %v expected %v",
			current, expected)
	}
}

// TestHandler_ConvertHandlerIntegration_0_5 tests the converter Rest Api from 0 to 5.
// It should raise a http.StatusBadRequest error as 4000 is not allowed number.
func TestHandler_ConvertHandlerIntegration_0_5(t *testing.T) {
	rr, _ := setupRequest("GET", "/convert/0/5")

	// Check the status code
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: cnv %v expected %v",
			status, http.StatusBadRequest)
	}
}

// TestHandler_ConvertHandlerIntegration_1_4000 tests the converter Rest Api from 1 to 4000.
// It should raise a http.StatusBadRequest error as 0 is not allowed number.
func TestHandler_ConvertHandlerIntegration_1_4000(t *testing.T) {
	rr, _ := setupRequest("GET", "/convert/1/4000")

	// Check the status code
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: cnv %v expected %v",
			status, http.StatusBadRequest)
	}
}

// TestHandler_ConvertHandlerIntegration_10_1 tests the converter Rest Api from 10 to 1.
// It should raise a http.StatusBadRequest error as the start number is higher than end.
func TestHandler_ConvertHandlerIntegration_10_1(t *testing.T) {
	rr, _ := setupRequest("GET", "/convert/10/1")

	// Check the status code
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: cnv %v expected %v",
			status, http.StatusBadRequest)
	}
}

// TestHandler_ConvertHandlerIntegration_1 tests the converter Rest Api for 1.
// It should simply work without any kind of failure.
func TestHandler_ConvertHandlerIntegration_1(t *testing.T) {
	rr, _ := setupRequest("GET", "/convert/1")

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: cnv %v expected %v",
			status, http.StatusOK)
	}

	// Check the response body
	expected := `{"1":"I"}`
	current := strings.TrimSpace(rr.Body.String())

	if current != expected {
		t.Errorf("handler returned unexpected body: cnv %v expected %v",
			current, expected)
	}
}

// TestHandler_ConvertHandlerIntegration_0 tests the converter Rest Api for 0.
// It should raise a http.StatusBadRequest error as 0 is not allowed number.
func TestHandler_ConvertHandlerIntegration_0(t *testing.T) {
	rr, _ := setupRequest("GET", "/convert/0")

	// Check the status code
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: cnv %v expected %v",
			status, http.StatusBadRequest)
	}
}

// TestHandler_ConvertHandlerIntegration_NaN tests the converter Rest Api for NaN.
// It should raise a http.StatusBadRequest error as NaN is not allowed input.
func TestHandler_ConvertHandlerIntegration_NaN(t *testing.T) {
	rr, _ := setupRequest("GET", "/convert/NaN")

	// Check the status code
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: cnv %v expected %v",
			status, http.StatusBadRequest)
	}
}

// TestHandler_ConvertHandlerIntegration_NaN_NaN tests the converter Rest Api for NaN and NaN.
// It should raise a http.StatusBadRequest error as either start or end are NaN and are not allowed input.
func TestHandler_ConvertHandlerIntegration_NaN_NaN(t *testing.T) {
	rr, _ := setupRequest("GET", "/convert/NaN/NaN")

	// Check the status code
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: cnv %v expected %v",
			status, http.StatusBadRequest)
	}
}

// TestHandler_ConvertHandlerIntegration_1_NaN tests the converter Rest Api for 1 and NaN.
// It should raise a http.StatusBadRequest error as end is NaN and is not allowed input.
func TestHandler_ConvertHandlerIntegration_1_NaN(t *testing.T) {
	rr, _ := setupRequest("GET", "/convert/1/NaN")

	// Check the status code
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: cnv %v expected %v",
			status, http.StatusBadRequest)
	}
}
