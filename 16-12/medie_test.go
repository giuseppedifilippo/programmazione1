package main
import "testing"
func TestAvg(t *testing.T) {
  v := avg([]int{30,30,30})
  if v != 30.0 {
    t.Error()
  }
}