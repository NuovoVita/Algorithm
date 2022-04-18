package sorting

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"
	"time"
)

func TestBubbleSort(t *testing.T) {
	var arr = sort.IntSlice{60, 50, 40, 30, 20, 10}
	var ans = sort.IntSlice{10, 20, 30, 40, 50, 60}
	BubbleSort(arr)
	if !reflect.DeepEqual(arr, ans) {
		t.Errorf("expected be %v, but %d got", ans, arr)
	}
}

func TestBubbleSort2(t *testing.T) {
	var arr = make(sort.IntSlice, 10, 100)
	var ans = make(sort.IntSlice, 10, 100)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(100)
	}

	copy(ans, arr)
	BubbleSort2(ans)
	arr.Sort()

	if !reflect.DeepEqual(arr, ans) {
		t.Errorf("expected be %v, but %d got", ans, arr)
	}
}

func TestBubbleSort3(t *testing.T) {
	var arr = make(sort.IntSlice, 10, 100)
	var ans = make(sort.IntSlice, 10, 100)

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(100)
	}

	copy(ans, arr)
	BubbleSort3(ans)
	arr.Sort()

	if !reflect.DeepEqual(arr, ans) {
		t.Errorf("expected be %v, but %d got", ans, arr)
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	var arr = make(sort.IntSlice, 100, 100)
	var ans = make(sort.IntSlice, 100, 100)

	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(100)
	}

	copy(ans, arr)
	ans.Sort()

	for n := 0; n < b.N; n++ {
		BubbleSort(arr)
		if !reflect.DeepEqual(arr, ans) {
			b.Errorf("expected be %v, but %d got", ans, arr)
		}
	}
}

func BenchmarkBubbleSort2(b *testing.B) {
	var arr = make(sort.IntSlice, 100, 100)
	var ans = make(sort.IntSlice, 100, 100)

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(100)
	}

	copy(ans, arr)
	arr.Sort()

	for n := 0; n < b.N; n++ {
		BubbleSort2(ans)
		if !reflect.DeepEqual(arr, ans) {
			b.Errorf("expected be %v, but %d got", ans, arr)
		}
	}
}

func BenchmarkBubbleSort3(b *testing.B) {
	var arr = make(sort.IntSlice, 100, 100)
	var ans = make(sort.IntSlice, 100, 100)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(100)
	}

	copy(ans, arr)
	arr.Sort()

	for n := 0; n < b.N; n++ {
		BubbleSort3(ans)
		if !reflect.DeepEqual(arr, ans) {
			b.Errorf("expected be %v, but %d got", ans, arr)
		}
	}
}
