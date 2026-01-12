# achcli

achcli is a command-line tool for displaying and manipulating Nacha-formatted ACH files in a human-readable format.
It is part of the [moov-io/ach](https://github.com/moov-io/ach) project, which provides Go libraries for creating,
parsing, and validating ACH files.

## Features

- Describe ACH files in a readable format.
- Mask sensitive information (e.g., account numbers, names).
- Compare (diff) two ACH files.
- Reformat ACH files to other formats (e.g., JSON).
- Merge multiple ACH files.
- Flatten batches in ACH files.
- Validate ACH files with custom options.
- Fix certain fields in ACH files (e.g., update Effective Entry Date).
- Pretty-print amounts and other values for better readability.

## Installation

achcli is built using Go. You can install it directly from the source:

```bash
go install github.com/moov-io/ach/cmd/achcli@latest
```

Ensure your Go environment is set up (Go 1.21+ recommended).

## Usage

```
achcli is a tool for displaying Nacha formatted ACH files in a human readable format.

USAGE
   achcli [-mask] [-pretty] [-validate opts.json] path/to/file.ach

EXAMPLES
  achcli -diff first.ach second.ach    Show the difference between two ACH files
  achcli -mask file.ach                Print file details with personally identifiable information partially removed
  achcli -reformat=json first.ach      Convert an incoming ACH file into another format (options: ach, json)
  achcli -validate opts.json file.ach  Read an ACH File with the provided ValidateOpts
  achcli -version                      Print the version of achcli (Example: v1.34.0)
  achcli 20060102.ach                  Summarize an ACH file for human readability

FLAGS
  -diff                        Compare two files against each other
  -fix                         Trigger fix tasks
  -flatten                     Flatten batches in each file
  -mask                        Mask/hide full account numbers and individual names
  -mask.accounts               Mask/hide full account numbers
  -mask.corrections            Mask/Hide Corrected Data in Addenda98 records
  -mask.names                  Mask/hide full individual names
  -merge                       Merge files before describing
  -pretty                      Display all values in their human readable format
  -pretty.amounts              Display human readable amounts instead of exact values
  -reformat string             Reformat an incoming ACH file to another format
  -skip-validation             Skip all validation checks
  -update-eed string           Set the EffectiveEntryDate to a new value
  -v                           Print verbose details about each ACH file
  -validate string             Path to config file in json format to enable validation opts
  -version                     Print moov-io/ach cli version
```

### Validation Options (-validate)

Provide a JSON file with [validation options](https://moov-io.github.io/ach/create-file/#validate-options) for custom ACH validation. Example `opts.json`:

```json
{
  "RequireABAOrigin": true,
  "BypassOriginValidation": false,
  "BypassDestinationValidation": false,
  "CustomTraceNumbers": true,
  "CustomReturnCodes": true,
  "AllowZeroBatches": true,
  "AllowZeroEntries": true,
  "AllowMissingFileHeader": true,
  "AllowMissingFileControl": true,
  "SkipAll": false
}
```

Usage: `achcli -validate opts.json file.ach`

### Fixing Files (-fix)

Use `-fix` to modify ACH files. Currently supports updating the Effective Entry Date:

```bash
achcli -fix -update-eed=20260102 input.ach
```

This will create a fixed file and output its path.

## Examples

### Describe a File

```bash
achcli ppd-debit.ach
```

Output:
```
Describing ACH file 'ppd-debit.ach'

  Origin      OriginName               Destination  DestinationName          FileCreationDate  FileCreationTime
   121042882  My Bank Name              231380104   Federal Reserve Bank     190624            0000

  BatchNumber  SECCode  ServiceClassCode    CompanyName       DiscretionaryData     Identification  EntryDescription  EffectiveEntryDate  DescriptiveDate
  0000001      PPD      225 (Debits  Only)  Name on Account                         121042882       REG.SALARY        190625

    TransactionCode       RDFIIdentification  AccountNumber      Amount     Name                    IdentificationNumber  TraceNumber      Category
    27 (Checking Debit)   23138010            12345678           100000000  Receiver Account Name                         121042880000001  Forward

  ServiceClassCode    EntryAddendaCount  EntryHash   TotalDebits  TotalCredits  MACCode              ODFIIdentification  BatchNumber
  225 (Debits  Only)  000001             0023138010  100000000    0                                  12104288            0000001

  BatchCount  BlockCount  EntryAddendaCount  TotalDebitAmount  TotalCreditAmount
  000001      000001      00000001           100000000         0
```

### Mask Sensitive Data

```bash
achcli -mask example.ach
```

Masks account numbers, names, etc.

### Diff Two Files

```bash
achcli -diff file1.ach file2.ach
```

Shows differences in headers, batches, etc., using colored output.

### Reformat to JSON

```bash
achcli -reformat=json input.ach > output.json
```

### Merge and Flatten Files

```bash
achcli -merge -flatten file1.ach file2.ach
```

Merges files and flattens batches before describing.

## License

Copyright © 2020 The Moov Authors

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License. You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.

## Contributing

achcli is part of the moov-io/ach repository. Contributions are welcome! See the [main repository](https://github.com/moov-io/ach) for details.
