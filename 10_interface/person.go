package person

import (
	"errors"
	"io"
	"sort"
)

type Person interface {
	Age() int
}

type Employee struct {
	age int
}

type Customer struct {
	age int
}

func (c *Customer) Age() int {
	return c.age
}

func (e *Employee) Age() int {
	return e.age
}

// OldestAge() returns the age
// of the oldest among
// employees and customers
func OldestAge(p ...Person) (int, error) {
	if p == nil {
		return 0, errors.New("no argsuments were provided")
	}
	sort.Slice(p, func(i, j int) bool { return p[i].Age() >= p[j].Age() })
	return p[0].Age(), nil
}

// OldestObject() returns the interface
// of the oldest among
// employees and customers
func OldestObject(args ...interface{}) (interface{}, error) {
	if args == nil {
		return nil, errors.New("no argsuments were provided")
	}
	max := 0
	var i interface{}
	for _, p := range args {
		switch t := p.(type) {
		case Employee:
			if t.age > max {
				max = t.age
				i = t
			}
		case Customer:
			if t.age > max {
				max = t.age
				i = t
			}
		}
	}

	return i, nil
}

// Print() passes to io.Writer
// only arguments
// with a 'string' type
func Print(w io.Writer, args ...interface{}) {
	for _, i := range args {
		if str, ok := i.(string); ok {
			w.Write([]byte(str))
		}
	}
}
