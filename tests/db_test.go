package tests

import (
	"adv/api"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestEnvCheck(t *testing.T) {

	err := godotenv.Load(".env")
	assert.Nil(t, err)

}

func TestDbConnection(t *testing.T) {
	DriverName := os.Getenv("TEST_DBNAME")
	DataSourceName := os.Getenv("TEST_DBSOURCE")

	_, err := sqlx.Connect(DriverName, DataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	assert.Nil(t, err)
}

func TestPingRoute(t *testing.T) {
	router := api.GetRoutes()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/advertisements", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
