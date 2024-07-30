package position

type Position struct {
	idx  int
	ln   int
	col  int
	fn   string
	ftxt string
}

func (p *Position) Advance(currentChar rune) *Position {
	p.idx++
	p.col++

	if currentChar == '\n' {
		p.ln++
		p.col = 0
	}

	return p
}

func (p Position) Copy() Position {
	return Position{p.idx, p.ln, p.col, p.fn, p.ftxt}
}
