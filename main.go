package main

import (
	"fmt"
	"os"
	"unicode"
)

// ==================== TOKEN ====================
type TokenKind int

const (
	TK_LBRACE TokenKind = iota
	TK_RBRACE
	TK_COLON
	TK_STRING
	TK_EOF
)

type Token struct {
	Kind  TokenKind
	Value string
}

// ==================== LEXER ====================
type Lexer struct {
	src    string
	pos    int
	tokens []Token
}

func NewLexer(src string) *Lexer {
	l := &Lexer{src: src}
	l.run()
	return l
}

func (l *Lexer) run() {
	for l.pos < len(l.src) {
		ch := l.peek()

		if unicode.IsSpace(ch) {
			l.pos++
			continue
		}

		switch ch {
		case '{':
			l.emit(TK_LBRACE)
		case '}':
			l.emit(TK_RBRACE)
		case ':':
			l.emit(TK_COLON)
		case '"':
			if str, ok := l.readString(); ok {
				l.tokens = append(l.tokens, Token{Kind: TK_STRING, Value: str})
			} else {
				os.Exit(1)
			}
		default:
			os.Exit(1)
		}
	}
	l.emit(TK_EOF)
}

func (l *Lexer) peek() rune {
	if l.pos >= len(l.src) {
		return 0
	}
	return rune(l.src[l.pos])
}

func (l *Lexer) emit(kind TokenKind) {
	l.tokens = append(l.tokens, Token{Kind: kind})
	l.pos++
}

func (l *Lexer) readString() (string, bool) {
	if l.peek() != '"' {
		return "", false
	}
	l.pos++ // skip "

	start := l.pos
	for l.pos < len(l.src) && l.src[l.pos] != '"' {
		l.pos++
	}
	if l.pos >= len(l.src) {
		return "", false
	}
	value := l.src[start:l.pos]
	l.pos++ // skip "
	return value, true
}

func (l *Lexer) Next() Token {
	if len(l.tokens) == 0 {
		return Token{Kind: TK_EOF}
	}
	t := l.tokens[0]
	l.tokens = l.tokens[1:]
	return t
}

func (l *Lexer) Peek() Token {
	if len(l.tokens) == 0 {
		return Token{Kind: TK_EOF}
	}
	return l.tokens[0]
}

// ==================== PARSER ====================
type Parser struct {
	lexer *Lexer
}

func NewParser(l *Lexer) *Parser { return &Parser{lexer: l} }

func (p *Parser) Parse() bool {
	if !p.eat(TK_LBRACE) {
		return false
	}

	if p.tryEat(TK_RBRACE) {
		return p.eof()
	}

	if !p.eat(TK_STRING) {
		return false
	}
	if !p.eat(TK_COLON) {
		return false
	}
	if !p.eat(TK_STRING) {
		return false
	}
	if !p.eat(TK_RBRACE) {
		return false
	}

	return p.eof()
}

func (p *Parser) eat(want TokenKind) bool {
	t := p.lexer.Next()
	return t.Kind == want
}

func (p *Parser) tryEat(want TokenKind) bool {
	if p.lexer.Peek().Kind == want {
		p.lexer.Next()
		return true
	}
	return false
}

func (p *Parser) eof() bool {
	return p.lexer.Peek().Kind == TK_EOF
}

// ==================== MAIN ====================
func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <file>\n", os.Args[0])
		os.Exit(1)
	}
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		os.Exit(1)
	}

	lexer := NewLexer(string(data))
	parser := NewParser(lexer)

	if parser.Parse() {
		os.Exit(0)
	}
	os.Exit(1)
}
