package lexer

import (
	"monkey/token"
)

type Lexer struct {
	position     int
	readPosition int
	ch           rune
	characters   []rune
}

func New(input string) *Lexer {
	l := &Lexer{characters: []rune(input)}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.characters) {
		l.ch = rune(0)
	} else {
		l.ch = l.characters[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	l.skipWhitespace()
	switch l.ch {
	case rune('='):
		tok = newToken(token.ASSIGN, l.ch)
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case rune(';'):
		tok = newToken(token.SEMICOLON, l.ch)
	case rune('('):
		tok = newToken(token.LPAREN, l.ch)
	case rune(')'):
		tok = newToken(token.RPAREN, l.ch)
	case rune(','):
		tok = newToken(token.COMMA, l.ch)
	case rune('+'):
		tok = newToken(token.PLUS, l.ch)
	case rune('{'):
		tok = newToken(token.LBRACE, l.ch)
	case rune('}'):
		tok = newToken(token.RBRACE, l.ch)
	case rune('-'):
		tok = newToken(token.MINUS, l.ch)
	case rune('/'):
		tok = newToken(token.SLASH, l.ch)
	case rune('*'):
		tok = newToken(token.ASTERISK, l.ch)
	case rune('<'):
		tok = newToken(token.LT, l.ch)
	case rune('>'):
		tok = newToken(token.GT, l.ch)
	case rune('!'):
		if l.peekChar() == rune('=') {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.NOT_EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case rune('"'):
		tok.Type = token.STRING
		tok.Literal = l.readString()
	case rune('['):
		tok = newToken(token.LBRACKET, l.ch)
	case rune(']'):
		tok = newToken(token.RBRACKET, l.ch)
	case rune(':'):
		tok = newToken(token.COLON, l.ch)
	case rune('.'):
		l.readChar()
		tok.Type = token.FLOAT
		tok.Literal = "." + l.readNumber()
	case rune(0):
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isDigit(l.ch) {
			return l.readDecimal()
		} else {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		}
	}
	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch rune) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isIdentifier(l.ch) {
		l.readChar()
	}
	return string(l.characters[position:l.position])
}

func isIdentifier(ch rune) bool {
	return !isDigit(ch) && !isWhitespace(ch) && !isBrace(ch) && !isOperator(ch) && !isComparison(ch) && !isCompound(ch) && !isBrace(ch) && !isParen(ch) && !isBracket(ch) && !isEmpty(ch)
}

func isWhitespace(ch rune) bool {
	return ch == rune(' ') || ch == rune('\t') || ch == rune('\n') || ch == rune('\r')
}
func isOperator(ch rune) bool {
	return ch == rune('+') || ch == rune('-') || ch == rune('/') || ch == rune('*')
}
func isComparison(ch rune) bool {
	return ch == rune('=') || ch == rune('!') || ch == rune('>') || ch == rune('<')
}
func isCompound(ch rune) bool {
	return ch == rune('.') || ch == rune(',') || ch == rune(':') || ch == rune('"') || ch == rune(';')
}
func isBrace(ch rune) bool {
	return ch == rune('{') || ch == rune('}')
}
func isBracket(ch rune) bool {
	return ch == rune('[') || ch == rune(']')
}
func isParen(ch rune) bool {
	return ch == rune('(') || ch == rune(')')
}
func isLetter(ch rune) bool {
	return rune('a') <= ch && ch <= rune('z') || rune('A') <= ch && ch <= rune('Z') || ch == rune('_')
}
func isEmpty(ch rune) bool {
	return rune(0) == ch
}

func (l *Lexer) skipWhitespace() {
	for l.ch == rune(' ') || l.ch == rune('\t') || l.ch == rune('\n') || l.ch == rune('\r') {
		l.readChar()
	}
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return string(l.characters[position:l.position])
}

func (l *Lexer) readDecimal() token.Token {
	integer := l.readNumber()
	if l.ch != '.' {
		return token.Token{Type: token.INT, Literal: integer}
	} else {
		l.readChar()
		fraction := l.readNumber()
		return token.Token{Type: token.FLOAT, Literal: integer + "." + fraction}
	}
}

func (l *Lexer) readString() string {
	position := l.position + 1
	for {
		l.readChar()
		if l.ch == '"' {
			break
		}
	}
	return string(l.characters[position:l.position])
}

func isDigit(ch rune) bool {
	return rune('0') <= ch && ch <= rune('9')
}

func (l *Lexer) peekChar() rune {
	if l.readPosition >= len(l.characters) {
		return rune(0)
	} else {
		return l.characters[l.readPosition]
	}
}
