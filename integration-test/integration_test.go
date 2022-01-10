package integration_test

import (
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	. "github.com/Eun/go-hit"
)

const (
	// Attempts connection
	host       = "app:8080"
	healthPath = "http://" + host + "/healthz"
	attempts   = 40
	// HTTP REST
	basePath = "http://" + host
)

func TestMain(m *testing.M) {
	err := healthCheck(attempts)
	if err != nil {
		log.Fatalf("Integration tests: host %s is not available: %s", host, err)
	}

	log.Printf("Integration tests: host %s is available", host)

	code := m.Run()
	os.Exit(code)
}

func healthCheck(attempts int) error {
	var err error

	for attempts > 0 {
		err = Do(Get(healthPath), Expect().Status().Equal(http.StatusOK))
		if err == nil {
			return nil
		}

		log.Printf("Integration tests: url %s is not available, attempts left: %d", healthPath, attempts)

		time.Sleep(time.Second)

		attempts--
	}

	return err
}

// HTTP POST: /v1/comments
func TestCommentinsertReview(t *testing.T) {
	body := `{
		"userID":1,
		"contentId":"1",
		"rate":3,
		"comment":"It was so good, I strongly recommend to order from here!",
		"userName":"artist_boy_1995"
	}`
	Test(t,
		Description("comment insertReview Success"),
		Post(basePath+"/v1/comments"),
		Send().Headers("Content-Type").Add("application/json"),
		Send().Body().String(body),
		Expect().Status().Equal(http.StatusOK),
		Expect().Body().JSON().JQ(".userName").Equal("artist_boy_1995"),
	)

}
