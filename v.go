package wflow

// INIT
// v
type v[V any] struct {
	*step
	fn func() V
}

func V[V any](wf *Workflow, action string, fn func() V) *v[V] {
	step := wf.add(action)
	return &v[V]{step, fn}
}

type ve[V any] struct {
	*step
	fn func() (V, error)
}

func VE[V any](wf *Workflow, action string, fn func() (V, error)) *ve[V] {
	step := wf.add(action)
	return &ve[V]{step, fn}
}

// v1
type v1[V any, A1 any] struct {
	*step
	fn func(A1) V
}

func V1[V any, A1 any](wf *Workflow, action string, fn func(A1) V) *v1[V, A1] {
	step := wf.add(action)
	return &v1[V, A1]{step, fn}
}

type ve1[V any, A1 any] struct {
	*step
	fn func(A1) (V, error)
}

func VE1[V any, A1 any](wf *Workflow, action string, fn func(A1) (V, error)) *ve1[V, A1] {
	step := wf.add(action)
	return &ve1[V, A1]{step, fn}
}

// v2
type v2[V any, A1 any, A2 any] struct {
	*step
	fn func(A1, A2) V
}

func V2[V any, A1 any, A2 any](wf *Workflow, action string, fn func(A1, A2) V) *v2[V, A1, A2] {
	step := wf.add(action)
	return &v2[V, A1, A2]{step, fn}
}

type ve2[V any, A1 any, A2 any] struct {
	*step
	fn func(A1, A2) (V, error)
}

func VE2[V any, A1 any, A2 any](wf *Workflow, action string, fn func(A1, A2) (V, error)) *ve2[V, A1, A2] {
	step := wf.add(action)
	return &ve2[V, A1, A2]{step, fn}
}

// GO
// v
func (s *v[V]) Go() (*V, *step) {
	var v V
	s.runAsync(func() {
		v = s.fn()
	})
	return &v, s.step
}

func (s *ve[V]) Go() (*V, *step) {
	var v V
	s.runAsync(func() {
		result, err := s.fn()
		s.processError(err)
		v = result
	})
	return &v, s.step
}

// v1
func (s *v1[V, A1]) Go(a1 A1) (*V, *step) {
	var v V
	s.runAsync(func() {
		v = s.fn(a1)
	})
	return &v, s.step
}

func (s *ve1[V, A1]) Go(a1 A1) (*V, *step) {
	var v V
	s.runAsync(func() {
		result, err := s.fn(a1)
		s.processError(err)
		v = result
	})
	return &v, s.step
}

// v2
func (s *v2[V, A1, A2]) Go(a1 A1, a2 A2) (*V, *step) {
	var v V
	s.runAsync(func() {
		v = s.fn(a1, a2)
	})
	return &v, s.step
}

func (s *ve2[V, A1, A2]) Go(a1 A1, a2 A2) (*V, *step) {
	var v V
	s.runAsync(func() {
		result, err := s.fn(a1, a2)
		s.processError(err)
		v = result
	})
	return &v, s.step
}

// DO
// v
func (s *v[V]) Do() V {
	if s.shouldRun() {
		return s.fn()
	}
	return zero[V]()
}

func (s *ve[V]) Do() (V, error) {
	if s.shouldRun() {
		result, err := s.fn()
		return result, s.processError(err)
	}
	return zero2[V, error]()
}

// v1
func (s *v1[V, A1]) Do(a1 A1) V {
	if s.shouldRun() {
		return s.fn(a1)
	}
	return zero[V]()
}

func (s *ve1[V, A1]) Do(a1 A1) (V, error) {
	if s.shouldRun() {
		result, err := s.fn(a1)
		return result, s.processError(err)
	}
	return zero2[V, error]()
}

// v2
func (s *v2[V, A1, A2]) Do(a1 A1, a2 A2) V {
	if s.shouldRun() {
		return s.fn(a1, a2)
	}
	return zero[V]()
}

func (s *ve2[V, A1, A2]) Do(a1 A1, a2 A2) (V, error) {
	if s.shouldRun() {
		result, err := s.fn(a1, a2)
		return result, s.processError(err)
	}
	return zero2[V, error]()
}
