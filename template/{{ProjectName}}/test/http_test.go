package test

import (
    {{ if Http }}
	"bytes"
	"encoding/json"
	"{{ProjectName}}/internal"
	myhttp "{{ProjectName}}/api/http"
    {{ end }}
	"fmt"
	"github.com/stretchr/testify/assert"
	"{{ProjectName}}/pkg"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestHealthRoute(t *testing.T) {
	url := fmt.Sprint("http://127.0.0.1:4040/health")
	rsp, err := http.Get(url)
	if err != nil {
		pkg.GetLog().Fatal(err)
	}

	assert.Equal(t, http.StatusOK, rsp.StatusCode)
	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		pkg.GetLog().Fatal(err)
	}
	assert.Equal(t, string(body), "{\"status\":1}")
}

func TestExampleRoute(t *testing.T) {
	{{ if Http }}
	url := fmt.Sprint("http://127.0.0.1:4040/example")
	b, err := json.Marshal(myhttp.Example{Title: "helloTitle", Body: "helloBody"})

	rsp, err := http.Post(url, "application/json", bytes.NewBuffer(b))
	if err != nil {
		pkg.GetLog().Fatal(err)
	}

	assert.Equal(t, http.StatusOK, rsp.StatusCode)
	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		pkg.GetLog().Fatal(err)
	}
	assert.Contains(t, string(body), "helloTitle")
	{{ end }}
}


func TestAdminExampleRoute(t *testing.T) {
	{{ if Http }}
	config := internal.GetConfig()
	user := config.Endpoints.Http.User
	pass := config.Endpoints.Http.Pass
	url := fmt.Sprintf("http://%s:%s@127.0.0.1:4040/admin/example", user, pass)

	rsp, err := http.Post(url, "application/json", bytes.NewBuffer([]byte{}))
	if err != nil {
		pkg.GetLog().Fatal(err)
	}
	assert.Equal(t, http.StatusOK, rsp.StatusCode)
	{{ end }}
}