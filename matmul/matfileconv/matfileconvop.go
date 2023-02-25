package matfileconv

import (
  "bufio"
  "log"
  "os"
  "strconv"
  "strings"
)

// get_num_slice_frm_string takes a string as a parameter
// and returns a slice of int containing the int value of
// the string representation.
// The function panics in case of any error during conversion.
// Probably signalling that the NaN was obtained.
func get_num_slice_frm_string(str string) []int {
  strSlc := strings.Fields(str)
  numSlc := make([]int, len(strSlc))
  for i, v := range strSlc {
    val, err := strconv.Atoi(v)
    if err != nil {
      log.Fatal(err)
    }
    numSlc[i] = val
  }
  return numSlc
}

// get_str_from_num_slice appends the int slice into
// a single string, treating the numbers as words.
// The words are space separated.
func get_str_from_num_slice(slc []int) string {
  strSlc := make([]string, len(slc))
  for i, val := range slc {
    strSlc[i] = strconv.Itoa(val)
  }
  return strings.Join(strSlc, " ")
}

// Put_mat_to_file converts the matrix given as input, to its
// corresponding string representation and dumps the string
// into a file which is also suplied as parameter.
// The file is created, even if it is present. Thus, its
// contents are truncated and replaced with the new string
// representation of the matrix.
func Put_mat_to_file(matrix [][]int, file string) {
  f, err := os.Create(file)
  if err != nil {
    log.Fatal(err)
  }
  defer f.Close()
  writer := bufio.NewWriter(f)
  for _, slc := range matrix {
    _, err = writer.WriteString(get_str_from_num_slice(slc) + "\n")
    if err != nil {
      log.Fatal(err)
    }
  }
  writer.Flush()
}

// Get_mat_from_file reads the file, given as input
// parameter and converts the contents to an integer
// matrix. The dimension is determined by the contents
// in the file.
// The expectation is string representation of a number.
// In case, there is failure due to unexpected format, the
// function panics.
func Get_mat_from_file(file string) [][]int {
  f, err := os.Open(file)
  if err != nil {
    log.Fatal(err)
  }
  defer f.Close()

  var matrix [][]int

  scanner := bufio.NewScanner(f)
  for scanner.Scan() {
    numSlc := get_num_slice_frm_string(scanner.Text())
    matrix = append(matrix, numSlc)
  }
  if err = scanner.Err(); err != nil {
    log.Fatal(err)
  }
  return matrix
}
