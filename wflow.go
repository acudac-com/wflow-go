package wflow

import (
	"fmt"
	"sync"
)

const firstErrorKey string = "hasErrors"

type Workflow struct {
	state *sync.Map
	steps *sync.Map
}

func New() *Workflow {
	return &Workflow{&sync.Map{}, &sync.Map{}}
}

func (wf *Workflow) add(action string, waitFor ...*step) *step {
	step := &step{wf, action, nil, nil, waitFor}
	wf.steps.Store(step, true)
	return step
}

func (wf *Workflow) hasErrors() bool {
	_, ok := wf.state.Load(firstErrorKey)
	return ok
}

func (wf *Workflow) Wait() {
	wf.steps.Range(func(key, value any) bool {
		step := key.(*step)
		step.wait()
		return true
	})
}

// Will return an error if any step already failed and will not wait for remaining steps to finish.
// Will only wait for all steps if no error already found.
func (wf *Workflow) FirstError() error {
	if err, ok := wf.state.Load(firstErrorKey); ok {
		return err.(error)
	}
	wf.Wait()
	if err, ok := wf.state.Load(firstErrorKey); ok {
		return err.(error)
	}
	return nil
}

// Will wait for all steps to finish and return a new combined error of all the errors
func (wf *Workflow) AllErrors() error {
	var err error
	wf.steps.Range(func(key, value any) bool {
		step := key.(*step)
		step.wait()
		if step.err != nil {
			if err == nil {
				err = step.err
			} else {
				err = fmt.Errorf("%w; %s", err, step.err)
			}
		}
		return true
	})
	return err
}

type step struct {
	wf      *Workflow
	action  string
	wg      *sync.WaitGroup
	err     error
	waitFor []*step
}

// If the error is not nil, it is saved on the step and the workflow is marked as failed.
// Also prepends the action in the error description.
func (s *step) processError(err error) error {
	if err == nil {
		return nil
	}
	s.err = fmt.Errorf("%s: %w", s.action, err)
	s.wf.state.LoadOrStore(firstErrorKey, s.err)
	return s.err
}

// Returns whether the step should run, but first waits for dependancies to finish.
func (s *step) shouldRun() bool {
	for _, step := range s.waitFor {
		step.wait()
	}
	return !s.wf.hasErrors()
}

func (s *step) runAsync(f func()) {
	if s.shouldRun() {
		if s.wg == nil {
			s.wg = &sync.WaitGroup{}
		}
		s.wg.Add(1)
		go func() {
			defer s.wg.Done()
			f()
		}()
	}
}

func (s *step) wait() {
	if s.wg != nil {
		s.wg.Wait()
	}
}

func zero[T any]() T {
	var z T
	return z
}

func zero2[T1 any, T2 any]() (T1, T2) {
	var z1 T1
	var z2 T2
	return z1, z2
}

func zero3[T1 any, T2 any, T3 any]() (T1, T2, T3) {
	var z1 T1
	var z2 T2
	var z3 T3
	return z1, z2, z3
}
