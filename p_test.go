package wflow_test

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/acudac-com/wflow-go"
)

type User struct {
	name string
	age  int32
}

func log(thing string, object any) {
	println(fmt.Sprintf("\033[1;34m%s\033[0m: %v", thing, object))
}

func logErr(thing string, object any) {
	println(fmt.Sprintf("\033[1;31m%s error\033[0m: %v", thing, object))
}

func Test_P(t *testing.T) {
	getUser := func() *User {
		time.Sleep(500 * time.Millisecond)
		return &User{"John Doe", 30}
	}
	wf := wflow.New()
	user := wflow.P(wf, "reading user row", getUser).Do()
	log("syncUser", user)
	asyncUser, _ := wflow.P(wf, "reading user row", getUser).Go()
	println("waiting...")
	if err := wf.FirstError(); err != nil {
		logErr("wf", err)
	}
	log("asyncUser", asyncUser)
}

func Test_PE(t *testing.T) {
	getUser := func() (*User, error) {
		time.Sleep(500 * time.Millisecond)
		// return nil, errors.New("not found")
		return &User{"John Doe", 30}, nil
	}
	wf := wflow.New()
	asyncUser, _ := wflow.PE(wf, "reading user row", getUser).Go()
	user, err := wflow.PE(wf, "reading user row", getUser).Do()
	if err != nil {
		logErr("user", err)
	}
	log("user", user)
	println("waiting...")
	if err := wf.FirstError(); err != nil {
		logErr("wf", err)
	}
	log("asyncUser", asyncUser)
}

func Test_P1(t *testing.T) {
	getUser := func(name string) *User {
		time.Sleep(500 * time.Millisecond)
		return &User{name, 25}
	}
	wf := wflow.New()
	user := wflow.P1(wf, "fetching user by name", getUser).Do("Do")
	log("syncUser", user)
	asyncUser, _ := wflow.P1(wf, "fetching user by name", getUser).Go("Go")
	println("waiting...")
	if err := wf.FirstError(); err != nil {
		logErr("wf", err)
	}
	log("asyncUser", asyncUser)
}

func Test_PE1(t *testing.T) {
	getUser := func(name string) (*User, error) {
		time.Sleep(500 * time.Millisecond)
		if name == "" {
			return nil, errors.New("name cannot be empty")
		}
		return &User{name, 25}, nil
	}
	wf := wflow.New()
	user, err := wflow.PE1(wf, "fetching user by name", getUser).Do("Do")
	if err != nil {
		logErr("user", err)
	}
	log("user", user)
	asyncUser, _ := wflow.PE1(wf, "fetching user by name", getUser).Go("")
	println("waiting...")
	if err := wf.FirstError(); err != nil {
		logErr("wf", err)
	}
	log("asyncUser", asyncUser)
}

func Test_P2(t *testing.T) {
	getUser := func(name string, age int32) *User {
		time.Sleep(500 * time.Millisecond)
		return &User{name, age}
	}
	wf := wflow.New()
	user := wflow.P2(wf, "fetching user by name and age", getUser).Do("Do", 40)
	log("syncUser", user)
	asyncUser, _ := wflow.P2(wf, "fetching user by name and age", getUser).Go("Go", 40)
	println("waiting...")
	if err := wf.FirstError(); err != nil {
		logErr("wf", err)
	}
	log("asyncUser", *asyncUser)
}

func Test_PE2(t *testing.T) {
	getUser := func(name string, age int32) (*User, error) {
		time.Sleep(500 * time.Millisecond)
		if age < 0 {
			return nil, errors.New("age cannot be negative")
		}
		return &User{name, age}, nil
	}
	wf := wflow.New()
	asyncUser, _ := wflow.PE2(wf, "fetching go user by name and age", getUser).Go("Go", -35)
	user, err := wflow.PE2(wf, "fetching do user by name and age", getUser).Do("Do", -5)
	if err != nil {
		logErr("user", err)
	}
	log("user", user)
	println("waiting...")
	if err := wf.AllErrors(); err != nil {
		logErr("wf", err)
	}
	log("asyncUser", asyncUser)
}
