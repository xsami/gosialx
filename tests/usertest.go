package tests

import (
	"bytes"
	"fmt"
	"math/rand"
	"strings"

	"github.com/revel/revel/testing"
)

type UserTest struct {
	testing.TestSuite
}

func getRandomString() string {
	var output []string

	for i := 0; i < 10; i++ {
		output = append(output, string(rand.Intn(120)))
	}

	return strings.Join(output, "")
}

func (t *UserTest) Before() {
	println("Set up")
}

func (t *UserTest) TestListAllUsers() {
	t.Get("/users")
	t.AssertOk()
	t.AssertContentType("application/json; charset=utf-8")
}

func (t *UserTest) TestListAUser() {
	t.Get("/users/fdsfdsfderwrewew232")
	t.AssertOk()
	t.AssertContentType("application/json; charset=utf-8")
}

func (t *UserTest) TestCreateAUser() {

	t.Post("/users/create",
		"application/json",
		bytes.NewReader([]byte(fmt.Sprintf("{\"username\": %v}", getRandomString()))))
	t.AssertOk()
	t.AssertContentType("application/json; charset=utf-8")
}

func (t *UserTest) TestUpdateAUser() {

	date := "1995-02-13"
	userid := "dkasjdsdad9a89d87d6asdas"

	t.Post("/users/create",
		"application/json",
		bytes.NewReader([]byte(fmt.Sprintf("{\"birth\": %v, \"_id\": %v}", date, userid))))
	t.AssertOk()
	t.AssertContentType("application/json; charset=utf-8")
}

func (t *UserTest) TestDeleteAUser() {
	t.Delete("/users/delete/3432432dfdsfdsfds33")
	t.AssertOk()
	t.AssertContentType("application/json; charset=utf-8")
}

func (t *UserTest) After() {
	println("Tear down")
}
