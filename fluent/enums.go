package fluent

type MemberType int64

const (
	MemberTypeAll    MemberType = 0
	MemberTypeField  MemberType = 1
	MemberTypeMethod MemberType = 2
)

type AdditiveInverse bool

const (
	NegativeDefault AdditiveInverse = false
	NegativeInvert  AdditiveInverse = true
)

type Match int

const (
	MatchAny Match = 1
	MatchAll Match = 2
)

type paramType struct {
	Field  string
	Method string
}

var ParamType paramType = paramType{
	Field:  "Field",
	Method: "Method",
}
