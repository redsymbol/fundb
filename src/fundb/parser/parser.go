package parser

const whitespace = " \t\n"

type Scanner struct {
	src string
	pos int
}

func (s* Scanner) Init(src string) {
	s.src = src
	s.pos = 0
}

func (s* Scanner) Scan() (lit string) {
}

func ParseExpr(toparse string) {
}