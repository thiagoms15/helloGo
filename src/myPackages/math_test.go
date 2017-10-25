package math

import "testing"

func TestAverage(t *testing.T) {
  var v float64
  v = Average([]float64{1,2})
  if v != 1.5 {
    t.Error("Expected 1.5, got ", v)
  }

  v = Average([]float64{-10,3})
  if v != -3.5 {
    t.Error("Expected -3.5, got ", v)
  }
}