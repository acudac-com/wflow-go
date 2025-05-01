package wflow

// INIT
// e0
type e0 struct {
	*step
	fn func() error
}

func E0(wf *Workflow, action string, fn func() error, waitFor ...*step) *e0 {
	step := wf.add(action, waitFor...)
	return &e0{step, fn}
}

// e1
type e1[A1 any] struct {
	*step
	fn func(A1) error
}

func E1[A1 any](wf *Workflow, action string, fn func(A1) error, waitFor ...*step) *e1[A1] {
	step := wf.add(action, waitFor...)
	return &e1[A1]{step, fn}
}

// e2
type e2[A1, A2 any] struct {
	*step
	fn func(A1, A2) error
}

func E2[A1, A2 any](wf *Workflow, action string, fn func(A1, A2) error, waitFor ...*step) *e2[A1, A2] {
	step := wf.add(action, waitFor...)
	return &e2[A1, A2]{step, fn}
}

// e3
type e3[A1, A2, A3 any] struct {
	*step
	fn func(A1, A2, A3) error
}

func E3[A1, A2, A3 any](wf *Workflow, action string, fn func(A1, A2, A3) error, waitFor ...*step) *e3[A1, A2, A3] {
	step := wf.add(action, waitFor...)
	return &e3[A1, A2, A3]{step, fn}
}

// e4
type e4[A1, A2, A3, A4 any] struct {
	*step
	fn func(A1, A2, A3, A4) error
}

func E4[A1, A2, A3, A4 any](wf *Workflow, action string, fn func(A1, A2, A3, A4) error, waitFor ...*step) *e4[A1, A2, A3, A4] {
	step := wf.add(action, waitFor...)
	return &e4[A1, A2, A3, A4]{step, fn}
}

// e5
type e5[A1, A2, A3, A4, A5 any] struct {
	*step
	fn func(A1, A2, A3, A4, A5) error
}

func E5[A1, A2, A3, A4, A5 any](wf *Workflow, action string, fn func(A1, A2, A3, A4, A5) error, waitFor ...*step) *e5[A1, A2, A3, A4, A5] {
	step := wf.add(action, waitFor...)
	return &e5[A1, A2, A3, A4, A5]{step, fn}
}

// DO
// e0
func (s *e0) Do() error {
	if s.shouldRun() {
		return s.processError(s.fn())
	}
	return nil
}

// e1
func (s *e1[A1]) Do(a1 A1) error {
	if s.shouldRun() {
		return s.processError(s.fn(a1))
	}
	return nil
}

// e2
func (s *e2[A1, A2]) Do(a1 A1, a2 A2) error {
	if s.shouldRun() {
		return s.processError(s.fn(a1, a2))
	}
	return nil
}

// e3
func (s *e3[A1, A2, A3]) Do(a1 A1, a2 A2, a3 A3) error {
	if s.shouldRun() {
		return s.processError(s.fn(a1, a2, a3))
	}
	return nil
}

// e4
func (s *e4[A1, A2, A3, A4]) Do(a1 A1, a2 A2, a3 A3, a4 A4) error {
	if s.shouldRun() {
		return s.processError(s.fn(a1, a2, a3, a4))
	}
	return nil
}

// e5
func (s *e5[A1, A2, A3, A4, A5]) Do(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) error {
	if s.shouldRun() {
		return s.processError(s.fn(a1, a2, a3, a4, a5))
	}
	return nil
}

// GO
// e0
func (s *e0) Go() *step {
	s.runAsync(func() {
		s.processError(s.fn())
	})
	return s.step
}

// e1
func (s *e1[A1]) Go(a1 A1) *step {
	s.runAsync(func() {
		s.processError(s.fn(a1))
	})
	return s.step
}

// e2
func (s *e2[A1, A2]) Go(a1 A1, a2 A2) *step {
	s.runAsync(func() {
		s.processError(s.fn(a1, a2))
	})
	return s.step
}

// e3
func (s *e3[A1, A2, A3]) Go(a1 A1, a2 A2, a3 A3) *step {
	s.runAsync(func() {
		s.processError(s.fn(a1, a2, a3))
	})
	return s.step
}

// e4
func (s *e4[A1, A2, A3, A4]) Go(a1 A1, a2 A2, a3 A3, a4 A4) *step {
	s.runAsync(func() {
		s.processError(s.fn(a1, a2, a3, a4))
	})
	return s.step
}

// e5
func (s *e5[A1, A2, A3, A4, A5]) Go(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) *step {
	s.runAsync(func() {
		s.processError(s.fn(a1, a2, a3, a4, a5))
	})
	return s.step
}
