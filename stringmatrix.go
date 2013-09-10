package stringmatrix

import (
  "fmt"
)

type StringMatrix struct {
  Fields [][]string
  RowCount int
  ColCount int
}

func NewStringMatrix() StringMatrix {
  m := StringMatrix{}
  m.RowCount = 0
  m.ColCount = 0

  return m
}

func NewStringMatrixWithSize(h int, w int) StringMatrix {
  m := StringMatrix{}
  m.RowCount = h
  m.ColCount = w
  
  return m
}

func (m StringMatrix) NumRows() int {
  return m.RowCount
}

func (m StringMatrix) NumCols() int {
  return m.ColCount
}

func (m StringMatrix) Column(idx int) []string {
  if idx >= m.NumCols() {
    panic(fmt.Sprintf("Cannot extract column index %d from matrix with %d columns\n", idx, m.NumCols()))
  }

  out := make([]string, m.NumRows())
    
  for i:=0; i<m.NumRows(); i++ {
    out[i] = m.Fields[i][idx]
  }
  return out
}

func (m StringMatrix) Row(idx int) []string {
  if idx >= m.NumRows() {
    panic(fmt.Sprintf("Cannot extract row index %d from matrix with %d columns\n", idx, m.NumRows()))
  }

  return m.Fields[idx]
}

func (m *StringMatrix) AppendColumn(col []string) {
  for i:= range col {
    m.Fields[i] = append(m.Fields[i], col[i])
  }

  // fill out sparse matrix
  if len(col) != m.RowCount {
    for i := len(col); i<m.RowCount; i++ {
      m.Fields[i] = append(m.Fields[i], "")
    }
  }
  m.ColCount = m.ColCount + 1
}

func (m *StringMatrix) AppendRow(row []string) {
  m.Fields = append(m.Fields, row)
  m.RowCount = m.RowCount + 1
  if m.ColCount != 0  &&  m.ColCount != len(row) {
    panic(fmt.Sprintf("AppendRow with inconsistent # of columns %v expected %v",
                    len(row), m.ColCount))
  }
  m.ColCount = len(row)
}

func (m *StringMatrix) AppendEmptyRow() {
  m.Fields = append(m.Fields, make([]string, m.NumCols()))
  m.RowCount = m.RowCount + 1
}

func (m *StringMatrix) ReplaceLastColumn(col []string, startRow int, endRow int) {
  m.ReplaceColumn(col, m.ColCount-1, startRow, endRow)
}


func (m *StringMatrix) ReplaceColumn(col []string, colIdx int, startRow int, endRow int) {
  if len(col) != (endRow - startRow + 1) {
    panic(fmt.Sprintf("Inconsistent arguments:  len(col)=%v startRow %v endRow %v", len(col), startRow, endRow))
  }
  for i := startRow; i<= endRow; i++ {
    m.Fields[i][colIdx] = col[i - startRow]
  }
}

func (m *StringMatrix) ReplaceLastRow(row []string, startCol int, endCol int) {
  m.ReplaceRow(row, m.RowCount-1, startCol, endCol)
}

func (m *StringMatrix) ReplaceRow(row []string, rowIdx int, startCol int, endCol int) {
  if len(row) != (endCol - startCol + 1) {
    panic(fmt.Sprintf("Inconsistent arguments:  len(row)=%v startCol %v endCol %v", len(row), startCol, endCol))
  }
  for idx := startCol; idx<= endCol; idx++ {
    m.Fields[rowIdx][idx] = row[idx - startCol]
  }
}

func (m *StringMatrix) AppendToColumn(s string, colIdx int) {
  c := m.Column(colIdx)

  if len(c) == m.NumRows() {
    m.AppendEmptyRow()
  }

  m.Fields[m.NumRows()-1][colIdx] = s
}

func (m *StringMatrix) ReplaceArrayInColumn(col []string, colIdx int) {
  m.ReplaceColumn(col, colIdx, m.NumRows() - len(col), m.NumRows() - 1)
}

func (m *StringMatrix) AppendArrayToColumn(col []string, colIdx int) {

  newRows := len(col)
  curRows := 0
  if m.NumRows() > 0 {
    c := m.Column(colIdx)

    newRows = len(c) + len(col) - m.NumRows()
    curRows = m.NumRows()
  }

  for i := 0; i<newRows; i++ {
    m.AppendEmptyRow()
  }

  m.ReplaceColumn(col, colIdx, curRows, curRows + newRows - 1)
 
}

func (m *StringMatrix) Transpose() StringMatrix {
  out := NewStringMatrix()
  
  for i:=0; i<m.NumCols(); i++ {
    c := m.Column(i)
    out.AppendRow(c)
  }
  return out 
}
