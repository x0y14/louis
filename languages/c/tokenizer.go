package c

import (
	"fmt"
	"louis/interfaces"
	"strconv"
)

type Tokenizer struct {
	input   []rune
	curtPos int
}

func (t *Tokenizer) isEof() bool {
	return t.curtPos >= len(t.input)
}

func (t *Tokenizer) getCurrentRune() (*rune, error) {
	if t.curtPos < 0 || t.isEof() {
		return nil, fmt.Errorf("out of range: pos=%d, len=%d", t.curtPos, len(t.input))
	}
	return &t.input[t.curtPos], nil
}

func (t *Tokenizer) getNextRune() (*rune, error) {
	if t.curtPos+1 < 0 || t.curtPos+1 >= len(t.input) {
		return nil, fmt.Errorf("out of range: pos=%d, len=%d", t.curtPos+1, len(t.input))
	}
	return &t.input[t.curtPos+1], nil
}

func (t *Tokenizer) goAdvance() {
	t.curtPos++
}

func (t *Tokenizer) consumeWhitespace() ([]rune, error) {
	var whitespace []rune
Loop:
	for !t.isEof() {
		cr, err := t.getCurrentRune()
		if err != nil {
			return nil, err
		}
		switch *cr {
		case ' ', '\t':
			whitespace = append(whitespace, *cr)
			t.goAdvance()
		default:
			break Loop
		}
	}
	return whitespace, nil
}

func (t *Tokenizer) consumeNewline() ([]rune, error) {
	var newline []rune
Loop:
	for !t.isEof() {
		cr, err := t.getCurrentRune()
		if err != nil {
			return nil, err
		}
		switch *cr {
		case '\n', '\r':
			newline = append(newline, *cr)
			t.goAdvance()
		default:
			break Loop
		}
	}
	return newline, nil
}

func (t *Tokenizer) consumeComment() ([]rune, error) {
	var comment []rune
	t.goAdvance() // /
	t.goAdvance() // /
Loop:
	for !t.isEof() {
		cr, err := t.getCurrentRune()
		if err != nil {
			return nil, err
		}
		switch *cr {
		case '\n', '\r':
			break Loop
		default:
			comment = append(comment, *cr)
			t.goAdvance()
		}
	}
	return comment, nil
}

func (t *Tokenizer) isSingleSymbol() bool {
	cr, err := t.getCurrentRune()
	if err != nil {
		return false
	}

	switch string(*cr) {
	case ADD.String(), SUB.String(), MUL.String(), DIV.String(), MOD.String():
		return true
	case LSS.String(), GTR.String(), ASSIGN.String(), NOT.String():
		return true
	case LRB.String(), LSB.String(), LCB.String():
		return true
	case RRB.String(), RSB.String(), RCB.String():
		return true
	case DOT.String(), COMMA.String(), COLON.String(), SEMI.String():
		return true
	default:
		return false
	}
}

func (t *Tokenizer) isCompoundSymbol() bool {
	cr, err := t.getCurrentRune()
	if err != nil {
		return false
	}
	nr, err := t.getNextRune()
	if err != nil {
		return false
	}
	switch string(*cr) + string(*nr) {
	case EQL.String():
		return true
	case NEQ.String(), LEQ.String(), GEQ.String():
		return true
	case LAND.String(), LOR.String():
		return true
	default:
		return false
	}
}

func (t *Tokenizer) isSymbol() bool {
	isCompound := t.isCompoundSymbol()

	if isCompound {
		return true
	}

	isSingle := t.isSingleSymbol()

	if isSingle {
		return true
	}

	return false
}

func (t *Tokenizer) consumeSymbol() ([]rune, error) {
	// 複合記号
	isSym := t.isCompoundSymbol()
	if isSym {
		cr, err := t.getCurrentRune()
		if err != nil {
			return nil, err
		}
		nr, err := t.getNextRune()
		if err != nil {
			return nil, err
		}
		t.goAdvance()
		t.goAdvance()
		return []rune{*cr, *nr}, nil
	}

	// 単体記号
	isSym = t.isSingleSymbol()
	if isSym {
		cr, err := t.getCurrentRune()
		if err != nil {
			return nil, err
		}
		t.goAdvance()
		return []rune{*cr}, nil
	}

	return []rune{}, nil
}

func (t *Tokenizer) consumeString() ([]rune, error) {
	var str []rune
	t.goAdvance() // '"'
Loop:
	for !t.isEof() {
		cr, err := t.getCurrentRune()
		if err != nil {
			return nil, err
		}
		switch *cr {
		case '"':
			break Loop
		default:
			str = append(str, *cr)
			t.goAdvance()
		}
	}
	t.goAdvance() // '"'
	return str, nil
}

func (t *Tokenizer) consumeNumber() ([]rune, error) {
	var num []rune

	// check minus
	cr, err := t.getCurrentRune()
	if err != nil {
		return nil, err
	}
	if *cr == '-' {
		num = append(num, *cr)
		t.goAdvance() // '-'
	}
Loop:
	for !t.isEof() {
		cr, err = t.getCurrentRune()
		if err != nil {
			return nil, err
		}
		switch {
		case '0' <= *cr && *cr <= '9':
			num = append(num, *cr)
			t.goAdvance()
		case '.' == *cr:
			num = append(num, *cr)
			t.goAdvance()
		default:
			break Loop
		}
	}

	// 数値が.で終わっていないことを確認
	if len(num) > 0 && num[len(num)-1] == '.' {
		return nil, fmt.Errorf("numeric ending with dot: %s", string(num))
	}

	return num, nil
}

func (t *Tokenizer) consumeIdent() ([]rune, error) {
	var ident []rune
Loop:
	for !t.isEof() {
		cr, err := t.getCurrentRune()
		if err != nil {
			return nil, err
		}
		switch {
		case ('a' <= *cr && *cr <= 'z') || ('A' <= *cr && *cr <= 'Z'):
			ident = append(ident, *cr)
			t.goAdvance()
		case '_' == *cr:
			ident = append(ident, *cr)
			t.goAdvance()
		case '0' <= *cr && *cr <= '9':
			ident = append(ident, *cr)
			t.goAdvance()
		default:
			break Loop
		}
	}

	return ident, nil
}

func (t *Tokenizer) Tokenize(s string) (interfaces.Token, error) {
	t.input = []rune(s)
	t.curtPos = 0
	var head Token
	curt := &head

	//Loop:
	for !t.isEof() {
		cr, err := t.getCurrentRune()
		if err != nil {
			return nil, err
		}

		switch {
		// whitespace
		case *cr == ' ' || *cr == '\t':
			if _, err = t.consumeWhitespace(); err != nil {
				return nil, err
			}
		// newline
		case *cr == '\n' || *cr == '\r':
			if _, err = t.consumeNewline(); err != nil {
				return nil, err
			}
		// comment or symbol
		case *cr == '/':
			var nc *rune
			if nc, err = t.getNextRune(); err != nil {
				return nil, err
			}
			if *nc == '/' { // -> "// comment"
				if _, err = t.consumeComment(); err != nil {
					return nil, err
				}
			}
			fallthrough
		// symbol
		case t.isSymbol():
			symbol, err := t.consumeSymbol()
			if err != nil {
				return nil, err
			}
			kind, ok := symbolToTokenKind(string(symbol))
			if !ok {
				return nil, fmt.Errorf("unsupported symbol: %s", string(symbol))
			}
			tok := &Token{Kind: kind}
			curt.Next = tok
			curt = tok
		// string
		case *cr == '"':
			str, err := t.consumeString()
			if err != nil {
				return nil, err
			}
			tok := &Token{Kind: STRING, s: string(str)}
			curt.Next = tok
			curt = tok
		// number
		case *cr == '-' || ('0' <= *cr && *cr <= '9'):
			num, err := t.consumeNumber()
			if err != nil {
				return nil, err
			}
			tokType, ok := numericTokenKind(num)
			if !ok {
				return nil, fmt.Errorf("invalid numeric value: %s", string(num))
			}
			var tok *Token
			if tokType == FLOAT {
				f, _ := strconv.ParseFloat(string(num), 64)
				tok = &Token{Kind: FLOAT, f: f}
			} else {
				i, _ := strconv.Atoi(string(num))
				tok = &Token{Kind: INT, i: i}
			}
			curt.Next = tok
			curt = tok
		// ident
		case *cr == '_' || ('a' <= *cr && *cr <= 'z') || ('A' <= *cr && *cr <= 'Z'):
			ident, err := t.consumeIdent()
			if err != nil {
				return nil, err
			}
			tok := &Token{Kind: IDENT, s: string(ident)}
			curt.Next = tok
			curt = tok
		default:
			fmt.Println(string(*cr))
			t.goAdvance()
		}
	}

	tok := &Token{Kind: EOF}
	curt.Next = tok
	curt = tok

	return head.Next, nil
}
