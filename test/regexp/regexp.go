package main

import (
	"fmt"
	"regexp"
)

type S struct {
	Name string
}

func main() {
	s1 := S{Name: "lisi"}
	s2 := s1
	fmt.Println(s1, s2)
	s1.Name = "wangwu"
	fmt.Println(s1, s2)

	fmt.Println(regexp.MatchString("^/auth/*", "/api/v1/auth/login"))
}
