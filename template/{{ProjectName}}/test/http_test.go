package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
	myhttp "{{ProjectName}}/api/http"
)

func TestHealthRoute(t *testing.T) {
	url := fmt.Sprint("http://127.0.0.1:4040/health")
	rsp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, http.StatusOK, rsp.StatusCode)
	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, string(body), "{\"status\":1}")
}

func TestExampleRoute(t *testing.T) {

	url := fmt.Sprint("http://127.0.0.1:4040/example")
	b, err := json.Marshal(myhttp.Example{Title: "helloTitle", Body: "helloBody"})

	rsp, err := http.Post(url, "application/json", bytes.NewBuffer(b))
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, http.StatusOK, rsp.StatusCode)
	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		log.Fatal(err)
	}
	assert.Contains(t, string(body), "helloTitle")

}

func TestAdminExampleRoute(t *testing.T) {
	user := Conf.Endpoints.Http.User
	pass := Conf.Endpoints.Http.Pass
	url := fmt.Sprintf("http://%s:%s@127.0.0.1:4040/admin/example", user, pass)

	rsp, err := http.Post(url, "application/json", bytes.NewBuffer([]byte{}))
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, http.StatusOK, rsp.StatusCode)

}
