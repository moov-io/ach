window.BENCHMARK_DATA = {
  "lastUpdate": 1727903599479,
  "repoUrl": "https://github.com/moov-io/ach",
  "entries": {
    "My Project Go Benchmark": [
      {
        "commit": {
          "author": {
            "name": "Adam Shannon",
            "username": "adamdecaf",
            "email": "adamkshannon@gmail.com"
          },
          "committer": {
            "name": "Adam Shannon",
            "username": "adamdecaf",
            "email": "adamkshannon@gmail.com"
          },
          "id": "c2fe468fbac7362e4f77093b9e224777bc13d7d3",
          "message": "build: extend benchmark timeout, skip fetching branch for results",
          "timestamp": "2024-10-02T20:51:45Z",
          "url": "https://github.com/moov-io/ach/commit/c2fe468fbac7362e4f77093b9e224777bc13d7d3"
        },
        "date": 1727903599412,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkAddenda02ValidTypeCode",
            "value": 93.69,
            "unit": "ns/op\t      80 B/op\t       2 allocs/op",
            "extra": "11388410 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda02ValidTypeCode - ns/op",
            "value": 93.69,
            "unit": "ns/op",
            "extra": "11388410 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda02ValidTypeCode - B/op",
            "value": 80,
            "unit": "B/op",
            "extra": "11388410 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda02ValidTypeCode - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "11388410 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda02TypeCode02",
            "value": 97.7,
            "unit": "ns/op\t      80 B/op\t       2 allocs/op",
            "extra": "12408441 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda02TypeCode02 - ns/op",
            "value": 97.7,
            "unit": "ns/op",
            "extra": "12408441 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda02TypeCode02 - B/op",
            "value": 80,
            "unit": "B/op",
            "extra": "12408441 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda02TypeCode02 - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "12408441 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda02FieldInclusionTypeCode",
            "value": 70.86,
            "unit": "ns/op\t      64 B/op\t       1 allocs/op",
            "extra": "16664527 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda02FieldInclusionTypeCode - ns/op",
            "value": 70.86,
            "unit": "ns/op",
            "extra": "16664527 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda02FieldInclusionTypeCode - B/op",
            "value": 64,
            "unit": "B/op",
            "extra": "16664527 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda02FieldInclusionTypeCode - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "16664527 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda02TransactionSerialNumber",
            "value": 65.78,
            "unit": "ns/op\t      64 B/op\t       1 allocs/op",
            "extra": "18102366 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda02TransactionSerialNumber - ns/op",
            "value": 65.78,
            "unit": "ns/op",
            "extra": "18102366 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda02TransactionSerialNumber - B/op",
            "value": 64,
            "unit": "B/op",
            "extra": "18102366 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda02TransactionSerialNumber - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "18102366 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda02TransactionDate",
            "value": 65.98,
            "unit": "ns/op\t      64 B/op\t       1 allocs/op",
            "extra": "17930242 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda02TransactionDate - ns/op",
            "value": 65.98,
            "unit": "ns/op",
            "extra": "17930242 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda02TransactionDate - B/op",
            "value": 64,
            "unit": "B/op",
            "extra": "17930242 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda02TransactionDate - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "17930242 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda02TerminalLocation",
            "value": 66.18,
            "unit": "ns/op\t      64 B/op\t       1 allocs/op",
            "extra": "17918510 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda02TerminalLocation - ns/op",
            "value": 66.18,
            "unit": "ns/op",
            "extra": "17918510 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda02TerminalLocation - B/op",
            "value": 64,
            "unit": "B/op",
            "extra": "17918510 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda02TerminalLocation - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "17918510 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda02TerminalCity",
            "value": 68.49,
            "unit": "ns/op\t      64 B/op\t       1 allocs/op",
            "extra": "17415121 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda02TerminalCity - ns/op",
            "value": 68.49,
            "unit": "ns/op",
            "extra": "17415121 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda02TerminalCity - B/op",
            "value": 64,
            "unit": "B/op",
            "extra": "17415121 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda02TerminalCity - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "17415121 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda02TerminalState",
            "value": 65.76,
            "unit": "ns/op\t      64 B/op\t       1 allocs/op",
            "extra": "17838172 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda02TerminalState - ns/op",
            "value": 65.76,
            "unit": "ns/op",
            "extra": "17838172 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda02TerminalState - B/op",
            "value": 64,
            "unit": "B/op",
            "extra": "17838172 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda02TerminalState - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "17838172 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda02String",
            "value": 1309,
            "unit": "ns/op\t     256 B/op\t      14 allocs/op",
            "extra": "887920 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda02String - ns/op",
            "value": 1309,
            "unit": "ns/op",
            "extra": "887920 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda02String - B/op",
            "value": 256,
            "unit": "B/op",
            "extra": "887920 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda02String - allocs/op",
            "value": 14,
            "unit": "allocs/op",
            "extra": "887920 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda02TransactionDateMonth",
            "value": 197.3,
            "unit": "ns/op\t      80 B/op\t       2 allocs/op",
            "extra": "6048253 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda02TransactionDateMonth - ns/op",
            "value": 197.3,
            "unit": "ns/op",
            "extra": "6048253 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda02TransactionDateMonth - B/op",
            "value": 80,
            "unit": "B/op",
            "extra": "6048253 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda02TransactionDateMonth - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "6048253 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda02TransactionDateDay",
            "value": 147.4,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "8306462 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda02TransactionDateDay - ns/op",
            "value": 147.4,
            "unit": "ns/op",
            "extra": "8306462 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda02TransactionDateDay - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "8306462 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda02TransactionDateDay - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "8306462 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda02TransactionDateFeb",
            "value": 190.3,
            "unit": "ns/op\t      80 B/op\t       2 allocs/op",
            "extra": "6267718 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda02TransactionDateFeb - ns/op",
            "value": 190.3,
            "unit": "ns/op",
            "extra": "6267718 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda02TransactionDateFeb - B/op",
            "value": 80,
            "unit": "B/op",
            "extra": "6267718 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda02TransactionDateFeb - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "6267718 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda02TransactionDate30Day",
            "value": 146.4,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "8268751 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda02TransactionDate30Day - ns/op",
            "value": 146.4,
            "unit": "ns/op",
            "extra": "8268751 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda02TransactionDate30Day - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "8268751 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda02TransactionDate30Day - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "8268751 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda02TransactionDate31Day",
            "value": 145.8,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "8302448 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda02TransactionDate31Day - ns/op",
            "value": 145.8,
            "unit": "ns/op",
            "extra": "8302448 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda02TransactionDate31Day - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "8302448 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda02TransactionDate31Day - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "8302448 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda02TransactionDateInvalidDay",
            "value": 209.6,
            "unit": "ns/op\t      80 B/op\t       2 allocs/op",
            "extra": "5752274 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda02TransactionDateInvalidDay - ns/op",
            "value": 209.6,
            "unit": "ns/op",
            "extra": "5752274 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda02TransactionDateInvalidDay - B/op",
            "value": 80,
            "unit": "B/op",
            "extra": "5752274 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda02TransactionDateInvalidDay - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "5752274 times\n4 procs"
          },
          {
            "name": "BenchmarkReferenceInformationOneAlphaNumeric",
            "value": 326.1,
            "unit": "ns/op\t     160 B/op\t       4 allocs/op",
            "extra": "3658160 times\n4 procs"
          },
          {
            "name": "BenchmarkReferenceInformationOneAlphaNumeric - ns/op",
            "value": 326.1,
            "unit": "ns/op",
            "extra": "3658160 times\n4 procs"
          },
          {
            "name": "BenchmarkReferenceInformationOneAlphaNumeric - B/op",
            "value": 160,
            "unit": "B/op",
            "extra": "3658160 times\n4 procs"
          },
          {
            "name": "BenchmarkReferenceInformationOneAlphaNumeric - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "3658160 times\n4 procs"
          },
          {
            "name": "BenchmarkReferenceInformationTwoAlphaNumeric",
            "value": 341,
            "unit": "ns/op\t     160 B/op\t       4 allocs/op",
            "extra": "3532485 times\n4 procs"
          },
          {
            "name": "BenchmarkReferenceInformationTwoAlphaNumeric - ns/op",
            "value": 341,
            "unit": "ns/op",
            "extra": "3532485 times\n4 procs"
          },
          {
            "name": "BenchmarkReferenceInformationTwoAlphaNumeric - B/op",
            "value": 160,
            "unit": "B/op",
            "extra": "3532485 times\n4 procs"
          },
          {
            "name": "BenchmarkReferenceInformationTwoAlphaNumeric - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "3532485 times\n4 procs"
          },
          {
            "name": "BenchmarkTerminalIdentificationCodeAlphaNumeric",
            "value": 346.4,
            "unit": "ns/op\t     160 B/op\t       4 allocs/op",
            "extra": "3493114 times\n4 procs"
          },
          {
            "name": "BenchmarkTerminalIdentificationCodeAlphaNumeric - ns/op",
            "value": 346.4,
            "unit": "ns/op",
            "extra": "3493114 times\n4 procs"
          },
          {
            "name": "BenchmarkTerminalIdentificationCodeAlphaNumeric - B/op",
            "value": 160,
            "unit": "B/op",
            "extra": "3493114 times\n4 procs"
          },
          {
            "name": "BenchmarkTerminalIdentificationCodeAlphaNumeric - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "3493114 times\n4 procs"
          },
          {
            "name": "BenchmarkTransactionSerialNumberAlphaNumeric",
            "value": 347.1,
            "unit": "ns/op\t     160 B/op\t       4 allocs/op",
            "extra": "3451208 times\n4 procs"
          },
          {
            "name": "BenchmarkTransactionSerialNumberAlphaNumeric - ns/op",
            "value": 347.1,
            "unit": "ns/op",
            "extra": "3451208 times\n4 procs"
          },
          {
            "name": "BenchmarkTransactionSerialNumberAlphaNumeric - B/op",
            "value": 160,
            "unit": "B/op",
            "extra": "3451208 times\n4 procs"
          },
          {
            "name": "BenchmarkTransactionSerialNumberAlphaNumeric - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "3451208 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationCodeOrExpireDateAlphaNumeric",
            "value": 431.3,
            "unit": "ns/op\t     160 B/op\t       4 allocs/op",
            "extra": "2768270 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationCodeOrExpireDateAlphaNumeric - ns/op",
            "value": 431.3,
            "unit": "ns/op",
            "extra": "2768270 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationCodeOrExpireDateAlphaNumeric - B/op",
            "value": 160,
            "unit": "B/op",
            "extra": "2768270 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationCodeOrExpireDateAlphaNumeric - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "2768270 times\n4 procs"
          },
          {
            "name": "BenchmarkTerminalLocationAlphaNumeric",
            "value": 455,
            "unit": "ns/op\t     160 B/op\t       4 allocs/op",
            "extra": "2658763 times\n4 procs"
          },
          {
            "name": "BenchmarkTerminalLocationAlphaNumeric - ns/op",
            "value": 455,
            "unit": "ns/op",
            "extra": "2658763 times\n4 procs"
          },
          {
            "name": "BenchmarkTerminalLocationAlphaNumeric - B/op",
            "value": 160,
            "unit": "B/op",
            "extra": "2658763 times\n4 procs"
          },
          {
            "name": "BenchmarkTerminalLocationAlphaNumeric - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "2658763 times\n4 procs"
          },
          {
            "name": "BenchmarkTerminalCityAlphaNumeric",
            "value": 469.4,
            "unit": "ns/op\t     160 B/op\t       4 allocs/op",
            "extra": "2475808 times\n4 procs"
          },
          {
            "name": "BenchmarkTerminalCityAlphaNumeric - ns/op",
            "value": 469.4,
            "unit": "ns/op",
            "extra": "2475808 times\n4 procs"
          },
          {
            "name": "BenchmarkTerminalCityAlphaNumeric - B/op",
            "value": 160,
            "unit": "B/op",
            "extra": "2475808 times\n4 procs"
          },
          {
            "name": "BenchmarkTerminalCityAlphaNumeric - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "2475808 times\n4 procs"
          },
          {
            "name": "BenchmarkTerminalStateAlphaNumeric",
            "value": 491.4,
            "unit": "ns/op\t     160 B/op\t       4 allocs/op",
            "extra": "2436620 times\n4 procs"
          },
          {
            "name": "BenchmarkTerminalStateAlphaNumeric - ns/op",
            "value": 491.4,
            "unit": "ns/op",
            "extra": "2436620 times\n4 procs"
          },
          {
            "name": "BenchmarkTerminalStateAlphaNumeric - B/op",
            "value": 160,
            "unit": "B/op",
            "extra": "2436620 times\n4 procs"
          },
          {
            "name": "BenchmarkTerminalStateAlphaNumeric - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "2436620 times\n4 procs"
          },
          {
            "name": "BenchmarkParseAddenda05",
            "value": 6048,
            "unit": "ns/op\t   15916 B/op\t      31 allocs/op",
            "extra": "195538 times\n4 procs"
          },
          {
            "name": "BenchmarkParseAddenda05 - ns/op",
            "value": 6048,
            "unit": "ns/op",
            "extra": "195538 times\n4 procs"
          },
          {
            "name": "BenchmarkParseAddenda05 - B/op",
            "value": 15916,
            "unit": "B/op",
            "extra": "195538 times\n4 procs"
          },
          {
            "name": "BenchmarkParseAddenda05 - allocs/op",
            "value": 31,
            "unit": "allocs/op",
            "extra": "195538 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda05String",
            "value": 939.6,
            "unit": "ns/op\t     272 B/op\t       6 allocs/op",
            "extra": "1274988 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda05String - ns/op",
            "value": 939.6,
            "unit": "ns/op",
            "extra": "1274988 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda05String - B/op",
            "value": 272,
            "unit": "B/op",
            "extra": "1274988 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda05String - allocs/op",
            "value": 6,
            "unit": "allocs/op",
            "extra": "1274988 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda05TypeCodeNil",
            "value": 62.8,
            "unit": "ns/op\t      64 B/op\t       1 allocs/op",
            "extra": "18688287 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda05TypeCodeNil - ns/op",
            "value": 62.8,
            "unit": "ns/op",
            "extra": "18688287 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda05TypeCodeNil - B/op",
            "value": 64,
            "unit": "B/op",
            "extra": "18688287 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda05TypeCodeNil - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "18688287 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda05TypeCode05",
            "value": 91.24,
            "unit": "ns/op\t      80 B/op\t       2 allocs/op",
            "extra": "12966826 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda05TypeCode05 - ns/op",
            "value": 91.24,
            "unit": "ns/op",
            "extra": "12966826 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda05TypeCode05 - B/op",
            "value": 80,
            "unit": "B/op",
            "extra": "12966826 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda05TypeCode05 - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "12966826 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda10Parse",
            "value": 802.5,
            "unit": "ns/op\t      78 B/op\t       4 allocs/op",
            "extra": "1489652 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda10Parse - ns/op",
            "value": 802.5,
            "unit": "ns/op",
            "extra": "1489652 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda10Parse - B/op",
            "value": 78,
            "unit": "B/op",
            "extra": "1489652 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda10Parse - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "1489652 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda10ValidTypeCode",
            "value": 94.95,
            "unit": "ns/op\t      80 B/op\t       2 allocs/op",
            "extra": "12526634 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda10ValidTypeCode - ns/op",
            "value": 94.95,
            "unit": "ns/op",
            "extra": "12526634 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda10ValidTypeCode - B/op",
            "value": 80,
            "unit": "B/op",
            "extra": "12526634 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda10ValidTypeCode - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "12526634 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda10TypeCode10",
            "value": 95.83,
            "unit": "ns/op\t      80 B/op\t       2 allocs/op",
            "extra": "12329649 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda10TypeCode10 - ns/op",
            "value": 95.83,
            "unit": "ns/op",
            "extra": "12329649 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda10TypeCode10 - B/op",
            "value": 80,
            "unit": "B/op",
            "extra": "12329649 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda10TypeCode10 - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "12329649 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda10TransactionTypeCode",
            "value": 102.5,
            "unit": "ns/op\t      80 B/op\t       2 allocs/op",
            "extra": "11694895 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda10TransactionTypeCode - ns/op",
            "value": 102.5,
            "unit": "ns/op",
            "extra": "11694895 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda10TransactionTypeCode - B/op",
            "value": 80,
            "unit": "B/op",
            "extra": "11694895 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda10TransactionTypeCode - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "11694895 times\n4 procs"
          },
          {
            "name": "BenchmarkForeignTraceNumberAlphaNumeric",
            "value": 328.4,
            "unit": "ns/op\t     160 B/op\t       4 allocs/op",
            "extra": "3645949 times\n4 procs"
          },
          {
            "name": "BenchmarkForeignTraceNumberAlphaNumeric - ns/op",
            "value": 328.4,
            "unit": "ns/op",
            "extra": "3645949 times\n4 procs"
          },
          {
            "name": "BenchmarkForeignTraceNumberAlphaNumeric - B/op",
            "value": 160,
            "unit": "B/op",
            "extra": "3645949 times\n4 procs"
          },
          {
            "name": "BenchmarkForeignTraceNumberAlphaNumeric - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "3645949 times\n4 procs"
          },
          {
            "name": "BenchmarkNameAlphaNumeric",
            "value": 348.4,
            "unit": "ns/op\t     160 B/op\t       4 allocs/op",
            "extra": "3427744 times\n4 procs"
          },
          {
            "name": "BenchmarkNameAlphaNumeric - ns/op",
            "value": 348.4,
            "unit": "ns/op",
            "extra": "3427744 times\n4 procs"
          },
          {
            "name": "BenchmarkNameAlphaNumeric - B/op",
            "value": 160,
            "unit": "B/op",
            "extra": "3427744 times\n4 procs"
          },
          {
            "name": "BenchmarkNameAlphaNumeric - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "3427744 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda10FieldInclusionTypeCode",
            "value": 65.56,
            "unit": "ns/op\t      64 B/op\t       1 allocs/op",
            "extra": "18100273 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda10FieldInclusionTypeCode - ns/op",
            "value": 65.56,
            "unit": "ns/op",
            "extra": "18100273 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda10FieldInclusionTypeCode - B/op",
            "value": 64,
            "unit": "B/op",
            "extra": "18100273 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda10FieldInclusionTypeCode - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "18100273 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda10FieldInclusionTransactionTypeCode",
            "value": 65.97,
            "unit": "ns/op\t      64 B/op\t       1 allocs/op",
            "extra": "17789251 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda10FieldInclusionTransactionTypeCode - ns/op",
            "value": 65.97,
            "unit": "ns/op",
            "extra": "17789251 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda10FieldInclusionTransactionTypeCode - B/op",
            "value": 64,
            "unit": "B/op",
            "extra": "17789251 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda10FieldInclusionTransactionTypeCode - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "17789251 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda10FieldInclusionForeignPaymentAmount",
            "value": 50.77,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "23530722 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda10FieldInclusionForeignPaymentAmount - ns/op",
            "value": 50.77,
            "unit": "ns/op",
            "extra": "23530722 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda10FieldInclusionForeignPaymentAmount - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "23530722 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda10FieldInclusionForeignPaymentAmount - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "23530722 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda10FieldInclusionName",
            "value": 65.75,
            "unit": "ns/op\t      64 B/op\t       1 allocs/op",
            "extra": "17646574 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda10FieldInclusionName - ns/op",
            "value": 65.75,
            "unit": "ns/op",
            "extra": "17646574 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda10FieldInclusionName - B/op",
            "value": 64,
            "unit": "B/op",
            "extra": "17646574 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda10FieldInclusionName - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "17646574 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda10FieldInclusionEntryDetailSequenceNumber",
            "value": 131.6,
            "unit": "ns/op\t      88 B/op\t       3 allocs/op",
            "extra": "9070915 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda10FieldInclusionEntryDetailSequenceNumber - ns/op",
            "value": 131.6,
            "unit": "ns/op",
            "extra": "9070915 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda10FieldInclusionEntryDetailSequenceNumber - B/op",
            "value": 88,
            "unit": "B/op",
            "extra": "9070915 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda10FieldInclusionEntryDetailSequenceNumber - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "9070915 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda10String",
            "value": 1185,
            "unit": "ns/op\t     288 B/op\t      10 allocs/op",
            "extra": "938353 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda10String - ns/op",
            "value": 1185,
            "unit": "ns/op",
            "extra": "938353 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda10String - B/op",
            "value": 288,
            "unit": "B/op",
            "extra": "938353 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda10String - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "938353 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda11Parse",
            "value": 737.5,
            "unit": "ns/op\t      98 B/op\t       3 allocs/op",
            "extra": "1622049 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda11Parse - ns/op",
            "value": 737.5,
            "unit": "ns/op",
            "extra": "1622049 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda11Parse - B/op",
            "value": 98,
            "unit": "B/op",
            "extra": "1622049 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda11Parse - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "1622049 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda11ValidTypeCode",
            "value": 97.31,
            "unit": "ns/op\t      80 B/op\t       2 allocs/op",
            "extra": "12160258 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda11ValidTypeCode - ns/op",
            "value": 97.31,
            "unit": "ns/op",
            "extra": "12160258 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda11ValidTypeCode - B/op",
            "value": 80,
            "unit": "B/op",
            "extra": "12160258 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda11ValidTypeCode - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "12160258 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda11TypeCode11",
            "value": 93.53,
            "unit": "ns/op\t      80 B/op\t       2 allocs/op",
            "extra": "12724593 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda11TypeCode11 - ns/op",
            "value": 93.53,
            "unit": "ns/op",
            "extra": "12724593 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda11TypeCode11 - B/op",
            "value": 80,
            "unit": "B/op",
            "extra": "12724593 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda11TypeCode11 - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "12724593 times\n4 procs"
          },
          {
            "name": "BenchmarkOriginatorNameAlphaNumeric",
            "value": 310.8,
            "unit": "ns/op\t     160 B/op\t       4 allocs/op",
            "extra": "3855115 times\n4 procs"
          },
          {
            "name": "BenchmarkOriginatorNameAlphaNumeric - ns/op",
            "value": 310.8,
            "unit": "ns/op",
            "extra": "3855115 times\n4 procs"
          },
          {
            "name": "BenchmarkOriginatorNameAlphaNumeric - B/op",
            "value": 160,
            "unit": "B/op",
            "extra": "3855115 times\n4 procs"
          },
          {
            "name": "BenchmarkOriginatorNameAlphaNumeric - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "3855115 times\n4 procs"
          },
          {
            "name": "BenchmarkOriginatorStreetAddressAlphaNumeric",
            "value": 328.9,
            "unit": "ns/op\t     160 B/op\t       4 allocs/op",
            "extra": "3654564 times\n4 procs"
          },
          {
            "name": "BenchmarkOriginatorStreetAddressAlphaNumeric - ns/op",
            "value": 328.9,
            "unit": "ns/op",
            "extra": "3654564 times\n4 procs"
          },
          {
            "name": "BenchmarkOriginatorStreetAddressAlphaNumeric - B/op",
            "value": 160,
            "unit": "B/op",
            "extra": "3654564 times\n4 procs"
          },
          {
            "name": "BenchmarkOriginatorStreetAddressAlphaNumeric - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "3654564 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda11FieldInclusionTypeCode",
            "value": 65.4,
            "unit": "ns/op\t      64 B/op\t       1 allocs/op",
            "extra": "18072362 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda11FieldInclusionTypeCode - ns/op",
            "value": 65.4,
            "unit": "ns/op",
            "extra": "18072362 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda11FieldInclusionTypeCode - B/op",
            "value": 64,
            "unit": "B/op",
            "extra": "18072362 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda11FieldInclusionTypeCode - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "18072362 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda11FieldInclusionOriginatorName",
            "value": 63.68,
            "unit": "ns/op\t      64 B/op\t       1 allocs/op",
            "extra": "18565156 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda11FieldInclusionOriginatorName - ns/op",
            "value": 63.68,
            "unit": "ns/op",
            "extra": "18565156 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda11FieldInclusionOriginatorName - B/op",
            "value": 64,
            "unit": "B/op",
            "extra": "18565156 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda11FieldInclusionOriginatorName - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "18565156 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda11FieldInclusionOriginatorStreetAddress",
            "value": 65.43,
            "unit": "ns/op\t      64 B/op\t       1 allocs/op",
            "extra": "18217262 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda11FieldInclusionOriginatorStreetAddress - ns/op",
            "value": 65.43,
            "unit": "ns/op",
            "extra": "18217262 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda11FieldInclusionOriginatorStreetAddress - B/op",
            "value": 64,
            "unit": "B/op",
            "extra": "18217262 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda11FieldInclusionOriginatorStreetAddress - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "18217262 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda11FieldInclusionEntryDetailSequenceNumber",
            "value": 136.6,
            "unit": "ns/op\t      88 B/op\t       3 allocs/op",
            "extra": "8701993 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda11FieldInclusionEntryDetailSequenceNumber - ns/op",
            "value": 136.6,
            "unit": "ns/op",
            "extra": "8701993 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda11FieldInclusionEntryDetailSequenceNumber - B/op",
            "value": 88,
            "unit": "B/op",
            "extra": "8701993 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda11FieldInclusionEntryDetailSequenceNumber - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "8701993 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda11String",
            "value": 1020,
            "unit": "ns/op\t     304 B/op\t       7 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda11String - ns/op",
            "value": 1020,
            "unit": "ns/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda11String - B/op",
            "value": 304,
            "unit": "B/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda11String - allocs/op",
            "value": 7,
            "unit": "allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda12Parse",
            "value": 729.8,
            "unit": "ns/op\t      98 B/op\t       3 allocs/op",
            "extra": "1640653 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda12Parse - ns/op",
            "value": 729.8,
            "unit": "ns/op",
            "extra": "1640653 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda12Parse - B/op",
            "value": 98,
            "unit": "B/op",
            "extra": "1640653 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda12Parse - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "1640653 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda12ValidTypeCode",
            "value": 91.29,
            "unit": "ns/op\t      80 B/op\t       2 allocs/op",
            "extra": "12972306 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda12ValidTypeCode - ns/op",
            "value": 91.29,
            "unit": "ns/op",
            "extra": "12972306 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda12ValidTypeCode - B/op",
            "value": 80,
            "unit": "B/op",
            "extra": "12972306 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda12ValidTypeCode - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "12972306 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda12TypeCode12",
            "value": 91.91,
            "unit": "ns/op\t      80 B/op\t       2 allocs/op",
            "extra": "12823023 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda12TypeCode12 - ns/op",
            "value": 91.91,
            "unit": "ns/op",
            "extra": "12823023 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda12TypeCode12 - B/op",
            "value": 80,
            "unit": "B/op",
            "extra": "12823023 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda12TypeCode12 - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "12823023 times\n4 procs"
          },
          {
            "name": "BenchmarkOriginatorCityStateProvinceAlphaNumeric",
            "value": 319.8,
            "unit": "ns/op\t     160 B/op\t       4 allocs/op",
            "extra": "3736315 times\n4 procs"
          },
          {
            "name": "BenchmarkOriginatorCityStateProvinceAlphaNumeric - ns/op",
            "value": 319.8,
            "unit": "ns/op",
            "extra": "3736315 times\n4 procs"
          },
          {
            "name": "BenchmarkOriginatorCityStateProvinceAlphaNumeric - B/op",
            "value": 160,
            "unit": "B/op",
            "extra": "3736315 times\n4 procs"
          },
          {
            "name": "BenchmarkOriginatorCityStateProvinceAlphaNumeric - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "3736315 times\n4 procs"
          },
          {
            "name": "BenchmarkOriginatorCountryPostalCodeAlphaNumeric",
            "value": 341,
            "unit": "ns/op\t     160 B/op\t       4 allocs/op",
            "extra": "3529704 times\n4 procs"
          },
          {
            "name": "BenchmarkOriginatorCountryPostalCodeAlphaNumeric - ns/op",
            "value": 341,
            "unit": "ns/op",
            "extra": "3529704 times\n4 procs"
          },
          {
            "name": "BenchmarkOriginatorCountryPostalCodeAlphaNumeric - B/op",
            "value": 160,
            "unit": "B/op",
            "extra": "3529704 times\n4 procs"
          },
          {
            "name": "BenchmarkOriginatorCountryPostalCodeAlphaNumeric - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "3529704 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda12FieldInclusionTypeCode",
            "value": 64.54,
            "unit": "ns/op\t      64 B/op\t       1 allocs/op",
            "extra": "18377335 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda12FieldInclusionTypeCode - ns/op",
            "value": 64.54,
            "unit": "ns/op",
            "extra": "18377335 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda12FieldInclusionTypeCode - B/op",
            "value": 64,
            "unit": "B/op",
            "extra": "18377335 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda12FieldInclusionTypeCode - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "18377335 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda12FieldInclusionOriginatorCityStateProvince",
            "value": 64.51,
            "unit": "ns/op\t      64 B/op\t       1 allocs/op",
            "extra": "18517131 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda12FieldInclusionOriginatorCityStateProvince - ns/op",
            "value": 64.51,
            "unit": "ns/op",
            "extra": "18517131 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda12FieldInclusionOriginatorCityStateProvince - B/op",
            "value": 64,
            "unit": "B/op",
            "extra": "18517131 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda12FieldInclusionOriginatorCityStateProvince - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "18517131 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda12FieldInclusionOriginatorCountryPostalCode",
            "value": 64.56,
            "unit": "ns/op\t      64 B/op\t       1 allocs/op",
            "extra": "18087651 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda12FieldInclusionOriginatorCountryPostalCode - ns/op",
            "value": 64.56,
            "unit": "ns/op",
            "extra": "18087651 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda12FieldInclusionOriginatorCountryPostalCode - B/op",
            "value": 64,
            "unit": "B/op",
            "extra": "18087651 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda12FieldInclusionOriginatorCountryPostalCode - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "18087651 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda12FieldInclusionEntryDetailSequenceNumber",
            "value": 135.7,
            "unit": "ns/op\t      88 B/op\t       3 allocs/op",
            "extra": "8899286 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda12FieldInclusionEntryDetailSequenceNumber - ns/op",
            "value": 135.7,
            "unit": "ns/op",
            "extra": "8899286 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda12FieldInclusionEntryDetailSequenceNumber - B/op",
            "value": 88,
            "unit": "B/op",
            "extra": "8899286 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda12FieldInclusionEntryDetailSequenceNumber - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "8899286 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda12String",
            "value": 1008,
            "unit": "ns/op\t     304 B/op\t       7 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda12String - ns/op",
            "value": 1008,
            "unit": "ns/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda12String - B/op",
            "value": 304,
            "unit": "B/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda12String - allocs/op",
            "value": 7,
            "unit": "allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda13Parse",
            "value": 762,
            "unit": "ns/op\t     104 B/op\t       5 allocs/op",
            "extra": "1570140 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda13Parse - ns/op",
            "value": 762,
            "unit": "ns/op",
            "extra": "1570140 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda13Parse - B/op",
            "value": 104,
            "unit": "B/op",
            "extra": "1570140 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda13Parse - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "1570140 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda13ValidTypeCode",
            "value": 93.1,
            "unit": "ns/op\t      80 B/op\t       2 allocs/op",
            "extra": "12640966 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda13ValidTypeCode - ns/op",
            "value": 93.1,
            "unit": "ns/op",
            "extra": "12640966 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda13ValidTypeCode - B/op",
            "value": 80,
            "unit": "B/op",
            "extra": "12640966 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda13ValidTypeCode - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "12640966 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda13TypeCode13",
            "value": 92.55,
            "unit": "ns/op\t      80 B/op\t       2 allocs/op",
            "extra": "12809952 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda13TypeCode13 - ns/op",
            "value": 92.55,
            "unit": "ns/op",
            "extra": "12809952 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda13TypeCode13 - B/op",
            "value": 80,
            "unit": "B/op",
            "extra": "12809952 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda13TypeCode13 - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "12809952 times\n4 procs"
          },
          {
            "name": "BenchmarkODFINameAlphaNumeric",
            "value": 326.3,
            "unit": "ns/op\t     160 B/op\t       4 allocs/op",
            "extra": "3671908 times\n4 procs"
          },
          {
            "name": "BenchmarkODFINameAlphaNumeric - ns/op",
            "value": 326.3,
            "unit": "ns/op",
            "extra": "3671908 times\n4 procs"
          },
          {
            "name": "BenchmarkODFINameAlphaNumeric - B/op",
            "value": 160,
            "unit": "B/op",
            "extra": "3671908 times\n4 procs"
          },
          {
            "name": "BenchmarkODFINameAlphaNumeric - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "3671908 times\n4 procs"
          },
          {
            "name": "BenchmarkODFIIDNumberQualifierValid",
            "value": 109.3,
            "unit": "ns/op\t      80 B/op\t       2 allocs/op",
            "extra": "10322372 times\n4 procs"
          },
          {
            "name": "BenchmarkODFIIDNumberQualifierValid - ns/op",
            "value": 109.3,
            "unit": "ns/op",
            "extra": "10322372 times\n4 procs"
          },
          {
            "name": "BenchmarkODFIIDNumberQualifierValid - B/op",
            "value": 80,
            "unit": "B/op",
            "extra": "10322372 times\n4 procs"
          },
          {
            "name": "BenchmarkODFIIDNumberQualifierValid - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "10322372 times\n4 procs"
          },
          {
            "name": "BenchmarkODFIIdentificationAlphaNumeric",
            "value": 335.1,
            "unit": "ns/op\t     160 B/op\t       4 allocs/op",
            "extra": "3637825 times\n4 procs"
          },
          {
            "name": "BenchmarkODFIIdentificationAlphaNumeric - ns/op",
            "value": 335.1,
            "unit": "ns/op",
            "extra": "3637825 times\n4 procs"
          },
          {
            "name": "BenchmarkODFIIdentificationAlphaNumeric - B/op",
            "value": 160,
            "unit": "B/op",
            "extra": "3637825 times\n4 procs"
          },
          {
            "name": "BenchmarkODFIIdentificationAlphaNumeric - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "3637825 times\n4 procs"
          },
          {
            "name": "BenchmarkODFIBranchCountryCodeAlphaNumeric",
            "value": 373.2,
            "unit": "ns/op\t     160 B/op\t       4 allocs/op",
            "extra": "3422120 times\n4 procs"
          },
          {
            "name": "BenchmarkODFIBranchCountryCodeAlphaNumeric - ns/op",
            "value": 373.2,
            "unit": "ns/op",
            "extra": "3422120 times\n4 procs"
          },
          {
            "name": "BenchmarkODFIBranchCountryCodeAlphaNumeric - B/op",
            "value": 160,
            "unit": "B/op",
            "extra": "3422120 times\n4 procs"
          },
          {
            "name": "BenchmarkODFIBranchCountryCodeAlphaNumeric - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "3422120 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda13FieldInclusionTypeCode",
            "value": 64.49,
            "unit": "ns/op\t      64 B/op\t       1 allocs/op",
            "extra": "18297108 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda13FieldInclusionTypeCode - ns/op",
            "value": 64.49,
            "unit": "ns/op",
            "extra": "18297108 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda13FieldInclusionTypeCode - B/op",
            "value": 64,
            "unit": "B/op",
            "extra": "18297108 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda13FieldInclusionTypeCode - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "18297108 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda13FieldInclusionODFIName",
            "value": 64.13,
            "unit": "ns/op\t      64 B/op\t       1 allocs/op",
            "extra": "18529390 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda13FieldInclusionODFIName - ns/op",
            "value": 64.13,
            "unit": "ns/op",
            "extra": "18529390 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda13FieldInclusionODFIName - B/op",
            "value": 64,
            "unit": "B/op",
            "extra": "18529390 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda13FieldInclusionODFIName - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "18529390 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda13FieldInclusionODFIIDNumberQualifier",
            "value": 64.76,
            "unit": "ns/op\t      64 B/op\t       1 allocs/op",
            "extra": "18498902 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda13FieldInclusionODFIIDNumberQualifier - ns/op",
            "value": 64.76,
            "unit": "ns/op",
            "extra": "18498902 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda13FieldInclusionODFIIDNumberQualifier - B/op",
            "value": 64,
            "unit": "B/op",
            "extra": "18498902 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda13FieldInclusionODFIIDNumberQualifier - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "18498902 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda13FieldInclusionODFIIdentification",
            "value": 67.76,
            "unit": "ns/op\t      64 B/op\t       1 allocs/op",
            "extra": "17651121 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda13FieldInclusionODFIIdentification - ns/op",
            "value": 67.76,
            "unit": "ns/op",
            "extra": "17651121 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda13FieldInclusionODFIIdentification - B/op",
            "value": 64,
            "unit": "B/op",
            "extra": "17651121 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda13FieldInclusionODFIIdentification - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "17651121 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda13FieldInclusionODFIBranchCountryCode",
            "value": 64.96,
            "unit": "ns/op\t      64 B/op\t       1 allocs/op",
            "extra": "17933250 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda13FieldInclusionODFIBranchCountryCode - ns/op",
            "value": 64.96,
            "unit": "ns/op",
            "extra": "17933250 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda13FieldInclusionODFIBranchCountryCode - B/op",
            "value": 64,
            "unit": "B/op",
            "extra": "17933250 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda13FieldInclusionODFIBranchCountryCode - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "17933250 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda13FieldInclusionEntryDetailSequenceNumber",
            "value": 135.4,
            "unit": "ns/op\t      88 B/op\t       3 allocs/op",
            "extra": "8886169 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda13FieldInclusionEntryDetailSequenceNumber - ns/op",
            "value": 135.4,
            "unit": "ns/op",
            "extra": "8886169 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda13FieldInclusionEntryDetailSequenceNumber - B/op",
            "value": 88,
            "unit": "B/op",
            "extra": "8886169 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda13FieldInclusionEntryDetailSequenceNumber - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "8886169 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda13String",
            "value": 1035,
            "unit": "ns/op\t     256 B/op\t       8 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda13String - ns/op",
            "value": 1035,
            "unit": "ns/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda13String - B/op",
            "value": 256,
            "unit": "B/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda13String - allocs/op",
            "value": 8,
            "unit": "allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda14Parse",
            "value": 809.2,
            "unit": "ns/op\t     104 B/op\t       5 allocs/op",
            "extra": "1486220 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda14Parse - ns/op",
            "value": 809.2,
            "unit": "ns/op",
            "extra": "1486220 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda14Parse - B/op",
            "value": 104,
            "unit": "B/op",
            "extra": "1486220 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda14Parse - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "1486220 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda14ValidTypeCode",
            "value": 94.33,
            "unit": "ns/op\t      80 B/op\t       2 allocs/op",
            "extra": "12607418 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda14ValidTypeCode - ns/op",
            "value": 94.33,
            "unit": "ns/op",
            "extra": "12607418 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda14ValidTypeCode - B/op",
            "value": 80,
            "unit": "B/op",
            "extra": "12607418 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda14ValidTypeCode - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "12607418 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda14TypeCode14",
            "value": 93.15,
            "unit": "ns/op\t      80 B/op\t       2 allocs/op",
            "extra": "12790159 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda14TypeCode14 - ns/op",
            "value": 93.15,
            "unit": "ns/op",
            "extra": "12790159 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda14TypeCode14 - B/op",
            "value": 80,
            "unit": "B/op",
            "extra": "12790159 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda14TypeCode14 - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "12790159 times\n4 procs"
          },
          {
            "name": "BenchmarkRDFINameAlphaNumeric",
            "value": 315.6,
            "unit": "ns/op\t     160 B/op\t       4 allocs/op",
            "extra": "3781850 times\n4 procs"
          },
          {
            "name": "BenchmarkRDFINameAlphaNumeric - ns/op",
            "value": 315.6,
            "unit": "ns/op",
            "extra": "3781850 times\n4 procs"
          },
          {
            "name": "BenchmarkRDFINameAlphaNumeric - B/op",
            "value": 160,
            "unit": "B/op",
            "extra": "3781850 times\n4 procs"
          },
          {
            "name": "BenchmarkRDFINameAlphaNumeric - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "3781850 times\n4 procs"
          },
          {
            "name": "BenchmarkRDFIIDNumberQualifierValid",
            "value": 107.1,
            "unit": "ns/op\t      80 B/op\t       2 allocs/op",
            "extra": "11117917 times\n4 procs"
          },
          {
            "name": "BenchmarkRDFIIDNumberQualifierValid - ns/op",
            "value": 107.1,
            "unit": "ns/op",
            "extra": "11117917 times\n4 procs"
          },
          {
            "name": "BenchmarkRDFIIDNumberQualifierValid - B/op",
            "value": 80,
            "unit": "B/op",
            "extra": "11117917 times\n4 procs"
          },
          {
            "name": "BenchmarkRDFIIDNumberQualifierValid - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "11117917 times\n4 procs"
          },
          {
            "name": "BenchmarkRDFIIdentificationAlphaNumeric",
            "value": 335.6,
            "unit": "ns/op\t     160 B/op\t       4 allocs/op",
            "extra": "3566613 times\n4 procs"
          },
          {
            "name": "BenchmarkRDFIIdentificationAlphaNumeric - ns/op",
            "value": 335.6,
            "unit": "ns/op",
            "extra": "3566613 times\n4 procs"
          },
          {
            "name": "BenchmarkRDFIIdentificationAlphaNumeric - B/op",
            "value": 160,
            "unit": "B/op",
            "extra": "3566613 times\n4 procs"
          },
          {
            "name": "BenchmarkRDFIIdentificationAlphaNumeric - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "3566613 times\n4 procs"
          },
          {
            "name": "BenchmarkRDFIBranchCountryCodeAlphaNumeric",
            "value": 355.4,
            "unit": "ns/op\t     160 B/op\t       4 allocs/op",
            "extra": "3395323 times\n4 procs"
          },
          {
            "name": "BenchmarkRDFIBranchCountryCodeAlphaNumeric - ns/op",
            "value": 355.4,
            "unit": "ns/op",
            "extra": "3395323 times\n4 procs"
          },
          {
            "name": "BenchmarkRDFIBranchCountryCodeAlphaNumeric - B/op",
            "value": 160,
            "unit": "B/op",
            "extra": "3395323 times\n4 procs"
          },
          {
            "name": "BenchmarkRDFIBranchCountryCodeAlphaNumeric - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "3395323 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda14FieldInclusionTypeCode",
            "value": 65.02,
            "unit": "ns/op\t      64 B/op\t       1 allocs/op",
            "extra": "18292773 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda14FieldInclusionTypeCode - ns/op",
            "value": 65.02,
            "unit": "ns/op",
            "extra": "18292773 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda14FieldInclusionTypeCode - B/op",
            "value": 64,
            "unit": "B/op",
            "extra": "18292773 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda14FieldInclusionTypeCode - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "18292773 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda14FieldInclusionRDFIName",
            "value": 64.72,
            "unit": "ns/op\t      64 B/op\t       1 allocs/op",
            "extra": "18436807 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda14FieldInclusionRDFIName - ns/op",
            "value": 64.72,
            "unit": "ns/op",
            "extra": "18436807 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda14FieldInclusionRDFIName - B/op",
            "value": 64,
            "unit": "B/op",
            "extra": "18436807 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda14FieldInclusionRDFIName - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "18436807 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda14FieldInclusionRDFIIdNumberQualifier",
            "value": 65.54,
            "unit": "ns/op\t      64 B/op\t       1 allocs/op",
            "extra": "17997483 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda14FieldInclusionRDFIIdNumberQualifier - ns/op",
            "value": 65.54,
            "unit": "ns/op",
            "extra": "17997483 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda14FieldInclusionRDFIIdNumberQualifier - B/op",
            "value": 64,
            "unit": "B/op",
            "extra": "17997483 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda14FieldInclusionRDFIIdNumberQualifier - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "17997483 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda14FieldInclusionRDFIIdentification",
            "value": 66.26,
            "unit": "ns/op\t      64 B/op\t       1 allocs/op",
            "extra": "18038212 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda14FieldInclusionRDFIIdentification - ns/op",
            "value": 66.26,
            "unit": "ns/op",
            "extra": "18038212 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda14FieldInclusionRDFIIdentification - B/op",
            "value": 64,
            "unit": "B/op",
            "extra": "18038212 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda14FieldInclusionRDFIIdentification - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "18038212 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda14FieldInclusionRDFIBranchCountryCode",
            "value": 65.25,
            "unit": "ns/op\t      64 B/op\t       1 allocs/op",
            "extra": "18041781 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda14FieldInclusionRDFIBranchCountryCode - ns/op",
            "value": 65.25,
            "unit": "ns/op",
            "extra": "18041781 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda14FieldInclusionRDFIBranchCountryCode - B/op",
            "value": 64,
            "unit": "B/op",
            "extra": "18041781 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda14FieldInclusionRDFIBranchCountryCode - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "18041781 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda14FieldInclusionEntryDetailSequenceNumber",
            "value": 143.5,
            "unit": "ns/op\t      88 B/op\t       3 allocs/op",
            "extra": "8379324 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda14FieldInclusionEntryDetailSequenceNumber - ns/op",
            "value": 143.5,
            "unit": "ns/op",
            "extra": "8379324 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda14FieldInclusionEntryDetailSequenceNumber - B/op",
            "value": 88,
            "unit": "B/op",
            "extra": "8379324 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda14FieldInclusionEntryDetailSequenceNumber - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "8379324 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda14String",
            "value": 1084,
            "unit": "ns/op\t     256 B/op\t       8 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda14String - ns/op",
            "value": 1084,
            "unit": "ns/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda14String - B/op",
            "value": 256,
            "unit": "B/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda14String - allocs/op",
            "value": 8,
            "unit": "allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda15Parse",
            "value": 757.6,
            "unit": "ns/op\t     114 B/op\t       4 allocs/op",
            "extra": "1578363 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda15Parse - ns/op",
            "value": 757.6,
            "unit": "ns/op",
            "extra": "1578363 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda15Parse - B/op",
            "value": 114,
            "unit": "B/op",
            "extra": "1578363 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda15Parse - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "1578363 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda15ValidTypeCode",
            "value": 91.64,
            "unit": "ns/op\t      80 B/op\t       2 allocs/op",
            "extra": "13084708 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda15ValidTypeCode - ns/op",
            "value": 91.64,
            "unit": "ns/op",
            "extra": "13084708 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda15ValidTypeCode - B/op",
            "value": 80,
            "unit": "B/op",
            "extra": "13084708 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda15ValidTypeCode - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "13084708 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda15TypeCode15",
            "value": 91.96,
            "unit": "ns/op\t      80 B/op\t       2 allocs/op",
            "extra": "12902562 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda15TypeCode15 - ns/op",
            "value": 91.96,
            "unit": "ns/op",
            "extra": "12902562 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda15TypeCode15 - B/op",
            "value": 80,
            "unit": "B/op",
            "extra": "12902562 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda15TypeCode15 - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "12902562 times\n4 procs"
          },
          {
            "name": "BenchmarkReceiverIDNumberAlphaNumeric",
            "value": 344,
            "unit": "ns/op\t     160 B/op\t       4 allocs/op",
            "extra": "3480591 times\n4 procs"
          },
          {
            "name": "BenchmarkReceiverIDNumberAlphaNumeric - ns/op",
            "value": 344,
            "unit": "ns/op",
            "extra": "3480591 times\n4 procs"
          },
          {
            "name": "BenchmarkReceiverIDNumberAlphaNumeric - B/op",
            "value": 160,
            "unit": "B/op",
            "extra": "3480591 times\n4 procs"
          },
          {
            "name": "BenchmarkReceiverIDNumberAlphaNumeric - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "3480591 times\n4 procs"
          },
          {
            "name": "BenchmarkReceiverStreetAddressAlphaNumeric",
            "value": 345.8,
            "unit": "ns/op\t     160 B/op\t       4 allocs/op",
            "extra": "3474942 times\n4 procs"
          },
          {
            "name": "BenchmarkReceiverStreetAddressAlphaNumeric - ns/op",
            "value": 345.8,
            "unit": "ns/op",
            "extra": "3474942 times\n4 procs"
          },
          {
            "name": "BenchmarkReceiverStreetAddressAlphaNumeric - B/op",
            "value": 160,
            "unit": "B/op",
            "extra": "3474942 times\n4 procs"
          },
          {
            "name": "BenchmarkReceiverStreetAddressAlphaNumeric - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "3474942 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda15FieldInclusionTypeCode",
            "value": 64.71,
            "unit": "ns/op\t      64 B/op\t       1 allocs/op",
            "extra": "18226291 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda15FieldInclusionTypeCode - ns/op",
            "value": 64.71,
            "unit": "ns/op",
            "extra": "18226291 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda15FieldInclusionTypeCode - B/op",
            "value": 64,
            "unit": "B/op",
            "extra": "18226291 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda15FieldInclusionTypeCode - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "18226291 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda15FieldInclusionReceiverStreetAddress",
            "value": 64.38,
            "unit": "ns/op\t      64 B/op\t       1 allocs/op",
            "extra": "18315388 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda15FieldInclusionReceiverStreetAddress - ns/op",
            "value": 64.38,
            "unit": "ns/op",
            "extra": "18315388 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda15FieldInclusionReceiverStreetAddress - B/op",
            "value": 64,
            "unit": "B/op",
            "extra": "18315388 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda15FieldInclusionReceiverStreetAddress - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "18315388 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda15FieldInclusionEntryDetailSequenceNumber",
            "value": 131.5,
            "unit": "ns/op\t      88 B/op\t       3 allocs/op",
            "extra": "9088922 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda15FieldInclusionEntryDetailSequenceNumber - ns/op",
            "value": 131.5,
            "unit": "ns/op",
            "extra": "9088922 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda15FieldInclusionEntryDetailSequenceNumber - B/op",
            "value": 88,
            "unit": "B/op",
            "extra": "9088922 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda15FieldInclusionEntryDetailSequenceNumber - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "9088922 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda15String",
            "value": 994.4,
            "unit": "ns/op\t     272 B/op\t       7 allocs/op",
            "extra": "1207669 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda15String - ns/op",
            "value": 994.4,
            "unit": "ns/op",
            "extra": "1207669 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda15String - B/op",
            "value": 272,
            "unit": "B/op",
            "extra": "1207669 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda15String - allocs/op",
            "value": 7,
            "unit": "allocs/op",
            "extra": "1207669 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda16Parse",
            "value": 752,
            "unit": "ns/op\t      98 B/op\t       3 allocs/op",
            "extra": "1616006 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda16Parse - ns/op",
            "value": 752,
            "unit": "ns/op",
            "extra": "1616006 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda16Parse - B/op",
            "value": 98,
            "unit": "B/op",
            "extra": "1616006 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda16Parse - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "1616006 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda16ValidTypeCode",
            "value": 91.75,
            "unit": "ns/op\t      80 B/op\t       2 allocs/op",
            "extra": "12857374 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda16ValidTypeCode - ns/op",
            "value": 91.75,
            "unit": "ns/op",
            "extra": "12857374 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda16ValidTypeCode - B/op",
            "value": 80,
            "unit": "B/op",
            "extra": "12857374 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda16ValidTypeCode - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "12857374 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda16TypeCode16",
            "value": 92.38,
            "unit": "ns/op\t      80 B/op\t       2 allocs/op",
            "extra": "12875512 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda16TypeCode16 - ns/op",
            "value": 92.38,
            "unit": "ns/op",
            "extra": "12875512 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda16TypeCode16 - B/op",
            "value": 80,
            "unit": "B/op",
            "extra": "12875512 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda16TypeCode16 - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "12875512 times\n4 procs"
          },
          {
            "name": "BenchmarkReceiverCityStateProvinceAlphaNumeric",
            "value": 320.8,
            "unit": "ns/op\t     160 B/op\t       4 allocs/op",
            "extra": "3746605 times\n4 procs"
          },
          {
            "name": "BenchmarkReceiverCityStateProvinceAlphaNumeric - ns/op",
            "value": 320.8,
            "unit": "ns/op",
            "extra": "3746605 times\n4 procs"
          },
          {
            "name": "BenchmarkReceiverCityStateProvinceAlphaNumeric - B/op",
            "value": 160,
            "unit": "B/op",
            "extra": "3746605 times\n4 procs"
          },
          {
            "name": "BenchmarkReceiverCityStateProvinceAlphaNumeric - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "3746605 times\n4 procs"
          },
          {
            "name": "BenchmarkReceiverCountryPostalCodeAlphaNumeric",
            "value": 344.2,
            "unit": "ns/op\t     160 B/op\t       4 allocs/op",
            "extra": "3414952 times\n4 procs"
          },
          {
            "name": "BenchmarkReceiverCountryPostalCodeAlphaNumeric - ns/op",
            "value": 344.2,
            "unit": "ns/op",
            "extra": "3414952 times\n4 procs"
          },
          {
            "name": "BenchmarkReceiverCountryPostalCodeAlphaNumeric - B/op",
            "value": 160,
            "unit": "B/op",
            "extra": "3414952 times\n4 procs"
          },
          {
            "name": "BenchmarkReceiverCountryPostalCodeAlphaNumeric - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "3414952 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda16FieldInclusionTypeCode",
            "value": 65.64,
            "unit": "ns/op\t      64 B/op\t       1 allocs/op",
            "extra": "18082388 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda16FieldInclusionTypeCode - ns/op",
            "value": 65.64,
            "unit": "ns/op",
            "extra": "18082388 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda16FieldInclusionTypeCode - B/op",
            "value": 64,
            "unit": "B/op",
            "extra": "18082388 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda16FieldInclusionTypeCode - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "18082388 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda16FieldInclusionReceiverCityStateProvince",
            "value": 64.52,
            "unit": "ns/op\t      64 B/op\t       1 allocs/op",
            "extra": "18590235 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda16FieldInclusionReceiverCityStateProvince - ns/op",
            "value": 64.52,
            "unit": "ns/op",
            "extra": "18590235 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda16FieldInclusionReceiverCityStateProvince - B/op",
            "value": 64,
            "unit": "B/op",
            "extra": "18590235 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda16FieldInclusionReceiverCityStateProvince - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "18590235 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda16FieldInclusionReceiverCountryPostalCode",
            "value": 64.95,
            "unit": "ns/op\t      64 B/op\t       1 allocs/op",
            "extra": "18126613 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda16FieldInclusionReceiverCountryPostalCode - ns/op",
            "value": 64.95,
            "unit": "ns/op",
            "extra": "18126613 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda16FieldInclusionReceiverCountryPostalCode - B/op",
            "value": 64,
            "unit": "B/op",
            "extra": "18126613 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda16FieldInclusionReceiverCountryPostalCode - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "18126613 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda16FieldInclusionEntryDetailSequenceNumber",
            "value": 129.3,
            "unit": "ns/op\t      88 B/op\t       3 allocs/op",
            "extra": "9214998 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda16FieldInclusionEntryDetailSequenceNumber - ns/op",
            "value": 129.3,
            "unit": "ns/op",
            "extra": "9214998 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda16FieldInclusionEntryDetailSequenceNumber - B/op",
            "value": 88,
            "unit": "B/op",
            "extra": "9214998 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda16FieldInclusionEntryDetailSequenceNumber - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "9214998 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda16String",
            "value": 1022,
            "unit": "ns/op\t     304 B/op\t       7 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda16String - ns/op",
            "value": 1022,
            "unit": "ns/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda16String - B/op",
            "value": 304,
            "unit": "B/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda16String - allocs/op",
            "value": 7,
            "unit": "allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda17Parse",
            "value": 717.6,
            "unit": "ns/op\t      82 B/op\t       2 allocs/op",
            "extra": "1657794 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda17Parse - ns/op",
            "value": 717.6,
            "unit": "ns/op",
            "extra": "1657794 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda17Parse - B/op",
            "value": 82,
            "unit": "B/op",
            "extra": "1657794 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda17Parse - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "1657794 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda17String",
            "value": 974,
            "unit": "ns/op\t     272 B/op\t       6 allocs/op",
            "extra": "1229100 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda17String - ns/op",
            "value": 974,
            "unit": "ns/op",
            "extra": "1229100 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda17String - B/op",
            "value": 272,
            "unit": "B/op",
            "extra": "1229100 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda17String - allocs/op",
            "value": 6,
            "unit": "allocs/op",
            "extra": "1229100 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda17PaymentRelatedInformationAlphaNumeric",
            "value": 309.6,
            "unit": "ns/op\t     160 B/op\t       4 allocs/op",
            "extra": "3874782 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda17PaymentRelatedInformationAlphaNumeric - ns/op",
            "value": 309.6,
            "unit": "ns/op",
            "extra": "3874782 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda17PaymentRelatedInformationAlphaNumeric - B/op",
            "value": 160,
            "unit": "B/op",
            "extra": "3874782 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda17PaymentRelatedInformationAlphaNumeric - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "3874782 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda17ValidTypeCode",
            "value": 90.42,
            "unit": "ns/op\t      80 B/op\t       2 allocs/op",
            "extra": "13113166 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda17ValidTypeCode - ns/op",
            "value": 90.42,
            "unit": "ns/op",
            "extra": "13113166 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda17ValidTypeCode - B/op",
            "value": 80,
            "unit": "B/op",
            "extra": "13113166 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda17ValidTypeCode - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "13113166 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda17TypeCode17",
            "value": 90.49,
            "unit": "ns/op\t      80 B/op\t       2 allocs/op",
            "extra": "13070604 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda17TypeCode17 - ns/op",
            "value": 90.49,
            "unit": "ns/op",
            "extra": "13070604 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda17TypeCode17 - B/op",
            "value": 80,
            "unit": "B/op",
            "extra": "13070604 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda17TypeCode17 - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "13070604 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda18Parse",
            "value": 786,
            "unit": "ns/op\t     104 B/op\t       5 allocs/op",
            "extra": "1520641 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda18Parse - ns/op",
            "value": 786,
            "unit": "ns/op",
            "extra": "1520641 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda18Parse - B/op",
            "value": 104,
            "unit": "B/op",
            "extra": "1520641 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda18Parse - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "1520641 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda18String",
            "value": 1135,
            "unit": "ns/op\t     264 B/op\t      10 allocs/op",
            "extra": "985938 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda18String - ns/op",
            "value": 1135,
            "unit": "ns/op",
            "extra": "985938 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda18String - B/op",
            "value": 264,
            "unit": "B/op",
            "extra": "985938 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda18String - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "985938 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda18ForeignCorrespondentBankNameAlphaNumeric",
            "value": 315.7,
            "unit": "ns/op\t     160 B/op\t       4 allocs/op",
            "extra": "3806520 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda18ForeignCorrespondentBankNameAlphaNumeric - ns/op",
            "value": 315.7,
            "unit": "ns/op",
            "extra": "3806520 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda18ForeignCorrespondentBankNameAlphaNumeric - B/op",
            "value": 160,
            "unit": "B/op",
            "extra": "3806520 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda18ForeignCorrespondentBankNameAlphaNumeric - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "3806520 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda18ForeignCorrespondentBankIDQualifierAlphaNumeric",
            "value": 337.9,
            "unit": "ns/op\t     160 B/op\t       4 allocs/op",
            "extra": "3559489 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda18ForeignCorrespondentBankIDQualifierAlphaNumeric - ns/op",
            "value": 337.9,
            "unit": "ns/op",
            "extra": "3559489 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda18ForeignCorrespondentBankIDQualifierAlphaNumeric - B/op",
            "value": 160,
            "unit": "B/op",
            "extra": "3559489 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda18ForeignCorrespondentBankIDQualifierAlphaNumeric - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "3559489 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda18ForeignCorrespondentBankBranchCountryCodeAlphaNumeric",
            "value": 356.6,
            "unit": "ns/op\t     160 B/op\t       4 allocs/op",
            "extra": "3352178 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda18ForeignCorrespondentBankBranchCountryCodeAlphaNumeric - ns/op",
            "value": 356.6,
            "unit": "ns/op",
            "extra": "3352178 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda18ForeignCorrespondentBankBranchCountryCodeAlphaNumeric - B/op",
            "value": 160,
            "unit": "B/op",
            "extra": "3352178 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda18ForeignCorrespondentBankBranchCountryCodeAlphaNumeric - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "3352178 times\n4 procs"
          },
          {
            "name": "BenchmarkAddendaForeignCorrespondentBankIDNumberAlphaNumeric",
            "value": 341.8,
            "unit": "ns/op\t     160 B/op\t       4 allocs/op",
            "extra": "3503941 times\n4 procs"
          },
          {
            "name": "BenchmarkAddendaForeignCorrespondentBankIDNumberAlphaNumeric - ns/op",
            "value": 341.8,
            "unit": "ns/op",
            "extra": "3503941 times\n4 procs"
          },
          {
            "name": "BenchmarkAddendaForeignCorrespondentBankIDNumberAlphaNumeric - B/op",
            "value": 160,
            "unit": "B/op",
            "extra": "3503941 times\n4 procs"
          },
          {
            "name": "BenchmarkAddendaForeignCorrespondentBankIDNumberAlphaNumeric - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "3503941 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda18ValidTypeCode",
            "value": 92.93,
            "unit": "ns/op\t      80 B/op\t       2 allocs/op",
            "extra": "12820251 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda18ValidTypeCode - ns/op",
            "value": 92.93,
            "unit": "ns/op",
            "extra": "12820251 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda18ValidTypeCode - B/op",
            "value": 80,
            "unit": "B/op",
            "extra": "12820251 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda18ValidTypeCode - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "12820251 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda18TypeCode18",
            "value": 92.07,
            "unit": "ns/op\t      80 B/op\t       2 allocs/op",
            "extra": "12820851 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda18TypeCode18 - ns/op",
            "value": 92.07,
            "unit": "ns/op",
            "extra": "12820851 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda18TypeCode18 - B/op",
            "value": 80,
            "unit": "B/op",
            "extra": "12820851 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda18TypeCode18 - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "12820851 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda98Parse",
            "value": 831.8,
            "unit": "ns/op\t      88 B/op\t       7 allocs/op",
            "extra": "1446936 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda98Parse - ns/op",
            "value": 831.8,
            "unit": "ns/op",
            "extra": "1446936 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda98Parse - B/op",
            "value": 88,
            "unit": "B/op",
            "extra": "1446936 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda98Parse - allocs/op",
            "value": 7,
            "unit": "allocs/op",
            "extra": "1446936 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda98String",
            "value": 1100,
            "unit": "ns/op\t     216 B/op\t       9 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda98String - ns/op",
            "value": 1100,
            "unit": "ns/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda98String - B/op",
            "value": 216,
            "unit": "B/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda98String - allocs/op",
            "value": 9,
            "unit": "allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda98ValidTypeCode",
            "value": 90.26,
            "unit": "ns/op\t      80 B/op\t       2 allocs/op",
            "extra": "13153432 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda98ValidTypeCode - ns/op",
            "value": 90.26,
            "unit": "ns/op",
            "extra": "13153432 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda98ValidTypeCode - B/op",
            "value": 80,
            "unit": "B/op",
            "extra": "13153432 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda98ValidTypeCode - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "13153432 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda98ValidCorrectedData",
            "value": 78.78,
            "unit": "ns/op\t      64 B/op\t       1 allocs/op",
            "extra": "15198274 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda98ValidCorrectedData - ns/op",
            "value": 78.78,
            "unit": "ns/op",
            "extra": "15198274 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda98ValidCorrectedData - B/op",
            "value": 64,
            "unit": "B/op",
            "extra": "15198274 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda98ValidCorrectedData - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "15198274 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda98ValidateTrue",
            "value": 18.22,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "66636124 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda98ValidateTrue - ns/op",
            "value": 18.22,
            "unit": "ns/op",
            "extra": "66636124 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda98ValidateTrue - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "66636124 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda98ValidateTrue - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "66636124 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda98ValidateChangeCodeFalse",
            "value": 110.1,
            "unit": "ns/op\t      80 B/op\t       2 allocs/op",
            "extra": "10909263 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda98ValidateChangeCodeFalse - ns/op",
            "value": 110.1,
            "unit": "ns/op",
            "extra": "10909263 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda98ValidateChangeCodeFalse - B/op",
            "value": 80,
            "unit": "B/op",
            "extra": "10909263 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda98ValidateChangeCodeFalse - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "10909263 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda98OriginalTraceField",
            "value": 67.35,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "17600794 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda98OriginalTraceField - ns/op",
            "value": 67.35,
            "unit": "ns/op",
            "extra": "17600794 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda98OriginalTraceField - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "17600794 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda98OriginalTraceField - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "17600794 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda98OriginalDFIField",
            "value": 46.14,
            "unit": "ns/op\t       8 B/op\t       1 allocs/op",
            "extra": "24292959 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda98OriginalDFIField - ns/op",
            "value": 46.14,
            "unit": "ns/op",
            "extra": "24292959 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda98OriginalDFIField - B/op",
            "value": 8,
            "unit": "B/op",
            "extra": "24292959 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda98OriginalDFIField - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "24292959 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda98CorrectedDataField",
            "value": 67.64,
            "unit": "ns/op\t      32 B/op\t       1 allocs/op",
            "extra": "17279661 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda98CorrectedDataField - ns/op",
            "value": 67.64,
            "unit": "ns/op",
            "extra": "17279661 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda98CorrectedDataField - B/op",
            "value": 32,
            "unit": "B/op",
            "extra": "17279661 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda98CorrectedDataField - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "17279661 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda98TraceNumberField",
            "value": 59.27,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "19941122 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda98TraceNumberField - ns/op",
            "value": 59.27,
            "unit": "ns/op",
            "extra": "19941122 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda98TraceNumberField - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "19941122 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda98TraceNumberField - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "19941122 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda98TypeCodeNil",
            "value": 63,
            "unit": "ns/op\t      64 B/op\t       1 allocs/op",
            "extra": "18622518 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda98TypeCodeNil - ns/op",
            "value": 63,
            "unit": "ns/op",
            "extra": "18622518 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda98TypeCodeNil - B/op",
            "value": 64,
            "unit": "B/op",
            "extra": "18622518 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda98TypeCodeNil - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "18622518 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda99Parse",
            "value": 974.9,
            "unit": "ns/op\t     200 B/op\t      10 allocs/op",
            "extra": "1234621 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda99Parse - ns/op",
            "value": 974.9,
            "unit": "ns/op",
            "extra": "1234621 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda99Parse - B/op",
            "value": 200,
            "unit": "B/op",
            "extra": "1234621 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda99Parse - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "1234621 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda99String",
            "value": 1169,
            "unit": "ns/op\t     296 B/op\t      11 allocs/op",
            "extra": "957730 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda99String - ns/op",
            "value": 1169,
            "unit": "ns/op",
            "extra": "957730 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda99String - B/op",
            "value": 296,
            "unit": "B/op",
            "extra": "957730 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda99String - allocs/op",
            "value": 11,
            "unit": "allocs/op",
            "extra": "957730 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda99MakeReturnCodeDict",
            "value": 7169,
            "unit": "ns/op\t   11469 B/op\t      10 allocs/op",
            "extra": "164106 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda99MakeReturnCodeDict - ns/op",
            "value": 7169,
            "unit": "ns/op",
            "extra": "164106 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda99MakeReturnCodeDict - B/op",
            "value": 11469,
            "unit": "B/op",
            "extra": "164106 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda99MakeReturnCodeDict - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "164106 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda99ValidateTrue",
            "value": 18.49,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "65228493 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda99ValidateTrue - ns/op",
            "value": 18.49,
            "unit": "ns/op",
            "extra": "65228493 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda99ValidateTrue - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "65228493 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda99ValidateTrue - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "65228493 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda99ValidateReturnCodeFalse",
            "value": 78.99,
            "unit": "ns/op\t      64 B/op\t       1 allocs/op",
            "extra": "15111051 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda99ValidateReturnCodeFalse - ns/op",
            "value": 78.99,
            "unit": "ns/op",
            "extra": "15111051 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda99ValidateReturnCodeFalse - B/op",
            "value": 64,
            "unit": "B/op",
            "extra": "15111051 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda99ValidateReturnCodeFalse - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "15111051 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda99OriginalTraceField",
            "value": 50.41,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "23914306 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda99OriginalTraceField - ns/op",
            "value": 50.41,
            "unit": "ns/op",
            "extra": "23914306 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda99OriginalTraceField - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "23914306 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda99OriginalTraceField - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "23914306 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda99DateOfDeathField",
            "value": 23.62,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "51215217 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda99DateOfDeathField - ns/op",
            "value": 23.62,
            "unit": "ns/op",
            "extra": "51215217 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda99DateOfDeathField - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "51215217 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda99DateOfDeathField - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "51215217 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda99OriginalDFIField",
            "value": 46.09,
            "unit": "ns/op\t       8 B/op\t       1 allocs/op",
            "extra": "25274578 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda99OriginalDFIField - ns/op",
            "value": 46.09,
            "unit": "ns/op",
            "extra": "25274578 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda99OriginalDFIField - B/op",
            "value": 8,
            "unit": "B/op",
            "extra": "25274578 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda99OriginalDFIField - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "25274578 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda99AddendaInformationField",
            "value": 69.06,
            "unit": "ns/op\t      48 B/op\t       1 allocs/op",
            "extra": "17142194 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda99AddendaInformationField - ns/op",
            "value": 69.06,
            "unit": "ns/op",
            "extra": "17142194 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda99AddendaInformationField - B/op",
            "value": 48,
            "unit": "B/op",
            "extra": "17142194 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda99AddendaInformationField - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "17142194 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda99TraceNumberField",
            "value": 58.78,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "20230592 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda99TraceNumberField - ns/op",
            "value": 58.78,
            "unit": "ns/op",
            "extra": "20230592 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda99TraceNumberField - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "20230592 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda99TraceNumberField - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "20230592 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda99TypeCode99",
            "value": 89.55,
            "unit": "ns/op\t      80 B/op\t       2 allocs/op",
            "extra": "13204461 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda99TypeCode99 - ns/op",
            "value": 89.55,
            "unit": "ns/op",
            "extra": "13204461 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda99TypeCode99 - B/op",
            "value": 80,
            "unit": "B/op",
            "extra": "13204461 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda99TypeCode99 - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "13204461 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda99TypeCodeNil",
            "value": 63.29,
            "unit": "ns/op\t      64 B/op\t       1 allocs/op",
            "extra": "18721712 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda99TypeCodeNil - ns/op",
            "value": 63.29,
            "unit": "ns/op",
            "extra": "18721712 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda99TypeCodeNil - B/op",
            "value": 64,
            "unit": "B/op",
            "extra": "18721712 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda99TypeCodeNil - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "18721712 times\n4 procs"
          },
          {
            "name": "BenchmarkMockADVBatchControl",
            "value": 18.4,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "65303958 times\n4 procs"
          },
          {
            "name": "BenchmarkMockADVBatchControl - ns/op",
            "value": 18.4,
            "unit": "ns/op",
            "extra": "65303958 times\n4 procs"
          },
          {
            "name": "BenchmarkMockADVBatchControl - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "65303958 times\n4 procs"
          },
          {
            "name": "BenchmarkMockADVBatchControl - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "65303958 times\n4 procs"
          },
          {
            "name": "BenchmarkParseADVBatchControl",
            "value": 5913,
            "unit": "ns/op\t   15532 B/op\t      29 allocs/op",
            "extra": "205069 times\n4 procs"
          },
          {
            "name": "BenchmarkParseADVBatchControl - ns/op",
            "value": 5913,
            "unit": "ns/op",
            "extra": "205069 times\n4 procs"
          },
          {
            "name": "BenchmarkParseADVBatchControl - B/op",
            "value": 15532,
            "unit": "B/op",
            "extra": "205069 times\n4 procs"
          },
          {
            "name": "BenchmarkParseADVBatchControl - allocs/op",
            "value": 29,
            "unit": "allocs/op",
            "extra": "205069 times\n4 procs"
          },
          {
            "name": "BenchmarkADVBCString",
            "value": 6062,
            "unit": "ns/op\t   15648 B/op\t      31 allocs/op",
            "extra": "196341 times\n4 procs"
          },
          {
            "name": "BenchmarkADVBCString - ns/op",
            "value": 6062,
            "unit": "ns/op",
            "extra": "196341 times\n4 procs"
          },
          {
            "name": "BenchmarkADVBCString - B/op",
            "value": 15648,
            "unit": "B/op",
            "extra": "196341 times\n4 procs"
          },
          {
            "name": "BenchmarkADVBCString - allocs/op",
            "value": 31,
            "unit": "allocs/op",
            "extra": "196341 times\n4 procs"
          },
          {
            "name": "BenchmarkADVBCisServiceClassErr",
            "value": 117.5,
            "unit": "ns/op\t      83 B/op\t       3 allocs/op",
            "extra": "10203698 times\n4 procs"
          },
          {
            "name": "BenchmarkADVBCisServiceClassErr - ns/op",
            "value": 117.5,
            "unit": "ns/op",
            "extra": "10203698 times\n4 procs"
          },
          {
            "name": "BenchmarkADVBCisServiceClassErr - B/op",
            "value": 83,
            "unit": "B/op",
            "extra": "10203698 times\n4 procs"
          },
          {
            "name": "BenchmarkADVBCisServiceClassErr - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "10203698 times\n4 procs"
          },
          {
            "name": "BenchmarkADVBCBatchNumber",
            "value": 19.01,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "63099778 times\n4 procs"
          },
          {
            "name": "BenchmarkADVBCBatchNumber - ns/op",
            "value": 19.01,
            "unit": "ns/op",
            "extra": "63099778 times\n4 procs"
          },
          {
            "name": "BenchmarkADVBCBatchNumber - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "63099778 times\n4 procs"
          },
          {
            "name": "BenchmarkADVBCBatchNumber - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "63099778 times\n4 procs"
          },
          {
            "name": "BenchmarkADVACHOperatorDataAlphaNumeric",
            "value": 309.5,
            "unit": "ns/op\t     160 B/op\t       4 allocs/op",
            "extra": "3889466 times\n4 procs"
          },
          {
            "name": "BenchmarkADVACHOperatorDataAlphaNumeric - ns/op",
            "value": 309.5,
            "unit": "ns/op",
            "extra": "3889466 times\n4 procs"
          },
          {
            "name": "BenchmarkADVACHOperatorDataAlphaNumeric - B/op",
            "value": 160,
            "unit": "B/op",
            "extra": "3889466 times\n4 procs"
          },
          {
            "name": "BenchmarkADVACHOperatorDataAlphaNumeric - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "3889466 times\n4 procs"
          },
          {
            "name": "BenchmarkADVBCFieldInclusionServiceClassCode",
            "value": 96.95,
            "unit": "ns/op\t      80 B/op\t       2 allocs/op",
            "extra": "12279691 times\n4 procs"
          },
          {
            "name": "BenchmarkADVBCFieldInclusionServiceClassCode - ns/op",
            "value": 96.95,
            "unit": "ns/op",
            "extra": "12279691 times\n4 procs"
          },
          {
            "name": "BenchmarkADVBCFieldInclusionServiceClassCode - B/op",
            "value": 80,
            "unit": "B/op",
            "extra": "12279691 times\n4 procs"
          },
          {
            "name": "BenchmarkADVBCFieldInclusionServiceClassCode - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "12279691 times\n4 procs"
          },
          {
            "name": "BenchmarkADVBCFieldInclusionODFIIdentification",
            "value": 104.2,
            "unit": "ns/op\t      80 B/op\t       2 allocs/op",
            "extra": "11584516 times\n4 procs"
          },
          {
            "name": "BenchmarkADVBCFieldInclusionODFIIdentification - ns/op",
            "value": 104.2,
            "unit": "ns/op",
            "extra": "11584516 times\n4 procs"
          },
          {
            "name": "BenchmarkADVBCFieldInclusionODFIIdentification - B/op",
            "value": 80,
            "unit": "B/op",
            "extra": "11584516 times\n4 procs"
          },
          {
            "name": "BenchmarkADVBCFieldInclusionODFIIdentification - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "11584516 times\n4 procs"
          },
          {
            "name": "BenchmarkADVBatchControlLength",
            "value": 453.6,
            "unit": "ns/op\t     192 B/op\t       8 allocs/op",
            "extra": "2640439 times\n4 procs"
          },
          {
            "name": "BenchmarkADVBatchControlLength - ns/op",
            "value": 453.6,
            "unit": "ns/op",
            "extra": "2640439 times\n4 procs"
          },
          {
            "name": "BenchmarkADVBatchControlLength - B/op",
            "value": 192,
            "unit": "B/op",
            "extra": "2640439 times\n4 procs"
          },
          {
            "name": "BenchmarkADVBatchControlLength - allocs/op",
            "value": 8,
            "unit": "allocs/op",
            "extra": "2640439 times\n4 procs"
          },
          {
            "name": "BenchmarkMockADVEntryDetail",
            "value": 238.3,
            "unit": "ns/op\t     240 B/op\t       1 allocs/op",
            "extra": "5065575 times\n4 procs"
          },
          {
            "name": "BenchmarkMockADVEntryDetail - ns/op",
            "value": 238.3,
            "unit": "ns/op",
            "extra": "5065575 times\n4 procs"
          },
          {
            "name": "BenchmarkMockADVEntryDetail - B/op",
            "value": 240,
            "unit": "B/op",
            "extra": "5065575 times\n4 procs"
          },
          {
            "name": "BenchmarkMockADVEntryDetail - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "5065575 times\n4 procs"
          },
          {
            "name": "BenchmarkADVEDString",
            "value": 6611,
            "unit": "ns/op\t   16320 B/op\t      39 allocs/op",
            "extra": "174934 times\n4 procs"
          },
          {
            "name": "BenchmarkADVEDString - ns/op",
            "value": 6611,
            "unit": "ns/op",
            "extra": "174934 times\n4 procs"
          },
          {
            "name": "BenchmarkADVEDString - B/op",
            "value": 16320,
            "unit": "B/op",
            "extra": "174934 times\n4 procs"
          },
          {
            "name": "BenchmarkADVEDString - allocs/op",
            "value": 39,
            "unit": "allocs/op",
            "extra": "174934 times\n4 procs"
          },
          {
            "name": "BenchmarkMockADVFileControl",
            "value": 5.939,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "202427846 times\n4 procs"
          },
          {
            "name": "BenchmarkMockADVFileControl - ns/op",
            "value": 5.939,
            "unit": "ns/op",
            "extra": "202427846 times\n4 procs"
          },
          {
            "name": "BenchmarkMockADVFileControl - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "202427846 times\n4 procs"
          },
          {
            "name": "BenchmarkMockADVFileControl - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "202427846 times\n4 procs"
          },
          {
            "name": "BenchmarkParseADVFileControl",
            "value": 6987,
            "unit": "ns/op\t   16136 B/op\t      32 allocs/op",
            "extra": "170492 times\n4 procs"
          },
          {
            "name": "BenchmarkParseADVFileControl - ns/op",
            "value": 6987,
            "unit": "ns/op",
            "extra": "170492 times\n4 procs"
          },
          {
            "name": "BenchmarkParseADVFileControl - B/op",
            "value": 16136,
            "unit": "B/op",
            "extra": "170492 times\n4 procs"
          },
          {
            "name": "BenchmarkParseADVFileControl - allocs/op",
            "value": 32,
            "unit": "allocs/op",
            "extra": "170492 times\n4 procs"
          },
          {
            "name": "BenchmarkADVFCString",
            "value": 7222,
            "unit": "ns/op\t   16236 B/op\t      33 allocs/op",
            "extra": "168116 times\n4 procs"
          },
          {
            "name": "BenchmarkADVFCString - ns/op",
            "value": 7222,
            "unit": "ns/op",
            "extra": "168116 times\n4 procs"
          },
          {
            "name": "BenchmarkADVFCString - B/op",
            "value": 16236,
            "unit": "B/op",
            "extra": "168116 times\n4 procs"
          },
          {
            "name": "BenchmarkADVFCString - allocs/op",
            "value": 33,
            "unit": "allocs/op",
            "extra": "168116 times\n4 procs"
          },
          {
            "name": "BenchmarkADVFCFieldInclusion",
            "value": 129.3,
            "unit": "ns/op\t      88 B/op\t       3 allocs/op",
            "extra": "9184525 times\n4 procs"
          },
          {
            "name": "BenchmarkADVFCFieldInclusion - ns/op",
            "value": 129.3,
            "unit": "ns/op",
            "extra": "9184525 times\n4 procs"
          },
          {
            "name": "BenchmarkADVFCFieldInclusion - B/op",
            "value": 88,
            "unit": "B/op",
            "extra": "9184525 times\n4 procs"
          },
          {
            "name": "BenchmarkADVFCFieldInclusion - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "9184525 times\n4 procs"
          },
          {
            "name": "BenchmarkADVFCFieldInclusionBlockCount",
            "value": 127.2,
            "unit": "ns/op\t      88 B/op\t       3 allocs/op",
            "extra": "9351590 times\n4 procs"
          },
          {
            "name": "BenchmarkADVFCFieldInclusionBlockCount - ns/op",
            "value": 127.2,
            "unit": "ns/op",
            "extra": "9351590 times\n4 procs"
          },
          {
            "name": "BenchmarkADVFCFieldInclusionBlockCount - B/op",
            "value": 88,
            "unit": "B/op",
            "extra": "9351590 times\n4 procs"
          },
          {
            "name": "BenchmarkADVFCFieldInclusionBlockCount - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "9351590 times\n4 procs"
          },
          {
            "name": "BenchmarkADVFCFieldInclusionEntryAddendaCount",
            "value": 127.7,
            "unit": "ns/op\t      88 B/op\t       3 allocs/op",
            "extra": "9322231 times\n4 procs"
          },
          {
            "name": "BenchmarkADVFCFieldInclusionEntryAddendaCount - ns/op",
            "value": 127.7,
            "unit": "ns/op",
            "extra": "9322231 times\n4 procs"
          },
          {
            "name": "BenchmarkADVFCFieldInclusionEntryAddendaCount - B/op",
            "value": 88,
            "unit": "B/op",
            "extra": "9322231 times\n4 procs"
          },
          {
            "name": "BenchmarkADVFCFieldInclusionEntryAddendaCount - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "9322231 times\n4 procs"
          },
          {
            "name": "BenchmarkADVFCFieldInclusionEntryHash",
            "value": 138.8,
            "unit": "ns/op\t      96 B/op\t       3 allocs/op",
            "extra": "8605002 times\n4 procs"
          },
          {
            "name": "BenchmarkADVFCFieldInclusionEntryHash - ns/op",
            "value": 138.8,
            "unit": "ns/op",
            "extra": "8605002 times\n4 procs"
          },
          {
            "name": "BenchmarkADVFCFieldInclusionEntryHash - B/op",
            "value": 96,
            "unit": "B/op",
            "extra": "8605002 times\n4 procs"
          },
          {
            "name": "BenchmarkADVFCFieldInclusionEntryHash - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "8605002 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchACKHeader",
            "value": 163.8,
            "unit": "ns/op\t     448 B/op\t       3 allocs/op",
            "extra": "7329343 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchACKHeader - ns/op",
            "value": 163.8,
            "unit": "ns/op",
            "extra": "7329343 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchACKHeader - B/op",
            "value": 448,
            "unit": "B/op",
            "extra": "7329343 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchACKHeader - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "7329343 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchACKAddendumCount",
            "value": 2922,
            "unit": "ns/op\t    1440 B/op\t      20 allocs/op",
            "extra": "386179 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchACKAddendumCount - ns/op",
            "value": 2922,
            "unit": "ns/op",
            "extra": "386179 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchACKAddendumCount - B/op",
            "value": 1440,
            "unit": "B/op",
            "extra": "386179 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchACKAddendumCount - allocs/op",
            "value": 20,
            "unit": "allocs/op",
            "extra": "386179 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchACKReceivingCompanyName",
            "value": 1580,
            "unit": "ns/op\t    1104 B/op\t      13 allocs/op",
            "extra": "704678 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchACKReceivingCompanyName - ns/op",
            "value": 1580,
            "unit": "ns/op",
            "extra": "704678 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchACKReceivingCompanyName - B/op",
            "value": 1104,
            "unit": "B/op",
            "extra": "704678 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchACKReceivingCompanyName - allocs/op",
            "value": 13,
            "unit": "allocs/op",
            "extra": "704678 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchACKAddendaTypeCode",
            "value": 1758,
            "unit": "ns/op\t    1120 B/op\t      14 allocs/op",
            "extra": "651459 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchACKAddendaTypeCode - ns/op",
            "value": 1758,
            "unit": "ns/op",
            "extra": "651459 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchACKAddendaTypeCode - B/op",
            "value": 1120,
            "unit": "B/op",
            "extra": "651459 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchACKAddendaTypeCode - allocs/op",
            "value": 14,
            "unit": "allocs/op",
            "extra": "651459 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchACKSEC",
            "value": 1918,
            "unit": "ns/op\t    1048 B/op\t      13 allocs/op",
            "extra": "593350 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchACKSEC - ns/op",
            "value": 1918,
            "unit": "ns/op",
            "extra": "593350 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchACKSEC - B/op",
            "value": 1048,
            "unit": "B/op",
            "extra": "593350 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchACKSEC - allocs/op",
            "value": 13,
            "unit": "allocs/op",
            "extra": "593350 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchACKAddendaCount",
            "value": 2951,
            "unit": "ns/op\t    1440 B/op\t      20 allocs/op",
            "extra": "394332 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchACKAddendaCount - ns/op",
            "value": 2951,
            "unit": "ns/op",
            "extra": "394332 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchACKAddendaCount - B/op",
            "value": 1440,
            "unit": "B/op",
            "extra": "394332 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchACKAddendaCount - allocs/op",
            "value": 20,
            "unit": "allocs/op",
            "extra": "394332 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchACKServiceClassCode",
            "value": 1480,
            "unit": "ns/op\t    1040 B/op\t      13 allocs/op",
            "extra": "734703 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchACKServiceClassCode - ns/op",
            "value": 1480,
            "unit": "ns/op",
            "extra": "734703 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchACKServiceClassCode - B/op",
            "value": 1040,
            "unit": "B/op",
            "extra": "734703 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchACKServiceClassCode - allocs/op",
            "value": 13,
            "unit": "allocs/op",
            "extra": "734703 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchACKReceivingCompanyField",
            "value": 1452,
            "unit": "ns/op\t     984 B/op\t      12 allocs/op",
            "extra": "752770 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchACKReceivingCompanyField - ns/op",
            "value": 1452,
            "unit": "ns/op",
            "extra": "752770 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchACKReceivingCompanyField - B/op",
            "value": 984,
            "unit": "B/op",
            "extra": "752770 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchACKReceivingCompanyField - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "752770 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchADVHeader",
            "value": 343.6,
            "unit": "ns/op\t     440 B/op\t       4 allocs/op",
            "extra": "3496394 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchADVHeader - ns/op",
            "value": 343.6,
            "unit": "ns/op",
            "extra": "3496394 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchADVHeader - B/op",
            "value": 440,
            "unit": "B/op",
            "extra": "3496394 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchADVHeader - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "3496394 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchADVSEC",
            "value": 1439,
            "unit": "ns/op\t     880 B/op\t       8 allocs/op",
            "extra": "780776 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchADVSEC - ns/op",
            "value": 1439,
            "unit": "ns/op",
            "extra": "780776 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchADVSEC - B/op",
            "value": 880,
            "unit": "B/op",
            "extra": "780776 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchADVSEC - allocs/op",
            "value": 8,
            "unit": "allocs/op",
            "extra": "780776 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchADVServiceClassCode",
            "value": 1585,
            "unit": "ns/op\t     992 B/op\t       9 allocs/op",
            "extra": "661899 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchADVServiceClassCode - ns/op",
            "value": 1585,
            "unit": "ns/op",
            "extra": "661899 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchADVServiceClassCode - B/op",
            "value": 992,
            "unit": "B/op",
            "extra": "661899 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchADVServiceClassCode - allocs/op",
            "value": 9,
            "unit": "allocs/op",
            "extra": "661899 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchARCHeader",
            "value": 164.7,
            "unit": "ns/op\t     448 B/op\t       3 allocs/op",
            "extra": "7228138 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchARCHeader - ns/op",
            "value": 164.7,
            "unit": "ns/op",
            "extra": "7228138 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchARCHeader - B/op",
            "value": 448,
            "unit": "B/op",
            "extra": "7228138 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchARCHeader - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "7228138 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchARCCreate",
            "value": 1943,
            "unit": "ns/op\t     992 B/op\t       9 allocs/op",
            "extra": "583821 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchARCCreate - ns/op",
            "value": 1943,
            "unit": "ns/op",
            "extra": "583821 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchARCCreate - B/op",
            "value": 992,
            "unit": "B/op",
            "extra": "583821 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchARCCreate - allocs/op",
            "value": 9,
            "unit": "allocs/op",
            "extra": "583821 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchARCStandardEntryClassCode",
            "value": 2051,
            "unit": "ns/op\t    1072 B/op\t      10 allocs/op",
            "extra": "561334 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchARCStandardEntryClassCode - ns/op",
            "value": 2051,
            "unit": "ns/op",
            "extra": "561334 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchARCStandardEntryClassCode - B/op",
            "value": 1072,
            "unit": "B/op",
            "extra": "561334 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchARCStandardEntryClassCode - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "561334 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchARCServiceClassCodeEquality",
            "value": 2127,
            "unit": "ns/op\t    1088 B/op\t      12 allocs/op",
            "extra": "537276 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchARCServiceClassCodeEquality - ns/op",
            "value": 2127,
            "unit": "ns/op",
            "extra": "537276 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchARCServiceClassCodeEquality - B/op",
            "value": 1088,
            "unit": "B/op",
            "extra": "537276 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchARCServiceClassCodeEquality - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "537276 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchARCMixedCreditsAndDebits",
            "value": 2152,
            "unit": "ns/op\t    1088 B/op\t      12 allocs/op",
            "extra": "525685 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchARCMixedCreditsAndDebits - ns/op",
            "value": 2152,
            "unit": "ns/op",
            "extra": "525685 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchARCMixedCreditsAndDebits - B/op",
            "value": 1088,
            "unit": "B/op",
            "extra": "525685 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchARCMixedCreditsAndDebits - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "525685 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchARCCreditsOnly",
            "value": 2160,
            "unit": "ns/op\t    1088 B/op\t      12 allocs/op",
            "extra": "523891 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchARCCreditsOnly - ns/op",
            "value": 2160,
            "unit": "ns/op",
            "extra": "523891 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchARCCreditsOnly - B/op",
            "value": 1088,
            "unit": "B/op",
            "extra": "523891 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchARCCreditsOnly - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "523891 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchARCAutomatedAccountingAdvices",
            "value": 2155,
            "unit": "ns/op\t    1096 B/op\t      13 allocs/op",
            "extra": "533875 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchARCAutomatedAccountingAdvices - ns/op",
            "value": 2155,
            "unit": "ns/op",
            "extra": "533875 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchARCAutomatedAccountingAdvices - B/op",
            "value": 1096,
            "unit": "B/op",
            "extra": "533875 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchARCAutomatedAccountingAdvices - allocs/op",
            "value": 13,
            "unit": "allocs/op",
            "extra": "533875 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchARCAmount",
            "value": 2640,
            "unit": "ns/op\t    1296 B/op\t      17 allocs/op",
            "extra": "429318 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchARCAmount - ns/op",
            "value": 2640,
            "unit": "ns/op",
            "extra": "429318 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchARCAmount - B/op",
            "value": 1296,
            "unit": "B/op",
            "extra": "429318 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchARCAmount - allocs/op",
            "value": 17,
            "unit": "allocs/op",
            "extra": "429318 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchARCCheckSerialNumber",
            "value": 1770,
            "unit": "ns/op\t     944 B/op\t       9 allocs/op",
            "extra": "638433 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchARCCheckSerialNumber - ns/op",
            "value": 1770,
            "unit": "ns/op",
            "extra": "638433 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchARCCheckSerialNumber - B/op",
            "value": 944,
            "unit": "B/op",
            "extra": "638433 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchARCCheckSerialNumber - allocs/op",
            "value": 9,
            "unit": "allocs/op",
            "extra": "638433 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchARCTransactionCode",
            "value": 1107,
            "unit": "ns/op\t     944 B/op\t       9 allocs/op",
            "extra": "957127 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchARCTransactionCode - ns/op",
            "value": 1107,
            "unit": "ns/op",
            "extra": "957127 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchARCTransactionCode - B/op",
            "value": 944,
            "unit": "B/op",
            "extra": "957127 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchARCTransactionCode - allocs/op",
            "value": 9,
            "unit": "allocs/op",
            "extra": "957127 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchARCAddendaCount",
            "value": 2381,
            "unit": "ns/op\t    1184 B/op\t      14 allocs/op",
            "extra": "479074 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchARCAddendaCount - ns/op",
            "value": 2381,
            "unit": "ns/op",
            "extra": "479074 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchARCAddendaCount - B/op",
            "value": 1184,
            "unit": "B/op",
            "extra": "479074 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchARCAddendaCount - allocs/op",
            "value": 14,
            "unit": "allocs/op",
            "extra": "479074 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchARCInvalidBuild",
            "value": 1471,
            "unit": "ns/op\t     928 B/op\t       9 allocs/op",
            "extra": "759771 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchARCInvalidBuild - ns/op",
            "value": 1471,
            "unit": "ns/op",
            "extra": "759771 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchARCInvalidBuild - B/op",
            "value": 928,
            "unit": "B/op",
            "extra": "759771 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchARCInvalidBuild - allocs/op",
            "value": 9,
            "unit": "allocs/op",
            "extra": "759771 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchATXHeader",
            "value": 164.9,
            "unit": "ns/op\t     448 B/op\t       3 allocs/op",
            "extra": "7213524 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchATXHeader - ns/op",
            "value": 164.9,
            "unit": "ns/op",
            "extra": "7213524 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchATXHeader - B/op",
            "value": 448,
            "unit": "B/op",
            "extra": "7213524 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchATXHeader - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "7213524 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchATXCreate",
            "value": 2189,
            "unit": "ns/op\t    1040 B/op\t      16 allocs/op",
            "extra": "520737 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchATXCreate - ns/op",
            "value": 2189,
            "unit": "ns/op",
            "extra": "520737 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchATXCreate - B/op",
            "value": 1040,
            "unit": "B/op",
            "extra": "520737 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchATXCreate - allocs/op",
            "value": 16,
            "unit": "allocs/op",
            "extra": "520737 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchATXStandardEntryClassCode",
            "value": 2512,
            "unit": "ns/op\t    1248 B/op\t      18 allocs/op",
            "extra": "463618 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchATXStandardEntryClassCode - ns/op",
            "value": 2512,
            "unit": "ns/op",
            "extra": "463618 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchATXStandardEntryClassCode - B/op",
            "value": 1248,
            "unit": "B/op",
            "extra": "463618 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchATXStandardEntryClassCode - allocs/op",
            "value": 18,
            "unit": "allocs/op",
            "extra": "463618 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchATXServiceClassCodeEquality",
            "value": 2442,
            "unit": "ns/op\t    1256 B/op\t      19 allocs/op",
            "extra": "471484 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchATXServiceClassCodeEquality - ns/op",
            "value": 2442,
            "unit": "ns/op",
            "extra": "471484 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchATXServiceClassCodeEquality - B/op",
            "value": 1256,
            "unit": "B/op",
            "extra": "471484 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchATXServiceClassCodeEquality - allocs/op",
            "value": 19,
            "unit": "allocs/op",
            "extra": "471484 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchATXAddendaCount",
            "value": 3284,
            "unit": "ns/op\t    1512 B/op\t      24 allocs/op",
            "extra": "353768 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchATXAddendaCount - ns/op",
            "value": 3284,
            "unit": "ns/op",
            "extra": "353768 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchATXAddendaCount - B/op",
            "value": 1512,
            "unit": "B/op",
            "extra": "353768 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchATXAddendaCount - allocs/op",
            "value": 24,
            "unit": "allocs/op",
            "extra": "353768 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchATXAddendaCountZero",
            "value": 1604,
            "unit": "ns/op\t     904 B/op\t      14 allocs/op",
            "extra": "710245 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchATXAddendaCountZero - ns/op",
            "value": 1604,
            "unit": "ns/op",
            "extra": "710245 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchATXAddendaCountZero - B/op",
            "value": 904,
            "unit": "B/op",
            "extra": "710245 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchATXAddendaCountZero - allocs/op",
            "value": 14,
            "unit": "allocs/op",
            "extra": "710245 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchATXInvalidAddenda",
            "value": 1487,
            "unit": "ns/op\t    1152 B/op\t      18 allocs/op",
            "extra": "743017 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchATXInvalidAddenda - ns/op",
            "value": 1487,
            "unit": "ns/op",
            "extra": "743017 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchATXInvalidAddenda - B/op",
            "value": 1152,
            "unit": "B/op",
            "extra": "743017 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchATXInvalidAddenda - allocs/op",
            "value": 18,
            "unit": "allocs/op",
            "extra": "743017 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchATXInvalidBuild",
            "value": 1741,
            "unit": "ns/op\t    1096 B/op\t      16 allocs/op",
            "extra": "651858 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchATXInvalidBuild - ns/op",
            "value": 1741,
            "unit": "ns/op",
            "extra": "651858 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchATXInvalidBuild - B/op",
            "value": 1096,
            "unit": "B/op",
            "extra": "651858 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchATXInvalidBuild - allocs/op",
            "value": 16,
            "unit": "allocs/op",
            "extra": "651858 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchATXAddenda10000",
            "value": 2009473,
            "unit": "ns/op\t 1191843 B/op\t   20039 allocs/op",
            "extra": "598 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchATXAddenda10000 - ns/op",
            "value": 2009473,
            "unit": "ns/op",
            "extra": "598 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchATXAddenda10000 - B/op",
            "value": 1191843,
            "unit": "B/op",
            "extra": "598 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchATXAddenda10000 - allocs/op",
            "value": 20039,
            "unit": "allocs/op",
            "extra": "598 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchATXAddendaRecords",
            "value": 106547,
            "unit": "ns/op\t   59864 B/op\t    1152 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchATXAddendaRecords - ns/op",
            "value": 106547,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchATXAddendaRecords - B/op",
            "value": 59864,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchATXAddendaRecords - allocs/op",
            "value": 1152,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchATXReceivingCompany",
            "value": 2466,
            "unit": "ns/op\t    1072 B/op\t      17 allocs/op",
            "extra": "458428 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchATXReceivingCompany - ns/op",
            "value": 2466,
            "unit": "ns/op",
            "extra": "458428 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchATXReceivingCompany - B/op",
            "value": 1072,
            "unit": "B/op",
            "extra": "458428 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchATXReceivingCompany - allocs/op",
            "value": 17,
            "unit": "allocs/op",
            "extra": "458428 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchATXReserved",
            "value": 1655,
            "unit": "ns/op\t    1032 B/op\t      15 allocs/op",
            "extra": "657079 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchATXReserved - ns/op",
            "value": 1655,
            "unit": "ns/op",
            "extra": "657079 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchATXReserved - B/op",
            "value": 1032,
            "unit": "B/op",
            "extra": "657079 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchATXReserved - allocs/op",
            "value": 15,
            "unit": "allocs/op",
            "extra": "657079 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchATXZeroAddendaRecords",
            "value": 1608,
            "unit": "ns/op\t     904 B/op\t      14 allocs/op",
            "extra": "699469 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchATXZeroAddendaRecords - ns/op",
            "value": 1608,
            "unit": "ns/op",
            "extra": "699469 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchATXZeroAddendaRecords - B/op",
            "value": 904,
            "unit": "B/op",
            "extra": "699469 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchATXZeroAddendaRecords - allocs/op",
            "value": 14,
            "unit": "allocs/op",
            "extra": "699469 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchATXTransactionCode",
            "value": 1614,
            "unit": "ns/op\t     984 B/op\t      15 allocs/op",
            "extra": "676844 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchATXTransactionCode - ns/op",
            "value": 1614,
            "unit": "ns/op",
            "extra": "676844 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchATXTransactionCode - B/op",
            "value": 984,
            "unit": "B/op",
            "extra": "676844 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchATXTransactionCode - allocs/op",
            "value": 15,
            "unit": "allocs/op",
            "extra": "676844 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchBOCHeader",
            "value": 171.2,
            "unit": "ns/op\t     448 B/op\t       3 allocs/op",
            "extra": "7112397 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchBOCHeader - ns/op",
            "value": 171.2,
            "unit": "ns/op",
            "extra": "7112397 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchBOCHeader - B/op",
            "value": 448,
            "unit": "B/op",
            "extra": "7112397 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchBOCHeader - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "7112397 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchBOCCreate",
            "value": 1987,
            "unit": "ns/op\t     992 B/op\t       9 allocs/op",
            "extra": "561014 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchBOCCreate - ns/op",
            "value": 1987,
            "unit": "ns/op",
            "extra": "561014 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchBOCCreate - B/op",
            "value": 992,
            "unit": "B/op",
            "extra": "561014 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchBOCCreate - allocs/op",
            "value": 9,
            "unit": "allocs/op",
            "extra": "561014 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchBOCStandardEntryClassCode",
            "value": 2047,
            "unit": "ns/op\t    1072 B/op\t      10 allocs/op",
            "extra": "540192 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchBOCStandardEntryClassCode - ns/op",
            "value": 2047,
            "unit": "ns/op",
            "extra": "540192 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchBOCStandardEntryClassCode - B/op",
            "value": 1072,
            "unit": "B/op",
            "extra": "540192 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchBOCStandardEntryClassCode - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "540192 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchBOCServiceClassCodeEquality",
            "value": 2150,
            "unit": "ns/op\t    1088 B/op\t      12 allocs/op",
            "extra": "526233 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchBOCServiceClassCodeEquality - ns/op",
            "value": 2150,
            "unit": "ns/op",
            "extra": "526233 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchBOCServiceClassCodeEquality - B/op",
            "value": 1088,
            "unit": "B/op",
            "extra": "526233 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchBOCServiceClassCodeEquality - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "526233 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchBOCMixedCreditsAndDebits",
            "value": 2153,
            "unit": "ns/op\t    1088 B/op\t      12 allocs/op",
            "extra": "527402 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchBOCMixedCreditsAndDebits - ns/op",
            "value": 2153,
            "unit": "ns/op",
            "extra": "527402 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchBOCMixedCreditsAndDebits - B/op",
            "value": 1088,
            "unit": "B/op",
            "extra": "527402 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchBOCMixedCreditsAndDebits - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "527402 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchBOCCreditsOnly",
            "value": 2164,
            "unit": "ns/op\t    1088 B/op\t      12 allocs/op",
            "extra": "530757 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchBOCCreditsOnly - ns/op",
            "value": 2164,
            "unit": "ns/op",
            "extra": "530757 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchBOCCreditsOnly - B/op",
            "value": 1088,
            "unit": "B/op",
            "extra": "530757 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchBOCCreditsOnly - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "530757 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchBOCAutomatedAccountingAdvices",
            "value": 2174,
            "unit": "ns/op\t    1096 B/op\t      13 allocs/op",
            "extra": "529060 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchBOCAutomatedAccountingAdvices - ns/op",
            "value": 2174,
            "unit": "ns/op",
            "extra": "529060 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchBOCAutomatedAccountingAdvices - B/op",
            "value": 1096,
            "unit": "B/op",
            "extra": "529060 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchBOCAutomatedAccountingAdvices - allocs/op",
            "value": 13,
            "unit": "allocs/op",
            "extra": "529060 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchBOCAmount",
            "value": 2666,
            "unit": "ns/op\t    1296 B/op\t      17 allocs/op",
            "extra": "428738 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchBOCAmount - ns/op",
            "value": 2666,
            "unit": "ns/op",
            "extra": "428738 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchBOCAmount - B/op",
            "value": 1296,
            "unit": "B/op",
            "extra": "428738 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchBOCAmount - allocs/op",
            "value": 17,
            "unit": "allocs/op",
            "extra": "428738 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchBOCCheckSerialNumber",
            "value": 1798,
            "unit": "ns/op\t     944 B/op\t       9 allocs/op",
            "extra": "626514 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchBOCCheckSerialNumber - ns/op",
            "value": 1798,
            "unit": "ns/op",
            "extra": "626514 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchBOCCheckSerialNumber - B/op",
            "value": 944,
            "unit": "B/op",
            "extra": "626514 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchBOCCheckSerialNumber - allocs/op",
            "value": 9,
            "unit": "allocs/op",
            "extra": "626514 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchBOCTransactionCode",
            "value": 1129,
            "unit": "ns/op\t     944 B/op\t       9 allocs/op",
            "extra": "973654 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchBOCTransactionCode - ns/op",
            "value": 1129,
            "unit": "ns/op",
            "extra": "973654 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchBOCTransactionCode - B/op",
            "value": 944,
            "unit": "B/op",
            "extra": "973654 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchBOCTransactionCode - allocs/op",
            "value": 9,
            "unit": "allocs/op",
            "extra": "973654 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchBOCAddenda05",
            "value": 2403,
            "unit": "ns/op\t    1184 B/op\t      14 allocs/op",
            "extra": "485020 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchBOCAddenda05 - ns/op",
            "value": 2403,
            "unit": "ns/op",
            "extra": "485020 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchBOCAddenda05 - B/op",
            "value": 1184,
            "unit": "B/op",
            "extra": "485020 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchBOCAddenda05 - allocs/op",
            "value": 14,
            "unit": "allocs/op",
            "extra": "485020 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchBOCInvalidBuild",
            "value": 1496,
            "unit": "ns/op\t     928 B/op\t       9 allocs/op",
            "extra": "758239 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchBOCInvalidBuild - ns/op",
            "value": 1496,
            "unit": "ns/op",
            "extra": "758239 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchBOCInvalidBuild - B/op",
            "value": 928,
            "unit": "B/op",
            "extra": "758239 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchBOCInvalidBuild - allocs/op",
            "value": 9,
            "unit": "allocs/op",
            "extra": "758239 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCCDHeader",
            "value": 171.6,
            "unit": "ns/op\t     448 B/op\t       3 allocs/op",
            "extra": "7140642 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCCDHeader - ns/op",
            "value": 171.6,
            "unit": "ns/op",
            "extra": "7140642 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCCDHeader - B/op",
            "value": 448,
            "unit": "B/op",
            "extra": "7140642 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCCDHeader - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "7140642 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCCDAddendumCount",
            "value": 2615,
            "unit": "ns/op\t    1312 B/op\t      17 allocs/op",
            "extra": "407760 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCCDAddendumCount - ns/op",
            "value": 2615,
            "unit": "ns/op",
            "extra": "407760 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCCDAddendumCount - B/op",
            "value": 1312,
            "unit": "B/op",
            "extra": "407760 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCCDAddendumCount - allocs/op",
            "value": 17,
            "unit": "allocs/op",
            "extra": "407760 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCCDReceivingCompanyName",
            "value": 1867,
            "unit": "ns/op\t    1104 B/op\t      13 allocs/op",
            "extra": "616869 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCCDReceivingCompanyName - ns/op",
            "value": 1867,
            "unit": "ns/op",
            "extra": "616869 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCCDReceivingCompanyName - B/op",
            "value": 1104,
            "unit": "B/op",
            "extra": "616869 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCCDReceivingCompanyName - allocs/op",
            "value": 13,
            "unit": "allocs/op",
            "extra": "616869 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCCDAddendaTypeCode",
            "value": 2028,
            "unit": "ns/op\t    1120 B/op\t      14 allocs/op",
            "extra": "556620 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCCDAddendaTypeCode - ns/op",
            "value": 2028,
            "unit": "ns/op",
            "extra": "556620 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCCDAddendaTypeCode - B/op",
            "value": 1120,
            "unit": "B/op",
            "extra": "556620 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCCDAddendaTypeCode - allocs/op",
            "value": 14,
            "unit": "allocs/op",
            "extra": "556620 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCCDSEC",
            "value": 2205,
            "unit": "ns/op\t    1048 B/op\t      13 allocs/op",
            "extra": "516302 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCCDSEC - ns/op",
            "value": 2205,
            "unit": "ns/op",
            "extra": "516302 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCCDSEC - B/op",
            "value": 1048,
            "unit": "B/op",
            "extra": "516302 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCCDSEC - allocs/op",
            "value": 13,
            "unit": "allocs/op",
            "extra": "516302 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCCDAddendaCount",
            "value": 3223,
            "unit": "ns/op\t    1440 B/op\t      20 allocs/op",
            "extra": "354150 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCCDAddendaCount - ns/op",
            "value": 3223,
            "unit": "ns/op",
            "extra": "354150 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCCDAddendaCount - B/op",
            "value": 1440,
            "unit": "B/op",
            "extra": "354150 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCCDAddendaCount - allocs/op",
            "value": 20,
            "unit": "allocs/op",
            "extra": "354150 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCCDCreate",
            "value": 1817,
            "unit": "ns/op\t    1040 B/op\t      13 allocs/op",
            "extra": "632589 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCCDCreate - ns/op",
            "value": 1817,
            "unit": "ns/op",
            "extra": "632589 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCCDCreate - B/op",
            "value": 1040,
            "unit": "B/op",
            "extra": "632589 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCCDCreate - allocs/op",
            "value": 13,
            "unit": "allocs/op",
            "extra": "632589 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCCDReceivingCompanyField",
            "value": 1704,
            "unit": "ns/op\t     984 B/op\t      12 allocs/op",
            "extra": "650484 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCCDReceivingCompanyField - ns/op",
            "value": 1704,
            "unit": "ns/op",
            "extra": "650484 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCCDReceivingCompanyField - B/op",
            "value": 984,
            "unit": "B/op",
            "extra": "650484 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCCDReceivingCompanyField - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "650484 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCIEHeader",
            "value": 171,
            "unit": "ns/op\t     448 B/op\t       3 allocs/op",
            "extra": "6935592 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCIEHeader - ns/op",
            "value": 171,
            "unit": "ns/op",
            "extra": "6935592 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCIEHeader - B/op",
            "value": 448,
            "unit": "B/op",
            "extra": "6935592 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCIEHeader - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "6935592 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCIECreate",
            "value": 2484,
            "unit": "ns/op\t    1096 B/op\t      13 allocs/op",
            "extra": "471834 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCIECreate - ns/op",
            "value": 2484,
            "unit": "ns/op",
            "extra": "471834 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCIECreate - B/op",
            "value": 1096,
            "unit": "B/op",
            "extra": "471834 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCIECreate - allocs/op",
            "value": 13,
            "unit": "allocs/op",
            "extra": "471834 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCIEStandardEntryClassCode",
            "value": 2654,
            "unit": "ns/op\t    1176 B/op\t      14 allocs/op",
            "extra": "390064 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCIEStandardEntryClassCode - ns/op",
            "value": 2654,
            "unit": "ns/op",
            "extra": "390064 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCIEStandardEntryClassCode - B/op",
            "value": 1176,
            "unit": "B/op",
            "extra": "390064 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCIEStandardEntryClassCode - allocs/op",
            "value": 14,
            "unit": "allocs/op",
            "extra": "390064 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCIEServiceClassCodeEquality",
            "value": 2503,
            "unit": "ns/op\t    1184 B/op\t      15 allocs/op",
            "extra": "452047 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCIEServiceClassCodeEquality - ns/op",
            "value": 2503,
            "unit": "ns/op",
            "extra": "452047 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCIEServiceClassCodeEquality - B/op",
            "value": 1184,
            "unit": "B/op",
            "extra": "452047 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCIEServiceClassCodeEquality - allocs/op",
            "value": 15,
            "unit": "allocs/op",
            "extra": "452047 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCIEMixedCreditsAndDebits",
            "value": 2510,
            "unit": "ns/op\t    1184 B/op\t      15 allocs/op",
            "extra": "465026 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCIEMixedCreditsAndDebits - ns/op",
            "value": 2510,
            "unit": "ns/op",
            "extra": "465026 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCIEMixedCreditsAndDebits - B/op",
            "value": 1184,
            "unit": "B/op",
            "extra": "465026 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCIEMixedCreditsAndDebits - allocs/op",
            "value": 15,
            "unit": "allocs/op",
            "extra": "465026 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCIEDebitsOnly",
            "value": 2502,
            "unit": "ns/op\t    1184 B/op\t      15 allocs/op",
            "extra": "455288 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCIEDebitsOnly - ns/op",
            "value": 2502,
            "unit": "ns/op",
            "extra": "455288 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCIEDebitsOnly - B/op",
            "value": 1184,
            "unit": "B/op",
            "extra": "455288 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCIEDebitsOnly - allocs/op",
            "value": 15,
            "unit": "allocs/op",
            "extra": "455288 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCIEAutomatedAccountingAdvices",
            "value": 2538,
            "unit": "ns/op\t    1192 B/op\t      16 allocs/op",
            "extra": "446713 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCIEAutomatedAccountingAdvices - ns/op",
            "value": 2538,
            "unit": "ns/op",
            "extra": "446713 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCIEAutomatedAccountingAdvices - B/op",
            "value": 1192,
            "unit": "B/op",
            "extra": "446713 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCIEAutomatedAccountingAdvices - allocs/op",
            "value": 16,
            "unit": "allocs/op",
            "extra": "446713 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCIETransactionCode",
            "value": 2574,
            "unit": "ns/op\t    1176 B/op\t      14 allocs/op",
            "extra": "447025 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCIETransactionCode - ns/op",
            "value": 2574,
            "unit": "ns/op",
            "extra": "447025 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCIETransactionCode - B/op",
            "value": 1176,
            "unit": "B/op",
            "extra": "447025 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCIETransactionCode - allocs/op",
            "value": 14,
            "unit": "allocs/op",
            "extra": "447025 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCIEAddendaCount",
            "value": 3273,
            "unit": "ns/op\t    1440 B/op\t      20 allocs/op",
            "extra": "348741 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCIEAddendaCount - ns/op",
            "value": 3273,
            "unit": "ns/op",
            "extra": "348741 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCIEAddendaCount - B/op",
            "value": 1440,
            "unit": "B/op",
            "extra": "348741 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCIEAddendaCount - allocs/op",
            "value": 20,
            "unit": "allocs/op",
            "extra": "348741 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCIEAddendaCountZero",
            "value": 1041,
            "unit": "ns/op\t     736 B/op\t       7 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCIEAddendaCountZero - ns/op",
            "value": 1041,
            "unit": "ns/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCIEAddendaCountZero - B/op",
            "value": 736,
            "unit": "B/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCIEAddendaCountZero - allocs/op",
            "value": 7,
            "unit": "allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCIEInvalidAddendum",
            "value": 1447,
            "unit": "ns/op\t    1040 B/op\t      10 allocs/op",
            "extra": "776931 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCIEInvalidAddendum - ns/op",
            "value": 1447,
            "unit": "ns/op",
            "extra": "776931 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCIEInvalidAddendum - B/op",
            "value": 1040,
            "unit": "B/op",
            "extra": "776931 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCIEInvalidAddendum - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "776931 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCIEInvalidAddenda",
            "value": 1119,
            "unit": "ns/op\t     984 B/op\t      12 allocs/op",
            "extra": "952833 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCIEInvalidAddenda - ns/op",
            "value": 1119,
            "unit": "ns/op",
            "extra": "952833 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCIEInvalidAddenda - B/op",
            "value": 984,
            "unit": "B/op",
            "extra": "952833 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCIEInvalidAddenda - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "952833 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCIEInvalidBuild",
            "value": 1774,
            "unit": "ns/op\t    1024 B/op\t      12 allocs/op",
            "extra": "620923 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCIEInvalidBuild - ns/op",
            "value": 1774,
            "unit": "ns/op",
            "extra": "620923 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCIEInvalidBuild - B/op",
            "value": 1024,
            "unit": "B/op",
            "extra": "620923 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCIEInvalidBuild - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "620923 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCIECardTransactionType",
            "value": 2203,
            "unit": "ns/op\t     968 B/op\t      12 allocs/op",
            "extra": "509856 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCIECardTransactionType - ns/op",
            "value": 2203,
            "unit": "ns/op",
            "extra": "509856 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCIECardTransactionType - B/op",
            "value": 968,
            "unit": "B/op",
            "extra": "509856 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCIECardTransactionType - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "509856 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCORHeader",
            "value": 171.1,
            "unit": "ns/op\t     448 B/op\t       3 allocs/op",
            "extra": "7153450 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCORHeader - ns/op",
            "value": 171.1,
            "unit": "ns/op",
            "extra": "7153450 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCORHeader - B/op",
            "value": 448,
            "unit": "B/op",
            "extra": "7153450 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCORHeader - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "7153450 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCORSEC",
            "value": 1894,
            "unit": "ns/op\t    1088 B/op\t      10 allocs/op",
            "extra": "593324 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCORSEC - ns/op",
            "value": 1894,
            "unit": "ns/op",
            "extra": "593324 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCORSEC - B/op",
            "value": 1088,
            "unit": "B/op",
            "extra": "593324 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCORSEC - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "593324 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCORAddendumCountTwo",
            "value": 2146,
            "unit": "ns/op\t    1280 B/op\t      11 allocs/op",
            "extra": "531573 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCORAddendumCountTwo - ns/op",
            "value": 2146,
            "unit": "ns/op",
            "extra": "531573 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCORAddendumCountTwo - B/op",
            "value": 1280,
            "unit": "B/op",
            "extra": "531573 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCORAddendumCountTwo - allocs/op",
            "value": 11,
            "unit": "allocs/op",
            "extra": "531573 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCORAddendaCountZero",
            "value": 1051,
            "unit": "ns/op\t     816 B/op\t       8 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCORAddendaCountZero - ns/op",
            "value": 1051,
            "unit": "ns/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCORAddendaCountZero - B/op",
            "value": 816,
            "unit": "B/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCORAddendaCountZero - allocs/op",
            "value": 8,
            "unit": "allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCORAddendaType",
            "value": 1301,
            "unit": "ns/op\t     912 B/op\t      11 allocs/op",
            "extra": "864688 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCORAddendaType - ns/op",
            "value": 1301,
            "unit": "ns/op",
            "extra": "864688 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCORAddendaType - B/op",
            "value": 912,
            "unit": "B/op",
            "extra": "864688 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCORAddendaType - allocs/op",
            "value": 11,
            "unit": "allocs/op",
            "extra": "864688 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCORAddendaTypeCode",
            "value": 1840,
            "unit": "ns/op\t    1168 B/op\t      12 allocs/op",
            "extra": "613628 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCORAddendaTypeCode - ns/op",
            "value": 1840,
            "unit": "ns/op",
            "extra": "613628 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCORAddendaTypeCode - B/op",
            "value": 1168,
            "unit": "B/op",
            "extra": "613628 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCORAddendaTypeCode - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "613628 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCORAmount",
            "value": 2870,
            "unit": "ns/op\t    1440 B/op\t      15 allocs/op",
            "extra": "402280 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCORAmount - ns/op",
            "value": 2870,
            "unit": "ns/op",
            "extra": "402280 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCORAmount - B/op",
            "value": 1440,
            "unit": "B/op",
            "extra": "402280 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCORAmount - allocs/op",
            "value": 15,
            "unit": "allocs/op",
            "extra": "402280 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCORTransactionCode27",
            "value": 2153,
            "unit": "ns/op\t    1216 B/op\t      11 allocs/op",
            "extra": "524101 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCORTransactionCode27 - ns/op",
            "value": 2153,
            "unit": "ns/op",
            "extra": "524101 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCORTransactionCode27 - B/op",
            "value": 1216,
            "unit": "B/op",
            "extra": "524101 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCORTransactionCode27 - allocs/op",
            "value": 11,
            "unit": "allocs/op",
            "extra": "524101 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCORTransactionCode21",
            "value": 2063,
            "unit": "ns/op\t    1136 B/op\t      10 allocs/op",
            "extra": "543004 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCORTransactionCode21 - ns/op",
            "value": 2063,
            "unit": "ns/op",
            "extra": "543004 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCORTransactionCode21 - B/op",
            "value": 1136,
            "unit": "B/op",
            "extra": "543004 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCORTransactionCode21 - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "543004 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCORCreate",
            "value": 1550,
            "unit": "ns/op\t    1072 B/op\t      10 allocs/op",
            "extra": "713216 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCORCreate - ns/op",
            "value": 1550,
            "unit": "ns/op",
            "extra": "713216 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCORCreate - B/op",
            "value": 1072,
            "unit": "B/op",
            "extra": "713216 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCORCreate - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "713216 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCORServiceClassCodeEquality",
            "value": 2252,
            "unit": "ns/op\t    1232 B/op\t      13 allocs/op",
            "extra": "515366 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCORServiceClassCodeEquality - ns/op",
            "value": 2252,
            "unit": "ns/op",
            "extra": "515366 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCORServiceClassCodeEquality - B/op",
            "value": 1232,
            "unit": "B/op",
            "extra": "515366 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCORServiceClassCodeEquality - allocs/op",
            "value": 13,
            "unit": "allocs/op",
            "extra": "515366 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCTXHeader",
            "value": 172.7,
            "unit": "ns/op\t     448 B/op\t       3 allocs/op",
            "extra": "6947420 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCTXHeader - ns/op",
            "value": 172.7,
            "unit": "ns/op",
            "extra": "6947420 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCTXHeader - B/op",
            "value": 448,
            "unit": "B/op",
            "extra": "6947420 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCTXHeader - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "6947420 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCTXCreate",
            "value": 2483,
            "unit": "ns/op\t    1040 B/op\t      16 allocs/op",
            "extra": "460885 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCTXCreate - ns/op",
            "value": 2483,
            "unit": "ns/op",
            "extra": "460885 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCTXCreate - B/op",
            "value": 1040,
            "unit": "B/op",
            "extra": "460885 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCTXCreate - allocs/op",
            "value": 16,
            "unit": "allocs/op",
            "extra": "460885 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCTXStandardEntryClassCode",
            "value": 2789,
            "unit": "ns/op\t    1248 B/op\t      18 allocs/op",
            "extra": "410331 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCTXStandardEntryClassCode - ns/op",
            "value": 2789,
            "unit": "ns/op",
            "extra": "410331 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCTXStandardEntryClassCode - B/op",
            "value": 1248,
            "unit": "B/op",
            "extra": "410331 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCTXStandardEntryClassCode - allocs/op",
            "value": 18,
            "unit": "allocs/op",
            "extra": "410331 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCTXServiceClassCodeEquality",
            "value": 2767,
            "unit": "ns/op\t    1256 B/op\t      19 allocs/op",
            "extra": "422010 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCTXServiceClassCodeEquality - ns/op",
            "value": 2767,
            "unit": "ns/op",
            "extra": "422010 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCTXServiceClassCodeEquality - B/op",
            "value": 1256,
            "unit": "B/op",
            "extra": "422010 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCTXServiceClassCodeEquality - allocs/op",
            "value": 19,
            "unit": "allocs/op",
            "extra": "422010 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCTXAddendaCount",
            "value": 4923,
            "unit": "ns/op\t    1688 B/op\t      28 allocs/op",
            "extra": "237256 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCTXAddendaCount - ns/op",
            "value": 4923,
            "unit": "ns/op",
            "extra": "237256 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCTXAddendaCount - B/op",
            "value": 1688,
            "unit": "B/op",
            "extra": "237256 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCTXAddendaCount - allocs/op",
            "value": 28,
            "unit": "allocs/op",
            "extra": "237256 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCTXAddendaCountZero",
            "value": 1802,
            "unit": "ns/op\t    1048 B/op\t      15 allocs/op",
            "extra": "636421 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCTXAddendaCountZero - ns/op",
            "value": 1802,
            "unit": "ns/op",
            "extra": "636421 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCTXAddendaCountZero - B/op",
            "value": 1048,
            "unit": "B/op",
            "extra": "636421 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCTXAddendaCountZero - allocs/op",
            "value": 15,
            "unit": "allocs/op",
            "extra": "636421 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCTXInvalidAddenda",
            "value": 1339,
            "unit": "ns/op\t    1056 B/op\t      16 allocs/op",
            "extra": "796377 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCTXInvalidAddenda - ns/op",
            "value": 1339,
            "unit": "ns/op",
            "extra": "796377 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCTXInvalidAddenda - B/op",
            "value": 1056,
            "unit": "B/op",
            "extra": "796377 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCTXInvalidAddenda - allocs/op",
            "value": 16,
            "unit": "allocs/op",
            "extra": "796377 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCTXInvalidBuild",
            "value": 2083,
            "unit": "ns/op\t    1096 B/op\t      16 allocs/op",
            "extra": "561676 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCTXInvalidBuild - ns/op",
            "value": 2083,
            "unit": "ns/op",
            "extra": "561676 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCTXInvalidBuild - B/op",
            "value": 1096,
            "unit": "B/op",
            "extra": "561676 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCTXInvalidBuild - allocs/op",
            "value": 16,
            "unit": "allocs/op",
            "extra": "561676 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCTXAddenda10000",
            "value": 2017768,
            "unit": "ns/op\t 1191866 B/op\t   20039 allocs/op",
            "extra": "597 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCTXAddenda10000 - ns/op",
            "value": 2017768,
            "unit": "ns/op",
            "extra": "597 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCTXAddenda10000 - B/op",
            "value": 1191866,
            "unit": "B/op",
            "extra": "597 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCTXAddenda10000 - allocs/op",
            "value": 20039,
            "unit": "allocs/op",
            "extra": "597 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCTXAddendaRecords",
            "value": 107971,
            "unit": "ns/op\t   60138 B/op\t    1158 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCTXAddendaRecords - ns/op",
            "value": 107971,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCTXAddendaRecords - B/op",
            "value": 60138,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCTXAddendaRecords - allocs/op",
            "value": 1158,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCTXReceivingCompany",
            "value": 2922,
            "unit": "ns/op\t    1072 B/op\t      17 allocs/op",
            "extra": "396492 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCTXReceivingCompany - ns/op",
            "value": 2922,
            "unit": "ns/op",
            "extra": "396492 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCTXReceivingCompany - B/op",
            "value": 1072,
            "unit": "B/op",
            "extra": "396492 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCTXReceivingCompany - allocs/op",
            "value": 17,
            "unit": "allocs/op",
            "extra": "396492 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCTXReserved",
            "value": 1970,
            "unit": "ns/op\t    1032 B/op\t      15 allocs/op",
            "extra": "557625 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCTXReserved - ns/op",
            "value": 1970,
            "unit": "ns/op",
            "extra": "557625 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCTXReserved - B/op",
            "value": 1032,
            "unit": "B/op",
            "extra": "557625 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCTXReserved - allocs/op",
            "value": 15,
            "unit": "allocs/op",
            "extra": "557625 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCTXZeroAddendaRecords",
            "value": 1792,
            "unit": "ns/op\t    1048 B/op\t      15 allocs/op",
            "extra": "644899 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCTXZeroAddendaRecords - ns/op",
            "value": 1792,
            "unit": "ns/op",
            "extra": "644899 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCTXZeroAddendaRecords - B/op",
            "value": 1048,
            "unit": "B/op",
            "extra": "644899 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCTXZeroAddendaRecords - allocs/op",
            "value": 15,
            "unit": "allocs/op",
            "extra": "644899 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCTXPrenoteAddendaRecords",
            "value": 1561,
            "unit": "ns/op\t     904 B/op\t      14 allocs/op",
            "extra": "716827 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCTXPrenoteAddendaRecords - ns/op",
            "value": 1561,
            "unit": "ns/op",
            "extra": "716827 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCTXPrenoteAddendaRecords - B/op",
            "value": 904,
            "unit": "B/op",
            "extra": "716827 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCTXPrenoteAddendaRecords - allocs/op",
            "value": 14,
            "unit": "allocs/op",
            "extra": "716827 times\n4 procs"
          },
          {
            "name": "BenchmarkMockBatchControl",
            "value": 22.73,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "52501510 times\n4 procs"
          },
          {
            "name": "BenchmarkMockBatchControl - ns/op",
            "value": 22.73,
            "unit": "ns/op",
            "extra": "52501510 times\n4 procs"
          },
          {
            "name": "BenchmarkMockBatchControl - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "52501510 times\n4 procs"
          },
          {
            "name": "BenchmarkMockBatchControl - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "52501510 times\n4 procs"
          },
          {
            "name": "BenchmarkParseBatchControl",
            "value": 5363,
            "unit": "ns/op\t   15512 B/op\t      27 allocs/op",
            "extra": "218395 times\n4 procs"
          },
          {
            "name": "BenchmarkParseBatchControl - ns/op",
            "value": 5363,
            "unit": "ns/op",
            "extra": "218395 times\n4 procs"
          },
          {
            "name": "BenchmarkParseBatchControl - B/op",
            "value": 15512,
            "unit": "B/op",
            "extra": "218395 times\n4 procs"
          },
          {
            "name": "BenchmarkParseBatchControl - allocs/op",
            "value": 27,
            "unit": "allocs/op",
            "extra": "218395 times\n4 procs"
          },
          {
            "name": "BenchmarkBCString",
            "value": 5678,
            "unit": "ns/op\t   15617 B/op\t      29 allocs/op",
            "extra": "212926 times\n4 procs"
          },
          {
            "name": "BenchmarkBCString - ns/op",
            "value": 5678,
            "unit": "ns/op",
            "extra": "212926 times\n4 procs"
          },
          {
            "name": "BenchmarkBCString - B/op",
            "value": 15617,
            "unit": "B/op",
            "extra": "212926 times\n4 procs"
          },
          {
            "name": "BenchmarkBCString - allocs/op",
            "value": 29,
            "unit": "allocs/op",
            "extra": "212926 times\n4 procs"
          },
          {
            "name": "BenchmarkBCisServiceClassErr",
            "value": 114.8,
            "unit": "ns/op\t      83 B/op\t       3 allocs/op",
            "extra": "10363183 times\n4 procs"
          },
          {
            "name": "BenchmarkBCisServiceClassErr - ns/op",
            "value": 114.8,
            "unit": "ns/op",
            "extra": "10363183 times\n4 procs"
          },
          {
            "name": "BenchmarkBCisServiceClassErr - B/op",
            "value": 83,
            "unit": "B/op",
            "extra": "10363183 times\n4 procs"
          },
          {
            "name": "BenchmarkBCisServiceClassErr - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "10363183 times\n4 procs"
          },
          {
            "name": "BenchmarkBCBatchNumber",
            "value": 24.36,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "50023254 times\n4 procs"
          },
          {
            "name": "BenchmarkBCBatchNumber - ns/op",
            "value": 24.36,
            "unit": "ns/op",
            "extra": "50023254 times\n4 procs"
          },
          {
            "name": "BenchmarkBCBatchNumber - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "50023254 times\n4 procs"
          },
          {
            "name": "BenchmarkBCBatchNumber - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "50023254 times\n4 procs"
          },
          {
            "name": "BenchmarkBCCompanyIdentificationAlphaNumeric",
            "value": 311.2,
            "unit": "ns/op\t     160 B/op\t       4 allocs/op",
            "extra": "3844401 times\n4 procs"
          },
          {
            "name": "BenchmarkBCCompanyIdentificationAlphaNumeric - ns/op",
            "value": 311.2,
            "unit": "ns/op",
            "extra": "3844401 times\n4 procs"
          },
          {
            "name": "BenchmarkBCCompanyIdentificationAlphaNumeric - B/op",
            "value": 160,
            "unit": "B/op",
            "extra": "3844401 times\n4 procs"
          },
          {
            "name": "BenchmarkBCCompanyIdentificationAlphaNumeric - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "3844401 times\n4 procs"
          },
          {
            "name": "BenchmarkBCMessageAuthenticationCodeAlphaNumeric",
            "value": 323.2,
            "unit": "ns/op\t     160 B/op\t       4 allocs/op",
            "extra": "3714115 times\n4 procs"
          },
          {
            "name": "BenchmarkBCMessageAuthenticationCodeAlphaNumeric - ns/op",
            "value": 323.2,
            "unit": "ns/op",
            "extra": "3714115 times\n4 procs"
          },
          {
            "name": "BenchmarkBCMessageAuthenticationCodeAlphaNumeric - B/op",
            "value": 160,
            "unit": "B/op",
            "extra": "3714115 times\n4 procs"
          },
          {
            "name": "BenchmarkBCMessageAuthenticationCodeAlphaNumeric - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "3714115 times\n4 procs"
          },
          {
            "name": "BenchmarkBCFieldInclusionServiceClassCode",
            "value": 92.86,
            "unit": "ns/op\t      80 B/op\t       2 allocs/op",
            "extra": "12817113 times\n4 procs"
          },
          {
            "name": "BenchmarkBCFieldInclusionServiceClassCode - ns/op",
            "value": 92.86,
            "unit": "ns/op",
            "extra": "12817113 times\n4 procs"
          },
          {
            "name": "BenchmarkBCFieldInclusionServiceClassCode - B/op",
            "value": 80,
            "unit": "B/op",
            "extra": "12817113 times\n4 procs"
          },
          {
            "name": "BenchmarkBCFieldInclusionServiceClassCode - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "12817113 times\n4 procs"
          },
          {
            "name": "BenchmarkBCFieldInclusionODFIIdentification",
            "value": 102.2,
            "unit": "ns/op\t      80 B/op\t       2 allocs/op",
            "extra": "11680209 times\n4 procs"
          },
          {
            "name": "BenchmarkBCFieldInclusionODFIIdentification - ns/op",
            "value": 102.2,
            "unit": "ns/op",
            "extra": "11680209 times\n4 procs"
          },
          {
            "name": "BenchmarkBCFieldInclusionODFIIdentification - B/op",
            "value": 80,
            "unit": "B/op",
            "extra": "11680209 times\n4 procs"
          },
          {
            "name": "BenchmarkBCFieldInclusionODFIIdentification - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "11680209 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchControlLength",
            "value": 475.2,
            "unit": "ns/op\t     160 B/op\t       7 allocs/op",
            "extra": "2521429 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchControlLength - ns/op",
            "value": 475.2,
            "unit": "ns/op",
            "extra": "2521429 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchControlLength - B/op",
            "value": 160,
            "unit": "B/op",
            "extra": "2521429 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchControlLength - allocs/op",
            "value": 7,
            "unit": "allocs/op",
            "extra": "2521429 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchDNEHeader",
            "value": 352,
            "unit": "ns/op\t     456 B/op\t       4 allocs/op",
            "extra": "3413731 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchDNEHeader - ns/op",
            "value": 352,
            "unit": "ns/op",
            "extra": "3413731 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchDNEHeader - B/op",
            "value": 456,
            "unit": "B/op",
            "extra": "3413731 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchDNEHeader - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "3413731 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchDNEAddendumCount",
            "value": 2993,
            "unit": "ns/op\t    1320 B/op\t      18 allocs/op",
            "extra": "387812 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchDNEAddendumCount - ns/op",
            "value": 2993,
            "unit": "ns/op",
            "extra": "387812 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchDNEAddendumCount - B/op",
            "value": 1320,
            "unit": "B/op",
            "extra": "387812 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchDNEAddendumCount - allocs/op",
            "value": 18,
            "unit": "allocs/op",
            "extra": "387812 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchDNEReceivingCompanyName",
            "value": 2210,
            "unit": "ns/op\t    1112 B/op\t      14 allocs/op",
            "extra": "519055 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchDNEReceivingCompanyName - ns/op",
            "value": 2210,
            "unit": "ns/op",
            "extra": "519055 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchDNEReceivingCompanyName - B/op",
            "value": 1112,
            "unit": "B/op",
            "extra": "519055 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchDNEReceivingCompanyName - allocs/op",
            "value": 14,
            "unit": "allocs/op",
            "extra": "519055 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchDNEAddendaTypeCode",
            "value": 2524,
            "unit": "ns/op\t     976 B/op\t      13 allocs/op",
            "extra": "453930 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchDNEAddendaTypeCode - ns/op",
            "value": 2524,
            "unit": "ns/op",
            "extra": "453930 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchDNEAddendaTypeCode - B/op",
            "value": 976,
            "unit": "B/op",
            "extra": "453930 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchDNEAddendaTypeCode - allocs/op",
            "value": 13,
            "unit": "allocs/op",
            "extra": "453930 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchDNESEC",
            "value": 2583,
            "unit": "ns/op\t    1056 B/op\t      14 allocs/op",
            "extra": "442609 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchDNESEC - ns/op",
            "value": 2583,
            "unit": "ns/op",
            "extra": "442609 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchDNESEC - B/op",
            "value": 1056,
            "unit": "B/op",
            "extra": "442609 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchDNESEC - allocs/op",
            "value": 14,
            "unit": "allocs/op",
            "extra": "442609 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchDNEAddendaCount",
            "value": 3677,
            "unit": "ns/op\t    1448 B/op\t      21 allocs/op",
            "extra": "320336 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchDNEAddendaCount - ns/op",
            "value": 3677,
            "unit": "ns/op",
            "extra": "320336 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchDNEAddendaCount - B/op",
            "value": 1448,
            "unit": "B/op",
            "extra": "320336 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchDNEAddendaCount - allocs/op",
            "value": 21,
            "unit": "allocs/op",
            "extra": "320336 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchDNEServiceClassCode",
            "value": 2138,
            "unit": "ns/op\t    1048 B/op\t      14 allocs/op",
            "extra": "525448 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchDNEServiceClassCode - ns/op",
            "value": 2138,
            "unit": "ns/op",
            "extra": "525448 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchDNEServiceClassCode - B/op",
            "value": 1048,
            "unit": "B/op",
            "extra": "525448 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchDNEServiceClassCode - allocs/op",
            "value": 14,
            "unit": "allocs/op",
            "extra": "525448 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchENRHeader",
            "value": 169.8,
            "unit": "ns/op\t     448 B/op\t       3 allocs/op",
            "extra": "7200122 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchENRHeader - ns/op",
            "value": 169.8,
            "unit": "ns/op",
            "extra": "7200122 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchENRHeader - B/op",
            "value": 448,
            "unit": "B/op",
            "extra": "7200122 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchENRHeader - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "7200122 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchENRAddendumCount",
            "value": 3706,
            "unit": "ns/op\t    1216 B/op\t      18 allocs/op",
            "extra": "311008 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchENRAddendumCount - ns/op",
            "value": 3706,
            "unit": "ns/op",
            "extra": "311008 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchENRAddendumCount - B/op",
            "value": 1216,
            "unit": "B/op",
            "extra": "311008 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchENRAddendumCount - allocs/op",
            "value": 18,
            "unit": "allocs/op",
            "extra": "311008 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchENRCompanyEntryDescription",
            "value": 2711,
            "unit": "ns/op\t    1192 B/op\t      15 allocs/op",
            "extra": "429652 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchENRCompanyEntryDescription - ns/op",
            "value": 2711,
            "unit": "ns/op",
            "extra": "429652 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchENRCompanyEntryDescription - B/op",
            "value": 1192,
            "unit": "B/op",
            "extra": "429652 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchENRCompanyEntryDescription - allocs/op",
            "value": 15,
            "unit": "allocs/op",
            "extra": "429652 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchENRAddendaTypeCode",
            "value": 2234,
            "unit": "ns/op\t    1120 B/op\t      14 allocs/op",
            "extra": "517078 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchENRAddendaTypeCode - ns/op",
            "value": 2234,
            "unit": "ns/op",
            "extra": "517078 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchENRAddendaTypeCode - B/op",
            "value": 1120,
            "unit": "B/op",
            "extra": "517078 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchENRAddendaTypeCode - allocs/op",
            "value": 14,
            "unit": "allocs/op",
            "extra": "517078 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchENRSEC",
            "value": 2391,
            "unit": "ns/op\t    1048 B/op\t      13 allocs/op",
            "extra": "469490 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchENRSEC - ns/op",
            "value": 2391,
            "unit": "ns/op",
            "extra": "469490 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchENRSEC - B/op",
            "value": 1048,
            "unit": "B/op",
            "extra": "469490 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchENRSEC - allocs/op",
            "value": 13,
            "unit": "allocs/op",
            "extra": "469490 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchENRAddendaCount",
            "value": 2999,
            "unit": "ns/op\t    1200 B/op\t      16 allocs/op",
            "extra": "388152 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchENRAddendaCount - ns/op",
            "value": 2999,
            "unit": "ns/op",
            "extra": "388152 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchENRAddendaCount - B/op",
            "value": 1200,
            "unit": "B/op",
            "extra": "388152 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchENRAddendaCount - allocs/op",
            "value": 16,
            "unit": "allocs/op",
            "extra": "388152 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchENRServiceClassCode",
            "value": 1932,
            "unit": "ns/op\t    1040 B/op\t      13 allocs/op",
            "extra": "587575 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchENRServiceClassCode - ns/op",
            "value": 1932,
            "unit": "ns/op",
            "extra": "587575 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchENRServiceClassCode - B/op",
            "value": 1040,
            "unit": "B/op",
            "extra": "587575 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchENRServiceClassCode - allocs/op",
            "value": 13,
            "unit": "allocs/op",
            "extra": "587575 times\n4 procs"
          },
          {
            "name": "BenchmarkMockBatchHeader",
            "value": 56.54,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "21296158 times\n4 procs"
          },
          {
            "name": "BenchmarkMockBatchHeader - ns/op",
            "value": 56.54,
            "unit": "ns/op",
            "extra": "21296158 times\n4 procs"
          },
          {
            "name": "BenchmarkMockBatchHeader - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "21296158 times\n4 procs"
          },
          {
            "name": "BenchmarkMockBatchHeader - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "21296158 times\n4 procs"
          },
          {
            "name": "BenchmarkParseBatchHeader",
            "value": 6123,
            "unit": "ns/op\t   15292 B/op\t      29 allocs/op",
            "extra": "194025 times\n4 procs"
          },
          {
            "name": "BenchmarkParseBatchHeader - ns/op",
            "value": 6123,
            "unit": "ns/op",
            "extra": "194025 times\n4 procs"
          },
          {
            "name": "BenchmarkParseBatchHeader - B/op",
            "value": 15292,
            "unit": "B/op",
            "extra": "194025 times\n4 procs"
          },
          {
            "name": "BenchmarkParseBatchHeader - allocs/op",
            "value": 29,
            "unit": "allocs/op",
            "extra": "194025 times\n4 procs"
          },
          {
            "name": "BenchmarkBHString",
            "value": 6424,
            "unit": "ns/op\t   15408 B/op\t      31 allocs/op",
            "extra": "187689 times\n4 procs"
          },
          {
            "name": "BenchmarkBHString - ns/op",
            "value": 6424,
            "unit": "ns/op",
            "extra": "187689 times\n4 procs"
          },
          {
            "name": "BenchmarkBHString - B/op",
            "value": 15408,
            "unit": "B/op",
            "extra": "187689 times\n4 procs"
          },
          {
            "name": "BenchmarkBHString - allocs/op",
            "value": 31,
            "unit": "allocs/op",
            "extra": "187689 times\n4 procs"
          },
          {
            "name": "BenchmarkInvalidServiceCode",
            "value": 67.81,
            "unit": "ns/op\t      64 B/op\t       1 allocs/op",
            "extra": "17363430 times\n4 procs"
          },
          {
            "name": "BenchmarkInvalidServiceCode - ns/op",
            "value": 67.81,
            "unit": "ns/op",
            "extra": "17363430 times\n4 procs"
          },
          {
            "name": "BenchmarkInvalidServiceCode - B/op",
            "value": 64,
            "unit": "B/op",
            "extra": "17363430 times\n4 procs"
          },
          {
            "name": "BenchmarkInvalidServiceCode - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "17363430 times\n4 procs"
          },
          {
            "name": "BenchmarkInvalidSECCode",
            "value": 96.38,
            "unit": "ns/op\t      80 B/op\t       2 allocs/op",
            "extra": "12341701 times\n4 procs"
          },
          {
            "name": "BenchmarkInvalidSECCode - ns/op",
            "value": 96.38,
            "unit": "ns/op",
            "extra": "12341701 times\n4 procs"
          },
          {
            "name": "BenchmarkInvalidSECCode - B/op",
            "value": 80,
            "unit": "B/op",
            "extra": "12341701 times\n4 procs"
          },
          {
            "name": "BenchmarkInvalidSECCode - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "12341701 times\n4 procs"
          },
          {
            "name": "BenchmarkInvalidOrigStatusCode",
            "value": 71.84,
            "unit": "ns/op\t      64 B/op\t       1 allocs/op",
            "extra": "16524681 times\n4 procs"
          },
          {
            "name": "BenchmarkInvalidOrigStatusCode - ns/op",
            "value": 71.84,
            "unit": "ns/op",
            "extra": "16524681 times\n4 procs"
          },
          {
            "name": "BenchmarkInvalidOrigStatusCode - B/op",
            "value": 64,
            "unit": "B/op",
            "extra": "16524681 times\n4 procs"
          },
          {
            "name": "BenchmarkInvalidOrigStatusCode - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "16524681 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchHeaderFieldInclusion",
            "value": 56.64,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "21396159 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchHeaderFieldInclusion - ns/op",
            "value": 56.64,
            "unit": "ns/op",
            "extra": "21396159 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchHeaderFieldInclusion - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "21396159 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchHeaderFieldInclusion - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "21396159 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchHeaderCompanyNameAlphaNumeric",
            "value": 329.1,
            "unit": "ns/op\t     160 B/op\t       4 allocs/op",
            "extra": "3643471 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchHeaderCompanyNameAlphaNumeric - ns/op",
            "value": 329.1,
            "unit": "ns/op",
            "extra": "3643471 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchHeaderCompanyNameAlphaNumeric - B/op",
            "value": 160,
            "unit": "B/op",
            "extra": "3643471 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchHeaderCompanyNameAlphaNumeric - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "3643471 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCompanyDiscretionaryDataAlphaNumeric",
            "value": 342.5,
            "unit": "ns/op\t     160 B/op\t       4 allocs/op",
            "extra": "3479984 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCompanyDiscretionaryDataAlphaNumeric - ns/op",
            "value": 342.5,
            "unit": "ns/op",
            "extra": "3479984 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCompanyDiscretionaryDataAlphaNumeric - B/op",
            "value": 160,
            "unit": "B/op",
            "extra": "3479984 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCompanyDiscretionaryDataAlphaNumeric - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "3479984 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCompanyIdentificationAlphaNumeric",
            "value": 674.1,
            "unit": "ns/op\t     160 B/op\t       4 allocs/op",
            "extra": "1786003 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCompanyIdentificationAlphaNumeric - ns/op",
            "value": 674.1,
            "unit": "ns/op",
            "extra": "1786003 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCompanyIdentificationAlphaNumeric - B/op",
            "value": 160,
            "unit": "B/op",
            "extra": "1786003 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCompanyIdentificationAlphaNumeric - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "1786003 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCompanyEntryDescriptionAlphaNumeric",
            "value": 707.9,
            "unit": "ns/op\t     160 B/op\t       4 allocs/op",
            "extra": "1697269 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCompanyEntryDescriptionAlphaNumeric - ns/op",
            "value": 707.9,
            "unit": "ns/op",
            "extra": "1697269 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCompanyEntryDescriptionAlphaNumeric - B/op",
            "value": 160,
            "unit": "B/op",
            "extra": "1697269 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCompanyEntryDescriptionAlphaNumeric - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "1697269 times\n4 procs"
          },
          {
            "name": "BenchmarkBHFieldInclusionCompanyName",
            "value": 68.8,
            "unit": "ns/op\t      64 B/op\t       1 allocs/op",
            "extra": "17183823 times\n4 procs"
          },
          {
            "name": "BenchmarkBHFieldInclusionCompanyName - ns/op",
            "value": 68.8,
            "unit": "ns/op",
            "extra": "17183823 times\n4 procs"
          },
          {
            "name": "BenchmarkBHFieldInclusionCompanyName - B/op",
            "value": 64,
            "unit": "B/op",
            "extra": "17183823 times\n4 procs"
          },
          {
            "name": "BenchmarkBHFieldInclusionCompanyName - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "17183823 times\n4 procs"
          },
          {
            "name": "BenchmarkBHFieldInclusionCompanyIdentification",
            "value": 67.69,
            "unit": "ns/op\t      64 B/op\t       1 allocs/op",
            "extra": "16975027 times\n4 procs"
          },
          {
            "name": "BenchmarkBHFieldInclusionCompanyIdentification - ns/op",
            "value": 67.69,
            "unit": "ns/op",
            "extra": "16975027 times\n4 procs"
          },
          {
            "name": "BenchmarkBHFieldInclusionCompanyIdentification - B/op",
            "value": 64,
            "unit": "B/op",
            "extra": "16975027 times\n4 procs"
          },
          {
            "name": "BenchmarkBHFieldInclusionCompanyIdentification - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "16975027 times\n4 procs"
          },
          {
            "name": "BenchmarkBHFieldInclusionStandardEntryClassCode",
            "value": 70.48,
            "unit": "ns/op\t      64 B/op\t       1 allocs/op",
            "extra": "16964412 times\n4 procs"
          },
          {
            "name": "BenchmarkBHFieldInclusionStandardEntryClassCode - ns/op",
            "value": 70.48,
            "unit": "ns/op",
            "extra": "16964412 times\n4 procs"
          },
          {
            "name": "BenchmarkBHFieldInclusionStandardEntryClassCode - B/op",
            "value": 64,
            "unit": "B/op",
            "extra": "16964412 times\n4 procs"
          },
          {
            "name": "BenchmarkBHFieldInclusionStandardEntryClassCode - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "16964412 times\n4 procs"
          },
          {
            "name": "BenchmarkBHFieldInclusionCompanyEntryDescription",
            "value": 65.65,
            "unit": "ns/op\t      64 B/op\t       1 allocs/op",
            "extra": "17968701 times\n4 procs"
          },
          {
            "name": "BenchmarkBHFieldInclusionCompanyEntryDescription - ns/op",
            "value": 65.65,
            "unit": "ns/op",
            "extra": "17968701 times\n4 procs"
          },
          {
            "name": "BenchmarkBHFieldInclusionCompanyEntryDescription - B/op",
            "value": 64,
            "unit": "B/op",
            "extra": "17968701 times\n4 procs"
          },
          {
            "name": "BenchmarkBHFieldInclusionCompanyEntryDescription - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "17968701 times\n4 procs"
          },
          {
            "name": "BenchmarkBHFieldInclusionOriginatorStatusCode",
            "value": 72.14,
            "unit": "ns/op\t      64 B/op\t       1 allocs/op",
            "extra": "16001385 times\n4 procs"
          },
          {
            "name": "BenchmarkBHFieldInclusionOriginatorStatusCode - ns/op",
            "value": 72.14,
            "unit": "ns/op",
            "extra": "16001385 times\n4 procs"
          },
          {
            "name": "BenchmarkBHFieldInclusionOriginatorStatusCode - B/op",
            "value": 64,
            "unit": "B/op",
            "extra": "16001385 times\n4 procs"
          },
          {
            "name": "BenchmarkBHFieldInclusionOriginatorStatusCode - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "16001385 times\n4 procs"
          },
          {
            "name": "BenchmarkBHFieldInclusionODFIIdentification",
            "value": 113,
            "unit": "ns/op\t      80 B/op\t       2 allocs/op",
            "extra": "10564023 times\n4 procs"
          },
          {
            "name": "BenchmarkBHFieldInclusionODFIIdentification - ns/op",
            "value": 113,
            "unit": "ns/op",
            "extra": "10564023 times\n4 procs"
          },
          {
            "name": "BenchmarkBHFieldInclusionODFIIdentification - B/op",
            "value": 80,
            "unit": "B/op",
            "extra": "10564023 times\n4 procs"
          },
          {
            "name": "BenchmarkBHFieldInclusionODFIIdentification - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "10564023 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchMTEHeader",
            "value": 349.1,
            "unit": "ns/op\t     456 B/op\t       4 allocs/op",
            "extra": "3457491 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchMTEHeader - ns/op",
            "value": 349.1,
            "unit": "ns/op",
            "extra": "3457491 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchMTEHeader - B/op",
            "value": 456,
            "unit": "B/op",
            "extra": "3457491 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchMTEHeader - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "3457491 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchMTEReceivingCompanyName",
            "value": 2240,
            "unit": "ns/op\t    1224 B/op\t      12 allocs/op",
            "extra": "511720 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchMTEReceivingCompanyName - ns/op",
            "value": 2240,
            "unit": "ns/op",
            "extra": "511720 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchMTEReceivingCompanyName - B/op",
            "value": 1224,
            "unit": "B/op",
            "extra": "511720 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchMTEReceivingCompanyName - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "511720 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchMTEAddendaTypeCode",
            "value": 2410,
            "unit": "ns/op\t    1240 B/op\t      13 allocs/op",
            "extra": "476786 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchMTEAddendaTypeCode - ns/op",
            "value": 2410,
            "unit": "ns/op",
            "extra": "476786 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchMTEAddendaTypeCode - B/op",
            "value": 1240,
            "unit": "B/op",
            "extra": "476786 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchMTEAddendaTypeCode - allocs/op",
            "value": 13,
            "unit": "allocs/op",
            "extra": "476786 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchMTESEC",
            "value": 2566,
            "unit": "ns/op\t    1160 B/op\t      11 allocs/op",
            "extra": "451789 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchMTESEC - ns/op",
            "value": 2566,
            "unit": "ns/op",
            "extra": "451789 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchMTESEC - B/op",
            "value": 1160,
            "unit": "B/op",
            "extra": "451789 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchMTESEC - allocs/op",
            "value": 11,
            "unit": "allocs/op",
            "extra": "451789 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchMTEServiceClassCode",
            "value": 2144,
            "unit": "ns/op\t    1160 B/op\t      12 allocs/op",
            "extra": "537081 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchMTEServiceClassCode - ns/op",
            "value": 2144,
            "unit": "ns/op",
            "extra": "537081 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchMTEServiceClassCode - B/op",
            "value": 1160,
            "unit": "B/op",
            "extra": "537081 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchMTEServiceClassCode - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "537081 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOPHeader",
            "value": 171.8,
            "unit": "ns/op\t     448 B/op\t       3 allocs/op",
            "extra": "7051086 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOPHeader - ns/op",
            "value": 171.8,
            "unit": "ns/op",
            "extra": "7051086 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOPHeader - B/op",
            "value": 448,
            "unit": "B/op",
            "extra": "7051086 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOPHeader - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "7051086 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOPCreate",
            "value": 2201,
            "unit": "ns/op\t    1024 B/op\t      11 allocs/op",
            "extra": "513400 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOPCreate - ns/op",
            "value": 2201,
            "unit": "ns/op",
            "extra": "513400 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOPCreate - B/op",
            "value": 1024,
            "unit": "B/op",
            "extra": "513400 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOPCreate - allocs/op",
            "value": 11,
            "unit": "allocs/op",
            "extra": "513400 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOPStandardEntryClassCode",
            "value": 2275,
            "unit": "ns/op\t    1104 B/op\t      12 allocs/op",
            "extra": "493735 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOPStandardEntryClassCode - ns/op",
            "value": 2275,
            "unit": "ns/op",
            "extra": "493735 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOPStandardEntryClassCode - B/op",
            "value": 1104,
            "unit": "B/op",
            "extra": "493735 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOPStandardEntryClassCode - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "493735 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOPServiceClassCodeEquality",
            "value": 2364,
            "unit": "ns/op\t    1120 B/op\t      14 allocs/op",
            "extra": "488364 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOPServiceClassCodeEquality - ns/op",
            "value": 2364,
            "unit": "ns/op",
            "extra": "488364 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOPServiceClassCodeEquality - B/op",
            "value": 1120,
            "unit": "B/op",
            "extra": "488364 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOPServiceClassCodeEquality - allocs/op",
            "value": 14,
            "unit": "allocs/op",
            "extra": "488364 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOPMixedCreditsAndDebits",
            "value": 2358,
            "unit": "ns/op\t    1120 B/op\t      14 allocs/op",
            "extra": "490704 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOPMixedCreditsAndDebits - ns/op",
            "value": 2358,
            "unit": "ns/op",
            "extra": "490704 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOPMixedCreditsAndDebits - B/op",
            "value": 1120,
            "unit": "B/op",
            "extra": "490704 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOPMixedCreditsAndDebits - allocs/op",
            "value": 14,
            "unit": "allocs/op",
            "extra": "490704 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOPCreditsOnly",
            "value": 2372,
            "unit": "ns/op\t    1120 B/op\t      14 allocs/op",
            "extra": "486439 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOPCreditsOnly - ns/op",
            "value": 2372,
            "unit": "ns/op",
            "extra": "486439 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOPCreditsOnly - B/op",
            "value": 1120,
            "unit": "B/op",
            "extra": "486439 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOPCreditsOnly - allocs/op",
            "value": 14,
            "unit": "allocs/op",
            "extra": "486439 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOPAutomatedAccountingAdvices",
            "value": 2358,
            "unit": "ns/op\t    1128 B/op\t      15 allocs/op",
            "extra": "490100 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOPAutomatedAccountingAdvices - ns/op",
            "value": 2358,
            "unit": "ns/op",
            "extra": "490100 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOPAutomatedAccountingAdvices - B/op",
            "value": 1128,
            "unit": "B/op",
            "extra": "490100 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOPAutomatedAccountingAdvices - allocs/op",
            "value": 15,
            "unit": "allocs/op",
            "extra": "490100 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOPAmount",
            "value": 2877,
            "unit": "ns/op\t    1328 B/op\t      19 allocs/op",
            "extra": "393792 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOPAmount - ns/op",
            "value": 2877,
            "unit": "ns/op",
            "extra": "393792 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOPAmount - B/op",
            "value": 1328,
            "unit": "B/op",
            "extra": "393792 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOPAmount - allocs/op",
            "value": 19,
            "unit": "allocs/op",
            "extra": "393792 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOPCheckSerialNumber",
            "value": 1997,
            "unit": "ns/op\t     976 B/op\t      11 allocs/op",
            "extra": "555498 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOPCheckSerialNumber - ns/op",
            "value": 1997,
            "unit": "ns/op",
            "extra": "555498 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOPCheckSerialNumber - B/op",
            "value": 976,
            "unit": "B/op",
            "extra": "555498 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOPCheckSerialNumber - allocs/op",
            "value": 11,
            "unit": "allocs/op",
            "extra": "555498 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOPCheckSerialNumberField",
            "value": 1583,
            "unit": "ns/op\t     896 B/op\t      10 allocs/op",
            "extra": "708646 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOPCheckSerialNumberField - ns/op",
            "value": 1583,
            "unit": "ns/op",
            "extra": "708646 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOPCheckSerialNumberField - B/op",
            "value": 896,
            "unit": "B/op",
            "extra": "708646 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOPCheckSerialNumberField - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "708646 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOPTerminalCityField",
            "value": 1586,
            "unit": "ns/op\t     896 B/op\t      10 allocs/op",
            "extra": "707618 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOPTerminalCityField - ns/op",
            "value": 1586,
            "unit": "ns/op",
            "extra": "707618 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOPTerminalCityField - B/op",
            "value": 896,
            "unit": "B/op",
            "extra": "707618 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOPTerminalCityField - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "707618 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOPTerminalStateField",
            "value": 1581,
            "unit": "ns/op\t     896 B/op\t      10 allocs/op",
            "extra": "701548 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOPTerminalStateField - ns/op",
            "value": 1581,
            "unit": "ns/op",
            "extra": "701548 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOPTerminalStateField - B/op",
            "value": 896,
            "unit": "B/op",
            "extra": "701548 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOPTerminalStateField - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "701548 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOPTransactionCode",
            "value": 1288,
            "unit": "ns/op\t     984 B/op\t      12 allocs/op",
            "extra": "837244 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOPTransactionCode - ns/op",
            "value": 1288,
            "unit": "ns/op",
            "extra": "837244 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOPTransactionCode - B/op",
            "value": 984,
            "unit": "B/op",
            "extra": "837244 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOPTransactionCode - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "837244 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOPAddendaCount",
            "value": 2396,
            "unit": "ns/op\t    1184 B/op\t      14 allocs/op",
            "extra": "481926 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOPAddendaCount - ns/op",
            "value": 2396,
            "unit": "ns/op",
            "extra": "481926 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOPAddendaCount - B/op",
            "value": 1184,
            "unit": "B/op",
            "extra": "481926 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOPAddendaCount - allocs/op",
            "value": 14,
            "unit": "allocs/op",
            "extra": "481926 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOPInvalidBuild",
            "value": 1667,
            "unit": "ns/op\t     960 B/op\t      11 allocs/op",
            "extra": "643213 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOPInvalidBuild - ns/op",
            "value": 1667,
            "unit": "ns/op",
            "extra": "643213 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOPInvalidBuild - B/op",
            "value": 960,
            "unit": "B/op",
            "extra": "643213 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOPInvalidBuild - allocs/op",
            "value": 11,
            "unit": "allocs/op",
            "extra": "643213 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOSHeader",
            "value": 170.3,
            "unit": "ns/op\t     448 B/op\t       3 allocs/op",
            "extra": "7111778 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOSHeader - ns/op",
            "value": 170.3,
            "unit": "ns/op",
            "extra": "7111778 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOSHeader - B/op",
            "value": 448,
            "unit": "B/op",
            "extra": "7111778 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOSHeader - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "7111778 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOSCreate",
            "value": 2543,
            "unit": "ns/op\t    1200 B/op\t      10 allocs/op",
            "extra": "452827 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOSCreate - ns/op",
            "value": 2543,
            "unit": "ns/op",
            "extra": "452827 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOSCreate - B/op",
            "value": 1200,
            "unit": "B/op",
            "extra": "452827 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOSCreate - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "452827 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOSStandardEntryClassCode",
            "value": 2559,
            "unit": "ns/op\t    1280 B/op\t      11 allocs/op",
            "extra": "444489 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOSStandardEntryClassCode - ns/op",
            "value": 2559,
            "unit": "ns/op",
            "extra": "444489 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOSStandardEntryClassCode - B/op",
            "value": 1280,
            "unit": "B/op",
            "extra": "444489 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOSStandardEntryClassCode - allocs/op",
            "value": 11,
            "unit": "allocs/op",
            "extra": "444489 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOSServiceClassCodeEquality",
            "value": 2647,
            "unit": "ns/op\t    1296 B/op\t      13 allocs/op",
            "extra": "438864 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOSServiceClassCodeEquality - ns/op",
            "value": 2647,
            "unit": "ns/op",
            "extra": "438864 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOSServiceClassCodeEquality - B/op",
            "value": 1296,
            "unit": "B/op",
            "extra": "438864 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOSServiceClassCodeEquality - allocs/op",
            "value": 13,
            "unit": "allocs/op",
            "extra": "438864 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOSMixedCreditsAndDebits",
            "value": 2668,
            "unit": "ns/op\t    1296 B/op\t      13 allocs/op",
            "extra": "419683 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOSMixedCreditsAndDebits - ns/op",
            "value": 2668,
            "unit": "ns/op",
            "extra": "419683 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOSMixedCreditsAndDebits - B/op",
            "value": 1296,
            "unit": "B/op",
            "extra": "419683 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOSMixedCreditsAndDebits - allocs/op",
            "value": 13,
            "unit": "allocs/op",
            "extra": "419683 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOSCreditsOnly",
            "value": 2661,
            "unit": "ns/op\t    1296 B/op\t      13 allocs/op",
            "extra": "430728 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOSCreditsOnly - ns/op",
            "value": 2661,
            "unit": "ns/op",
            "extra": "430728 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOSCreditsOnly - B/op",
            "value": 1296,
            "unit": "B/op",
            "extra": "430728 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOSCreditsOnly - allocs/op",
            "value": 13,
            "unit": "allocs/op",
            "extra": "430728 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOSAutomatedAccountingAdvices",
            "value": 2675,
            "unit": "ns/op\t    1304 B/op\t      14 allocs/op",
            "extra": "431646 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOSAutomatedAccountingAdvices - ns/op",
            "value": 2675,
            "unit": "ns/op",
            "extra": "431646 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOSAutomatedAccountingAdvices - B/op",
            "value": 1304,
            "unit": "B/op",
            "extra": "431646 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOSAutomatedAccountingAdvices - allocs/op",
            "value": 14,
            "unit": "allocs/op",
            "extra": "431646 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOSAddendaCount",
            "value": 2618,
            "unit": "ns/op\t    1408 B/op\t      11 allocs/op",
            "extra": "435421 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOSAddendaCount - ns/op",
            "value": 2618,
            "unit": "ns/op",
            "extra": "435421 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOSAddendaCount - B/op",
            "value": 1408,
            "unit": "B/op",
            "extra": "435421 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOSAddendaCount - allocs/op",
            "value": 11,
            "unit": "allocs/op",
            "extra": "435421 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOSAddendaCountZero",
            "value": 1337,
            "unit": "ns/op\t     944 B/op\t       8 allocs/op",
            "extra": "843064 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOSAddendaCountZero - ns/op",
            "value": 1337,
            "unit": "ns/op",
            "extra": "843064 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOSAddendaCountZero - B/op",
            "value": 944,
            "unit": "B/op",
            "extra": "843064 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOSAddendaCountZero - allocs/op",
            "value": 8,
            "unit": "allocs/op",
            "extra": "843064 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOSInvalidAddendum",
            "value": 1686,
            "unit": "ns/op\t    1136 B/op\t      13 allocs/op",
            "extra": "643791 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOSInvalidAddendum - ns/op",
            "value": 1686,
            "unit": "ns/op",
            "extra": "643791 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOSInvalidAddendum - B/op",
            "value": 1136,
            "unit": "B/op",
            "extra": "643791 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOSInvalidAddendum - allocs/op",
            "value": 13,
            "unit": "allocs/op",
            "extra": "643791 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOSInvalidAddenda",
            "value": 1066,
            "unit": "ns/op\t    1104 B/op\t      11 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOSInvalidAddenda - ns/op",
            "value": 1066,
            "unit": "ns/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOSInvalidAddenda - B/op",
            "value": 1104,
            "unit": "B/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOSInvalidAddenda - allocs/op",
            "value": 11,
            "unit": "allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOSInvalidBuild",
            "value": 1811,
            "unit": "ns/op\t    1136 B/op\t      10 allocs/op",
            "extra": "616629 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOSInvalidBuild - ns/op",
            "value": 1811,
            "unit": "ns/op",
            "extra": "616629 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOSInvalidBuild - B/op",
            "value": 1136,
            "unit": "B/op",
            "extra": "616629 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOSInvalidBuild - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "616629 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOSCardTransactionType",
            "value": 2368,
            "unit": "ns/op\t    1168 B/op\t      11 allocs/op",
            "extra": "486494 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOSCardTransactionType - ns/op",
            "value": 2368,
            "unit": "ns/op",
            "extra": "486494 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOSCardTransactionType - B/op",
            "value": 1168,
            "unit": "B/op",
            "extra": "486494 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOSCardTransactionType - allocs/op",
            "value": 11,
            "unit": "allocs/op",
            "extra": "486494 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOSMixedDebitsAndCredits",
            "value": 4727,
            "unit": "ns/op\t    1704 B/op\t      15 allocs/op",
            "extra": "246001 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOSMixedDebitsAndCredits - ns/op",
            "value": 4727,
            "unit": "ns/op",
            "extra": "246001 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOSMixedDebitsAndCredits - B/op",
            "value": 1704,
            "unit": "B/op",
            "extra": "246001 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPOSMixedDebitsAndCredits - allocs/op",
            "value": 15,
            "unit": "allocs/op",
            "extra": "246001 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchError",
            "value": 291.6,
            "unit": "ns/op\t     160 B/op\t       3 allocs/op",
            "extra": "4108681 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchError - ns/op",
            "value": 291.6,
            "unit": "ns/op",
            "extra": "4108681 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchError - B/op",
            "value": 160,
            "unit": "B/op",
            "extra": "4108681 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchError - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "4108681 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchServiceClassCodeEquality",
            "value": 2953,
            "unit": "ns/op\t    1296 B/op\t      15 allocs/op",
            "extra": "391899 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchServiceClassCodeEquality - ns/op",
            "value": 2953,
            "unit": "ns/op",
            "extra": "391899 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchServiceClassCodeEquality - B/op",
            "value": 1296,
            "unit": "B/op",
            "extra": "391899 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchServiceClassCodeEquality - allocs/op",
            "value": 15,
            "unit": "allocs/op",
            "extra": "391899 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPPDCreate",
            "value": 2271,
            "unit": "ns/op\t    1152 B/op\t      13 allocs/op",
            "extra": "502646 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPPDCreate - ns/op",
            "value": 2271,
            "unit": "ns/op",
            "extra": "502646 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPPDCreate - B/op",
            "value": 1152,
            "unit": "B/op",
            "extra": "502646 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPPDCreate - allocs/op",
            "value": 13,
            "unit": "allocs/op",
            "extra": "502646 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPPDTypeCode",
            "value": 2940,
            "unit": "ns/op\t    1448 B/op\t      17 allocs/op",
            "extra": "390421 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPPDTypeCode - ns/op",
            "value": 2940,
            "unit": "ns/op",
            "extra": "390421 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPPDTypeCode - B/op",
            "value": 1448,
            "unit": "B/op",
            "extra": "390421 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPPDTypeCode - allocs/op",
            "value": 17,
            "unit": "allocs/op",
            "extra": "390421 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCompanyIdentification",
            "value": 2929,
            "unit": "ns/op\t    1328 B/op\t      17 allocs/op",
            "extra": "394405 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCompanyIdentification - ns/op",
            "value": 2929,
            "unit": "ns/op",
            "extra": "394405 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCompanyIdentification - B/op",
            "value": 1328,
            "unit": "B/op",
            "extra": "394405 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCompanyIdentification - allocs/op",
            "value": 17,
            "unit": "allocs/op",
            "extra": "394405 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchODFIIDMismatch",
            "value": 2950,
            "unit": "ns/op\t    1360 B/op\t      17 allocs/op",
            "extra": "390754 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchODFIIDMismatch - ns/op",
            "value": 2950,
            "unit": "ns/op",
            "extra": "390754 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchODFIIDMismatch - B/op",
            "value": 1360,
            "unit": "B/op",
            "extra": "390754 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchODFIIDMismatch - allocs/op",
            "value": 17,
            "unit": "allocs/op",
            "extra": "390754 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchBuild",
            "value": 1841,
            "unit": "ns/op\t    1040 B/op\t      13 allocs/op",
            "extra": "600157 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchBuild - ns/op",
            "value": 1841,
            "unit": "ns/op",
            "extra": "600157 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchBuild - B/op",
            "value": 1040,
            "unit": "B/op",
            "extra": "600157 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchBuild - allocs/op",
            "value": 13,
            "unit": "allocs/op",
            "extra": "600157 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPPDAddendaCount",
            "value": 3249,
            "unit": "ns/op\t    1512 B/op\t      19 allocs/op",
            "extra": "358474 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPPDAddendaCount - ns/op",
            "value": 3249,
            "unit": "ns/op",
            "extra": "358474 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPPDAddendaCount - B/op",
            "value": 1512,
            "unit": "B/op",
            "extra": "358474 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchPPDAddendaCount - allocs/op",
            "value": 19,
            "unit": "allocs/op",
            "extra": "358474 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchRCKHeader",
            "value": 169.8,
            "unit": "ns/op\t     448 B/op\t       3 allocs/op",
            "extra": "7023342 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchRCKHeader - ns/op",
            "value": 169.8,
            "unit": "ns/op",
            "extra": "7023342 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchRCKHeader - B/op",
            "value": 448,
            "unit": "B/op",
            "extra": "7023342 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchRCKHeader - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "7023342 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchRCKCreate",
            "value": 2004,
            "unit": "ns/op\t     992 B/op\t       9 allocs/op",
            "extra": "555789 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchRCKCreate - ns/op",
            "value": 2004,
            "unit": "ns/op",
            "extra": "555789 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchRCKCreate - B/op",
            "value": 992,
            "unit": "B/op",
            "extra": "555789 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchRCKCreate - allocs/op",
            "value": 9,
            "unit": "allocs/op",
            "extra": "555789 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchRCKStandardEntryClassCode",
            "value": 2069,
            "unit": "ns/op\t    1072 B/op\t      10 allocs/op",
            "extra": "536260 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchRCKStandardEntryClassCode - ns/op",
            "value": 2069,
            "unit": "ns/op",
            "extra": "536260 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchRCKStandardEntryClassCode - B/op",
            "value": 1072,
            "unit": "B/op",
            "extra": "536260 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchRCKStandardEntryClassCode - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "536260 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchRCKServiceClassCodeEquality",
            "value": 2169,
            "unit": "ns/op\t    1088 B/op\t      12 allocs/op",
            "extra": "519982 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchRCKServiceClassCodeEquality - ns/op",
            "value": 2169,
            "unit": "ns/op",
            "extra": "519982 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchRCKServiceClassCodeEquality - B/op",
            "value": 1088,
            "unit": "B/op",
            "extra": "519982 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchRCKServiceClassCodeEquality - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "519982 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchRCKMixedCreditsAndDebits",
            "value": 2169,
            "unit": "ns/op\t    1088 B/op\t      12 allocs/op",
            "extra": "520822 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchRCKMixedCreditsAndDebits - ns/op",
            "value": 2169,
            "unit": "ns/op",
            "extra": "520822 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchRCKMixedCreditsAndDebits - B/op",
            "value": 1088,
            "unit": "B/op",
            "extra": "520822 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchRCKMixedCreditsAndDebits - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "520822 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchRCKCreditsOnly",
            "value": 2174,
            "unit": "ns/op\t    1088 B/op\t      12 allocs/op",
            "extra": "523543 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchRCKCreditsOnly - ns/op",
            "value": 2174,
            "unit": "ns/op",
            "extra": "523543 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchRCKCreditsOnly - B/op",
            "value": 1088,
            "unit": "B/op",
            "extra": "523543 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchRCKCreditsOnly - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "523543 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchRCKAutomatedAccountingAdvices",
            "value": 2170,
            "unit": "ns/op\t    1096 B/op\t      13 allocs/op",
            "extra": "516405 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchRCKAutomatedAccountingAdvices - ns/op",
            "value": 2170,
            "unit": "ns/op",
            "extra": "516405 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchRCKAutomatedAccountingAdvices - B/op",
            "value": 1096,
            "unit": "B/op",
            "extra": "516405 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchRCKAutomatedAccountingAdvices - allocs/op",
            "value": 13,
            "unit": "allocs/op",
            "extra": "516405 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchRCKCompanyEntryDescription",
            "value": 2119,
            "unit": "ns/op\t    1088 B/op\t      11 allocs/op",
            "extra": "540631 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchRCKCompanyEntryDescription - ns/op",
            "value": 2119,
            "unit": "ns/op",
            "extra": "540631 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchRCKCompanyEntryDescription - B/op",
            "value": 1088,
            "unit": "B/op",
            "extra": "540631 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchRCKCompanyEntryDescription - allocs/op",
            "value": 11,
            "unit": "allocs/op",
            "extra": "540631 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchRCKAmount",
            "value": 2670,
            "unit": "ns/op\t    1296 B/op\t      17 allocs/op",
            "extra": "426943 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchRCKAmount - ns/op",
            "value": 2670,
            "unit": "ns/op",
            "extra": "426943 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchRCKAmount - B/op",
            "value": 1296,
            "unit": "B/op",
            "extra": "426943 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchRCKAmount - allocs/op",
            "value": 17,
            "unit": "allocs/op",
            "extra": "426943 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchRCKCheckSerialNumber",
            "value": 1814,
            "unit": "ns/op\t     944 B/op\t       9 allocs/op",
            "extra": "619560 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchRCKCheckSerialNumber - ns/op",
            "value": 1814,
            "unit": "ns/op",
            "extra": "619560 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchRCKCheckSerialNumber - B/op",
            "value": 944,
            "unit": "B/op",
            "extra": "619560 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchRCKCheckSerialNumber - allocs/op",
            "value": 9,
            "unit": "allocs/op",
            "extra": "619560 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchRCKTransactionCode",
            "value": 1120,
            "unit": "ns/op\t     944 B/op\t       9 allocs/op",
            "extra": "982364 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchRCKTransactionCode - ns/op",
            "value": 1120,
            "unit": "ns/op",
            "extra": "982364 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchRCKTransactionCode - B/op",
            "value": 944,
            "unit": "B/op",
            "extra": "982364 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchRCKTransactionCode - allocs/op",
            "value": 9,
            "unit": "allocs/op",
            "extra": "982364 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchRCKAddendaCount",
            "value": 2432,
            "unit": "ns/op\t    1184 B/op\t      14 allocs/op",
            "extra": "467908 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchRCKAddendaCount - ns/op",
            "value": 2432,
            "unit": "ns/op",
            "extra": "467908 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchRCKAddendaCount - B/op",
            "value": 1184,
            "unit": "B/op",
            "extra": "467908 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchRCKAddendaCount - allocs/op",
            "value": 14,
            "unit": "allocs/op",
            "extra": "467908 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchRCKParseCheckSerialNumber",
            "value": 2060,
            "unit": "ns/op\t    1008 B/op\t      10 allocs/op",
            "extra": "550920 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchRCKParseCheckSerialNumber - ns/op",
            "value": 2060,
            "unit": "ns/op",
            "extra": "550920 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchRCKParseCheckSerialNumber - B/op",
            "value": 1008,
            "unit": "B/op",
            "extra": "550920 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchRCKParseCheckSerialNumber - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "550920 times\n4 procs"
          },
          {
            "name": "BenchmarkRCKBatchInvalidBuild",
            "value": 1499,
            "unit": "ns/op\t     928 B/op\t       9 allocs/op",
            "extra": "762660 times\n4 procs"
          },
          {
            "name": "BenchmarkRCKBatchInvalidBuild - ns/op",
            "value": 1499,
            "unit": "ns/op",
            "extra": "762660 times\n4 procs"
          },
          {
            "name": "BenchmarkRCKBatchInvalidBuild - B/op",
            "value": 928,
            "unit": "B/op",
            "extra": "762660 times\n4 procs"
          },
          {
            "name": "BenchmarkRCKBatchInvalidBuild - allocs/op",
            "value": 9,
            "unit": "allocs/op",
            "extra": "762660 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchSHRHeader",
            "value": 170.5,
            "unit": "ns/op\t     448 B/op\t       3 allocs/op",
            "extra": "6922866 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchSHRHeader - ns/op",
            "value": 170.5,
            "unit": "ns/op",
            "extra": "6922866 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchSHRHeader - B/op",
            "value": 448,
            "unit": "B/op",
            "extra": "6922866 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchSHRHeader - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "6922866 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchSHRCreate",
            "value": 2925,
            "unit": "ns/op\t    1240 B/op\t      12 allocs/op",
            "extra": "400976 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchSHRCreate - ns/op",
            "value": 2925,
            "unit": "ns/op",
            "extra": "400976 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchSHRCreate - B/op",
            "value": 1240,
            "unit": "B/op",
            "extra": "400976 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchSHRCreate - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "400976 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchSHRStandardEntryClassCode",
            "value": 2862,
            "unit": "ns/op\t    1320 B/op\t      13 allocs/op",
            "extra": "406784 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchSHRStandardEntryClassCode - ns/op",
            "value": 2862,
            "unit": "ns/op",
            "extra": "406784 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchSHRStandardEntryClassCode - B/op",
            "value": 1320,
            "unit": "B/op",
            "extra": "406784 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchSHRStandardEntryClassCode - allocs/op",
            "value": 13,
            "unit": "allocs/op",
            "extra": "406784 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchSHRServiceClassCodeEquality",
            "value": 2924,
            "unit": "ns/op\t    1336 B/op\t      15 allocs/op",
            "extra": "348658 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchSHRServiceClassCodeEquality - ns/op",
            "value": 2924,
            "unit": "ns/op",
            "extra": "348658 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchSHRServiceClassCodeEquality - B/op",
            "value": 1336,
            "unit": "B/op",
            "extra": "348658 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchSHRServiceClassCodeEquality - allocs/op",
            "value": 15,
            "unit": "allocs/op",
            "extra": "348658 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchSHRMixedCreditsAndDebits",
            "value": 2943,
            "unit": "ns/op\t    1336 B/op\t      15 allocs/op",
            "extra": "393546 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchSHRMixedCreditsAndDebits - ns/op",
            "value": 2943,
            "unit": "ns/op",
            "extra": "393546 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchSHRMixedCreditsAndDebits - B/op",
            "value": 1336,
            "unit": "B/op",
            "extra": "393546 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchSHRMixedCreditsAndDebits - allocs/op",
            "value": 15,
            "unit": "allocs/op",
            "extra": "393546 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchSHRCreditsOnly",
            "value": 2930,
            "unit": "ns/op\t    1336 B/op\t      15 allocs/op",
            "extra": "380215 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchSHRCreditsOnly - ns/op",
            "value": 2930,
            "unit": "ns/op",
            "extra": "380215 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchSHRCreditsOnly - B/op",
            "value": 1336,
            "unit": "B/op",
            "extra": "380215 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchSHRCreditsOnly - allocs/op",
            "value": 15,
            "unit": "allocs/op",
            "extra": "380215 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchSHRAutomatedAccountingAdvices",
            "value": 2937,
            "unit": "ns/op\t    1344 B/op\t      16 allocs/op",
            "extra": "393130 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchSHRAutomatedAccountingAdvices - ns/op",
            "value": 2937,
            "unit": "ns/op",
            "extra": "393130 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchSHRAutomatedAccountingAdvices - B/op",
            "value": 1344,
            "unit": "B/op",
            "extra": "393130 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchSHRAutomatedAccountingAdvices - allocs/op",
            "value": 16,
            "unit": "allocs/op",
            "extra": "393130 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchSHRTransactionCode",
            "value": 2858,
            "unit": "ns/op\t    1320 B/op\t      13 allocs/op",
            "extra": "398882 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchSHRTransactionCode - ns/op",
            "value": 2858,
            "unit": "ns/op",
            "extra": "398882 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchSHRTransactionCode - B/op",
            "value": 1320,
            "unit": "B/op",
            "extra": "398882 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchSHRTransactionCode - allocs/op",
            "value": 13,
            "unit": "allocs/op",
            "extra": "398882 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchSHRAddendaCount",
            "value": 2991,
            "unit": "ns/op\t    1448 B/op\t      13 allocs/op",
            "extra": "389152 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchSHRAddendaCount - ns/op",
            "value": 2991,
            "unit": "ns/op",
            "extra": "389152 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchSHRAddendaCount - B/op",
            "value": 1448,
            "unit": "B/op",
            "extra": "389152 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchSHRAddendaCount - allocs/op",
            "value": 13,
            "unit": "allocs/op",
            "extra": "389152 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchSHRAddendaCountZero",
            "value": 1490,
            "unit": "ns/op\t    1080 B/op\t      13 allocs/op",
            "extra": "760218 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchSHRAddendaCountZero - ns/op",
            "value": 1490,
            "unit": "ns/op",
            "extra": "760218 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchSHRAddendaCountZero - B/op",
            "value": 1080,
            "unit": "B/op",
            "extra": "760218 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchSHRAddendaCountZero - allocs/op",
            "value": 13,
            "unit": "allocs/op",
            "extra": "760218 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchSHRInvalidAddendum",
            "value": 1936,
            "unit": "ns/op\t    1176 B/op\t      15 allocs/op",
            "extra": "572733 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchSHRInvalidAddendum - ns/op",
            "value": 1936,
            "unit": "ns/op",
            "extra": "572733 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchSHRInvalidAddendum - B/op",
            "value": 1176,
            "unit": "B/op",
            "extra": "572733 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchSHRInvalidAddendum - allocs/op",
            "value": 15,
            "unit": "allocs/op",
            "extra": "572733 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchSHRInvalidAddenda",
            "value": 1223,
            "unit": "ns/op\t    1144 B/op\t      13 allocs/op",
            "extra": "901797 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchSHRInvalidAddenda - ns/op",
            "value": 1223,
            "unit": "ns/op",
            "extra": "901797 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchSHRInvalidAddenda - B/op",
            "value": 1144,
            "unit": "B/op",
            "extra": "901797 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchSHRInvalidAddenda - allocs/op",
            "value": 13,
            "unit": "allocs/op",
            "extra": "901797 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchSHRInvalidBuild",
            "value": 2081,
            "unit": "ns/op\t    1176 B/op\t      12 allocs/op",
            "extra": "552744 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchSHRInvalidBuild - ns/op",
            "value": 2081,
            "unit": "ns/op",
            "extra": "552744 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchSHRInvalidBuild - B/op",
            "value": 1176,
            "unit": "B/op",
            "extra": "552744 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchSHRInvalidBuild - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "552744 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchSHRCardTransactionType",
            "value": 2631,
            "unit": "ns/op\t    1208 B/op\t      13 allocs/op",
            "extra": "436702 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchSHRCardTransactionType - ns/op",
            "value": 2631,
            "unit": "ns/op",
            "extra": "436702 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchSHRCardTransactionType - B/op",
            "value": 1208,
            "unit": "B/op",
            "extra": "436702 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchSHRCardTransactionType - allocs/op",
            "value": 13,
            "unit": "allocs/op",
            "extra": "436702 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchSHRCardExpirationDateField",
            "value": 2006,
            "unit": "ns/op\t    1112 B/op\t      11 allocs/op",
            "extra": "570368 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchSHRCardExpirationDateField - ns/op",
            "value": 2006,
            "unit": "ns/op",
            "extra": "570368 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchSHRCardExpirationDateField - B/op",
            "value": 1112,
            "unit": "B/op",
            "extra": "570368 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchSHRCardExpirationDateField - allocs/op",
            "value": 11,
            "unit": "allocs/op",
            "extra": "570368 times\n4 procs"
          },
          {
            "name": "BenchmarkSHRDocumentReferenceNumberField",
            "value": 1998,
            "unit": "ns/op\t    1112 B/op\t      11 allocs/op",
            "extra": "551803 times\n4 procs"
          },
          {
            "name": "BenchmarkSHRDocumentReferenceNumberField - ns/op",
            "value": 1998,
            "unit": "ns/op",
            "extra": "551803 times\n4 procs"
          },
          {
            "name": "BenchmarkSHRDocumentReferenceNumberField - B/op",
            "value": 1112,
            "unit": "B/op",
            "extra": "551803 times\n4 procs"
          },
          {
            "name": "BenchmarkSHRDocumentReferenceNumberField - allocs/op",
            "value": 11,
            "unit": "allocs/op",
            "extra": "551803 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchSHRDocumentReferenceNumberField",
            "value": 2017,
            "unit": "ns/op\t    1112 B/op\t      11 allocs/op",
            "extra": "552649 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchSHRDocumentReferenceNumberField - ns/op",
            "value": 2017,
            "unit": "ns/op",
            "extra": "552649 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchSHRDocumentReferenceNumberField - B/op",
            "value": 1112,
            "unit": "B/op",
            "extra": "552649 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchSHRDocumentReferenceNumberField - allocs/op",
            "value": 11,
            "unit": "allocs/op",
            "extra": "552649 times\n4 procs"
          },
          {
            "name": "BenchmarkSHRCardExpirationDateMonth",
            "value": 2687,
            "unit": "ns/op\t    1192 B/op\t      13 allocs/op",
            "extra": "433074 times\n4 procs"
          },
          {
            "name": "BenchmarkSHRCardExpirationDateMonth - ns/op",
            "value": 2687,
            "unit": "ns/op",
            "extra": "433074 times\n4 procs"
          },
          {
            "name": "BenchmarkSHRCardExpirationDateMonth - B/op",
            "value": 1192,
            "unit": "B/op",
            "extra": "433074 times\n4 procs"
          },
          {
            "name": "BenchmarkSHRCardExpirationDateMonth - allocs/op",
            "value": 13,
            "unit": "allocs/op",
            "extra": "433074 times\n4 procs"
          },
          {
            "name": "BenchmarkSHRCardExpirationDateYear",
            "value": 2700,
            "unit": "ns/op\t    1192 B/op\t      13 allocs/op",
            "extra": "426614 times\n4 procs"
          },
          {
            "name": "BenchmarkSHRCardExpirationDateYear - ns/op",
            "value": 2700,
            "unit": "ns/op",
            "extra": "426614 times\n4 procs"
          },
          {
            "name": "BenchmarkSHRCardExpirationDateYear - B/op",
            "value": 1192,
            "unit": "B/op",
            "extra": "426614 times\n4 procs"
          },
          {
            "name": "BenchmarkSHRCardExpirationDateYear - allocs/op",
            "value": 13,
            "unit": "allocs/op",
            "extra": "426614 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTELHeader",
            "value": 170.4,
            "unit": "ns/op\t     448 B/op\t       3 allocs/op",
            "extra": "7072216 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTELHeader - ns/op",
            "value": 170.4,
            "unit": "ns/op",
            "extra": "7072216 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTELHeader - B/op",
            "value": 448,
            "unit": "B/op",
            "extra": "7072216 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTELHeader - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "7072216 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTELCreate",
            "value": 1562,
            "unit": "ns/op\t     944 B/op\t      10 allocs/op",
            "extra": "705778 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTELCreate - ns/op",
            "value": 1562,
            "unit": "ns/op",
            "extra": "705778 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTELCreate - B/op",
            "value": 944,
            "unit": "B/op",
            "extra": "705778 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTELCreate - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "705778 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTELAddendaCount",
            "value": 2523,
            "unit": "ns/op\t    1328 B/op\t      13 allocs/op",
            "extra": "449829 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTELAddendaCount - ns/op",
            "value": 2523,
            "unit": "ns/op",
            "extra": "449829 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTELAddendaCount - B/op",
            "value": 1328,
            "unit": "B/op",
            "extra": "449829 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTELAddendaCount - allocs/op",
            "value": 13,
            "unit": "allocs/op",
            "extra": "449829 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTELSEC",
            "value": 1861,
            "unit": "ns/op\t     944 B/op\t       9 allocs/op",
            "extra": "610671 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTELSEC - ns/op",
            "value": 1861,
            "unit": "ns/op",
            "extra": "610671 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTELSEC - B/op",
            "value": 944,
            "unit": "B/op",
            "extra": "610671 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTELSEC - allocs/op",
            "value": 9,
            "unit": "allocs/op",
            "extra": "610671 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTELDebit",
            "value": 2126,
            "unit": "ns/op\t    1072 B/op\t      10 allocs/op",
            "extra": "539666 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTELDebit - ns/op",
            "value": 2126,
            "unit": "ns/op",
            "extra": "539666 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTELDebit - B/op",
            "value": 1072,
            "unit": "B/op",
            "extra": "539666 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTELDebit - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "539666 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTELPaymentType",
            "value": 1803,
            "unit": "ns/op\t     864 B/op\t       8 allocs/op",
            "extra": "634194 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTELPaymentType - ns/op",
            "value": 1803,
            "unit": "ns/op",
            "extra": "634194 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTELPaymentType - B/op",
            "value": 864,
            "unit": "B/op",
            "extra": "634194 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTELPaymentType - allocs/op",
            "value": 8,
            "unit": "allocs/op",
            "extra": "634194 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRCHeader",
            "value": 171,
            "unit": "ns/op\t     448 B/op\t       3 allocs/op",
            "extra": "6939392 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRCHeader - ns/op",
            "value": 171,
            "unit": "ns/op",
            "extra": "6939392 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRCHeader - B/op",
            "value": 448,
            "unit": "B/op",
            "extra": "6939392 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRCHeader - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "6939392 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRCCreate",
            "value": 2195,
            "unit": "ns/op\t    1032 B/op\t      11 allocs/op",
            "extra": "520028 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRCCreate - ns/op",
            "value": 2195,
            "unit": "ns/op",
            "extra": "520028 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRCCreate - B/op",
            "value": 1032,
            "unit": "B/op",
            "extra": "520028 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRCCreate - allocs/op",
            "value": 11,
            "unit": "allocs/op",
            "extra": "520028 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRCStandardEntryClassCode",
            "value": 2244,
            "unit": "ns/op\t    1112 B/op\t      12 allocs/op",
            "extra": "503740 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRCStandardEntryClassCode - ns/op",
            "value": 2244,
            "unit": "ns/op",
            "extra": "503740 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRCStandardEntryClassCode - B/op",
            "value": 1112,
            "unit": "B/op",
            "extra": "503740 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRCStandardEntryClassCode - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "503740 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRCServiceClassCodeEquality",
            "value": 2336,
            "unit": "ns/op\t    1128 B/op\t      14 allocs/op",
            "extra": "491994 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRCServiceClassCodeEquality - ns/op",
            "value": 2336,
            "unit": "ns/op",
            "extra": "491994 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRCServiceClassCodeEquality - B/op",
            "value": 1128,
            "unit": "B/op",
            "extra": "491994 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRCServiceClassCodeEquality - allocs/op",
            "value": 14,
            "unit": "allocs/op",
            "extra": "491994 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRCMixedCreditsAndDebits",
            "value": 2329,
            "unit": "ns/op\t    1128 B/op\t      14 allocs/op",
            "extra": "486949 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRCMixedCreditsAndDebits - ns/op",
            "value": 2329,
            "unit": "ns/op",
            "extra": "486949 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRCMixedCreditsAndDebits - B/op",
            "value": 1128,
            "unit": "B/op",
            "extra": "486949 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRCMixedCreditsAndDebits - allocs/op",
            "value": 14,
            "unit": "allocs/op",
            "extra": "486949 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRCCreditsOnly",
            "value": 2335,
            "unit": "ns/op\t    1128 B/op\t      14 allocs/op",
            "extra": "491170 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRCCreditsOnly - ns/op",
            "value": 2335,
            "unit": "ns/op",
            "extra": "491170 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRCCreditsOnly - B/op",
            "value": 1128,
            "unit": "B/op",
            "extra": "491170 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRCCreditsOnly - allocs/op",
            "value": 14,
            "unit": "allocs/op",
            "extra": "491170 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRCAutomatedAccountingAdvices",
            "value": 2340,
            "unit": "ns/op\t    1136 B/op\t      15 allocs/op",
            "extra": "482026 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRCAutomatedAccountingAdvices - ns/op",
            "value": 2340,
            "unit": "ns/op",
            "extra": "482026 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRCAutomatedAccountingAdvices - B/op",
            "value": 1136,
            "unit": "B/op",
            "extra": "482026 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRCAutomatedAccountingAdvices - allocs/op",
            "value": 15,
            "unit": "allocs/op",
            "extra": "482026 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRCCheckSerialNumber",
            "value": 1928,
            "unit": "ns/op\t     904 B/op\t      10 allocs/op",
            "extra": "589221 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRCCheckSerialNumber - ns/op",
            "value": 1928,
            "unit": "ns/op",
            "extra": "589221 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRCCheckSerialNumber - B/op",
            "value": 904,
            "unit": "B/op",
            "extra": "589221 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRCCheckSerialNumber - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "589221 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRCTransactionCode",
            "value": 1239,
            "unit": "ns/op\t     984 B/op\t      11 allocs/op",
            "extra": "877894 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRCTransactionCode - ns/op",
            "value": 1239,
            "unit": "ns/op",
            "extra": "877894 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRCTransactionCode - B/op",
            "value": 984,
            "unit": "B/op",
            "extra": "877894 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRCTransactionCode - allocs/op",
            "value": 11,
            "unit": "allocs/op",
            "extra": "877894 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRCAddendaCount",
            "value": 2612,
            "unit": "ns/op\t    1224 B/op\t      16 allocs/op",
            "extra": "447073 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRCAddendaCount - ns/op",
            "value": 2612,
            "unit": "ns/op",
            "extra": "447073 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRCAddendaCount - B/op",
            "value": 1224,
            "unit": "B/op",
            "extra": "447073 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRCAddendaCount - allocs/op",
            "value": 16,
            "unit": "allocs/op",
            "extra": "447073 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRCInvalidBuild",
            "value": 1661,
            "unit": "ns/op\t     968 B/op\t      11 allocs/op",
            "extra": "677341 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRCInvalidBuild - ns/op",
            "value": 1661,
            "unit": "ns/op",
            "extra": "677341 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRCInvalidBuild - B/op",
            "value": 968,
            "unit": "B/op",
            "extra": "677341 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRCInvalidBuild - allocs/op",
            "value": 11,
            "unit": "allocs/op",
            "extra": "677341 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRXHeader",
            "value": 172.1,
            "unit": "ns/op\t     448 B/op\t       3 allocs/op",
            "extra": "7064368 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRXHeader - ns/op",
            "value": 172.1,
            "unit": "ns/op",
            "extra": "7064368 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRXHeader - B/op",
            "value": 448,
            "unit": "B/op",
            "extra": "7064368 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRXHeader - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "7064368 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRXCreate",
            "value": 2494,
            "unit": "ns/op\t    1040 B/op\t      16 allocs/op",
            "extra": "455558 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRXCreate - ns/op",
            "value": 2494,
            "unit": "ns/op",
            "extra": "455558 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRXCreate - B/op",
            "value": 1040,
            "unit": "B/op",
            "extra": "455558 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRXCreate - allocs/op",
            "value": 16,
            "unit": "allocs/op",
            "extra": "455558 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRXStandardEntryClassCode",
            "value": 2849,
            "unit": "ns/op\t    1248 B/op\t      18 allocs/op",
            "extra": "405314 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRXStandardEntryClassCode - ns/op",
            "value": 2849,
            "unit": "ns/op",
            "extra": "405314 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRXStandardEntryClassCode - B/op",
            "value": 1248,
            "unit": "B/op",
            "extra": "405314 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRXStandardEntryClassCode - allocs/op",
            "value": 18,
            "unit": "allocs/op",
            "extra": "405314 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRXServiceClassCodeEquality",
            "value": 2819,
            "unit": "ns/op\t    1256 B/op\t      19 allocs/op",
            "extra": "395876 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRXServiceClassCodeEquality - ns/op",
            "value": 2819,
            "unit": "ns/op",
            "extra": "395876 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRXServiceClassCodeEquality - B/op",
            "value": 1256,
            "unit": "B/op",
            "extra": "395876 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRXServiceClassCodeEquality - allocs/op",
            "value": 19,
            "unit": "allocs/op",
            "extra": "395876 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRXAddendaCount",
            "value": 3624,
            "unit": "ns/op\t    1512 B/op\t      24 allocs/op",
            "extra": "323715 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRXAddendaCount - ns/op",
            "value": 3624,
            "unit": "ns/op",
            "extra": "323715 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRXAddendaCount - B/op",
            "value": 1512,
            "unit": "B/op",
            "extra": "323715 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRXAddendaCount - allocs/op",
            "value": 24,
            "unit": "allocs/op",
            "extra": "323715 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRXAddendaCountZero",
            "value": 1837,
            "unit": "ns/op\t    1048 B/op\t      15 allocs/op",
            "extra": "605624 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRXAddendaCountZero - ns/op",
            "value": 1837,
            "unit": "ns/op",
            "extra": "605624 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRXAddendaCountZero - B/op",
            "value": 1048,
            "unit": "B/op",
            "extra": "605624 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRXAddendaCountZero - allocs/op",
            "value": 15,
            "unit": "allocs/op",
            "extra": "605624 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRXInvalidAddendum",
            "value": 2126,
            "unit": "ns/op\t    1256 B/op\t      16 allocs/op",
            "extra": "492058 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRXInvalidAddendum - ns/op",
            "value": 2126,
            "unit": "ns/op",
            "extra": "492058 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRXInvalidAddendum - B/op",
            "value": 1256,
            "unit": "B/op",
            "extra": "492058 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRXInvalidAddendum - allocs/op",
            "value": 16,
            "unit": "allocs/op",
            "extra": "492058 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRXInvalidAddenda",
            "value": 1355,
            "unit": "ns/op\t    1056 B/op\t      16 allocs/op",
            "extra": "799663 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRXInvalidAddenda - ns/op",
            "value": 1355,
            "unit": "ns/op",
            "extra": "799663 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRXInvalidAddenda - B/op",
            "value": 1056,
            "unit": "B/op",
            "extra": "799663 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRXInvalidAddenda - allocs/op",
            "value": 16,
            "unit": "allocs/op",
            "extra": "799663 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRXInvalidBuild",
            "value": 2098,
            "unit": "ns/op\t    1096 B/op\t      16 allocs/op",
            "extra": "555800 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRXInvalidBuild - ns/op",
            "value": 2098,
            "unit": "ns/op",
            "extra": "555800 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRXInvalidBuild - B/op",
            "value": 1096,
            "unit": "B/op",
            "extra": "555800 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRXInvalidBuild - allocs/op",
            "value": 16,
            "unit": "allocs/op",
            "extra": "555800 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRXAddenda10000",
            "value": 2009563,
            "unit": "ns/op\t 1191852 B/op\t   20039 allocs/op",
            "extra": "592 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRXAddenda10000 - ns/op",
            "value": 2009563,
            "unit": "ns/op",
            "extra": "592 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRXAddenda10000 - B/op",
            "value": 1191852,
            "unit": "B/op",
            "extra": "592 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRXAddenda10000 - allocs/op",
            "value": 20039,
            "unit": "allocs/op",
            "extra": "592 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRXAddendaRecords",
            "value": 107924,
            "unit": "ns/op\t   60153 B/op\t    1160 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRXAddendaRecords - ns/op",
            "value": 107924,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRXAddendaRecords - B/op",
            "value": 60153,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRXAddendaRecords - allocs/op",
            "value": 1160,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRXReceivingCompany",
            "value": 2965,
            "unit": "ns/op\t    1072 B/op\t      17 allocs/op",
            "extra": "388650 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRXReceivingCompany - ns/op",
            "value": 2965,
            "unit": "ns/op",
            "extra": "388650 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRXReceivingCompany - B/op",
            "value": 1072,
            "unit": "B/op",
            "extra": "388650 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRXReceivingCompany - allocs/op",
            "value": 17,
            "unit": "allocs/op",
            "extra": "388650 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRXReserved",
            "value": 1975,
            "unit": "ns/op\t    1032 B/op\t      15 allocs/op",
            "extra": "565255 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRXReserved - ns/op",
            "value": 1975,
            "unit": "ns/op",
            "extra": "565255 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRXReserved - B/op",
            "value": 1032,
            "unit": "B/op",
            "extra": "565255 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRXReserved - allocs/op",
            "value": 15,
            "unit": "allocs/op",
            "extra": "565255 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRXZeroAddendaRecords",
            "value": 1821,
            "unit": "ns/op\t    1048 B/op\t      15 allocs/op",
            "extra": "631010 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRXZeroAddendaRecords - ns/op",
            "value": 1821,
            "unit": "ns/op",
            "extra": "631010 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRXZeroAddendaRecords - B/op",
            "value": 1048,
            "unit": "B/op",
            "extra": "631010 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRXZeroAddendaRecords - allocs/op",
            "value": 15,
            "unit": "allocs/op",
            "extra": "631010 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRXTransactionCode",
            "value": 1252,
            "unit": "ns/op\t     984 B/op\t      11 allocs/op",
            "extra": "880761 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRXTransactionCode - ns/op",
            "value": 1252,
            "unit": "ns/op",
            "extra": "880761 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRXTransactionCode - B/op",
            "value": 984,
            "unit": "B/op",
            "extra": "880761 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRXTransactionCode - allocs/op",
            "value": 11,
            "unit": "allocs/op",
            "extra": "880761 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRXCreditsOnly",
            "value": 2805,
            "unit": "ns/op\t    1256 B/op\t      19 allocs/op",
            "extra": "400156 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRXCreditsOnly - ns/op",
            "value": 2805,
            "unit": "ns/op",
            "extra": "400156 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRXCreditsOnly - B/op",
            "value": 1256,
            "unit": "B/op",
            "extra": "400156 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRXCreditsOnly - allocs/op",
            "value": 19,
            "unit": "allocs/op",
            "extra": "400156 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRXAutomatedAccountingAdvices",
            "value": 2805,
            "unit": "ns/op\t    1264 B/op\t      20 allocs/op",
            "extra": "404973 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRXAutomatedAccountingAdvices - ns/op",
            "value": 2805,
            "unit": "ns/op",
            "extra": "404973 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRXAutomatedAccountingAdvices - B/op",
            "value": 1264,
            "unit": "B/op",
            "extra": "404973 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTRXAutomatedAccountingAdvices - allocs/op",
            "value": 20,
            "unit": "allocs/op",
            "extra": "404973 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchWebAddenda",
            "value": 3293,
            "unit": "ns/op\t    1440 B/op\t      20 allocs/op",
            "extra": "350802 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchWebAddenda - ns/op",
            "value": 3293,
            "unit": "ns/op",
            "extra": "350802 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchWebAddenda - B/op",
            "value": 1440,
            "unit": "B/op",
            "extra": "350802 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchWebAddenda - allocs/op",
            "value": 20,
            "unit": "allocs/op",
            "extra": "350802 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchWebIndividualNameRequired",
            "value": 1900,
            "unit": "ns/op\t    1104 B/op\t      13 allocs/op",
            "extra": "609036 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchWebIndividualNameRequired - ns/op",
            "value": 1900,
            "unit": "ns/op",
            "extra": "609036 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchWebIndividualNameRequired - B/op",
            "value": 1104,
            "unit": "B/op",
            "extra": "609036 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchWebIndividualNameRequired - allocs/op",
            "value": 13,
            "unit": "allocs/op",
            "extra": "609036 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchWEBAddendaTypeCode",
            "value": 2170,
            "unit": "ns/op\t    1216 B/op\t      16 allocs/op",
            "extra": "530998 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchWEBAddendaTypeCode - ns/op",
            "value": 2170,
            "unit": "ns/op",
            "extra": "530998 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchWEBAddendaTypeCode - B/op",
            "value": 1216,
            "unit": "B/op",
            "extra": "530998 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchWEBAddendaTypeCode - allocs/op",
            "value": 16,
            "unit": "allocs/op",
            "extra": "530998 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchWebSEC",
            "value": 2238,
            "unit": "ns/op\t    1048 B/op\t      13 allocs/op",
            "extra": "504409 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchWebSEC - ns/op",
            "value": 2238,
            "unit": "ns/op",
            "extra": "504409 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchWebSEC - B/op",
            "value": 1048,
            "unit": "B/op",
            "extra": "504409 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchWebSEC - allocs/op",
            "value": 13,
            "unit": "allocs/op",
            "extra": "504409 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchWebPaymentType",
            "value": 2196,
            "unit": "ns/op\t     968 B/op\t      12 allocs/op",
            "extra": "510570 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchWebPaymentType - ns/op",
            "value": 2196,
            "unit": "ns/op",
            "extra": "510570 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchWebPaymentType - B/op",
            "value": 968,
            "unit": "B/op",
            "extra": "510570 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchWebPaymentType - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "510570 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchWebCreate",
            "value": 1756,
            "unit": "ns/op\t    1024 B/op\t      12 allocs/op",
            "extra": "632079 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchWebCreate - ns/op",
            "value": 1756,
            "unit": "ns/op",
            "extra": "632079 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchWebCreate - B/op",
            "value": 1024,
            "unit": "B/op",
            "extra": "632079 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchWebCreate - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "632079 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchWebPaymentTypeR",
            "value": 2187,
            "unit": "ns/op\t     968 B/op\t      12 allocs/op",
            "extra": "513644 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchWebPaymentTypeR - ns/op",
            "value": 2187,
            "unit": "ns/op",
            "extra": "513644 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchWebPaymentTypeR - B/op",
            "value": 968,
            "unit": "B/op",
            "extra": "513644 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchWebPaymentTypeR - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "513644 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchXCKHeader",
            "value": 172.3,
            "unit": "ns/op\t     448 B/op\t       3 allocs/op",
            "extra": "7007635 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchXCKHeader - ns/op",
            "value": 172.3,
            "unit": "ns/op",
            "extra": "7007635 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchXCKHeader - B/op",
            "value": 448,
            "unit": "B/op",
            "extra": "7007635 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchXCKHeader - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "7007635 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchXCKCreate",
            "value": 2155,
            "unit": "ns/op\t    1032 B/op\t      11 allocs/op",
            "extra": "526374 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchXCKCreate - ns/op",
            "value": 2155,
            "unit": "ns/op",
            "extra": "526374 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchXCKCreate - B/op",
            "value": 1032,
            "unit": "B/op",
            "extra": "526374 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchXCKCreate - allocs/op",
            "value": 11,
            "unit": "allocs/op",
            "extra": "526374 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchXCKStandardEntryClassCode",
            "value": 2184,
            "unit": "ns/op\t    1112 B/op\t      12 allocs/op",
            "extra": "519334 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchXCKStandardEntryClassCode - ns/op",
            "value": 2184,
            "unit": "ns/op",
            "extra": "519334 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchXCKStandardEntryClassCode - B/op",
            "value": 1112,
            "unit": "B/op",
            "extra": "519334 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchXCKStandardEntryClassCode - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "519334 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchXCKServiceClassCodeEquality",
            "value": 2299,
            "unit": "ns/op\t    1128 B/op\t      14 allocs/op",
            "extra": "501945 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchXCKServiceClassCodeEquality - ns/op",
            "value": 2299,
            "unit": "ns/op",
            "extra": "501945 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchXCKServiceClassCodeEquality - B/op",
            "value": 1128,
            "unit": "B/op",
            "extra": "501945 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchXCKServiceClassCodeEquality - allocs/op",
            "value": 14,
            "unit": "allocs/op",
            "extra": "501945 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchXCKMixedCreditsAndDebits",
            "value": 2322,
            "unit": "ns/op\t    1128 B/op\t      14 allocs/op",
            "extra": "491365 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchXCKMixedCreditsAndDebits - ns/op",
            "value": 2322,
            "unit": "ns/op",
            "extra": "491365 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchXCKMixedCreditsAndDebits - B/op",
            "value": 1128,
            "unit": "B/op",
            "extra": "491365 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchXCKMixedCreditsAndDebits - allocs/op",
            "value": 14,
            "unit": "allocs/op",
            "extra": "491365 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchXCKCreditsOnly",
            "value": 2337,
            "unit": "ns/op\t    1128 B/op\t      14 allocs/op",
            "extra": "500089 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchXCKCreditsOnly - ns/op",
            "value": 2337,
            "unit": "ns/op",
            "extra": "500089 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchXCKCreditsOnly - B/op",
            "value": 1128,
            "unit": "B/op",
            "extra": "500089 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchXCKCreditsOnly - allocs/op",
            "value": 14,
            "unit": "allocs/op",
            "extra": "500089 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchXCKAutomatedAccountingAdvices",
            "value": 2329,
            "unit": "ns/op\t    1136 B/op\t      15 allocs/op",
            "extra": "492856 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchXCKAutomatedAccountingAdvices - ns/op",
            "value": 2329,
            "unit": "ns/op",
            "extra": "492856 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchXCKAutomatedAccountingAdvices - B/op",
            "value": 1136,
            "unit": "B/op",
            "extra": "492856 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchXCKAutomatedAccountingAdvices - allocs/op",
            "value": 15,
            "unit": "allocs/op",
            "extra": "492856 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchXCKCheckSerialNumber",
            "value": 1894,
            "unit": "ns/op\t     904 B/op\t      10 allocs/op",
            "extra": "604615 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchXCKCheckSerialNumber - ns/op",
            "value": 1894,
            "unit": "ns/op",
            "extra": "604615 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchXCKCheckSerialNumber - B/op",
            "value": 904,
            "unit": "B/op",
            "extra": "604615 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchXCKCheckSerialNumber - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "604615 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchXCKTransactionCode",
            "value": 1256,
            "unit": "ns/op\t     984 B/op\t      11 allocs/op",
            "extra": "857887 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchXCKTransactionCode - ns/op",
            "value": 1256,
            "unit": "ns/op",
            "extra": "857887 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchXCKTransactionCode - B/op",
            "value": 984,
            "unit": "B/op",
            "extra": "857887 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchXCKTransactionCode - allocs/op",
            "value": 11,
            "unit": "allocs/op",
            "extra": "857887 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchXCKAddendaCount",
            "value": 2576,
            "unit": "ns/op\t    1224 B/op\t      16 allocs/op",
            "extra": "442334 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchXCKAddendaCount - ns/op",
            "value": 2576,
            "unit": "ns/op",
            "extra": "442334 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchXCKAddendaCount - B/op",
            "value": 1224,
            "unit": "B/op",
            "extra": "442334 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchXCKAddendaCount - allocs/op",
            "value": 16,
            "unit": "allocs/op",
            "extra": "442334 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchXCKInvalidBuild",
            "value": 1642,
            "unit": "ns/op\t     968 B/op\t      11 allocs/op",
            "extra": "664093 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchXCKInvalidBuild - ns/op",
            "value": 1642,
            "unit": "ns/op",
            "extra": "664093 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchXCKInvalidBuild - B/op",
            "value": 968,
            "unit": "B/op",
            "extra": "664093 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchXCKInvalidBuild - allocs/op",
            "value": 11,
            "unit": "allocs/op",
            "extra": "664093 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchNumberMismatch",
            "value": 1654,
            "unit": "ns/op\t     960 B/op\t      11 allocs/op",
            "extra": "676422 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchNumberMismatch - ns/op",
            "value": 1654,
            "unit": "ns/op",
            "extra": "676422 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchNumberMismatch - B/op",
            "value": 960,
            "unit": "B/op",
            "extra": "676422 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchNumberMismatch - allocs/op",
            "value": 11,
            "unit": "allocs/op",
            "extra": "676422 times\n4 procs"
          },
          {
            "name": "BenchmarkCreditBatchIsBatchAmount",
            "value": 2597,
            "unit": "ns/op\t    1624 B/op\t      18 allocs/op",
            "extra": "443910 times\n4 procs"
          },
          {
            "name": "BenchmarkCreditBatchIsBatchAmount - ns/op",
            "value": 2597,
            "unit": "ns/op",
            "extra": "443910 times\n4 procs"
          },
          {
            "name": "BenchmarkCreditBatchIsBatchAmount - B/op",
            "value": 1624,
            "unit": "B/op",
            "extra": "443910 times\n4 procs"
          },
          {
            "name": "BenchmarkCreditBatchIsBatchAmount - allocs/op",
            "value": 18,
            "unit": "allocs/op",
            "extra": "443910 times\n4 procs"
          },
          {
            "name": "BenchmarkSavingsBatchIsBatchAmount",
            "value": 2596,
            "unit": "ns/op\t    1624 B/op\t      18 allocs/op",
            "extra": "445760 times\n4 procs"
          },
          {
            "name": "BenchmarkSavingsBatchIsBatchAmount - ns/op",
            "value": 2596,
            "unit": "ns/op",
            "extra": "445760 times\n4 procs"
          },
          {
            "name": "BenchmarkSavingsBatchIsBatchAmount - B/op",
            "value": 1624,
            "unit": "B/op",
            "extra": "445760 times\n4 procs"
          },
          {
            "name": "BenchmarkSavingsBatchIsBatchAmount - allocs/op",
            "value": 18,
            "unit": "allocs/op",
            "extra": "445760 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchIsEntryHash",
            "value": 1786,
            "unit": "ns/op\t    1000 B/op\t      12 allocs/op",
            "extra": "627283 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchIsEntryHash - ns/op",
            "value": 1786,
            "unit": "ns/op",
            "extra": "627283 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchIsEntryHash - B/op",
            "value": 1000,
            "unit": "B/op",
            "extra": "627283 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchIsEntryHash - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "627283 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchDNEMismatch",
            "value": 1939,
            "unit": "ns/op\t    1320 B/op\t      14 allocs/op",
            "extra": "592519 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchDNEMismatch - ns/op",
            "value": 1939,
            "unit": "ns/op",
            "extra": "592519 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchDNEMismatch - B/op",
            "value": 1320,
            "unit": "B/op",
            "extra": "592519 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchDNEMismatch - allocs/op",
            "value": 14,
            "unit": "allocs/op",
            "extra": "592519 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTraceNumberNotODFI",
            "value": 1973,
            "unit": "ns/op\t    1080 B/op\t      17 allocs/op",
            "extra": "564004 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTraceNumberNotODFI - ns/op",
            "value": 1973,
            "unit": "ns/op",
            "extra": "564004 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTraceNumberNotODFI - B/op",
            "value": 1080,
            "unit": "B/op",
            "extra": "564004 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTraceNumberNotODFI - allocs/op",
            "value": 17,
            "unit": "allocs/op",
            "extra": "564004 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchEntryCountEquality",
            "value": 3486,
            "unit": "ns/op\t    1728 B/op\t      20 allocs/op",
            "extra": "334184 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchEntryCountEquality - ns/op",
            "value": 3486,
            "unit": "ns/op",
            "extra": "334184 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchEntryCountEquality - B/op",
            "value": 1728,
            "unit": "B/op",
            "extra": "334184 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchEntryCountEquality - allocs/op",
            "value": 20,
            "unit": "allocs/op",
            "extra": "334184 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchAddendaIndicator",
            "value": 1487,
            "unit": "ns/op\t     904 B/op\t      10 allocs/op",
            "extra": "761313 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchAddendaIndicator - ns/op",
            "value": 1487,
            "unit": "ns/op",
            "extra": "761313 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchAddendaIndicator - B/op",
            "value": 904,
            "unit": "B/op",
            "extra": "761313 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchAddendaIndicator - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "761313 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchIsAddendaSeqAscending",
            "value": 2508,
            "unit": "ns/op\t    1344 B/op\t      17 allocs/op",
            "extra": "454663 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchIsAddendaSeqAscending - ns/op",
            "value": 2508,
            "unit": "ns/op",
            "extra": "454663 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchIsAddendaSeqAscending - B/op",
            "value": 1344,
            "unit": "B/op",
            "extra": "454663 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchIsAddendaSeqAscending - allocs/op",
            "value": 17,
            "unit": "allocs/op",
            "extra": "454663 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchIsSequenceAscending",
            "value": 2272,
            "unit": "ns/op\t    1384 B/op\t      17 allocs/op",
            "extra": "502094 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchIsSequenceAscending - ns/op",
            "value": 2272,
            "unit": "ns/op",
            "extra": "502094 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchIsSequenceAscending - B/op",
            "value": 1384,
            "unit": "B/op",
            "extra": "502094 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchIsSequenceAscending - allocs/op",
            "value": 17,
            "unit": "allocs/op",
            "extra": "502094 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchAddendaTraceNumber",
            "value": 2335,
            "unit": "ns/op\t    1248 B/op\t      15 allocs/op",
            "extra": "484137 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchAddendaTraceNumber - ns/op",
            "value": 2335,
            "unit": "ns/op",
            "extra": "484137 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchAddendaTraceNumber - B/op",
            "value": 1248,
            "unit": "B/op",
            "extra": "484137 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchAddendaTraceNumber - allocs/op",
            "value": 15,
            "unit": "allocs/op",
            "extra": "484137 times\n4 procs"
          },
          {
            "name": "BenchmarkNewBatchDefault",
            "value": 585.6,
            "unit": "ns/op\t     360 B/op\t       7 allocs/op",
            "extra": "2044814 times\n4 procs"
          },
          {
            "name": "BenchmarkNewBatchDefault - ns/op",
            "value": 585.6,
            "unit": "ns/op",
            "extra": "2044814 times\n4 procs"
          },
          {
            "name": "BenchmarkNewBatchDefault - B/op",
            "value": 360,
            "unit": "B/op",
            "extra": "2044814 times\n4 procs"
          },
          {
            "name": "BenchmarkNewBatchDefault - allocs/op",
            "value": 7,
            "unit": "allocs/op",
            "extra": "2044814 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCategory",
            "value": 1531,
            "unit": "ns/op\t    1304 B/op\t      13 allocs/op",
            "extra": "719948 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCategory - ns/op",
            "value": 1531,
            "unit": "ns/op",
            "extra": "719948 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCategory - B/op",
            "value": 1304,
            "unit": "B/op",
            "extra": "719948 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCategory - allocs/op",
            "value": 13,
            "unit": "allocs/op",
            "extra": "719948 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCategoryForwardReturn",
            "value": 2815,
            "unit": "ns/op\t    1640 B/op\t      22 allocs/op",
            "extra": "408544 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCategoryForwardReturn - ns/op",
            "value": 2815,
            "unit": "ns/op",
            "extra": "408544 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCategoryForwardReturn - B/op",
            "value": 1640,
            "unit": "B/op",
            "extra": "408544 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchCategoryForwardReturn - allocs/op",
            "value": 22,
            "unit": "allocs/op",
            "extra": "408544 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTraceNumberExists",
            "value": 1527,
            "unit": "ns/op\t    1160 B/op\t      12 allocs/op",
            "extra": "743472 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTraceNumberExists - ns/op",
            "value": 1527,
            "unit": "ns/op",
            "extra": "743472 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTraceNumberExists - B/op",
            "value": 1160,
            "unit": "B/op",
            "extra": "743472 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchTraceNumberExists - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "743472 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchFieldInclusion",
            "value": 1081,
            "unit": "ns/op\t     896 B/op\t      10 allocs/op",
            "extra": "996356 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchFieldInclusion - ns/op",
            "value": 1081,
            "unit": "ns/op",
            "extra": "996356 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchFieldInclusion - B/op",
            "value": 896,
            "unit": "B/op",
            "extra": "996356 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchFieldInclusion - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "996356 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchInvalidTraceNumberODFI",
            "value": 750.6,
            "unit": "ns/op\t     768 B/op\t      10 allocs/op",
            "extra": "1596129 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchInvalidTraceNumberODFI - ns/op",
            "value": 750.6,
            "unit": "ns/op",
            "extra": "1596129 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchInvalidTraceNumberODFI - B/op",
            "value": 768,
            "unit": "B/op",
            "extra": "1596129 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchInvalidTraceNumberODFI - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "1596129 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchNoEntry",
            "value": 245.8,
            "unit": "ns/op\t     352 B/op\t       3 allocs/op",
            "extra": "4876195 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchNoEntry - ns/op",
            "value": 245.8,
            "unit": "ns/op",
            "extra": "4876195 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchNoEntry - B/op",
            "value": 352,
            "unit": "B/op",
            "extra": "4876195 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchNoEntry - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "4876195 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchControl",
            "value": 1639,
            "unit": "ns/op\t     976 B/op\t      12 allocs/op",
            "extra": "695906 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchControl - ns/op",
            "value": 1639,
            "unit": "ns/op",
            "extra": "695906 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchControl - B/op",
            "value": 976,
            "unit": "B/op",
            "extra": "695906 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchControl - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "695906 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatch",
            "value": 247.2,
            "unit": "ns/op\t     200 B/op\t       2 allocs/op",
            "extra": "4866328 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatch - ns/op",
            "value": 247.2,
            "unit": "ns/op",
            "extra": "4866328 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatch - B/op",
            "value": 200,
            "unit": "B/op",
            "extra": "4866328 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatch - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "4866328 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldShort",
            "value": 1188,
            "unit": "ns/op\t      96 B/op\t       4 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldShort - ns/op",
            "value": 1188,
            "unit": "ns/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldShort - B/op",
            "value": 96,
            "unit": "B/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldShort - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong",
            "value": 13.31,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "88533391 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong - ns/op",
            "value": 13.31,
            "unit": "ns/op",
            "extra": "88533391 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "88533391 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "88533391 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort",
            "value": 62.24,
            "unit": "ns/op\t      16 B/op\t       2 allocs/op",
            "extra": "19075627 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort - ns/op",
            "value": 62.24,
            "unit": "ns/op",
            "extra": "19075627 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "19075627 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "19075627 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong",
            "value": 30.41,
            "unit": "ns/op\t       8 B/op\t       1 allocs/op",
            "extra": "39368860 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong - ns/op",
            "value": 30.41,
            "unit": "ns/op",
            "extra": "39368860 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong - B/op",
            "value": 8,
            "unit": "B/op",
            "extra": "39368860 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "39368860 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField",
            "value": 13.38,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "89248480 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField - ns/op",
            "value": 13.38,
            "unit": "ns/op",
            "extra": "89248480 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "89248480 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "89248480 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField",
            "value": 5.916,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "202249029 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField - ns/op",
            "value": 5.916,
            "unit": "ns/op",
            "extra": "202249029 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "202249029 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "202249029 times\n4 procs"
          },
          {
            "name": "BenchmarkRTNFieldShort",
            "value": 42.81,
            "unit": "ns/op\t       8 B/op\t       1 allocs/op",
            "extra": "27180520 times\n4 procs"
          },
          {
            "name": "BenchmarkRTNFieldShort - ns/op",
            "value": 42.81,
            "unit": "ns/op",
            "extra": "27180520 times\n4 procs"
          },
          {
            "name": "BenchmarkRTNFieldShort - B/op",
            "value": 8,
            "unit": "B/op",
            "extra": "27180520 times\n4 procs"
          },
          {
            "name": "BenchmarkRTNFieldShort - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "27180520 times\n4 procs"
          },
          {
            "name": "BenchmarkRTNFieldLong",
            "value": 12.52,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "94690303 times\n4 procs"
          },
          {
            "name": "BenchmarkRTNFieldLong - ns/op",
            "value": 12.52,
            "unit": "ns/op",
            "extra": "94690303 times\n4 procs"
          },
          {
            "name": "BenchmarkRTNFieldLong - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "94690303 times\n4 procs"
          },
          {
            "name": "BenchmarkRTNFieldLong - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "94690303 times\n4 procs"
          },
          {
            "name": "BenchmarkRTNFieldExact",
            "value": 23.89,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "50949036 times\n4 procs"
          },
          {
            "name": "BenchmarkRTNFieldExact - ns/op",
            "value": 23.89,
            "unit": "ns/op",
            "extra": "50949036 times\n4 procs"
          },
          {
            "name": "BenchmarkRTNFieldExact - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "50949036 times\n4 procs"
          },
          {
            "name": "BenchmarkRTNFieldExact - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "50949036 times\n4 procs"
          },
          {
            "name": "BenchmarkMockEntryDetail",
            "value": 342.3,
            "unit": "ns/op\t     280 B/op\t       3 allocs/op",
            "extra": "3513469 times\n4 procs"
          },
          {
            "name": "BenchmarkMockEntryDetail - ns/op",
            "value": 342.3,
            "unit": "ns/op",
            "extra": "3513469 times\n4 procs"
          },
          {
            "name": "BenchmarkMockEntryDetail - B/op",
            "value": 280,
            "unit": "B/op",
            "extra": "3513469 times\n4 procs"
          },
          {
            "name": "BenchmarkMockEntryDetail - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "3513469 times\n4 procs"
          },
          {
            "name": "BenchmarkParseEntryDetail",
            "value": 6250,
            "unit": "ns/op\t   15748 B/op\t      30 allocs/op",
            "extra": "190770 times\n4 procs"
          },
          {
            "name": "BenchmarkParseEntryDetail - ns/op",
            "value": 6250,
            "unit": "ns/op",
            "extra": "190770 times\n4 procs"
          },
          {
            "name": "BenchmarkParseEntryDetail - B/op",
            "value": 15748,
            "unit": "B/op",
            "extra": "190770 times\n4 procs"
          },
          {
            "name": "BenchmarkParseEntryDetail - allocs/op",
            "value": 30,
            "unit": "allocs/op",
            "extra": "190770 times\n4 procs"
          },
          {
            "name": "BenchmarkEDString",
            "value": 6620,
            "unit": "ns/op\t   15856 B/op\t      32 allocs/op",
            "extra": "184627 times\n4 procs"
          },
          {
            "name": "BenchmarkEDString - ns/op",
            "value": 6620,
            "unit": "ns/op",
            "extra": "184627 times\n4 procs"
          },
          {
            "name": "BenchmarkEDString - B/op",
            "value": 15856,
            "unit": "B/op",
            "extra": "184627 times\n4 procs"
          },
          {
            "name": "BenchmarkEDString - allocs/op",
            "value": 32,
            "unit": "allocs/op",
            "extra": "184627 times\n4 procs"
          },
          {
            "name": "BenchmarkValidateEDTransactionCode",
            "value": 318.9,
            "unit": "ns/op\t     360 B/op\t       5 allocs/op",
            "extra": "3791217 times\n4 procs"
          },
          {
            "name": "BenchmarkValidateEDTransactionCode - ns/op",
            "value": 318.9,
            "unit": "ns/op",
            "extra": "3791217 times\n4 procs"
          },
          {
            "name": "BenchmarkValidateEDTransactionCode - B/op",
            "value": 360,
            "unit": "B/op",
            "extra": "3791217 times\n4 procs"
          },
          {
            "name": "BenchmarkValidateEDTransactionCode - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "3791217 times\n4 procs"
          },
          {
            "name": "BenchmarkEDFieldInclusion",
            "value": 336.8,
            "unit": "ns/op\t     280 B/op\t       3 allocs/op",
            "extra": "3562562 times\n4 procs"
          },
          {
            "name": "BenchmarkEDFieldInclusion - ns/op",
            "value": 336.8,
            "unit": "ns/op",
            "extra": "3562562 times\n4 procs"
          },
          {
            "name": "BenchmarkEDFieldInclusion - B/op",
            "value": 280,
            "unit": "B/op",
            "extra": "3562562 times\n4 procs"
          },
          {
            "name": "BenchmarkEDFieldInclusion - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "3562562 times\n4 procs"
          },
          {
            "name": "BenchmarkEDdfiAccountNumberAlphaNumeric",
            "value": 593.4,
            "unit": "ns/op\t     440 B/op\t       7 allocs/op",
            "extra": "2025885 times\n4 procs"
          },
          {
            "name": "BenchmarkEDdfiAccountNumberAlphaNumeric - ns/op",
            "value": 593.4,
            "unit": "ns/op",
            "extra": "2025885 times\n4 procs"
          },
          {
            "name": "BenchmarkEDdfiAccountNumberAlphaNumeric - B/op",
            "value": 440,
            "unit": "B/op",
            "extra": "2025885 times\n4 procs"
          },
          {
            "name": "BenchmarkEDdfiAccountNumberAlphaNumeric - allocs/op",
            "value": 7,
            "unit": "allocs/op",
            "extra": "2025885 times\n4 procs"
          },
          {
            "name": "BenchmarkEDIdentificationNumberAlphaNumeric",
            "value": 631.6,
            "unit": "ns/op\t     440 B/op\t       7 allocs/op",
            "extra": "1910606 times\n4 procs"
          },
          {
            "name": "BenchmarkEDIdentificationNumberAlphaNumeric - ns/op",
            "value": 631.6,
            "unit": "ns/op",
            "extra": "1910606 times\n4 procs"
          },
          {
            "name": "BenchmarkEDIdentificationNumberAlphaNumeric - B/op",
            "value": 440,
            "unit": "B/op",
            "extra": "1910606 times\n4 procs"
          },
          {
            "name": "BenchmarkEDIdentificationNumberAlphaNumeric - allocs/op",
            "value": 7,
            "unit": "allocs/op",
            "extra": "1910606 times\n4 procs"
          },
          {
            "name": "BenchmarkEDIndividualNameAlphaNumeric",
            "value": 638.2,
            "unit": "ns/op\t     440 B/op\t       7 allocs/op",
            "extra": "1878580 times\n4 procs"
          },
          {
            "name": "BenchmarkEDIndividualNameAlphaNumeric - ns/op",
            "value": 638.2,
            "unit": "ns/op",
            "extra": "1878580 times\n4 procs"
          },
          {
            "name": "BenchmarkEDIndividualNameAlphaNumeric - B/op",
            "value": 440,
            "unit": "B/op",
            "extra": "1878580 times\n4 procs"
          },
          {
            "name": "BenchmarkEDIndividualNameAlphaNumeric - allocs/op",
            "value": 7,
            "unit": "allocs/op",
            "extra": "1878580 times\n4 procs"
          },
          {
            "name": "BenchmarkEDDiscretionaryDataAlphaNumeric",
            "value": 639.7,
            "unit": "ns/op\t     440 B/op\t       7 allocs/op",
            "extra": "1884846 times\n4 procs"
          },
          {
            "name": "BenchmarkEDDiscretionaryDataAlphaNumeric - ns/op",
            "value": 639.7,
            "unit": "ns/op",
            "extra": "1884846 times\n4 procs"
          },
          {
            "name": "BenchmarkEDDiscretionaryDataAlphaNumeric - B/op",
            "value": 440,
            "unit": "B/op",
            "extra": "1884846 times\n4 procs"
          },
          {
            "name": "BenchmarkEDDiscretionaryDataAlphaNumeric - allocs/op",
            "value": 7,
            "unit": "allocs/op",
            "extra": "1884846 times\n4 procs"
          },
          {
            "name": "BenchmarkEDisCheckDigit",
            "value": 747.6,
            "unit": "ns/op\t     480 B/op\t       8 allocs/op",
            "extra": "1597192 times\n4 procs"
          },
          {
            "name": "BenchmarkEDisCheckDigit - ns/op",
            "value": 747.6,
            "unit": "ns/op",
            "extra": "1597192 times\n4 procs"
          },
          {
            "name": "BenchmarkEDisCheckDigit - B/op",
            "value": 480,
            "unit": "B/op",
            "extra": "1597192 times\n4 procs"
          },
          {
            "name": "BenchmarkEDisCheckDigit - allocs/op",
            "value": 8,
            "unit": "allocs/op",
            "extra": "1597192 times\n4 procs"
          },
          {
            "name": "BenchmarkEDSetRDFI",
            "value": 35.74,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "33972601 times\n4 procs"
          },
          {
            "name": "BenchmarkEDSetRDFI - ns/op",
            "value": 35.74,
            "unit": "ns/op",
            "extra": "33972601 times\n4 procs"
          },
          {
            "name": "BenchmarkEDSetRDFI - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "33972601 times\n4 procs"
          },
          {
            "name": "BenchmarkEDSetRDFI - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "33972601 times\n4 procs"
          },
          {
            "name": "BenchmarkEDFieldInclusionTransactionCode",
            "value": 311.1,
            "unit": "ns/op\t     360 B/op\t       5 allocs/op",
            "extra": "3852862 times\n4 procs"
          },
          {
            "name": "BenchmarkEDFieldInclusionTransactionCode - ns/op",
            "value": 311.1,
            "unit": "ns/op",
            "extra": "3852862 times\n4 procs"
          },
          {
            "name": "BenchmarkEDFieldInclusionTransactionCode - B/op",
            "value": 360,
            "unit": "B/op",
            "extra": "3852862 times\n4 procs"
          },
          {
            "name": "BenchmarkEDFieldInclusionTransactionCode - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "3852862 times\n4 procs"
          },
          {
            "name": "BenchmarkEDFieldInclusionRDFIIdentification",
            "value": 341.3,
            "unit": "ns/op\t     360 B/op\t       5 allocs/op",
            "extra": "3542397 times\n4 procs"
          },
          {
            "name": "BenchmarkEDFieldInclusionRDFIIdentification - ns/op",
            "value": 341.3,
            "unit": "ns/op",
            "extra": "3542397 times\n4 procs"
          },
          {
            "name": "BenchmarkEDFieldInclusionRDFIIdentification - B/op",
            "value": 360,
            "unit": "B/op",
            "extra": "3542397 times\n4 procs"
          },
          {
            "name": "BenchmarkEDFieldInclusionRDFIIdentification - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "3542397 times\n4 procs"
          },
          {
            "name": "BenchmarkEDFieldInclusionDFIAccountNumber",
            "value": 286.4,
            "unit": "ns/op\t     344 B/op\t       4 allocs/op",
            "extra": "4231279 times\n4 procs"
          },
          {
            "name": "BenchmarkEDFieldInclusionDFIAccountNumber - ns/op",
            "value": 286.4,
            "unit": "ns/op",
            "extra": "4231279 times\n4 procs"
          },
          {
            "name": "BenchmarkEDFieldInclusionDFIAccountNumber - B/op",
            "value": 344,
            "unit": "B/op",
            "extra": "4231279 times\n4 procs"
          },
          {
            "name": "BenchmarkEDFieldInclusionDFIAccountNumber - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "4231279 times\n4 procs"
          },
          {
            "name": "BenchmarkEDFieldInclusionIndividualName",
            "value": 280.2,
            "unit": "ns/op\t     344 B/op\t       4 allocs/op",
            "extra": "4284913 times\n4 procs"
          },
          {
            "name": "BenchmarkEDFieldInclusionIndividualName - ns/op",
            "value": 280.2,
            "unit": "ns/op",
            "extra": "4284913 times\n4 procs"
          },
          {
            "name": "BenchmarkEDFieldInclusionIndividualName - B/op",
            "value": 344,
            "unit": "B/op",
            "extra": "4284913 times\n4 procs"
          },
          {
            "name": "BenchmarkEDFieldInclusionIndividualName - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "4284913 times\n4 procs"
          },
          {
            "name": "BenchmarkEDFieldInclusionTraceNumber",
            "value": 344.7,
            "unit": "ns/op\t     280 B/op\t       3 allocs/op",
            "extra": "3473344 times\n4 procs"
          },
          {
            "name": "BenchmarkEDFieldInclusionTraceNumber - ns/op",
            "value": 344.7,
            "unit": "ns/op",
            "extra": "3473344 times\n4 procs"
          },
          {
            "name": "BenchmarkEDFieldInclusionTraceNumber - B/op",
            "value": 280,
            "unit": "B/op",
            "extra": "3473344 times\n4 procs"
          },
          {
            "name": "BenchmarkEDFieldInclusionTraceNumber - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "3473344 times\n4 procs"
          },
          {
            "name": "BenchmarkEDAddAddenda99",
            "value": 271.2,
            "unit": "ns/op\t     424 B/op\t       4 allocs/op",
            "extra": "4415230 times\n4 procs"
          },
          {
            "name": "BenchmarkEDAddAddenda99 - ns/op",
            "value": 271.2,
            "unit": "ns/op",
            "extra": "4415230 times\n4 procs"
          },
          {
            "name": "BenchmarkEDAddAddenda99 - B/op",
            "value": 424,
            "unit": "B/op",
            "extra": "4415230 times\n4 procs"
          },
          {
            "name": "BenchmarkEDAddAddenda99 - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "4415230 times\n4 procs"
          },
          {
            "name": "BenchmarkEDAddAddenda99Twice",
            "value": 330.4,
            "unit": "ns/op\t     568 B/op\t       5 allocs/op",
            "extra": "3621966 times\n4 procs"
          },
          {
            "name": "BenchmarkEDAddAddenda99Twice - ns/op",
            "value": 330.4,
            "unit": "ns/op",
            "extra": "3621966 times\n4 procs"
          },
          {
            "name": "BenchmarkEDAddAddenda99Twice - B/op",
            "value": 568,
            "unit": "B/op",
            "extra": "3621966 times\n4 procs"
          },
          {
            "name": "BenchmarkEDAddAddenda99Twice - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "3621966 times\n4 procs"
          },
          {
            "name": "BenchmarkEDCreditOrDebit",
            "value": 917.7,
            "unit": "ns/op\t     726 B/op\t       4 allocs/op",
            "extra": "1309394 times\n4 procs"
          },
          {
            "name": "BenchmarkEDCreditOrDebit - ns/op",
            "value": 917.7,
            "unit": "ns/op",
            "extra": "1309394 times\n4 procs"
          },
          {
            "name": "BenchmarkEDCreditOrDebit - B/op",
            "value": 726,
            "unit": "B/op",
            "extra": "1309394 times\n4 procs"
          },
          {
            "name": "BenchmarkEDCreditOrDebit - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "1309394 times\n4 procs"
          },
          {
            "name": "BenchmarkValidateEDCheckDigit",
            "value": 480,
            "unit": "ns/op\t     416 B/op\t       7 allocs/op",
            "extra": "2489624 times\n4 procs"
          },
          {
            "name": "BenchmarkValidateEDCheckDigit - ns/op",
            "value": 480,
            "unit": "ns/op",
            "extra": "2489624 times\n4 procs"
          },
          {
            "name": "BenchmarkValidateEDCheckDigit - B/op",
            "value": 416,
            "unit": "B/op",
            "extra": "2489624 times\n4 procs"
          },
          {
            "name": "BenchmarkValidateEDCheckDigit - allocs/op",
            "value": 7,
            "unit": "allocs/op",
            "extra": "2489624 times\n4 procs"
          },
          {
            "name": "BenchmarkMockFileControl",
            "value": 12.8,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "93765309 times\n4 procs"
          },
          {
            "name": "BenchmarkMockFileControl - ns/op",
            "value": 12.8,
            "unit": "ns/op",
            "extra": "93765309 times\n4 procs"
          },
          {
            "name": "BenchmarkMockFileControl - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "93765309 times\n4 procs"
          },
          {
            "name": "BenchmarkMockFileControl - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "93765309 times\n4 procs"
          },
          {
            "name": "BenchmarkParseFileControl",
            "value": 5446,
            "unit": "ns/op\t   14835 B/op\t      23 allocs/op",
            "extra": "220837 times\n4 procs"
          },
          {
            "name": "BenchmarkParseFileControl - ns/op",
            "value": 5446,
            "unit": "ns/op",
            "extra": "220837 times\n4 procs"
          },
          {
            "name": "BenchmarkParseFileControl - B/op",
            "value": 14835,
            "unit": "B/op",
            "extra": "220837 times\n4 procs"
          },
          {
            "name": "BenchmarkParseFileControl - allocs/op",
            "value": 23,
            "unit": "allocs/op",
            "extra": "220837 times\n4 procs"
          },
          {
            "name": "BenchmarkFCString",
            "value": 5582,
            "unit": "ns/op\t   14931 B/op\t      24 allocs/op",
            "extra": "217821 times\n4 procs"
          },
          {
            "name": "BenchmarkFCString - ns/op",
            "value": 5582,
            "unit": "ns/op",
            "extra": "217821 times\n4 procs"
          },
          {
            "name": "BenchmarkFCString - B/op",
            "value": 14931,
            "unit": "B/op",
            "extra": "217821 times\n4 procs"
          },
          {
            "name": "BenchmarkFCString - allocs/op",
            "value": 24,
            "unit": "allocs/op",
            "extra": "217821 times\n4 procs"
          },
          {
            "name": "BenchmarkFCFieldInclusion",
            "value": 128,
            "unit": "ns/op\t      88 B/op\t       3 allocs/op",
            "extra": "9334570 times\n4 procs"
          },
          {
            "name": "BenchmarkFCFieldInclusion - ns/op",
            "value": 128,
            "unit": "ns/op",
            "extra": "9334570 times\n4 procs"
          },
          {
            "name": "BenchmarkFCFieldInclusion - B/op",
            "value": 88,
            "unit": "B/op",
            "extra": "9334570 times\n4 procs"
          },
          {
            "name": "BenchmarkFCFieldInclusion - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "9334570 times\n4 procs"
          },
          {
            "name": "BenchmarkFCFieldInclusionBlockCount",
            "value": 129,
            "unit": "ns/op\t      88 B/op\t       3 allocs/op",
            "extra": "9242815 times\n4 procs"
          },
          {
            "name": "BenchmarkFCFieldInclusionBlockCount - ns/op",
            "value": 129,
            "unit": "ns/op",
            "extra": "9242815 times\n4 procs"
          },
          {
            "name": "BenchmarkFCFieldInclusionBlockCount - B/op",
            "value": 88,
            "unit": "B/op",
            "extra": "9242815 times\n4 procs"
          },
          {
            "name": "BenchmarkFCFieldInclusionBlockCount - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "9242815 times\n4 procs"
          },
          {
            "name": "BenchmarkFCFieldInclusionEntryAddendaCount",
            "value": 128.2,
            "unit": "ns/op\t      88 B/op\t       3 allocs/op",
            "extra": "9285014 times\n4 procs"
          },
          {
            "name": "BenchmarkFCFieldInclusionEntryAddendaCount - ns/op",
            "value": 128.2,
            "unit": "ns/op",
            "extra": "9285014 times\n4 procs"
          },
          {
            "name": "BenchmarkFCFieldInclusionEntryAddendaCount - B/op",
            "value": 88,
            "unit": "B/op",
            "extra": "9285014 times\n4 procs"
          },
          {
            "name": "BenchmarkFCFieldInclusionEntryAddendaCount - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "9285014 times\n4 procs"
          },
          {
            "name": "BenchmarkFCFieldInclusionEntryHash",
            "value": 127.9,
            "unit": "ns/op\t      88 B/op\t       3 allocs/op",
            "extra": "9359946 times\n4 procs"
          },
          {
            "name": "BenchmarkFCFieldInclusionEntryHash - ns/op",
            "value": 127.9,
            "unit": "ns/op",
            "extra": "9359946 times\n4 procs"
          },
          {
            "name": "BenchmarkFCFieldInclusionEntryHash - B/op",
            "value": 88,
            "unit": "B/op",
            "extra": "9359946 times\n4 procs"
          },
          {
            "name": "BenchmarkFCFieldInclusionEntryHash - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "9359946 times\n4 procs"
          },
          {
            "name": "BenchmarkMockFileHeader",
            "value": 297.9,
            "unit": "ns/op\t       8 B/op\t       1 allocs/op",
            "extra": "4016888 times\n4 procs"
          },
          {
            "name": "BenchmarkMockFileHeader - ns/op",
            "value": 297.9,
            "unit": "ns/op",
            "extra": "4016888 times\n4 procs"
          },
          {
            "name": "BenchmarkMockFileHeader - B/op",
            "value": 8,
            "unit": "B/op",
            "extra": "4016888 times\n4 procs"
          },
          {
            "name": "BenchmarkMockFileHeader - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "4016888 times\n4 procs"
          },
          {
            "name": "BenchmarkParseFileHeader",
            "value": 6190,
            "unit": "ns/op\t   15299 B/op\t      27 allocs/op",
            "extra": "192729 times\n4 procs"
          },
          {
            "name": "BenchmarkParseFileHeader - ns/op",
            "value": 6190,
            "unit": "ns/op",
            "extra": "192729 times\n4 procs"
          },
          {
            "name": "BenchmarkParseFileHeader - B/op",
            "value": 15299,
            "unit": "B/op",
            "extra": "192729 times\n4 procs"
          },
          {
            "name": "BenchmarkParseFileHeader - allocs/op",
            "value": 27,
            "unit": "allocs/op",
            "extra": "192729 times\n4 procs"
          },
          {
            "name": "BenchmarkFHString",
            "value": 5765,
            "unit": "ns/op\t   15219 B/op\t      23 allocs/op",
            "extra": "209732 times\n4 procs"
          },
          {
            "name": "BenchmarkFHString - ns/op",
            "value": 5765,
            "unit": "ns/op",
            "extra": "209732 times\n4 procs"
          },
          {
            "name": "BenchmarkFHString - B/op",
            "value": 15219,
            "unit": "B/op",
            "extra": "209732 times\n4 procs"
          },
          {
            "name": "BenchmarkFHString - allocs/op",
            "value": 23,
            "unit": "allocs/op",
            "extra": "209732 times\n4 procs"
          },
          {
            "name": "BenchmarkValidateIDModifier",
            "value": 576.9,
            "unit": "ns/op\t     152 B/op\t       5 allocs/op",
            "extra": "2077201 times\n4 procs"
          },
          {
            "name": "BenchmarkValidateIDModifier - ns/op",
            "value": 576.9,
            "unit": "ns/op",
            "extra": "2077201 times\n4 procs"
          },
          {
            "name": "BenchmarkValidateIDModifier - B/op",
            "value": 152,
            "unit": "B/op",
            "extra": "2077201 times\n4 procs"
          },
          {
            "name": "BenchmarkValidateIDModifier - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "2077201 times\n4 procs"
          },
          {
            "name": "BenchmarkValidateRecordSize",
            "value": 310.2,
            "unit": "ns/op\t      88 B/op\t       3 allocs/op",
            "extra": "3847062 times\n4 procs"
          },
          {
            "name": "BenchmarkValidateRecordSize - ns/op",
            "value": 310.2,
            "unit": "ns/op",
            "extra": "3847062 times\n4 procs"
          },
          {
            "name": "BenchmarkValidateRecordSize - B/op",
            "value": 88,
            "unit": "B/op",
            "extra": "3847062 times\n4 procs"
          },
          {
            "name": "BenchmarkValidateRecordSize - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "3847062 times\n4 procs"
          },
          {
            "name": "BenchmarkBlockingFactor",
            "value": 308.5,
            "unit": "ns/op\t      88 B/op\t       3 allocs/op",
            "extra": "3887740 times\n4 procs"
          },
          {
            "name": "BenchmarkBlockingFactor - ns/op",
            "value": 308.5,
            "unit": "ns/op",
            "extra": "3887740 times\n4 procs"
          },
          {
            "name": "BenchmarkBlockingFactor - B/op",
            "value": 88,
            "unit": "B/op",
            "extra": "3887740 times\n4 procs"
          },
          {
            "name": "BenchmarkBlockingFactor - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "3887740 times\n4 procs"
          },
          {
            "name": "BenchmarkFormatCode",
            "value": 309.6,
            "unit": "ns/op\t      88 B/op\t       3 allocs/op",
            "extra": "3881529 times\n4 procs"
          },
          {
            "name": "BenchmarkFormatCode - ns/op",
            "value": 309.6,
            "unit": "ns/op",
            "extra": "3881529 times\n4 procs"
          },
          {
            "name": "BenchmarkFormatCode - B/op",
            "value": 88,
            "unit": "B/op",
            "extra": "3881529 times\n4 procs"
          },
          {
            "name": "BenchmarkFormatCode - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "3881529 times\n4 procs"
          },
          {
            "name": "BenchmarkFHFieldInclusion",
            "value": 321.7,
            "unit": "ns/op\t      88 B/op\t       3 allocs/op",
            "extra": "3738847 times\n4 procs"
          },
          {
            "name": "BenchmarkFHFieldInclusion - ns/op",
            "value": 321.7,
            "unit": "ns/op",
            "extra": "3738847 times\n4 procs"
          },
          {
            "name": "BenchmarkFHFieldInclusion - B/op",
            "value": 88,
            "unit": "B/op",
            "extra": "3738847 times\n4 procs"
          },
          {
            "name": "BenchmarkFHFieldInclusion - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "3738847 times\n4 procs"
          },
          {
            "name": "BenchmarkUpperLengthFileID",
            "value": 927.4,
            "unit": "ns/op\t     288 B/op\t      10 allocs/op",
            "extra": "1298484 times\n4 procs"
          },
          {
            "name": "BenchmarkUpperLengthFileID - ns/op",
            "value": 927.4,
            "unit": "ns/op",
            "extra": "1298484 times\n4 procs"
          },
          {
            "name": "BenchmarkUpperLengthFileID - B/op",
            "value": 288,
            "unit": "B/op",
            "extra": "1298484 times\n4 procs"
          },
          {
            "name": "BenchmarkUpperLengthFileID - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "1298484 times\n4 procs"
          },
          {
            "name": "BenchmarkImmediateDestinationNameAlphaNumeric",
            "value": 600.4,
            "unit": "ns/op\t     168 B/op\t       5 allocs/op",
            "extra": "1995770 times\n4 procs"
          },
          {
            "name": "BenchmarkImmediateDestinationNameAlphaNumeric - ns/op",
            "value": 600.4,
            "unit": "ns/op",
            "extra": "1995770 times\n4 procs"
          },
          {
            "name": "BenchmarkImmediateDestinationNameAlphaNumeric - B/op",
            "value": 168,
            "unit": "B/op",
            "extra": "1995770 times\n4 procs"
          },
          {
            "name": "BenchmarkImmediateDestinationNameAlphaNumeric - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "1995770 times\n4 procs"
          },
          {
            "name": "BenchmarkImmediateOriginNameAlphaNumeric",
            "value": 719.5,
            "unit": "ns/op\t     168 B/op\t       5 allocs/op",
            "extra": "1676548 times\n4 procs"
          },
          {
            "name": "BenchmarkImmediateOriginNameAlphaNumeric - ns/op",
            "value": 719.5,
            "unit": "ns/op",
            "extra": "1676548 times\n4 procs"
          },
          {
            "name": "BenchmarkImmediateOriginNameAlphaNumeric - B/op",
            "value": 168,
            "unit": "B/op",
            "extra": "1676548 times\n4 procs"
          },
          {
            "name": "BenchmarkImmediateOriginNameAlphaNumeric - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "1676548 times\n4 procs"
          },
          {
            "name": "BenchmarkImmediateReferenceCodeAlphaNumeric",
            "value": 720.8,
            "unit": "ns/op\t     168 B/op\t       5 allocs/op",
            "extra": "1661072 times\n4 procs"
          },
          {
            "name": "BenchmarkImmediateReferenceCodeAlphaNumeric - ns/op",
            "value": 720.8,
            "unit": "ns/op",
            "extra": "1661072 times\n4 procs"
          },
          {
            "name": "BenchmarkImmediateReferenceCodeAlphaNumeric - B/op",
            "value": 168,
            "unit": "B/op",
            "extra": "1661072 times\n4 procs"
          },
          {
            "name": "BenchmarkImmediateReferenceCodeAlphaNumeric - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "1661072 times\n4 procs"
          },
          {
            "name": "BenchmarkFHFieldInclusionImmediateDestination",
            "value": 317.7,
            "unit": "ns/op\t      88 B/op\t       3 allocs/op",
            "extra": "3778566 times\n4 procs"
          },
          {
            "name": "BenchmarkFHFieldInclusionImmediateDestination - ns/op",
            "value": 317.7,
            "unit": "ns/op",
            "extra": "3778566 times\n4 procs"
          },
          {
            "name": "BenchmarkFHFieldInclusionImmediateDestination - B/op",
            "value": 88,
            "unit": "B/op",
            "extra": "3778566 times\n4 procs"
          },
          {
            "name": "BenchmarkFHFieldInclusionImmediateDestination - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "3778566 times\n4 procs"
          },
          {
            "name": "BenchmarkFHFieldInclusionFileIDModifier",
            "value": 266.5,
            "unit": "ns/op\t      72 B/op\t       2 allocs/op",
            "extra": "4501069 times\n4 procs"
          },
          {
            "name": "BenchmarkFHFieldInclusionFileIDModifier - ns/op",
            "value": 266.5,
            "unit": "ns/op",
            "extra": "4501069 times\n4 procs"
          },
          {
            "name": "BenchmarkFHFieldInclusionFileIDModifier - B/op",
            "value": 72,
            "unit": "B/op",
            "extra": "4501069 times\n4 procs"
          },
          {
            "name": "BenchmarkFHFieldInclusionFileIDModifier - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "4501069 times\n4 procs"
          },
          {
            "name": "BenchmarkFHFieldInclusionRecordSize",
            "value": 266.5,
            "unit": "ns/op\t      72 B/op\t       2 allocs/op",
            "extra": "4440192 times\n4 procs"
          },
          {
            "name": "BenchmarkFHFieldInclusionRecordSize - ns/op",
            "value": 266.5,
            "unit": "ns/op",
            "extra": "4440192 times\n4 procs"
          },
          {
            "name": "BenchmarkFHFieldInclusionRecordSize - B/op",
            "value": 72,
            "unit": "B/op",
            "extra": "4440192 times\n4 procs"
          },
          {
            "name": "BenchmarkFHFieldInclusionRecordSize - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "4440192 times\n4 procs"
          },
          {
            "name": "BenchmarkFHFieldInclusionBlockingFactor",
            "value": 268.5,
            "unit": "ns/op\t      72 B/op\t       2 allocs/op",
            "extra": "4421199 times\n4 procs"
          },
          {
            "name": "BenchmarkFHFieldInclusionBlockingFactor - ns/op",
            "value": 268.5,
            "unit": "ns/op",
            "extra": "4421199 times\n4 procs"
          },
          {
            "name": "BenchmarkFHFieldInclusionBlockingFactor - B/op",
            "value": 72,
            "unit": "B/op",
            "extra": "4421199 times\n4 procs"
          },
          {
            "name": "BenchmarkFHFieldInclusionBlockingFactor - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "4421199 times\n4 procs"
          },
          {
            "name": "BenchmarkFHFieldInclusionFormatCode",
            "value": 271.6,
            "unit": "ns/op\t      72 B/op\t       2 allocs/op",
            "extra": "4443573 times\n4 procs"
          },
          {
            "name": "BenchmarkFHFieldInclusionFormatCode - ns/op",
            "value": 271.6,
            "unit": "ns/op",
            "extra": "4443573 times\n4 procs"
          },
          {
            "name": "BenchmarkFHFieldInclusionFormatCode - B/op",
            "value": 72,
            "unit": "B/op",
            "extra": "4443573 times\n4 procs"
          },
          {
            "name": "BenchmarkFHFieldInclusionFormatCode - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "4443573 times\n4 procs"
          },
          {
            "name": "BenchmarkFileHeaderCreationDate",
            "value": 1493,
            "unit": "ns/op\t     168 B/op\t       9 allocs/op",
            "extra": "791096 times\n4 procs"
          },
          {
            "name": "BenchmarkFileHeaderCreationDate - ns/op",
            "value": 1493,
            "unit": "ns/op",
            "extra": "791096 times\n4 procs"
          },
          {
            "name": "BenchmarkFileHeaderCreationDate - B/op",
            "value": 168,
            "unit": "B/op",
            "extra": "791096 times\n4 procs"
          },
          {
            "name": "BenchmarkFileHeaderCreationDate - allocs/op",
            "value": 9,
            "unit": "allocs/op",
            "extra": "791096 times\n4 procs"
          },
          {
            "name": "BenchmarkFileHeaderCreationTime",
            "value": 1301,
            "unit": "ns/op\t     152 B/op\t       8 allocs/op",
            "extra": "895989 times\n4 procs"
          },
          {
            "name": "BenchmarkFileHeaderCreationTime - ns/op",
            "value": 1301,
            "unit": "ns/op",
            "extra": "895989 times\n4 procs"
          },
          {
            "name": "BenchmarkFileHeaderCreationTime - B/op",
            "value": 152,
            "unit": "B/op",
            "extra": "895989 times\n4 procs"
          },
          {
            "name": "BenchmarkFileHeaderCreationTime - allocs/op",
            "value": 8,
            "unit": "allocs/op",
            "extra": "895989 times\n4 procs"
          },
          {
            "name": "BenchmarkFileError",
            "value": 148.1,
            "unit": "ns/op\t      56 B/op\t       3 allocs/op",
            "extra": "8049858 times\n4 procs"
          },
          {
            "name": "BenchmarkFileError - ns/op",
            "value": 148.1,
            "unit": "ns/op",
            "extra": "8049858 times\n4 procs"
          },
          {
            "name": "BenchmarkFileError - B/op",
            "value": 56,
            "unit": "B/op",
            "extra": "8049858 times\n4 procs"
          },
          {
            "name": "BenchmarkFileError - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "8049858 times\n4 procs"
          },
          {
            "name": "BenchmarkFileBatchCount",
            "value": 8470,
            "unit": "ns/op\t    2952 B/op\t      35 allocs/op",
            "extra": "139538 times\n4 procs"
          },
          {
            "name": "BenchmarkFileBatchCount - ns/op",
            "value": 8470,
            "unit": "ns/op",
            "extra": "139538 times\n4 procs"
          },
          {
            "name": "BenchmarkFileBatchCount - B/op",
            "value": 2952,
            "unit": "B/op",
            "extra": "139538 times\n4 procs"
          },
          {
            "name": "BenchmarkFileBatchCount - allocs/op",
            "value": 35,
            "unit": "allocs/op",
            "extra": "139538 times\n4 procs"
          },
          {
            "name": "BenchmarkFileEntryAddenda",
            "value": 6451,
            "unit": "ns/op\t    1880 B/op\t      23 allocs/op",
            "extra": "181398 times\n4 procs"
          },
          {
            "name": "BenchmarkFileEntryAddenda - ns/op",
            "value": 6451,
            "unit": "ns/op",
            "extra": "181398 times\n4 procs"
          },
          {
            "name": "BenchmarkFileEntryAddenda - B/op",
            "value": 1880,
            "unit": "B/op",
            "extra": "181398 times\n4 procs"
          },
          {
            "name": "BenchmarkFileEntryAddenda - allocs/op",
            "value": 23,
            "unit": "allocs/op",
            "extra": "181398 times\n4 procs"
          },
          {
            "name": "BenchmarkFileDebitAmount",
            "value": 6450,
            "unit": "ns/op\t    1912 B/op\t      23 allocs/op",
            "extra": "180663 times\n4 procs"
          },
          {
            "name": "BenchmarkFileDebitAmount - ns/op",
            "value": 6450,
            "unit": "ns/op",
            "extra": "180663 times\n4 procs"
          },
          {
            "name": "BenchmarkFileDebitAmount - B/op",
            "value": 1912,
            "unit": "B/op",
            "extra": "180663 times\n4 procs"
          },
          {
            "name": "BenchmarkFileDebitAmount - allocs/op",
            "value": 23,
            "unit": "allocs/op",
            "extra": "180663 times\n4 procs"
          },
          {
            "name": "BenchmarkFileCreditAmount",
            "value": 6535,
            "unit": "ns/op\t    1928 B/op\t      25 allocs/op",
            "extra": "179186 times\n4 procs"
          },
          {
            "name": "BenchmarkFileCreditAmount - ns/op",
            "value": 6535,
            "unit": "ns/op",
            "extra": "179186 times\n4 procs"
          },
          {
            "name": "BenchmarkFileCreditAmount - B/op",
            "value": 1928,
            "unit": "B/op",
            "extra": "179186 times\n4 procs"
          },
          {
            "name": "BenchmarkFileCreditAmount - allocs/op",
            "value": 25,
            "unit": "allocs/op",
            "extra": "179186 times\n4 procs"
          },
          {
            "name": "BenchmarkFileEntryHash",
            "value": 9596,
            "unit": "ns/op\t    3000 B/op\t      37 allocs/op",
            "extra": "124040 times\n4 procs"
          },
          {
            "name": "BenchmarkFileEntryHash - ns/op",
            "value": 9596,
            "unit": "ns/op",
            "extra": "124040 times\n4 procs"
          },
          {
            "name": "BenchmarkFileEntryHash - B/op",
            "value": 3000,
            "unit": "B/op",
            "extra": "124040 times\n4 procs"
          },
          {
            "name": "BenchmarkFileEntryHash - allocs/op",
            "value": 37,
            "unit": "allocs/op",
            "extra": "124040 times\n4 procs"
          },
          {
            "name": "BenchmarkFileBlockCount10",
            "value": 8096,
            "unit": "ns/op\t    2984 B/op\t      47 allocs/op",
            "extra": "146079 times\n4 procs"
          },
          {
            "name": "BenchmarkFileBlockCount10 - ns/op",
            "value": 8096,
            "unit": "ns/op",
            "extra": "146079 times\n4 procs"
          },
          {
            "name": "BenchmarkFileBlockCount10 - B/op",
            "value": 2984,
            "unit": "B/op",
            "extra": "146079 times\n4 procs"
          },
          {
            "name": "BenchmarkFileBlockCount10 - allocs/op",
            "value": 47,
            "unit": "allocs/op",
            "extra": "146079 times\n4 procs"
          },
          {
            "name": "BenchmarkFileBuildBadFileHeader",
            "value": 121.3,
            "unit": "ns/op\t      80 B/op\t       2 allocs/op",
            "extra": "9792028 times\n4 procs"
          },
          {
            "name": "BenchmarkFileBuildBadFileHeader - ns/op",
            "value": 121.3,
            "unit": "ns/op",
            "extra": "9792028 times\n4 procs"
          },
          {
            "name": "BenchmarkFileBuildBadFileHeader - B/op",
            "value": 80,
            "unit": "B/op",
            "extra": "9792028 times\n4 procs"
          },
          {
            "name": "BenchmarkFileBuildBadFileHeader - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "9792028 times\n4 procs"
          },
          {
            "name": "BenchmarkFileBuildNoBatch",
            "value": 339,
            "unit": "ns/op\t       8 B/op\t       1 allocs/op",
            "extra": "3519092 times\n4 procs"
          },
          {
            "name": "BenchmarkFileBuildNoBatch - ns/op",
            "value": 339,
            "unit": "ns/op",
            "extra": "3519092 times\n4 procs"
          },
          {
            "name": "BenchmarkFileBuildNoBatch - B/op",
            "value": 8,
            "unit": "B/op",
            "extra": "3519092 times\n4 procs"
          },
          {
            "name": "BenchmarkFileBuildNoBatch - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "3519092 times\n4 procs"
          },
          {
            "name": "BenchmarkFileNotificationOfChange",
            "value": 1421,
            "unit": "ns/op\t     984 B/op\t      13 allocs/op",
            "extra": "763456 times\n4 procs"
          },
          {
            "name": "BenchmarkFileNotificationOfChange - ns/op",
            "value": 1421,
            "unit": "ns/op",
            "extra": "763456 times\n4 procs"
          },
          {
            "name": "BenchmarkFileNotificationOfChange - B/op",
            "value": 984,
            "unit": "B/op",
            "extra": "763456 times\n4 procs"
          },
          {
            "name": "BenchmarkFileNotificationOfChange - allocs/op",
            "value": 13,
            "unit": "allocs/op",
            "extra": "763456 times\n4 procs"
          },
          {
            "name": "BenchmarkFileReturnEntries",
            "value": 1799,
            "unit": "ns/op\t    1048 B/op\t      12 allocs/op",
            "extra": "628347 times\n4 procs"
          },
          {
            "name": "BenchmarkFileReturnEntries - ns/op",
            "value": 1799,
            "unit": "ns/op",
            "extra": "628347 times\n4 procs"
          },
          {
            "name": "BenchmarkFileReturnEntries - B/op",
            "value": 1048,
            "unit": "B/op",
            "extra": "628347 times\n4 procs"
          },
          {
            "name": "BenchmarkFileReturnEntries - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "628347 times\n4 procs"
          },
          {
            "name": "BenchmarkMockIATBatchHeaderFF",
            "value": 340.4,
            "unit": "ns/op\t      38 B/op\t       4 allocs/op",
            "extra": "3517086 times\n4 procs"
          },
          {
            "name": "BenchmarkMockIATBatchHeaderFF - ns/op",
            "value": 340.4,
            "unit": "ns/op",
            "extra": "3517086 times\n4 procs"
          },
          {
            "name": "BenchmarkMockIATBatchHeaderFF - B/op",
            "value": 38,
            "unit": "B/op",
            "extra": "3517086 times\n4 procs"
          },
          {
            "name": "BenchmarkMockIATBatchHeaderFF - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "3517086 times\n4 procs"
          },
          {
            "name": "BenchmarkParseIATBatchHeader",
            "value": 6212,
            "unit": "ns/op\t   15240 B/op\t      33 allocs/op",
            "extra": "191692 times\n4 procs"
          },
          {
            "name": "BenchmarkParseIATBatchHeader - ns/op",
            "value": 6212,
            "unit": "ns/op",
            "extra": "191692 times\n4 procs"
          },
          {
            "name": "BenchmarkParseIATBatchHeader - B/op",
            "value": 15240,
            "unit": "B/op",
            "extra": "191692 times\n4 procs"
          },
          {
            "name": "BenchmarkParseIATBatchHeader - allocs/op",
            "value": 33,
            "unit": "allocs/op",
            "extra": "191692 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBHString",
            "value": 6723,
            "unit": "ns/op\t   15351 B/op\t      36 allocs/op",
            "extra": "175248 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBHString - ns/op",
            "value": 6723,
            "unit": "ns/op",
            "extra": "175248 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBHString - B/op",
            "value": 15351,
            "unit": "B/op",
            "extra": "175248 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBHString - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "175248 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBHFVString",
            "value": 6780,
            "unit": "ns/op\t   15351 B/op\t      36 allocs/op",
            "extra": "171072 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBHFVString - ns/op",
            "value": 6780,
            "unit": "ns/op",
            "extra": "171072 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBHFVString - B/op",
            "value": 15351,
            "unit": "B/op",
            "extra": "171072 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBHFVString - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "171072 times\n4 procs"
          },
          {
            "name": "BenchmarkValidateIATBHServiceClassCode",
            "value": 124.8,
            "unit": "ns/op\t      83 B/op\t       3 allocs/op",
            "extra": "9402367 times\n4 procs"
          },
          {
            "name": "BenchmarkValidateIATBHServiceClassCode - ns/op",
            "value": 124.8,
            "unit": "ns/op",
            "extra": "9402367 times\n4 procs"
          },
          {
            "name": "BenchmarkValidateIATBHServiceClassCode - B/op",
            "value": 83,
            "unit": "B/op",
            "extra": "9402367 times\n4 procs"
          },
          {
            "name": "BenchmarkValidateIATBHServiceClassCode - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "9402367 times\n4 procs"
          },
          {
            "name": "BenchmarkValidateIATBHForeignExchangeIndicator",
            "value": 96.39,
            "unit": "ns/op\t      80 B/op\t       2 allocs/op",
            "extra": "12340862 times\n4 procs"
          },
          {
            "name": "BenchmarkValidateIATBHForeignExchangeIndicator - ns/op",
            "value": 96.39,
            "unit": "ns/op",
            "extra": "12340862 times\n4 procs"
          },
          {
            "name": "BenchmarkValidateIATBHForeignExchangeIndicator - B/op",
            "value": 80,
            "unit": "B/op",
            "extra": "12340862 times\n4 procs"
          },
          {
            "name": "BenchmarkValidateIATBHForeignExchangeIndicator - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "12340862 times\n4 procs"
          },
          {
            "name": "BenchmarkValidateIATBHForeignExchangeReferenceIndicator",
            "value": 100.2,
            "unit": "ns/op\t      80 B/op\t       2 allocs/op",
            "extra": "11966965 times\n4 procs"
          },
          {
            "name": "BenchmarkValidateIATBHForeignExchangeReferenceIndicator - ns/op",
            "value": 100.2,
            "unit": "ns/op",
            "extra": "11966965 times\n4 procs"
          },
          {
            "name": "BenchmarkValidateIATBHForeignExchangeReferenceIndicator - B/op",
            "value": 80,
            "unit": "B/op",
            "extra": "11966965 times\n4 procs"
          },
          {
            "name": "BenchmarkValidateIATBHForeignExchangeReferenceIndicator - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "11966965 times\n4 procs"
          },
          {
            "name": "BenchmarkValidateIATBHISODestinationCountryCode",
            "value": 135.1,
            "unit": "ns/op\t      80 B/op\t       2 allocs/op",
            "extra": "8823482 times\n4 procs"
          },
          {
            "name": "BenchmarkValidateIATBHISODestinationCountryCode - ns/op",
            "value": 135.1,
            "unit": "ns/op",
            "extra": "8823482 times\n4 procs"
          },
          {
            "name": "BenchmarkValidateIATBHISODestinationCountryCode - B/op",
            "value": 80,
            "unit": "B/op",
            "extra": "8823482 times\n4 procs"
          },
          {
            "name": "BenchmarkValidateIATBHISODestinationCountryCode - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "8823482 times\n4 procs"
          },
          {
            "name": "BenchmarkValidateIATBHOriginatorIdentification",
            "value": 329,
            "unit": "ns/op\t      38 B/op\t       4 allocs/op",
            "extra": "3642969 times\n4 procs"
          },
          {
            "name": "BenchmarkValidateIATBHOriginatorIdentification - ns/op",
            "value": 329,
            "unit": "ns/op",
            "extra": "3642969 times\n4 procs"
          },
          {
            "name": "BenchmarkValidateIATBHOriginatorIdentification - B/op",
            "value": 38,
            "unit": "B/op",
            "extra": "3642969 times\n4 procs"
          },
          {
            "name": "BenchmarkValidateIATBHOriginatorIdentification - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "3642969 times\n4 procs"
          },
          {
            "name": "BenchmarkValidateIATBHStandardEntryClassCode",
            "value": 119.4,
            "unit": "ns/op\t      80 B/op\t       2 allocs/op",
            "extra": "9969762 times\n4 procs"
          },
          {
            "name": "BenchmarkValidateIATBHStandardEntryClassCode - ns/op",
            "value": 119.4,
            "unit": "ns/op",
            "extra": "9969762 times\n4 procs"
          },
          {
            "name": "BenchmarkValidateIATBHStandardEntryClassCode - B/op",
            "value": 80,
            "unit": "B/op",
            "extra": "9969762 times\n4 procs"
          },
          {
            "name": "BenchmarkValidateIATBHStandardEntryClassCode - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "9969762 times\n4 procs"
          },
          {
            "name": "BenchmarkValidateIATBHCompanyEntryDescription",
            "value": 378.1,
            "unit": "ns/op\t     160 B/op\t       4 allocs/op",
            "extra": "3155280 times\n4 procs"
          },
          {
            "name": "BenchmarkValidateIATBHCompanyEntryDescription - ns/op",
            "value": 378.1,
            "unit": "ns/op",
            "extra": "3155280 times\n4 procs"
          },
          {
            "name": "BenchmarkValidateIATBHCompanyEntryDescription - B/op",
            "value": 160,
            "unit": "B/op",
            "extra": "3155280 times\n4 procs"
          },
          {
            "name": "BenchmarkValidateIATBHCompanyEntryDescription - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "3155280 times\n4 procs"
          },
          {
            "name": "BenchmarkValidateIATBHISOOriginatingCurrencyCode",
            "value": 374.9,
            "unit": "ns/op\t     100 B/op\t       4 allocs/op",
            "extra": "3192037 times\n4 procs"
          },
          {
            "name": "BenchmarkValidateIATBHISOOriginatingCurrencyCode - ns/op",
            "value": 374.9,
            "unit": "ns/op",
            "extra": "3192037 times\n4 procs"
          },
          {
            "name": "BenchmarkValidateIATBHISOOriginatingCurrencyCode - B/op",
            "value": 100,
            "unit": "B/op",
            "extra": "3192037 times\n4 procs"
          },
          {
            "name": "BenchmarkValidateIATBHISOOriginatingCurrencyCode - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "3192037 times\n4 procs"
          },
          {
            "name": "BenchmarkValidateIATBHISODestinationCurrencyCode",
            "value": 538.6,
            "unit": "ns/op\t     120 B/op\t       6 allocs/op",
            "extra": "2230220 times\n4 procs"
          },
          {
            "name": "BenchmarkValidateIATBHISODestinationCurrencyCode - ns/op",
            "value": 538.6,
            "unit": "ns/op",
            "extra": "2230220 times\n4 procs"
          },
          {
            "name": "BenchmarkValidateIATBHISODestinationCurrencyCode - B/op",
            "value": 120,
            "unit": "B/op",
            "extra": "2230220 times\n4 procs"
          },
          {
            "name": "BenchmarkValidateIATBHISODestinationCurrencyCode - allocs/op",
            "value": 6,
            "unit": "allocs/op",
            "extra": "2230220 times\n4 procs"
          },
          {
            "name": "BenchmarkValidateIATBHOriginatorStatusCode",
            "value": 442.4,
            "unit": "ns/op\t     118 B/op\t       6 allocs/op",
            "extra": "2692219 times\n4 procs"
          },
          {
            "name": "BenchmarkValidateIATBHOriginatorStatusCode - ns/op",
            "value": 442.4,
            "unit": "ns/op",
            "extra": "2692219 times\n4 procs"
          },
          {
            "name": "BenchmarkValidateIATBHOriginatorStatusCode - B/op",
            "value": 118,
            "unit": "B/op",
            "extra": "2692219 times\n4 procs"
          },
          {
            "name": "BenchmarkValidateIATBHOriginatorStatusCode - allocs/op",
            "value": 6,
            "unit": "allocs/op",
            "extra": "2692219 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBHServiceClassCode",
            "value": 95.62,
            "unit": "ns/op\t      80 B/op\t       2 allocs/op",
            "extra": "12410007 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBHServiceClassCode - ns/op",
            "value": 95.62,
            "unit": "ns/op",
            "extra": "12410007 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBHServiceClassCode - B/op",
            "value": 80,
            "unit": "B/op",
            "extra": "12410007 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBHServiceClassCode - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "12410007 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBHForeignExchangeIndicator",
            "value": 66.53,
            "unit": "ns/op\t      64 B/op\t       1 allocs/op",
            "extra": "17744398 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBHForeignExchangeIndicator - ns/op",
            "value": 66.53,
            "unit": "ns/op",
            "extra": "17744398 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBHForeignExchangeIndicator - B/op",
            "value": 64,
            "unit": "B/op",
            "extra": "17744398 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBHForeignExchangeIndicator - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "17744398 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBHForeignExchangeReferenceIndicator",
            "value": 100.3,
            "unit": "ns/op\t      80 B/op\t       2 allocs/op",
            "extra": "11844198 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBHForeignExchangeReferenceIndicator - ns/op",
            "value": 100.3,
            "unit": "ns/op",
            "extra": "11844198 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBHForeignExchangeReferenceIndicator - B/op",
            "value": 80,
            "unit": "B/op",
            "extra": "11844198 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBHForeignExchangeReferenceIndicator - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "11844198 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBHISODestinationCountryCode",
            "value": 68.63,
            "unit": "ns/op\t      64 B/op\t       1 allocs/op",
            "extra": "17375691 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBHISODestinationCountryCode - ns/op",
            "value": 68.63,
            "unit": "ns/op",
            "extra": "17375691 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBHISODestinationCountryCode - B/op",
            "value": 64,
            "unit": "B/op",
            "extra": "17375691 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBHISODestinationCountryCode - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "17375691 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBHOriginatorIdentification",
            "value": 68.43,
            "unit": "ns/op\t      64 B/op\t       1 allocs/op",
            "extra": "17399191 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBHOriginatorIdentification - ns/op",
            "value": 68.43,
            "unit": "ns/op",
            "extra": "17399191 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBHOriginatorIdentification - B/op",
            "value": 64,
            "unit": "B/op",
            "extra": "17399191 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBHOriginatorIdentification - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "17399191 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBHStandardEntryClassCode",
            "value": 72.82,
            "unit": "ns/op\t      64 B/op\t       1 allocs/op",
            "extra": "16730994 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBHStandardEntryClassCode - ns/op",
            "value": 72.82,
            "unit": "ns/op",
            "extra": "16730994 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBHStandardEntryClassCode - B/op",
            "value": 64,
            "unit": "B/op",
            "extra": "16730994 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBHStandardEntryClassCode - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "16730994 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBHCompanyEntryDescription",
            "value": 73.22,
            "unit": "ns/op\t      64 B/op\t       1 allocs/op",
            "extra": "16205394 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBHCompanyEntryDescription - ns/op",
            "value": 73.22,
            "unit": "ns/op",
            "extra": "16205394 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBHCompanyEntryDescription - B/op",
            "value": 64,
            "unit": "B/op",
            "extra": "16205394 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBHCompanyEntryDescription - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "16205394 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBHISOOriginatingCurrencyCode",
            "value": 67.58,
            "unit": "ns/op\t      64 B/op\t       1 allocs/op",
            "extra": "17686522 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBHISOOriginatingCurrencyCode - ns/op",
            "value": 67.58,
            "unit": "ns/op",
            "extra": "17686522 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBHISOOriginatingCurrencyCode - B/op",
            "value": 64,
            "unit": "B/op",
            "extra": "17686522 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBHISOOriginatingCurrencyCode - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "17686522 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBHISODestinationCurrencyCode",
            "value": 68.04,
            "unit": "ns/op\t      64 B/op\t       1 allocs/op",
            "extra": "17397615 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBHISODestinationCurrencyCode - ns/op",
            "value": 68.04,
            "unit": "ns/op",
            "extra": "17397615 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBHISODestinationCurrencyCode - B/op",
            "value": 64,
            "unit": "B/op",
            "extra": "17397615 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBHISODestinationCurrencyCode - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "17397615 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBHODFIIdentification",
            "value": 113.5,
            "unit": "ns/op\t      80 B/op\t       2 allocs/op",
            "extra": "10497039 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBHODFIIdentification - ns/op",
            "value": 113.5,
            "unit": "ns/op",
            "extra": "10497039 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBHODFIIdentification - B/op",
            "value": 80,
            "unit": "B/op",
            "extra": "10497039 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBHODFIIdentification - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "10497039 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchAddenda10Error",
            "value": 2305,
            "unit": "ns/op\t    1772 B/op\t      26 allocs/op",
            "extra": "500038 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchAddenda10Error - ns/op",
            "value": 2305,
            "unit": "ns/op",
            "extra": "500038 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchAddenda10Error - B/op",
            "value": 1772,
            "unit": "B/op",
            "extra": "500038 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchAddenda10Error - allocs/op",
            "value": 26,
            "unit": "allocs/op",
            "extra": "500038 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchAddenda11Error",
            "value": 2296,
            "unit": "ns/op\t    1772 B/op\t      26 allocs/op",
            "extra": "511558 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchAddenda11Error - ns/op",
            "value": 2296,
            "unit": "ns/op",
            "extra": "511558 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchAddenda11Error - B/op",
            "value": 1772,
            "unit": "B/op",
            "extra": "511558 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchAddenda11Error - allocs/op",
            "value": 26,
            "unit": "allocs/op",
            "extra": "511558 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchAddenda12Error",
            "value": 2299,
            "unit": "ns/op\t    1772 B/op\t      26 allocs/op",
            "extra": "493945 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchAddenda12Error - ns/op",
            "value": 2299,
            "unit": "ns/op",
            "extra": "493945 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchAddenda12Error - B/op",
            "value": 1772,
            "unit": "B/op",
            "extra": "493945 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchAddenda12Error - allocs/op",
            "value": 26,
            "unit": "allocs/op",
            "extra": "493945 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchAddenda13Error",
            "value": 2290,
            "unit": "ns/op\t    1772 B/op\t      26 allocs/op",
            "extra": "512163 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchAddenda13Error - ns/op",
            "value": 2290,
            "unit": "ns/op",
            "extra": "512163 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchAddenda13Error - B/op",
            "value": 1772,
            "unit": "B/op",
            "extra": "512163 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchAddenda13Error - allocs/op",
            "value": 26,
            "unit": "allocs/op",
            "extra": "512163 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchAddenda14Error",
            "value": 2295,
            "unit": "ns/op\t    1772 B/op\t      26 allocs/op",
            "extra": "507246 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchAddenda14Error - ns/op",
            "value": 2295,
            "unit": "ns/op",
            "extra": "507246 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchAddenda14Error - B/op",
            "value": 1772,
            "unit": "B/op",
            "extra": "507246 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchAddenda14Error - allocs/op",
            "value": 26,
            "unit": "allocs/op",
            "extra": "507246 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchAddenda15Error",
            "value": 2304,
            "unit": "ns/op\t    1772 B/op\t      26 allocs/op",
            "extra": "506278 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchAddenda15Error - ns/op",
            "value": 2304,
            "unit": "ns/op",
            "extra": "506278 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchAddenda15Error - B/op",
            "value": 1772,
            "unit": "B/op",
            "extra": "506278 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchAddenda15Error - allocs/op",
            "value": 26,
            "unit": "allocs/op",
            "extra": "506278 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchAddenda16Error",
            "value": 2378,
            "unit": "ns/op\t    1772 B/op\t      26 allocs/op",
            "extra": "503568 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchAddenda16Error - ns/op",
            "value": 2378,
            "unit": "ns/op",
            "extra": "503568 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchAddenda16Error - B/op",
            "value": 1772,
            "unit": "B/op",
            "extra": "503568 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchAddenda16Error - allocs/op",
            "value": 26,
            "unit": "allocs/op",
            "extra": "503568 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda10EntryDetailSequenceNumber",
            "value": 3514,
            "unit": "ns/op\t    2000 B/op\t      34 allocs/op",
            "extra": "339417 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda10EntryDetailSequenceNumber - ns/op",
            "value": 3514,
            "unit": "ns/op",
            "extra": "339417 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda10EntryDetailSequenceNumber - B/op",
            "value": 2000,
            "unit": "B/op",
            "extra": "339417 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda10EntryDetailSequenceNumber - allocs/op",
            "value": 34,
            "unit": "allocs/op",
            "extra": "339417 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda11EntryDetailSequenceNumber",
            "value": 3522,
            "unit": "ns/op\t    2008 B/op\t      35 allocs/op",
            "extra": "327976 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda11EntryDetailSequenceNumber - ns/op",
            "value": 3522,
            "unit": "ns/op",
            "extra": "327976 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda11EntryDetailSequenceNumber - B/op",
            "value": 2008,
            "unit": "B/op",
            "extra": "327976 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda11EntryDetailSequenceNumber - allocs/op",
            "value": 35,
            "unit": "allocs/op",
            "extra": "327976 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda12EntryDetailSequenceNumber",
            "value": 3588,
            "unit": "ns/op\t    2016 B/op\t      36 allocs/op",
            "extra": "327094 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda12EntryDetailSequenceNumber - ns/op",
            "value": 3588,
            "unit": "ns/op",
            "extra": "327094 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda12EntryDetailSequenceNumber - B/op",
            "value": 2016,
            "unit": "B/op",
            "extra": "327094 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda12EntryDetailSequenceNumber - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "327094 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda13EntryDetailSequenceNumber",
            "value": 3625,
            "unit": "ns/op\t    2024 B/op\t      37 allocs/op",
            "extra": "321259 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda13EntryDetailSequenceNumber - ns/op",
            "value": 3625,
            "unit": "ns/op",
            "extra": "321259 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda13EntryDetailSequenceNumber - B/op",
            "value": 2024,
            "unit": "B/op",
            "extra": "321259 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda13EntryDetailSequenceNumber - allocs/op",
            "value": 37,
            "unit": "allocs/op",
            "extra": "321259 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda14EntryDetailSequenceNumber",
            "value": 3670,
            "unit": "ns/op\t    2032 B/op\t      38 allocs/op",
            "extra": "318090 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda14EntryDetailSequenceNumber - ns/op",
            "value": 3670,
            "unit": "ns/op",
            "extra": "318090 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda14EntryDetailSequenceNumber - B/op",
            "value": 2032,
            "unit": "B/op",
            "extra": "318090 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda14EntryDetailSequenceNumber - allocs/op",
            "value": 38,
            "unit": "allocs/op",
            "extra": "318090 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda15EntryDetailSequenceNumber",
            "value": 3745,
            "unit": "ns/op\t    2040 B/op\t      39 allocs/op",
            "extra": "312886 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda15EntryDetailSequenceNumber - ns/op",
            "value": 3745,
            "unit": "ns/op",
            "extra": "312886 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda15EntryDetailSequenceNumber - B/op",
            "value": 2040,
            "unit": "B/op",
            "extra": "312886 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda15EntryDetailSequenceNumber - allocs/op",
            "value": 39,
            "unit": "allocs/op",
            "extra": "312886 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda16EntryDetailSequenceNumber",
            "value": 3807,
            "unit": "ns/op\t    2048 B/op\t      40 allocs/op",
            "extra": "309423 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda16EntryDetailSequenceNumber - ns/op",
            "value": 3807,
            "unit": "ns/op",
            "extra": "309423 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda16EntryDetailSequenceNumber - B/op",
            "value": 2048,
            "unit": "B/op",
            "extra": "309423 times\n4 procs"
          },
          {
            "name": "BenchmarkAddenda16EntryDetailSequenceNumber - allocs/op",
            "value": 40,
            "unit": "allocs/op",
            "extra": "309423 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchNumberMismatch",
            "value": 3041,
            "unit": "ns/op\t    1852 B/op\t      28 allocs/op",
            "extra": "352125 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchNumberMismatch - ns/op",
            "value": 3041,
            "unit": "ns/op",
            "extra": "352125 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchNumberMismatch - B/op",
            "value": 1852,
            "unit": "B/op",
            "extra": "352125 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchNumberMismatch - allocs/op",
            "value": 28,
            "unit": "allocs/op",
            "extra": "352125 times\n4 procs"
          },
          {
            "name": "BenchmarkIATServiceClassCodeMismatch",
            "value": 3059,
            "unit": "ns/op\t    1852 B/op\t      28 allocs/op",
            "extra": "387344 times\n4 procs"
          },
          {
            "name": "BenchmarkIATServiceClassCodeMismatch - ns/op",
            "value": 3059,
            "unit": "ns/op",
            "extra": "387344 times\n4 procs"
          },
          {
            "name": "BenchmarkIATServiceClassCodeMismatch - B/op",
            "value": 1852,
            "unit": "B/op",
            "extra": "387344 times\n4 procs"
          },
          {
            "name": "BenchmarkIATServiceClassCodeMismatch - allocs/op",
            "value": 28,
            "unit": "allocs/op",
            "extra": "387344 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchCreditIsBatchAmount",
            "value": 5384,
            "unit": "ns/op\t    3264 B/op\t      50 allocs/op",
            "extra": "218341 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchCreditIsBatchAmount - ns/op",
            "value": 5384,
            "unit": "ns/op",
            "extra": "218341 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchCreditIsBatchAmount - B/op",
            "value": 3264,
            "unit": "B/op",
            "extra": "218341 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchCreditIsBatchAmount - allocs/op",
            "value": 50,
            "unit": "allocs/op",
            "extra": "218341 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchDebitIsBatchAmount",
            "value": 5412,
            "unit": "ns/op\t    3264 B/op\t      50 allocs/op",
            "extra": "217846 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchDebitIsBatchAmount - ns/op",
            "value": 5412,
            "unit": "ns/op",
            "extra": "217846 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchDebitIsBatchAmount - B/op",
            "value": 3264,
            "unit": "B/op",
            "extra": "217846 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchDebitIsBatchAmount - allocs/op",
            "value": 50,
            "unit": "allocs/op",
            "extra": "217846 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchFieldInclusion",
            "value": 5133,
            "unit": "ns/op\t    3520 B/op\t      56 allocs/op",
            "extra": "231200 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchFieldInclusion - ns/op",
            "value": 5133,
            "unit": "ns/op",
            "extra": "231200 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchFieldInclusion - B/op",
            "value": 3520,
            "unit": "B/op",
            "extra": "231200 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchFieldInclusion - allocs/op",
            "value": 56,
            "unit": "allocs/op",
            "extra": "231200 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchBuild",
            "value": 498,
            "unit": "ns/op\t     374 B/op\t       6 allocs/op",
            "extra": "2416350 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchBuild - ns/op",
            "value": 498,
            "unit": "ns/op",
            "extra": "2416350 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchBuild - B/op",
            "value": 374,
            "unit": "B/op",
            "extra": "2416350 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchBuild - allocs/op",
            "value": 6,
            "unit": "allocs/op",
            "extra": "2416350 times\n4 procs"
          },
          {
            "name": "BenchmarkIATODFIIdentificationMismatch",
            "value": 3133,
            "unit": "ns/op\t    1884 B/op\t      30 allocs/op",
            "extra": "373365 times\n4 procs"
          },
          {
            "name": "BenchmarkIATODFIIdentificationMismatch - ns/op",
            "value": 3133,
            "unit": "ns/op",
            "extra": "373365 times\n4 procs"
          },
          {
            "name": "BenchmarkIATODFIIdentificationMismatch - B/op",
            "value": 1884,
            "unit": "B/op",
            "extra": "373365 times\n4 procs"
          },
          {
            "name": "BenchmarkIATODFIIdentificationMismatch - allocs/op",
            "value": 30,
            "unit": "allocs/op",
            "extra": "373365 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchAddendaRecordIndicator",
            "value": 2858,
            "unit": "ns/op\t    1708 B/op\t      25 allocs/op",
            "extra": "408723 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchAddendaRecordIndicator - ns/op",
            "value": 2858,
            "unit": "ns/op",
            "extra": "408723 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchAddendaRecordIndicator - B/op",
            "value": 1708,
            "unit": "B/op",
            "extra": "408723 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchAddendaRecordIndicator - allocs/op",
            "value": 25,
            "unit": "allocs/op",
            "extra": "408723 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchInvalidTraceNumberODFI",
            "value": 3548,
            "unit": "ns/op\t    1984 B/op\t      35 allocs/op",
            "extra": "316164 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchInvalidTraceNumberODFI - ns/op",
            "value": 3548,
            "unit": "ns/op",
            "extra": "316164 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchInvalidTraceNumberODFI - B/op",
            "value": 1984,
            "unit": "B/op",
            "extra": "316164 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchInvalidTraceNumberODFI - allocs/op",
            "value": 35,
            "unit": "allocs/op",
            "extra": "316164 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchControl",
            "value": 3034,
            "unit": "ns/op\t    1868 B/op\t      29 allocs/op",
            "extra": "377055 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchControl - ns/op",
            "value": 3034,
            "unit": "ns/op",
            "extra": "377055 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchControl - B/op",
            "value": 1868,
            "unit": "B/op",
            "extra": "377055 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchControl - allocs/op",
            "value": 29,
            "unit": "allocs/op",
            "extra": "377055 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchEntryCountEquality",
            "value": 3941,
            "unit": "ns/op\t    2243 B/op\t      36 allocs/op",
            "extra": "296252 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchEntryCountEquality - ns/op",
            "value": 3941,
            "unit": "ns/op",
            "extra": "296252 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchEntryCountEquality - B/op",
            "value": 2243,
            "unit": "B/op",
            "extra": "296252 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchEntryCountEquality - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "296252 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchisEntryHash",
            "value": 3199,
            "unit": "ns/op\t    1896 B/op\t      29 allocs/op",
            "extra": "363130 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchisEntryHash - ns/op",
            "value": 3199,
            "unit": "ns/op",
            "extra": "363130 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchisEntryHash - B/op",
            "value": 1896,
            "unit": "B/op",
            "extra": "363130 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchisEntryHash - allocs/op",
            "value": 29,
            "unit": "allocs/op",
            "extra": "363130 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchIsSequenceAscending",
            "value": 4298,
            "unit": "ns/op\t    2968 B/op\t      41 allocs/op",
            "extra": "272744 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchIsSequenceAscending - ns/op",
            "value": 4298,
            "unit": "ns/op",
            "extra": "272744 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchIsSequenceAscending - B/op",
            "value": 2968,
            "unit": "B/op",
            "extra": "272744 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchIsSequenceAscending - allocs/op",
            "value": 41,
            "unit": "allocs/op",
            "extra": "272744 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchIsCategory",
            "value": 6051,
            "unit": "ns/op\t    4649 B/op\t      63 allocs/op",
            "extra": "192715 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchIsCategory - ns/op",
            "value": 6051,
            "unit": "ns/op",
            "extra": "192715 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchIsCategory - B/op",
            "value": 4649,
            "unit": "B/op",
            "extra": "192715 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchIsCategory - allocs/op",
            "value": 63,
            "unit": "allocs/op",
            "extra": "192715 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchCategory",
            "value": 2546,
            "unit": "ns/op\t    1948 B/op\t      28 allocs/op",
            "extra": "456160 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchCategory - ns/op",
            "value": 2546,
            "unit": "ns/op",
            "extra": "456160 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchCategory - B/op",
            "value": 1948,
            "unit": "B/op",
            "extra": "456160 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchCategory - allocs/op",
            "value": 28,
            "unit": "allocs/op",
            "extra": "456160 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchValidateEntry",
            "value": 2211,
            "unit": "ns/op\t    1788 B/op\t      27 allocs/op",
            "extra": "521326 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchValidateEntry - ns/op",
            "value": 2211,
            "unit": "ns/op",
            "extra": "521326 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchValidateEntry - B/op",
            "value": 1788,
            "unit": "B/op",
            "extra": "521326 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchValidateEntry - allocs/op",
            "value": 27,
            "unit": "allocs/op",
            "extra": "521326 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchValidateAddenda10",
            "value": 5398,
            "unit": "ns/op\t    4665 B/op\t      64 allocs/op",
            "extra": "218020 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchValidateAddenda10 - ns/op",
            "value": 5398,
            "unit": "ns/op",
            "extra": "218020 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchValidateAddenda10 - B/op",
            "value": 4665,
            "unit": "B/op",
            "extra": "218020 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchValidateAddenda10 - allocs/op",
            "value": 64,
            "unit": "allocs/op",
            "extra": "218020 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchValidateAddenda11",
            "value": 5502,
            "unit": "ns/op\t    4665 B/op\t      64 allocs/op",
            "extra": "217768 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchValidateAddenda11 - ns/op",
            "value": 5502,
            "unit": "ns/op",
            "extra": "217768 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchValidateAddenda11 - B/op",
            "value": 4665,
            "unit": "B/op",
            "extra": "217768 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchValidateAddenda11 - allocs/op",
            "value": 64,
            "unit": "allocs/op",
            "extra": "217768 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchValidateAddenda12",
            "value": 5510,
            "unit": "ns/op\t    4665 B/op\t      64 allocs/op",
            "extra": "211737 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchValidateAddenda12 - ns/op",
            "value": 5510,
            "unit": "ns/op",
            "extra": "211737 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchValidateAddenda12 - B/op",
            "value": 4665,
            "unit": "B/op",
            "extra": "211737 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchValidateAddenda12 - allocs/op",
            "value": 64,
            "unit": "allocs/op",
            "extra": "211737 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchValidateAddenda13",
            "value": 5571,
            "unit": "ns/op\t    4665 B/op\t      64 allocs/op",
            "extra": "211032 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchValidateAddenda13 - ns/op",
            "value": 5571,
            "unit": "ns/op",
            "extra": "211032 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchValidateAddenda13 - B/op",
            "value": 4665,
            "unit": "B/op",
            "extra": "211032 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchValidateAddenda13 - allocs/op",
            "value": 64,
            "unit": "allocs/op",
            "extra": "211032 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchValidateAddenda14",
            "value": 5603,
            "unit": "ns/op\t    4665 B/op\t      64 allocs/op",
            "extra": "210934 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchValidateAddenda14 - ns/op",
            "value": 5603,
            "unit": "ns/op",
            "extra": "210934 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchValidateAddenda14 - B/op",
            "value": 4665,
            "unit": "B/op",
            "extra": "210934 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchValidateAddenda14 - allocs/op",
            "value": 64,
            "unit": "allocs/op",
            "extra": "210934 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchValidateAddenda15",
            "value": 5661,
            "unit": "ns/op\t    4665 B/op\t      64 allocs/op",
            "extra": "207496 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchValidateAddenda15 - ns/op",
            "value": 5661,
            "unit": "ns/op",
            "extra": "207496 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchValidateAddenda15 - B/op",
            "value": 4665,
            "unit": "B/op",
            "extra": "207496 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchValidateAddenda15 - allocs/op",
            "value": 64,
            "unit": "allocs/op",
            "extra": "207496 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchValidateAddenda16",
            "value": 5703,
            "unit": "ns/op\t    4665 B/op\t      64 allocs/op",
            "extra": "208803 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchValidateAddenda16 - ns/op",
            "value": 5703,
            "unit": "ns/op",
            "extra": "208803 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchValidateAddenda16 - B/op",
            "value": 4665,
            "unit": "B/op",
            "extra": "208803 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchValidateAddenda16 - allocs/op",
            "value": 64,
            "unit": "allocs/op",
            "extra": "208803 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchValidateAddenda17",
            "value": 2867,
            "unit": "ns/op\t    1876 B/op\t      29 allocs/op",
            "extra": "401578 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchValidateAddenda17 - ns/op",
            "value": 2867,
            "unit": "ns/op",
            "extra": "401578 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchValidateAddenda17 - B/op",
            "value": 1876,
            "unit": "B/op",
            "extra": "401578 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchValidateAddenda17 - allocs/op",
            "value": 29,
            "unit": "allocs/op",
            "extra": "401578 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchCreate",
            "value": 2202,
            "unit": "ns/op\t    1760 B/op\t      24 allocs/op",
            "extra": "524104 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchCreate - ns/op",
            "value": 2202,
            "unit": "ns/op",
            "extra": "524104 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchCreate - B/op",
            "value": 1760,
            "unit": "B/op",
            "extra": "524104 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchCreate - allocs/op",
            "value": 24,
            "unit": "allocs/op",
            "extra": "524104 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchValidate",
            "value": 3487,
            "unit": "ns/op\t    1944 B/op\t      30 allocs/op",
            "extra": "334414 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchValidate - ns/op",
            "value": 3487,
            "unit": "ns/op",
            "extra": "334414 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchValidate - B/op",
            "value": 1944,
            "unit": "B/op",
            "extra": "334414 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchValidate - allocs/op",
            "value": 30,
            "unit": "allocs/op",
            "extra": "334414 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchEntryAddendum",
            "value": 6842,
            "unit": "ns/op\t    3504 B/op\t      67 allocs/op",
            "extra": "170972 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchEntryAddendum - ns/op",
            "value": 6842,
            "unit": "ns/op",
            "extra": "170972 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchEntryAddendum - B/op",
            "value": 3504,
            "unit": "B/op",
            "extra": "170972 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchEntryAddendum - allocs/op",
            "value": 67,
            "unit": "allocs/op",
            "extra": "170972 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchAddenda17EDSequenceNumber",
            "value": 4232,
            "unit": "ns/op\t    2248 B/op\t      46 allocs/op",
            "extra": "277299 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchAddenda17EDSequenceNumber - ns/op",
            "value": 4232,
            "unit": "ns/op",
            "extra": "277299 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchAddenda17EDSequenceNumber - B/op",
            "value": 2248,
            "unit": "B/op",
            "extra": "277299 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchAddenda17EDSequenceNumber - allocs/op",
            "value": 46,
            "unit": "allocs/op",
            "extra": "277299 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchAddenda17Sequence",
            "value": 4113,
            "unit": "ns/op\t    2176 B/op\t      41 allocs/op",
            "extra": "285570 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchAddenda17Sequence - ns/op",
            "value": 4113,
            "unit": "ns/op",
            "extra": "285570 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchAddenda17Sequence - B/op",
            "value": 2176,
            "unit": "B/op",
            "extra": "285570 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchAddenda17Sequence - allocs/op",
            "value": 41,
            "unit": "allocs/op",
            "extra": "285570 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchAddenda18EDSequenceNumber",
            "value": 4795,
            "unit": "ns/op\t    2544 B/op\t      52 allocs/op",
            "extra": "244545 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchAddenda18EDSequenceNumber - ns/op",
            "value": 4795,
            "unit": "ns/op",
            "extra": "244545 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchAddenda18EDSequenceNumber - B/op",
            "value": 2544,
            "unit": "B/op",
            "extra": "244545 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchAddenda18EDSequenceNumber - allocs/op",
            "value": 52,
            "unit": "allocs/op",
            "extra": "244545 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchAddenda18Sequence",
            "value": 4671,
            "unit": "ns/op\t    2472 B/op\t      47 allocs/op",
            "extra": "246854 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchAddenda18Sequence - ns/op",
            "value": 4671,
            "unit": "ns/op",
            "extra": "246854 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchAddenda18Sequence - B/op",
            "value": 2472,
            "unit": "B/op",
            "extra": "246854 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchAddenda18Sequence - allocs/op",
            "value": 47,
            "unit": "allocs/op",
            "extra": "246854 times\n4 procs"
          },
          {
            "name": "BenchmarkIATNoEntry",
            "value": 143.2,
            "unit": "ns/op\t     336 B/op\t       2 allocs/op",
            "extra": "8416939 times\n4 procs"
          },
          {
            "name": "BenchmarkIATNoEntry - ns/op",
            "value": 143.2,
            "unit": "ns/op",
            "extra": "8416939 times\n4 procs"
          },
          {
            "name": "BenchmarkIATNoEntry - B/op",
            "value": 336,
            "unit": "B/op",
            "extra": "8416939 times\n4 procs"
          },
          {
            "name": "BenchmarkIATNoEntry - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "8416939 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchAddendumTypeCode",
            "value": 4300,
            "unit": "ns/op\t    2144 B/op\t      42 allocs/op",
            "extra": "270784 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchAddendumTypeCode - ns/op",
            "value": 4300,
            "unit": "ns/op",
            "extra": "270784 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchAddendumTypeCode - B/op",
            "value": 2144,
            "unit": "B/op",
            "extra": "270784 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchAddendumTypeCode - allocs/op",
            "value": 42,
            "unit": "allocs/op",
            "extra": "270784 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchAddenda17Count",
            "value": 4339,
            "unit": "ns/op\t    2248 B/op\t      44 allocs/op",
            "extra": "260264 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchAddenda17Count - ns/op",
            "value": 4339,
            "unit": "ns/op",
            "extra": "260264 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchAddenda17Count - B/op",
            "value": 2248,
            "unit": "B/op",
            "extra": "260264 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchAddenda17Count - allocs/op",
            "value": 44,
            "unit": "allocs/op",
            "extra": "260264 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchAddenda18Count",
            "value": 5274,
            "unit": "ns/op\t    2960 B/op\t      54 allocs/op",
            "extra": "227521 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchAddenda18Count - ns/op",
            "value": 5274,
            "unit": "ns/op",
            "extra": "227521 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchAddenda18Count - B/op",
            "value": 2960,
            "unit": "B/op",
            "extra": "227521 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchAddenda18Count - allocs/op",
            "value": 54,
            "unit": "allocs/op",
            "extra": "227521 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchBuildAddendaError",
            "value": 802.1,
            "unit": "ns/op\t     680 B/op\t      10 allocs/op",
            "extra": "1500004 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchBuildAddendaError - ns/op",
            "value": 802.1,
            "unit": "ns/op",
            "extra": "1500004 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchBuildAddendaError - B/op",
            "value": 680,
            "unit": "B/op",
            "extra": "1500004 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchBuildAddendaError - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "1500004 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchBHODFI",
            "value": 3504,
            "unit": "ns/op\t    1976 B/op\t      34 allocs/op",
            "extra": "330579 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchBHODFI - ns/op",
            "value": 3504,
            "unit": "ns/op",
            "extra": "330579 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchBHODFI - B/op",
            "value": 1976,
            "unit": "B/op",
            "extra": "330579 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchBHODFI - allocs/op",
            "value": 34,
            "unit": "allocs/op",
            "extra": "330579 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchAddenda99Count",
            "value": 3762,
            "unit": "ns/op\t    2048 B/op\t      39 allocs/op",
            "extra": "312993 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchAddenda99Count - ns/op",
            "value": 3762,
            "unit": "ns/op",
            "extra": "312993 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchAddenda99Count - B/op",
            "value": 2048,
            "unit": "B/op",
            "extra": "312993 times\n4 procs"
          },
          {
            "name": "BenchmarkIATBatchAddenda99Count - allocs/op",
            "value": 39,
            "unit": "allocs/op",
            "extra": "312993 times\n4 procs"
          },
          {
            "name": "BenchmarkIATMockEntryDetail",
            "value": 313.6,
            "unit": "ns/op\t     312 B/op\t       3 allocs/op",
            "extra": "3836100 times\n4 procs"
          },
          {
            "name": "BenchmarkIATMockEntryDetail - ns/op",
            "value": 313.6,
            "unit": "ns/op",
            "extra": "3836100 times\n4 procs"
          },
          {
            "name": "BenchmarkIATMockEntryDetail - B/op",
            "value": 312,
            "unit": "B/op",
            "extra": "3836100 times\n4 procs"
          },
          {
            "name": "BenchmarkIATMockEntryDetail - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "3836100 times\n4 procs"
          },
          {
            "name": "BenchmarkParseIATEntryDetail",
            "value": 5991,
            "unit": "ns/op\t   15788 B/op\t      26 allocs/op",
            "extra": "199993 times\n4 procs"
          },
          {
            "name": "BenchmarkParseIATEntryDetail - ns/op",
            "value": 5991,
            "unit": "ns/op",
            "extra": "199993 times\n4 procs"
          },
          {
            "name": "BenchmarkParseIATEntryDetail - B/op",
            "value": 15788,
            "unit": "B/op",
            "extra": "199993 times\n4 procs"
          },
          {
            "name": "BenchmarkParseIATEntryDetail - allocs/op",
            "value": 26,
            "unit": "allocs/op",
            "extra": "199993 times\n4 procs"
          },
          {
            "name": "BenchmarkIATEDString",
            "value": 6327,
            "unit": "ns/op\t   15840 B/op\t      27 allocs/op",
            "extra": "189798 times\n4 procs"
          },
          {
            "name": "BenchmarkIATEDString - ns/op",
            "value": 6327,
            "unit": "ns/op",
            "extra": "189798 times\n4 procs"
          },
          {
            "name": "BenchmarkIATEDString - B/op",
            "value": 15840,
            "unit": "B/op",
            "extra": "189798 times\n4 procs"
          },
          {
            "name": "BenchmarkIATEDString - allocs/op",
            "value": 27,
            "unit": "allocs/op",
            "extra": "189798 times\n4 procs"
          },
          {
            "name": "BenchmarkIATEDInvalidTransactionCode",
            "value": 324.5,
            "unit": "ns/op\t     392 B/op\t       5 allocs/op",
            "extra": "3694142 times\n4 procs"
          },
          {
            "name": "BenchmarkIATEDInvalidTransactionCode - ns/op",
            "value": 324.5,
            "unit": "ns/op",
            "extra": "3694142 times\n4 procs"
          },
          {
            "name": "BenchmarkIATEDInvalidTransactionCode - B/op",
            "value": 392,
            "unit": "B/op",
            "extra": "3694142 times\n4 procs"
          },
          {
            "name": "BenchmarkIATEDInvalidTransactionCode - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "3694142 times\n4 procs"
          },
          {
            "name": "BenchmarkEDIATDFIAccountNumberAlphaNumeric",
            "value": 589.3,
            "unit": "ns/op\t     472 B/op\t       7 allocs/op",
            "extra": "2029227 times\n4 procs"
          },
          {
            "name": "BenchmarkEDIATDFIAccountNumberAlphaNumeric - ns/op",
            "value": 589.3,
            "unit": "ns/op",
            "extra": "2029227 times\n4 procs"
          },
          {
            "name": "BenchmarkEDIATDFIAccountNumberAlphaNumeric - B/op",
            "value": 472,
            "unit": "B/op",
            "extra": "2029227 times\n4 procs"
          },
          {
            "name": "BenchmarkEDIATDFIAccountNumberAlphaNumeric - allocs/op",
            "value": 7,
            "unit": "allocs/op",
            "extra": "2029227 times\n4 procs"
          },
          {
            "name": "BenchmarkEDIATisCheckDigit",
            "value": 707.3,
            "unit": "ns/op\t     512 B/op\t       8 allocs/op",
            "extra": "1693952 times\n4 procs"
          },
          {
            "name": "BenchmarkEDIATisCheckDigit - ns/op",
            "value": 707.3,
            "unit": "ns/op",
            "extra": "1693952 times\n4 procs"
          },
          {
            "name": "BenchmarkEDIATisCheckDigit - B/op",
            "value": 512,
            "unit": "B/op",
            "extra": "1693952 times\n4 procs"
          },
          {
            "name": "BenchmarkEDIATisCheckDigit - allocs/op",
            "value": 8,
            "unit": "allocs/op",
            "extra": "1693952 times\n4 procs"
          },
          {
            "name": "BenchmarkEDIATSetRDFI",
            "value": 35.76,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "33513937 times\n4 procs"
          },
          {
            "name": "BenchmarkEDIATSetRDFI - ns/op",
            "value": 35.76,
            "unit": "ns/op",
            "extra": "33513937 times\n4 procs"
          },
          {
            "name": "BenchmarkEDIATSetRDFI - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "33513937 times\n4 procs"
          },
          {
            "name": "BenchmarkEDIATSetRDFI - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "33513937 times\n4 procs"
          },
          {
            "name": "BenchmarkValidateEDIATCheckDigit",
            "value": 456.6,
            "unit": "ns/op\t     448 B/op\t       7 allocs/op",
            "extra": "2625961 times\n4 procs"
          },
          {
            "name": "BenchmarkValidateEDIATCheckDigit - ns/op",
            "value": 456.6,
            "unit": "ns/op",
            "extra": "2625961 times\n4 procs"
          },
          {
            "name": "BenchmarkValidateEDIATCheckDigit - B/op",
            "value": 448,
            "unit": "B/op",
            "extra": "2625961 times\n4 procs"
          },
          {
            "name": "BenchmarkValidateEDIATCheckDigit - allocs/op",
            "value": 7,
            "unit": "allocs/op",
            "extra": "2625961 times\n4 procs"
          },
          {
            "name": "BenchmarkIATEDTransactionCode",
            "value": 319,
            "unit": "ns/op\t     392 B/op\t       5 allocs/op",
            "extra": "3760237 times\n4 procs"
          },
          {
            "name": "BenchmarkIATEDTransactionCode - ns/op",
            "value": 319,
            "unit": "ns/op",
            "extra": "3760237 times\n4 procs"
          },
          {
            "name": "BenchmarkIATEDTransactionCode - B/op",
            "value": 392,
            "unit": "B/op",
            "extra": "3760237 times\n4 procs"
          },
          {
            "name": "BenchmarkIATEDTransactionCode - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "3760237 times\n4 procs"
          },
          {
            "name": "BenchmarkIATEDRDFIIdentification",
            "value": 334.3,
            "unit": "ns/op\t     392 B/op\t       5 allocs/op",
            "extra": "3598219 times\n4 procs"
          },
          {
            "name": "BenchmarkIATEDRDFIIdentification - ns/op",
            "value": 334.3,
            "unit": "ns/op",
            "extra": "3598219 times\n4 procs"
          },
          {
            "name": "BenchmarkIATEDRDFIIdentification - B/op",
            "value": 392,
            "unit": "B/op",
            "extra": "3598219 times\n4 procs"
          },
          {
            "name": "BenchmarkIATEDRDFIIdentification - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "3598219 times\n4 procs"
          },
          {
            "name": "BenchmarkIATEDAddendaRecords",
            "value": 325.7,
            "unit": "ns/op\t     392 B/op\t       5 allocs/op",
            "extra": "3702969 times\n4 procs"
          },
          {
            "name": "BenchmarkIATEDAddendaRecords - ns/op",
            "value": 325.7,
            "unit": "ns/op",
            "extra": "3702969 times\n4 procs"
          },
          {
            "name": "BenchmarkIATEDAddendaRecords - B/op",
            "value": 392,
            "unit": "B/op",
            "extra": "3702969 times\n4 procs"
          },
          {
            "name": "BenchmarkIATEDAddendaRecords - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "3702969 times\n4 procs"
          },
          {
            "name": "BenchmarkIATEDDFIAccountNumber",
            "value": 288.8,
            "unit": "ns/op\t     376 B/op\t       4 allocs/op",
            "extra": "4171476 times\n4 procs"
          },
          {
            "name": "BenchmarkIATEDDFIAccountNumber - ns/op",
            "value": 288.8,
            "unit": "ns/op",
            "extra": "4171476 times\n4 procs"
          },
          {
            "name": "BenchmarkIATEDDFIAccountNumber - B/op",
            "value": 376,
            "unit": "B/op",
            "extra": "4171476 times\n4 procs"
          },
          {
            "name": "BenchmarkIATEDDFIAccountNumber - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "4171476 times\n4 procs"
          },
          {
            "name": "BenchmarkIATEDTraceNumber",
            "value": 314,
            "unit": "ns/op\t     312 B/op\t       3 allocs/op",
            "extra": "3817599 times\n4 procs"
          },
          {
            "name": "BenchmarkIATEDTraceNumber - ns/op",
            "value": 314,
            "unit": "ns/op",
            "extra": "3817599 times\n4 procs"
          },
          {
            "name": "BenchmarkIATEDTraceNumber - B/op",
            "value": 312,
            "unit": "B/op",
            "extra": "3817599 times\n4 procs"
          },
          {
            "name": "BenchmarkIATEDTraceNumber - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "3817599 times\n4 procs"
          },
          {
            "name": "BenchmarkIATEDAddendaRecordIndicator",
            "value": 321.5,
            "unit": "ns/op\t     392 B/op\t       5 allocs/op",
            "extra": "3705136 times\n4 procs"
          },
          {
            "name": "BenchmarkIATEDAddendaRecordIndicator - ns/op",
            "value": 321.5,
            "unit": "ns/op",
            "extra": "3705136 times\n4 procs"
          },
          {
            "name": "BenchmarkIATEDAddendaRecordIndicator - B/op",
            "value": 392,
            "unit": "B/op",
            "extra": "3705136 times\n4 procs"
          },
          {
            "name": "BenchmarkIATEDAddendaRecordIndicator - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "3705136 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles",
            "value": 297006,
            "unit": "ns/op\t   56823 B/op\t     637 allocs/op",
            "extra": "4128 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles - ns/op",
            "value": 297006,
            "unit": "ns/op",
            "extra": "4128 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles - B/op",
            "value": 56823,
            "unit": "B/op",
            "extra": "4128 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles - allocs/op",
            "value": 637,
            "unit": "allocs/op",
            "extra": "4128 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts",
            "value": 297765,
            "unit": "ns/op\t   56847 B/op\t     637 allocs/op",
            "extra": "4027 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts - ns/op",
            "value": 297765,
            "unit": "ns/op",
            "extra": "4027 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts - B/op",
            "value": 56847,
            "unit": "B/op",
            "extra": "4027 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts - allocs/op",
            "value": 637,
            "unit": "allocs/op",
            "extra": "4027 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir",
            "value": 178467,
            "unit": "ns/op\t   57011 B/op\t     640 allocs/op",
            "extra": "6525 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir - ns/op",
            "value": 178467,
            "unit": "ns/op",
            "extra": "6525 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir - B/op",
            "value": 57011,
            "unit": "B/op",
            "extra": "6525 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir - allocs/op",
            "value": 640,
            "unit": "allocs/op",
            "extra": "6525 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts",
            "value": 169223,
            "unit": "ns/op\t   57049 B/op\t     640 allocs/op",
            "extra": "6600 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts - ns/op",
            "value": 169223,
            "unit": "ns/op",
            "extra": "6600 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts - B/op",
            "value": 57049,
            "unit": "B/op",
            "extra": "6600 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts - allocs/op",
            "value": 640,
            "unit": "allocs/op",
            "extra": "6600 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups",
            "value": 367036,
            "unit": "ns/op\t   62546 B/op\t     697 allocs/op",
            "extra": "3368 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups - ns/op",
            "value": 367036,
            "unit": "ns/op",
            "extra": "3368 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups - B/op",
            "value": 62546,
            "unit": "B/op",
            "extra": "3368 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups - allocs/op",
            "value": 697,
            "unit": "allocs/op",
            "extra": "3368 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups",
            "value": 360076,
            "unit": "ns/op\t   62526 B/op\t     697 allocs/op",
            "extra": "3525 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups - ns/op",
            "value": 360076,
            "unit": "ns/op",
            "extra": "3525 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups - B/op",
            "value": 62526,
            "unit": "B/op",
            "extra": "3525 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups - allocs/op",
            "value": 697,
            "unit": "allocs/op",
            "extra": "3525 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups",
            "value": 350578,
            "unit": "ns/op\t   62534 B/op\t     697 allocs/op",
            "extra": "3459 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups - ns/op",
            "value": 350578,
            "unit": "ns/op",
            "extra": "3459 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups - B/op",
            "value": 62534,
            "unit": "B/op",
            "extra": "3459 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups - allocs/op",
            "value": 697,
            "unit": "allocs/op",
            "extra": "3459 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups",
            "value": 351320,
            "unit": "ns/op\t   62535 B/op\t     697 allocs/op",
            "extra": "4069 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups - ns/op",
            "value": 351320,
            "unit": "ns/op",
            "extra": "4069 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups - B/op",
            "value": 62535,
            "unit": "B/op",
            "extra": "4069 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups - allocs/op",
            "value": 697,
            "unit": "allocs/op",
            "extra": "4069 times\n4 procs"
          },
          {
            "name": "BenchmarkLineCount",
            "value": 15343,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "78660 times\n4 procs"
          },
          {
            "name": "BenchmarkLineCount - ns/op",
            "value": 15343,
            "unit": "ns/op",
            "extra": "78660 times\n4 procs"
          },
          {
            "name": "BenchmarkLineCount - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "78660 times\n4 procs"
          },
          {
            "name": "BenchmarkLineCount - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "78660 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead",
            "value": 40954,
            "unit": "ns/op\t   21523 B/op\t      61 allocs/op",
            "extra": "29415 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead - ns/op",
            "value": 40954,
            "unit": "ns/op",
            "extra": "29415 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead - B/op",
            "value": 21523,
            "unit": "B/op",
            "extra": "29415 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead - allocs/op",
            "value": 61,
            "unit": "allocs/op",
            "extra": "29415 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead",
            "value": 74752,
            "unit": "ns/op\t   25388 B/op\t     136 allocs/op",
            "extra": "16048 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead - ns/op",
            "value": 74752,
            "unit": "ns/op",
            "extra": "16048 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead - B/op",
            "value": 25388,
            "unit": "B/op",
            "extra": "16048 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead - allocs/op",
            "value": 136,
            "unit": "allocs/op",
            "extra": "16048 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead",
            "value": 31008,
            "unit": "ns/op\t   20931 B/op\t      54 allocs/op",
            "extra": "38606 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead - ns/op",
            "value": 31008,
            "unit": "ns/op",
            "extra": "38606 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead - B/op",
            "value": 20931,
            "unit": "B/op",
            "extra": "38606 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead - allocs/op",
            "value": 54,
            "unit": "allocs/op",
            "extra": "38606 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordTypeUnknown",
            "value": 8152,
            "unit": "ns/op\t   19170 B/op\t      25 allocs/op",
            "extra": "144574 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordTypeUnknown - ns/op",
            "value": 8152,
            "unit": "ns/op",
            "extra": "144574 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordTypeUnknown - B/op",
            "value": 19170,
            "unit": "B/op",
            "extra": "144574 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordTypeUnknown - allocs/op",
            "value": 25,
            "unit": "allocs/op",
            "extra": "144574 times\n4 procs"
          },
          {
            "name": "BenchmarkTwoFileHeaders",
            "value": 11396,
            "unit": "ns/op\t   19776 B/op\t      30 allocs/op",
            "extra": "105258 times\n4 procs"
          },
          {
            "name": "BenchmarkTwoFileHeaders - ns/op",
            "value": 11396,
            "unit": "ns/op",
            "extra": "105258 times\n4 procs"
          },
          {
            "name": "BenchmarkTwoFileHeaders - B/op",
            "value": 19776,
            "unit": "B/op",
            "extra": "105258 times\n4 procs"
          },
          {
            "name": "BenchmarkTwoFileHeaders - allocs/op",
            "value": 30,
            "unit": "allocs/op",
            "extra": "105258 times\n4 procs"
          },
          {
            "name": "BenchmarkTwoFileControls",
            "value": 10521,
            "unit": "ns/op\t   19956 B/op\t      29 allocs/op",
            "extra": "114036 times\n4 procs"
          },
          {
            "name": "BenchmarkTwoFileControls - ns/op",
            "value": 10521,
            "unit": "ns/op",
            "extra": "114036 times\n4 procs"
          },
          {
            "name": "BenchmarkTwoFileControls - B/op",
            "value": 19956,
            "unit": "B/op",
            "extra": "114036 times\n4 procs"
          },
          {
            "name": "BenchmarkTwoFileControls - allocs/op",
            "value": 29,
            "unit": "allocs/op",
            "extra": "114036 times\n4 procs"
          },
          {
            "name": "BenchmarkFileLineEmpty",
            "value": 1986,
            "unit": "ns/op\t    6089 B/op\t       9 allocs/op",
            "extra": "583537 times\n4 procs"
          },
          {
            "name": "BenchmarkFileLineEmpty - ns/op",
            "value": 1986,
            "unit": "ns/op",
            "extra": "583537 times\n4 procs"
          },
          {
            "name": "BenchmarkFileLineEmpty - B/op",
            "value": 6089,
            "unit": "B/op",
            "extra": "583537 times\n4 procs"
          },
          {
            "name": "BenchmarkFileLineEmpty - allocs/op",
            "value": 9,
            "unit": "allocs/op",
            "extra": "583537 times\n4 procs"
          },
          {
            "name": "BenchmarkFileLineShort",
            "value": 9420,
            "unit": "ns/op\t   20016 B/op\t      39 allocs/op",
            "extra": "127582 times\n4 procs"
          },
          {
            "name": "BenchmarkFileLineShort - ns/op",
            "value": 9420,
            "unit": "ns/op",
            "extra": "127582 times\n4 procs"
          },
          {
            "name": "BenchmarkFileLineShort - B/op",
            "value": 20016,
            "unit": "B/op",
            "extra": "127582 times\n4 procs"
          },
          {
            "name": "BenchmarkFileLineShort - allocs/op",
            "value": 39,
            "unit": "allocs/op",
            "extra": "127582 times\n4 procs"
          },
          {
            "name": "BenchmarkFileLineLong",
            "value": 12657,
            "unit": "ns/op\t   20647 B/op\t      52 allocs/op",
            "extra": "93566 times\n4 procs"
          },
          {
            "name": "BenchmarkFileLineLong - ns/op",
            "value": 12657,
            "unit": "ns/op",
            "extra": "93566 times\n4 procs"
          },
          {
            "name": "BenchmarkFileLineLong - B/op",
            "value": 20647,
            "unit": "B/op",
            "extra": "93566 times\n4 procs"
          },
          {
            "name": "BenchmarkFileLineLong - allocs/op",
            "value": 52,
            "unit": "allocs/op",
            "extra": "93566 times\n4 procs"
          },
          {
            "name": "BenchmarkFileFileHeaderErr",
            "value": 10232,
            "unit": "ns/op\t   19761 B/op\t      36 allocs/op",
            "extra": "115903 times\n4 procs"
          },
          {
            "name": "BenchmarkFileFileHeaderErr - ns/op",
            "value": 10232,
            "unit": "ns/op",
            "extra": "115903 times\n4 procs"
          },
          {
            "name": "BenchmarkFileFileHeaderErr - B/op",
            "value": 19761,
            "unit": "B/op",
            "extra": "115903 times\n4 procs"
          },
          {
            "name": "BenchmarkFileFileHeaderErr - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "115903 times\n4 procs"
          },
          {
            "name": "BenchmarkFileBatchHeaderErr",
            "value": 10162,
            "unit": "ns/op\t   19796 B/op\t      41 allocs/op",
            "extra": "119276 times\n4 procs"
          },
          {
            "name": "BenchmarkFileBatchHeaderErr - ns/op",
            "value": 10162,
            "unit": "ns/op",
            "extra": "119276 times\n4 procs"
          },
          {
            "name": "BenchmarkFileBatchHeaderErr - B/op",
            "value": 19796,
            "unit": "B/op",
            "extra": "119276 times\n4 procs"
          },
          {
            "name": "BenchmarkFileBatchHeaderErr - allocs/op",
            "value": 41,
            "unit": "allocs/op",
            "extra": "119276 times\n4 procs"
          },
          {
            "name": "BenchmarkFileBatchHeaderDuplicate",
            "value": 9062,
            "unit": "ns/op\t   19649 B/op\t      30 allocs/op",
            "extra": "134526 times\n4 procs"
          },
          {
            "name": "BenchmarkFileBatchHeaderDuplicate - ns/op",
            "value": 9062,
            "unit": "ns/op",
            "extra": "134526 times\n4 procs"
          },
          {
            "name": "BenchmarkFileBatchHeaderDuplicate - B/op",
            "value": 19649,
            "unit": "B/op",
            "extra": "134526 times\n4 procs"
          },
          {
            "name": "BenchmarkFileBatchHeaderDuplicate - allocs/op",
            "value": 30,
            "unit": "allocs/op",
            "extra": "134526 times\n4 procs"
          },
          {
            "name": "BenchmarkFileEntryDetailOutsideBatch",
            "value": 8744,
            "unit": "ns/op\t   19497 B/op\t      29 allocs/op",
            "extra": "137850 times\n4 procs"
          },
          {
            "name": "BenchmarkFileEntryDetailOutsideBatch - ns/op",
            "value": 8744,
            "unit": "ns/op",
            "extra": "137850 times\n4 procs"
          },
          {
            "name": "BenchmarkFileEntryDetailOutsideBatch - B/op",
            "value": 19497,
            "unit": "B/op",
            "extra": "137850 times\n4 procs"
          },
          {
            "name": "BenchmarkFileEntryDetailOutsideBatch - allocs/op",
            "value": 29,
            "unit": "allocs/op",
            "extra": "137850 times\n4 procs"
          },
          {
            "name": "BenchmarkFileEntryDetail",
            "value": 11119,
            "unit": "ns/op\t   20958 B/op\t      51 allocs/op",
            "extra": "108836 times\n4 procs"
          },
          {
            "name": "BenchmarkFileEntryDetail - ns/op",
            "value": 11119,
            "unit": "ns/op",
            "extra": "108836 times\n4 procs"
          },
          {
            "name": "BenchmarkFileEntryDetail - B/op",
            "value": 20958,
            "unit": "B/op",
            "extra": "108836 times\n4 procs"
          },
          {
            "name": "BenchmarkFileEntryDetail - allocs/op",
            "value": 51,
            "unit": "allocs/op",
            "extra": "108836 times\n4 procs"
          },
          {
            "name": "BenchmarkFileAddenda05",
            "value": 17372,
            "unit": "ns/op\t   21572 B/op\t      70 allocs/op",
            "extra": "69406 times\n4 procs"
          },
          {
            "name": "BenchmarkFileAddenda05 - ns/op",
            "value": 17372,
            "unit": "ns/op",
            "extra": "69406 times\n4 procs"
          },
          {
            "name": "BenchmarkFileAddenda05 - B/op",
            "value": 21572,
            "unit": "B/op",
            "extra": "69406 times\n4 procs"
          },
          {
            "name": "BenchmarkFileAddenda05 - allocs/op",
            "value": 70,
            "unit": "allocs/op",
            "extra": "69406 times\n4 procs"
          },
          {
            "name": "BenchmarkFileAddenda02invalid",
            "value": 17540,
            "unit": "ns/op\t   21660 B/op\t      70 allocs/op",
            "extra": "69460 times\n4 procs"
          },
          {
            "name": "BenchmarkFileAddenda02invalid - ns/op",
            "value": 17540,
            "unit": "ns/op",
            "extra": "69460 times\n4 procs"
          },
          {
            "name": "BenchmarkFileAddenda02invalid - B/op",
            "value": 21660,
            "unit": "B/op",
            "extra": "69460 times\n4 procs"
          },
          {
            "name": "BenchmarkFileAddenda02invalid - allocs/op",
            "value": 70,
            "unit": "allocs/op",
            "extra": "69460 times\n4 procs"
          },
          {
            "name": "BenchmarkFileAddenda02",
            "value": 17646,
            "unit": "ns/op\t   21661 B/op\t      70 allocs/op",
            "extra": "69205 times\n4 procs"
          },
          {
            "name": "BenchmarkFileAddenda02 - ns/op",
            "value": 17646,
            "unit": "ns/op",
            "extra": "69205 times\n4 procs"
          },
          {
            "name": "BenchmarkFileAddenda02 - B/op",
            "value": 21661,
            "unit": "B/op",
            "extra": "69205 times\n4 procs"
          },
          {
            "name": "BenchmarkFileAddenda02 - allocs/op",
            "value": 70,
            "unit": "allocs/op",
            "extra": "69205 times\n4 procs"
          },
          {
            "name": "BenchmarkFileAddenda98invalid",
            "value": 17907,
            "unit": "ns/op\t   21893 B/op\t      70 allocs/op",
            "extra": "66826 times\n4 procs"
          },
          {
            "name": "BenchmarkFileAddenda98invalid - ns/op",
            "value": 17907,
            "unit": "ns/op",
            "extra": "66826 times\n4 procs"
          },
          {
            "name": "BenchmarkFileAddenda98invalid - B/op",
            "value": 21893,
            "unit": "B/op",
            "extra": "66826 times\n4 procs"
          },
          {
            "name": "BenchmarkFileAddenda98invalid - allocs/op",
            "value": 70,
            "unit": "allocs/op",
            "extra": "66826 times\n4 procs"
          },
          {
            "name": "BenchmarkFileAddenda98",
            "value": 17369,
            "unit": "ns/op\t   21612 B/op\t      70 allocs/op",
            "extra": "68923 times\n4 procs"
          },
          {
            "name": "BenchmarkFileAddenda98 - ns/op",
            "value": 17369,
            "unit": "ns/op",
            "extra": "68923 times\n4 procs"
          },
          {
            "name": "BenchmarkFileAddenda98 - B/op",
            "value": 21612,
            "unit": "B/op",
            "extra": "68923 times\n4 procs"
          },
          {
            "name": "BenchmarkFileAddenda98 - allocs/op",
            "value": 70,
            "unit": "allocs/op",
            "extra": "68923 times\n4 procs"
          },
          {
            "name": "BenchmarkFileAddenda99invalid",
            "value": 18117,
            "unit": "ns/op\t   21909 B/op\t      70 allocs/op",
            "extra": "66206 times\n4 procs"
          },
          {
            "name": "BenchmarkFileAddenda99invalid - ns/op",
            "value": 18117,
            "unit": "ns/op",
            "extra": "66206 times\n4 procs"
          },
          {
            "name": "BenchmarkFileAddenda99invalid - B/op",
            "value": 21909,
            "unit": "B/op",
            "extra": "66206 times\n4 procs"
          },
          {
            "name": "BenchmarkFileAddenda99invalid - allocs/op",
            "value": 70,
            "unit": "allocs/op",
            "extra": "66206 times\n4 procs"
          },
          {
            "name": "BenchmarkFileAddenda99",
            "value": 17459,
            "unit": "ns/op\t   21628 B/op\t      70 allocs/op",
            "extra": "68564 times\n4 procs"
          },
          {
            "name": "BenchmarkFileAddenda99 - ns/op",
            "value": 17459,
            "unit": "ns/op",
            "extra": "68564 times\n4 procs"
          },
          {
            "name": "BenchmarkFileAddenda99 - B/op",
            "value": 21628,
            "unit": "B/op",
            "extra": "68564 times\n4 procs"
          },
          {
            "name": "BenchmarkFileAddenda99 - allocs/op",
            "value": 70,
            "unit": "allocs/op",
            "extra": "68564 times\n4 procs"
          },
          {
            "name": "BenchmarkFileAddendaOutsideBatch",
            "value": 11134,
            "unit": "ns/op\t   19969 B/op\t      35 allocs/op",
            "extra": "109471 times\n4 procs"
          },
          {
            "name": "BenchmarkFileAddendaOutsideBatch - ns/op",
            "value": 11134,
            "unit": "ns/op",
            "extra": "109471 times\n4 procs"
          },
          {
            "name": "BenchmarkFileAddendaOutsideBatch - B/op",
            "value": 19969,
            "unit": "B/op",
            "extra": "109471 times\n4 procs"
          },
          {
            "name": "BenchmarkFileAddendaOutsideBatch - allocs/op",
            "value": 35,
            "unit": "allocs/op",
            "extra": "109471 times\n4 procs"
          },
          {
            "name": "BenchmarkFileAddendaNoIndicator",
            "value": 17351,
            "unit": "ns/op\t   21484 B/op\t      68 allocs/op",
            "extra": "69733 times\n4 procs"
          },
          {
            "name": "BenchmarkFileAddendaNoIndicator - ns/op",
            "value": 17351,
            "unit": "ns/op",
            "extra": "69733 times\n4 procs"
          },
          {
            "name": "BenchmarkFileAddendaNoIndicator - B/op",
            "value": 21484,
            "unit": "B/op",
            "extra": "69733 times\n4 procs"
          },
          {
            "name": "BenchmarkFileAddendaNoIndicator - allocs/op",
            "value": 68,
            "unit": "allocs/op",
            "extra": "69733 times\n4 procs"
          },
          {
            "name": "BenchmarkFileFileControlErr",
            "value": 9260,
            "unit": "ns/op\t   19335 B/op\t      33 allocs/op",
            "extra": "122236 times\n4 procs"
          },
          {
            "name": "BenchmarkFileFileControlErr - ns/op",
            "value": 9260,
            "unit": "ns/op",
            "extra": "122236 times\n4 procs"
          },
          {
            "name": "BenchmarkFileFileControlErr - B/op",
            "value": 19335,
            "unit": "B/op",
            "extra": "122236 times\n4 procs"
          },
          {
            "name": "BenchmarkFileFileControlErr - allocs/op",
            "value": 33,
            "unit": "allocs/op",
            "extra": "122236 times\n4 procs"
          },
          {
            "name": "BenchmarkFileBatchHeaderSEC",
            "value": 10131,
            "unit": "ns/op\t   19716 B/op\t      42 allocs/op",
            "extra": "119066 times\n4 procs"
          },
          {
            "name": "BenchmarkFileBatchHeaderSEC - ns/op",
            "value": 10131,
            "unit": "ns/op",
            "extra": "119066 times\n4 procs"
          },
          {
            "name": "BenchmarkFileBatchHeaderSEC - B/op",
            "value": 19716,
            "unit": "B/op",
            "extra": "119066 times\n4 procs"
          },
          {
            "name": "BenchmarkFileBatchHeaderSEC - allocs/op",
            "value": 42,
            "unit": "allocs/op",
            "extra": "119066 times\n4 procs"
          },
          {
            "name": "BenchmarkFileBatchControlNoCurrentBatch",
            "value": 8416,
            "unit": "ns/op\t   19217 B/op\t      28 allocs/op",
            "extra": "134772 times\n4 procs"
          },
          {
            "name": "BenchmarkFileBatchControlNoCurrentBatch - ns/op",
            "value": 8416,
            "unit": "ns/op",
            "extra": "134772 times\n4 procs"
          },
          {
            "name": "BenchmarkFileBatchControlNoCurrentBatch - B/op",
            "value": 19217,
            "unit": "B/op",
            "extra": "134772 times\n4 procs"
          },
          {
            "name": "BenchmarkFileBatchControlNoCurrentBatch - allocs/op",
            "value": 28,
            "unit": "allocs/op",
            "extra": "134772 times\n4 procs"
          },
          {
            "name": "BenchmarkFileBatchControlValidate",
            "value": 16286,
            "unit": "ns/op\t   13343 B/op\t      72 allocs/op",
            "extra": "73809 times\n4 procs"
          },
          {
            "name": "BenchmarkFileBatchControlValidate - ns/op",
            "value": 16286,
            "unit": "ns/op",
            "extra": "73809 times\n4 procs"
          },
          {
            "name": "BenchmarkFileBatchControlValidate - B/op",
            "value": 13343,
            "unit": "B/op",
            "extra": "73809 times\n4 procs"
          },
          {
            "name": "BenchmarkFileBatchControlValidate - allocs/op",
            "value": 72,
            "unit": "allocs/op",
            "extra": "73809 times\n4 procs"
          },
          {
            "name": "BenchmarkFileAddBatchValidation",
            "value": 18632,
            "unit": "ns/op\t   21645 B/op\t      75 allocs/op",
            "extra": "64598 times\n4 procs"
          },
          {
            "name": "BenchmarkFileAddBatchValidation - ns/op",
            "value": 18632,
            "unit": "ns/op",
            "extra": "64598 times\n4 procs"
          },
          {
            "name": "BenchmarkFileAddBatchValidation - B/op",
            "value": 21645,
            "unit": "B/op",
            "extra": "64598 times\n4 procs"
          },
          {
            "name": "BenchmarkFileAddBatchValidation - allocs/op",
            "value": 75,
            "unit": "allocs/op",
            "extra": "64598 times\n4 procs"
          },
          {
            "name": "BenchmarkFileLongErr",
            "value": 12943,
            "unit": "ns/op\t   20011 B/op\t      42 allocs/op",
            "extra": "91989 times\n4 procs"
          },
          {
            "name": "BenchmarkFileLongErr - ns/op",
            "value": 12943,
            "unit": "ns/op",
            "extra": "91989 times\n4 procs"
          },
          {
            "name": "BenchmarkFileLongErr - B/op",
            "value": 20011,
            "unit": "B/op",
            "extra": "91989 times\n4 procs"
          },
          {
            "name": "BenchmarkFileLongErr - allocs/op",
            "value": 42,
            "unit": "allocs/op",
            "extra": "91989 times\n4 procs"
          },
          {
            "name": "BenchmarkFileAddendaOutsideEntry",
            "value": 12769,
            "unit": "ns/op\t   20340 B/op\t      48 allocs/op",
            "extra": "93774 times\n4 procs"
          },
          {
            "name": "BenchmarkFileAddendaOutsideEntry - ns/op",
            "value": 12769,
            "unit": "ns/op",
            "extra": "93774 times\n4 procs"
          },
          {
            "name": "BenchmarkFileAddendaOutsideEntry - B/op",
            "value": 20340,
            "unit": "B/op",
            "extra": "93774 times\n4 procs"
          },
          {
            "name": "BenchmarkFileAddendaOutsideEntry - allocs/op",
            "value": 48,
            "unit": "allocs/op",
            "extra": "93774 times\n4 procs"
          },
          {
            "name": "BenchmarkFileFHImmediateOrigin",
            "value": 10287,
            "unit": "ns/op\t   19761 B/op\t      36 allocs/op",
            "extra": "116112 times\n4 procs"
          },
          {
            "name": "BenchmarkFileFHImmediateOrigin - ns/op",
            "value": 10287,
            "unit": "ns/op",
            "extra": "116112 times\n4 procs"
          },
          {
            "name": "BenchmarkFileFHImmediateOrigin - B/op",
            "value": 19761,
            "unit": "B/op",
            "extra": "116112 times\n4 procs"
          },
          {
            "name": "BenchmarkFileFHImmediateOrigin - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "116112 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead",
            "value": 303414,
            "unit": "ns/op\t   56273 B/op\t     743 allocs/op",
            "extra": "3816 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead - ns/op",
            "value": 303414,
            "unit": "ns/op",
            "extra": "3816 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead - B/op",
            "value": 56273,
            "unit": "B/op",
            "extra": "3816 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead - allocs/op",
            "value": 743,
            "unit": "allocs/op",
            "extra": "3816 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2",
            "value": 304610,
            "unit": "ns/op\t   56274 B/op\t     743 allocs/op",
            "extra": "3993 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2 - ns/op",
            "value": 304610,
            "unit": "ns/op",
            "extra": "3993 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2 - B/op",
            "value": 56274,
            "unit": "B/op",
            "extra": "3993 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2 - allocs/op",
            "value": 743,
            "unit": "allocs/op",
            "extra": "3993 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3",
            "value": 99410,
            "unit": "ns/op\t   27428 B/op\t     199 allocs/op",
            "extra": "12128 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3 - ns/op",
            "value": 99410,
            "unit": "ns/op",
            "extra": "12128 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3 - B/op",
            "value": 27428,
            "unit": "B/op",
            "extra": "12128 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3 - allocs/op",
            "value": 199,
            "unit": "allocs/op",
            "extra": "12128 times\n4 procs"
          },
          {
            "name": "BenchmarkACHIATAddenda17",
            "value": 103941,
            "unit": "ns/op\t   27925 B/op\t     213 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkACHIATAddenda17 - ns/op",
            "value": 103941,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkACHIATAddenda17 - B/op",
            "value": 27925,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkACHIATAddenda17 - allocs/op",
            "value": 213,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkACHIATAddenda1718",
            "value": 136596,
            "unit": "ns/op\t   31719 B/op\t     306 allocs/op",
            "extra": "8164 times\n4 procs"
          },
          {
            "name": "BenchmarkACHIATAddenda1718 - ns/op",
            "value": 136596,
            "unit": "ns/op",
            "extra": "8164 times\n4 procs"
          },
          {
            "name": "BenchmarkACHIATAddenda1718 - B/op",
            "value": 31719,
            "unit": "B/op",
            "extra": "8164 times\n4 procs"
          },
          {
            "name": "BenchmarkACHIATAddenda1718 - allocs/op",
            "value": 306,
            "unit": "allocs/op",
            "extra": "8164 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileIATBatchHeader",
            "value": 89222,
            "unit": "ns/op\t   26171 B/op\t     158 allocs/op",
            "extra": "13508 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileIATBatchHeader - ns/op",
            "value": 89222,
            "unit": "ns/op",
            "extra": "13508 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileIATBatchHeader - B/op",
            "value": 26171,
            "unit": "B/op",
            "extra": "13508 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileIATBatchHeader - allocs/op",
            "value": 158,
            "unit": "allocs/op",
            "extra": "13508 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileIATEntryDetail",
            "value": 92956,
            "unit": "ns/op\t   26804 B/op\t     170 allocs/op",
            "extra": "13024 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileIATEntryDetail - ns/op",
            "value": 92956,
            "unit": "ns/op",
            "extra": "13024 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileIATEntryDetail - B/op",
            "value": 26804,
            "unit": "B/op",
            "extra": "13024 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileIATEntryDetail - allocs/op",
            "value": 170,
            "unit": "allocs/op",
            "extra": "13024 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileIATBC",
            "value": 100337,
            "unit": "ns/op\t   27982 B/op\t     207 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileIATBC - ns/op",
            "value": 100337,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileIATBC - B/op",
            "value": 27982,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileIATBC - allocs/op",
            "value": 207,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileIATBH",
            "value": 94950,
            "unit": "ns/op\t   27189 B/op\t     183 allocs/op",
            "extra": "12663 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileIATBH - ns/op",
            "value": 94950,
            "unit": "ns/op",
            "extra": "12663 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileIATBH - B/op",
            "value": 27189,
            "unit": "B/op",
            "extra": "12663 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileIATBH - allocs/op",
            "value": 183,
            "unit": "allocs/op",
            "extra": "12663 times\n4 procs"
          },
          {
            "name": "BenchmarkFileRecord",
            "value": 326.5,
            "unit": "ns/op\t       8 B/op\t       1 allocs/op",
            "extra": "3665876 times\n4 procs"
          },
          {
            "name": "BenchmarkFileRecord - ns/op",
            "value": 326.5,
            "unit": "ns/op",
            "extra": "3665876 times\n4 procs"
          },
          {
            "name": "BenchmarkFileRecord - B/op",
            "value": 8,
            "unit": "B/op",
            "extra": "3665876 times\n4 procs"
          },
          {
            "name": "BenchmarkFileRecord - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "3665876 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchRecord",
            "value": 429.3,
            "unit": "ns/op\t     456 B/op\t       4 allocs/op",
            "extra": "2808530 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchRecord - ns/op",
            "value": 429.3,
            "unit": "ns/op",
            "extra": "2808530 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchRecord - B/op",
            "value": 456,
            "unit": "B/op",
            "extra": "2808530 times\n4 procs"
          },
          {
            "name": "BenchmarkBatchRecord - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "2808530 times\n4 procs"
          },
          {
            "name": "BenchmarkEntryDetail",
            "value": 354.8,
            "unit": "ns/op\t     280 B/op\t       3 allocs/op",
            "extra": "3381554 times\n4 procs"
          },
          {
            "name": "BenchmarkEntryDetail - ns/op",
            "value": 354.8,
            "unit": "ns/op",
            "extra": "3381554 times\n4 procs"
          },
          {
            "name": "BenchmarkEntryDetail - B/op",
            "value": 280,
            "unit": "B/op",
            "extra": "3381554 times\n4 procs"
          },
          {
            "name": "BenchmarkEntryDetail - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "3381554 times\n4 procs"
          },
          {
            "name": "BenchmarkEntryDetailPaymentType",
            "value": 357,
            "unit": "ns/op\t     280 B/op\t       3 allocs/op",
            "extra": "3382722 times\n4 procs"
          },
          {
            "name": "BenchmarkEntryDetailPaymentType - ns/op",
            "value": 357,
            "unit": "ns/op",
            "extra": "3382722 times\n4 procs"
          },
          {
            "name": "BenchmarkEntryDetailPaymentType - B/op",
            "value": 280,
            "unit": "B/op",
            "extra": "3382722 times\n4 procs"
          },
          {
            "name": "BenchmarkEntryDetailPaymentType - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "3382722 times\n4 procs"
          },
          {
            "name": "BenchmarkEntryDetailReceivingCompany",
            "value": 364.5,
            "unit": "ns/op\t     280 B/op\t       3 allocs/op",
            "extra": "3263918 times\n4 procs"
          },
          {
            "name": "BenchmarkEntryDetailReceivingCompany - ns/op",
            "value": 364.5,
            "unit": "ns/op",
            "extra": "3263918 times\n4 procs"
          },
          {
            "name": "BenchmarkEntryDetailReceivingCompany - B/op",
            "value": 280,
            "unit": "B/op",
            "extra": "3263918 times\n4 procs"
          },
          {
            "name": "BenchmarkEntryDetailReceivingCompany - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "3263918 times\n4 procs"
          },
          {
            "name": "BenchmarkAddendaRecord",
            "value": 73.13,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "17003000 times\n4 procs"
          },
          {
            "name": "BenchmarkAddendaRecord - ns/op",
            "value": 73.13,
            "unit": "ns/op",
            "extra": "17003000 times\n4 procs"
          },
          {
            "name": "BenchmarkAddendaRecord - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "17003000 times\n4 procs"
          },
          {
            "name": "BenchmarkAddendaRecord - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "17003000 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile",
            "value": 12937,
            "unit": "ns/op\t    9668 B/op\t      99 allocs/op",
            "extra": "90888 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile - ns/op",
            "value": 12937,
            "unit": "ns/op",
            "extra": "90888 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile - B/op",
            "value": 9668,
            "unit": "B/op",
            "extra": "90888 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile - allocs/op",
            "value": 99,
            "unit": "allocs/op",
            "extra": "90888 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid",
            "value": 33.29,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "35556913 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid - ns/op",
            "value": 33.29,
            "unit": "ns/op",
            "extra": "35556913 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "35556913 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "35556913 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite",
            "value": 46281,
            "unit": "ns/op\t   31592 B/op\t     130 allocs/op",
            "extra": "25153 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite - ns/op",
            "value": 46281,
            "unit": "ns/op",
            "extra": "25153 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite - B/op",
            "value": 31592,
            "unit": "B/op",
            "extra": "25153 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite - allocs/op",
            "value": 130,
            "unit": "allocs/op",
            "extra": "25153 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite",
            "value": 229656,
            "unit": "ns/op\t   53917 B/op\t    2041 allocs/op",
            "extra": "5407 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite - ns/op",
            "value": 229656,
            "unit": "ns/op",
            "extra": "5407 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite - B/op",
            "value": 53917,
            "unit": "B/op",
            "extra": "5407 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite - allocs/op",
            "value": 2041,
            "unit": "allocs/op",
            "extra": "5407 times\n4 procs"
          },
          {
            "name": "BenchmarkFileWriteErr",
            "value": 5935,
            "unit": "ns/op\t    6145 B/op\t      25 allocs/op",
            "extra": "197046 times\n4 procs"
          },
          {
            "name": "BenchmarkFileWriteErr - ns/op",
            "value": 5935,
            "unit": "ns/op",
            "extra": "197046 times\n4 procs"
          },
          {
            "name": "BenchmarkFileWriteErr - B/op",
            "value": 6145,
            "unit": "B/op",
            "extra": "197046 times\n4 procs"
          },
          {
            "name": "BenchmarkFileWriteErr - allocs/op",
            "value": 25,
            "unit": "allocs/op",
            "extra": "197046 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite",
            "value": 155820,
            "unit": "ns/op\t   57108 B/op\t     612 allocs/op",
            "extra": "7460 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite - ns/op",
            "value": 155820,
            "unit": "ns/op",
            "extra": "7460 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite - B/op",
            "value": 57108,
            "unit": "B/op",
            "extra": "7460 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite - allocs/op",
            "value": 612,
            "unit": "allocs/op",
            "extra": "7460 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDIATWrite",
            "value": 231380104,
            "unit": "1210428822410032113A094101Federal Reserve Bank   My Bank Name                   ",
            "extra": "101 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDIATWrite - 1210428822410032113A094101Federal",
            "value": 231380104,
            "unit": "1210428822410032113A094101Federal",
            "extra": "101 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDIATWrite - Bank",
            "value": null,
            "unit": "Bank",
            "extra": "101 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDIATWrite - Bank",
            "value": null,
            "unit": "Bank",
            "extra": "101 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDIATWrite - ",
            "value": null,
            "unit": "",
            "extra": "101 times\n4 procs"
          },
          {
            "name": "BenchmarkIATReturn",
            "value": 231380104,
            "unit": "1210428822410032113A094101Federal Reserve Bank   My Bank Name                   ",
            "extra": "101 times\n4 procs"
          },
          {
            "name": "BenchmarkIATReturn - 1210428822410032113A094101Federal",
            "value": 231380104,
            "unit": "1210428822410032113A094101Federal",
            "extra": "101 times\n4 procs"
          },
          {
            "name": "BenchmarkIATReturn - Bank",
            "value": null,
            "unit": "Bank",
            "extra": "101 times\n4 procs"
          },
          {
            "name": "BenchmarkIATReturn - Bank",
            "value": null,
            "unit": "Bank",
            "extra": "101 times\n4 procs"
          },
          {
            "name": "BenchmarkIATReturn - ",
            "value": null,
            "unit": "",
            "extra": "101 times\n4 procs"
          },
          {
            "name": "Benchmark__FlattenBigFile",
            "value": 5031444916,
            "unit": "ns/op\t3210912560 B/op\t 2019845 allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark__FlattenBigFile - ns/op",
            "value": 5031444916,
            "unit": "ns/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark__FlattenBigFile - B/op",
            "value": 3210912560,
            "unit": "B/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark__FlattenBigFile - allocs/op",
            "value": 2019845,
            "unit": "allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "BenchmarkParsing/Read_../testdata/20110805A.ach",
            "value": 333311,
            "unit": "ns/op\t   56026 B/op\t     737 allocs/op",
            "extra": "3364 times\n4 procs"
          },
          {
            "name": "BenchmarkParsing/Read_../testdata/20110805A.ach - ns/op",
            "value": 333311,
            "unit": "ns/op",
            "extra": "3364 times\n4 procs"
          },
          {
            "name": "BenchmarkParsing/Read_../testdata/20110805A.ach - B/op",
            "value": 56026,
            "unit": "B/op",
            "extra": "3364 times\n4 procs"
          },
          {
            "name": "BenchmarkParsing/Read_../testdata/20110805A.ach - allocs/op",
            "value": 737,
            "unit": "allocs/op",
            "extra": "3364 times\n4 procs"
          },
          {
            "name": "BenchmarkParsing/Iterator_../testdata/20110805A.ach",
            "value": 166805,
            "unit": "ns/op\t   43081 B/op\t     725 allocs/op",
            "extra": "8193 times\n4 procs"
          },
          {
            "name": "BenchmarkParsing/Iterator_../testdata/20110805A.ach - ns/op",
            "value": 166805,
            "unit": "ns/op",
            "extra": "8193 times\n4 procs"
          },
          {
            "name": "BenchmarkParsing/Iterator_../testdata/20110805A.ach - B/op",
            "value": 43081,
            "unit": "B/op",
            "extra": "8193 times\n4 procs"
          },
          {
            "name": "BenchmarkParsing/Iterator_../testdata/20110805A.ach - allocs/op",
            "value": 725,
            "unit": "allocs/op",
            "extra": "8193 times\n4 procs"
          },
          {
            "name": "BenchmarkParsing/IAT",
            "value": 144123,
            "unit": "ns/op\t   32197 B/op\t     307 allocs/op",
            "extra": "8773 times\n4 procs"
          },
          {
            "name": "BenchmarkParsing/IAT - ns/op",
            "value": 144123,
            "unit": "ns/op",
            "extra": "8773 times\n4 procs"
          },
          {
            "name": "BenchmarkParsing/IAT - B/op",
            "value": 32197,
            "unit": "B/op",
            "extra": "8773 times\n4 procs"
          },
          {
            "name": "BenchmarkParsing/IAT - allocs/op",
            "value": 307,
            "unit": "allocs/op",
            "extra": "8773 times\n4 procs"
          },
          {
            "name": "BenchmarkFile/String",
            "value": 54523,
            "unit": "ns/op\t   19367 B/op\t     444 allocs/op",
            "extra": "21562 times\n4 procs"
          },
          {
            "name": "BenchmarkFile/String - ns/op",
            "value": 54523,
            "unit": "ns/op",
            "extra": "21562 times\n4 procs"
          },
          {
            "name": "BenchmarkFile/String - B/op",
            "value": 19367,
            "unit": "B/op",
            "extra": "21562 times\n4 procs"
          },
          {
            "name": "BenchmarkFile/String - allocs/op",
            "value": 444,
            "unit": "allocs/op",
            "extra": "21562 times\n4 procs"
          }
        ]
      }
    ]
  }
}