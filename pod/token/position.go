package token

import "fmt"

// Pos describes an arbitrary source position
// including the file, line, and offset location.
// A Position is valid if the line number is > 0.
type Pos struct {
	Filename string // filename, if any
	Line     int    // line number, starting at 1
	Offset   int    // offset, starting at 0
}

// IsValid returns true if the position is valid.
func (p *Pos) IsValid() bool { return p.Line > 0 }

// String returns a string in one of several forms:
//
//	file:line:offset    valid position with file name
//	line:offset         valid position without file name
//	file                invalid position with file name
//	-                   invalid position without file name
func (p Pos) String() string {
	s := p.Filename
	if p.IsValid() {
		if s != "" {
			s += ":"
		}
		s += fmt.Sprintf("%d:%d", p.Line, p.Offset)
	}
	if s == "" {
		s = "-"
	}
	return s
}

// Before reports whether the position p is before u.
func (p Pos) Before(u Pos) bool {
	if u.Line == p.Line {
		return u.Offset > p.Offset
	}
	return u.Line > p.Line
}

// After reports whether the position p is after u.
func (p Pos) After(u Pos) bool {
	if u.Line == p.Line {
		return u.Offset < p.Offset
	}
	return u.Line < p.Line
}
