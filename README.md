moov-io/ach
===
[![GoDoc](https://godoc.org/github.com/moov-io/ach?status.svg)](https://godoc.org/github.com/moov-io/ach)
[![Build Status](https://travis-ci.com/moov-io/ach.svg?branch=master)](https://travis-ci.com/moov-io/ach)
[![Coverage Status](https://coveralls.io/repos/github/moov-io/ach/badge.svg?branch=master)](https://coveralls.io/github/moov-io/ach?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/moov-io/ach)](https://goreportcard.com/report/github.com/moov-io/ach)
[![Apache 2 licensed](https://img.shields.io/badge/license-Apache2-blue.svg)](https://raw.githubusercontent.com/moov-io/ach/master/LICENSE)
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fmoov-io%2Fach.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Fmoov-io%2Fach?ref=badge_shield)

Package `github.com/moov-io/ach` implements a file reader and writer for parsing [ACH](https://en.wikipedia.org/wiki/Automated_Clearing_House) Automated Clearing House files. ACH is the primary method of electronic money movement throughout the United States.

Docs: [docs.moov.io](http://docs.moov.io/en/latest/)

## Project Status

ACH is under active development but already in production for multiple companies. Please star the project if you are interested in its progress.

* Library currently supports the reading and writing
	* ARC (Accounts Receivable Entry)
	* BOC (Back Office Conversion)
	* CCD (Corporate credit or debit)
	* CIE (Customer-Initiated Entry)
	* COR (Automated Notification of Change(NOC))
	* CTX (Corporate Trade Exchange)
	* IAT (International ACH Transactions)
	* POP (Point of Purchase)
	* POS (Point of Sale)
	* PPD (Prearranged payment and deposits)
	* RCK (Represented Check Entries)
	* SHR (Shared Network Entry)
	* TEL (Telephone-Initiated Entry)
	* WEB (Internet-initiated Entries)
	* Return Entries
	* Addenda Type Code 02
	* Addenda Type Code 05
	* Addenda Type Code 10 (IAT)
	* Addenda Type Code 11 (IAT)
	* Addenda Type Code 12 (IAT)
	* Addenda Type Code 13 (IAT)
	* Addenda Type Code 14 (IAT)
	* Addenda Type Code 15 (IAT)
	* Addenda Type Code 16 (IAT)
	* Addenda Type Code 17 (IAT Optional)
	* Addenda Type Code 18 (IAT Optional)
	* Addenda Type Code 98 (NOC)
	* Addenda Type Code 99 (Return)

## Project Roadmap
* Additional SEC codes will be added based on library users needs. Please open an issue with a valid test file.
* Review the project issues for more detailed information

## Usage and tests
The following is a high level of reading and writing an ach file. Tests exist in projects [test](https://github.com/moov-io/ach/tree/master/test) folder with more details.

### Read a file

```go
// open a file for reading or pass any io.Reader NewReader()
f, err := os.Open("name-of-your-ach-file.ach")
if err != nil {
	log.Panicf("Can not open local file: %s: \n", err)
}
r := ach.NewReader(f)
achFile, err := r.Read()
if err != nil {
	fmt.Printf("Issue reading file: %+v \n", err)
}
// ensure we have a validated file structure
if achFile.Validate(); err != nil {
	fmt.Printf("Could not validate entire read file: %v", err)
}
// Check if any Notifications Of Change exist in the file
if len(achFile.NotificationOfChange) > 0 {
	for _, batch := range achFile.NotificationOfChange {
		a98 := batch.GetEntries()[0].Addendum[0].(*Addenda98)
		println(a98.CorrectedData)
	}
}
// Check if any Return Entries exist in the file
if len(achFile.ReturnEntries) > 0 {
	for _, batch := range achFile.ReturnEntries {
		aReturn := batch.GetEntries()[0].Addendum[0].(*Addenda99)
		println(aReturn.ReturnCode)
	}
}
```

### Create a file
The following is based on [simple file creation](https://github.com/moov-io/ach/tree/master/test/simple-file-creation)

 ```go
	fh := ach.NewFileHeader()
	fh.ImmediateDestination = "9876543210" // A blank space followed by your ODFI's transit/routing number
	fh.ImmediateOrigin = "1234567890"      // Organization or Company FED ID usually 1 and FEIN/SSN. Assigned by your ODFI
	fh.FileCreationDate = time.Now()     // Todays Date
	fh.ImmediateDestinationName = "Federal Reserve Bank"
	fh.ImmediateOriginName = "My Bank Name")

	file := ach.NewFile()
	file.SetHeader(fh)
```

Explicitly create a PPD batch file.

Errors only if payment type is not supported

 ```go
func mockBatchPPDHeader() *BatchHeader {
	bh := NewBatchHeader()
	bh.ServiceClassCode = 220
	bh.StandardEntryClassCode = "PPD"
	bh.CompanyName = "ACME Corporation"
	bh.CompanyIdentification = "123456789"
	bh.CompanyEntryDescription = "PAYROLL"
	bh.EffectiveEntryDate = time.Now()
	bh.ODFIIdentification = "6200001"
	return bh
}

mockBatch := NewBatch(mockBatchPPDHeader())
```

OR use the NewBatch factory

 ```go
func mockBatchPPDHeader() *BatchHeader {
	bh := NewBatchHeader()
	bh.ServiceClassCode = 220
	bh.StandardEntryClassCode = "PPD"
	bh.CompanyName = "ACME Corporation"
	bh.CompanyIdentification = "123456789"
	bh.CompanyEntryDescription = "PAYROLL"
	bh.EffectiveEntryDate = time.Now()
	bh.ODFIIdentification = "6200001"
	return bh
}

mockBatch, _ := ach.NewBatch(mockBatchPPDHeader())
```


To create an entry

 ```go
entry := ach.NewEntryDetail()
entry.TransactionCode = 22
entry.SetRDFI("009101298")
entry.DFIAccountNumber = "123456789"
entry.Amount = 100000000
entry.IndividualName = "Wade Arnold"
entry.SetTraceNumber(bh.ODFIIdentification, 1)
entry.IdentificationNumber = "ABC##jvkdjfuiwn"
entry.Category = ach.CategoryForward
```

To add one or more optional addenda records for an entry

 ```go
addenda := NewAddenda05()
addenda.PaymentRelatedInformation = "Bonus pay for amazing work on #OSS"
```
Add the addenda record to the detail entry

 ```go
entry.AddAddenda(addenda)
```

Entries are added to batches like so:

 ```go
batch.AddEntry(entry)
```

When all of the Entries are added to the batch we can create the batch.

 ```go
if err := batch.Create(); err != nil {
	fmt.Printf("%T: %s", err, err)
}
  ```

And batches are added to files much the same way:

 ```go
file.AddBatch(batch)
```

Now add a new batch for accepting payments on the WEB

```go
func mockBatchWEBHeader() *BatchHeader {
	bh := NewBatchHeader()
	bh.ServiceClassCode = 220
	bh.StandardEntryClassCode = "WEB"
	bh.CompanyName = "Your Company, inc"
	bh.CompanyIdentification = "123456789"
	bh.CompanyEntryDescription = "Online Order"
	bh.ODFIIdentification = "6200001"
	return bh
}

batch2, _ := ach.NewBatch(mockBatchWEBHeader())
```

Add an entry and define if it is a single or reoccurring payment. The following is a reoccurring payment for $7.99

```go
entry2 := ach.NewEntryDetail()
entry2.TransactionCode = 22
entry2.SetRDFI(102001017)
entry2.DFIAccountNumber = "5343121"
entry2.Amount = 799
entry2.IndividualName = "Wade Arnold"
entry2.SetTraceNumber(bh.ODFIIdentification, 1)
entry2.IdentificationNumber = "#123456"
entry.DiscretionaryData = "R"
entry2.Category = ach.CategoryForward


addenda2 := NewAddenda05()
addenda2.PaymentRelatedInformation = "Monthly Membership Subscription"
```

Add the entry to the batch
```go
entry2.AddAddenda(addenda2)
```

Create and add the second batch

```go
batch2.AddEntry(entry2)
if err := batch2.Create(); err != nil {
	fmt.Printf("%T: %s", err, err)
}
file.AddBatch(batch2)
```

Once we added all our batches we must build the file

 ```go
if err := file.Create(); err != nil {
	fmt.Printf("%T: %s", err, err)
}
```

Finally we want to write the file to an io.Writer

 ```go
w := ach.NewWriter(os.Stdout)
if err := w.Write(file); err != nil {
	fmt.Printf("%T: %s", err, err)
}
w.Flush()
}
```

Which will generate a well formed ACH flat file.

```text
101 210000890 1234567891708290000A094101Your Bank              Your Company           #00000A1
5200Your Company                        123456789 PPDTrans. DesOct 23010101   1234567890000001
6271020010175343121          0000017500#456789        Bob Smith             B11234567890000001
705bonus pay for amazing work on #OSS                                              00010000001
82000000020010200101000000017500000000000000123456789                          234567890000001
5220Your Company                        123456789 WEBsubscr    Oct 23010101   1234567890000002
6221020010175343121          0000000799#123456        Wade Arnold           R 1234567890000001
705Monthly Membership Subscription                                                 00010000001
82200000020010200101000000000000000000000799123456789                          234567890000002
9000002000001000000040020400202000000017500000000000799
```

### Create an IAT file

```go
    file := NewFile().SetHeader(mockFileHeader())
	iatBatch := IATBatch{}
	iatBatch.SetHeader(mockIATBatchHeaderFF())
	iatBatch.AddEntry(mockIATEntryDetail())
	iatBatch.Entries[0].Addenda10 = mockAddenda10()
	iatBatch.Entries[0].Addenda11 = mockAddenda11()
	iatBatch.Entries[0].Addenda12 = mockAddenda12()
	iatBatch.Entries[0].Addenda13 = mockAddenda13()
	iatBatch.Entries[0].Addenda14 = mockAddenda14()
	iatBatch.Entries[0].Addenda15 = mockAddenda15()
	iatBatch.Entries[0].Addenda16 = mockAddenda16()
	iatBatch.Entries[0].AddIATAddenda(mockAddenda17())
	iatBatch.Entries[0].AddIATAddenda(mockAddenda17B())
	iatBatch.Entries[0].AddIATAddenda(mockAddenda18())
	iatBatch.Entries[0].AddIATAddenda(mockAddenda18B())
	iatBatch.Entries[0].AddIATAddenda(mockAddenda18C())
	iatBatch.Entries[0].AddIATAddenda(mockAddenda18D())
	iatBatch.Entries[0].AddIATAddenda(mockAddenda18E())
	iatBatch.Create()
	file.AddIATBatch(iatBatch)

	iatBatch2 := IATBatch{}
	iatBatch2.SetHeader(mockIATBatchHeaderFF())
	iatBatch2.AddEntry(mockIATEntryDetail())
	iatBatch2.GetEntries()[0].TransactionCode = 27
	iatBatch2.GetEntries()[0].Amount = 2000
	iatBatch2.Entries[0].Addenda10 = mockAddenda10()
	iatBatch2.Entries[0].Addenda11 = mockAddenda11()
	iatBatch2.Entries[0].Addenda12 = mockAddenda12()
	iatBatch2.Entries[0].Addenda13 = mockAddenda13()
	iatBatch2.Entries[0].Addenda14 = mockAddenda14()
	iatBatch2.Entries[0].Addenda15 = mockAddenda15()
	iatBatch2.Entries[0].Addenda16 = mockAddenda16()
	iatBatch2.Entries[0].AddIATAddenda(mockAddenda17())
	iatBatch2.Entries[0].AddIATAddenda(mockAddenda17B())
	iatBatch2.Entries[0].AddIATAddenda(mockAddenda18())
	iatBatch2.Entries[0].AddIATAddenda(mockAddenda18B())
	iatBatch2.Entries[0].AddIATAddenda(mockAddenda18C())
	iatBatch2.Entries[0].AddIATAddenda(mockAddenda18D())
	iatBatch2.Entries[0].AddIATAddenda(mockAddenda18E())
	iatBatch2.Create()
	file.AddIATBatch(iatBatch2)

	if err := file.Create(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
	if err := file.Validate(); err != nil {
		t.Errorf("%T: %s", err, err)
	}

	b := &bytes.Buffer{}
	f := NewWriter(b)

	if err := f.Write(file); err != nil {
		t.Errorf("%T: %s", err, err)
	}

	r := NewReader(strings.NewReader(b.String()))
	_, err := r.Read()
	if err != nil {
		t.Errorf("%T: %s", err, err)
	}
	if err = r.File.Validate(); err != nil {
		t.Errorf("%T: %s", err, err)
	}

	// Write IAT records to standard output. Anything io.Writer
	w := NewWriter(os.Stdout)
	if err := w.Write(file); err != nil {
		log.Fatalf("Unexpected error: %s\n", err)
	}
	w.Flush()
```

This will generate a well formed flat IAT ACH file

```text
101 987654321 1234567891807200000A094101Federal Reserve Bank   My Bank Name
5220                FF3               US123456789 IATTRADEPAYMTCADUSD010101   0231380100000001
6221210428820007             0000100000123456789                              1231380100000001
710ANN000000000000100000928383-23938          BEK Enterprises                          0000001
711BEK Solutions                      15 West Place Street                             0000001
712JacobsTown*PA\                     US*19305\                                        0000001
713Wells Fargo                        01121042882                         US           0000001
714Citadel Bank                       01231380104                         US           0000001
7159874654932139872121 Front Street                                                    0000001
716LetterTown*AB\                     CA*80014\                                        0000001
717This is an international payment                                                00010000001
717Transfer of money from one country to another                                   00020000001
718Bank of Germany                    01987987987654654                   DE       00010000001
718Bank of Spain                      01987987987123123                   ES       00020000001
718Bank of France                     01456456456987987                   FR       00030000001
718Bank of Turkey                     0112312345678910                    TR       00040000001
718Bank of United Kingdom             011234567890123456789012345678901234GB       00050000001
82200000150012104288000000000000000000100000                                   231380100000001
5220                FF3               US123456789 IATTRADEPAYMTCADUSD010101   0231380100000002
6271210428820007             0000002000123456789                              1231380100000001
710ANN000000000000100000928383-23938          BEK Enterprises                          0000001
711BEK Solutions                      15 West Place Street                             0000001
712JacobsTown*PA\                     US*19305\                                        0000001
713Wells Fargo                        01121042882                         US           0000001
714Citadel Bank                       01231380104                         US           0000001
7159874654932139872121 Front Street                                                    0000001
716LetterTown*AB\                     CA*80014\                                        0000001
717This is an international payment                                                00010000001
717Transfer of money from one country to another                                   00020000001
718Bank of Germany                    01987987987654654                   DE       00010000001
718Bank of Spain                      01987987987123123                   ES       00020000001
718Bank of France                     01456456456987987                   FR       00030000001
718Bank of Turkey                     0112312345678910                    TR       00040000001
718Bank of United Kingdom             011234567890123456789012345678901234GB       00050000001
82200000150012104288000000002000000000000000                                   231380100000002
9000002000004000000300024208576000000002000000000100000
9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999
9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999
9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999
9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999
```

## Getting Help

 channel | info
 ------- | -------
 [Project Documentation](http://docs.moov.io/en/latest/) | Our project documentation available online.
 Google Group [moov-users](https://groups.google.com/forum/#!forum/moov-users)| The Moov users Google group is for contributors other people contributing to the Moov project. You can join them without a google account by sending an email to [moov-users+subscribe@googlegroups.com](mailto:moov-users+subscribe@googlegroups.com). After receiving the join-request message, you can simply reply to that to confirm the subscription.
Twitter [@moov_io](https://twitter.com/moov_io)	| You can follow Moov.IO's Twitter feed to get updates on our project(s). You can also tweet us questions or just share blogs or stories.
[GitHub Issue](https://github.com/moov-io) | If you are able to reproduce an problem please open a GitHub Issue under the specific project that caused the error.
[moov-io slack](http://moov-io.slack.com/) | Join our slack channel to have an interactive discussion about the development of the project. [Request and invote to the slack channel](https://join.slack.com/t/moov-io/shared_invite/enQtNDE5NzIwNTYxODEwLTRkYTcyZDI5ZTlkZWRjMzlhMWVhMGZlOTZiOTk4MmM3MmRhZDY4OTJiMDVjOTE2MGEyNWYzYzY1MGMyMThiZjg)

## Supported and Tested Platforms

- 64-bit Linux (Ubuntu, Debian), macOS, and Windows
- Rasberry Pi

Note: 32-bit platforms have known issues and is not supported.

## Contributing

Yes please! Please review our [Contributing guide](CONTRIBUTING.md) and [Code of Conduct](CODE_OF_CONDUCT.md) to get started!

Note: This project requires Go 1.10 or higher to compile.

## License

Apache License 2.0 See [LICENSE](LICENSE) for details.

[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fmoov-io%2Fach.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fmoov-io%2Fach?ref=badge_large)
