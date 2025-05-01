package wflow

// INIT
// e0
type re0[R1 any] struct {
	*step
	fn func() (R1, error)
}

// re0
func RE0[R1 any](wf *Workflow, action string, fn func() (R1, error), waitFor ...*step) *re0[R1] {
	step := wf.add(action, waitFor...)
	return &re0[R1]{step, fn}
}

// re1
type re1[R1, A1 any] struct {
	*step
	fn func(A1) (R1, error)
}

func RE1[R1, A1 any](wf *Workflow, action string, fn func(A1) (R1, error), waitFor ...*step) *re1[R1, A1] {
	step := wf.add(action, waitFor...)
	return &re1[R1, A1]{step, fn}
}

// re2
type re2[R1, A1, A2 any] struct {
	*step
	fn func(A1, A2) (R1, error)
}

func RE2[R1, A1, A2 any](wf *Workflow, action string, fn func(A1, A2) (R1, error), waitFor ...*step) *re2[R1, A1, A2] {
	step := wf.add(action, waitFor...)
	return &re2[R1, A1, A2]{step, fn}
}

// re3
type re3[R1, A1, A2, A3 any] struct {
	*step
	fn func(A1, A2, A3) (R1, error)
}

func RE3[R1, A1, A2, A3 any](wf *Workflow, action string, fn func(A1, A2, A3) (R1, error), waitFor ...*step) *re3[R1, A1, A2, A3] {
	step := wf.add(action, waitFor...)
	return &re3[R1, A1, A2, A3]{step, fn}
}

// re4
type re4[R1, A1, A2, A3, A4 any] struct {
	*step
	fn func(A1, A2, A3, A4) (R1, error)
}

func RE4[R1, A1, A2, A3, A4 any](wf *Workflow, action string, fn func(A1, A2, A3, A4) (R1, error), waitFor ...*step) *re4[R1, A1, A2, A3, A4] {
	step := wf.add(action, waitFor...)
	return &re4[R1, A1, A2, A3, A4]{step, fn}
}

// re5
type re5[R1, A1, A2, A3, A4, A5 any] struct {
	*step
	fn func(A1, A2, A3, A4, A5) (R1, error)
}

func RE5[R1, A1, A2, A3, A4, A5 any](wf *Workflow, action string, fn func(A1, A2, A3, A4, A5) (R1, error), waitFor ...*step) *re5[R1, A1, A2, A3, A4, A5] {
	step := wf.add(action, waitFor...)
	return &re5[R1, A1, A2, A3, A4, A5]{step, fn}
}

// rre0
type rre0[R1 any, R2 any] struct {
	*step
	fn func() (R1, R2, error)
}

func RRE0[R1 any, R2 any](wf *Workflow, action string, fn func() (R1, R2, error), waitFor ...*step) *rre0[R1, R2] {
	step := wf.add(action, waitFor...)
	return &rre0[R1, R2]{step, fn}
}

// rre1
type rre1[R1 any, R2 any, A1 any] struct {
	*step
	fn func(A1) (R1, R2, error)
}

func RRE1[R1 any, R2 any, A1 any](wf *Workflow, action string, fn func(A1) (R1, R2, error), waitFor ...*step) *rre1[R1, R2, A1] {
	step := wf.add(action, waitFor...)
	return &rre1[R1, R2, A1]{step, fn}
}

// rre2
type rre2[R1 any, R2 any, A1 any, A2 any] struct {
	*step
	fn func(A1, A2) (R1, R2, error)
}

func RRE2[R1 any, R2 any, A1 any, A2 any](wf *Workflow, action string, fn func(A1, A2) (R1, R2, error), waitFor ...*step) *rre2[R1, R2, A1, A2] {
	step := wf.add(action, waitFor...)
	return &rre2[R1, R2, A1, A2]{step, fn}
}

// rre3
type rre3[R1 any, R2 any, A1 any, A2 any, A3 any] struct {
	*step
	fn func(A1, A2, A3) (R1, R2, error)
}

func RRE3[R1 any, R2 any, A1 any, A2 any, A3 any](wf *Workflow, action string, fn func(A1, A2, A3) (R1, R2, error), waitFor ...*step) *rre3[R1, R2, A1, A2, A3] {
	step := wf.add(action, waitFor...)
	return &rre3[R1, R2, A1, A2, A3]{step, fn}
}

// rre4
type rre4[R1 any, R2 any, A1 any, A2 any, A3 any, A4 any] struct {
	*step
	fn func(A1, A2, A3, A4) (R1, R2, error)
}

func RRE4[R1 any, R2 any, A1 any, A2 any, A3 any, A4 any](wf *Workflow, action string, fn func(A1, A2, A3, A4) (R1, R2, error), waitFor ...*step) *rre4[R1, R2, A1, A2, A3, A4] {
	step := wf.add(action, waitFor...)
	return &rre4[R1, R2, A1, A2, A3, A4]{step, fn}
}

// rre5
type rre5[R1 any, R2 any, A1 any, A2 any, A3 any, A4 any, A5 any] struct {
	*step
	fn func(A1, A2, A3, A4, A5) (R1, R2, error)
}

func RRE5[R1 any, R2 any, A1 any, A2 any, A3 any, A4 any, A5 any](wf *Workflow, action string, fn func(A1, A2, A3, A4, A5) (R1, R2, error), waitFor ...*step) *rre5[R1, R2, A1, A2, A3, A4, A5] {
	step := wf.add(action, waitFor...)
	return &rre5[R1, R2, A1, A2, A3, A4, A5]{step, fn}
}

// rrre0
type rrre0[R1 any, R2 any, R3 any] struct {
	*step
	fn func() (R1, R2, R3, error)
}

func RRRE0[R1 any, R2 any, R3 any](wf *Workflow, action string, fn func() (R1, R2, R3, error), waitFor ...*step) *rrre0[R1, R2, R3] {
	step := wf.add(action, waitFor...)
	return &rrre0[R1, R2, R3]{step, fn}
}

// rrre1
type rrre1[R1 any, R2 any, R3 any, A1 any] struct {
	*step
	fn func(A1) (R1, R2, R3, error)
}

func RRRE1[R1 any, R2 any, R3 any, A1 any](wf *Workflow, action string, fn func(A1) (R1, R2, R3, error), waitFor ...*step) *rrre1[R1, R2, R3, A1] {
	step := wf.add(action, waitFor...)
	return &rrre1[R1, R2, R3, A1]{step, fn}
}

// rrre2
type rrre2[R1 any, R2 any, R3 any, A1 any, A2 any] struct {
	*step
	fn func(A1, A2) (R1, R2, R3, error)
}

func RRRE2[R1 any, R2 any, R3 any, A1 any, A2 any](wf *Workflow, action string, fn func(A1, A2) (R1, R2, R3, error), waitFor ...*step) *rrre2[R1, R2, R3, A1, A2] {
	step := wf.add(action, waitFor...)
	return &rrre2[R1, R2, R3, A1, A2]{step, fn}
}

// rrre3
type rrre3[R1 any, R2 any, R3 any, A1 any, A2 any, A3 any] struct {
	*step
	fn func(A1, A2, A3) (R1, R2, R3, error)
}

func RRRE3[R1 any, R2 any, R3 any, A1 any, A2 any, A3 any](wf *Workflow, action string, fn func(A1, A2, A3) (R1, R2, R3, error), waitFor ...*step) *rrre3[R1, R2, R3, A1, A2, A3] {
	step := wf.add(action, waitFor...)
	return &rrre3[R1, R2, R3, A1, A2, A3]{step, fn}
}

// rrre4
type rrre4[R1 any, R2 any, R3 any, A1 any, A2 any, A3 any, A4 any] struct {
	*step
	fn func(A1, A2, A3, A4) (R1, R2, R3, error)
}

func RRRE4[R1 any, R2 any, R3 any, A1 any, A2 any, A3 any, A4 any](wf *Workflow, action string, fn func(A1, A2, A3, A4) (R1, R2, R3, error), waitFor ...*step) *rrre4[R1, R2, R3, A1, A2, A3, A4] {
	step := wf.add(action, waitFor...)
	return &rrre4[R1, R2, R3, A1, A2, A3, A4]{step, fn}
}

// rrre5
type rrre5[R1 any, R2 any, R3 any, A1 any, A2 any, A3 any, A4 any, A5 any] struct {
	*step
	fn func(A1, A2, A3, A4, A5) (R1, R2, R3, error)
}

func RRRE5[R1 any, R2 any, R3 any, A1 any, A2 any, A3 any, A4 any, A5 any](wf *Workflow, action string, fn func(A1, A2, A3, A4, A5) (R1, R2, R3, error), waitFor ...*step) *rrre5[R1, R2, R3, A1, A2, A3, A4, A5] {
	step := wf.add(action, waitFor...)
	return &rrre5[R1, R2, R3, A1, A2, A3, A4, A5]{step, fn}
}

// DO
// re0
func (s *re0[R1]) Do() (R1, error) {
	if s.shouldRun() {
		r1, err := s.fn()
		return r1, s.processError(err)
	}
	return zero[R1](), nil
}

// re1
func (s *re1[R1, A1]) Do(a1 A1) (R1, error) {
	if s.shouldRun() {
		r1, err := s.fn(a1)
		return r1, s.processError(err)
	}
	return zero[R1](), nil
}

// re2
func (s *re2[R1, A1, A2]) Do(a1 A1, a2 A2) (R1, error) {
	if s.shouldRun() {
		r1, err := s.fn(a1, a2)
		return r1, s.processError(err)
	}
	return zero[R1](), nil
}

// re3
func (s *re3[R1, A1, A2, A3]) Do(a1 A1, a2 A2, a3 A3) (R1, error) {
	if s.shouldRun() {
		r1, err := s.fn(a1, a2, a3)
		return r1, s.processError(err)
	}
	return zero[R1](), nil
}

// re4
func (s *re4[R1, A1, A2, A3, A4]) Do(a1 A1, a2 A2, a3 A3, a4 A4) (R1, error) {
	if s.shouldRun() {
		r1, err := s.fn(a1, a2, a3, a4)
		return r1, s.processError(err)
	}
	return zero[R1](), nil
}

// re5
func (s *re5[R1, A1, A2, A3, A4, A5]) Do(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) (R1, error) {
	if s.shouldRun() {
		r1, err := s.fn(a1, a2, a3, a4, a5)
		return r1, s.processError(err)
	}
	return zero[R1](), nil
}

// rre0
func (s *rre0[R1, R2]) Do() (R1, R2, error) {
	if s.shouldRun() {
		r1, r2, err := s.fn()
		return r1, r2, s.processError(err)
	}
	return zero[R1](), zero[R2](), nil
}

// rre1
func (s *rre1[R1, R2, A1]) Do(a1 A1) (R1, R2, error) {
	if s.shouldRun() {
		r1, r2, err := s.fn(a1)
		return r1, r2, s.processError(err)
	}
	return zero[R1](), zero[R2](), nil
}

// rre2
func (s *rre2[R1, R2, A1, A2]) Do(a1 A1, a2 A2) (R1, R2, error) {
	if s.shouldRun() {
		r1, r2, err := s.fn(a1, a2)
		return r1, r2, s.processError(err)
	}
	return zero[R1](), zero[R2](), nil
}

// rre3
func (s *rre3[R1, R2, A1, A2, A3]) Do(a1 A1, a2 A2, a3 A3) (R1, R2, error) {
	if s.shouldRun() {
		r1, r2, err := s.fn(a1, a2, a3)
		return r1, r2, s.processError(err)
	}
	return zero[R1](), zero[R2](), nil
}

// rre4
func (s *rre4[R1, R2, A1, A2, A3, A4]) Do(a1 A1, a2 A2, a3 A3, a4 A4) (R1, R2, error) {
	if s.shouldRun() {
		r1, r2, err := s.fn(a1, a2, a3, a4)
		return r1, r2, s.processError(err)
	}
	return zero[R1](), zero[R2](), nil
}

// rre5
func (s *rre5[R1, R2, A1, A2, A3, A4, A5]) Do(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) (R1, R2, error) {
	if s.shouldRun() {
		r1, r2, err := s.fn(a1, a2, a3, a4, a5)
		return r1, r2, s.processError(err)
	}
	return zero[R1](), zero[R2](), nil
}

// rrre0
func (s *rrre0[R1, R2, R3]) Do() (R1, R2, R3, error) {
	if s.shouldRun() {
		r1, r2, r3, err := s.fn()
		return r1, r2, r3, s.processError(err)
	}
	return zero[R1](), zero[R2](), zero[R3](), nil
}

// rrre1
func (s *rrre1[R1, R2, R3, A1]) Do(a1 A1) (R1, R2, R3, error) {
	if s.shouldRun() {
		r1, r2, r3, err := s.fn(a1)
		return r1, r2, r3, s.processError(err)
	}
	return zero[R1](), zero[R2](), zero[R3](), nil
}

// rrre2
func (s *rrre2[R1, R2, R3, A1, A2]) Do(a1 A1, a2 A2) (R1, R2, R3, error) {
	if s.shouldRun() {
		r1, r2, r3, err := s.fn(a1, a2)
		return r1, r2, r3, s.processError(err)
	}
	return zero[R1](), zero[R2](), zero[R3](), nil
}

// rrre3
func (s *rrre3[R1, R2, R3, A1, A2, A3]) Do(a1 A1, a2 A2, a3 A3) (R1, R2, R3, error) {
	if s.shouldRun() {
		r1, r2, r3, err := s.fn(a1, a2, a3)
		return r1, r2, r3, s.processError(err)
	}
	return zero[R1](), zero[R2](), zero[R3](), nil
}

// rrre4
func (s *rrre4[R1, R2, R3, A1, A2, A3, A4]) Do(a1 A1, a2 A2, a3 A3, a4 A4) (R1, R2, R3, error) {
	if s.shouldRun() {
		r1, r2, r3, err := s.fn(a1, a2, a3, a4)
		return r1, r2, r3, s.processError(err)
	}
	return zero[R1](), zero[R2](), zero[R3](), nil
}

// rrre5
func (s *rrre5[R1, R2, R3, A1, A2, A3, A4, A5]) Do(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) (R1, R2, R3, error) {
	if s.shouldRun() {
		r1, r2, r3, err := s.fn(a1, a2, a3, a4, a5)
		return r1, r2, r3, s.processError(err)
	}
	return zero[R1](), zero[R2](), zero[R3](), nil
}

// GO
// re0
func (s *re0[R1]) Go() (*R1, *step) {
	var r1 R1
	s.runAsync(func() {
		var err error
		r1, err = s.fn()
		s.processError(err)
	})
	return &r1, s.step
}

// re1
func (s *re1[R1, A1]) Go(a1 A1) (*R1, *step) {
	var r1 R1
	s.runAsync(func() {
		var err error
		r1, err = s.fn(a1)
		s.processError(err)
	})
	return &r1, s.step
}

// re2
func (s *re2[R1, A1, A2]) Go(a1 A1, a2 A2) (*R1, *step) {
	var r1 R1
	s.runAsync(func() {
		var err error
		r1, err = s.fn(a1, a2)
		s.processError(err)
	})
	return &r1, s.step
}

// re3
func (s *re3[R1, A1, A2, A3]) Go(a1 A1, a2 A2, a3 A3) (*R1, *step) {
	var r1 R1
	s.runAsync(func() {
		var err error
		r1, err = s.fn(a1, a2, a3)
		s.processError(err)
	})
	return &r1, s.step
}

// re4
func (s *re4[R1, A1, A2, A3, A4]) Go(a1 A1, a2 A2, a3 A3, a4 A4) (*R1, *step) {
	var r1 R1
	s.runAsync(func() {
		var err error
		r1, err = s.fn(a1, a2, a3, a4)
		s.processError(err)
	})
	return &r1, s.step
}

// re5
func (s *re5[R1, A1, A2, A3, A4, A5]) Go(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) (*R1, *step) {
	var r1 R1
	s.runAsync(func() {
		var err error
		r1, err = s.fn(a1, a2, a3, a4, a5)
		s.processError(err)
	})
	return &r1, s.step
}

// rre0
func (s *rre0[R1, R2]) Go() (*R1, *R2, *step) {
	var r1 R1
	var r2 R2
	s.runAsync(func() {
		var err error
		r1, r2, err = s.fn()
		s.processError(err)
	})
	return &r1, &r2, s.step
}

// rre1
func (s *rre1[R1, R2, A1]) Go(a1 A1) (*R1, *R2, *step) {
	var r1 R1
	var r2 R2
	s.runAsync(func() {
		var err error
		r1, r2, err = s.fn(a1)
		s.processError(err)
	})
	return &r1, &r2, s.step
}

// rre2
func (s *rre2[R1, R2, A1, A2]) Go(a1 A1, a2 A2) (*R1, *R2, *step) {
	var r1 R1
	var r2 R2
	s.runAsync(func() {
		var err error
		r1, r2, err = s.fn(a1, a2)
		s.processError(err)
	})
	return &r1, &r2, s.step
}

// rre3
func (s *rre3[R1, R2, A1, A2, A3]) Go(a1 A1, a2 A2, a3 A3) (*R1, *R2, *step) {
	var r1 R1
	var r2 R2
	s.runAsync(func() {
		var err error
		r1, r2, err = s.fn(a1, a2, a3)
		s.processError(err)
	})
	return &r1, &r2, s.step
}

// rre4
func (s *rre4[R1, R2, A1, A2, A3, A4]) Go(a1 A1, a2 A2, a3 A3, a4 A4) (*R1, *R2, *step) {
	var r1 R1
	var r2 R2
	s.runAsync(func() {
		var err error
		r1, r2, err = s.fn(a1, a2, a3, a4)
		s.processError(err)
	})
	return &r1, &r2, s.step
}

// rre5
func (s *rre5[R1, R2, A1, A2, A3, A4, A5]) Go(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) (*R1, *R2, *step) {
	var r1 R1
	var r2 R2
	s.runAsync(func() {
		var err error
		r1, r2, err = s.fn(a1, a2, a3, a4, a5)
		s.processError(err)
	})
	return &r1, &r2, s.step
}

// rrre0
func (s *rrre0[R1, R2, R3]) Go() (*R1, *R2, *R3, *step) {
	var r1 R1
	var r2 R2
	var r3 R3
	s.runAsync(func() {
		var err error
		r1, r2, r3, err = s.fn()
		s.processError(err)
	})
	return &r1, &r2, &r3, s.step
}

// rrre1
func (s *rrre1[R1, R2, R3, A1]) Go(a1 A1) (*R1, *R2, *R3, *step) {
	var r1 R1
	var r2 R2
	var r3 R3
	s.runAsync(func() {
		var err error
		r1, r2, r3, err = s.fn(a1)
		s.processError(err)
	})
	return &r1, &r2, &r3, s.step
}

// rrre2
func (s *rrre2[R1, R2, R3, A1, A2]) Go(a1 A1, a2 A2) (*R1, *R2, *R3, *step) {
	var r1 R1
	var r2 R2
	var r3 R3
	s.runAsync(func() {
		var err error
		r1, r2, r3, err = s.fn(a1, a2)
		s.processError(err)
	})
	return &r1, &r2, &r3, s.step
}

// rrre3
func (s *rrre3[R1, R2, R3, A1, A2, A3]) Go(a1 A1, a2 A2, a3 A3) (*R1, *R2, *R3, *step) {
	var r1 R1
	var r2 R2
	var r3 R3
	s.runAsync(func() {
		var err error
		r1, r2, r3, err = s.fn(a1, a2, a3)
		s.processError(err)
	})
	return &r1, &r2, &r3, s.step
}

// rrre4
func (s *rrre4[R1, R2, R3, A1, A2, A3, A4]) Go(a1 A1, a2 A2, a3 A3, a4 A4) (*R1, *R2, *R3, *step) {
	var r1 R1
	var r2 R2
	var r3 R3
	s.runAsync(func() {
		var err error
		r1, r2, r3, err = s.fn(a1, a2, a3, a4)
		s.processError(err)
	})
	return &r1, &r2, &r3, s.step
}

// rrre5
func (s *rrre5[R1, R2, R3, A1, A2, A3, A4, A5]) Go(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) (*R1, *R2, *R3, *step) {
	var r1 R1
	var r2 R2
	var r3 R3
	s.runAsync(func() {
		var err error
		r1, r2, r3, err = s.fn(a1, a2, a3, a4, a5)
		s.processError(err)
	})
	return &r1, &r2, &r3, s.step
}
