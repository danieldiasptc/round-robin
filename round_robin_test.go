package roundrobin

import (
	"fmt"
	"reflect"
	"sync"
	"testing"
)

func TestRoundRobin(t *testing.T) {
	tests := []struct {
		values   []string
		iserr    bool
		expected []string
		want     []string
	}{
		{
			values: []string{
				"192.168.33.10",
				"192.168.33.11",
				"192.168.33.12",
			},
			iserr: false,
			want: []string{
				"192.168.33.10",
				"192.168.33.11",
				"192.168.33.12",
				"192.168.33.10",
			},
		},
		{
			values: []string{},
			iserr:  true,
			want:   []string{},
		},
	}

	for i, test := range tests {
		rr, err := New(test.values...)

		if got, want := !(err == nil), test.iserr; got != want {
			t.Errorf("tests[%d] - RoundRobin iserr is wrong. want: %v, but got: %v", i, test.want, got)
		}

		gots := make([]string, 0, len(test.want))
		for j := 0; j < len(test.want); j++ {
			gots = append(gots, rr.Next())
		}

		if got, want := gots, test.want; !reflect.DeepEqual(got, want) {
			t.Errorf("tests[%d] - RoundRobin is wrong. want: %v, got: %v", i, want, got)
		}
	}
}

func BenchmarkRoundRobinSync(b *testing.B) {
	resources := []string{
		"127.0.0.1",
		"127.0.0.2",
		"127.0.0.3",
		"127.0.0.4",
		"127.0.0.5",
		"127.0.0.6",
		"127.0.0.7",
		"127.0.0.8",
		"127.0.0.9",
		"127.0.0.10",
	}

	for i := 1; i < len(resources)+1; i++ {
		b.Run(fmt.Sprintf("RoundRobinSliceOfSize(%d)", i), func(b *testing.B) {
			rr, err := New(resources[:i]...)
			if err != nil {
				b.Fatal(err)
			}
			// Adding WaitGroup complexity as this helps in comparing Sync and Async RoundRobinAccess (see BenchmarkRoundRobinASync as well)
			wg := &sync.WaitGroup{}
			for i := 0; i < b.N; i++ {
				wg.Add(1)
				defer wg.Done()
				rr.Next()
			}
		})
	}
}

func BenchmarkRoundRobinASync(b *testing.B) {
	resources := []string{
		"127.0.0.1",
		"127.0.0.2",
		"127.0.0.3",
		"127.0.0.4",
		"127.0.0.5",
		"127.0.0.6",
		"127.0.0.7",
		"127.0.0.8",
		"127.0.0.9",
		"127.0.0.10",
	}

	for i := 1; i < len(resources)+1; i++ {
		b.Run(fmt.Sprintf("RoundRobinSliceOfSize(%d)", i), func(b *testing.B) {
			rr, err := New(resources[:i]...)
			if err != nil {
				b.Fatal(err)
			}
			wg := &sync.WaitGroup{}
			for i := 0; i < b.N; i++ {
				wg.Add(1)
				go func() {
					defer wg.Done()
					rr.Next()
				}()
			}
			wg.Wait()
		})
	}
}
