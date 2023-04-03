package lexer

import "cs4215/goophy/pkg/token"

type Lexer struct {
	input   string
	pos     int  // Current position in input (points to current char)
	readPos int  // Current read pos in input(after current char)
	ch      byte // Current char under examination
}

func NewLexer(input string) *Lexer {
	l := &Lexer{
		input: input,
	}
	l.readChar()
	return l
}

// Lexer only supports acii instead o fthe full unicode range.
func (l *Lexer) readChar() {
	if l.readPos >= len(l.input) {
		l.ch = 0 // ascii code for null char, set it to null when readPos exceeds string.
	} else {
		l.ch = l.input[l.readPos] // if not we set the curr char to the char of the input.
	}
	l.pos = l.readPos
	l.readPos++
}

func (l *Lexer) NextToken() token.Token {
	l.eatWhitespace()
	var tok token.Token
	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			l.readChar()
			tok = token.NewTokenWithStr(token.EQ, "==")
		} else {
			tok = token.NewToken(token.ASSIGN, l.ch)
		}
	case ';':
		tok = token.NewToken(token.SEMICOLON, l.ch)
	case '-':
		tok = token.NewToken(token.MINUS, l.ch)
	case '!':
		if l.peekChar() == '=' {
			l.readChar()
			tok = token.NewTokenWithStr(token.NOT_EQ, "!=")
		} else {
			tok = token.NewToken(token.BANG, l.ch)
		}
	case '/':
		tok = token.NewToken(token.SLASH, l.ch)
	case '*':
		tok = token.NewToken(token.ASTERISK, l.ch)
	case '<':
		tok = token.NewToken(token.LT, l.ch)
	case '>':
		tok = token.NewToken(token.GT, l.ch)
	case '(':
		tok = token.NewToken(token.LPAREN, l.ch)
	case ')':
		tok = token.NewToken(token.RPAREN, l.ch)
	case ',':
		tok = token.NewToken(token.COMMA, l.ch)
	case '+':
		tok = token.NewToken(token.PLUS, l.ch)
	case '{':
		tok = token.NewToken(token.LBRACE, l.ch)
	case '}':
		tok = token.NewToken(token.RBRACE, l.ch)
	case 0:
		tok = token.NewTokenWithStr(token.EOF, "")
	default:

		if isLetter(l.ch) {
			// Read identifier returns a string
			literal := l.readIdentifier()
			return token.NewTokenWithStr(token.LookUpIdentifier(literal), literal)
		} else if isDigit(l.ch) {
			return token.NewTokenWithStr(token.INT, l.readNumber())
		} else {
			// decalre token as ileegal if cannot be identified properly/
			tok = token.NewToken(token.ILLEGAL, l.ch)
		}
	}
	l.readChar()
	return tok
}

// This function returns a string for a given identifier, it keeps increasing the read position
// while the char at the current pos is a letter or an underscore.
func (l *Lexer) readIdentifier() string {
	position := l.pos
	for isLetterOrUnderscore(l.ch) {
		l.readChar()
	}
	return l.input[position:l.pos]
}

func (l *Lexer) eatWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) peekChar() byte {
	if l.readPos >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPos]
	}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z'
}

func isLetterOrUnderscore(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_' || isDigit(ch)
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) readNumber() string {
	pos := l.pos
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[pos:l.pos]
}
