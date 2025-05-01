package wflow_test

import (
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

func logErr(thing string, err error) {
	if err != nil {
		println(fmt.Sprintf("\033[1;31m%s error\033[0m: %v", thing, err))
	}
}

func Test_R0(t *testing.T) {
	getUser := func() *User {
		time.Sleep(500 * time.Millisecond)
		return &User{"John Doe", 30}
	}
	wf := wflow.New()
	user := wflow.R0(wf, "reading user row", getUser).Do()
	log("do", user)
	userPtr, _ := wflow.R0(wf, "reading user row", getUser).Go()
	println("waiting...")
	wf.Wait()
	log("go", *userPtr)
}

func Test_R1(t *testing.T) {
	getUser := func(name string) *User {
		time.Sleep(500 * time.Millisecond)
		return &User{name, 30}
	}
	wf := wflow.New()
	user := wflow.R1(wf, "reading user row", getUser).Do("Do")
	log("do", user)
	userPtr, _ := wflow.R1(wf, "reading user row", getUser).Go("Go")
	println("waiting...")
	wf.Wait()
	log("go", *userPtr)
}

func Test_RR2(t *testing.T) {
	getUser := func(name string, age int32) (*User, error) {
		time.Sleep(500 * time.Millisecond)
		// return nil, errors.New("not found")
		return &User{name, age}, nil
	}
	wf := wflow.New()
	user, err := wflow.RR2(wf, "reading user row", getUser).Do("Do", 30)
	log("do", user)
	logErr("do", err)
	userPtr, errPtr, _ := wflow.RR2(wf, "reading user row", getUser).Go("Go", 30)
	println("waiting...")
	wf.Wait()
	log("go", *userPtr)
	logErr("go", *errPtr)
}
