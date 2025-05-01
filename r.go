package wflow

// INIT
// r0
type r0[R1 any] struct {
	*step
	fn func() R1
}

func R0[R1 any](wf *Workflow, action string, fn func() R1, waitFor ...*step) *r0[R1] {
	step := wf.add(action, waitFor...)
	return &r0[R1]{step, fn}
}

// r1
type r1[R1 any, A1 any] struct {
	*step
	fn func(A1) R1
}

func R1[R1 any, A1 any](wf *Workflow, action string, fn func(A1) R1, waitFor ...*step) *r1[R1, A1] {
	step := wf.add(action, waitFor...)
	return &r1[R1, A1]{step, fn}
}

// r2
type r2[R1 any, A1 any, A2 any] struct {
	*step
	fn func(A1, A2) R1
}

func R2[R1 any, A1 any, A2 any](wf *Workflow, action string, fn func(A1, A2) R1, waitFor ...*step) *r2[R1, A1, A2] {
	step := wf.add(action, waitFor...)
	return &r2[R1, A1, A2]{step, fn}
}

// r3
type r3[R1 any, A1 any, A2 any, A3 any] struct {
	*step
	fn func(A1, A2, A3) R1
}

func R3[R1 any, A1 any, A2 any, A3 any](wf *Workflow, action string, fn func(A1, A2, A3) R1, waitFor ...*step) *r3[R1, A1, A2, A3] {
	step := wf.add(action, waitFor...)
	return &r3[R1, A1, A2, A3]{step, fn}
}

// r4
type r4[R1 any, A1 any, A2 any, A3 any, A4 any] struct {
	*step
	fn func(A1, A2, A3, A4) R1
}

func R4[R1 any, A1 any, A2 any, A3 any, A4 any](wf *Workflow, action string, fn func(A1, A2, A3, A4) R1, waitFor ...*step) *r4[R1, A1, A2, A3, A4] {
	step := wf.add(action, waitFor...)
	return &r4[R1, A1, A2, A3, A4]{step, fn}
}

// r5
type r5[R1 any, A1 any, A2 any, A3 any, A4 any, A5 any] struct {
	*step
	fn func(A1, A2, A3, A4, A5) R1
}

func R5[R1 any, A1 any, A2 any, A3 any, A4 any, A5 any](wf *Workflow, action string, fn func(A1, A2, A3, A4, A5) R1, waitFor ...*step) *r5[R1, A1, A2, A3, A4, A5] {
	step := wf.add(action, waitFor...)
	return &r5[R1, A1, A2, A3, A4, A5]{step, fn}
}

// rr0
type rr0[R1 any, R2 any] struct {
	*step
	fn func() (R1, R2)
}

func RR0[R1 any, R2 any](wf *Workflow, action string, fn func() (R1, R2), waitFor ...*step) *rr0[R1, R2] {
	step := wf.add(action, waitFor...)
	return &rr0[R1, R2]{step, fn}
}

// rr1
type rr1[R1 any, R2 any, A1 any] struct {
	*step
	fn func(A1) (R1, R2)
}

func RR1[R1 any, R2 any, A1 any](wf *Workflow, action string, fn func(A1) (R1, R2), waitFor ...*step) *rr1[R1, R2, A1] {
	step := wf.add(action, waitFor...)
	return &rr1[R1, R2, A1]{step, fn}
}

// rr2
type rr2[R1 any, R2 any, A1 any, A2 any] struct {
	*step
	fn func(A1, A2) (R1, R2)
}

func RR2[R1 any, R2 any, A1 any, A2 any](wf *Workflow, action string, fn func(A1, A2) (R1, R2), waitFor ...*step) *rr2[R1, R2, A1, A2] {
	step := wf.add(action, waitFor...)
	return &rr2[R1, R2, A1, A2]{step, fn}
}

// rr3
type rr3[R1 any, R2 any, A1 any, A2 any, A3 any] struct {
	*step
	fn func(A1, A2, A3) (R1, R2)
}

func RR3[R1 any, R2 any, A1 any, A2 any, A3 any](wf *Workflow, action string, fn func(A1, A2, A3) (R1, R2), waitFor ...*step) *rr3[R1, R2, A1, A2, A3] {
	step := wf.add(action, waitFor...)
	return &rr3[R1, R2, A1, A2, A3]{step, fn}
}

// rr4
type rr4[R1 any, R2 any, A1 any, A2 any, A3 any, A4 any] struct {
	*step
	fn func(A1, A2, A3, A4) (R1, R2)
}

func RR4[R1 any, R2 any, A1 any, A2 any, A3 any, A4 any](wf *Workflow, action string, fn func(A1, A2, A3, A4) (R1, R2), waitFor ...*step) *rr4[R1, R2, A1, A2, A3, A4] {
	step := wf.add(action, waitFor...)
	return &rr4[R1, R2, A1, A2, A3, A4]{step, fn}
}

// rr5
type rr5[R1 any, R2 any, A1 any, A2 any, A3 any, A4 any, A5 any] struct {
	*step
	fn func(A1, A2, A3, A4, A5) (R1, R2)
}

func RR5[R1 any, R2 any, A1 any, A2 any, A3 any, A4 any, A5 any](wf *Workflow, action string, fn func(A1, A2, A3, A4, A5) (R1, R2), waitFor ...*step) *rr5[R1, R2, A1, A2, A3, A4, A5] {
	step := wf.add(action, waitFor...)
	return &rr5[R1, R2, A1, A2, A3, A4, A5]{step, fn}
}

// rrr0
type rrr0[R1 any, R2 any, R3 any] struct {
	*step
	fn func() (R1, R2, R3)
}

func RRR0[R1 any, R2 any, R3 any](wf *Workflow, action string, fn func() (R1, R2, R3), waitFor ...*step) *rrr0[R1, R2, R3] {
	step := wf.add(action, waitFor...)
	return &rrr0[R1, R2, R3]{step, fn}
}

// rrr1
type rrr1[R1 any, R2 any, R3 any, A1 any] struct {
	*step
	fn func(A1) (R1, R2, R3)
}

func RRR1[R1 any, R2 any, R3 any, A1 any](wf *Workflow, action string, fn func(A1) (R1, R2, R3), waitFor ...*step) *rrr1[R1, R2, R3, A1] {
	step := wf.add(action, waitFor...)
	return &rrr1[R1, R2, R3, A1]{step, fn}
}

// rrr2
type rrr2[R1 any, R2 any, R3 any, A1 any, A2 any] struct {
	*step
	fn func(A1, A2) (R1, R2, R3)
}

func RRR2[R1 any, R2 any, R3 any, A1 any, A2 any](wf *Workflow, action string, fn func(A1, A2) (R1, R2, R3), waitFor ...*step) *rrr2[R1, R2, R3, A1, A2] {
	step := wf.add(action, waitFor...)
	return &rrr2[R1, R2, R3, A1, A2]{step, fn}
}

// rrr3
type rrr3[R1 any, R2 any, R3 any, A1 any, A2 any, A3 any] struct {
	*step
	fn func(A1, A2, A3) (R1, R2, R3)
}

func RRR3[R1 any, R2 any, R3 any, A1 any, A2 any, A3 any](wf *Workflow, action string, fn func(A1, A2, A3) (R1, R2, R3), waitFor ...*step) *rrr3[R1, R2, R3, A1, A2, A3] {
	step := wf.add(action, waitFor...)
	return &rrr3[R1, R2, R3, A1, A2, A3]{step, fn}
}

// rrr4
type rrr4[R1 any, R2 any, R3 any, A1 any, A2 any, A3 any, A4 any] struct {
	*step
	fn func(A1, A2, A3, A4) (R1, R2, R3)
}

func RRR4[R1 any, R2 any, R3 any, A1 any, A2 any, A3 any, A4 any](wf *Workflow, action string, fn func(A1, A2, A3, A4) (R1, R2, R3), waitFor ...*step) *rrr4[R1, R2, R3, A1, A2, A3, A4] {
	step := wf.add(action, waitFor...)
	return &rrr4[R1, R2, R3, A1, A2, A3, A4]{step, fn}
}

// rrr5
type rrr5[R1 any, R2 any, R3 any, A1 any, A2 any, A3 any, A4 any, A5 any] struct {
	*step
	fn func(A1, A2, A3, A4, A5) (R1, R2, R3)
}

func RRR5[R1 any, R2 any, R3 any, A1 any, A2 any, A3 any, A4 any, A5 any](wf *Workflow, action string, fn func(A1, A2, A3, A4, A5) (R1, R2, R3), waitFor ...*step) *rrr5[R1, R2, R3, A1, A2, A3, A4, A5] {
	step := wf.add(action, waitFor...)
	return &rrr5[R1, R2, R3, A1, A2, A3, A4, A5]{step, fn}
}

// DO
// r0
func (s *r0[R1]) Do() R1 {
	if s.shouldRun() {
		return s.fn()
	}
	return zero[R1]()
}

// r1
func (s *r1[R1, A1]) Do(a1 A1) R1 {
	if s.shouldRun() {
		return s.fn(a1)
	}
	return zero[R1]()
}

// r2
func (s *r2[R1, A1, A2]) Do(a1 A1, a2 A2) R1 {
	if s.shouldRun() {
		return s.fn(a1, a2)
	}
	return zero[R1]()
}

// r3
func (s *r3[R1, A1, A2, A3]) Do(a1 A1, a2 A2, a3 A3) R1 {
	if s.shouldRun() {
		return s.fn(a1, a2, a3)
	}
	return zero[R1]()
}

// r4
func (s *r4[R1, A1, A2, A3, A4]) Do(a1 A1, a2 A2, a3 A3, a4 A4) R1 {
	if s.shouldRun() {
		return s.fn(a1, a2, a3, a4)
	}
	return zero[R1]()
}

// r5
func (s *r5[R1, A1, A2, A3, A4, A5]) Do(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) R1 {
	if s.shouldRun() {
		return s.fn(a1, a2, a3, a4, a5)
	}
	return zero[R1]()
}

// rr0
func (s *rr0[R1, R2]) Do() (R1, R2) {
	if s.shouldRun() {
		return s.fn()
	}
	return zero2[R1, R2]()
}

// rr1
func (s *rr1[R1, R2, A1]) Do(a1 A1) (R1, R2) {
	if s.shouldRun() {
		return s.fn(a1)
	}
	return zero2[R1, R2]()
}

// rr2
func (s *rr2[R1, R2, A1, A2]) Do(a1 A1, a2 A2) (R1, R2) {
	if s.shouldRun() {
		return s.fn(a1, a2)
	}
	return zero2[R1, R2]()
}

// rr3
func (s *rr3[R1, R2, A1, A2, A3]) Do(a1 A1, a2 A2, a3 A3) (R1, R2) {
	if s.shouldRun() {
		return s.fn(a1, a2, a3)
	}
	return zero2[R1, R2]()
}

// rr4
func (s *rr4[R1, R2, A1, A2, A3, A4]) Do(a1 A1, a2 A2, a3 A3, a4 A4) (R1, R2) {
	if s.shouldRun() {
		return s.fn(a1, a2, a3, a4)
	}
	return zero2[R1, R2]()
}

// rr5
func (s *rr5[R1, R2, A1, A2, A3, A4, A5]) Do(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) (R1, R2) {
	if s.shouldRun() {
		return s.fn(a1, a2, a3, a4, a5)
	}
	return zero2[R1, R2]()
}

// rrr0
func (s *rrr0[R1, R2, R3]) Do() (R1, R2, R3) {
	if s.shouldRun() {
		return s.fn()
	}
	return zero3[R1, R2, R3]()
}

// rrr1
func (s *rrr1[R1, R2, R3, A1]) Do(a1 A1) (R1, R2, R3) {
	if s.shouldRun() {
		return s.fn(a1)
	}
	return zero3[R1, R2, R3]()
}

// rrr2
func (s *rrr2[R1, R2, R3, A1, A2]) Do(a1 A1, a2 A2) (R1, R2, R3) {
	if s.shouldRun() {
		return s.fn(a1, a2)
	}
	return zero3[R1, R2, R3]()
}

// rrr3
func (s *rrr3[R1, R2, R3, A1, A2, A3]) Do(a1 A1, a2 A2, a3 A3) (R1, R2, R3) {
	if s.shouldRun() {
		return s.fn(a1, a2, a3)
	}
	return zero3[R1, R2, R3]()
}

// rrr4
func (s *rrr4[R1, R2, R3, A1, A2, A3, A4]) Do(a1 A1, a2 A2, a3 A3, a4 A4) (R1, R2, R3) {
	if s.shouldRun() {
		return s.fn(a1, a2, a3, a4)
	}
	return zero3[R1, R2, R3]()
}

// rrr5
func (s *rrr5[R1, R2, R3, A1, A2, A3, A4, A5]) Do(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) (R1, R2, R3) {
	if s.shouldRun() {
		return s.fn(a1, a2, a3, a4, a5)
	}
	return zero3[R1, R2, R3]()
}

// GO
// r0
func (s *r0[R1]) Go() (*R1, *step) {
	var r1 R1
	s.runAsync(func() {
		r1 = s.fn()
	})
	return &r1, s.step
}

// r1
func (s *r1[R1, A1]) Go(a1 A1) (*R1, *step) {
	var r1 R1
	s.runAsync(func() {
		r1 = s.fn(a1)
	})
	return &r1, s.step
}

// r2
func (s *r2[R1, A1, A2]) Go(a1 A1, a2 A2) (*R1, *step) {
	var r1 R1
	s.runAsync(func() {
		r1 = s.fn(a1, a2)
	})
	return &r1, s.step
}

// r3
func (s *r3[R1, A1, A2, A3]) Go(a1 A1, a2 A2, a3 A3) (*R1, *step) {
	var r1 R1
	s.runAsync(func() {
		r1 = s.fn(a1, a2, a3)
	})
	return &r1, s.step
}

// r4
func (s *r4[R1, A1, A2, A3, A4]) Go(a1 A1, a2 A2, a3 A3, a4 A4) (*R1, *step) {
	var r1 R1
	s.runAsync(func() {
		r1 = s.fn(a1, a2, a3, a4)
	})
	return &r1, s.step
}

// r5
func (s *r5[R1, A1, A2, A3, A4, A5]) Go(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) (*R1, *step) {
	var r1 R1
	s.runAsync(func() {
		r1 = s.fn(a1, a2, a3, a4, a5)
	})
	return &r1, s.step
}

// rr0
func (s *rr0[R1, R2]) Go() (*R1, *R2, *step) {
	var r1 R1
	var r2 R2
	s.runAsync(func() {
		r1, r2 = s.fn()
	})
	return &r1, &r2, s.step
}

// rr1
func (s *rr1[R1, R2, A1]) Go(a1 A1) (*R1, *R2, *step) {
	var r1 R1
	var r2 R2
	s.runAsync(func() {
		r1, r2 = s.fn(a1)
	})
	return &r1, &r2, s.step
}

// rr2
func (s *rr2[R1, R2, A1, A2]) Go(a1 A1, a2 A2) (*R1, *R2, *step) {
	var r1 R1
	var r2 R2
	s.runAsync(func() {
		r1, r2 = s.fn(a1, a2)
	})
	return &r1, &r2, s.step
}

// rr3
func (s *rr3[R1, R2, A1, A2, A3]) Go(a1 A1, a2 A2, a3 A3) (*R1, *R2, *step) {
	var r1 R1
	var r2 R2
	s.runAsync(func() {
		r1, r2 = s.fn(a1, a2, a3)
	})
	return &r1, &r2, s.step
}

// rr4
func (s *rr4[R1, R2, A1, A2, A3, A4]) Go(a1 A1, a2 A2, a3 A3, a4 A4) (*R1, *R2, *step) {
	var r1 R1
	var r2 R2
	s.runAsync(func() {
		r1, r2 = s.fn(a1, a2, a3, a4)
	})
	return &r1, &r2, s.step
}

// rr5
func (s *rr5[R1, R2, A1, A2, A3, A4, A5]) Go(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) (*R1, *R2, *step) {
	var r1 R1
	var r2 R2
	s.runAsync(func() {
		r1, r2 = s.fn(a1, a2, a3, a4, a5)
	})
	return &r1, &r2, s.step
}

// rrr0
func (s *rrr0[R1, R2, R3]) Go() (*R1, *R2, *R3, *step) {
	var r1 R1
	var r2 R2
	var r3 R3
	s.runAsync(func() {
		r1, r2, r3 = s.fn()
	})
	return &r1, &r2, &r3, s.step
}

// rrr1
func (s *rrr1[R1, R2, R3, A1]) Go(a1 A1) (*R1, *R2, *R3, *step) {
	var r1 R1
	var r2 R2
	var r3 R3
	s.runAsync(func() {
		r1, r2, r3 = s.fn(a1)
	})
	return &r1, &r2, &r3, s.step
}

// rrr2
func (s *rrr2[R1, R2, R3, A1, A2]) Go(a1 A1, a2 A2) (*R1, *R2, *R3, *step) {
	var r1 R1
	var r2 R2
	var r3 R3
	s.runAsync(func() {
		r1, r2, r3 = s.fn(a1, a2)
	})
	return &r1, &r2, &r3, s.step
}

// rrr3
func (s *rrr3[R1, R2, R3, A1, A2, A3]) Go(a1 A1, a2 A2, a3 A3) (*R1, *R2, *R3, *step) {
	var r1 R1
	var r2 R2
	var r3 R3
	s.runAsync(func() {
		r1, r2, r3 = s.fn(a1, a2, a3)
	})
	return &r1, &r2, &r3, s.step
}

// rrr4
func (s *rrr4[R1, R2, R3, A1, A2, A3, A4]) Go(a1 A1, a2 A2, a3 A3, a4 A4) (*R1, *R2, *R3, *step) {
	var r1 R1
	var r2 R2
	var r3 R3
	s.runAsync(func() {
		r1, r2, r3 = s.fn(a1, a2, a3, a4)
	})
	return &r1, &r2, &r3, s.step
}

// rrr5
func (s *rrr5[R1, R2, R3, A1, A2, A3, A4, A5]) Go(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) (*R1, *R2, *R3, *step) {
	var r1 R1
	var r2 R2
	var r3 R3
	s.runAsync(func() {
		r1, r2, r3 = s.fn(a1, a2, a3, a4, a5)
	})
	return &r1, &r2, &r3, s.step
}
