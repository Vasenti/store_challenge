package domain

import "fmt"

type Kind int

const (
	KindUnknown Kind = iota
	KindCredit
	KindDebit
)

func (k Kind) String() string {
	switch k {
	case KindCredit:
		return "CREDIT"
	case KindDebit:
		return "DEBIT"
	default:
		return "UNKNOWN"
	}
}

func ParseKind(s string) (Kind, error) {
	switch s {
	case "CREDIT", "Credit", "credit":
		return KindCredit, nil
	case "DEBIT", "Debit", "debit":
		return KindDebit, nil
	default:
		return KindUnknown, fmt.Errorf("invalid kind: %s", s)
	}
}