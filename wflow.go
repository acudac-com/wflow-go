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
		s.processError(err)
		if result != nil {
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
		s.processError(err)
		v = result
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

// pve: Pointer, Value, Error - function returns (*T, V, error)
type pve[T any, V any] struct {
	*Step
	fn func() (*T, V, error)
}

func PVE[T any, V any](wf *Workflow, action string, fn func() (*T, V, error)) *pve[T, V] {
	step := wf.add(action)
	return &pve[T, V]{step, fn}
}

func (s *pve[T, V]) Go() (*T, *V, *Step) {
	var p T
	var v V
	s.runAsync(func() {
		resultPtr, resultVal, err := s.fn()
		s.processError(err)
		if resultPtr != nil {
			p = *resultPtr
		}
		v = resultVal
	})
	return &p, &v, s.Step
}

func (s *pve[T, V]) Do() (*T, V, error) {
	if s.wf.hasErrors() {
		var zeroV V
		return nil, zeroV, nil
	}
	resultPtr, resultVal, err := s.fn()
	return resultPtr, resultVal, s.processError(err)
}

// pv: Pointer, Value - function returns (*T, V)
type pv[T any, V any] struct {
	*Step
	fn func() (*T, V)
}

func PV[T any, V any](wf *Workflow, action string, fn func() (*T, V)) *pv[T, V] {
	step := wf.add(action)
	return &pv[T, V]{step, fn}
}

func (s *pv[T, V]) Go() (*T, *V, *Step) {
	var p T
	var v V
	s.runAsync(func() {
		resultPtr, resultVal := s.fn()
		if resultPtr != nil {
			p = *resultPtr
		}
		v = resultVal
	})
	return &p, &v, s.Step
}

func (s *pv[T, V]) Do() (*T, V) {
	if s.wf.hasErrors() {
		var zeroV V
		return nil, zeroV
	}
	resultPtr, resultVal := s.fn()
	return resultPtr, resultVal
}

// PE1: Pointer, Error - function takes 1 argument and returns (*T, error)
type pe1[A any, T any] struct {
	*Step
	fn  func(A) (*T, error)
	arg A
}

func PE1[A any, T any](wf *Workflow, action string, fn func(A) (*T, error), arg A) *pe1[A, T] {
	step := wf.add(action)
	return &pe1[A, T]{step, fn, arg}
}

func (s *pe1[A, T]) Go() (*T, *Step) {
	var p T
	s.runAsync(func() {
		result, err := s.fn(s.arg)
		s.processError(err)
		if result != nil {
			p = *result
		}
	})
	return &p, s.Step
}

func (s *pe1[A, T]) Do() (*T, error) {
	if s.wf.hasErrors() {
		return nil, nil
	}
	result, err := s.fn(s.arg)
	return result, s.processError(err)
}

// PE2: Pointer, Error - function takes 2 arguments and returns (*T, error)
type pe2[A any, B any, T any] struct {
	*Step
	fn   func(A, B) (*T, error)
	arg1 A
	arg2 B
}

func PE2[A any, B any, T any](wf *Workflow, action string, fn func(A, B) (*T, error), arg1 A, arg2 B) *pe2[A, B, T] {
	step := wf.add(action)
	return &pe2[A, B, T]{step, fn, arg1, arg2}
}

func (s *pe2[A, B, T]) Go() (*T, *Step) {
	var p T
	s.runAsync(func() {
		result, err := s.fn(s.arg1, s.arg2)
		s.processError(err)
		if result != nil {
			p = *result
		}
	})
	return &p, s.Step
}

func (s *pe2[A, B, T]) Do() (*T, error) {
	if s.wf.hasErrors() {
		return nil, nil
	}
	result, err := s.fn(s.arg1, s.arg2)
	return result, s.processError(err)
}

// PE3: Pointer, Error - function takes 3 arguments and returns (*T, error)
type pe3[A any, B any, C any, T any] struct {
	*Step
	fn   func(A, B, C) (*T, error)
	arg1 A
	arg2 B
	arg3 C
}

func PE3[A any, B any, C any, T any](wf *Workflow, action string, fn func(A, B, C) (*T, error), arg1 A, arg2 B, arg3 C) *pe3[A, B, C, T] {
	step := wf.add(action)
	return &pe3[A, B, C, T]{step, fn, arg1, arg2, arg3}
}

func (s *pe3[A, B, C, T]) Go() (*T, *Step) {
	var p T
	s.runAsync(func() {
		result, err := s.fn(s.arg1, s.arg2, s.arg3)
		s.processError(err)
		if result != nil {
			p = *result
		}
	})
	return &p, s.Step
}

func (s *pe3[A, B, C, T]) Do() (*T, error) {
	if s.wf.hasErrors() {
		return nil, nil
	}
	result, err := s.fn(s.arg1, s.arg2, s.arg3)
	return result, s.processError(err)
}

// PE4: Pointer, Error - function takes 4 arguments and returns (*T, error)
type pe4[A any, B any, C any, D any, T any] struct {
	*Step
	fn   func(A, B, C, D) (*T, error)
	arg1 A
	arg2 B
	arg3 C
	arg4 D
}

func PE4[A any, B any, C any, D any, T any](wf *Workflow, action string, fn func(A, B, C, D) (*T, error), arg1 A, arg2 B, arg3 C, arg4 D) *pe4[A, B, C, D, T] {
	step := wf.add(action)
	return &pe4[A, B, C, D, T]{step, fn, arg1, arg2, arg3, arg4}
}

func (s *pe4[A, B, C, D, T]) Go() (*T, *Step) {
	var p T
	s.runAsync(func() {
		result, err := s.fn(s.arg1, s.arg2, s.arg3, s.arg4)
		s.processError(err)
		if result != nil {
			p = *result
		}
	})
	return &p, s.Step
}

func (s *pe4[A, B, C, D, T]) Do() (*T, error) {
	if s.wf.hasErrors() {
		return nil, nil
	}
	result, err := s.fn(s.arg1, s.arg2, s.arg3, s.arg4)
	return result, s.processError(err)
}

// PE5: Pointer, Error - function takes 5 arguments and returns (*T, error)
type pe5[A any, B any, C any, D any, E any, T any] struct {
	*Step
	fn   func(A, B, C, D, E) (*T, error)
	arg1 A
	arg2 B
	arg3 C
	arg4 D
	arg5 E
}

func PE5[A any, B any, C any, D any, E any, T any](wf *Workflow, action string, fn func(A, B, C, D, E) (*T, error), arg1 A, arg2 B, arg3 C, arg4 D, arg5 E) *pe5[A, B, C, D, E, T] {
	step := wf.add(action)
	return &pe5[A, B, C, D, E, T]{step, fn, arg1, arg2, arg3, arg4, arg5}
}

func (s *pe5[A, B, C, D, E, T]) Go() (*T, *Step) {
	var p T
	s.runAsync(func() {
		result, err := s.fn(s.arg1, s.arg2, s.arg3, s.arg4, s.arg5)
		s.processError(err)
		if result != nil {
			p = *result
		}
	})
	return &p, s.Step
}

func (s *pe5[A, B, C, D, E, T]) Do() (*T, error) {
	if s.wf.hasErrors() {
		return nil, nil
	}
	result, err := s.fn(s.arg1, s.arg2, s.arg3, s.arg4, s.arg5)
	return result, s.processError(err)
}
