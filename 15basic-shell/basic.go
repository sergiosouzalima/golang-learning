package basic

import (
	"fmt"
	"strconv"
	"unicode"
)

const DIGITS = "0123456789"

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

func (p *Position) Copy() Position {
	return Position{
		idx:  p.idx,
		ln:   p.ln,
		col:  p.col,
		fn:   p.fn,
		ftxt: p.ftxt,
	}
}

type Error struct {
	posStart  Position
	posEnd    Position
	errorName string
	details   string
}

func (e *Error) Error() string {
	result := fmt.Sprintf("%s: %s\n", e.errorName, e.details)
	result += fmt.Sprintf("File %s, line %d", e.posStart.fn, e.posStart.ln+1)
	return result
}

type IllegalCharError struct {
	Error
}

func NewIllegalCharError(posStart, posEnd Position, details string) *IllegalCharError {
	return &IllegalCharError{
		Error{
			posStart:  posStart,
			posEnd:    posEnd,
			errorName: "Illegal Character",
			details:   details,
		},
	}
}

const (
	TT_INT    = "INT"
	TT_FLOAT  = "FLOAT"
	TT_PLUS   = "PLUS"
	TT_MINUS  = "MINUS"
	TT_MUL    = "MUL"
	TT_DIV    = "DIV"
	TT_LPAREN = "LPAREN"
	TT_RPAREN = "RPAREN"
)

type Token struct {
	Type  string
	Value interface{}
}

func (t *Token) String() string {
	if t.Value != nil {
		return fmt.Sprintf("%s:%v", t.Type, t.Value)
	}
	return t.Type
}

type Lexer struct {
	fn          string
	text        string
	pos         Position
	currentChar rune
}

func NewLexer(fn, text string) *Lexer {
	l := &Lexer{
		fn:          fn,
		text:        text,
		pos:         Position{-1, 0, -1, fn, text},
		currentChar: rune(0),
	}
	l.Advance()
	return l
}

func (l *Lexer) Advance() {
	l.pos = *l.pos.Advance(l.currentChar)
	if l.pos.idx < len(l.text) {
		l.currentChar = rune(l.text[l.pos.idx])
	} else {
		l.currentChar = rune(0)
	}
}

func (l *Lexer) MakeTokens() ([]Token, error) {
	var tokens []Token

	for l.currentChar != rune(0) {
		if unicode.IsSpace(l.currentChar) {
			l.Advance()
		} else if unicode.IsDigit(l.currentChar) || l.currentChar == '.' {
			tokens = append(tokens, l.MakeNumber())
		} else {
			switch l.currentChar {
			case '+':
				tokens = append(tokens, Token{Type: TT_PLUS})
				l.Advance()
			case '-':
				tokens = append(tokens, Token{Type: TT_MINUS})
				l.Advance()
			case '*':
				tokens = append(tokens, Token{Type: TT_MUL})
				l.Advance()
			case '/':
				tokens = append(tokens, Token{Type: TT_DIV})
				l.Advance()
			case '(':
				tokens = append(tokens, Token{Type: TT_LPAREN})
				l.Advance()
			case ')':
				tokens = append(tokens, Token{Type: TT_RPAREN})
				l.Advance()
			default:
				posStart := l.pos.Copy()
				char := l.currentChar
				l.Advance()
				return nil, NewIllegalCharError(posStart, l.pos, string(char))
			}
		}
	}

	return tokens, nil
}

func (l *Lexer) MakeNumber() Token {
	numStr := ""
	dotCount := 0

	for l.currentChar != rune(0) && (unicode.IsDigit(l.currentChar) || l.currentChar == '.') {
		if l.currentChar == '.' {
			if dotCount == 1 {
				break
			}
			dotCount++
		}
		numStr += string(l.currentChar)
		l.Advance()
	}

	if dotCount == 0 {
		i, _ := strconv.Atoi(numStr)
		return Token{Type: TT_INT, Value: i}
	}
	f, _ := strconv.ParseFloat(numStr, 64)
	return Token{Type: TT_FLOAT, Value: f}
}

func Run(fn, text string) ([]Token, error) {
	lexer := NewLexer(fn, text)
	tokens, err := lexer.MakeTokens()
	return tokens, err
}
