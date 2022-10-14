package roundrobin

import (
	"errors"
	"sync/atomic"
)

var ErrEmptyValues = errors.New("no values were provided")

// RoundRobin is an interface for representing round-robin balancing.
type RoundRobin interface {
	Next() string
}

type roundrobin struct {
	values []string
	next   uint32
}

// New returns RoundRobin implementation(*roundrobin).
func New(values ...string) (RoundRobin, error) {
	if len(values) == 0 {
		return nil, ErrEmptyValues
	}

	return &roundrobin{
		values: values,
	}, nil
}

// Next returns next value
func (r *roundrobin) Next() string {
	n := atomic.AddUint32(&r.next, 1)
	return r.values[(int(n)-1)%len(r.values)]
}
