package wflow

import (
	"fmt"
	"sync"
)

const hasErrorsKey string = "hasErrors"

type Workflow struct {
	steps []*Step
	state *sync.Map
}

func New() *Workflow {
	return &Workflow{}
}

func (wf *Workflow) add(action string) *Step {
	step := &Step{wf, action, &sync.WaitGroup{}, nil}
	wf.steps = append(wf.steps, step)
	return step
}

func (wf *Workflow) hasErrors() bool {
	_, ok := wf.state.Load(hasErrorsKey)
	return ok
}

func (wf *Workflow) Wait() {
	for _, step := range wf.steps {
		step.wg.Wait()
	}
}

func (wf *Workflow) Error() error {
	wf.Wait()
	for _, step := range wf.steps {
		if step.err != nil {
			return step.err
		}
	}
	return nil
}

type Step struct {
	wf     *Workflow
	action string
	wg     *sync.WaitGroup
	err    error
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

func (s *Step) runAsync(f func()) {
	if !s.wf.hasErrors() {
		s.wg.Add(1)
		go func() {
			defer s.wg.Done()
			f()
		}()
	}
}

type pe[T any] struct {
	*Step
	fn func() (*T, error)
}

func PE[T any](wf *Workflow, action string, fn func() (*T, error)) *pe[T] {
	step := wf.add(action)
	return &pe[T]{step, fn}
}

func (s *pe[T]) Go() (*T, *Step) {
	var p T
	s.runAsync(func() {
		result, err := s.fn()
		if s.processError(err) != nil {
			p = *result
		}
	})
	return &p, s.Step
}

func (s *pe[T]) Do() (*T, error) {
	if s.wf.hasErrors() {
		return nil, nil
	}
	result, err := s.fn()
	return result, s.processError(err)
}

type ve[T any] struct {
	*Step
	fn func() (T, error)
}

func VE[T any](wf *Workflow, action string, fn func() (T, error)) *ve[T] {
	step := wf.add(action)
	return &ve[T]{step, fn}
}

func (s *ve[T]) Go() (*T, *Step) {
	var v T
	s.runAsync(func() {
		result, err := s.fn()
		if s.processError(err) != nil {
			v = result
		}
	})
	return &v, s.Step
}

func (s *ve[T]) Do() (T, error) {
	if s.wf.hasErrors() {
		var zero T
		return zero, nil
	}
	result, err := s.fn()
	return result, s.processError(err)
}

type p[T any] struct {
	*Step
	fn func() *T
}

func P[T any](wf *Workflow, action string, fn func() *T) *p[T] {
	step := wf.add(action)
	return &p[T]{step, fn}
}

func (s *p[T]) Go() (*T, *Step) {
	var p T
	s.runAsync(func() {
		result := s.fn()
		p = *result
	})
	return &p, s.Step
}

func (s *p[T]) Do() *T {
	if s.wf.hasErrors() {
		return nil
	}
	return s.fn()
}

type v[T any] struct {
	*Step
	fn func() T
}

func V[T any](wf *Workflow, action string, fn func() T) *v[T] {
	step := wf.add(action)
	return &v[T]{step, fn}
}

func (s *v[T]) Go() (*T, *Step) {
	var v T
	s.runAsync(func() {
		v = s.fn()
	})
	return &v, s.Step
}

func (s *v[T]) Do() T {
	if s.wf.hasErrors() {
		var zero T
		return zero
	}
	return s.fn()
}
