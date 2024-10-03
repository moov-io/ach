window.BENCHMARK_DATA = {
  "lastUpdate": 1727963854600,
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
          "id": "92740d316a8f2cb754b577c05984007625305bcc",
          "message": "docs/bench: track fewer benchmarks",
          "timestamp": "2024-10-03T13:54:10Z",
          "url": "https://github.com/moov-io/ach/commit/92740d316a8f2cb754b577c05984007625305bcc"
        },
        "date": 1727963854556,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkBuildFile",
            "value": 13008,
            "unit": "ns/op\t    9668 B/op\t      99 allocs/op",
            "extra": "91822 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile - ns/op",
            "value": 13008,
            "unit": "ns/op",
            "extra": "91822 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile - B/op",
            "value": 9668,
            "unit": "B/op",
            "extra": "91822 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile - allocs/op",
            "value": 99,
            "unit": "allocs/op",
            "extra": "91822 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead",
            "value": 41219,
            "unit": "ns/op\t   21523 B/op\t      61 allocs/op",
            "extra": "29629 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead - ns/op",
            "value": 41219,
            "unit": "ns/op",
            "extra": "29629 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead - B/op",
            "value": 21523,
            "unit": "B/op",
            "extra": "29629 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead - allocs/op",
            "value": 61,
            "unit": "allocs/op",
            "extra": "29629 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead",
            "value": 74058,
            "unit": "ns/op\t   25388 B/op\t     136 allocs/op",
            "extra": "16208 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead - ns/op",
            "value": 74058,
            "unit": "ns/op",
            "extra": "16208 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead - B/op",
            "value": 25388,
            "unit": "B/op",
            "extra": "16208 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead - allocs/op",
            "value": 136,
            "unit": "allocs/op",
            "extra": "16208 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead",
            "value": 31208,
            "unit": "ns/op\t   20931 B/op\t      54 allocs/op",
            "extra": "38118 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead - ns/op",
            "value": 31208,
            "unit": "ns/op",
            "extra": "38118 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead - B/op",
            "value": 20931,
            "unit": "B/op",
            "extra": "38118 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead - allocs/op",
            "value": 54,
            "unit": "allocs/op",
            "extra": "38118 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead",
            "value": 323753,
            "unit": "ns/op\t   56275 B/op\t     743 allocs/op",
            "extra": "3706 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead - ns/op",
            "value": 323753,
            "unit": "ns/op",
            "extra": "3706 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead - B/op",
            "value": 56275,
            "unit": "B/op",
            "extra": "3706 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead - allocs/op",
            "value": 743,
            "unit": "allocs/op",
            "extra": "3706 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2",
            "value": 303713,
            "unit": "ns/op\t   56275 B/op\t     743 allocs/op",
            "extra": "3820 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2 - ns/op",
            "value": 303713,
            "unit": "ns/op",
            "extra": "3820 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2 - B/op",
            "value": 56275,
            "unit": "B/op",
            "extra": "3820 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2 - allocs/op",
            "value": 743,
            "unit": "allocs/op",
            "extra": "3820 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3",
            "value": 101592,
            "unit": "ns/op\t   27427 B/op\t     199 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3 - ns/op",
            "value": 101592,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3 - B/op",
            "value": 27427,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3 - allocs/op",
            "value": 199,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite",
            "value": 46690,
            "unit": "ns/op\t   31593 B/op\t     130 allocs/op",
            "extra": "26035 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite - ns/op",
            "value": 46690,
            "unit": "ns/op",
            "extra": "26035 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite - B/op",
            "value": 31593,
            "unit": "B/op",
            "extra": "26035 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite - allocs/op",
            "value": 130,
            "unit": "allocs/op",
            "extra": "26035 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite",
            "value": 231002,
            "unit": "ns/op\t   53917 B/op\t    2041 allocs/op",
            "extra": "5407 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite - ns/op",
            "value": 231002,
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
            "value": 5858,
            "unit": "ns/op\t    6145 B/op\t      25 allocs/op",
            "extra": "197896 times\n4 procs"
          },
          {
            "name": "BenchmarkFileWriteErr - ns/op",
            "value": 5858,
            "unit": "ns/op",
            "extra": "197896 times\n4 procs"
          },
          {
            "name": "BenchmarkFileWriteErr - B/op",
            "value": 6145,
            "unit": "B/op",
            "extra": "197896 times\n4 procs"
          },
          {
            "name": "BenchmarkFileWriteErr - allocs/op",
            "value": 25,
            "unit": "allocs/op",
            "extra": "197896 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite",
            "value": 157673,
            "unit": "ns/op\t   57106 B/op\t     612 allocs/op",
            "extra": "7772 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite - ns/op",
            "value": 157673,
            "unit": "ns/op",
            "extra": "7772 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite - B/op",
            "value": 57106,
            "unit": "B/op",
            "extra": "7772 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite - allocs/op",
            "value": 612,
            "unit": "allocs/op",
            "extra": "7772 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDIATWrite",
            "value": 231380104,
            "unit": "1210428822410041356A094101Federal Reserve Bank   My Bank Name                   ",
            "extra": "101 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDIATWrite - 1210428822410041356A094101Federal",
            "value": 231380104,
            "unit": "1210428822410041356A094101Federal",
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
            "name": "BenchmarkAlphaFieldShort",
            "value": 1237,
            "unit": "ns/op\t      96 B/op\t       4 allocs/op",
            "extra": "922606 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldShort - ns/op",
            "value": 1237,
            "unit": "ns/op",
            "extra": "922606 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldShort - B/op",
            "value": 96,
            "unit": "B/op",
            "extra": "922606 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldShort - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "922606 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong",
            "value": 13.26,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "91214818 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong - ns/op",
            "value": 13.26,
            "unit": "ns/op",
            "extra": "91214818 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "91214818 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "91214818 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort",
            "value": 63.99,
            "unit": "ns/op\t      16 B/op\t       2 allocs/op",
            "extra": "18386058 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort - ns/op",
            "value": 63.99,
            "unit": "ns/op",
            "extra": "18386058 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "18386058 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "18386058 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong",
            "value": 30.62,
            "unit": "ns/op\t       8 B/op\t       1 allocs/op",
            "extra": "37462834 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong - ns/op",
            "value": 30.62,
            "unit": "ns/op",
            "extra": "37462834 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong - B/op",
            "value": 8,
            "unit": "B/op",
            "extra": "37462834 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "37462834 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField",
            "value": 14.55,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "83876121 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField - ns/op",
            "value": 14.55,
            "unit": "ns/op",
            "extra": "83876121 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "83876121 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "83876121 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField",
            "value": 6.534,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "184392856 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField - ns/op",
            "value": 6.534,
            "unit": "ns/op",
            "extra": "184392856 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "184392856 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "184392856 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles",
            "value": 319188,
            "unit": "ns/op\t   56828 B/op\t     637 allocs/op",
            "extra": "3954 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles - ns/op",
            "value": 319188,
            "unit": "ns/op",
            "extra": "3954 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles - B/op",
            "value": 56828,
            "unit": "B/op",
            "extra": "3954 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles - allocs/op",
            "value": 637,
            "unit": "allocs/op",
            "extra": "3954 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts",
            "value": 309645,
            "unit": "ns/op\t   56827 B/op\t     637 allocs/op",
            "extra": "4015 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts - ns/op",
            "value": 309645,
            "unit": "ns/op",
            "extra": "4015 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts - B/op",
            "value": 56827,
            "unit": "B/op",
            "extra": "4015 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts - allocs/op",
            "value": 637,
            "unit": "allocs/op",
            "extra": "4015 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir",
            "value": 179402,
            "unit": "ns/op\t   57019 B/op\t     640 allocs/op",
            "extra": "6537 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir - ns/op",
            "value": 179402,
            "unit": "ns/op",
            "extra": "6537 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir - B/op",
            "value": 57019,
            "unit": "B/op",
            "extra": "6537 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir - allocs/op",
            "value": 640,
            "unit": "allocs/op",
            "extra": "6537 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts",
            "value": 197737,
            "unit": "ns/op\t   57054 B/op\t     640 allocs/op",
            "extra": "6049 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts - ns/op",
            "value": 197737,
            "unit": "ns/op",
            "extra": "6049 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts - B/op",
            "value": 57054,
            "unit": "B/op",
            "extra": "6049 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts - allocs/op",
            "value": 640,
            "unit": "allocs/op",
            "extra": "6049 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups",
            "value": 376747,
            "unit": "ns/op\t   62545 B/op\t     697 allocs/op",
            "extra": "3168 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups - ns/op",
            "value": 376747,
            "unit": "ns/op",
            "extra": "3168 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups - B/op",
            "value": 62545,
            "unit": "B/op",
            "extra": "3168 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups - allocs/op",
            "value": 697,
            "unit": "allocs/op",
            "extra": "3168 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups",
            "value": 378191,
            "unit": "ns/op\t   62529 B/op\t     697 allocs/op",
            "extra": "3187 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups - ns/op",
            "value": 378191,
            "unit": "ns/op",
            "extra": "3187 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups - B/op",
            "value": 62529,
            "unit": "B/op",
            "extra": "3187 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups - allocs/op",
            "value": 697,
            "unit": "allocs/op",
            "extra": "3187 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups",
            "value": 373000,
            "unit": "ns/op\t   62565 B/op\t     697 allocs/op",
            "extra": "3344 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups - ns/op",
            "value": 373000,
            "unit": "ns/op",
            "extra": "3344 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups - B/op",
            "value": 62565,
            "unit": "B/op",
            "extra": "3344 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups - allocs/op",
            "value": 697,
            "unit": "allocs/op",
            "extra": "3344 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups",
            "value": 371661,
            "unit": "ns/op\t   62539 B/op\t     697 allocs/op",
            "extra": "4033 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups - ns/op",
            "value": 371661,
            "unit": "ns/op",
            "extra": "4033 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups - B/op",
            "value": 62539,
            "unit": "B/op",
            "extra": "4033 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups - allocs/op",
            "value": 697,
            "unit": "allocs/op",
            "extra": "4033 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid",
            "value": 33.57,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "35654580 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid - ns/op",
            "value": 33.57,
            "unit": "ns/op",
            "extra": "35654580 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "35654580 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "35654580 times\n4 procs"
          },
          {
            "name": "Benchmark__FlattenBigFile",
            "value": 4887804618,
            "unit": "ns/op\t3211097024 B/op\t 2020946 allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark__FlattenBigFile - ns/op",
            "value": 4887804618,
            "unit": "ns/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark__FlattenBigFile - B/op",
            "value": 3211097024,
            "unit": "B/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark__FlattenBigFile - allocs/op",
            "value": 2020946,
            "unit": "allocs/op",
            "extra": "1 times\n4 procs"
          }
        ]
      }
    ]
  }
}