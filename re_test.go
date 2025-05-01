package wflow_test

import (
	"testing"
	"time"

	"github.com/acudac-com/wflow-go"
)

func Test_RE1(t *testing.T) {
	getUser := func(name string) (*User, error) {
		time.Sleep(500 * time.Millisecond)
		// return nil, errors.New("not found")
		return &User{name, 30}, nil
	}
	wf := wflow.New()
	userPtr, _ := wflow.RE1(wf, "reading go row", getUser).Go("Go")
	time.Sleep(20 * time.Millisecond)
	user, err := wflow.RE1(wf, "reading do row", getUser).Do("Do")
	log("do", user)
	logErr("do", err)
	println("waiting...")
	if err := wf.AllErrors(); err != nil {
		logErr("wf", err)
	}
	log("go", *userPtr)
}
