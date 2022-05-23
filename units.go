package main

import "fmt"

type Unit struct {
	V int
	U string
}

func (u Unit) String() string {
	return fmt.Sprintf("%d%s", u.V, u.U)
}
