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

ACH is at an early stage and under active development. Please star the project if you are interested in its progress.

* Library currently supports the reading and writing of PPD (Prearranged payment and deposits)


## Project Roadmap
* Support SEC WEB
* Additional SEC codes will be added based on library users needs. Please open an issue with a valid test file.

## Usage and examples
Examples exist in projects [example](https://github.com/moov-io/ach/tree/master/example) folder. The following is based on [simple file creation](https://github.com/moov-io/ach/tree/master/example/simple-file-creation)

 To create a file
 
 ```go
	file := ach.NewFile(ach.FileParam{
		ImmediateDestination:     "0210000890",
		ImmediateOrigin:          "123456789",
		ImmediateDestinationName: "Your Bank",
		ImmediateOriginName:      "Your Company",
		ReferenceCode:            "#00000A1"})
```

To create a batch

 ```go
	batch := ach.NewBatchPPD(ach.BatchParam{
		ServiceClassCode:        "220",
		CompanyName:             "Your Company",
		StandardEntryClass:      "PPD",
		CompanyIdentification:   "123456789",
		CompanyEntryDescription: "Trans. Description",
		CompanyDescriptiveDate:  "Oct 23",
		ODFIIdentification:      "123456789"})
```

To create an entry

 ```go
	entry := ach.NewEntryDetail(ach.EntryParam{
		ReceivingDFI:      "102001017",
		RDFIAccount:       "5343121",
		Amount:            "17500",
		TransactionCode:   "22",
		IDNumber:          "ABC##jvkdjfuiwn",
		IndividualName:    "Bob Smith",
		DiscretionaryData: "B1"})
```

To add one or more optional addenda records for an entry

 ```go
	addenda := ach.NewAddenda(ach.AddendaParam{
		PaymentRelatedInfo: "bonus pay for amazing work on #OSS"})
	entry.AddAddenda(addenda)
```

Entries are added to batches like so:

 ```go
	batch.AddEntry(entry)
```

When all of the Entries are added to the batch we can build the batch.

 ```go
	if err := batch.Build(); err != nil {
		fmt.Printf("%T: %s", err, err)
	}
  ```

And batches are added to files much the same way:

 ```go
	file.AddBatch(batch)
```

Once we added all our batches we must build the file

 ```go
	if err := file.Build(); err != nil {
		fmt.Printf("%T: %s", err, err)
	}
```

Finally we wnt to write the file to an io.Writer

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
101 210000890 1234567891708080000A094101Your Bank              Your Company           #00000A1
5220Your Company                        123456789 PPDTrans. DesOct 23010101   1234567890000001
6221020010175343121          0000017500ABC##jvkdjfuiwnBob Smith             B11234567890000001
705bonus pay for amazing work on #OSS                                              00010000001
82200000020010200101000000000000000000017500123456789                          234567890000001
9000001000001000000020010200101000000000000000000017500                                       
9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999
9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999
9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999
9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999
```

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
