package wflow

import (
	"fmt"
	"sync"
)

const hasErrorsKey string = "hasErrors"

type Workflow struct {
	state *sync.Map
	steps *sync.Map
}

func New() *Workflow {
	return &Workflow{}
}

func (wf *Workflow) add(action string) *Step {
	step := &Step{wf, action, nil, nil, nil}
	wf.steps.Store(step, true)
	return step
}

func (wf *Workflow) hasErrors() bool {
	_, ok := wf.state.Load(hasErrorsKey)
	return ok
}

func (wf *Workflow) Wait() {
	wf.steps.Range(func(key, value any) bool {
		step := key.(*Step)
		step.wait()
		return true
	})
}

// Will wait for all steps to finish and return first error
func (wf *Workflow) FirstError() error {
	var err error
	wf.steps.Range(func(key, value any) bool {
		step := key.(*Step)
		step.wait()
		if step.err != nil {
			err = step.err
			return false
		}
		return true
	})
	return err
}

// Will wait for all steps to finish and return a new combined error of all the errors
func (wf *Workflow) AllErrors() error {
	var err error
	wf.steps.Range(func(key, value any) bool {
		step := key.(*Step)
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

type Step struct {
	wf     *Workflow
	action string
	wg     *sync.WaitGroup
	err    error
	after  []*Step
}

// Returns true if err==nil and stores the error in the step otherwise.
// Also marks the workflow as failed if err != nil, so that other unstarted steps do not start.
func (s *Step) processError(err error) error {
	if err == nil {
		return nil
	}
	s.err = fmt.Errorf("%s: %w", s.action, err)
	s.wf.state.LoadOrStore(hasErrorsKey, true)
	return s.err
}

// Returns whether the step should run, but first waits for dependancies to finish.
func (s *Step) shouldRun() bool {
	for _, step := range s.after {
		step.wait()
	}
	return !s.wf.hasErrors()
}

func (s *Step) runAsync(f func()) {
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

func (s *Step) wait() {
	if s.wg != nil {
		s.wg.Wait()
	}
}

// Marks the steps this step should wait for before running. Returns immediately
func (s *Step) After(steps ...*Step) {
	s.after = steps
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
