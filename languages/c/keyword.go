package c

type Keyword int

const (
	KW_ILLEGAL Keyword = iota
	KW_RETURN
)

func (kw Keyword) String() string {
	k := [...]string{
		KW_ILLEGAL: "illegal",
		KW_RETURN:  "return",
	}
	return k[kw]
}
