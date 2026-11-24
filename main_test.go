package main

import (
	"reflect"
	"testing"
)

func TestMaxXorQueries(t *testing.T) {
	got := MaxXorQueries([]int{0, 1, 2, 3, 4}, []Query{{3, 1}, {1, 3}, {5, 6}})
	want := []int{3, 3, 7}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("MaxXorQueries = %v, want %v", got, want)
	}

	got = MaxXorQueries([]int{5, 2, 4, 6, 6, 3}, []Query{{12, 4}, {8, 1}, {6, 3}})
	want = []int{15, -1, 5}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("MaxXorQueries = %v, want %v", got, want)
	}

	// Bound below every element: all -1.
	got = MaxXorQueries([]int{10, 20}, []Query{{1, 5}, {99, 9}})
	if !reflect.DeepEqual(got, []int{-1, -1}) {
		t.Errorf("below-all bounds = %v", got)
	}

	// Answers must return in the original query order despite sorting.
	got = MaxXorQueries([]int{1, 2, 3}, []Query{{1, 3}, {1, 1}})
	if !reflect.DeepEqual(got, []int{3, 0}) {
		t.Errorf("order preservation = %v", got)
	}
}
