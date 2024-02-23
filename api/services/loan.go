package services

import (
	"math"
	"time"

	"github.com/jojiejo/hibank-loan/api/helpers"
)

const monthPerYear = 12

type Loan struct {
	Plafond      float64    `json:"plafond,omitempty"`
	Duration     int        `json:"duration,omitempty"`
	InterestRate float64    `json:"interest_rate,omitempty"`
	StartDate    *time.Time `json:"start_date,omitempty"`
}

type Installment struct {
	InstallmentNumber    int     `json:"installment_number"`
	Date                 string  `json:"date"`
	Annuity              float64 `json:"annuity"`
	InstallmentPrinciple float64 `json:"installment_principle"`
	InstallmentInterest  float64 `json:"installment_interest"`
	InstallmentBalance   float64 `json:"installment_balance"`
}

func (l *Loan) CalculateAnnuity() float64 {
	interestRate := l.InterestRate / 100 / monthPerYear
	return (l.Plafond * interestRate) / (1 - math.Pow((1+interestRate), -float64(l.Duration)))
}

func (l *Loan) CalculateInstallments() []Installment {
	installmentTotal := l.CalculateAnnuity()
	currentInstallmentBalance := l.Plafond

	var installments []Installment
	for j := 1; j <= l.Duration; j++ {
		currentInterest := (l.InterestRate / 100 / 360) * 30 * currentInstallmentBalance
		currentPrinciple := installmentTotal - currentInterest
		installments = append(installments, Installment{
			InstallmentNumber:    j,
			Date:                 l.StartDate.AddDate(0, j-1, 0).Format("2006-01-02"),
			Annuity:              helpers.Round(installmentTotal, 2),
			InstallmentPrinciple: helpers.Round(currentPrinciple, 0),
			InstallmentInterest:  helpers.Round(currentInterest, 0),
			InstallmentBalance:   helpers.Round((currentInstallmentBalance - currentPrinciple), 0),
		})
		currentInstallmentBalance -= currentPrinciple
	}

	return installments
}
