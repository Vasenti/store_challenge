package domain

type MonthlyStats struct {
	Month           MonthKey
	TxnCount        int
	AvgCreditAmount  float64
	AvgDebitAmount   float64
}