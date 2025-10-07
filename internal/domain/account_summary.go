package domain

type AccountSummary struct {
	TotalBalance float64
	ByMonth           []MonthlyStats
}