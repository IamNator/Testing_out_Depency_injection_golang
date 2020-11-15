package main

import (
	"log"
	"os"
)

func main() {

	x := 23
	xr := log.New(os.Stdout, "Dependency injection", log.LstdFlags)
	lx := prinn(xr)
	str := lx.evalInt(x)
	log.Println(str)

}

type prin struct {
	l *log.Logger
}

func prinn(s *log.Logger) *prin {
	return &prin{s}
}

func (u prin) evalInt(n int) string {

	if n > 10 {
		u.l.Println("x is greater than 10")
		return "ok"
	}

	u.l.Println("x is less than 10")
	return "ok"
}
