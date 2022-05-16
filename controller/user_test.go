package controller

import (
	"bytes"
	"encoding/json"
	"eraport/config"
	"eraport/model"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
)

func initTestEcho() *echo.Echo{
	config.InitDB()
	e:=echo.New()

	return e
}

func TestPersonAddValid(t *testing.T) {	
	e := initTestEcho()

	// compose request
	newPerson, err := json.Marshal(map[string]string{
		"nama":     "dono",
		"email":    "dono@warkop.id",
		"password": "rahasia",
	})
	if err != nil {
		t.Errorf("marshalling new person failed")
	}
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(newPerson))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/user")

	// send request
	if err = CreateUserController(c); err != nil {
		t.Errorf("should not get error, get error: %s", err)
		return
	}

	// compare status
	if rec.Code != 200 {
		t.Errorf("should return 200, get: %d", rec.Code)
	}

	// compare response
	var p model.Users
	if err = json.Unmarshal(rec.Body.Bytes(), &p); err != nil {
		t.Errorf("unmarshalling returned person failed")
	}
	expectedNama := "dono"
	if p.Nama != expectedNama {
		t.Errorf("person name should be %s, get: %s", expectedNama, p.Nama)
	}
	expectedEmail := "dono@warkop.id"
	if p.Email != expectedEmail {
		t.Errorf("person email should be %s, get: %s", expectedEmail, p.Email)
	}
	expectedPassword := "rahasia"
	if p.Password != expectedPassword {
		t.Errorf("person pasword should be %s, get: %s", expectedPassword, p.Password)
	}
}


