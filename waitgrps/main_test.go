package main

import "testing"

func TestAdd(t *testing.T) {

	got := Add(3, 5)
	want := 8
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}

}
func TestMain(t *testing.T) {
	main()
}

func BenchmarkAdd(b *testing.B){
	for i := 0; i < b.N; i++ {
		Add(4,3)
		
	}
}