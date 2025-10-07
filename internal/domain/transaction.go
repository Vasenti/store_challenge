package domain

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID uuid.UUID
	Date time.Time
	Amount float64
	Kind    Kind   
}

// Validate verifica reglas básicas de negocio.
func (t Transaction) Validate() error {
	if t.Kind != KindCredit && t.Kind != KindDebit {
		return fmt.Errorf("invalid transaction kind: %s", t.Kind.String())
	}
	if t.Amount <= 0 {
		return fmt.Errorf("amount must be positive")
	}
	return nil
}

func (t Transaction) SignedAmount() float64 {
	if t.Kind == KindDebit {
		return -t.Amount
	}
	return t.Amount
}