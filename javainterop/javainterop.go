package javainterop

import (
	"errors"
	"io"
	"time"

	"github.com/moov-io/ach"
)

type Addenda02 struct{ *ach.Addenda02 }

func NewAddenda02() *Addenda02 {
	return &Addenda02{ach.NewAddenda02()}
}

func (addenda02 *Addenda02) Parse(record string) {
	addenda02.Addenda02.Parse(record)
}

func (addenda02 *Addenda02) String() string {
	return addenda02.Addenda02.String()
}

func (addenda02 *Addenda02) Validate() error {
	return addenda02.Addenda02.Validate()
}

func (addenda02 *Addenda02) ReferenceInformationOneField() string {
	return addenda02.Addenda02.ReferenceInformationOneField()
}

func (addenda02 *Addenda02) ReferenceInformationTwoField() string {
	return addenda02.Addenda02.ReferenceInformationTwoField()
}

func (addenda02 *Addenda02) TerminalIdentificationCodeField() string {
	return addenda02.Addenda02.TerminalIdentificationCodeField()
}

func (addenda02 *Addenda02) TransactionSerialNumberField() string {
	return addenda02.Addenda02.TransactionSerialNumberField()
}

func (addenda02 *Addenda02) TransactionDateField() string {
	return addenda02.Addenda02.TransactionDateField()
}

func (addenda02 *Addenda02) AuthorizationCodeOrExpireDateField() string {
	return addenda02.Addenda02.AuthorizationCodeOrExpireDateField()
}

func (addenda02 *Addenda02) TerminalLocationField() string {
	return addenda02.Addenda02.TerminalLocationField()
}

func (addenda02 *Addenda02) TerminalCityField() string {
	return addenda02.Addenda02.TerminalCityField()
}

func (addenda02 *Addenda02) TerminalStateField() string {
	return addenda02.Addenda02.TerminalStateField()
}

func (addenda02 *Addenda02) TraceNumberField() string {
	return addenda02.Addenda02.TraceNumberField()
}

type Addenda05 struct{ *ach.Addenda05 }

func NewAddenda05() *Addenda05 {
	return &Addenda05{ach.NewAddenda05()}
}

func (addenda05 *Addenda05) Parse(record string) {
	addenda05.Addenda05.Parse(record)
}

func (addenda05 *Addenda05) String() string {
	return addenda05.Addenda05.String()
}

func (addenda05 *Addenda05) Validate() error {
	return addenda05.Addenda05.Validate()
}

func (addenda05 *Addenda05) PaymentRelatedInformationField() string {
	return addenda05.Addenda05.PaymentRelatedInformationField()
}

func (addenda05 *Addenda05) SequenceNumberField() string {
	return addenda05.Addenda05.SequenceNumberField()
}

func (addenda05 *Addenda05) EntryDetailSequenceNumberField() string {
	return addenda05.Addenda05.EntryDetailSequenceNumberField()
}

type Addenda10 struct{ *ach.Addenda10 }

func (addenda10 *Addenda10) Parse(record string) {
	addenda10.Addenda10.Parse(record)
}

func (addenda10 *Addenda10) String() string {
	return addenda10.Addenda10.String()
}

func (addenda10 *Addenda10) Validate() error {
	return addenda10.Addenda10.Validate()
}

func (addenda10 *Addenda10) ForeignPaymentAmountField() string {
	return addenda10.Addenda10.ForeignPaymentAmountField()
}

func (addenda10 *Addenda10) ForeignTraceNumberField() string {
	return addenda10.Addenda10.ForeignTraceNumberField()
}

func (addenda10 *Addenda10) NameField() string {
	return addenda10.Addenda10.NameField()
}

func (addenda10 *Addenda10) EntryDetailSequenceNumberField() string {
	return addenda10.Addenda10.EntryDetailSequenceNumberField()
}

type Addenda11 struct{ *ach.Addenda11 }

func NewAddenda11() *Addenda11 {
	return &Addenda11{ach.NewAddenda11()}
}

func (addenda11 *Addenda11) Parse(record string) {
	addenda11.Addenda11.Parse(record)
}

func (addenda11 *Addenda11) String() string {
	return addenda11.Addenda11.String()
}

func (addenda11 *Addenda11) Validate() error {
	return addenda11.Addenda11.Validate()
}

func (addenda11 *Addenda11) OriginatorNameField() string {
	return addenda11.Addenda11.OriginatorNameField()
}

func (addenda11 *Addenda11) OriginatorStreetAddressField() string {
	return addenda11.Addenda11.OriginatorStreetAddressField()
}

func (addenda11 *Addenda11) EntryDetailSequenceNumberField() string {
	return addenda11.Addenda11.EntryDetailSequenceNumberField()
}

type Addenda12 struct{ *ach.Addenda12 }

func NewAddenda12() *Addenda12 {
	return &Addenda12{ach.NewAddenda12()}
}

func (addenda12 *Addenda12) Parse(record string) {
	addenda12.Addenda12.Parse(record)
}

func (addenda12 *Addenda12) String() string {
	return addenda12.Addenda12.String()
}

func (addenda12 *Addenda12) Validate() error {
	return addenda12.Addenda12.Validate()
}

func (addenda12 *Addenda12) OriginatorCityStateProvinceField() string {
	return addenda12.Addenda12.OriginatorCityStateProvinceField()
}

func (addenda12 *Addenda12) OriginatorCountryPostalCodeField() string {
	return addenda12.Addenda12.OriginatorCountryPostalCodeField()
}

func (addenda12 *Addenda12) EntryDetailSequenceNumberField() string {
	return addenda12.Addenda12.EntryDetailSequenceNumberField()
}

type Addenda13 struct{ *ach.Addenda13 }

func NewAddenda13() *Addenda13 {
	return &Addenda13{ach.NewAddenda13()}
}

func (addenda13 *Addenda13) Parse(record string) {
	addenda13.Addenda13.Parse(record)
}

func (addenda13 *Addenda13) String() string {
	return addenda13.Addenda13.String()
}

func (addenda13 *Addenda13) Validate() error {
	return addenda13.Addenda13.Validate()
}

func (addenda13 *Addenda13) ODFINameField() string {
	return addenda13.Addenda13.ODFINameField()
}

func (addenda13 *Addenda13) ODFIIDNumberQualifierField() string {
	return addenda13.Addenda13.ODFIIDNumberQualifierField()
}

func (addenda13 *Addenda13) ODFIIdentificationField() string {
	return addenda13.Addenda13.ODFIIdentificationField()
}

func (addenda13 *Addenda13) ODFIBranchCountryCodeField() string {
	return addenda13.Addenda13.ODFIBranchCountryCodeField()
}

func (addenda13 *Addenda13) EntryDetailSequenceNumberField() string {
	return addenda13.Addenda13.EntryDetailSequenceNumberField()
}

type Addenda14 struct{ *ach.Addenda14 }

func NewAddenda14() *Addenda14 {
	return &Addenda14{ach.NewAddenda14()}
}

func (addenda14 *Addenda14) Parse(record string) {
	addenda14.Addenda14.Parse(record)
}

func (addenda14 *Addenda14) String() string {
	return addenda14.Addenda14.String()
}

func (addenda14 *Addenda14) Validate() error {
	return addenda14.Addenda14.Validate()
}
func (addenda14 *Addenda14) RDFINameField() string {
	return addenda14.Addenda14.RDFINameField()
}
func (addenda14 *Addenda14) RDFIIDNumberQualifierField() string {
	return addenda14.Addenda14.RDFIIDNumberQualifierField()
}

func (addenda14 *Addenda14) RDFIIdentificationField() string {
	return addenda14.Addenda14.RDFIIdentificationField()
}

func (addenda14 *Addenda14) RDFIBranchCountryCodeField() string {
	return addenda14.Addenda14.RDFIBranchCountryCodeField()
}

func (addenda14 *Addenda14) EntryDetailSequenceNumberField() string {
	return addenda14.Addenda14.EntryDetailSequenceNumberField()
}

type Addenda15 struct{ *ach.Addenda15 }

func NewAddenda15() *Addenda15 {
	return &Addenda15{ach.NewAddenda15()}
}

func (addenda15 *Addenda15) Parse(record string) {
	addenda15.Addenda15.Parse(record)
}

func (addenda15 *Addenda15) String() string {
	return addenda15.Addenda15.String()
}

func (addenda15 *Addenda15) Validate() error {
	return addenda15.Addenda15.Validate()
}

func (addenda15 *Addenda15) ReceiverIDNumberField() string {
	return addenda15.Addenda15.ReceiverIDNumberField()
}

func (addenda15 *Addenda15) ReceiverStreetAddressField() string {
	return addenda15.Addenda15.ReceiverStreetAddressField()
}

func (addenda15 *Addenda15) EntryDetailSequenceNumberField() string {
	return addenda15.Addenda15.EntryDetailSequenceNumberField()
}

type Addenda16 struct{ *ach.Addenda16 }

func NewAddenda16() *Addenda16 {
	return &Addenda16{ach.NewAddenda16()}
}

func (addenda16 *Addenda16) Parse(record string) {
	addenda16.Addenda16.Parse(record)
}

func (addenda16 *Addenda16) String() string {
	return addenda16.Addenda16.String()
}

func (addenda16 *Addenda16) Validate() error {
	return addenda16.Addenda16.Validate()
}
func (addenda16 *Addenda16) ReceiverCityStateProvinceField() string {
	return addenda16.Addenda16.ReceiverCityStateProvinceField()
}
func (addenda16 *Addenda16) ReceiverCountryPostalCodeField() string {
	return addenda16.Addenda16.ReceiverCountryPostalCodeField()
}

func (addenda16 *Addenda16) EntryDetailSequenceNumberField() string {
	return addenda16.Addenda16.EntryDetailSequenceNumberField()
}

type Addenda17 struct{ *ach.Addenda17 }

func NewAddenda17() *Addenda17 {
	return &Addenda17{ach.NewAddenda17()}
}

func (addenda17 *Addenda17) Parse(record string) {
	addenda17.Addenda17.Parse(record)
}

func (addenda17 *Addenda17) String() string {
	return addenda17.Addenda17.String()
}

func (addenda17 *Addenda17) Validate() error {
	return addenda17.Addenda17.Validate()
}
func (addenda17 *Addenda17) PaymentRelatedInformationField() string {
	return addenda17.Addenda17.PaymentRelatedInformationField()
}
func (addenda17 *Addenda17) SequenceNumberField() string {
	return addenda17.Addenda17.SequenceNumberField()
}

func (addenda17 *Addenda17) EntryDetailSequenceNumberField() string {
	return addenda17.Addenda17.EntryDetailSequenceNumberField()
}

type Addenda18 struct{ *ach.Addenda18 }

func NewAddenda18() *Addenda18 {
	return &Addenda18{ach.NewAddenda18()}
}

func (addenda18 *Addenda18) Parse(record string) {
	addenda18.Addenda18.Parse(record)
}

func (addenda18 *Addenda18) String() string {
	return addenda18.Addenda18.String()
}

func (addenda18 *Addenda18) Validate() error {
	return addenda18.Addenda18.Validate()
}

func (addenda18 *Addenda18) ForeignCorrespondentBankNameField() string {
	return addenda18.Addenda18.ForeignCorrespondentBankNameField()
}
func (addenda18 *Addenda18) ForeignCorrespondentBankIDNumberQualifierField() string {
	return addenda18.Addenda18.ForeignCorrespondentBankIDNumberQualifierField()
}

func (addenda18 *Addenda18) ForeignCorrespondentBankIDNumberField() string {
	return addenda18.Addenda18.ForeignCorrespondentBankIDNumberField()
}

func (addenda18 *Addenda18) ForeignCorrespondentBankBranchCountryCodeField() string {
	return addenda18.Addenda18.ForeignCorrespondentBankBranchCountryCodeField()
}

func (addenda18 *Addenda18) SequenceNumberField() string {
	return addenda18.Addenda18.SequenceNumberField()
}

func (addenda18 *Addenda18) EntryDetailSequenceNumberField() string {
	return addenda18.Addenda18.EntryDetailSequenceNumberField()
}

type Addenda98 struct{ *ach.Addenda98 }

func NewAddenda98() *Addenda98 {
	return &Addenda98{ach.NewAddenda98()}
}

type ChangeCode struct{ *ach.ChangeCode }

func (addenda98 *Addenda98) Parse(record string) {
	addenda98.Addenda98.Parse(record)
}

func (addenda98 *Addenda98) String() string {
	return addenda98.Addenda98.String()
}

func (addenda98 *Addenda98) Validate() error {
	return addenda98.Addenda98.Validate()
}

func (addenda98 *Addenda98) OriginalTraceField() string {
	return addenda98.Addenda98.OriginalTraceField()
}
func (addenda98 *Addenda98) OriginalDFIField() string {
	return addenda98.Addenda98.OriginalDFIField()
}

func (addenda98 *Addenda98) CorrectedDataField() string {
	return addenda98.Addenda98.CorrectedDataField()
}

func (addenda98 *Addenda98) TraceNumberField() string {
	return addenda98.Addenda98.TraceNumberField()
}

func (addenda98 *Addenda98) ChangeCodeField() *ChangeCode {
	return &ChangeCode{addenda98.Addenda98.ChangeCodeField()}
}

func LookupChangeCode(code string) *ChangeCode {
	return &ChangeCode{ach.LookupChangeCode(code)}
}

func IsRefusedChangeCode(code string) bool {
	return ach.IsRefusedChangeCode(code)
}

type CorrectedData struct{ *ach.CorrectedData }

func (addenda98 *Addenda98) ParseCorrectedData() *CorrectedData {
	return &CorrectedData{addenda98.Addenda98.ParseCorrectedData()}
}

func WriteCorrectionData(code string, data *ach.CorrectedData) string {
	return ach.WriteCorrectionData(code, data)

}

type Addenda98Refused struct{ *ach.Addenda98Refused }

func NewAddenda98Refused() *Addenda98Refused {
	return &Addenda98Refused{ach.NewAddenda98Refused()}
}

func (addenda98Refused *Addenda98Refused) Parse(record string) {
	addenda98Refused.Addenda98Refused.Parse(record)
}

func (addenda98Refused *Addenda98Refused) String() string {
	return addenda98Refused.Addenda98Refused.String()
}

func (addenda98Refused *Addenda98Refused) Validate() error {
	return addenda98Refused.Addenda98Refused.Validate()
}

func (addenda98Refused *Addenda98Refused) RefusedChangeCodeField() *ChangeCode {
	return &ChangeCode{addenda98Refused.Addenda98Refused.RefusedChangeCodeField()}
}

func (Addenda98Refused *Addenda98Refused) OriginalDFIField() string {
	return Addenda98Refused.Addenda98Refused.OriginalDFIField()
}

func (Addenda98Refused *Addenda98Refused) OriginalTraceField() string {
	return Addenda98Refused.Addenda98Refused.OriginalTraceField()
}

func (Addenda98Refused *Addenda98Refused) CorrectedDataField() string {
	return Addenda98Refused.Addenda98Refused.CorrectedDataField()
}

func (Addenda98Refused *Addenda98Refused) ChangeCodeField() *ChangeCode {
	return &ChangeCode{Addenda98Refused.Addenda98Refused.ChangeCodeField()}
}

func (Addenda98Refused *Addenda98Refused) TraceNumberField() string {
	return Addenda98Refused.Addenda98Refused.TraceNumberField()
}

func (Addenda98Refused *Addenda98Refused) TraceSequenceNumberField() string {
	return Addenda98Refused.Addenda98Refused.TraceSequenceNumberField()
}

type Addenda99 struct{ *ach.Addenda99 }

func NewAddenda99() *Addenda99 {
	return &Addenda99{ach.NewAddenda99()}
}

type ReturnCode struct{ *ach.ReturnCode }

func (addenda99 *Addenda99) Parse(record string) {
	addenda99.Addenda99.Parse(record)
}

func (addenda99 *Addenda99) String() string {
	return addenda99.Addenda99.String()
}

func (addenda99 *Addenda99) Validate() error {
	return addenda99.Addenda99.Validate()
}

func (addenda99 *Addenda99) SetValidation(opts *ach.ValidateOpts) {
	addenda99.Addenda99.SetValidation(opts)
}

func (addenda99 *Addenda99) OriginalTraceField() string {
	return addenda99.Addenda99.OriginalTraceField()
}
func (addenda99 *Addenda99) OriginalDFIField() string {
	return addenda99.Addenda99.OriginalDFIField()
}

func (addenda99 *Addenda99) DateOfDeathField() string {
	return addenda99.Addenda99.DateOfDeathField()
}

func (addenda99 *Addenda99) TraceNumberField() string {
	return addenda99.Addenda99.TraceNumberField()
}

func (addenda99 *Addenda99) AddendaInformationField() string {
	return addenda99.Addenda99.AddendaInformationField()
}

func (addenda99 *Addenda99) IATPaymentAmount(s string) {
	addenda99.Addenda99.IATPaymentAmount(s)
}

func (addenda99 *Addenda99) IATPaymentAmountField() int {
	return addenda99.Addenda99.IATPaymentAmountField()
}

func (addenda99 *Addenda99) IATAddendaInformation(s string) {
	addenda99.Addenda99.IATAddendaInformation(s)
}

func (addenda99 *Addenda99) IATAddendaInformationField() string {
	return addenda99.Addenda99.IATAddendaInformationField()
}

func (addenda99 *Addenda99) SetDishonoredAddendaInformation(
	returnTraceNumber string,
	returnSettlementDate string,
	returnReasonCode string,
	addenda string,
) {
	addenda99.Addenda99.SetDishonoredAddendaInformation(
		returnTraceNumber,
		returnSettlementDate,
		returnReasonCode,
		addenda,
	)
}

func (addenda99 *Addenda99) SetContestedAddendaInformation(
	originalSettlementDate string,
	returnTraceNumber string,
	returnSettlementDate string,
	returnReasonCode string,
	dishonoredReturnTraceNumber string,
	dishonoredReturnSettlementDate string,
	dishonoredReturnReasonCode string,
) {
	addenda99.Addenda99.SetContestedAddendaInformation(
		originalSettlementDate,
		returnTraceNumber,
		returnSettlementDate,
		returnReasonCode,
		dishonoredReturnTraceNumber,
		dishonoredReturnSettlementDate,
		dishonoredReturnReasonCode,
	)
}

func (addenda99 *Addenda99) AddendaInformationReturnTraceNumber() string {
	return addenda99.Addenda99.AddendaInformationReturnTraceNumber()
}

func (addenda99 *Addenda99) AddendaInformationReturnSettlementDate() string {
	return addenda99.Addenda99.AddendaInformationReturnSettlementDate()
}

func (addenda99 *Addenda99) AddendaInformationReturnReasonCode() string {
	return addenda99.Addenda99.AddendaInformationReturnReasonCode()
}

func (addenda99 *Addenda99) AddendaInformationExtra() string {
	return addenda99.Addenda99.AddendaInformationExtra()
}

func (addenda99 *Addenda99) SetOriginalEntryReturnDate(date string) {
	addenda99.Addenda99.SetOriginalEntryReturnDate(date)
}

func (addenda99 *Addenda99) OriginalEntryReturnDate() string {
	return addenda99.Addenda99.OriginalEntryReturnDate()
}

func (addenda99 *Addenda99) ReturnCodeField() *ReturnCode {
	return &ReturnCode{addenda99.Addenda99.ReturnCodeField()}
}

func LookupReturnCode(code string) *ReturnCode {
	return &ReturnCode{ach.LookupReturnCode(code)}
}

type Addenda99Dishonored struct{ *ach.Addenda99Dishonored }

func NewAddenda99Dishonored() *Addenda99Dishonored {
	return &Addenda99Dishonored{ach.NewAddenda99Dishonored()}
}

func (addenda99Dishonored *Addenda99Dishonored) Parse(record string) {
	addenda99Dishonored.Addenda99Dishonored.Parse(record)
}

func (addenda99Dishonored *Addenda99Dishonored) String() string {
	return addenda99Dishonored.Addenda99Dishonored.String()
}

func (addenda99Dishonored *Addenda99Dishonored) Validate() error {
	return addenda99Dishonored.Addenda99Dishonored.Validate()
}

func (Addenda99Dishonored *Addenda99Dishonored) TraceNumberField() string {
	return Addenda99Dishonored.Addenda99Dishonored.TraceNumberField()
}

func (Addenda99Dishonored *Addenda99Dishonored) DishonoredReturnReasonCodeField() string {
	return Addenda99Dishonored.Addenda99Dishonored.DishonoredReturnReasonCodeField()
}

func (Addenda99Dishonored *Addenda99Dishonored) OriginalEntryTraceNumberField() string {
	return Addenda99Dishonored.Addenda99Dishonored.OriginalEntryTraceNumberField()
}

func (Addenda99Dishonored *Addenda99Dishonored) OriginalReceivingDFIIdentificationField() string {
	return Addenda99Dishonored.Addenda99Dishonored.OriginalReceivingDFIIdentificationField()
}

func (Addenda99Dishonored *Addenda99Dishonored) ReturnTraceNumberField() string {
	return Addenda99Dishonored.Addenda99Dishonored.DishonoredReturnReasonCodeField()
}

func (Addenda99Dishonored *Addenda99Dishonored) ReturnSettlementDateField() string {
	return Addenda99Dishonored.Addenda99Dishonored.ReturnSettlementDateField()
}

func (Addenda99Dishonored *Addenda99Dishonored) ReturnReasonCodeField() string {
	return Addenda99Dishonored.Addenda99Dishonored.ReturnReasonCodeField()
}

func (Addenda99Dishonored *Addenda99Dishonored) AddendaInformationField() string {
	return Addenda99Dishonored.Addenda99Dishonored.AddendaInformationField()
}

func (Addenda99Dishonored *Addenda99Dishonored) SetValidation(opts *ach.ValidateOpts) {
	Addenda99Dishonored.Addenda99Dishonored.SetValidation(opts)
}

func IsDishonoredReturnCode(code string) bool {
	return ach.IsDishonoredReturnCode(code)
}

type ADVBatchControl struct{ *ach.ADVBatchControl }

func NewADVBatchControl() *ADVBatchControl {
	return &ADVBatchControl{ach.NewADVBatchControl()}
}

func (ADVBatchControl *ADVBatchControl) Parse(record string) {
	ADVBatchControl.ADVBatchControl.Parse(record)
}

func (ADVBatchControl *ADVBatchControl) String() string {
	return ADVBatchControl.ADVBatchControl.String()
}

func (ADVBatchControl *ADVBatchControl) Validate() error {
	return ADVBatchControl.ADVBatchControl.Validate()
}

func (bc *ADVBatchControl) EntryAddendaCountField() string {
	return bc.ADVBatchControl.EntryAddendaCountField()
}

func (bc *ADVBatchControl) EntryHashField() string {
	return bc.ADVBatchControl.EntryHashField()
}

func (bc *ADVBatchControl) TotalDebitEntryDollarAmountField() string {
	return bc.ADVBatchControl.TotalDebitEntryDollarAmountField()
}

func (bc *ADVBatchControl) TotalCreditEntryDollarAmountField() string {
	return bc.ADVBatchControl.TotalCreditEntryDollarAmountField()
}

func (bc *ADVBatchControl) ACHOperatorDataField() string {
	return bc.ADVBatchControl.ACHOperatorDataField()
}

func (bc *ADVBatchControl) ODFIIdentificationField() string {
	return bc.ADVBatchControl.ODFIIdentificationField()
}

func (bc *ADVBatchControl) BatchNumberField() string {
	return bc.ADVBatchControl.BatchNumberField()
}

type ADVEntryDetail struct{ *ach.ADVEntryDetail }

func NewADVEntryDetail() *ADVEntryDetail {
	return &ADVEntryDetail{ach.NewADVEntryDetail()}
}

func (ed *ADVEntryDetail) Parse(record string) {
	ed.ADVEntryDetail.Parse(record)
}

func (ed *ADVEntryDetail) String() string {
	return ed.ADVEntryDetail.String()
}

func (ed *ADVEntryDetail) Validate() error {
	return ed.ADVEntryDetail.Validate()
}

const (
	CreditForDebitsOriginated = 81
	CreditForCreditsReceived  = 83
	CreditForCreditsRejected  = 85
	CreditSummary             = 87

	DebitForCreditsOriginated     = 82
	DebitForDebitsReceived        = 84
	DebitForDebitsRejectedBatches = 86
	DebitSummary                  = 88
)

func (ed *ADVEntryDetail) SetRDFI(rdfi string) *ADVEntryDetail {
	return &ADVEntryDetail{ed.ADVEntryDetail.SetRDFI(rdfi)}
}

func (ed *ADVEntryDetail) RDFIIdentificationField() string {
	return ed.ADVEntryDetail.RDFIIdentificationField()
}

func (ed *ADVEntryDetail) DFIAccountNumberField() string {
	return ed.ADVEntryDetail.DFIAccountNumberField()
}

func (ed *ADVEntryDetail) AmountField() string {
	return ed.ADVEntryDetail.AmountField()
}

func (ed *ADVEntryDetail) AdviceRoutingNumberField() string {
	return ed.ADVEntryDetail.AdviceRoutingNumberField()
}

func (ed *ADVEntryDetail) FileIdentificationField() string {
	return ed.ADVEntryDetail.FileIdentificationField()
}

func (ed *ADVEntryDetail) ACHOperatorDataField() string {
	return ed.ADVEntryDetail.ACHOperatorDataField()
}

func (ed *ADVEntryDetail) IndividualNameField() string {
	return ed.ADVEntryDetail.IndividualNameField()
}

func (ed *ADVEntryDetail) DiscretionaryDataField() string {
	return ed.ADVEntryDetail.DiscretionaryDataField()
}

func (ed *ADVEntryDetail) ACHOperatorRoutingNumberField() string {
	return ed.ADVEntryDetail.ACHOperatorRoutingNumberField()
}

func (ed *ADVEntryDetail) JulianDateDayField() string {
	return ed.ADVEntryDetail.JulianDateDayField()
}

func (ed *ADVEntryDetail) SequenceNumberField() string {
	return ed.ADVEntryDetail.SequenceNumberField()
}

type ADVFileControl struct{ ach.ADVFileControl }

func NewADVFileControl() ADVFileControl {
	return ADVFileControl{ach.NewADVFileControl()}
}

func (fc *ADVFileControl) Parse(record string) {
	fc.ADVFileControl.Parse(record)
}

func (fc *ADVFileControl) String() string {
	return fc.ADVFileControl.String()
}

func (fc *ADVFileControl) Validate() error {
	return fc.ADVFileControl.Validate()
}

func (fc *ADVFileControl) BatchCountField() string {
	return fc.ADVFileControl.BatchCountField()
}

func (fc *ADVFileControl) BlockCountField() string {
	return fc.ADVFileControl.BlockCountField()
}

func (fc *ADVFileControl) EntryAddendaCountField() string {
	return fc.ADVFileControl.EntryAddendaCountField()
}

func (fc *ADVFileControl) EntryHashField() string {
	return fc.ADVFileControl.EntryHashField()
}

func (fc *ADVFileControl) TotalDebitEntryDollarAmountInFileField() string {
	return fc.ADVFileControl.TotalDebitEntryDollarAmountInFileField()
}

func (fc *ADVFileControl) TotalCreditEntryDollarAmountInFileField() string {
	return fc.ADVFileControl.TotalCreditEntryDollarAmountInFileField()
}

type Batch struct{ ach.Batch }

const (
	ACK = "ACK"
	ADV = "ADV"
	ARC = "ARC"
	ATX = "ATX"
	BOC = "BOC"
	CCD = "CCD"
	CIE = "CIE"
	COR = "COR"
	CTX = "CTX"
	DNE = "DNE"
	ENR = "ENR"
	IAT = "IAT"
	MTE = "MTE"
	POP = "POP"
	POS = "POS"
	PPD = "PPD"
	RCK = "RCK"
	SHR = "SHR"
	TEL = "TEL"
	TRC = "TRC"
	TRX = "TRX"
	WEB = "WEB"
	XCK = "XCK"
)

func (batch *Batch) MarshalJSON() ([]byte, error) {
	return batch.Batch.MarshalJSON()
}

func (batch *Batch) UnmarshalJSON(p []byte) error {
	return batch.Batch.UnmarshalJSON(p)
}

type BatchHeader struct{ *ach.BatchHeader }
type Batcher struct{ ach.Batcher }
type BatchControl struct{ *ach.BatchControl }

func NewBatch(bh *BatchHeader) (Batcher, error) {
	bt, err := ach.NewBatch(bh.BatchHeader)
	return Batcher{bt}, err
}

func (batch *Batch) Create() error {
	return batch.Batch.Create()
}

func (batch *Batch) Validate() error {
	return batch.Batch.Validate()
}

func (batch *Batch) SetValidation(opts *ach.ValidateOpts) {
	batch.Batch.SetValidation(opts)
}

func (batch *Batch) SetHeader(batchHeader *BatchHeader) {
	batch.Batch.SetHeader(batchHeader.BatchHeader)
}

func (batch *Batch) GetHeader() *BatchHeader {
	return &BatchHeader{batch.Batch.Header}
}

func (batch *Batch) SetControl(batchControl *BatchControl) {
	batch.Control = batchControl.BatchControl
}

func (batch *Batch) GetControl() *BatchControl {
	return &BatchControl{batch.Batch.GetControl()}
}

func (batch *Batch) SetADVControl(batchADVControl *ADVBatchControl) {
	batch.ADVControl = batchADVControl.ADVBatchControl
}

func (batch *Batch) GetADVControl() *ADVBatchControl {
	return &ADVBatchControl{batch.Batch.ADVControl}
}

type EntryDetail struct{ ach.EntryDetail }

func (batch *Batch) GetEntries() []*EntryDetail {
	entries := make([]*EntryDetail, 0, len(batch.Batch.GetEntries()))
	for _, v := range batch.Batch.GetEntries() {
		entries = append(entries, &EntryDetail{*v})
	}
	return entries
}

func (batch *Batch) AddEntry(entry *EntryDetail) {
	batch.Batch.AddEntry(&entry.EntryDetail)
}

func (batch *Batch) AddADVEntry(entry *ADVEntryDetail) {
	batch.Batch.AddADVEntry(entry.ADVEntryDetail)
}

func (batch *Batch) GetADVEntries() []*ADVEntryDetail {
	entries := make([]*ADVEntryDetail, 0, len(batch.Batch.GetADVEntries()))
	for _, v := range batch.Batch.GetADVEntries() {
		entries = append(entries, &ADVEntryDetail{v})
	}
	return entries
}

func (batch *Batch) Category() string {
	return batch.Batch.Category()
}

func (batch *Batch) ID() string {
	return batch.Batch.ID()
}

func (batch *Batch) SetID(id string) {
	batch.Batch.SetID(id)
}

func (batch *Batch) IsADV() bool {
	return batch.Batch.IsADV()
}

func (batch *Batch) ValidAmountForCodes(entry *EntryDetail) error {
	return batch.Batch.ValidAmountForCodes(&entry.EntryDetail)
}

func (batch *Batch) ValidTranCodeForServiceClassCode(entry *EntryDetail) error {
	return batch.Batch.ValidTranCodeForServiceClassCode(&entry.EntryDetail)
}

func (batch *Batch) Equal(other Batcher) bool {
	return batch.Batch.Equal(other)
}

type Offset = *ach.Offset

func (b *Batch) WithOffset(off Offset) {
	b.Batch.WithOffset(off)
}

//const offsetIndividualName = "OFFSET"

type BatchACK struct {
	ach.BatchACK
}

func NewBatchACK(bh BatchHeader) *BatchACK {
	return &BatchACK{*ach.NewBatchACK(bh.BatchHeader)}
}

func (batch *BatchACK) Validate() error {
	return batch.Batch.Validate()
}

func (batch *BatchACK) Create() error {
	return batch.Batch.Create()
}

type BatchADV struct {
	ach.BatchADV
}

func NewBatchADV(bh *BatchHeader) *BatchADV {
	return &BatchADV{*ach.NewBatchADV(bh.BatchHeader)}
}

func (batch *BatchADV) Validate() error {
	return batch.Batch.Validate()
}

func (batch *BatchADV) Create() error {
	return batch.Batch.Create()
}

type BatchARC struct {
	ach.BatchARC
}

func NewBatchARC(bh *BatchHeader) *BatchARC {
	return &BatchARC{*ach.NewBatchARC(bh.BatchHeader)}
}

func (batch *BatchARC) Validate() error {
	return batch.Batch.Validate()
}

func (batch *BatchARC) Create() error {
	return batch.Batch.Create()
}

type BatchATX struct {
	ach.BatchATX
}

func NewBatchATX(bh *BatchHeader) *BatchATX {
	return &BatchATX{*ach.NewBatchATX(bh.BatchHeader)}
}

func (batch *BatchATX) Validate() error {
	return batch.Batch.Validate()
}

func (batch *BatchATX) Create() error {
	return batch.Batch.Create()
}

type BatchBOC struct {
	ach.BatchBOC
}

func NewBatchBOC(bh *BatchHeader) *BatchBOC {
	return &BatchBOC{*ach.NewBatchBOC(bh.BatchHeader)}
}

func (batch *BatchBOC) Validate() error {
	return batch.Batch.Validate()
}

func (batch *BatchBOC) Create() error {
	return batch.Batch.Create()
}

type BatchCCD struct {
	ach.BatchCCD
}

func NewBatchCCD(bh *BatchHeader) *BatchCCD {
	return &BatchCCD{*ach.NewBatchCCD(bh.BatchHeader)}
}

func (batch *BatchCCD) Validate() error {
	return batch.Batch.Validate()
}

func (batch *BatchCCD) Create() error {
	return batch.Batch.Create()
}

type BatchCIE struct {
	ach.BatchCIE
}

func NewBatchCIE(bh *BatchHeader) *BatchCIE {
	return &BatchCIE{*ach.NewBatchCIE(bh.BatchHeader)}
}

func (batch *BatchCIE) Validate() error {
	return batch.Batch.Validate()
}

func (batch *BatchCIE) Create() error {
	return batch.Batch.Create()
}

type BatchCOR struct {
	ach.BatchCOR
}

func NewBatchCOR(bh *BatchHeader) *BatchCOR {
	return &BatchCOR{*ach.NewBatchCOR(bh.BatchHeader)}
}

func (batch *BatchCOR) Validate() error {
	return batch.Batch.Validate()
}

func (batch *BatchCOR) Create() error {
	return batch.Batch.Create()
}

type BatchCTX struct {
	ach.BatchCTX
}

func NewBatchCTX(bh *BatchHeader) *BatchCTX {
	return &BatchCTX{*ach.NewBatchCTX(bh.BatchHeader)}
}

func (batch *BatchCTX) Validate() error {
	return batch.Batch.Validate()
}

func (batch *BatchCTX) Create() error {
	return batch.Batch.Create()
}

type ValidateOpts struct{ *ach.ValidateOpts }

func (bc *BatchControl) SetValidation(opts *ach.ValidateOpts) {

}

func (bc *BatchControl) String() string {
	return bc.BatchControl.String()
}

func (bc *BatchControl) EntryAddendaCountField() string {
	return bc.BatchControl.EntryAddendaCountField()
}

func (bc *BatchControl) EntryHashField() string {
	return bc.BatchControl.EntryHashField()
}

func (bc *BatchControl) TotalDebitEntryDollarAmountField() string {
	return bc.BatchControl.TotalDebitEntryDollarAmountField()
}

func (bc *BatchControl) TotalCreditEntryDollarAmountField() string {
	return bc.BatchControl.TotalCreditEntryDollarAmountField()
}

func (bc *BatchControl) CompanyIdentificationField() string {
	return bc.BatchControl.CompanyIdentificationField()
}

func (bc *BatchControl) MessageAuthenticationCodeField() string {
	return bc.BatchControl.MessageAuthenticationCodeField()
}

func (bc *BatchControl) ODFIIdentificationField() string {
	return bc.BatchControl.ODFIIdentificationField()
}

func (bc *BatchControl) BatchNumberField() string {
	return bc.BatchControl.BatchNumberField()
}

type BatchDNE struct {
	ach.BatchDNE
}

func NewBatchDNE(bh *BatchHeader) *BatchDNE {
	return &BatchDNE{*ach.NewBatchDNE(bh.BatchHeader)}
}

func (batch *BatchDNE) Validate() error {
	return batch.Batch.Validate()
}

func (batch *BatchDNE) Create() error {
	return batch.Batch.Create()
}

func (batch *BatchDNE) DateOfDeath() string {
	return batch.BatchDNE.DateOfDeath()
}

func (batch *BatchDNE) CustomerSSN() string {
	return batch.BatchDNE.CustomerSSN()
}

func (batch *BatchDNE) Amount() string {
	return batch.BatchDNE.Amount()
}

type BatchENR struct {
	ach.BatchENR
}

func NewBatchENR(bh *BatchHeader) *BatchENR {
	return &BatchENR{*ach.NewBatchENR(bh.BatchHeader)}
}

func (batch *BatchENR) Validate() error {
	return batch.Batch.Validate()
}

func (batch *BatchENR) Create() error {
	return batch.Batch.Create()
}

type ENRPaymentInformation struct {
	ach.ENRPaymentInformation
}

func (info *ENRPaymentInformation) String() string {
	return info.ENRPaymentInformation.String()
}

func (batch *BatchENR) ParsePaymentInformation(addenda05 *Addenda05) (*ach.ENRPaymentInformation, error) {
	return batch.BatchENR.ParsePaymentInformation(addenda05.Addenda05)
}

var (
	ErrBatchNoEntries                         = errors.New("must have Entry Record(s) to be built")
	ErrBatchADVCount                          = errors.New("there can be a maximum of 9999 ADV Sequence Numbers (ADV Entry Detail Records)")
	ErrBatchAddendaIndicator                  = errors.New("is 0 but found addenda record(s)")
	ErrBatchOriginatorDNE                     = errors.New("only government agencies (originator status code 2) can originate a DNE")
	ErrBatchInvalidCardTransactionType        = errors.New("invalid card transaction type")
	ErrBatchDebitOnly                         = errors.New("this batch type does not allow credit transaction codes")
	ErrBatchCheckSerialNumber                 = errors.New("this batch type requires entries to have Check Serial Numbers")
	ErrBatchSECType                           = errors.New("header SEC does not match this batch's type")
	ErrBatchServiceClassCode                  = errors.New("header SCC is not valid for this batch's type")
	ErrBatchTransactionCode                   = errors.New("transaction code is not valid for this batch's type")
	ErrBatchTransactionCodeAddenda            = errors.New("this batch type does not allow an addenda for this transaction code")
	ErrBatchAmountNonZero                     = errors.New("this batch type requires that the amount is zero")
	ErrBatchAmountZero                        = errors.New("this batch type requires that the amount is non-zero")
	ErrBatchCompanyEntryDescriptionAutoenroll = errors.New("this batch type requires that the Company Entry Description is AUTOENROLL")
	ErrBatchCompanyEntryDescriptionREDEPCHECK = errors.New("this batch type requires that the Company Entry Description is REDEPCHECK")
	ErrBatchAddendaCategory                   = errors.New("this batch type does not allow this addenda for category")
)

type BatchError struct {
	ach.BatchError
}

func (e *BatchError) Error() string {
	return e.BatchError.Error()
}

func (e *BatchError) Unwrap() error {
	return e.BatchError.Unwrap()
}

type IATBatch struct {
	ach.IATBatch
}

func (iatBatch *IATBatch) Error(field string, err error, values ...interface{}) error {
	return iatBatch.IATBatch.Error(field, err, values...)
}

type ErrBatchHeaderControlEquality struct {
	ach.ErrBatchHeaderControlEquality
}

func NewErrBatchHeaderControlEquality(header, control interface{}) *ErrBatchHeaderControlEquality {
	return &ErrBatchHeaderControlEquality{ach.NewErrBatchHeaderControlEquality(header, control)}
}

func (e ErrBatchHeaderControlEquality) Error() string {
	return e.ErrBatchHeaderControlEquality.Error()
}

type ErrBatchCalculatedControlEquality struct {
	ach.ErrBatchCalculatedControlEquality
}

func NewErrBatchCalculatedControlEquality(calculated, control interface{}) *ErrBatchCalculatedControlEquality {
	return &ErrBatchCalculatedControlEquality{ach.NewErrBatchCalculatedControlEquality(calculated, control)}
}

func (e ErrBatchCalculatedControlEquality) Error() string {
	return e.ErrBatchCalculatedControlEquality.Error()
}

type ErrBatchAscending struct {
	ach.ErrBatchAscending
}

func NewErrBatchAscending(header, control interface{}) *ErrBatchAscending {
	return &ErrBatchAscending{ach.NewErrBatchAscending(header, control)}
}

func (e ErrBatchAscending) Error() string {
	return e.ErrBatchAscending.Error()
}

type ErrBatchCategory struct {
	ach.ErrBatchCategory
}

func NewErrBatchCategory(categoryA, categoryB string) *ErrBatchCategory {
	return &ErrBatchCategory{ach.NewErrBatchCategory(categoryA, categoryB)}
}

func (e ErrBatchCategory) Error() string {
	return e.ErrBatchCategory.Error()
}

type ErrBatchTraceNumberNotODFI struct {
	ach.ErrBatchTraceNumberNotODFI
}

func NewErrBatchTraceNumberNotODFI(odfi, trace string) *ErrBatchTraceNumberNotODFI {
	return &ErrBatchTraceNumberNotODFI{ach.NewErrBatchTraceNumberNotODFI(odfi, trace)}
}

func (e ErrBatchTraceNumberNotODFI) Error() string {
	return e.ErrBatchTraceNumberNotODFI.Error()
}

type ErrBatchAddendaTraceNumber struct {
	ach.ErrBatchAddendaTraceNumber
}

func NewErrBatchAddendaTraceNumber(entryDetail, trace string) *ErrBatchAddendaTraceNumber {
	return &ErrBatchAddendaTraceNumber{ach.NewErrBatchAddendaTraceNumber(entryDetail, trace)}
}

func (e ErrBatchAddendaTraceNumber) Error() string {
	return e.ErrBatchAddendaTraceNumber.Error()
}

type ErrBatchAddendaCount struct {
	ach.ErrBatchAddendaCount
}

func NewErrBatchAddendaCount(found, allowed int) *ErrBatchAddendaCount {
	return &ErrBatchAddendaCount{ach.NewErrBatchAddendaCount(found, allowed)}
}

func (e ErrBatchAddendaCount) Error() string {
	return e.ErrBatchAddendaCount.Error()
}

type ErrBatchRequiredAddendaCount struct {
	ach.ErrBatchRequiredAddendaCount
}

func NewErrBatchRequiredAddendaCount(found, required int) *ErrBatchRequiredAddendaCount {
	return &ErrBatchRequiredAddendaCount{ach.NewErrBatchRequiredAddendaCount(found, required)}
}

func (e ErrBatchRequiredAddendaCount) Error() string {
	return e.ErrBatchRequiredAddendaCount.Error()
}

type ErrBatchExpectedAddendaCount struct {
	ach.ErrBatchExpectedAddendaCount
}

func NewErrBatchExpectedAddendaCount(found, expected int) *ErrBatchExpectedAddendaCount {
	return &ErrBatchExpectedAddendaCount{ach.NewErrBatchExpectedAddendaCount(found, expected)}
}

func (e ErrBatchExpectedAddendaCount) Error() string {
	return e.ErrBatchExpectedAddendaCount.Error()
}

type ErrBatchServiceClassTranCode struct {
	ach.ErrBatchServiceClassTranCode
}

func NewErrBatchServiceClassTranCode(serviceClassCode, transactionCode int) *ErrBatchServiceClassTranCode {
	return &ErrBatchServiceClassTranCode{ach.NewErrBatchServiceClassTranCode(serviceClassCode, transactionCode)}
}

func (e ErrBatchServiceClassTranCode) Error() string {
	return e.ErrBatchServiceClassTranCode.Error()
}

type ErrBatchAmount struct {
	ach.ErrBatchAmount
}

func NewErrBatchAmount(amount, limit int) *ErrBatchAmount {
	return &ErrBatchAmount{ach.NewErrBatchAmount(amount, limit)}
}

func (e ErrBatchAmount) Error() string {
	return e.ErrBatchAmount.Error()
}

type ErrBatchIATNOC struct {
	ach.ErrBatchIATNOC
}

func NewErrBatchIATNOC(found, expected interface{}) *ErrBatchIATNOC {
	return &ErrBatchIATNOC{ach.NewErrBatchIATNOC(found, expected)}
}

func (e ErrBatchIATNOC) Error() string {
	return e.ErrBatchIATNOC.Error()
}

func NewBatchHeader() *BatchHeader {
	return &BatchHeader{ach.NewBatchHeader()}
}

func (BatchHeader *BatchHeader) Parse(record string) {
	BatchHeader.BatchHeader.Parse(record)
}

func (BatchHeader *BatchHeader) String() string {
	return BatchHeader.BatchHeader.String()
}

func (BatchHeader *BatchHeader) Equal(other *BatchHeader) bool {
	return BatchHeader.BatchHeader.Equal(other.BatchHeader)
}

func (BatchHeader *BatchHeader) SetValidation(opts *ValidateOpts) {
	BatchHeader.BatchHeader.SetValidation(opts.ValidateOpts)
}

func (BatchHeader *BatchHeader) Validate() error {
	return BatchHeader.BatchHeader.Validate()
}

const (
	MixedDebitsAndCredits      = 200
	CreditsOnly                = 220
	DebitsOnly                 = 225
	AutomatedAccountingAdvices = 280
)

func (bh *BatchHeader) CompanyNameField() string {
	return bh.BatchHeader.CompanyNameField()
}

func (bh *BatchHeader) CompanyDiscretionaryDataField() string {
	return bh.BatchHeader.CompanyDiscretionaryDataField()
}

func (bh *BatchHeader) CompanyIdentificationField() string {
	return bh.BatchHeader.CompanyIdentificationField()
}

func (bh *BatchHeader) CompanyEntryDescriptionField() string {
	return bh.BatchHeader.CompanyEntryDescriptionField()
}

func (bh *BatchHeader) CompanyDescriptiveDateField() string {
	return bh.BatchHeader.CompanyDescriptiveDateField()
}

func (bh *BatchHeader) EffectiveEntryDateField() string {
	return bh.BatchHeader.EffectiveEntryDateField()
}

func (bh *BatchHeader) ODFIIdentificationField() string {
	return bh.BatchHeader.ODFIIdentificationField()
}

func (bh *BatchHeader) BatchNumberField() string {
	return bh.BatchHeader.BatchNumberField()
}

func (bh *BatchHeader) SettlementDateField() string {
	return bh.BatchHeader.SettlementDateField()
}

func (bh *BatchHeader) LiftEffectiveEntryDate() (time.Time, error) {
	return bh.BatchHeader.LiftEffectiveEntryDate()
}

type BatchMTE struct{ ach.BatchMTE }

func NewBatchMTE(bh *BatchHeader) *BatchMTE {
	return &BatchMTE{*ach.NewBatchMTE(bh.BatchHeader)}
}

func (batch *BatchMTE) Validate() error {
	return batch.BatchMTE.Validate()
}

func (batch *BatchMTE) Create() error {
	return batch.BatchMTE.Create()
}

type BatchPOP struct{ ach.BatchPOP }

func NewBatchPOP(bh *BatchHeader) *BatchPOP {
	return &BatchPOP{*ach.NewBatchPOP(bh.BatchHeader)}
}

func (batch *BatchPOP) Validate() error {
	return batch.BatchPOP.Validate()
}

func (batch *BatchPOP) Create() error {
	return batch.BatchPOP.Create()
}

type BatchPOS struct{ ach.BatchPOS }

func NewBatchPOS(bh *BatchHeader) *BatchPOS {
	return &BatchPOS{*ach.NewBatchPOS(bh.BatchHeader)}
}

func (batch *BatchPOS) Validate() error {
	return batch.BatchPOS.Validate()
}

func (batch *BatchPOS) Create() error {
	return batch.BatchPOS.Create()
}

type BatchPPD struct{ ach.BatchPPD }

func NewBatchPPD(bh *BatchHeader) *BatchPPD {
	return &BatchPPD{*ach.NewBatchPPD(bh.BatchHeader)}
}

func (batch *BatchPPD) Validate() error {
	return batch.BatchPPD.Validate()
}

func (batch *BatchPPD) Create() error {
	return batch.BatchPPD.Create()
}

type BatchRCK struct{ ach.BatchRCK }

func NewBatchRCK(bh *BatchHeader) *BatchRCK {
	return &BatchRCK{*ach.NewBatchRCK(bh.BatchHeader)}
}

func (batch *BatchRCK) Validate() error {
	return batch.BatchRCK.Validate()
}

func (batch *BatchRCK) Create() error {
	return batch.BatchRCK.Create()
}

type BatchSHR struct{ ach.BatchSHR }

func NewBatchSHR(bh *BatchHeader) *BatchSHR {
	return &BatchSHR{*ach.NewBatchSHR(bh.BatchHeader)}
}

func (batch *BatchSHR) Validate() error {
	return batch.BatchSHR.Validate()
}

func (batch *BatchSHR) Create() error {
	return batch.BatchSHR.Create()
}

type BatchTEL struct{ ach.BatchTEL }

func NewBatchTEL(bh *BatchHeader) *BatchTEL {
	return &BatchTEL{*ach.NewBatchTEL(bh.BatchHeader)}
}

func (batch *BatchTEL) Validate() error {
	return batch.BatchTEL.Validate()
}

func (batch *BatchTEL) Create() error {
	return batch.BatchTEL.Create()
}

type BatchTRC struct{ ach.BatchTRC }

func NewBatchTRC(bh *BatchHeader) *BatchTRC {
	return &BatchTRC{*ach.NewBatchTRC(bh.BatchHeader)}
}

func (batch *BatchTRC) Validate() error {
	return batch.BatchTRC.Validate()
}

func (batch *BatchTRC) Create() error {
	return batch.BatchTRC.Create()
}

type BatchTRX struct{ ach.BatchTRX }

func NewBatchTRX(bh *BatchHeader) *BatchTRX {
	return &BatchTRX{*ach.NewBatchTRX(bh.BatchHeader)}
}

func (batch *BatchTRX) Validate() error {
	return batch.BatchTRX.Validate()
}

func (batch *BatchTRX) Create() error {
	return batch.BatchTRX.Create()
}

type BatchWEB struct{ ach.BatchWEB }

func NewBatchWEB(bh *BatchHeader) *BatchWEB {
	return &BatchWEB{*ach.NewBatchWEB(bh.BatchHeader)}
}

func (batch *BatchWEB) Validate() error {
	return batch.BatchWEB.Validate()
}

func (batch *BatchWEB) Create() error {
	return batch.BatchWEB.Create()
}

type OffsetAccountType ach.OffsetAccountType

const (
	OffsetChecking OffsetAccountType = "checking"
	OffsetSavings  OffsetAccountType = "savings"
)

func ReadDir(dir string) ([]*File, error) {
	achFiles, err := ach.ReadDir(dir)
	files := make([]*File, 0, len(achFiles))
	for _, v := range achFiles {
		files = append(files, &File{*v})
	}
	return files, err
}

const (
	CategoryForward                   = "Forward"
	CategoryReturn                    = "Return"
	CategoryNOC                       = "NOC"
	CategoryDishonoredReturn          = "DishonoredReturn"
	CategoryDishonoredReturnContested = "DishonoredReturnContested"

	CheckingCredit                     = 22
	CheckingReturnNOCCredit            = 21
	CheckingPrenoteCredit              = 23
	CheckingZeroDollarRemittanceCredit = 24
	CheckingDebit                      = 27
	CheckingReturnNOCDebit             = 26
	CheckingPrenoteDebit               = 28
	CheckingZeroDollarRemittanceDebit  = 29

	SavingsCredit                     = 32
	SavingsReturnNOCCredit            = 31
	SavingsPrenoteCredit              = 33
	SavingsZeroDollarRemittanceCredit = 34
	SavingsDebit                      = 37
	SavingsReturnNOCDebit             = 36
	SavingsPrenoteDebit               = 38
	SavingsZeroDollarRemittanceDebit  = 39

	GLCredit                     = 42
	GLReturnNOCCredit            = 41
	GLPrenoteCredit              = 43
	GLZeroDollarRemittanceCredit = 44
	GLDebit                      = 47
	GLReturnNOCDebit             = 46
	GLPrenoteDebit               = 48
	GLZeroDollarRemittanceDebit  = 49

	LoanCredit                     = 52
	LoanReturnNOCCredit            = 51
	LoanPrenoteCredit              = 53
	LoanZeroDollarRemittanceCredit = 54
	LoanDebit                      = 55
	LoanReturnNOCDebit             = 56
)

func NewEntryDetail() *EntryDetail {
	return &EntryDetail{*ach.NewEntryDetail()}
}

func (ed *EntryDetail) Parse(record string) {
	ed.EntryDetail.Parse(record)
}

func (ed *EntryDetail) String() string {
	return ed.EntryDetail.String()
}

func (ed *EntryDetail) SetValidation(opts *ValidateOpts) {
	ed.EntryDetail.SetValidation(opts.ValidateOpts)
}

func (ed *EntryDetail) SetRDFI(rdfi string) *EntryDetail {
	return &EntryDetail{*ed.EntryDetail.SetRDFI(rdfi)}
}

func (ed *EntryDetail) SetTraceNumber(ODFIIdentification string, seq int) {
	ed.EntryDetail.SetTraceNumber(ODFIIdentification, seq)
}

func (ed *EntryDetail) RDFIIdentificationField() string {
	return ed.EntryDetail.RDFIIdentificationField()
}

func (ed *EntryDetail) DFIAccountNumberField() string {
	return ed.EntryDetail.DFIAccountNumberField()
}

func (ed *EntryDetail) AmountField() string {
	return ed.EntryDetail.AmountField()
}

func (ed *EntryDetail) IdentificationNumberField() string {
	return ed.EntryDetail.IdentificationNumberField()
}

func (ed *EntryDetail) CheckSerialNumberField() string {
	return ed.EntryDetail.CheckSerialNumberField()
}

func (ed *EntryDetail) SetCheckSerialNumber(s string) {
	ed.EntryDetail.SetCheckSerialNumber(s)
}

func (ed *EntryDetail) SetPOPCheckSerialNumber(s string) {
	ed.EntryDetail.SetPOPCheckSerialNumber(s)
}

func (ed *EntryDetail) SetPOPTerminalCity(s string) {
	ed.EntryDetail.SetPOPTerminalCity(s)
}

func (ed *EntryDetail) SetPOPTerminalState(s string) {
	ed.EntryDetail.SetPOPTerminalState(s)
}

func (ed *EntryDetail) POPCheckSerialNumberField() string {
	return ed.EntryDetail.POPCheckSerialNumberField()
}

func (ed *EntryDetail) POPTerminalCityField() string {
	return ed.EntryDetail.POPTerminalCityField()
}

func (ed *EntryDetail) POPTerminalStateField() string {
	return ed.EntryDetail.POPTerminalStateField()
}

func (ed *EntryDetail) SetSHRCardExpirationDate(s string) {
	ed.EntryDetail.SetSHRCardExpirationDate(s)
}

func (ed *EntryDetail) SetSHRDocumentReferenceNumber(s string) {
	ed.EntryDetail.SetSHRDocumentReferenceNumber(s)
}

func (ed *EntryDetail) SetSHRIndividualCardAccountNumber(s string) {
	ed.EntryDetail.SetSHRDocumentReferenceNumber(s)
}

func (ed *EntryDetail) SHRCardExpirationDateField() string {
	return ed.EntryDetail.SHRCardExpirationDateField()
}

func (ed *EntryDetail) SHRDocumentReferenceNumberField() string {
	return ed.EntryDetail.SHRDocumentReferenceNumberField()
}

func (ed *EntryDetail) SHRIndividualCardAccountNumberField() string {
	return ed.EntryDetail.SHRIndividualCardAccountNumberField()
}

func (ed *EntryDetail) IndividualNameField() string {
	return ed.EntryDetail.IndividualNameField()
}

func (ed *EntryDetail) ReceivingCompanyField() string {
	return ed.EntryDetail.ReceivingCompanyField()
}

func (ed *EntryDetail) SetReceivingCompany(s string) {
	ed.EntryDetail.SetReceivingCompany(s)
}

func (ed *EntryDetail) OriginalTraceNumberField() string {
	return ed.EntryDetail.OriginalTraceNumberField()
}

func (ed *EntryDetail) SetOriginalTraceNumber(s string) {
	ed.EntryDetail.SetOriginalTraceNumber(s)
}

func (ed *EntryDetail) SetCATXAddendaRecords(i int) {
	ed.EntryDetail.SetCATXAddendaRecords(i)
}

func (ed *EntryDetail) SetCATXReceivingCompany(s string) {
	ed.EntryDetail.SetCATXReceivingCompany(s)
}

func (ed *EntryDetail) CATXAddendaRecordsField() string {
	return ed.EntryDetail.CATXAddendaRecordsField()
}

func (ed *EntryDetail) CATXReceivingCompanyField() string {
	return ed.EntryDetail.CATXReceivingCompanyField()
}

func (ed *EntryDetail) CATXReservedField() string {
	return ed.EntryDetail.CATXReservedField()
}

func (ed *EntryDetail) DiscretionaryDataField() string {
	return ed.EntryDetail.DiscretionaryDataField()
}

func (ed *EntryDetail) PaymentTypeField() string {
	return ed.EntryDetail.PaymentTypeField()
}

func (ed *EntryDetail) SetPaymentType(t string) {
	ed.EntryDetail.SetPaymentType(t)
}

func (ed *EntryDetail) SetProcessControlField(s string) {
	ed.EntryDetail.SetProcessControlField(s)
}

func (ed *EntryDetail) SetItemResearchNumber(s string) {
	ed.EntryDetail.SetItemResearchNumber(s)
}

func (ed *EntryDetail) SetItemTypeIndicator(s string) {
	ed.EntryDetail.SetItemTypeIndicator(s)
}

func (ed *EntryDetail) ProcessControlField() string {
	return ed.EntryDetail.ProcessControlField()
}

func (ed *EntryDetail) ItemResearchNumber() string {
	return ed.EntryDetail.ItemResearchNumber()
}

func (ed *EntryDetail) ItemTypeIndicator() string {
	return ed.EntryDetail.ItemTypeIndicator()
}

func (ed *EntryDetail) TraceNumberField() string {
	return ed.EntryDetail.TraceNumberField()
}

func (ed *EntryDetail) CreditOrDebit() string {
	return ed.EntryDetail.CreditOrDebit()
}

func (ed *EntryDetail) AddAddenda05(addenda05 *Addenda05) {
	ed.EntryDetail.AddAddenda05(addenda05.Addenda05)
}

var (
	ErrNonAlphanumeric      = errors.New("has non alphanumeric characters")
	ErrUpperAlpha           = errors.New("is not uppercase A-Z or 0-9")
	ErrFieldInclusion       = errors.New("is a mandatory field and has a default value")
	ErrConstructor          = errors.New("is a mandatory field and has a default value, did you use the constructor?")
	ErrFieldRequired        = errors.New("is a required field")
	ErrServiceClass         = errors.New("is an invalid Service Class Code")
	ErrSECCode              = errors.New("is an invalid Standard Entry Class Code")
	ErrOrigStatusCode       = errors.New("is an invalid Originator Status Code")
	ErrAddendaTypeCode      = errors.New("is an invalid Addenda Type Code")
	ErrTransactionCode      = errors.New("is an invalid Transaction Code")
	ErrIdentificationNumber = errors.New("is an invalid identification number")
	ErrCardTransactionType  = errors.New("is an invalid Card Transaction Type")
	ErrValidMonth           = errors.New("is an invalid month")
	ErrValidDay             = errors.New("is an invalid day")
	ErrValidYear            = errors.New("is an invalid year")
	ErrValidState           = errors.New("is an invalid US state or territory")
	ErrValidISO3166         = errors.New("is an invalid ISO 3166-1-alpha-2 code")
	ErrValidISO4217         = errors.New("is an invalid ISO 4217 code")

	ErrNegativeAmount = errors.New("amounts cannot be negative")

	ErrAddenda98ChangeCode                 = errors.New("found is not a valid addenda Change Code")
	ErrAddenda98RefusedChangeCode          = errors.New("found is not a valid addenda Refused Change Code")
	ErrAddenda98RefusedTraceSequenceNumber = errors.New("found is not a valid addenda trace sequence number")
	ErrAddenda98CorrectedData              = errors.New("must contain the corrected information corresponding to the Change Code")
	ErrAddenda99ReturnCode                 = errors.New("found is not a valid return code")
	ErrAddenda99DishonoredReturnCode       = errors.New("found is not a valid dishonored return code")
	ErrAddenda99ContestedReturnCode        = errors.New("found is not a valid contested dishonored return code")
	ErrBatchCORAddenda                     = errors.New("one Addenda98 or Addenda98Refused record is required for each entry in SEC Type COR")

	ErrRecordSize     = errors.New("is not 094")
	ErrBlockingFactor = errors.New("is not 10")
	ErrFormatCode     = errors.New("is not 1")

	ErrForeignExchangeIndicator          = errors.New("is an invalid Foreign Exchange Indicator")
	ErrForeignExchangeReferenceIndicator = errors.New("is an invalid Foreign Exchange Reference Indicator")
	ErrTransactionTypeCode               = errors.New("is an invalid Addenda10 Transaction Type Code")
	ErrIDNumberQualifier                 = errors.New("is an invalid Identification Number Qualifier")
	ErrIATBatchAddendaIndicator          = errors.New("is invalid for addenda record(s) found")
)

type FieldError struct{ ach.FieldError }

func (e *FieldError) Error() string {
	return e.FieldError.Error()
}

func (e *FieldError) Unwrap() error {
	return e.FieldError.Unwrap()
}

type ErrValidCheckDigit struct{ ach.ErrValidCheckDigit }

func NewErrValidCheckDigit(digit int) *ErrValidCheckDigit {
	return &ErrValidCheckDigit{ach.NewErrValidCheckDigit(digit)}
}

func (e ErrValidCheckDigit) Error() string {
	return e.ErrValidCheckDigit.Error()
}

type ErrValidFieldLength struct{ ach.ErrValidFieldLength }

func NewErrValidFieldLength(digit int) *ErrValidFieldLength {
	return &ErrValidFieldLength{ach.NewErrValidFieldLength(digit)}
}

func (e ErrValidFieldLength) Error() string {
	return e.ErrValidFieldLength.Error()
}

type ErrRecordType struct{ ach.ErrRecordType }

func NewErrRecordType(digit int) *ErrRecordType {
	return &ErrRecordType{ach.NewErrRecordType(digit)}
}

func (e ErrRecordType) Error() string {
	return e.ErrRecordType.Error()
}

type FileError struct{ ach.FileError }

func (e FileError) Error() string {
	return e.FileError.Error()
}

type File struct{ ach.File }

func FileFromJSON(bs []byte) (*File, error) {
	file, err := ach.FileFromJSON(bs)
	return &File{*file}, err
}

func ReadJSONFile(path string) (*File, error) {
	file, err := ach.ReadJSONFile(path)
	return &File{*file}, err
}

func ReadJSONFileWith(path string, opts *ValidateOpts) (*File, error) {
	file, err := ach.ReadJSONFileWith(path, opts.ValidateOpts)
	return &File{*file}, err
}

func FileFromJSONWith(bs []byte, opts *ValidateOpts) (*File, error) {
	file, err := ach.FileFromJSONWith(bs, opts.ValidateOpts)
	return &File{*file}, err
}

func (f *File) MarshalJSON() ([]byte, error) {
	return f.File.MarshalJSON()
}

func (f *File) UnmarshalJSON(p []byte) error {
	return f.File.UnmarshalJSON(p)
}

func (f *File) Create() error {
	return f.File.Create()
}

func (f *File) AddBatch(batch Batcher) []Batcher {
	var batchers []Batcher
	for _, b := range f.File.AddBatch(batch.Batcher) {
		batchers = append(batchers, Batcher{b})
	}
	return batchers
}

func (f *File) RemoveBatch(batch Batcher) {
	f.File.RemoveBatch(batch.Batcher)
}

func (f *File) AddIATBatch(iatBatch IATBatch) []IATBatch {
	batchers := make([]IATBatch, 0, len(f.File.AddIATBatch(iatBatch.IATBatch)))
	for _, b := range f.File.AddIATBatch(iatBatch.IATBatch) {
		batchers = append(batchers, IATBatch{b})
	}
	return batchers
}

func (f *File) SetValidation(opts *ValidateOpts) {
	f.File.SetValidation(opts.ValidateOpts)
}

func (f *File) ValidateWith(opts *ValidateOpts) error {
	return f.File.ValidateWith(opts.ValidateOpts)
}

func (f *File) IsADV() bool {
	return f.File.IsADV()
}

type SegmentFileResponse struct {
	CreditEntryFile  *File
	DebitEntryFile   *File
	SegmentFileError error
}

func (f *File) SegmentFile(config *ach.SegmentFileConfiguration) SegmentFileResponse {
	creditFile, debitFile, err := ach.NewFile().SegmentFile(config)

	return SegmentFileResponse{
		CreditEntryFile:  &File{*creditFile},
		DebitEntryFile:   &File{*debitFile},
		SegmentFileError: err,
	}
}

type FileControl struct{ ach.FileControl }

func (fc *FileControl) Parse(record string) {
	fc.FileControl.Parse(record)
}

func NewFileControl() FileControl {
	return FileControl{ach.NewFileControl()}
}

func (fc *FileControl) String() string {
	return fc.FileControl.String()
}

func (fc *FileControl) Validate() error {
	return fc.FileControl.Validate()
}

func (fc *FileControl) BatchCountField() string {
	return fc.FileControl.BatchCountField()
}

func (fc *FileControl) BlockCountField() string {
	return fc.FileControl.BatchCountField()
}

func (fc *FileControl) EntryAddendaCountField() string {
	return fc.FileControl.EntryAddendaCountField()
}

func (fc *FileControl) EntryHashField() string {
	return fc.FileControl.EntryHashField()
}

func (fc *FileControl) TotalDebitEntryDollarAmountInFileField() string {
	return fc.FileControl.TotalDebitEntryDollarAmountInFileField()
}

func (fc *FileControl) TotalCreditEntryDollarAmountInFileField() string {
	return fc.FileControl.TotalCreditEntryDollarAmountInFileField()
}

var (
	ErrFileTooLong                  = errors.New("file exceeds maximum possible number of lines")
	ErrFileHeader                   = errors.New("none or more than one file headers exists")
	ErrFileControl                  = errors.New("none or more than one file control exists")
	ErrFileEntryOutsideBatch        = errors.New("entry outside of batch")
	ErrFileAddendaOutsideBatch      = errors.New("addenda outside of batch")
	ErrFileAddendaOutsideEntry      = errors.New("addenda outside of entry")
	ErrFileBatchControlOutsideBatch = errors.New("batch control outside of batch")
	ErrFileConsecutiveBatchHeaders  = errors.New("consecutive Batch Headers in file")
	ErrFileADVOnly                  = errors.New("file can only have ADV Batches")
	ErrFileIATSEC                   = errors.New("IAT Standard Entry Class Code should use iatBatch")
	ErrFileNoBatches                = errors.New("must have []*Batches or []*IATBatches to be built")

	ErrInvalidJSON = errors.New("invalid JSON")
)

type RecordWrongLengthErr struct{ ach.RecordWrongLengthErr }

func NewRecordWrongLengthErr(length int) RecordWrongLengthErr {
	return RecordWrongLengthErr{ach.NewRecordWrongLengthErr(length)}
}

func (e RecordWrongLengthErr) Error() string {
	return e.RecordWrongLengthErr.Error()
}

type ErrUnknownRecordType struct{ ach.ErrUnknownRecordType }

func NewErrUnknownRecordType(recordType string) ErrUnknownRecordType {
	return ErrUnknownRecordType{ach.NewErrUnknownRecordType(recordType)}
}

func (e ErrUnknownRecordType) Error() string {
	return e.ErrUnknownRecordType.Error()
}

type ErrFileUnknownSEC struct{ ach.ErrFileUnknownSEC }

func NewErrFileUnknownSEC(secType string) ErrFileUnknownSEC {
	return ErrFileUnknownSEC{ach.NewErrFileUnknownSEC(secType)}
}

func (e ErrFileUnknownSEC) Error() string {
	return e.ErrFileUnknownSEC.Error()
}

type ErrFileCalculatedControlEquality struct {
	ach.ErrFileCalculatedControlEquality
}

func NewErrFileCalculatedControlEquality(field string, calculated int, control int) ErrFileCalculatedControlEquality {
	return ErrFileCalculatedControlEquality{ach.NewErrFileCalculatedControlEquality(field, calculated, control)}
}

func (e ErrFileCalculatedControlEquality) Error() string {
	return e.ErrFileCalculatedControlEquality.Error()
}

type ErrFileBatchNumberAscending struct {
	ach.ErrFileBatchNumberAscending
}

func NewErrFileBatchNumberAscending(previous, current int) ErrFileBatchNumberAscending {
	return ErrFileBatchNumberAscending{ach.NewErrFileBatchNumberAscending(previous, current)}
}

func (e ErrFileBatchNumberAscending) Error() string {
	return e.ErrFileBatchNumberAscending.Error()
}

type FileHeader struct{ ach.FileHeader }

func NewFileHeader() FileHeader {
	return FileHeader{ach.NewFileHeader()}
}

func (fh *FileHeader) Parse(record string) {
	fh.FileHeader.Parse(record)
}

func (fh *FileHeader) String() string {
	return fh.FileHeader.String()
}

func (fh *FileHeader) SetValidation(opts *ValidateOpts) {
	fh.FileHeader.SetValidation(opts.ValidateOpts)
}

func (fh *FileHeader) Validate() error {
	return fh.FileHeader.Validate()
}

func (fh *FileHeader) ValidateWith(opts *ValidateOpts) error {
	return fh.FileHeader.ValidateWith(opts.ValidateOpts)
}

func (fh *FileHeader) ImmediateDestinationField() string {
	return fh.FileHeader.ImmediateDestinationField()
}

func (fh *FileHeader) ImmediateOriginField() string {
	return fh.FileHeader.ImmediateOriginField()
}

func (fh *FileHeader) FileCreationTimeField() string {
	return fh.FileHeader.FileCreationTimeField()
}

func (fh *FileHeader) ImmediateDestinationNameField() string {
	return fh.FileHeader.ImmediateDestinationField()
}

func (fh *FileHeader) ImmediateOriginNameField() string {
	return fh.FileHeader.ImmediateOriginNameField()
}

func (fh *FileHeader) ReferenceCodeField() string {
	return fh.FileHeader.ReferenceCodeField()
}

func Flatten(originalFile *File) (*File, error) {
	file, err := ach.Flatten(&originalFile.File)
	return &File{*file}, err
}

func (f *File) FlattenBatches() (*File, error) {
	file, err := f.File.FlattenBatches()
	return &File{*file}, err
}

type IATBatchHeader struct{ ach.IATBatchHeader }

func NewIATBatch(bh *IATBatchHeader) IATBatch {
	return IATBatch{ach.NewIATBatch(&bh.IATBatchHeader)}
}

func (iatBatch *IATBatch) UnmarshalJSON(p []byte) error {
	return iatBatch.IATBatch.UnmarshalJSON(p)
}

func (iatBatch *IATBatch) SetHeader(batchHeader *IATBatchHeader) {
	iatBatch.IATBatch.SetHeader(&batchHeader.IATBatchHeader)
}

func (iatBatch *IATBatch) GetHeader() *IATBatchHeader {
	return &IATBatchHeader{*iatBatch.IATBatch.GetHeader()}
}

func (iatBatch *IATBatch) SetControl(batchControl *BatchControl) {
	iatBatch.IATBatch.SetControl(batchControl.BatchControl)
}

func (iatBatch *IATBatch) GetControl() *BatchControl {
	return &BatchControl{iatBatch.IATBatch.GetControl()}
}

func (iatBatch *IATBatch) GetEntries() []*IATEntryDetail {
	entries := make([]*IATEntryDetail, 0, len(iatBatch.IATBatch.GetEntries()))
	for _, v := range iatBatch.IATBatch.GetEntries() {
		entries = append(entries, &IATEntryDetail{*v})
	}
	return entries
}

type IATEntryDetail struct{ ach.IATEntryDetail }

func (iatBatch *IATBatch) AddEntry(entry *IATEntryDetail) {
	iatBatch.IATBatch.AddEntry(&entry.IATEntryDetail)
}

func (iatBatch *IATBatch) Category() string {
	return iatBatch.IATBatch.Category()
}

func (iatBatch *IATBatch) Create() error {
	return iatBatch.IATBatch.Create()
}

func (iatBatch *IATBatch) Validate() error {
	return iatBatch.IATBatch.Validate()
}

func (iatBatch *IATBatch) SetValidation(opts *ValidateOpts) {
	iatBatch.IATBatch.SetValidation(opts.ValidateOpts)
}

const (
	IATCOR = "IATCOR"
)

func NewIATBatchHeader() *IATBatchHeader {
	return &IATBatchHeader{*ach.NewIATBatchHeader()}
}

func (iatBh *IATBatchHeader) Parse(record string) {
	iatBh.IATBatchHeader.Parse(record)
}

func (iatBh *IATBatchHeader) String() string {
	return iatBh.IATBatchHeader.String()
}

func (iatBh *IATBatchHeader) Validate() error {
	return iatBh.IATBatchHeader.Validate()
}

func (iatBh *IATBatchHeader) IATIndicatorField() string {
	return iatBh.IATBatchHeader.IATIndicatorField()
}

func (iatBh *IATBatchHeader) ForeignExchangeIndicatorField() string {
	return iatBh.IATBatchHeader.ForeignExchangeIndicatorField()
}

func (iatBh *IATBatchHeader) ForeignExchangeReferenceIndicatorField() string {
	return iatBh.IATBatchHeader.ForeignExchangeReferenceIndicatorField()
}

func (iatBh *IATBatchHeader) ForeignExchangeReferenceField() string {
	return iatBh.IATBatchHeader.ForeignExchangeReferenceField()
}

func (iatBh *IATBatchHeader) ISODestinationCountryCodeField() string {
	return iatBh.IATBatchHeader.ISODestinationCountryCodeField()
}

func (iatBh *IATBatchHeader) OriginatorIdentificationField() string {
	return iatBh.IATBatchHeader.OriginatorIdentificationField()
}

func (iatBh *IATBatchHeader) CompanyEntryDescriptionField() string {
	return iatBh.IATBatchHeader.CompanyEntryDescriptionField()
}

func (iatBh *IATBatchHeader) ISOOriginatingCurrencyCodeField() string {
	return iatBh.IATBatchHeader.ISOOriginatingCurrencyCodeField()
}

func (iatBh *IATBatchHeader) ISODestinationCurrencyCodeField() string {
	return iatBh.IATBatchHeader.ISODestinationCurrencyCodeField()
}

func (iatBh *IATBatchHeader) EffectiveEntryDateField() string {
	return iatBh.IATBatchHeader.EffectiveEntryDateField()
}

func (iatBh *IATBatchHeader) ODFIIdentificationField() string {
	return iatBh.IATBatchHeader.ODFIIdentificationField()
}

func (iatBh *IATBatchHeader) BatchNumberField() string {
	return iatBh.IATBatchHeader.BatchNumberField()
}

func (iatBh *IATBatchHeader) SettlementDateField() string {
	return iatBh.IATBatchHeader.SettlementDateField()
}

func NewIATEntryDetail() *IATEntryDetail {
	return &IATEntryDetail{*ach.NewIATEntryDetail()}
}

func (iatEd *IATEntryDetail) Parse(record string) {
	iatEd.IATEntryDetail.Parse(record)
}

func (iatEd *IATEntryDetail) String() string {
	return iatEd.IATEntryDetail.String()
}

func (iatEd *IATEntryDetail) SetValidation(opts *ValidateOpts) {
	iatEd.IATEntryDetail.SetValidation(opts.ValidateOpts)
}

func (iatEd *IATEntryDetail) Validate() error {
	return iatEd.IATEntryDetail.Validate()
}

func (iatEd *IATEntryDetail) SetRDFI(rdfi string) *IATEntryDetail {
	return &IATEntryDetail{*iatEd.IATEntryDetail.SetRDFI(rdfi)}
}

func (iatEd *IATEntryDetail) SetTraceNumber(ODFIIdentification string, seq int) {
	iatEd.IATEntryDetail.SetTraceNumber(ODFIIdentification, seq)
}

func (iatEd *IATEntryDetail) RDFIIdentificationField() string {
	return iatEd.IATEntryDetail.RDFIIdentificationField()
}

func (iatEd *IATEntryDetail) AddendaRecordsField() string {
	return iatEd.IATEntryDetail.AddendaRecordsField()
}

func (iatEd *IATEntryDetail) AmountField() string {
	return iatEd.IATEntryDetail.AmountField()
}

func (iatEd *IATEntryDetail) DFIAccountNumberField() string {
	return iatEd.IATEntryDetail.DFIAccountNumberField()
}

func (iatEd *IATEntryDetail) OFACScreeningIndicatorField() string {
	return iatEd.IATEntryDetail.OFACScreeningIndicatorField()
}

func (iatEd *IATEntryDetail) SecondaryOFACScreeningIndicatorField() string {
	return iatEd.IATEntryDetail.SecondaryOFACScreeningIndicatorField()
}

func (iatEd *IATEntryDetail) TraceNumberField() string {
	return iatEd.IATEntryDetail.TraceNumberField()
}

func (iatEd *IATEntryDetail) AddAddenda17(addenda17 *Addenda17) {
	iatEd.IATEntryDetail.AddAddenda17(addenda17.Addenda17)
}

func (iatEd *IATEntryDetail) AddAddenda18(addenda18 *Addenda18) {
	iatEd.IATEntryDetail.AddAddenda18(addenda18.Addenda18)
}

type Iterator struct{ ach.Iterator }

func NewIterator(r io.Reader) *Iterator {
	return &Iterator{*ach.NewIterator(r)}
}

func (i *Iterator) SetValidation(opts *ValidateOpts) {
	i.Iterator.SetValidation(opts.ValidateOpts)
}

func (i *Iterator) GetHeader() *FileHeader {
	return &FileHeader{*i.Iterator.GetHeader()}
}

func (i *Iterator) GetControl() *FileControl {
	return &FileControl{*i.Iterator.GetControl()}
}

type NextEntryResponse struct {
	NextBatchHeader BatchHeader
	NextEntryDetail EntryDetail
	NextEntryError  error
}

func (i *Iterator) NextEntry(r io.Reader) NextEntryResponse {
	batchHeader, entryFile, err := ach.NewIterator(r).NextEntry()

	return NextEntryResponse{
		NextBatchHeader: BatchHeader{batchHeader},
		NextEntryDetail: EntryDetail{*entryFile},
		NextEntryError:  err,
	}
}

const NACHAFileLineLimit = 10000

type AchFiles []*ach.File

func MergeFiles(files []*File) ([]*File, error) {
	achFiles := make([]*ach.File, 0, len(files))
	for _, v := range files {
		achFiles = append(achFiles, (*ach.File)(&v.File))
	}
	res, err := ach.MergeFiles(achFiles)
	newFiles := make([]*File, 0, len(res))
	for _, v := range res {
		newFiles = append(newFiles, &File{*v})
	}
	return newFiles, err
}

type Merger ach.Merger

func NewMerger(opts *ValidateOpts) Merger {
	return ach.NewMerger(opts.ValidateOpts)
}

type Conditions struct{ ach.Conditions }

func MergeFilesWith(files []*File, conditions Conditions) ([]*File, error) {
	var achFiles []*ach.File
	for _, v := range files {
		achFiles = append(achFiles, (*ach.File)(&v.File))
	}
	res, err := ach.MergeFilesWith(achFiles, conditions.Conditions)
	var newFiles []*File
	for _, v := range res {
		newFiles = append(newFiles, &File{*v})
	}
	return newFiles, err
}

type Reader struct{ ach.Reader }

func (r *Reader) SetValidation(opts *ValidateOpts) {
	r.Reader.SetValidation(opts.ValidateOpts)
}

func ReadFile(path string) (*File, error) {
	file, err := ach.ReadFile(path)
	return &File{*file}, err
}

func ReadFiles(paths []string) ([]*File, error) {
	files := make([]*File, 0, len(paths))
	achFiles, err := ach.ReadFiles(paths)
	for _, v := range achFiles {
		files = append(files, &File{*v})
	}
	return files, err
}

func NewReader(r io.Reader) *Reader {
	return &Reader{*ach.NewReader(r)}
}

func (r *Reader) SetMaxLines(max int) {
	r.Reader.SetMaxLines(max)
}

func (r *Reader) Read() (File, error) {
	file, err := r.Reader.Read()
	return File{file}, err
}

func (f *File) Reversal(effectiveEntryDate time.Time) error {
	return f.File.Reversal(effectiveEntryDate)
}

type SegmentFileConfiguration struct{ ach.SegmentFileConfiguration }

func CalculateCheckDigit(routingNumber string) int {
	return ach.CalculateCheckDigit(routingNumber)
}

func CheckRoutingNumber(routingNumber string) error {
	return ach.CheckRoutingNumber(routingNumber)
}

type Writer struct{ ach.Writer }

func NewWriter(w io.Writer) *Writer {
	return &Writer{*ach.NewWriter(w)}
}

func (w *Writer) Write(file *File) error {
	return w.Writer.Write(&file.File)
}

func (w *Writer) Flush() error {
	return w.Writer.Flush()
}
