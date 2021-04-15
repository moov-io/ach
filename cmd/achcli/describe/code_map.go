// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package describe

import (
	"fmt"

	"github.com/moov-io/ach"
)

var serviceClassCodes = map[int]string{
	ach.MixedDebitsAndCredits:      "(Mixed Debits and Credits)",
	ach.CreditsOnly:                "(Credits Only)",
	ach.DebitsOnly:                 "(Debits Only)",
	ach.AutomatedAccountingAdvices: "(Automated Accounting Analysis)",
}

type transactionType string

const (
	debit  transactionType = "Debit"
	credit transactionType = "Credit"
)

var transactionCodes = map[int]string{
	ach.CheckingCredit:                     entry("Checking", credit),
	ach.CheckingReturnNOCCredit:            noc("Checking", credit),
	ach.CheckingPrenoteCredit:              prenote("Checking", credit),
	ach.CheckingZeroDollarRemittanceCredit: remittance("Checking", credit),

	ach.CheckingDebit:                     entry("Checking", debit),
	ach.CheckingReturnNOCDebit:            noc("Checking", debit),
	ach.CheckingPrenoteDebit:              prenote("Checking", debit),
	ach.CheckingZeroDollarRemittanceDebit: remittance("Checking", debit),

	ach.SavingsCredit:                     entry("Savings", credit),
	ach.SavingsReturnNOCCredit:            noc("Savings", credit),
	ach.SavingsPrenoteCredit:              prenote("Savings", credit),
	ach.SavingsZeroDollarRemittanceCredit: remittance("Savings", credit),

	ach.SavingsDebit:                     entry("Savings", debit),
	ach.SavingsReturnNOCDebit:            noc("Savings", debit),
	ach.SavingsPrenoteDebit:              prenote("Savings", debit),
	ach.SavingsZeroDollarRemittanceDebit: remittance("Savings", debit),

	ach.GLCredit:                     entry("GL", credit),
	ach.GLReturnNOCCredit:            noc("GL", credit),
	ach.GLPrenoteCredit:              prenote("GL", credit),
	ach.GLZeroDollarRemittanceCredit: remittance("Gl", credit),

	ach.GLDebit:                     entry("GL", debit),
	ach.GLReturnNOCDebit:            noc("GL", debit),
	ach.GLPrenoteDebit:              prenote("GL", debit),
	ach.GLZeroDollarRemittanceDebit: remittance("Gl", debit),

	ach.LoanCredit:                     entry("Loan", credit),
	ach.LoanReturnNOCCredit:            noc("Loan", credit),
	ach.LoanPrenoteCredit:              prenote("Loan", credit),
	ach.LoanZeroDollarRemittanceCredit: remittance("Loan", credit),

	ach.LoanDebit:          entry("Loan", debit),
	ach.LoanReturnNOCDebit: noc("Loan", debit),
}

func entry(s string, t transactionType) string {
	return fmt.Sprintf("(%s %s)", s, t)
}
func noc(s string, t transactionType) string {
	return fmt.Sprintf("(%s Return NOC %s)", s, t)
}
func prenote(s string, t transactionType) string {
	return fmt.Sprintf("(%s Prenote %s)", s, t)
}
func remittance(s string, t transactionType) string {
	return fmt.Sprintf("(%s Zero Dollar Remittance %s)", s, t)
}
