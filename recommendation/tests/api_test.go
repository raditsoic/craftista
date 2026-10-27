package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"

	"recommendation/api"
	"recommendation/data"
)

func newRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/api/origami-of-the-day", api.GetOrigamiOfTheDay)
	return r
}

func TestGetOrigamiOfTheDay(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/origami-of-the-day", nil)
	w := httptest.NewRecorder()
	newRouter().ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", w.Code, http.StatusOK)
	}

	var got data.Origami
	if err := json.Unmarshal(w.Body.Bytes(), &got); err != nil {
		t.Fatalf("invalid JSON body: %v", err)
	}

	for _, o := range data.GetDailyOrigami() {
		if o == got {
			return
		}
	}
	t.Errorf("returned origami %q is not in the daily list", got.Name)
}

func TestDailyOrigamiData(t *testing.T) {
	list := data.GetDailyOrigami()
	if len(list) == 0 {
		t.Fatal("daily origami list is empty")
	}
	for _, o := range list {
		if o.Name == "" || o.Description == "" {
			t.Errorf("origami %+v has empty name or description", o)
		}
		if !strings.HasPrefix(o.ImageUrl, "/static/") {
			t.Errorf("origami %q image_url = %q, want /static/ prefix", o.Name, o.ImageUrl)
		}
	}
}
