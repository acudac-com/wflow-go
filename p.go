package wflow

// INIT
// p
type p[P any] struct {
	*Step
	fn func() *P
}

func P[P any](wf *Workflow, action string, fn func() *P) *p[P] {
	step := wf.add(action)
	return &p[P]{step, fn}
}

type pe[T any] struct {
	*Step
	fn func() (*T, error)
}

func PE[T any](wf *Workflow, action string, fn func() (*T, error)) *pe[T] {
	step := wf.add(action)
	return &pe[T]{step, fn}
}

// p1
type p1[P any, A1 any] struct {
	*Step
	fn func(A1) *P
}

func P1[P any, A1 any](wf *Workflow, action string, fn func(A1) *P) *p1[P, A1] {
	step := wf.add(action)
	return &p1[P, A1]{step, fn}
}

type pe1[P any, A1 any] struct {
	*Step
	fn func(A1) (*P, error)
}

func PE1[P any, A1 any](wf *Workflow, action string, fn func(A1) (*P, error)) *pe1[P, A1] {
	step := wf.add(action)
	return &pe1[P, A1]{step, fn}
}

// p2
type p2[P any, A1 any, A2 any] struct {
	*Step
	fn func(A1, A2) *P
}

func P2[P any, A1 any, A2 any](wf *Workflow, action string, fn func(A1, A2) *P) *p2[P, A1, A2] {
	step := wf.add(action)
	return &p2[P, A1, A2]{step, fn}
}

type pe2[P any, A1 any, A2 any] struct {
	*Step
	fn func(A1, A2) (*P, error)
}

func PE2[P any, A1 any, A2 any](wf *Workflow, action string, fn func(A1, A2) (*P, error)) *pe2[P, A1, A2] {
	step := wf.add(action)
	return &pe2[P, A1, A2]{step, fn}
}

// GO
// p
func (s *p[P]) Go() (*P, *Step) {
	var p P
	s.runAsync(func() {
		result := s.fn()
		s.processError(nil)
		if result != nil {
			p = *result
		}
	})
	return &p, s.Step
}

func (s *pe[P]) Go() (*P, *Step) {
	var p P
	s.runAsync(func() {
		result, err := s.fn()
		s.processError(err)
		if result != nil {
			p = *result
		}
	})
	return &p, s.Step
}

// p1
func (s *p1[P, A1]) Go(a1 A1) (*P, *Step) {
	var p P
	s.runAsync(func() {
		result := s.fn(a1)
		s.processError(nil)
		if result != nil {
			p = *result
		}
	})
	return &p, s.Step
}

func (s *pe1[P, A1]) Go(a1 A1) (*P, *Step) {
	var p P
	s.runAsync(func() {
		result, err := s.fn(a1)
		s.processError(err)
		if result != nil {
			p = *result
		}
	})
	return &p, s.Step
}

// p2
func (s *p2[P, A1, A2]) Go(a1 A1, a2 A2) (*P, *Step) {
	var p P
	s.runAsync(func() {
		result := s.fn(a1, a2)
		s.processError(nil)
		if result != nil {
			p = *result
		}
	})
	return &p, s.Step
}

func (s *pe2[P, A1, A2]) Go(a1 A1, a2 A2) (*P, *Step) {
	var p P
	s.runAsync(func() {
		result, err := s.fn(a1, a2)
		s.processError(err)
		if result != nil {
			p = *result
		}
	})
	return &p, s.Step
}

// DO
// p
func (s *p[P]) Do() *P {
	if s.shouldRun() {
		result := s.fn()
		return result
	}
	return nil
}

func (s *pe[P]) Do() (*P, error) {
	if s.shouldRun() {
		result, err := s.fn()
		return result, s.processError(err)
	}
	return nil, nil
}

// p1
func (s *p1[P, A1]) Do(a1 A1) *P {
	if s.shouldRun() {
		result := s.fn(a1)
		return result
	}
	return nil
}

func (s *pe1[P, A1]) Do(a1 A1) (*P, error) {
	if s.shouldRun() {
		result, err := s.fn(a1)
		return result, s.processError(err)
	}
	return nil, nil
}

// p2
func (s *p2[P, A1, A2]) Do(a1 A1, a2 A2) *P {
	if s.shouldRun() {
		return s.fn(a1, a2)
	}
	return nil
}

func (s *pe2[P, A1, A2]) Do(a1 A1, a2 A2) (*P, error) {
	if s.shouldRun() {
		result, err := s.fn(a1, a2)
		return result, s.processError(err)
	}
	return nil, nil
}
