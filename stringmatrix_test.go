package stringmatrix

import (
       "testing"
)

func check_array(actual []string, expected []string, t *testing.T) {
  if len(actual) != len(expected) {
    t.Errorf("Actuals - num entries: %v Expected %v", len(actual), len(expected))
  }

  for i := range actual {
    if actual[i] != expected[i] {
      t.Errorf("Actual[%v] = %v expected %v", i, actual[i], expected[i])
    }
  }
}

  

func TestStringMatrix(t *testing.T) {
  m := NewStringMatrix()
  if m.NumRows() != 0 {
    t.Errorf("m.NumRows() = %v, want 0", m.NumRows())
  }  
  if m.NumCols() != 0 {
    t.Errorf("m.NumCols() = %v, want 0", m.NumCols())
  }
 
  r1 := make([]string, 2)
  r1[0] = "foo" 
  r1[1] = "bar" 

  m.AppendRow(r1)  


  if m.NumRows() != 1 {
    t.Errorf("m.NumRows() = %v, want 1", m.NumRows())
  }

  r2 := make([]string, 2)
  r2[0] = "bar"
  r2[1] = "baz" 

  m.AppendRow(r2)  

  check_array(r1, m.Row(0), t)
  check_array(r2, m.Row(1), t)

  m2 := m.Transpose()
  check_array(m2.Column(0), m.Row(0), t)
  check_array(m2.Column(1), m.Row(1), t)
}
