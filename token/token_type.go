package token

type TokenType string

var keyWords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keyWords[ident]; ok {
		return tok
	}

	return IDENT
}
