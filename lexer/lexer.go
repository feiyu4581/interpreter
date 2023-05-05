package lexer

import (
	"fmt"
	"interpreter/pkg/utils"
	"interpreter/token"
	"io"
)

type LexerI interface {
	NextToken() token.Token
}

type Lexer struct {
	reader      io.Reader
	currentChar byte
	nextChar    byte
	buffer      []byte
	isEnd       bool
}

func NewLexer(reader io.Reader) LexerI {
	lexer := &Lexer{
		reader: reader,
		buffer: make([]byte, 1),
	}

	// 读取第一个 nextChar
	lexer.readChar()
	// 读取第一个 currentChar
	lexer.readChar()

	return lexer
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func newTokenWithString(tokenType token.TokenType, literal string) token.Token {
	return token.Token{Type: tokenType, Literal: literal}
}

func (lexer *Lexer) getNextChar() byte {
	if _, err := lexer.reader.Read(lexer.buffer); err != nil {
		if err == io.EOF {
			lexer.isEnd = true
			return 0
		}
		panic(fmt.Sprintf("get next character error: %s", err.Error()))
	}

	return lexer.buffer[0]
}

func (lexer *Lexer) readChar() {
	if lexer.isEnd {
		lexer.currentChar, lexer.nextChar = 0, 0
		return
	}

	ch := lexer.getNextChar()
	lexer.currentChar = lexer.nextChar
	lexer.nextChar = ch
}

func (lexer *Lexer) readIdent(checkIdent func(byte) bool) string {
	if !checkIdent(lexer.currentChar) {
		return ""
	}

	current := []byte{lexer.currentChar}
	for checkIdent(lexer.nextChar) {
		lexer.readChar()
		current = append(current, lexer.currentChar)
	}

	return string(current)
}

func (lexer *Lexer) readNumber() string {
	return lexer.readIdent(utils.IsDigit)
}

func (lexer *Lexer) readIdentifier() string {
	return lexer.readIdent(utils.IsLetter)
}

func (lexer *Lexer) skipWhiteSpace() {
	for utils.IsWriteSpace(lexer.currentChar) {
		lexer.readChar()
	}
}

func (lexer *Lexer) NextToken() token.Token {
	var tok token.Token

	// 跳过所有没有意义的空格
	lexer.skipWhiteSpace()

	switch lexer.currentChar {
	case '=':
		if lexer.nextChar == '=' {
			lexer.readChar()
			tok = newTokenWithString(token.EQ, "==")
		} else {
			tok = newToken(token.ASSIGN, lexer.currentChar)
		}
	case ';':
		tok = newToken(token.SEMICOLON, lexer.currentChar)
	case '(':
		tok = newToken(token.LPAREN, lexer.currentChar)
	case ')':
		tok = newToken(token.RPAREN, lexer.currentChar)
	case ',':
		tok = newToken(token.COMMA, lexer.currentChar)
	case '+':
		tok = newToken(token.PLUS, lexer.currentChar)
	case '{':
		tok = newToken(token.LBRACE, lexer.currentChar)
	case '}':
		tok = newToken(token.RBRACE, lexer.currentChar)
	case '-':
		tok = newToken(token.MINUS, lexer.currentChar)
	case '!':
		if lexer.nextChar == '=' {
			lexer.readChar()
			tok = newTokenWithString(token.NOT_EQ, "!=")
		} else {
			tok = newToken(token.BANG, lexer.currentChar)
		}
	case '*':
		tok = newToken(token.ASTERISK, lexer.currentChar)
	case '/':
		tok = newToken(token.SLASH, lexer.currentChar)
	case '<':
		tok = newToken(token.LT, lexer.currentChar)
	case '>':
		tok = newToken(token.GT, lexer.currentChar)
	case 0:
		tok = newTokenWithString(token.EOF, "")
	default:
		if utils.IsLetter(lexer.currentChar) {
			// 读取所有连续的字符串
			literal := lexer.readIdentifier()
			// 使用文本匹配固定的字符串，找到特殊的关键字，例如 let、fn
			tok = newTokenWithString(token.LookupIdent(literal), literal)
		} else if utils.IsDigit(lexer.currentChar) {
			// 读取所哟连续的数字
			// 暂时只支持整数
			tok = newTokenWithString(token.INT, lexer.readNumber())
		} else {
			// 无法识别的范式
			tok = newToken(token.ILLEGAL, lexer.currentChar)
		}
	}

	// 读取下一个字符
	lexer.readChar()

	return tok
}
