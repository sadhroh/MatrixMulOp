package main

import (
  "fmt"
  "MatrixMulOp/matmul/matfileconv"
  // "MatrixMulOp/matmul/waitgrp"
  "MatrixMulOp/matmul/channel"
  "os"
)

func main() {
  if len(os.Args) < 4 {
    fmt.Println("Invalid number of arguments")
    fmt.Println("Enter file1 & file2 containing matrices")
    fmt.Println("file3 would be the output when the result is dumped")
    return
  }

  matrix1 := matfileconv.Get_mat_from_file(os.Args[1])
  matrix2 := matfileconv.Get_mat_from_file(os.Args[2])

  // resMatWaitGrp := waitgrp.Add_with_WaitGroup(matrix1, matrix2)
  resMatChannel := channel.Add_with_channel(matrix1, matrix2)

  matfileconv.Put_mat_to_file(resMatChannel, os.Args[3])

  fmt.Println("Result", resMatChannel)

}
