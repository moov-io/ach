moov-io/ach
===
[![GoDoc](https://godoc.org/github.com/moov-io/ach?status.svg)](https://godoc.org/github.com/moov-io/ach)
[![Build Status](https://travis-ci.org/moov-io/ach.svg?branch=master)](https://travis-ci.org/moov-io/ach)
[![Coverage Status](https://coveralls.io/repos/github/moov-io/ach/badge.svg?branch=master)](https://coveralls.io/github/moov-io/ach?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/moov-io/ach)](https://goreportcard.com/report/github.com/moov-io/ach)
[![Apache 2 licensed](https://img.shields.io/badge/license-Apache2-blue.svg)](https://raw.githubusercontent.com/moov-io/ach/master/LICENSE)


Package 'moov-io/ach' implements a file reader and writer for parsing [ACH](https://en.wikipedia.org/wiki/Automated_Clearing_House
) Automated Clearing House files. ACH is the primary method of electronic money movement throughout the United States.

## Project Status

ACH is under active development but already in production for multiple companies. Please star the project if you are interested in its progress.

* Library currently supports the reading and writing 
	* PPD (Prearranged payment and deposits)
	* WEB (Internet-initiated Entries )
	* CCD (Corporate credit or debit)
	* TEL (Telephone-Initiated Entry)
	* COR (Automated Notification of Change(NOC))
	* Return Entries


## Project Roadmap
* Additional SEC codes will be added based on library users needs. Please open an issue with a valid test file.
* Review the project issues for more detailed information

## Usage and examples
The following is a high level of reading and writing an ach file. Examples exist in projects [example](https://github.com/moov-io/ach/tree/master/example) folder with more details.

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
The following is based on [simple file creation](https://github.com/moov-io/ach/tree/master/example/simple-file-creation)
 
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
# Mailing lists

Users trade notes on the Google group moov-users (send mail to moov-users@googlegroups.com). You must join the [moov-users](https://groups.google.com/forum/#!forum/moov-users)forumn in order to post. 

# Contributing

We use GitHub to manage reviews of pull requests.

* If you have a trivial fix or improvement, go ahead and create a pull
  request, addressing (with `@...`) one or more of the maintainers
  (see [AUTHORS.md](AUTHORS.md)) in the description of the pull request.

* If you plan to do something more involved, first propose your ideas
  in a Github issue. This will avoid unnecessary work and surely give
  you and us a good deal of inspiration.

* Relevant coding style guidelines are the [Go Code Review
  Comments](https://code.google.com/p/go-wiki/wiki/CodeReviewComments)
  and the _Formatting and style_ section of Peter Bourgon's [Go: Best
  Practices for Production
  Environments](http://peter.bourgon.org/go-in-production/#formatting-and-style).

# Additional SEC (Standard Entry Class) code batch types. 
SEC type's in the Batch Header record define the payment type of the following Entry Details and Addenda. The format of the records in the batch is the same between all payment types but NACHA defines different rules for the values that are held in each record field. To add support for an additional SEC type you will need to implement NACHA rules for that type. The vast majority of rules are implemented in ach.batch and then composed into Batch(SEC) for reuse. All Batch(SEC) types must be a ach.Batcher.

1. Create a milestone for the new SEC type that you want supported. 
2. Add issues to that milestone to meet the NACHA rules for the batch type. 
3. Create a new struct of the batch type. In the following example we will use MTE(Machine Transfer Entry) as our example. 
4. The following code would be place in a new file batchMTE.go next to the existing batch types.
5. The code is stub code and the MTE type is not implemented. For concrete examples review the existing batch types in the source.  

Create a new struct and compose ach.batch 

```go
type BatchMTE struct {
	batch
}
```
Add the ability for the new type to be created. 

```go
func NewBatchMTE(bh *BatchHeader) *BatchMTE {
	batch := new(BatchMTE)
	batch.setControl(NewBatchControl)
	batch.SetHeader(bh)
	return batch
}
```

To support the Batcher interface you must add the following functions that are not implemented in ach.batch. 
* Validate() error
* Create() error 

Validate is designed to enforce the NACHA rules for the MTE payment type. Validate is run after a batch of this type is read from a file. If you are creating a batch from code call validate afterwards. 

```go
// Validate checks valid NACHA batch rules. Assumes properly parsed records.
func (batch *BatchMTE) Validate() error {
	// basic verification of the batch before we validate specific rules.
	if err := batch.verify(); err != nil {
		return err
	}
	// Add configuration based validation for this type.
	// ... batch.isAddendaCount(1)
	// Add type specific validation.
	// ...
	return nil
}
```
Create takes the Batch Header and Entry details and creates the proper sequence number and batch control. If additional logic specific to the SEC type is required it building a batch file it should be added here. 

```go
// Create takes Batch Header and Entries and builds a valid batch
func (batch *BatchMTE) Create() error {
	// generates sequence numbers and batch control
	if err := batch.build(); err != nil {
		return err
	}
	// Additional steps specific to batch type
	// ...

	if err := batch.Validate(); err != nil {
		return err
	}
	return nil
}
```

Finally add the batch type to the NewBatch factory in batch.go. 

```go
//...
case "MTE":
		return NewBatchMTE(bh), nil
//...
```

Pull request require a batchMTE_test.go file that covers the logic of the type. 

## References  
* [Wikipeda: Automated Clearing House](http://en.wikipedia.org/wiki/Automated_Clearing_House)
* [Nacha ACH Network: How it Works](https://www.nacha.org/ach-network)
* [Federal ACH Directory](https://www.frbservices.org/EPaymentsDirectory/search.html)

## Format Specification
* [NACHA ACH File Formatting](https://www.nacha.org/system/files/resources/AAP201%20-%20ACH%20File%20Formatting.pdf)
* [PNC ACH File Specification](http://content.pncmc.com/live/pnc/corporate/treasury-management/ach-conversion/ACH-File-Specifications.pdf)
* [Thomson Reuters ACH FIle Structure](http://cs.thomsonreuters.com/ua/acct_pr/acs/cs_us_en/pr/dd/ach_file_structure_and_content.htm)
* [Gusto: How ACH Works: A developer perspective](http://engineering.gusto.com/how-ach-works-a-developer-perspective-part-4/)

![ACH File Layout](https://github.com/moov-io/ach/blob/master/documentation/ach_file_structure_shg.gif)

## Insperation 
* [ACH:Builder - Tools for Building ACH](http://search.cpan.org/~tkeefer/ACH-Builder-0.03/lib/ACH/Builder.pm)
* [mosscode / ach](https://github.com/mosscode/ach)
* [Helper for building ACH files in Ruby](https://github.com/jm81/ach)
* [Glenselle / nACH2](https://github.com/glenselle/nACH2)

## License
Apache License 2.0 See [LICENSE](LICENSE) for details.
