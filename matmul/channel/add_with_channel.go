package channel

import (
  "log"
)

// Add_with_channel performs addition of matrix rows
// provided there is a one-to-one dimension match between
// them.
// It spins up as many goroutines as the number of rows
// in the matrix.
// To sync up the goroutines, channels are used primarily.
// The done channel signals the completion of all goroutines
// that are currently in execution, to the main.
// This prevents the main goroutine to prematurely terminate
// before the addition of the matrix is complete.
func Add_with_channel(mat1, mat2 [][]int) [][]int {
  if len(mat1) != len(mat2) {
    log.Fatal("Dimension mismatch")
  }

  resMat := make([][]int, len(mat1))
  done := make(chan bool, len(resMat))

  for idx := 0; idx < len(mat1); idx++ {

    resMat[idx] = make([]int, len(mat1[idx]))

    ch1 := make(chan []int, 1)
    ch2 := make(chan []int, 1)

    ch1 <- mat1[idx]
    ch2 <- mat2[idx]

    go func(ch1, ch2 <-chan []int, done chan<- bool, mat []int) {
      slc1 := <-ch1
      slc2 := <-ch2

      if len(slc1) != len(slc2) {
        log.Fatal("Dimension mismatch at row ", idx)
      }

      for i := 0; i < len(slc1); i++ {
        mat[i] = slc1[i] + slc2[i]
      }
      done <- true
    }(ch1, ch2, done, resMat[idx])
  }

  for k := 0; k < cap(done); k++ {
    if ! <-done{
      log.Fatal("Matrix addition failure")
    }
  }

  close (done)

  return resMat
}
