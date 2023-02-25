package waitgrp

import (
  "fmt"
  "log"
  "sync"
)

// Add_with_WaitGroup performs addition of matrix rows
// provided there is a one-to-one dimension match between
// them.
// It spins up as many goroutines as the number of rows
// in the matrix.
// To sync up the goroutines, WaitGroup is used to keep
// a track of the number and status of goroutines.
func Add_with_WaitGroup(mat1, mat2 [][]int) [][]int {

  resMat := make([][]int, len(mat1))
  var wg sync.WaitGroup

  for i := 0; i < len(resMat); i++ {
    wg.Add(1)

    resMat[i] = make([]int, len(mat1[i]))

    go func(slc1, slc2, resSlc []int) {
      defer wg.Done()
      fmt.Println(slc1, slc2)
      if len(slc1) != len(slc2) {
        log.Fatal("Dimension mismatch at row ", i)
      }
      for k := 0; k < len(slc1); k++ {
        resSlc[k] = slc1[k] + slc2[k]
      }
    }(mat1[i], mat2[i], resMat[i])
  }

  wg.Wait()

  return resMat
}
