// TODO(adam): doc header
package java

import (
       "github.com/moov-io/ach"
)

// Types copied from moov-io/ach
type ADVBatchControl ach.ADVBatchControl
type ADVEntryDetail ach.ADVEntryDetail
type ADVFileControl ach.ADVFileControl
type Addenda02 ach.Addenda02
type Addenda05 ach.Addenda05
type Addenda10 ach.Addenda10
type Addenda11 ach.Addenda11
type Addenda12 ach.Addenda12
type Addenda13 ach.Addenda13
type Addenda14 ach.Addenda14
type Addenda15 ach.Addenda15
type Addenda16 ach.Addenda16
type Addenda17 ach.Addenda17
type Addenda18 ach.Addenda18
type Addenda98 ach.Addenda98
type Addenda98Refused ach.Addenda98Refused
type Addenda99 ach.Addenda99
type Addenda99Contested ach.Addenda99Contested
type Addenda99Dishonored ach.Addenda99Dishonored
type Batch ach.Batch
type BatchACK ach.BatchACK
type BatchADV ach.BatchADV
type BatchARC ach.BatchARC
type BatchATX ach.BatchATX
type BatchBOC ach.BatchBOC
type BatchCCD ach.BatchCCD
type BatchCIE ach.BatchCIE
type BatchCOR ach.BatchCOR
type BatchCTX ach.BatchCTX
type BatchControl ach.BatchControl
type BatchDNE ach.BatchDNE
type BatchENR ach.BatchENR
type BatchError ach.BatchError
type BatchHeader ach.BatchHeader
type BatchMTE ach.BatchMTE
type BatchPOP ach.BatchPOP
type BatchPOS ach.BatchPOS
type BatchPPD ach.BatchPPD
type BatchRCK ach.BatchRCK
type BatchSHR ach.BatchSHR
type BatchTEL ach.BatchTEL
type BatchTRC ach.BatchTRC
type BatchTRX ach.BatchTRX
type BatchWEB ach.BatchWEB
type BatchXCK ach.BatchXCK
type ChangeCode ach.ChangeCode
type Conditions ach.Conditions
type CorrectedData ach.CorrectedData
type ENRPaymentInformation ach.ENRPaymentInformation
type EntryDetail ach.EntryDetail
type ErrBatchAddendaCount ach.ErrBatchAddendaCount
type ErrBatchAddendaTraceNumber ach.ErrBatchAddendaTraceNumber
type ErrBatchAmount ach.ErrBatchAmount
type ErrBatchAscending ach.ErrBatchAscending
type ErrBatchCalculatedControlEquality ach.ErrBatchCalculatedControlEquality
type ErrBatchCategory ach.ErrBatchCategory
type ErrBatchExpectedAddendaCount ach.ErrBatchExpectedAddendaCount
type ErrBatchHeaderControlEquality ach.ErrBatchHeaderControlEquality
type ErrBatchIATNOC ach.ErrBatchIATNOC
type ErrBatchRequiredAddendaCount ach.ErrBatchRequiredAddendaCount
type ErrBatchServiceClassTranCode ach.ErrBatchServiceClassTranCode
type ErrBatchTraceNumberNotODFI ach.ErrBatchTraceNumberNotODFI
type ErrFileBatchNumberAscending ach.ErrFileBatchNumberAscending
type ErrFileCalculatedControlEquality ach.ErrFileCalculatedControlEquality
type ErrFileUnknownSEC ach.ErrFileUnknownSEC
type ErrRecordType ach.ErrRecordType
type ErrUnknownRecordType ach.ErrUnknownRecordType
type ErrValidCheckDigit ach.ErrValidCheckDigit
type ErrValidFieldLength ach.ErrValidFieldLength
type FieldError ach.FieldError
type File ach.File
type FileControl ach.FileControl
type FileError ach.FileError
type FileHeader ach.FileHeader
type IATBatch ach.IATBatch
type IATBatchHeader ach.IATBatchHeader
type IATEntryDetail ach.IATEntryDetail
type Iterator ach.Iterator
type MergeDirOptions ach.MergeDirOptions
type Offset ach.Offset
type Reader ach.Reader
type RecordWrongLengthErr ach.RecordWrongLengthErr
type ReturnCode ach.ReturnCode
type SegmentFileConfiguration ach.SegmentFileConfiguration
type ValidateOpts ach.ValidateOpts
type Writer ach.Writer

// Functions copied from moov-io/ach

var CalculateCheckDigit = ach.CalculateCheckDigit
var CheckRoutingNumber = ach.CheckRoutingNumber
var ConvertBatchType = ach.ConvertBatchType
var DefaultFileAcceptor = ach.DefaultFileAcceptor
var FileFromJSON = ach.FileFromJSON
var FileFromJSONWith = ach.FileFromJSONWith
var Flatten = ach.Flatten
var IsContestedReturnCode = ach.IsContestedReturnCode
var IsDishonoredReturnCode = ach.IsDishonoredReturnCode
var IsRefusedChangeCode = ach.IsRefusedChangeCode
var LookupChangeCode = ach.LookupChangeCode
var LookupReturnCode = ach.LookupReturnCode
var MergeDir = ach.MergeDir
var MergeFiles = ach.MergeFiles
var MergeFilesWith = ach.MergeFilesWith
var NewADVBatchControl = ach.NewADVBatchControl
var NewADVEntryDetail = ach.NewADVEntryDetail
var NewADVFileControl = ach.NewADVFileControl
var NewAddenda02 = ach.NewAddenda02
var NewAddenda05 = ach.NewAddenda05
var NewAddenda10 = ach.NewAddenda10
var NewAddenda11 = ach.NewAddenda11
var NewAddenda12 = ach.NewAddenda12
var NewAddenda13 = ach.NewAddenda13
var NewAddenda14 = ach.NewAddenda14
var NewAddenda15 = ach.NewAddenda15
var NewAddenda16 = ach.NewAddenda16
var NewAddenda17 = ach.NewAddenda17
var NewAddenda18 = ach.NewAddenda18
var NewAddenda98 = ach.NewAddenda98
var NewAddenda98Refused = ach.NewAddenda98Refused
var NewAddenda99 = ach.NewAddenda99
var NewAddenda99Contested = ach.NewAddenda99Contested
var NewAddenda99Dishonored = ach.NewAddenda99Dishonored
var NewBatch = ach.NewBatch
var NewBatchACK = ach.NewBatchACK
var NewBatchADV = ach.NewBatchADV
var NewBatchARC = ach.NewBatchARC
var NewBatchATX = ach.NewBatchATX
var NewBatchBOC = ach.NewBatchBOC
var NewBatchCCD = ach.NewBatchCCD
var NewBatchCIE = ach.NewBatchCIE
var NewBatchCOR = ach.NewBatchCOR
var NewBatchCTX = ach.NewBatchCTX
var NewBatchControl = ach.NewBatchControl
var NewBatchDNE = ach.NewBatchDNE
var NewBatchENR = ach.NewBatchENR
var NewBatchHeader = ach.NewBatchHeader
var NewBatchMTE = ach.NewBatchMTE
var NewBatchPOP = ach.NewBatchPOP
var NewBatchPOS = ach.NewBatchPOS
var NewBatchPPD = ach.NewBatchPPD
var NewBatchRCK = ach.NewBatchRCK
var NewBatchSHR = ach.NewBatchSHR
var NewBatchTEL = ach.NewBatchTEL
var NewBatchTRC = ach.NewBatchTRC
var NewBatchTRX = ach.NewBatchTRX
var NewBatchWEB = ach.NewBatchWEB
var NewBatchXCK = ach.NewBatchXCK
var NewEntryDetail = ach.NewEntryDetail
var NewErrBatchAddendaCount = ach.NewErrBatchAddendaCount
var NewErrBatchAddendaTraceNumber = ach.NewErrBatchAddendaTraceNumber
var NewErrBatchAmount = ach.NewErrBatchAmount
var NewErrBatchAscending = ach.NewErrBatchAscending
var NewErrBatchCalculatedControlEquality = ach.NewErrBatchCalculatedControlEquality
var NewErrBatchCategory = ach.NewErrBatchCategory
var NewErrBatchExpectedAddendaCount = ach.NewErrBatchExpectedAddendaCount
var NewErrBatchHeaderControlEquality = ach.NewErrBatchHeaderControlEquality
var NewErrBatchIATNOC = ach.NewErrBatchIATNOC
var NewErrBatchRequiredAddendaCount = ach.NewErrBatchRequiredAddendaCount
var NewErrBatchServiceClassTranCode = ach.NewErrBatchServiceClassTranCode
var NewErrBatchTraceNumberNotODFI = ach.NewErrBatchTraceNumberNotODFI
var NewErrFileBatchNumberAscending = ach.NewErrFileBatchNumberAscending
var NewErrFileCalculatedControlEquality = ach.NewErrFileCalculatedControlEquality
var NewErrFileUnknownSEC = ach.NewErrFileUnknownSEC
var NewErrRecordType = ach.NewErrRecordType
var NewErrUnknownRecordType = ach.NewErrUnknownRecordType
var NewErrValidCheckDigit = ach.NewErrValidCheckDigit
var NewErrValidFieldLength = ach.NewErrValidFieldLength
var NewFile = ach.NewFile
var NewFileControl = ach.NewFileControl
var NewFileHeader = ach.NewFileHeader
var NewIATBatch = ach.NewIATBatch
var NewIATBatchHeader = ach.NewIATBatchHeader
var NewIATEntryDetail = ach.NewIATEntryDetail
var NewIterator = ach.NewIterator
var NewMerger = ach.NewMerger
var NewReader = ach.NewReader
var NewRecordWrongLengthErr = ach.NewRecordWrongLengthErr
var NewSegmentFileConfiguration = ach.NewSegmentFileConfiguration
var NewWriter = ach.NewWriter
var ReadDir = ach.ReadDir
var ReadFile = ach.ReadFile
var ReadFiles = ach.ReadFiles
var ReadJSONFile = ach.ReadJSONFile
var ReadJSONFileWith = ach.ReadJSONFileWith
var StandardTransactionCode = ach.StandardTransactionCode
var WriteCorrectionData = ach.WriteCorrectionData
