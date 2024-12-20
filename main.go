package main

import "fmt"

const prefix = "Hello, "

func Hello(n string) string {
  if n == "" {
    n = "unknown"
  }
  return prefix + n
}
func main(){
  fmt.Println(Hello("Quint"))
}
