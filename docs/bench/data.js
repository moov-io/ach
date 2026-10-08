window.BENCHMARK_DATA = {
  "lastUpdate": 1791431628853,
  "repoUrl": "https://github.com/moov-io/ach",
  "entries": {
    "moov-io/ach": [
      {
        "commit": {
          "author": {
            "name": "Adam Shannon",
            "username": "adamdecaf",
            "email": "adamkshannon@gmail.com"
          },
          "committer": {
            "name": "GitHub",
            "username": "web-flow",
            "email": "noreply@github.com"
          },
          "id": "7f8e51d9f6e2feb085f1e1ffab8706ab41d58692",
          "message": "Run hot-path Go benchmarks in this repository. (#1857)\n\nStore results in docs/bench and label them with the ACH commit that ran,\nnot a hash from moov-io/benchmarks.",
          "timestamp": "2026-09-14T15:13:48Z",
          "url": "https://github.com/moov-io/ach/commit/7f8e51d9f6e2feb085f1e1ffab8706ab41d58692"
        },
        "date": 1789399612352,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkAlphaFieldShort",
            "value": 801.3,
            "unit": "ns/op\t      80 B/op\t       3 allocs/op",
            "extra": "1498614 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldShort - ns/op",
            "value": 801.3,
            "unit": "ns/op",
            "extra": "1498614 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldShort - B/op",
            "value": 80,
            "unit": "B/op",
            "extra": "1498614 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldShort - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "1498614 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong",
            "value": 82.44,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "14049914 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong - ns/op",
            "value": 82.44,
            "unit": "ns/op",
            "extra": "14049914 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "14049914 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "14049914 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort",
            "value": 47.85,
            "unit": "ns/op\t      16 B/op\t       2 allocs/op",
            "extra": "23817126 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort - ns/op",
            "value": 47.85,
            "unit": "ns/op",
            "extra": "23817126 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "23817126 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "23817126 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong",
            "value": 21.86,
            "unit": "ns/op\t       8 B/op\t       1 allocs/op",
            "extra": "52570362 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong - ns/op",
            "value": 21.86,
            "unit": "ns/op",
            "extra": "52570362 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong - B/op",
            "value": 8,
            "unit": "B/op",
            "extra": "52570362 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "52570362 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField",
            "value": 14.7,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "87268905 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField - ns/op",
            "value": 14.7,
            "unit": "ns/op",
            "extra": "87268905 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "87268905 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "87268905 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField",
            "value": 4.736,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "251318556 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField - ns/op",
            "value": 4.736,
            "unit": "ns/op",
            "extra": "251318556 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "251318556 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "251318556 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles",
            "value": 207340,
            "unit": "ns/op\t   54143 B/op\t     306 allocs/op",
            "extra": "5997 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles - ns/op",
            "value": 207340,
            "unit": "ns/op",
            "extra": "5997 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles - B/op",
            "value": 54143,
            "unit": "B/op",
            "extra": "5997 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles - allocs/op",
            "value": 306,
            "unit": "allocs/op",
            "extra": "5997 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts",
            "value": 202404,
            "unit": "ns/op\t   54155 B/op\t     306 allocs/op",
            "extra": "5703 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts - ns/op",
            "value": 202404,
            "unit": "ns/op",
            "extra": "5703 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts - B/op",
            "value": 54155,
            "unit": "B/op",
            "extra": "5703 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts - allocs/op",
            "value": 306,
            "unit": "allocs/op",
            "extra": "5703 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir",
            "value": 120182,
            "unit": "ns/op\t   54549 B/op\t     310 allocs/op",
            "extra": "10051 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir - ns/op",
            "value": 120182,
            "unit": "ns/op",
            "extra": "10051 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir - B/op",
            "value": 54549,
            "unit": "B/op",
            "extra": "10051 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir - allocs/op",
            "value": 310,
            "unit": "allocs/op",
            "extra": "10051 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts",
            "value": 116515,
            "unit": "ns/op\t   54594 B/op\t     310 allocs/op",
            "extra": "8733 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts - ns/op",
            "value": 116515,
            "unit": "ns/op",
            "extra": "8733 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts - B/op",
            "value": 54594,
            "unit": "B/op",
            "extra": "8733 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts - allocs/op",
            "value": 310,
            "unit": "allocs/op",
            "extra": "8733 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups",
            "value": 261974,
            "unit": "ns/op\t   59885 B/op\t     366 allocs/op",
            "extra": "4516 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups - ns/op",
            "value": 261974,
            "unit": "ns/op",
            "extra": "4516 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups - B/op",
            "value": 59885,
            "unit": "B/op",
            "extra": "4516 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups - allocs/op",
            "value": 366,
            "unit": "allocs/op",
            "extra": "4516 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups",
            "value": 260252,
            "unit": "ns/op\t   59858 B/op\t     366 allocs/op",
            "extra": "4801 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups - ns/op",
            "value": 260252,
            "unit": "ns/op",
            "extra": "4801 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups - B/op",
            "value": 59858,
            "unit": "B/op",
            "extra": "4801 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups - allocs/op",
            "value": 366,
            "unit": "allocs/op",
            "extra": "4801 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups",
            "value": 252552,
            "unit": "ns/op\t   59819 B/op\t     366 allocs/op",
            "extra": "4675 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups - ns/op",
            "value": 252552,
            "unit": "ns/op",
            "extra": "4675 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups - B/op",
            "value": 59819,
            "unit": "B/op",
            "extra": "4675 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups - allocs/op",
            "value": 366,
            "unit": "allocs/op",
            "extra": "4675 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups",
            "value": 244913,
            "unit": "ns/op\t   59953 B/op\t     367 allocs/op",
            "extra": "5584 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups - ns/op",
            "value": 244913,
            "unit": "ns/op",
            "extra": "5584 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups - B/op",
            "value": 59953,
            "unit": "B/op",
            "extra": "5584 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups - allocs/op",
            "value": 367,
            "unit": "allocs/op",
            "extra": "5584 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader",
            "value": 210338067,
            "unit": "ns/op\t42450132 B/op\t  203653 allocs/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader - ns/op",
            "value": 210338067,
            "unit": "ns/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader - B/op",
            "value": 42450132,
            "unit": "B/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader - allocs/op",
            "value": 203653,
            "unit": "allocs/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator",
            "value": 66555586,
            "unit": "ns/op\t42424726 B/op\t  203601 allocs/op",
            "extra": "16 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator - ns/op",
            "value": 66555586,
            "unit": "ns/op",
            "extra": "16 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator - B/op",
            "value": 42424726,
            "unit": "B/op",
            "extra": "16 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator - allocs/op",
            "value": 203601,
            "unit": "allocs/op",
            "extra": "16 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader",
            "value": 391472299,
            "unit": "ns/op\t62131664 B/op\t  905793 allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader - ns/op",
            "value": 391472299,
            "unit": "ns/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader - B/op",
            "value": 62131664,
            "unit": "B/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader - allocs/op",
            "value": 905793,
            "unit": "allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator",
            "value": 128568767,
            "unit": "ns/op\t60508681 B/op\t  705838 allocs/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator - ns/op",
            "value": 128568767,
            "unit": "ns/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator - B/op",
            "value": 60508681,
            "unit": "B/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator - allocs/op",
            "value": 705838,
            "unit": "allocs/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader",
            "value": 1145815731,
            "unit": "ns/op\t215541288 B/op\t 1042866 allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader - ns/op",
            "value": 1145815731,
            "unit": "ns/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader - B/op",
            "value": 215541288,
            "unit": "B/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader - allocs/op",
            "value": 1042866,
            "unit": "allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator",
            "value": 340884878,
            "unit": "ns/op\t215424698 B/op\t 1042799 allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator - ns/op",
            "value": 340884878,
            "unit": "ns/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator - B/op",
            "value": 215424698,
            "unit": "B/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator - allocs/op",
            "value": 1042799,
            "unit": "allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader",
            "value": 1929381496,
            "unit": "ns/op\t313545016 B/op\t 4568034 allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader - ns/op",
            "value": 1929381496,
            "unit": "ns/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader - B/op",
            "value": 313545016,
            "unit": "B/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader - allocs/op",
            "value": 4568034,
            "unit": "allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator",
            "value": 654385182,
            "unit": "ns/op\t305430324 B/op\t 3568075 allocs/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator - ns/op",
            "value": 654385182,
            "unit": "ns/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator - B/op",
            "value": 305430324,
            "unit": "B/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator - allocs/op",
            "value": 3568075,
            "unit": "allocs/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead",
            "value": 32321,
            "unit": "ns/op\t   23312 B/op\t     114 allocs/op",
            "extra": "38272 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead - ns/op",
            "value": 32321,
            "unit": "ns/op",
            "extra": "38272 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead - B/op",
            "value": 23312,
            "unit": "B/op",
            "extra": "38272 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead - allocs/op",
            "value": 114,
            "unit": "allocs/op",
            "extra": "38272 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead",
            "value": 57182,
            "unit": "ns/op\t   26920 B/op\t     159 allocs/op",
            "extra": "20520 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead - ns/op",
            "value": 57182,
            "unit": "ns/op",
            "extra": "20520 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead - B/op",
            "value": 26920,
            "unit": "B/op",
            "extra": "20520 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead - allocs/op",
            "value": 159,
            "unit": "allocs/op",
            "extra": "20520 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead",
            "value": 23287,
            "unit": "ns/op\t   21848 B/op\t      77 allocs/op",
            "extra": "52671 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead - ns/op",
            "value": 23287,
            "unit": "ns/op",
            "extra": "52671 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead - B/op",
            "value": 21848,
            "unit": "B/op",
            "extra": "52671 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead - allocs/op",
            "value": 77,
            "unit": "allocs/op",
            "extra": "52671 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead",
            "value": 216119,
            "unit": "ns/op\t   56353 B/op\t     552 allocs/op",
            "extra": "5395 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead - ns/op",
            "value": 216119,
            "unit": "ns/op",
            "extra": "5395 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead - B/op",
            "value": 56353,
            "unit": "B/op",
            "extra": "5395 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead - allocs/op",
            "value": 552,
            "unit": "allocs/op",
            "extra": "5395 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2",
            "value": 216112,
            "unit": "ns/op\t   56352 B/op\t     552 allocs/op",
            "extra": "5422 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2 - ns/op",
            "value": 216112,
            "unit": "ns/op",
            "extra": "5422 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2 - B/op",
            "value": 56352,
            "unit": "B/op",
            "extra": "5422 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2 - allocs/op",
            "value": 552,
            "unit": "allocs/op",
            "extra": "5422 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3",
            "value": 77903,
            "unit": "ns/op\t   29752 B/op\t     264 allocs/op",
            "extra": "14914 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3 - ns/op",
            "value": 77903,
            "unit": "ns/op",
            "extra": "14914 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3 - B/op",
            "value": 29752,
            "unit": "B/op",
            "extra": "14914 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3 - allocs/op",
            "value": 264,
            "unit": "allocs/op",
            "extra": "14914 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile",
            "value": 12081,
            "unit": "ns/op\t   10760 B/op\t     131 allocs/op",
            "extra": "107479 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile - ns/op",
            "value": 12081,
            "unit": "ns/op",
            "extra": "107479 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile - B/op",
            "value": 10760,
            "unit": "B/op",
            "extra": "107479 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile - allocs/op",
            "value": 131,
            "unit": "allocs/op",
            "extra": "107479 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid",
            "value": 26.11,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "45391250 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid - ns/op",
            "value": 26.11,
            "unit": "ns/op",
            "extra": "45391250 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "45391250 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "45391250 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite",
            "value": 41116,
            "unit": "ns/op\t   34744 B/op\t     228 allocs/op",
            "extra": "26910 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite - ns/op",
            "value": 41116,
            "unit": "ns/op",
            "extra": "26910 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite - B/op",
            "value": 34744,
            "unit": "B/op",
            "extra": "26910 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite - allocs/op",
            "value": 228,
            "unit": "allocs/op",
            "extra": "26910 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite",
            "value": 191804,
            "unit": "ns/op\t   54680 B/op\t    2069 allocs/op",
            "extra": "6422 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite - ns/op",
            "value": 191804,
            "unit": "ns/op",
            "extra": "6422 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite - B/op",
            "value": 54680,
            "unit": "B/op",
            "extra": "6422 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite - allocs/op",
            "value": 2069,
            "unit": "allocs/op",
            "extra": "6422 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite",
            "value": 129836,
            "unit": "ns/op\t   61152 B/op\t     723 allocs/op",
            "extra": "8640 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite - ns/op",
            "value": 129836,
            "unit": "ns/op",
            "extra": "8640 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite - B/op",
            "value": 61152,
            "unit": "B/op",
            "extra": "8640 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite - allocs/op",
            "value": 723,
            "unit": "allocs/op",
            "extra": "8640 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDIATWrite",
            "value": 231380104,
            "unit": "1210428822609151526A094101Federal Reserve Bank   My Bank Name                   ",
            "extra": "101 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDIATWrite - 1210428822609151526A094101Federal",
            "value": 231380104,
            "unit": "1210428822609151526A094101Federal",
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
          }
        ]
      },
      {
        "commit": {
          "author": {
            "name": "Adam Shannon",
            "username": "adamdecaf",
            "email": "adamkshannon@gmail.com"
          },
          "committer": {
            "name": "GitHub",
            "username": "web-flow",
            "email": "noreply@github.com"
          },
          "id": "7f8e51d9f6e2feb085f1e1ffab8706ab41d58692",
          "message": "Run hot-path Go benchmarks in this repository. (#1857)\n\nStore results in docs/bench and label them with the ACH commit that ran,\nnot a hash from moov-io/benchmarks.",
          "timestamp": "2026-09-14T15:13:48Z",
          "url": "https://github.com/moov-io/ach/commit/7f8e51d9f6e2feb085f1e1ffab8706ab41d58692"
        },
        "date": 1789400882217,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkAlphaFieldShort",
            "value": 1089,
            "unit": "ns/op\t      80 B/op\t       3 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldShort - ns/op",
            "value": 1089,
            "unit": "ns/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldShort - B/op",
            "value": 80,
            "unit": "B/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldShort - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong",
            "value": 99.79,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "12047161 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong - ns/op",
            "value": 99.79,
            "unit": "ns/op",
            "extra": "12047161 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "12047161 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "12047161 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort",
            "value": 56.66,
            "unit": "ns/op\t      16 B/op\t       2 allocs/op",
            "extra": "20632106 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort - ns/op",
            "value": 56.66,
            "unit": "ns/op",
            "extra": "20632106 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "20632106 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "20632106 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong",
            "value": 27.46,
            "unit": "ns/op\t       8 B/op\t       1 allocs/op",
            "extra": "44207126 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong - ns/op",
            "value": 27.46,
            "unit": "ns/op",
            "extra": "44207126 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong - B/op",
            "value": 8,
            "unit": "B/op",
            "extra": "44207126 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "44207126 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField",
            "value": 15.16,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "80327280 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField - ns/op",
            "value": 15.16,
            "unit": "ns/op",
            "extra": "80327280 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "80327280 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "80327280 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField",
            "value": 5.916,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "202981340 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField - ns/op",
            "value": 5.916,
            "unit": "ns/op",
            "extra": "202981340 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "202981340 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "202981340 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles",
            "value": 254864,
            "unit": "ns/op\t   54150 B/op\t     306 allocs/op",
            "extra": "4058 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles - ns/op",
            "value": 254864,
            "unit": "ns/op",
            "extra": "4058 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles - B/op",
            "value": 54150,
            "unit": "B/op",
            "extra": "4058 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles - allocs/op",
            "value": 306,
            "unit": "allocs/op",
            "extra": "4058 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts",
            "value": 241504,
            "unit": "ns/op\t   54158 B/op\t     306 allocs/op",
            "extra": "4670 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts - ns/op",
            "value": 241504,
            "unit": "ns/op",
            "extra": "4670 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts - B/op",
            "value": 54158,
            "unit": "B/op",
            "extra": "4670 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts - allocs/op",
            "value": 306,
            "unit": "allocs/op",
            "extra": "4670 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir",
            "value": 146700,
            "unit": "ns/op\t   54543 B/op\t     310 allocs/op",
            "extra": "8101 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir - ns/op",
            "value": 146700,
            "unit": "ns/op",
            "extra": "8101 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir - B/op",
            "value": 54543,
            "unit": "B/op",
            "extra": "8101 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir - allocs/op",
            "value": 310,
            "unit": "allocs/op",
            "extra": "8101 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts",
            "value": 142714,
            "unit": "ns/op\t   54578 B/op\t     310 allocs/op",
            "extra": "8133 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts - ns/op",
            "value": 142714,
            "unit": "ns/op",
            "extra": "8133 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts - B/op",
            "value": 54578,
            "unit": "B/op",
            "extra": "8133 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts - allocs/op",
            "value": 310,
            "unit": "allocs/op",
            "extra": "8133 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups",
            "value": 311206,
            "unit": "ns/op\t   59834 B/op\t     366 allocs/op",
            "extra": "3920 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups - ns/op",
            "value": 311206,
            "unit": "ns/op",
            "extra": "3920 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups - B/op",
            "value": 59834,
            "unit": "B/op",
            "extra": "3920 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups - allocs/op",
            "value": 366,
            "unit": "allocs/op",
            "extra": "3920 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups",
            "value": 303385,
            "unit": "ns/op\t   59870 B/op\t     366 allocs/op",
            "extra": "4207 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups - ns/op",
            "value": 303385,
            "unit": "ns/op",
            "extra": "4207 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups - B/op",
            "value": 59870,
            "unit": "B/op",
            "extra": "4207 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups - allocs/op",
            "value": 366,
            "unit": "allocs/op",
            "extra": "4207 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups",
            "value": 294229,
            "unit": "ns/op\t   59881 B/op\t     366 allocs/op",
            "extra": "4171 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups - ns/op",
            "value": 294229,
            "unit": "ns/op",
            "extra": "4171 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups - B/op",
            "value": 59881,
            "unit": "B/op",
            "extra": "4171 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups - allocs/op",
            "value": 366,
            "unit": "allocs/op",
            "extra": "4171 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups",
            "value": 286290,
            "unit": "ns/op\t   59535 B/op\t     367 allocs/op",
            "extra": "4936 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups - ns/op",
            "value": 286290,
            "unit": "ns/op",
            "extra": "4936 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups - B/op",
            "value": 59535,
            "unit": "B/op",
            "extra": "4936 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups - allocs/op",
            "value": 367,
            "unit": "allocs/op",
            "extra": "4936 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader",
            "value": 241004243,
            "unit": "ns/op\t42450206 B/op\t  203654 allocs/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader - ns/op",
            "value": 241004243,
            "unit": "ns/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader - B/op",
            "value": 42450206,
            "unit": "B/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader - allocs/op",
            "value": 203654,
            "unit": "allocs/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator",
            "value": 78738996,
            "unit": "ns/op\t42424772 B/op\t  203601 allocs/op",
            "extra": "14 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator - ns/op",
            "value": 78738996,
            "unit": "ns/op",
            "extra": "14 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator - B/op",
            "value": 42424772,
            "unit": "B/op",
            "extra": "14 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator - allocs/op",
            "value": 203601,
            "unit": "allocs/op",
            "extra": "14 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader",
            "value": 451183130,
            "unit": "ns/op\t62131413 B/op\t  905790 allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader - ns/op",
            "value": 451183130,
            "unit": "ns/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader - B/op",
            "value": 62131413,
            "unit": "B/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader - allocs/op",
            "value": 905790,
            "unit": "allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator",
            "value": 150838720,
            "unit": "ns/op\t60508916 B/op\t  705841 allocs/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator - ns/op",
            "value": 150838720,
            "unit": "ns/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator - B/op",
            "value": 60508916,
            "unit": "B/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator - allocs/op",
            "value": 705841,
            "unit": "allocs/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader",
            "value": 1294147296,
            "unit": "ns/op\t215541432 B/op\t 1042868 allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader - ns/op",
            "value": 1294147296,
            "unit": "ns/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader - B/op",
            "value": 215541432,
            "unit": "B/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader - allocs/op",
            "value": 1042868,
            "unit": "allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator",
            "value": 390319644,
            "unit": "ns/op\t215424805 B/op\t 1042800 allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator - ns/op",
            "value": 390319644,
            "unit": "ns/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator - B/op",
            "value": 215424805,
            "unit": "B/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator - allocs/op",
            "value": 1042800,
            "unit": "allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader",
            "value": 2210218667,
            "unit": "ns/op\t313545464 B/op\t 4568040 allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader - ns/op",
            "value": 2210218667,
            "unit": "ns/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader - B/op",
            "value": 313545464,
            "unit": "B/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader - allocs/op",
            "value": 4568040,
            "unit": "allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator",
            "value": 754117453,
            "unit": "ns/op\t305430092 B/op\t 3568072 allocs/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator - ns/op",
            "value": 754117453,
            "unit": "ns/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator - B/op",
            "value": 305430092,
            "unit": "B/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator - allocs/op",
            "value": 3568072,
            "unit": "allocs/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead",
            "value": 41378,
            "unit": "ns/op\t   23312 B/op\t     114 allocs/op",
            "extra": "29547 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead - ns/op",
            "value": 41378,
            "unit": "ns/op",
            "extra": "29547 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead - B/op",
            "value": 23312,
            "unit": "B/op",
            "extra": "29547 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead - allocs/op",
            "value": 114,
            "unit": "allocs/op",
            "extra": "29547 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead",
            "value": 70630,
            "unit": "ns/op\t   26920 B/op\t     159 allocs/op",
            "extra": "16377 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead - ns/op",
            "value": 70630,
            "unit": "ns/op",
            "extra": "16377 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead - B/op",
            "value": 26920,
            "unit": "B/op",
            "extra": "16377 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead - allocs/op",
            "value": 159,
            "unit": "allocs/op",
            "extra": "16377 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead",
            "value": 30888,
            "unit": "ns/op\t   21848 B/op\t      77 allocs/op",
            "extra": "39459 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead - ns/op",
            "value": 30888,
            "unit": "ns/op",
            "extra": "39459 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead - B/op",
            "value": 21848,
            "unit": "B/op",
            "extra": "39459 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead - allocs/op",
            "value": 77,
            "unit": "allocs/op",
            "extra": "39459 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead",
            "value": 268894,
            "unit": "ns/op\t   56353 B/op\t     552 allocs/op",
            "extra": "4477 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead - ns/op",
            "value": 268894,
            "unit": "ns/op",
            "extra": "4477 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead - B/op",
            "value": 56353,
            "unit": "B/op",
            "extra": "4477 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead - allocs/op",
            "value": 552,
            "unit": "allocs/op",
            "extra": "4477 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2",
            "value": 268787,
            "unit": "ns/op\t   56353 B/op\t     552 allocs/op",
            "extra": "4690 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2 - ns/op",
            "value": 268787,
            "unit": "ns/op",
            "extra": "4690 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2 - B/op",
            "value": 56353,
            "unit": "B/op",
            "extra": "4690 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2 - allocs/op",
            "value": 552,
            "unit": "allocs/op",
            "extra": "4690 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3",
            "value": 102030,
            "unit": "ns/op\t   29752 B/op\t     264 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3 - ns/op",
            "value": 102030,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3 - B/op",
            "value": 29752,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3 - allocs/op",
            "value": 264,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile",
            "value": 14087,
            "unit": "ns/op\t   10760 B/op\t     131 allocs/op",
            "extra": "92336 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile - ns/op",
            "value": 14087,
            "unit": "ns/op",
            "extra": "92336 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile - B/op",
            "value": 10760,
            "unit": "B/op",
            "extra": "92336 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile - allocs/op",
            "value": 131,
            "unit": "allocs/op",
            "extra": "92336 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid",
            "value": 25.51,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "45283772 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid - ns/op",
            "value": 25.51,
            "unit": "ns/op",
            "extra": "45283772 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "45283772 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "45283772 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite",
            "value": 49241,
            "unit": "ns/op\t   34744 B/op\t     228 allocs/op",
            "extra": "23016 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite - ns/op",
            "value": 49241,
            "unit": "ns/op",
            "extra": "23016 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite - B/op",
            "value": 34744,
            "unit": "B/op",
            "extra": "23016 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite - allocs/op",
            "value": 228,
            "unit": "allocs/op",
            "extra": "23016 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite",
            "value": 216217,
            "unit": "ns/op\t   54680 B/op\t    2069 allocs/op",
            "extra": "5748 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite - ns/op",
            "value": 216217,
            "unit": "ns/op",
            "extra": "5748 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite - B/op",
            "value": 54680,
            "unit": "B/op",
            "extra": "5748 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite - allocs/op",
            "value": 2069,
            "unit": "allocs/op",
            "extra": "5748 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite",
            "value": 152373,
            "unit": "ns/op\t   61152 B/op\t     723 allocs/op",
            "extra": "7443 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite - ns/op",
            "value": 152373,
            "unit": "ns/op",
            "extra": "7443 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite - B/op",
            "value": 61152,
            "unit": "B/op",
            "extra": "7443 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite - allocs/op",
            "value": 723,
            "unit": "allocs/op",
            "extra": "7443 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDIATWrite",
            "value": 231380104,
            "unit": "1210428822609151547A094101Federal Reserve Bank   My Bank Name                   ",
            "extra": "101 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDIATWrite - 1210428822609151547A094101Federal",
            "value": 231380104,
            "unit": "1210428822609151547A094101Federal",
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
          }
        ]
      },
      {
        "commit": {
          "author": {
            "name": "Adam Shannon",
            "username": "adamdecaf",
            "email": "adamkshannon@gmail.com"
          },
          "committer": {
            "name": "GitHub",
            "username": "web-flow",
            "email": "noreply@github.com"
          },
          "id": "7f8e51d9f6e2feb085f1e1ffab8706ab41d58692",
          "message": "Run hot-path Go benchmarks in this repository. (#1857)\n\nStore results in docs/bench and label them with the ACH commit that ran,\nnot a hash from moov-io/benchmarks.",
          "timestamp": "2026-09-14T15:13:48Z",
          "url": "https://github.com/moov-io/ach/commit/7f8e51d9f6e2feb085f1e1ffab8706ab41d58692"
        },
        "date": 1789401022244,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkAlphaFieldShort",
            "value": 893.5,
            "unit": "ns/op\t      80 B/op\t       3 allocs/op",
            "extra": "1309104 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldShort - ns/op",
            "value": 893.5,
            "unit": "ns/op",
            "extra": "1309104 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldShort - B/op",
            "value": 80,
            "unit": "B/op",
            "extra": "1309104 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldShort - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "1309104 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong",
            "value": 100.8,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "11818070 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong - ns/op",
            "value": 100.8,
            "unit": "ns/op",
            "extra": "11818070 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "11818070 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "11818070 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort",
            "value": 55.22,
            "unit": "ns/op\t      16 B/op\t       2 allocs/op",
            "extra": "21071067 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort - ns/op",
            "value": 55.22,
            "unit": "ns/op",
            "extra": "21071067 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "21071067 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "21071067 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong",
            "value": 26.81,
            "unit": "ns/op\t       8 B/op\t       1 allocs/op",
            "extra": "45934122 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong - ns/op",
            "value": 26.81,
            "unit": "ns/op",
            "extra": "45934122 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong - B/op",
            "value": 8,
            "unit": "B/op",
            "extra": "45934122 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "45934122 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField",
            "value": 14.01,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "87025378 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField - ns/op",
            "value": 14.01,
            "unit": "ns/op",
            "extra": "87025378 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "87025378 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "87025378 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField",
            "value": 5.648,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "212922252 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField - ns/op",
            "value": 5.648,
            "unit": "ns/op",
            "extra": "212922252 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "212922252 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "212922252 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles",
            "value": 253681,
            "unit": "ns/op\t   54139 B/op\t     306 allocs/op",
            "extra": "5418 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles - ns/op",
            "value": 253681,
            "unit": "ns/op",
            "extra": "5418 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles - B/op",
            "value": 54139,
            "unit": "B/op",
            "extra": "5418 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles - allocs/op",
            "value": 306,
            "unit": "allocs/op",
            "extra": "5418 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts",
            "value": 253731,
            "unit": "ns/op\t   54158 B/op\t     306 allocs/op",
            "extra": "5016 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts - ns/op",
            "value": 253731,
            "unit": "ns/op",
            "extra": "5016 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts - B/op",
            "value": 54158,
            "unit": "B/op",
            "extra": "5016 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts - allocs/op",
            "value": 306,
            "unit": "allocs/op",
            "extra": "5016 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir",
            "value": 141260,
            "unit": "ns/op\t   54543 B/op\t     310 allocs/op",
            "extra": "8404 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir - ns/op",
            "value": 141260,
            "unit": "ns/op",
            "extra": "8404 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir - B/op",
            "value": 54543,
            "unit": "B/op",
            "extra": "8404 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir - allocs/op",
            "value": 310,
            "unit": "allocs/op",
            "extra": "8404 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts",
            "value": 139242,
            "unit": "ns/op\t   54579 B/op\t     310 allocs/op",
            "extra": "7537 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts - ns/op",
            "value": 139242,
            "unit": "ns/op",
            "extra": "7537 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts - B/op",
            "value": 54579,
            "unit": "B/op",
            "extra": "7537 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts - allocs/op",
            "value": 310,
            "unit": "allocs/op",
            "extra": "7537 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups",
            "value": 305077,
            "unit": "ns/op\t   59869 B/op\t     366 allocs/op",
            "extra": "4292 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups - ns/op",
            "value": 305077,
            "unit": "ns/op",
            "extra": "4292 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups - B/op",
            "value": 59869,
            "unit": "B/op",
            "extra": "4292 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups - allocs/op",
            "value": 366,
            "unit": "allocs/op",
            "extra": "4292 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups",
            "value": 294387,
            "unit": "ns/op\t   59857 B/op\t     366 allocs/op",
            "extra": "4267 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups - ns/op",
            "value": 294387,
            "unit": "ns/op",
            "extra": "4267 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups - B/op",
            "value": 59857,
            "unit": "B/op",
            "extra": "4267 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups - allocs/op",
            "value": 366,
            "unit": "allocs/op",
            "extra": "4267 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups",
            "value": 291383,
            "unit": "ns/op\t   59862 B/op\t     366 allocs/op",
            "extra": "4237 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups - ns/op",
            "value": 291383,
            "unit": "ns/op",
            "extra": "4237 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups - B/op",
            "value": 59862,
            "unit": "B/op",
            "extra": "4237 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups - allocs/op",
            "value": 366,
            "unit": "allocs/op",
            "extra": "4237 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups",
            "value": 280299,
            "unit": "ns/op\t   59511 B/op\t     367 allocs/op",
            "extra": "5017 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups - ns/op",
            "value": 280299,
            "unit": "ns/op",
            "extra": "5017 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups - B/op",
            "value": 59511,
            "unit": "B/op",
            "extra": "5017 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups - allocs/op",
            "value": 367,
            "unit": "allocs/op",
            "extra": "5017 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader",
            "value": 235258754,
            "unit": "ns/op\t42450097 B/op\t  203653 allocs/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader - ns/op",
            "value": 235258754,
            "unit": "ns/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader - B/op",
            "value": 42450097,
            "unit": "B/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader - allocs/op",
            "value": 203653,
            "unit": "allocs/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator",
            "value": 67295890,
            "unit": "ns/op\t42424987 B/op\t  203602 allocs/op",
            "extra": "16 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator - ns/op",
            "value": 67295890,
            "unit": "ns/op",
            "extra": "16 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator - B/op",
            "value": 42424987,
            "unit": "B/op",
            "extra": "16 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator - allocs/op",
            "value": 203602,
            "unit": "allocs/op",
            "extra": "16 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader",
            "value": 435683955,
            "unit": "ns/op\t62131690 B/op\t  905794 allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader - ns/op",
            "value": 435683955,
            "unit": "ns/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader - B/op",
            "value": 62131690,
            "unit": "B/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader - allocs/op",
            "value": 905794,
            "unit": "allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator",
            "value": 133813705,
            "unit": "ns/op\t60508718 B/op\t  705837 allocs/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator - ns/op",
            "value": 133813705,
            "unit": "ns/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator - B/op",
            "value": 60508718,
            "unit": "B/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator - allocs/op",
            "value": 705837,
            "unit": "allocs/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader",
            "value": 1288332004,
            "unit": "ns/op\t215541064 B/op\t 1042863 allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader - ns/op",
            "value": 1288332004,
            "unit": "ns/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader - B/op",
            "value": 215541064,
            "unit": "B/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader - allocs/op",
            "value": 1042863,
            "unit": "allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator",
            "value": 340589706,
            "unit": "ns/op\t215424794 B/op\t 1042800 allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator - ns/op",
            "value": 340589706,
            "unit": "ns/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator - B/op",
            "value": 215424794,
            "unit": "B/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator - allocs/op",
            "value": 1042800,
            "unit": "allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader",
            "value": 2149368545,
            "unit": "ns/op\t313545400 B/op\t 4568039 allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader - ns/op",
            "value": 2149368545,
            "unit": "ns/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader - B/op",
            "value": 313545400,
            "unit": "B/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader - allocs/op",
            "value": 4568039,
            "unit": "allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator",
            "value": 692797641,
            "unit": "ns/op\t305429948 B/op\t 3568070 allocs/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator - ns/op",
            "value": 692797641,
            "unit": "ns/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator - B/op",
            "value": 305429948,
            "unit": "B/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator - allocs/op",
            "value": 3568070,
            "unit": "allocs/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead",
            "value": 41632,
            "unit": "ns/op\t   23312 B/op\t     114 allocs/op",
            "extra": "29007 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead - ns/op",
            "value": 41632,
            "unit": "ns/op",
            "extra": "29007 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead - B/op",
            "value": 23312,
            "unit": "B/op",
            "extra": "29007 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead - allocs/op",
            "value": 114,
            "unit": "allocs/op",
            "extra": "29007 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead",
            "value": 69068,
            "unit": "ns/op\t   26920 B/op\t     159 allocs/op",
            "extra": "16674 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead - ns/op",
            "value": 69068,
            "unit": "ns/op",
            "extra": "16674 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead - B/op",
            "value": 26920,
            "unit": "B/op",
            "extra": "16674 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead - allocs/op",
            "value": 159,
            "unit": "allocs/op",
            "extra": "16674 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead",
            "value": 31839,
            "unit": "ns/op\t   21848 B/op\t      77 allocs/op",
            "extra": "38764 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead - ns/op",
            "value": 31839,
            "unit": "ns/op",
            "extra": "38764 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead - B/op",
            "value": 21848,
            "unit": "B/op",
            "extra": "38764 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead - allocs/op",
            "value": 77,
            "unit": "allocs/op",
            "extra": "38764 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead",
            "value": 257085,
            "unit": "ns/op\t   56353 B/op\t     552 allocs/op",
            "extra": "4854 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead - ns/op",
            "value": 257085,
            "unit": "ns/op",
            "extra": "4854 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead - B/op",
            "value": 56353,
            "unit": "B/op",
            "extra": "4854 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead - allocs/op",
            "value": 552,
            "unit": "allocs/op",
            "extra": "4854 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2",
            "value": 256772,
            "unit": "ns/op\t   56353 B/op\t     552 allocs/op",
            "extra": "4772 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2 - ns/op",
            "value": 256772,
            "unit": "ns/op",
            "extra": "4772 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2 - B/op",
            "value": 56353,
            "unit": "B/op",
            "extra": "4772 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2 - allocs/op",
            "value": 552,
            "unit": "allocs/op",
            "extra": "4772 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3",
            "value": 96738,
            "unit": "ns/op\t   29752 B/op\t     264 allocs/op",
            "extra": "12151 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3 - ns/op",
            "value": 96738,
            "unit": "ns/op",
            "extra": "12151 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3 - B/op",
            "value": 29752,
            "unit": "B/op",
            "extra": "12151 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3 - allocs/op",
            "value": 264,
            "unit": "allocs/op",
            "extra": "12151 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile",
            "value": 13421,
            "unit": "ns/op\t   10760 B/op\t     131 allocs/op",
            "extra": "99644 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile - ns/op",
            "value": 13421,
            "unit": "ns/op",
            "extra": "99644 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile - B/op",
            "value": 10760,
            "unit": "B/op",
            "extra": "99644 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile - allocs/op",
            "value": 131,
            "unit": "allocs/op",
            "extra": "99644 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid",
            "value": 24.28,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "49316632 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid - ns/op",
            "value": 24.28,
            "unit": "ns/op",
            "extra": "49316632 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "49316632 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "49316632 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite",
            "value": 46585,
            "unit": "ns/op\t   34744 B/op\t     228 allocs/op",
            "extra": "23826 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite - ns/op",
            "value": 46585,
            "unit": "ns/op",
            "extra": "23826 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite - B/op",
            "value": 34744,
            "unit": "B/op",
            "extra": "23826 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite - allocs/op",
            "value": 228,
            "unit": "allocs/op",
            "extra": "23826 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite",
            "value": 188838,
            "unit": "ns/op\t   54680 B/op\t    2069 allocs/op",
            "extra": "6640 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite - ns/op",
            "value": 188838,
            "unit": "ns/op",
            "extra": "6640 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite - B/op",
            "value": 54680,
            "unit": "B/op",
            "extra": "6640 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite - allocs/op",
            "value": 2069,
            "unit": "allocs/op",
            "extra": "6640 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite",
            "value": 145922,
            "unit": "ns/op\t   61152 B/op\t     723 allocs/op",
            "extra": "7870 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite - ns/op",
            "value": 145922,
            "unit": "ns/op",
            "extra": "7870 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite - B/op",
            "value": 61152,
            "unit": "B/op",
            "extra": "7870 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite - allocs/op",
            "value": 723,
            "unit": "allocs/op",
            "extra": "7870 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDIATWrite",
            "value": 231380104,
            "unit": "1210428822609151550A094101Federal Reserve Bank   My Bank Name                   ",
            "extra": "101 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDIATWrite - 1210428822609151550A094101Federal",
            "value": 231380104,
            "unit": "1210428822609151550A094101Federal",
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
          }
        ]
      },
      {
        "commit": {
          "author": {
            "name": "Adam Shannon",
            "username": "adamdecaf",
            "email": "adamkshannon@gmail.com"
          },
          "committer": {
            "name": "GitHub",
            "username": "web-flow",
            "email": "noreply@github.com"
          },
          "id": "7f8e51d9f6e2feb085f1e1ffab8706ab41d58692",
          "message": "Run hot-path Go benchmarks in this repository. (#1857)\n\nStore results in docs/bench and label them with the ACH commit that ran,\nnot a hash from moov-io/benchmarks.",
          "timestamp": "2026-09-14T15:13:48Z",
          "url": "https://github.com/moov-io/ach/commit/7f8e51d9f6e2feb085f1e1ffab8706ab41d58692"
        },
        "date": 1789440556630,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkAlphaFieldShort",
            "value": 911.7,
            "unit": "ns/op\t      80 B/op\t       3 allocs/op",
            "extra": "1300100 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldShort - ns/op",
            "value": 911.7,
            "unit": "ns/op",
            "extra": "1300100 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldShort - B/op",
            "value": 80,
            "unit": "B/op",
            "extra": "1300100 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldShort - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "1300100 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong",
            "value": 100.4,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "11814820 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong - ns/op",
            "value": 100.4,
            "unit": "ns/op",
            "extra": "11814820 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "11814820 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "11814820 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort",
            "value": 55.11,
            "unit": "ns/op\t      16 B/op\t       2 allocs/op",
            "extra": "20958662 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort - ns/op",
            "value": 55.11,
            "unit": "ns/op",
            "extra": "20958662 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "20958662 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "20958662 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong",
            "value": 25.2,
            "unit": "ns/op\t       8 B/op\t       1 allocs/op",
            "extra": "45423076 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong - ns/op",
            "value": 25.2,
            "unit": "ns/op",
            "extra": "45423076 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong - B/op",
            "value": 8,
            "unit": "B/op",
            "extra": "45423076 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "45423076 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField",
            "value": 13.76,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "87136432 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField - ns/op",
            "value": 13.76,
            "unit": "ns/op",
            "extra": "87136432 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "87136432 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "87136432 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField",
            "value": 5.632,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "209869243 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField - ns/op",
            "value": 5.632,
            "unit": "ns/op",
            "extra": "209869243 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "209869243 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "209869243 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles",
            "value": 253305,
            "unit": "ns/op\t   54139 B/op\t     306 allocs/op",
            "extra": "5480 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles - ns/op",
            "value": 253305,
            "unit": "ns/op",
            "extra": "5480 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles - B/op",
            "value": 54139,
            "unit": "B/op",
            "extra": "5480 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles - allocs/op",
            "value": 306,
            "unit": "allocs/op",
            "extra": "5480 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts",
            "value": 242421,
            "unit": "ns/op\t   54156 B/op\t     306 allocs/op",
            "extra": "5152 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts - ns/op",
            "value": 242421,
            "unit": "ns/op",
            "extra": "5152 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts - B/op",
            "value": 54156,
            "unit": "B/op",
            "extra": "5152 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts - allocs/op",
            "value": 306,
            "unit": "allocs/op",
            "extra": "5152 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir",
            "value": 146756,
            "unit": "ns/op\t   54564 B/op\t     310 allocs/op",
            "extra": "8754 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir - ns/op",
            "value": 146756,
            "unit": "ns/op",
            "extra": "8754 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir - B/op",
            "value": 54564,
            "unit": "B/op",
            "extra": "8754 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir - allocs/op",
            "value": 310,
            "unit": "allocs/op",
            "extra": "8754 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts",
            "value": 141098,
            "unit": "ns/op\t   54589 B/op\t     310 allocs/op",
            "extra": "8149 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts - ns/op",
            "value": 141098,
            "unit": "ns/op",
            "extra": "8149 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts - B/op",
            "value": 54589,
            "unit": "B/op",
            "extra": "8149 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts - allocs/op",
            "value": 310,
            "unit": "allocs/op",
            "extra": "8149 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups",
            "value": 307809,
            "unit": "ns/op\t   59863 B/op\t     366 allocs/op",
            "extra": "4220 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups - ns/op",
            "value": 307809,
            "unit": "ns/op",
            "extra": "4220 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups - B/op",
            "value": 59863,
            "unit": "B/op",
            "extra": "4220 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups - allocs/op",
            "value": 366,
            "unit": "allocs/op",
            "extra": "4220 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups",
            "value": 292267,
            "unit": "ns/op\t   59892 B/op\t     366 allocs/op",
            "extra": "4321 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups - ns/op",
            "value": 292267,
            "unit": "ns/op",
            "extra": "4321 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups - B/op",
            "value": 59892,
            "unit": "B/op",
            "extra": "4321 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups - allocs/op",
            "value": 366,
            "unit": "allocs/op",
            "extra": "4321 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups",
            "value": 294956,
            "unit": "ns/op\t   59850 B/op\t     366 allocs/op",
            "extra": "4215 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups - ns/op",
            "value": 294956,
            "unit": "ns/op",
            "extra": "4215 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups - B/op",
            "value": 59850,
            "unit": "B/op",
            "extra": "4215 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups - allocs/op",
            "value": 366,
            "unit": "allocs/op",
            "extra": "4215 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups",
            "value": 286024,
            "unit": "ns/op\t   60073 B/op\t     367 allocs/op",
            "extra": "5296 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups - ns/op",
            "value": 286024,
            "unit": "ns/op",
            "extra": "5296 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups - B/op",
            "value": 60073,
            "unit": "B/op",
            "extra": "5296 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups - allocs/op",
            "value": 367,
            "unit": "allocs/op",
            "extra": "5296 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader",
            "value": 234622604,
            "unit": "ns/op\t42450113 B/op\t  203653 allocs/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader - ns/op",
            "value": 234622604,
            "unit": "ns/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader - B/op",
            "value": 42450113,
            "unit": "B/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader - allocs/op",
            "value": 203653,
            "unit": "allocs/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator",
            "value": 67066800,
            "unit": "ns/op\t42424951 B/op\t  203603 allocs/op",
            "extra": "16 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator - ns/op",
            "value": 67066800,
            "unit": "ns/op",
            "extra": "16 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator - B/op",
            "value": 42424951,
            "unit": "B/op",
            "extra": "16 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator - allocs/op",
            "value": 203603,
            "unit": "allocs/op",
            "extra": "16 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader",
            "value": 434371211,
            "unit": "ns/op\t62131568 B/op\t  905792 allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader - ns/op",
            "value": 434371211,
            "unit": "ns/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader - B/op",
            "value": 62131568,
            "unit": "B/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader - allocs/op",
            "value": 905792,
            "unit": "allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator",
            "value": 134492137,
            "unit": "ns/op\t60508912 B/op\t  705840 allocs/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator - ns/op",
            "value": 134492137,
            "unit": "ns/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator - B/op",
            "value": 60508912,
            "unit": "B/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator - allocs/op",
            "value": 705840,
            "unit": "allocs/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader",
            "value": 1306172035,
            "unit": "ns/op\t215541432 B/op\t 1042868 allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader - ns/op",
            "value": 1306172035,
            "unit": "ns/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader - B/op",
            "value": 215541432,
            "unit": "B/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader - allocs/op",
            "value": 1042868,
            "unit": "allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator",
            "value": 339829235,
            "unit": "ns/op\t215424949 B/op\t 1042802 allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator - ns/op",
            "value": 339829235,
            "unit": "ns/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator - B/op",
            "value": 215424949,
            "unit": "B/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator - allocs/op",
            "value": 1042802,
            "unit": "allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader",
            "value": 2138969458,
            "unit": "ns/op\t313545192 B/op\t 4568035 allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader - ns/op",
            "value": 2138969458,
            "unit": "ns/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader - B/op",
            "value": 313545192,
            "unit": "B/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader - allocs/op",
            "value": 4568035,
            "unit": "allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator",
            "value": 680917921,
            "unit": "ns/op\t305430452 B/op\t 3568076 allocs/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator - ns/op",
            "value": 680917921,
            "unit": "ns/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator - B/op",
            "value": 305430452,
            "unit": "B/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator - allocs/op",
            "value": 3568076,
            "unit": "allocs/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead",
            "value": 41523,
            "unit": "ns/op\t   23312 B/op\t     114 allocs/op",
            "extra": "29179 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead - ns/op",
            "value": 41523,
            "unit": "ns/op",
            "extra": "29179 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead - B/op",
            "value": 23312,
            "unit": "B/op",
            "extra": "29179 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead - allocs/op",
            "value": 114,
            "unit": "allocs/op",
            "extra": "29179 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead",
            "value": 69258,
            "unit": "ns/op\t   26920 B/op\t     159 allocs/op",
            "extra": "16798 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead - ns/op",
            "value": 69258,
            "unit": "ns/op",
            "extra": "16798 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead - B/op",
            "value": 26920,
            "unit": "B/op",
            "extra": "16798 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead - allocs/op",
            "value": 159,
            "unit": "allocs/op",
            "extra": "16798 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead",
            "value": 31370,
            "unit": "ns/op\t   21848 B/op\t      77 allocs/op",
            "extra": "38682 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead - ns/op",
            "value": 31370,
            "unit": "ns/op",
            "extra": "38682 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead - B/op",
            "value": 21848,
            "unit": "B/op",
            "extra": "38682 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead - allocs/op",
            "value": 77,
            "unit": "allocs/op",
            "extra": "38682 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead",
            "value": 258358,
            "unit": "ns/op\t   56353 B/op\t     552 allocs/op",
            "extra": "4786 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead - ns/op",
            "value": 258358,
            "unit": "ns/op",
            "extra": "4786 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead - B/op",
            "value": 56353,
            "unit": "B/op",
            "extra": "4786 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead - allocs/op",
            "value": 552,
            "unit": "allocs/op",
            "extra": "4786 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2",
            "value": 256851,
            "unit": "ns/op\t   56353 B/op\t     552 allocs/op",
            "extra": "4790 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2 - ns/op",
            "value": 256851,
            "unit": "ns/op",
            "extra": "4790 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2 - B/op",
            "value": 56353,
            "unit": "B/op",
            "extra": "4790 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2 - allocs/op",
            "value": 552,
            "unit": "allocs/op",
            "extra": "4790 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3",
            "value": 96657,
            "unit": "ns/op\t   29752 B/op\t     264 allocs/op",
            "extra": "12198 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3 - ns/op",
            "value": 96657,
            "unit": "ns/op",
            "extra": "12198 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3 - B/op",
            "value": 29752,
            "unit": "B/op",
            "extra": "12198 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3 - allocs/op",
            "value": 264,
            "unit": "allocs/op",
            "extra": "12198 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile",
            "value": 13141,
            "unit": "ns/op\t   10760 B/op\t     131 allocs/op",
            "extra": "96823 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile - ns/op",
            "value": 13141,
            "unit": "ns/op",
            "extra": "96823 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile - B/op",
            "value": 10760,
            "unit": "B/op",
            "extra": "96823 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile - allocs/op",
            "value": 131,
            "unit": "allocs/op",
            "extra": "96823 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid",
            "value": 24.11,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "48878056 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid - ns/op",
            "value": 24.11,
            "unit": "ns/op",
            "extra": "48878056 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "48878056 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "48878056 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite",
            "value": 47170,
            "unit": "ns/op\t   34696 B/op\t     226 allocs/op",
            "extra": "23596 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite - ns/op",
            "value": 47170,
            "unit": "ns/op",
            "extra": "23596 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite - B/op",
            "value": 34696,
            "unit": "B/op",
            "extra": "23596 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite - allocs/op",
            "value": 226,
            "unit": "allocs/op",
            "extra": "23596 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite",
            "value": 189019,
            "unit": "ns/op\t   54680 B/op\t    2069 allocs/op",
            "extra": "6596 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite - ns/op",
            "value": 189019,
            "unit": "ns/op",
            "extra": "6596 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite - B/op",
            "value": 54680,
            "unit": "B/op",
            "extra": "6596 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite - allocs/op",
            "value": 2069,
            "unit": "allocs/op",
            "extra": "6596 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite",
            "value": 145597,
            "unit": "ns/op\t   61104 B/op\t     721 allocs/op",
            "extra": "7926 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite - ns/op",
            "value": 145597,
            "unit": "ns/op",
            "extra": "7926 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite - B/op",
            "value": 61104,
            "unit": "B/op",
            "extra": "7926 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite - allocs/op",
            "value": 721,
            "unit": "allocs/op",
            "extra": "7926 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDIATWrite",
            "value": 231380104,
            "unit": "1210428822609160249A094101Federal Reserve Bank   My Bank Name                   ",
            "extra": "101 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDIATWrite - 1210428822609160249A094101Federal",
            "value": 231380104,
            "unit": "1210428822609160249A094101Federal",
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
          }
        ]
      },
      {
        "commit": {
          "author": {
            "name": "moov-bot",
            "email": "oss@moov.io"
          },
          "committer": {
            "name": "moov-bot",
            "email": "oss@moov.io"
          },
          "id": "38284d4872094754f6bf6570a349893a98c6575b",
          "message": "chore: updating wasm webui [skip ci]",
          "timestamp": "2026-09-15T19:50:32Z",
          "url": "https://github.com/moov-io/ach/commit/38284d4872094754f6bf6570a349893a98c6575b"
        },
        "date": 1789526665228,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkAlphaFieldShort",
            "value": 1039,
            "unit": "ns/op\t      80 B/op\t       3 allocs/op",
            "extra": "1211766 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldShort - ns/op",
            "value": 1039,
            "unit": "ns/op",
            "extra": "1211766 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldShort - B/op",
            "value": 80,
            "unit": "B/op",
            "extra": "1211766 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldShort - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "1211766 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong",
            "value": 97.85,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "12157152 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong - ns/op",
            "value": 97.85,
            "unit": "ns/op",
            "extra": "12157152 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "12157152 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "12157152 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort",
            "value": 57.83,
            "unit": "ns/op\t      16 B/op\t       2 allocs/op",
            "extra": "20594568 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort - ns/op",
            "value": 57.83,
            "unit": "ns/op",
            "extra": "20594568 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "20594568 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "20594568 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong",
            "value": 25.85,
            "unit": "ns/op\t       8 B/op\t       1 allocs/op",
            "extra": "44405613 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong - ns/op",
            "value": 25.85,
            "unit": "ns/op",
            "extra": "44405613 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong - B/op",
            "value": 8,
            "unit": "B/op",
            "extra": "44405613 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "44405613 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField",
            "value": 15.27,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "79988598 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField - ns/op",
            "value": 15.27,
            "unit": "ns/op",
            "extra": "79988598 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "79988598 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "79988598 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField",
            "value": 5.988,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "201447102 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField - ns/op",
            "value": 5.988,
            "unit": "ns/op",
            "extra": "201447102 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "201447102 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "201447102 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles",
            "value": 240692,
            "unit": "ns/op\t   54143 B/op\t     306 allocs/op",
            "extra": "5134 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles - ns/op",
            "value": 240692,
            "unit": "ns/op",
            "extra": "5134 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles - B/op",
            "value": 54143,
            "unit": "B/op",
            "extra": "5134 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles - allocs/op",
            "value": 306,
            "unit": "allocs/op",
            "extra": "5134 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts",
            "value": 241868,
            "unit": "ns/op\t   54143 B/op\t     306 allocs/op",
            "extra": "5036 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts - ns/op",
            "value": 241868,
            "unit": "ns/op",
            "extra": "5036 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts - B/op",
            "value": 54143,
            "unit": "B/op",
            "extra": "5036 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts - allocs/op",
            "value": 306,
            "unit": "allocs/op",
            "extra": "5036 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir",
            "value": 147264,
            "unit": "ns/op\t   54557 B/op\t     310 allocs/op",
            "extra": "7556 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir - ns/op",
            "value": 147264,
            "unit": "ns/op",
            "extra": "7556 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir - B/op",
            "value": 54557,
            "unit": "B/op",
            "extra": "7556 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir - allocs/op",
            "value": 310,
            "unit": "allocs/op",
            "extra": "7556 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts",
            "value": 141910,
            "unit": "ns/op\t   54589 B/op\t     310 allocs/op",
            "extra": "7606 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts - ns/op",
            "value": 141910,
            "unit": "ns/op",
            "extra": "7606 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts - B/op",
            "value": 54589,
            "unit": "B/op",
            "extra": "7606 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts - allocs/op",
            "value": 310,
            "unit": "allocs/op",
            "extra": "7606 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups",
            "value": 311994,
            "unit": "ns/op\t   59885 B/op\t     366 allocs/op",
            "extra": "4017 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups - ns/op",
            "value": 311994,
            "unit": "ns/op",
            "extra": "4017 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups - B/op",
            "value": 59885,
            "unit": "B/op",
            "extra": "4017 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups - allocs/op",
            "value": 366,
            "unit": "allocs/op",
            "extra": "4017 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups",
            "value": 303734,
            "unit": "ns/op\t   59846 B/op\t     366 allocs/op",
            "extra": "4014 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups - ns/op",
            "value": 303734,
            "unit": "ns/op",
            "extra": "4014 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups - B/op",
            "value": 59846,
            "unit": "B/op",
            "extra": "4014 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups - allocs/op",
            "value": 366,
            "unit": "allocs/op",
            "extra": "4014 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups",
            "value": 295151,
            "unit": "ns/op\t   59859 B/op\t     366 allocs/op",
            "extra": "4215 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups - ns/op",
            "value": 295151,
            "unit": "ns/op",
            "extra": "4215 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups - B/op",
            "value": 59859,
            "unit": "B/op",
            "extra": "4215 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups - allocs/op",
            "value": 366,
            "unit": "allocs/op",
            "extra": "4215 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups",
            "value": 292193,
            "unit": "ns/op\t   59560 B/op\t     367 allocs/op",
            "extra": "4808 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups - ns/op",
            "value": 292193,
            "unit": "ns/op",
            "extra": "4808 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups - B/op",
            "value": 59560,
            "unit": "B/op",
            "extra": "4808 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups - allocs/op",
            "value": 367,
            "unit": "allocs/op",
            "extra": "4808 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader",
            "value": 245468906,
            "unit": "ns/op\t42450254 B/op\t  203655 allocs/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader - ns/op",
            "value": 245468906,
            "unit": "ns/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader - B/op",
            "value": 42450254,
            "unit": "B/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader - allocs/op",
            "value": 203655,
            "unit": "allocs/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator",
            "value": 77770814,
            "unit": "ns/op\t42424704 B/op\t  203600 allocs/op",
            "extra": "14 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator - ns/op",
            "value": 77770814,
            "unit": "ns/op",
            "extra": "14 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator - B/op",
            "value": 42424704,
            "unit": "B/op",
            "extra": "14 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator - allocs/op",
            "value": 203600,
            "unit": "allocs/op",
            "extra": "14 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader",
            "value": 443686065,
            "unit": "ns/op\t62131717 B/op\t  905794 allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader - ns/op",
            "value": 443686065,
            "unit": "ns/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader - B/op",
            "value": 62131717,
            "unit": "B/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader - allocs/op",
            "value": 905794,
            "unit": "allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator",
            "value": 144755171,
            "unit": "ns/op\t60508790 B/op\t  705839 allocs/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator - ns/op",
            "value": 144755171,
            "unit": "ns/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator - B/op",
            "value": 60508790,
            "unit": "B/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator - allocs/op",
            "value": 705839,
            "unit": "allocs/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader",
            "value": 1283056718,
            "unit": "ns/op\t215541432 B/op\t 1042868 allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader - ns/op",
            "value": 1283056718,
            "unit": "ns/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader - B/op",
            "value": 215541432,
            "unit": "B/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader - allocs/op",
            "value": 1042868,
            "unit": "allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator",
            "value": 383817403,
            "unit": "ns/op\t215424805 B/op\t 1042800 allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator - ns/op",
            "value": 383817403,
            "unit": "ns/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator - B/op",
            "value": 215424805,
            "unit": "B/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator - allocs/op",
            "value": 1042800,
            "unit": "allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader",
            "value": 2177834805,
            "unit": "ns/op\t313545480 B/op\t 4568039 allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader - ns/op",
            "value": 2177834805,
            "unit": "ns/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader - B/op",
            "value": 313545480,
            "unit": "B/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader - allocs/op",
            "value": 4568039,
            "unit": "allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator",
            "value": 725925171,
            "unit": "ns/op\t305429996 B/op\t 3568070 allocs/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator - ns/op",
            "value": 725925171,
            "unit": "ns/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator - B/op",
            "value": 305429996,
            "unit": "B/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator - allocs/op",
            "value": 3568070,
            "unit": "allocs/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead",
            "value": 40838,
            "unit": "ns/op\t   23312 B/op\t     114 allocs/op",
            "extra": "29539 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead - ns/op",
            "value": 40838,
            "unit": "ns/op",
            "extra": "29539 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead - B/op",
            "value": 23312,
            "unit": "B/op",
            "extra": "29539 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead - allocs/op",
            "value": 114,
            "unit": "allocs/op",
            "extra": "29539 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead",
            "value": 70194,
            "unit": "ns/op\t   26920 B/op\t     159 allocs/op",
            "extra": "16327 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead - ns/op",
            "value": 70194,
            "unit": "ns/op",
            "extra": "16327 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead - B/op",
            "value": 26920,
            "unit": "B/op",
            "extra": "16327 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead - allocs/op",
            "value": 159,
            "unit": "allocs/op",
            "extra": "16327 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead",
            "value": 30532,
            "unit": "ns/op\t   21848 B/op\t      77 allocs/op",
            "extra": "39822 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead - ns/op",
            "value": 30532,
            "unit": "ns/op",
            "extra": "39822 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead - B/op",
            "value": 21848,
            "unit": "B/op",
            "extra": "39822 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead - allocs/op",
            "value": 77,
            "unit": "allocs/op",
            "extra": "39822 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead",
            "value": 266878,
            "unit": "ns/op\t   56353 B/op\t     552 allocs/op",
            "extra": "4476 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead - ns/op",
            "value": 266878,
            "unit": "ns/op",
            "extra": "4476 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead - B/op",
            "value": 56353,
            "unit": "B/op",
            "extra": "4476 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead - allocs/op",
            "value": 552,
            "unit": "allocs/op",
            "extra": "4476 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2",
            "value": 267417,
            "unit": "ns/op\t   56353 B/op\t     552 allocs/op",
            "extra": "4670 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2 - ns/op",
            "value": 267417,
            "unit": "ns/op",
            "extra": "4670 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2 - B/op",
            "value": 56353,
            "unit": "B/op",
            "extra": "4670 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2 - allocs/op",
            "value": 552,
            "unit": "allocs/op",
            "extra": "4670 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3",
            "value": 101618,
            "unit": "ns/op\t   29752 B/op\t     264 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3 - ns/op",
            "value": 101618,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3 - B/op",
            "value": 29752,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3 - allocs/op",
            "value": 264,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile",
            "value": 13886,
            "unit": "ns/op\t   10760 B/op\t     131 allocs/op",
            "extra": "94444 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile - ns/op",
            "value": 13886,
            "unit": "ns/op",
            "extra": "94444 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile - B/op",
            "value": 10760,
            "unit": "B/op",
            "extra": "94444 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile - allocs/op",
            "value": 131,
            "unit": "allocs/op",
            "extra": "94444 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid",
            "value": 25.38,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "45284889 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid - ns/op",
            "value": 25.38,
            "unit": "ns/op",
            "extra": "45284889 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "45284889 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "45284889 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite",
            "value": 48346,
            "unit": "ns/op\t   34696 B/op\t     226 allocs/op",
            "extra": "23298 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite - ns/op",
            "value": 48346,
            "unit": "ns/op",
            "extra": "23298 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite - B/op",
            "value": 34696,
            "unit": "B/op",
            "extra": "23298 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite - allocs/op",
            "value": 226,
            "unit": "allocs/op",
            "extra": "23298 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite",
            "value": 212759,
            "unit": "ns/op\t   54680 B/op\t    2069 allocs/op",
            "extra": "5772 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite - ns/op",
            "value": 212759,
            "unit": "ns/op",
            "extra": "5772 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite - B/op",
            "value": 54680,
            "unit": "B/op",
            "extra": "5772 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite - allocs/op",
            "value": 2069,
            "unit": "allocs/op",
            "extra": "5772 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite",
            "value": 150710,
            "unit": "ns/op\t   61104 B/op\t     721 allocs/op",
            "extra": "7651 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite - ns/op",
            "value": 150710,
            "unit": "ns/op",
            "extra": "7651 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite - B/op",
            "value": 61104,
            "unit": "B/op",
            "extra": "7651 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite - allocs/op",
            "value": 721,
            "unit": "allocs/op",
            "extra": "7651 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDIATWrite",
            "value": 231380104,
            "unit": "1210428822609170244A094101Federal Reserve Bank   My Bank Name                   ",
            "extra": "101 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDIATWrite - 1210428822609170244A094101Federal",
            "value": 231380104,
            "unit": "1210428822609170244A094101Federal",
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
          }
        ]
      },
      {
        "commit": {
          "author": {
            "name": "moov-bot",
            "email": "oss@moov.io"
          },
          "committer": {
            "name": "moov-bot",
            "email": "oss@moov.io"
          },
          "id": "38284d4872094754f6bf6570a349893a98c6575b",
          "message": "chore: updating wasm webui [skip ci]",
          "timestamp": "2026-09-15T19:50:32Z",
          "url": "https://github.com/moov-io/ach/commit/38284d4872094754f6bf6570a349893a98c6575b"
        },
        "date": 1789613313205,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkAlphaFieldShort",
            "value": 989.3,
            "unit": "ns/op\t      80 B/op\t       3 allocs/op",
            "extra": "1218429 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldShort - ns/op",
            "value": 989.3,
            "unit": "ns/op",
            "extra": "1218429 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldShort - B/op",
            "value": 80,
            "unit": "B/op",
            "extra": "1218429 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldShort - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "1218429 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong",
            "value": 97.64,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "12188973 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong - ns/op",
            "value": 97.64,
            "unit": "ns/op",
            "extra": "12188973 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "12188973 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "12188973 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort",
            "value": 56.06,
            "unit": "ns/op\t      16 B/op\t       2 allocs/op",
            "extra": "20458972 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort - ns/op",
            "value": 56.06,
            "unit": "ns/op",
            "extra": "20458972 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "20458972 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "20458972 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong",
            "value": 26.36,
            "unit": "ns/op\t       8 B/op\t       1 allocs/op",
            "extra": "42937227 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong - ns/op",
            "value": 26.36,
            "unit": "ns/op",
            "extra": "42937227 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong - B/op",
            "value": 8,
            "unit": "B/op",
            "extra": "42937227 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "42937227 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField",
            "value": 14.99,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "79912086 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField - ns/op",
            "value": 14.99,
            "unit": "ns/op",
            "extra": "79912086 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "79912086 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "79912086 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField",
            "value": 5.955,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "202686501 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField - ns/op",
            "value": 5.955,
            "unit": "ns/op",
            "extra": "202686501 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "202686501 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "202686501 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles",
            "value": 262769,
            "unit": "ns/op\t   54141 B/op\t     306 allocs/op",
            "extra": "5212 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles - ns/op",
            "value": 262769,
            "unit": "ns/op",
            "extra": "5212 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles - B/op",
            "value": 54141,
            "unit": "B/op",
            "extra": "5212 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles - allocs/op",
            "value": 306,
            "unit": "allocs/op",
            "extra": "5212 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts",
            "value": 248240,
            "unit": "ns/op\t   54162 B/op\t     306 allocs/op",
            "extra": "4833 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts - ns/op",
            "value": 248240,
            "unit": "ns/op",
            "extra": "4833 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts - B/op",
            "value": 54162,
            "unit": "B/op",
            "extra": "4833 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts - allocs/op",
            "value": 306,
            "unit": "allocs/op",
            "extra": "4833 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir",
            "value": 147990,
            "unit": "ns/op\t   54550 B/op\t     310 allocs/op",
            "extra": "8620 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir - ns/op",
            "value": 147990,
            "unit": "ns/op",
            "extra": "8620 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir - B/op",
            "value": 54550,
            "unit": "B/op",
            "extra": "8620 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir - allocs/op",
            "value": 310,
            "unit": "allocs/op",
            "extra": "8620 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts",
            "value": 147952,
            "unit": "ns/op\t   54589 B/op\t     310 allocs/op",
            "extra": "7041 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts - ns/op",
            "value": 147952,
            "unit": "ns/op",
            "extra": "7041 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts - B/op",
            "value": 54589,
            "unit": "B/op",
            "extra": "7041 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts - allocs/op",
            "value": 310,
            "unit": "allocs/op",
            "extra": "7041 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups",
            "value": 307912,
            "unit": "ns/op\t   59877 B/op\t     366 allocs/op",
            "extra": "3994 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups - ns/op",
            "value": 307912,
            "unit": "ns/op",
            "extra": "3994 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups - B/op",
            "value": 59877,
            "unit": "B/op",
            "extra": "3994 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups - allocs/op",
            "value": 366,
            "unit": "allocs/op",
            "extra": "3994 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups",
            "value": 300705,
            "unit": "ns/op\t   59878 B/op\t     366 allocs/op",
            "extra": "4251 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups - ns/op",
            "value": 300705,
            "unit": "ns/op",
            "extra": "4251 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups - B/op",
            "value": 59878,
            "unit": "B/op",
            "extra": "4251 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups - allocs/op",
            "value": 366,
            "unit": "allocs/op",
            "extra": "4251 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups",
            "value": 298148,
            "unit": "ns/op\t   59944 B/op\t     366 allocs/op",
            "extra": "4063 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups - ns/op",
            "value": 298148,
            "unit": "ns/op",
            "extra": "4063 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups - B/op",
            "value": 59944,
            "unit": "B/op",
            "extra": "4063 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups - allocs/op",
            "value": 366,
            "unit": "allocs/op",
            "extra": "4063 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups",
            "value": 290355,
            "unit": "ns/op\t   59550 B/op\t     367 allocs/op",
            "extra": "4902 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups - ns/op",
            "value": 290355,
            "unit": "ns/op",
            "extra": "4902 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups - B/op",
            "value": 59550,
            "unit": "B/op",
            "extra": "4902 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups - allocs/op",
            "value": 367,
            "unit": "allocs/op",
            "extra": "4902 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader",
            "value": 246557179,
            "unit": "ns/op\t42450132 B/op\t  203653 allocs/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader - ns/op",
            "value": 246557179,
            "unit": "ns/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader - B/op",
            "value": 42450132,
            "unit": "B/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader - allocs/op",
            "value": 203653,
            "unit": "allocs/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator",
            "value": 78882893,
            "unit": "ns/op\t42424898 B/op\t  203602 allocs/op",
            "extra": "14 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator - ns/op",
            "value": 78882893,
            "unit": "ns/op",
            "extra": "14 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator - B/op",
            "value": 42424898,
            "unit": "B/op",
            "extra": "14 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator - allocs/op",
            "value": 203602,
            "unit": "allocs/op",
            "extra": "14 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader",
            "value": 447538351,
            "unit": "ns/op\t62131616 B/op\t  905793 allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader - ns/op",
            "value": 447538351,
            "unit": "ns/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader - B/op",
            "value": 62131616,
            "unit": "B/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader - allocs/op",
            "value": 905793,
            "unit": "allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator",
            "value": 147548091,
            "unit": "ns/op\t60508708 B/op\t  705838 allocs/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator - ns/op",
            "value": 147548091,
            "unit": "ns/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator - B/op",
            "value": 60508708,
            "unit": "B/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator - allocs/op",
            "value": 705838,
            "unit": "allocs/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader",
            "value": 1304479164,
            "unit": "ns/op\t215541288 B/op\t 1042866 allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader - ns/op",
            "value": 1304479164,
            "unit": "ns/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader - B/op",
            "value": 215541288,
            "unit": "B/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader - allocs/op",
            "value": 1042866,
            "unit": "allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator",
            "value": 396313298,
            "unit": "ns/op\t215424592 B/op\t 1042797 allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator - ns/op",
            "value": 396313298,
            "unit": "ns/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator - B/op",
            "value": 215424592,
            "unit": "B/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator - allocs/op",
            "value": 1042797,
            "unit": "allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader",
            "value": 2196389270,
            "unit": "ns/op\t313545480 B/op\t 4568039 allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader - ns/op",
            "value": 2196389270,
            "unit": "ns/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader - B/op",
            "value": 313545480,
            "unit": "B/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader - allocs/op",
            "value": 4568039,
            "unit": "allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator",
            "value": 743302380,
            "unit": "ns/op\t305430332 B/op\t 3568074 allocs/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator - ns/op",
            "value": 743302380,
            "unit": "ns/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator - B/op",
            "value": 305430332,
            "unit": "B/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator - allocs/op",
            "value": 3568074,
            "unit": "allocs/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead",
            "value": 41591,
            "unit": "ns/op\t   23312 B/op\t     114 allocs/op",
            "extra": "29146 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead - ns/op",
            "value": 41591,
            "unit": "ns/op",
            "extra": "29146 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead - B/op",
            "value": 23312,
            "unit": "B/op",
            "extra": "29146 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead - allocs/op",
            "value": 114,
            "unit": "allocs/op",
            "extra": "29146 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead",
            "value": 71582,
            "unit": "ns/op\t   26920 B/op\t     159 allocs/op",
            "extra": "16227 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead - ns/op",
            "value": 71582,
            "unit": "ns/op",
            "extra": "16227 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead - B/op",
            "value": 26920,
            "unit": "B/op",
            "extra": "16227 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead - allocs/op",
            "value": 159,
            "unit": "allocs/op",
            "extra": "16227 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead",
            "value": 31142,
            "unit": "ns/op\t   21848 B/op\t      77 allocs/op",
            "extra": "39888 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead - ns/op",
            "value": 31142,
            "unit": "ns/op",
            "extra": "39888 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead - B/op",
            "value": 21848,
            "unit": "B/op",
            "extra": "39888 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead - allocs/op",
            "value": 77,
            "unit": "allocs/op",
            "extra": "39888 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead",
            "value": 270835,
            "unit": "ns/op\t   56353 B/op\t     552 allocs/op",
            "extra": "4530 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead - ns/op",
            "value": 270835,
            "unit": "ns/op",
            "extra": "4530 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead - B/op",
            "value": 56353,
            "unit": "B/op",
            "extra": "4530 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead - allocs/op",
            "value": 552,
            "unit": "allocs/op",
            "extra": "4530 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2",
            "value": 270185,
            "unit": "ns/op\t   56353 B/op\t     552 allocs/op",
            "extra": "4557 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2 - ns/op",
            "value": 270185,
            "unit": "ns/op",
            "extra": "4557 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2 - B/op",
            "value": 56353,
            "unit": "B/op",
            "extra": "4557 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2 - allocs/op",
            "value": 552,
            "unit": "allocs/op",
            "extra": "4557 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3",
            "value": 102903,
            "unit": "ns/op\t   29752 B/op\t     264 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3 - ns/op",
            "value": 102903,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3 - B/op",
            "value": 29752,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3 - allocs/op",
            "value": 264,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile",
            "value": 14526,
            "unit": "ns/op\t   10760 B/op\t     131 allocs/op",
            "extra": "92566 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile - ns/op",
            "value": 14526,
            "unit": "ns/op",
            "extra": "92566 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile - B/op",
            "value": 10760,
            "unit": "B/op",
            "extra": "92566 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile - allocs/op",
            "value": 131,
            "unit": "allocs/op",
            "extra": "92566 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid",
            "value": 25.43,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "46266044 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid - ns/op",
            "value": 25.43,
            "unit": "ns/op",
            "extra": "46266044 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "46266044 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "46266044 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite",
            "value": 49458,
            "unit": "ns/op\t   34696 B/op\t     226 allocs/op",
            "extra": "22632 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite - ns/op",
            "value": 49458,
            "unit": "ns/op",
            "extra": "22632 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite - B/op",
            "value": 34696,
            "unit": "B/op",
            "extra": "22632 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite - allocs/op",
            "value": 226,
            "unit": "allocs/op",
            "extra": "22632 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite",
            "value": 211458,
            "unit": "ns/op\t   54680 B/op\t    2069 allocs/op",
            "extra": "5946 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite - ns/op",
            "value": 211458,
            "unit": "ns/op",
            "extra": "5946 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite - B/op",
            "value": 54680,
            "unit": "B/op",
            "extra": "5946 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite - allocs/op",
            "value": 2069,
            "unit": "allocs/op",
            "extra": "5946 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite",
            "value": 152083,
            "unit": "ns/op\t   61104 B/op\t     721 allocs/op",
            "extra": "7467 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite - ns/op",
            "value": 152083,
            "unit": "ns/op",
            "extra": "7467 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite - B/op",
            "value": 61104,
            "unit": "B/op",
            "extra": "7467 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite - allocs/op",
            "value": 721,
            "unit": "allocs/op",
            "extra": "7467 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDIATWrite",
            "value": 231380104,
            "unit": "1210428822609180248A094101Federal Reserve Bank   My Bank Name                   ",
            "extra": "101 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDIATWrite - 1210428822609180248A094101Federal",
            "value": 231380104,
            "unit": "1210428822609180248A094101Federal",
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
          }
        ]
      },
      {
        "commit": {
          "author": {
            "name": "moov-bot",
            "email": "oss@moov.io"
          },
          "committer": {
            "name": "moov-bot",
            "email": "oss@moov.io"
          },
          "id": "d6eb332385a106c242d3e43044aecc81a2d7a28e",
          "message": "chore: updating wasm webui [skip ci]",
          "timestamp": "2026-09-17T14:51:32Z",
          "url": "https://github.com/moov-io/ach/commit/d6eb332385a106c242d3e43044aecc81a2d7a28e"
        },
        "date": 1789698917363,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkAlphaFieldShort",
            "value": 977.1,
            "unit": "ns/op\t      80 B/op\t       3 allocs/op",
            "extra": "1084400 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldShort - ns/op",
            "value": 977.1,
            "unit": "ns/op",
            "extra": "1084400 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldShort - B/op",
            "value": 80,
            "unit": "B/op",
            "extra": "1084400 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldShort - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "1084400 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong",
            "value": 97.8,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "12186574 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong - ns/op",
            "value": 97.8,
            "unit": "ns/op",
            "extra": "12186574 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "12186574 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "12186574 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort",
            "value": 59.76,
            "unit": "ns/op\t      16 B/op\t       2 allocs/op",
            "extra": "19513970 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort - ns/op",
            "value": 59.76,
            "unit": "ns/op",
            "extra": "19513970 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "19513970 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "19513970 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong",
            "value": 25.4,
            "unit": "ns/op\t       8 B/op\t       1 allocs/op",
            "extra": "44931662 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong - ns/op",
            "value": 25.4,
            "unit": "ns/op",
            "extra": "44931662 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong - B/op",
            "value": 8,
            "unit": "B/op",
            "extra": "44931662 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "44931662 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField",
            "value": 14.96,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "81996679 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField - ns/op",
            "value": 14.96,
            "unit": "ns/op",
            "extra": "81996679 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "81996679 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "81996679 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField",
            "value": 5.616,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "213612170 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField - ns/op",
            "value": 5.616,
            "unit": "ns/op",
            "extra": "213612170 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "213612170 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "213612170 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles",
            "value": 259415,
            "unit": "ns/op\t   54145 B/op\t     306 allocs/op",
            "extra": "5114 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles - ns/op",
            "value": 259415,
            "unit": "ns/op",
            "extra": "5114 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles - B/op",
            "value": 54145,
            "unit": "B/op",
            "extra": "5114 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles - allocs/op",
            "value": 306,
            "unit": "allocs/op",
            "extra": "5114 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts",
            "value": 250674,
            "unit": "ns/op\t   54160 B/op\t     306 allocs/op",
            "extra": "4700 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts - ns/op",
            "value": 250674,
            "unit": "ns/op",
            "extra": "4700 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts - B/op",
            "value": 54160,
            "unit": "B/op",
            "extra": "4700 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts - allocs/op",
            "value": 306,
            "unit": "allocs/op",
            "extra": "4700 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir",
            "value": 142073,
            "unit": "ns/op\t   54546 B/op\t     310 allocs/op",
            "extra": "7935 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir - ns/op",
            "value": 142073,
            "unit": "ns/op",
            "extra": "7935 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir - B/op",
            "value": 54546,
            "unit": "B/op",
            "extra": "7935 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir - allocs/op",
            "value": 310,
            "unit": "allocs/op",
            "extra": "7935 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts",
            "value": 145122,
            "unit": "ns/op\t   54586 B/op\t     310 allocs/op",
            "extra": "7038 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts - ns/op",
            "value": 145122,
            "unit": "ns/op",
            "extra": "7038 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts - B/op",
            "value": 54586,
            "unit": "B/op",
            "extra": "7038 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts - allocs/op",
            "value": 310,
            "unit": "allocs/op",
            "extra": "7038 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups",
            "value": 312388,
            "unit": "ns/op\t   59867 B/op\t     366 allocs/op",
            "extra": "3866 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups - ns/op",
            "value": 312388,
            "unit": "ns/op",
            "extra": "3866 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups - B/op",
            "value": 59867,
            "unit": "B/op",
            "extra": "3866 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups - allocs/op",
            "value": 366,
            "unit": "allocs/op",
            "extra": "3866 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups",
            "value": 315208,
            "unit": "ns/op\t   59886 B/op\t     366 allocs/op",
            "extra": "4099 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups - ns/op",
            "value": 315208,
            "unit": "ns/op",
            "extra": "4099 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups - B/op",
            "value": 59886,
            "unit": "B/op",
            "extra": "4099 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups - allocs/op",
            "value": 366,
            "unit": "allocs/op",
            "extra": "4099 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups",
            "value": 304615,
            "unit": "ns/op\t   59903 B/op\t     366 allocs/op",
            "extra": "4146 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups - ns/op",
            "value": 304615,
            "unit": "ns/op",
            "extra": "4146 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups - B/op",
            "value": 59903,
            "unit": "B/op",
            "extra": "4146 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups - allocs/op",
            "value": 366,
            "unit": "allocs/op",
            "extra": "4146 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups",
            "value": 293571,
            "unit": "ns/op\t   59576 B/op\t     367 allocs/op",
            "extra": "4834 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups - ns/op",
            "value": 293571,
            "unit": "ns/op",
            "extra": "4834 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups - B/op",
            "value": 59576,
            "unit": "B/op",
            "extra": "4834 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups - allocs/op",
            "value": 367,
            "unit": "allocs/op",
            "extra": "4834 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader",
            "value": 244584228,
            "unit": "ns/op\t42450299 B/op\t  203656 allocs/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader - ns/op",
            "value": 244584228,
            "unit": "ns/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader - B/op",
            "value": 42450299,
            "unit": "B/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader - allocs/op",
            "value": 203656,
            "unit": "allocs/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator",
            "value": 77616066,
            "unit": "ns/op\t42424749 B/op\t  203601 allocs/op",
            "extra": "14 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator - ns/op",
            "value": 77616066,
            "unit": "ns/op",
            "extra": "14 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator - B/op",
            "value": 42424749,
            "unit": "B/op",
            "extra": "14 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator - allocs/op",
            "value": 203601,
            "unit": "allocs/op",
            "extra": "14 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader",
            "value": 447518355,
            "unit": "ns/op\t62131690 B/op\t  905794 allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader - ns/op",
            "value": 447518355,
            "unit": "ns/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader - B/op",
            "value": 62131690,
            "unit": "B/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader - allocs/op",
            "value": 905794,
            "unit": "allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator",
            "value": 148219227,
            "unit": "ns/op\t60508797 B/op\t  705839 allocs/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator - ns/op",
            "value": 148219227,
            "unit": "ns/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator - B/op",
            "value": 60508797,
            "unit": "B/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator - allocs/op",
            "value": 705839,
            "unit": "allocs/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader",
            "value": 1291341950,
            "unit": "ns/op\t215541064 B/op\t 1042863 allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader - ns/op",
            "value": 1291341950,
            "unit": "ns/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader - B/op",
            "value": 215541064,
            "unit": "B/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader - allocs/op",
            "value": 1042863,
            "unit": "allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator",
            "value": 391546141,
            "unit": "ns/op\t215424704 B/op\t 1042799 allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator - ns/op",
            "value": 391546141,
            "unit": "ns/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator - B/op",
            "value": 215424704,
            "unit": "B/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator - allocs/op",
            "value": 1042799,
            "unit": "allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader",
            "value": 2203817771,
            "unit": "ns/op\t313545480 B/op\t 4568039 allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader - ns/op",
            "value": 2203817771,
            "unit": "ns/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader - B/op",
            "value": 313545480,
            "unit": "B/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader - allocs/op",
            "value": 4568039,
            "unit": "allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator",
            "value": 743548352,
            "unit": "ns/op\t305430068 B/op\t 3568071 allocs/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator - ns/op",
            "value": 743548352,
            "unit": "ns/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator - B/op",
            "value": 305430068,
            "unit": "B/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator - allocs/op",
            "value": 3568071,
            "unit": "allocs/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead",
            "value": 41208,
            "unit": "ns/op\t   23312 B/op\t     114 allocs/op",
            "extra": "29512 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead - ns/op",
            "value": 41208,
            "unit": "ns/op",
            "extra": "29512 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead - B/op",
            "value": 23312,
            "unit": "B/op",
            "extra": "29512 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead - allocs/op",
            "value": 114,
            "unit": "allocs/op",
            "extra": "29512 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead",
            "value": 70653,
            "unit": "ns/op\t   26920 B/op\t     159 allocs/op",
            "extra": "16075 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead - ns/op",
            "value": 70653,
            "unit": "ns/op",
            "extra": "16075 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead - B/op",
            "value": 26920,
            "unit": "B/op",
            "extra": "16075 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead - allocs/op",
            "value": 159,
            "unit": "allocs/op",
            "extra": "16075 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead",
            "value": 30976,
            "unit": "ns/op\t   21848 B/op\t      77 allocs/op",
            "extra": "40188 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead - ns/op",
            "value": 30976,
            "unit": "ns/op",
            "extra": "40188 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead - B/op",
            "value": 21848,
            "unit": "B/op",
            "extra": "40188 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead - allocs/op",
            "value": 77,
            "unit": "allocs/op",
            "extra": "40188 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead",
            "value": 269580,
            "unit": "ns/op\t   56353 B/op\t     552 allocs/op",
            "extra": "4510 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead - ns/op",
            "value": 269580,
            "unit": "ns/op",
            "extra": "4510 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead - B/op",
            "value": 56353,
            "unit": "B/op",
            "extra": "4510 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead - allocs/op",
            "value": 552,
            "unit": "allocs/op",
            "extra": "4510 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2",
            "value": 270115,
            "unit": "ns/op\t   56353 B/op\t     552 allocs/op",
            "extra": "4368 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2 - ns/op",
            "value": 270115,
            "unit": "ns/op",
            "extra": "4368 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2 - B/op",
            "value": 56353,
            "unit": "B/op",
            "extra": "4368 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2 - allocs/op",
            "value": 552,
            "unit": "allocs/op",
            "extra": "4368 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3",
            "value": 102722,
            "unit": "ns/op\t   29752 B/op\t     264 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3 - ns/op",
            "value": 102722,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3 - B/op",
            "value": 29752,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3 - allocs/op",
            "value": 264,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile",
            "value": 13782,
            "unit": "ns/op\t   10760 B/op\t     131 allocs/op",
            "extra": "93872 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile - ns/op",
            "value": 13782,
            "unit": "ns/op",
            "extra": "93872 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile - B/op",
            "value": 10760,
            "unit": "B/op",
            "extra": "93872 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile - allocs/op",
            "value": 131,
            "unit": "allocs/op",
            "extra": "93872 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid",
            "value": 25.05,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "47558137 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid - ns/op",
            "value": 25.05,
            "unit": "ns/op",
            "extra": "47558137 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "47558137 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "47558137 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite",
            "value": 48506,
            "unit": "ns/op\t   34696 B/op\t     226 allocs/op",
            "extra": "23079 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite - ns/op",
            "value": 48506,
            "unit": "ns/op",
            "extra": "23079 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite - B/op",
            "value": 34696,
            "unit": "B/op",
            "extra": "23079 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite - allocs/op",
            "value": 226,
            "unit": "allocs/op",
            "extra": "23079 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite",
            "value": 211419,
            "unit": "ns/op\t   54680 B/op\t    2069 allocs/op",
            "extra": "5878 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite - ns/op",
            "value": 211419,
            "unit": "ns/op",
            "extra": "5878 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite - B/op",
            "value": 54680,
            "unit": "B/op",
            "extra": "5878 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite - allocs/op",
            "value": 2069,
            "unit": "allocs/op",
            "extra": "5878 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite",
            "value": 152166,
            "unit": "ns/op\t   61104 B/op\t     721 allocs/op",
            "extra": "7597 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite - ns/op",
            "value": 152166,
            "unit": "ns/op",
            "extra": "7597 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite - B/op",
            "value": 61104,
            "unit": "B/op",
            "extra": "7597 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite - allocs/op",
            "value": 721,
            "unit": "allocs/op",
            "extra": "7597 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDIATWrite",
            "value": 231380104,
            "unit": "1210428822609190235A094101Federal Reserve Bank   My Bank Name                   ",
            "extra": "101 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDIATWrite - 1210428822609190235A094101Federal",
            "value": 231380104,
            "unit": "1210428822609190235A094101Federal",
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
          }
        ]
      },
      {
        "commit": {
          "author": {
            "name": "renovate[bot]",
            "username": "renovate[bot]",
            "email": "29139614+renovate[bot]@users.noreply.github.com"
          },
          "committer": {
            "name": "GitHub",
            "username": "web-flow",
            "email": "noreply@github.com"
          },
          "id": "f36ebb7ef2e6d678ef95ca6278d2b1f2deae4299",
          "message": "chore(deps): update github/codeql-action action to v4.38.1 (#1863)\n\nCo-authored-by: renovate[bot] <29139614+renovate[bot]@users.noreply.github.com>",
          "timestamp": "2026-09-19T02:30:43Z",
          "url": "https://github.com/moov-io/ach/commit/f36ebb7ef2e6d678ef95ca6278d2b1f2deae4299"
        },
        "date": 1789785266781,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkAlphaFieldShort",
            "value": 983.4,
            "unit": "ns/op\t      80 B/op\t       3 allocs/op",
            "extra": "1213219 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldShort - ns/op",
            "value": 983.4,
            "unit": "ns/op",
            "extra": "1213219 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldShort - B/op",
            "value": 80,
            "unit": "B/op",
            "extra": "1213219 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldShort - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "1213219 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong",
            "value": 98.32,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "12110565 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong - ns/op",
            "value": 98.32,
            "unit": "ns/op",
            "extra": "12110565 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "12110565 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "12110565 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort",
            "value": 59.53,
            "unit": "ns/op\t      16 B/op\t       2 allocs/op",
            "extra": "19210101 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort - ns/op",
            "value": 59.53,
            "unit": "ns/op",
            "extra": "19210101 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "19210101 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "19210101 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong",
            "value": 25.64,
            "unit": "ns/op\t       8 B/op\t       1 allocs/op",
            "extra": "44306924 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong - ns/op",
            "value": 25.64,
            "unit": "ns/op",
            "extra": "44306924 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong - B/op",
            "value": 8,
            "unit": "B/op",
            "extra": "44306924 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "44306924 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField",
            "value": 14.96,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "81932872 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField - ns/op",
            "value": 14.96,
            "unit": "ns/op",
            "extra": "81932872 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "81932872 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "81932872 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField",
            "value": 5.607,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "212988339 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField - ns/op",
            "value": 5.607,
            "unit": "ns/op",
            "extra": "212988339 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "212988339 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "212988339 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles",
            "value": 252814,
            "unit": "ns/op\t   54145 B/op\t     306 allocs/op",
            "extra": "5112 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles - ns/op",
            "value": 252814,
            "unit": "ns/op",
            "extra": "5112 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles - B/op",
            "value": 54145,
            "unit": "B/op",
            "extra": "5112 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles - allocs/op",
            "value": 306,
            "unit": "allocs/op",
            "extra": "5112 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts",
            "value": 256172,
            "unit": "ns/op\t   54157 B/op\t     306 allocs/op",
            "extra": "4429 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts - ns/op",
            "value": 256172,
            "unit": "ns/op",
            "extra": "4429 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts - B/op",
            "value": 54157,
            "unit": "B/op",
            "extra": "4429 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts - allocs/op",
            "value": 306,
            "unit": "allocs/op",
            "extra": "4429 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir",
            "value": 146937,
            "unit": "ns/op\t   54551 B/op\t     310 allocs/op",
            "extra": "7876 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir - ns/op",
            "value": 146937,
            "unit": "ns/op",
            "extra": "7876 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir - B/op",
            "value": 54551,
            "unit": "B/op",
            "extra": "7876 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir - allocs/op",
            "value": 310,
            "unit": "allocs/op",
            "extra": "7876 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts",
            "value": 142366,
            "unit": "ns/op\t   54578 B/op\t     310 allocs/op",
            "extra": "7700 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts - ns/op",
            "value": 142366,
            "unit": "ns/op",
            "extra": "7700 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts - B/op",
            "value": 54578,
            "unit": "B/op",
            "extra": "7700 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts - allocs/op",
            "value": 310,
            "unit": "allocs/op",
            "extra": "7700 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups",
            "value": 317861,
            "unit": "ns/op\t   59882 B/op\t     366 allocs/op",
            "extra": "4141 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups - ns/op",
            "value": 317861,
            "unit": "ns/op",
            "extra": "4141 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups - B/op",
            "value": 59882,
            "unit": "B/op",
            "extra": "4141 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups - allocs/op",
            "value": 366,
            "unit": "allocs/op",
            "extra": "4141 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups",
            "value": 325138,
            "unit": "ns/op\t   59907 B/op\t     366 allocs/op",
            "extra": "4149 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups - ns/op",
            "value": 325138,
            "unit": "ns/op",
            "extra": "4149 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups - B/op",
            "value": 59907,
            "unit": "B/op",
            "extra": "4149 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups - allocs/op",
            "value": 366,
            "unit": "allocs/op",
            "extra": "4149 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups",
            "value": 299463,
            "unit": "ns/op\t   59851 B/op\t     366 allocs/op",
            "extra": "3987 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups - ns/op",
            "value": 299463,
            "unit": "ns/op",
            "extra": "3987 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups - B/op",
            "value": 59851,
            "unit": "B/op",
            "extra": "3987 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups - allocs/op",
            "value": 366,
            "unit": "allocs/op",
            "extra": "3987 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups",
            "value": 288370,
            "unit": "ns/op\t   59569 B/op\t     367 allocs/op",
            "extra": "4806 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups - ns/op",
            "value": 288370,
            "unit": "ns/op",
            "extra": "4806 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups - B/op",
            "value": 59569,
            "unit": "B/op",
            "extra": "4806 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups - allocs/op",
            "value": 367,
            "unit": "allocs/op",
            "extra": "4806 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader",
            "value": 258666610,
            "unit": "ns/op\t42450798 B/op\t  203656 allocs/op",
            "extra": "4 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader - ns/op",
            "value": 258666610,
            "unit": "ns/op",
            "extra": "4 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader - B/op",
            "value": 42450798,
            "unit": "B/op",
            "extra": "4 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader - allocs/op",
            "value": 203656,
            "unit": "allocs/op",
            "extra": "4 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator",
            "value": 79237772,
            "unit": "ns/op\t42424664 B/op\t  203600 allocs/op",
            "extra": "13 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator - ns/op",
            "value": 79237772,
            "unit": "ns/op",
            "extra": "13 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator - B/op",
            "value": 42424664,
            "unit": "B/op",
            "extra": "13 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator - allocs/op",
            "value": 203600,
            "unit": "allocs/op",
            "extra": "13 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader",
            "value": 451616555,
            "unit": "ns/op\t62131664 B/op\t  905793 allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader - ns/op",
            "value": 451616555,
            "unit": "ns/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader - B/op",
            "value": 62131664,
            "unit": "B/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader - allocs/op",
            "value": 905793,
            "unit": "allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator",
            "value": 149884000,
            "unit": "ns/op\t60508793 B/op\t  705839 allocs/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator - ns/op",
            "value": 149884000,
            "unit": "ns/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator - B/op",
            "value": 60508793,
            "unit": "B/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator - allocs/op",
            "value": 705839,
            "unit": "allocs/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader",
            "value": 1291143946,
            "unit": "ns/op\t215541288 B/op\t 1042866 allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader - ns/op",
            "value": 1291143946,
            "unit": "ns/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader - B/op",
            "value": 215541288,
            "unit": "B/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader - allocs/op",
            "value": 1042866,
            "unit": "allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator",
            "value": 397508589,
            "unit": "ns/op\t215424901 B/op\t 1042801 allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator - ns/op",
            "value": 397508589,
            "unit": "ns/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator - B/op",
            "value": 215424901,
            "unit": "B/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator - allocs/op",
            "value": 1042801,
            "unit": "allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader",
            "value": 2224472894,
            "unit": "ns/op\t313545480 B/op\t 4568039 allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader - ns/op",
            "value": 2224472894,
            "unit": "ns/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader - B/op",
            "value": 313545480,
            "unit": "B/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader - allocs/op",
            "value": 4568039,
            "unit": "allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator",
            "value": 752624394,
            "unit": "ns/op\t305430260 B/op\t 3568073 allocs/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator - ns/op",
            "value": 752624394,
            "unit": "ns/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator - B/op",
            "value": 305430260,
            "unit": "B/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator - allocs/op",
            "value": 3568073,
            "unit": "allocs/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead",
            "value": 41339,
            "unit": "ns/op\t   23312 B/op\t     114 allocs/op",
            "extra": "28980 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead - ns/op",
            "value": 41339,
            "unit": "ns/op",
            "extra": "28980 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead - B/op",
            "value": 23312,
            "unit": "B/op",
            "extra": "28980 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead - allocs/op",
            "value": 114,
            "unit": "allocs/op",
            "extra": "28980 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead",
            "value": 71509,
            "unit": "ns/op\t   26920 B/op\t     159 allocs/op",
            "extra": "15788 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead - ns/op",
            "value": 71509,
            "unit": "ns/op",
            "extra": "15788 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead - B/op",
            "value": 26920,
            "unit": "B/op",
            "extra": "15788 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead - allocs/op",
            "value": 159,
            "unit": "allocs/op",
            "extra": "15788 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead",
            "value": 30697,
            "unit": "ns/op\t   21848 B/op\t      77 allocs/op",
            "extra": "39763 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead - ns/op",
            "value": 30697,
            "unit": "ns/op",
            "extra": "39763 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead - B/op",
            "value": 21848,
            "unit": "B/op",
            "extra": "39763 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead - allocs/op",
            "value": 77,
            "unit": "allocs/op",
            "extra": "39763 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead",
            "value": 270390,
            "unit": "ns/op\t   56353 B/op\t     552 allocs/op",
            "extra": "4453 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead - ns/op",
            "value": 270390,
            "unit": "ns/op",
            "extra": "4453 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead - B/op",
            "value": 56353,
            "unit": "B/op",
            "extra": "4453 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead - allocs/op",
            "value": 552,
            "unit": "allocs/op",
            "extra": "4453 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2",
            "value": 270064,
            "unit": "ns/op\t   56353 B/op\t     552 allocs/op",
            "extra": "4388 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2 - ns/op",
            "value": 270064,
            "unit": "ns/op",
            "extra": "4388 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2 - B/op",
            "value": 56353,
            "unit": "B/op",
            "extra": "4388 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2 - allocs/op",
            "value": 552,
            "unit": "allocs/op",
            "extra": "4388 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3",
            "value": 102975,
            "unit": "ns/op\t   29752 B/op\t     264 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3 - ns/op",
            "value": 102975,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3 - B/op",
            "value": 29752,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3 - allocs/op",
            "value": 264,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile",
            "value": 14022,
            "unit": "ns/op\t   10760 B/op\t     131 allocs/op",
            "extra": "93858 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile - ns/op",
            "value": 14022,
            "unit": "ns/op",
            "extra": "93858 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile - B/op",
            "value": 10760,
            "unit": "B/op",
            "extra": "93858 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile - allocs/op",
            "value": 131,
            "unit": "allocs/op",
            "extra": "93858 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid",
            "value": 25.1,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "46996730 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid - ns/op",
            "value": 25.1,
            "unit": "ns/op",
            "extra": "46996730 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "46996730 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "46996730 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite",
            "value": 49181,
            "unit": "ns/op\t   34696 B/op\t     226 allocs/op",
            "extra": "23186 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite - ns/op",
            "value": 49181,
            "unit": "ns/op",
            "extra": "23186 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite - B/op",
            "value": 34696,
            "unit": "B/op",
            "extra": "23186 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite - allocs/op",
            "value": 226,
            "unit": "allocs/op",
            "extra": "23186 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite",
            "value": 206111,
            "unit": "ns/op\t   54680 B/op\t    2069 allocs/op",
            "extra": "5904 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite - ns/op",
            "value": 206111,
            "unit": "ns/op",
            "extra": "5904 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite - B/op",
            "value": 54680,
            "unit": "B/op",
            "extra": "5904 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite - allocs/op",
            "value": 2069,
            "unit": "allocs/op",
            "extra": "5904 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite",
            "value": 151200,
            "unit": "ns/op\t   61104 B/op\t     721 allocs/op",
            "extra": "7821 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite - ns/op",
            "value": 151200,
            "unit": "ns/op",
            "extra": "7821 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite - B/op",
            "value": 61104,
            "unit": "B/op",
            "extra": "7821 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite - allocs/op",
            "value": 721,
            "unit": "allocs/op",
            "extra": "7821 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDIATWrite",
            "value": 231380104,
            "unit": "1210428822609200234A094101Federal Reserve Bank   My Bank Name                   ",
            "extra": "101 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDIATWrite - 1210428822609200234A094101Federal",
            "value": 231380104,
            "unit": "1210428822609200234A094101Federal",
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
          }
        ]
      },
      {
        "commit": {
          "author": {
            "name": "renovate[bot]",
            "username": "renovate[bot]",
            "email": "29139614+renovate[bot]@users.noreply.github.com"
          },
          "committer": {
            "name": "GitHub",
            "username": "web-flow",
            "email": "noreply@github.com"
          },
          "id": "f36ebb7ef2e6d678ef95ca6278d2b1f2deae4299",
          "message": "chore(deps): update github/codeql-action action to v4.38.1 (#1863)\n\nCo-authored-by: renovate[bot] <29139614+renovate[bot]@users.noreply.github.com>",
          "timestamp": "2026-09-19T02:30:43Z",
          "url": "https://github.com/moov-io/ach/commit/f36ebb7ef2e6d678ef95ca6278d2b1f2deae4299"
        },
        "date": 1789872357320,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkAlphaFieldShort",
            "value": 979.7,
            "unit": "ns/op\t      80 B/op\t       3 allocs/op",
            "extra": "1219075 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldShort - ns/op",
            "value": 979.7,
            "unit": "ns/op",
            "extra": "1219075 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldShort - B/op",
            "value": 80,
            "unit": "B/op",
            "extra": "1219075 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldShort - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "1219075 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong",
            "value": 97.87,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "12185284 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong - ns/op",
            "value": 97.87,
            "unit": "ns/op",
            "extra": "12185284 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "12185284 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "12185284 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort",
            "value": 59.65,
            "unit": "ns/op\t      16 B/op\t       2 allocs/op",
            "extra": "19537617 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort - ns/op",
            "value": 59.65,
            "unit": "ns/op",
            "extra": "19537617 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "19537617 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "19537617 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong",
            "value": 25.42,
            "unit": "ns/op\t       8 B/op\t       1 allocs/op",
            "extra": "46547578 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong - ns/op",
            "value": 25.42,
            "unit": "ns/op",
            "extra": "46547578 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong - B/op",
            "value": 8,
            "unit": "B/op",
            "extra": "46547578 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "46547578 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField",
            "value": 15.3,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "81630358 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField - ns/op",
            "value": 15.3,
            "unit": "ns/op",
            "extra": "81630358 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "81630358 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "81630358 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField",
            "value": 5.627,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "213365851 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField - ns/op",
            "value": 5.627,
            "unit": "ns/op",
            "extra": "213365851 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "213365851 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "213365851 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles",
            "value": 245606,
            "unit": "ns/op\t   54146 B/op\t     306 allocs/op",
            "extra": "4162 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles - ns/op",
            "value": 245606,
            "unit": "ns/op",
            "extra": "4162 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles - B/op",
            "value": 54146,
            "unit": "B/op",
            "extra": "4162 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles - allocs/op",
            "value": 306,
            "unit": "allocs/op",
            "extra": "4162 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts",
            "value": 244256,
            "unit": "ns/op\t   54159 B/op\t     306 allocs/op",
            "extra": "5116 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts - ns/op",
            "value": 244256,
            "unit": "ns/op",
            "extra": "5116 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts - B/op",
            "value": 54159,
            "unit": "B/op",
            "extra": "5116 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts - allocs/op",
            "value": 306,
            "unit": "allocs/op",
            "extra": "5116 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir",
            "value": 147457,
            "unit": "ns/op\t   54550 B/op\t     310 allocs/op",
            "extra": "7760 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir - ns/op",
            "value": 147457,
            "unit": "ns/op",
            "extra": "7760 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir - B/op",
            "value": 54550,
            "unit": "B/op",
            "extra": "7760 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir - allocs/op",
            "value": 310,
            "unit": "allocs/op",
            "extra": "7760 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts",
            "value": 138995,
            "unit": "ns/op\t   54588 B/op\t     310 allocs/op",
            "extra": "7389 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts - ns/op",
            "value": 138995,
            "unit": "ns/op",
            "extra": "7389 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts - B/op",
            "value": 54588,
            "unit": "B/op",
            "extra": "7389 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts - allocs/op",
            "value": 310,
            "unit": "allocs/op",
            "extra": "7389 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups",
            "value": 316262,
            "unit": "ns/op\t   59870 B/op\t     366 allocs/op",
            "extra": "3818 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups - ns/op",
            "value": 316262,
            "unit": "ns/op",
            "extra": "3818 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups - B/op",
            "value": 59870,
            "unit": "B/op",
            "extra": "3818 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups - allocs/op",
            "value": 366,
            "unit": "allocs/op",
            "extra": "3818 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups",
            "value": 302584,
            "unit": "ns/op\t   59893 B/op\t     366 allocs/op",
            "extra": "4130 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups - ns/op",
            "value": 302584,
            "unit": "ns/op",
            "extra": "4130 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups - B/op",
            "value": 59893,
            "unit": "B/op",
            "extra": "4130 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups - allocs/op",
            "value": 366,
            "unit": "allocs/op",
            "extra": "4130 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups",
            "value": 297645,
            "unit": "ns/op\t   59889 B/op\t     366 allocs/op",
            "extra": "4164 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups - ns/op",
            "value": 297645,
            "unit": "ns/op",
            "extra": "4164 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups - B/op",
            "value": 59889,
            "unit": "B/op",
            "extra": "4164 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups - allocs/op",
            "value": 366,
            "unit": "allocs/op",
            "extra": "4164 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups",
            "value": 285253,
            "unit": "ns/op\t   59504 B/op\t     367 allocs/op",
            "extra": "5000 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups - ns/op",
            "value": 285253,
            "unit": "ns/op",
            "extra": "5000 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups - B/op",
            "value": 59504,
            "unit": "B/op",
            "extra": "5000 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups - allocs/op",
            "value": 367,
            "unit": "allocs/op",
            "extra": "5000 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader",
            "value": 241085558,
            "unit": "ns/op\t42449980 B/op\t  203653 allocs/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader - ns/op",
            "value": 241085558,
            "unit": "ns/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader - B/op",
            "value": 42449980,
            "unit": "B/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader - allocs/op",
            "value": 203653,
            "unit": "allocs/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator",
            "value": 77586106,
            "unit": "ns/op\t42424718 B/op\t  203601 allocs/op",
            "extra": "14 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator - ns/op",
            "value": 77586106,
            "unit": "ns/op",
            "extra": "14 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator - B/op",
            "value": 42424718,
            "unit": "B/op",
            "extra": "14 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator - allocs/op",
            "value": 203601,
            "unit": "allocs/op",
            "extra": "14 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader",
            "value": 454905450,
            "unit": "ns/op\t62131557 B/op\t  905792 allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader - ns/op",
            "value": 454905450,
            "unit": "ns/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader - B/op",
            "value": 62131557,
            "unit": "B/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader - allocs/op",
            "value": 905792,
            "unit": "allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator",
            "value": 150068331,
            "unit": "ns/op\t60508768 B/op\t  705839 allocs/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator - ns/op",
            "value": 150068331,
            "unit": "ns/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator - B/op",
            "value": 60508768,
            "unit": "B/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator - allocs/op",
            "value": 705839,
            "unit": "allocs/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader",
            "value": 1286959031,
            "unit": "ns/op\t215541576 B/op\t 1042870 allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader - ns/op",
            "value": 1286959031,
            "unit": "ns/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader - B/op",
            "value": 215541576,
            "unit": "B/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader - allocs/op",
            "value": 1042870,
            "unit": "allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator",
            "value": 395167890,
            "unit": "ns/op\t215424730 B/op\t 1042799 allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator - ns/op",
            "value": 395167890,
            "unit": "ns/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator - B/op",
            "value": 215424730,
            "unit": "B/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator - allocs/op",
            "value": 1042799,
            "unit": "allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader",
            "value": 2208151066,
            "unit": "ns/op\t313545416 B/op\t 4568038 allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader - ns/op",
            "value": 2208151066,
            "unit": "ns/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader - B/op",
            "value": 313545416,
            "unit": "B/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader - allocs/op",
            "value": 4568038,
            "unit": "allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator",
            "value": 751602136,
            "unit": "ns/op\t305430116 B/op\t 3568071 allocs/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator - ns/op",
            "value": 751602136,
            "unit": "ns/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator - B/op",
            "value": 305430116,
            "unit": "B/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator - allocs/op",
            "value": 3568071,
            "unit": "allocs/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead",
            "value": 42265,
            "unit": "ns/op\t   23312 B/op\t     114 allocs/op",
            "extra": "28890 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead - ns/op",
            "value": 42265,
            "unit": "ns/op",
            "extra": "28890 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead - B/op",
            "value": 23312,
            "unit": "B/op",
            "extra": "28890 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead - allocs/op",
            "value": 114,
            "unit": "allocs/op",
            "extra": "28890 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead",
            "value": 71551,
            "unit": "ns/op\t   26920 B/op\t     159 allocs/op",
            "extra": "16134 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead - ns/op",
            "value": 71551,
            "unit": "ns/op",
            "extra": "16134 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead - B/op",
            "value": 26920,
            "unit": "B/op",
            "extra": "16134 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead - allocs/op",
            "value": 159,
            "unit": "allocs/op",
            "extra": "16134 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead",
            "value": 31301,
            "unit": "ns/op\t   21848 B/op\t      77 allocs/op",
            "extra": "39559 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead - ns/op",
            "value": 31301,
            "unit": "ns/op",
            "extra": "39559 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead - B/op",
            "value": 21848,
            "unit": "B/op",
            "extra": "39559 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead - allocs/op",
            "value": 77,
            "unit": "allocs/op",
            "extra": "39559 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead",
            "value": 273527,
            "unit": "ns/op\t   56353 B/op\t     552 allocs/op",
            "extra": "4520 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead - ns/op",
            "value": 273527,
            "unit": "ns/op",
            "extra": "4520 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead - B/op",
            "value": 56353,
            "unit": "B/op",
            "extra": "4520 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead - allocs/op",
            "value": 552,
            "unit": "allocs/op",
            "extra": "4520 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2",
            "value": 273916,
            "unit": "ns/op\t   56353 B/op\t     552 allocs/op",
            "extra": "4452 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2 - ns/op",
            "value": 273916,
            "unit": "ns/op",
            "extra": "4452 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2 - B/op",
            "value": 56353,
            "unit": "B/op",
            "extra": "4452 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2 - allocs/op",
            "value": 552,
            "unit": "allocs/op",
            "extra": "4452 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3",
            "value": 102821,
            "unit": "ns/op\t   29752 B/op\t     264 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3 - ns/op",
            "value": 102821,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3 - B/op",
            "value": 29752,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3 - allocs/op",
            "value": 264,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile",
            "value": 13993,
            "unit": "ns/op\t   10760 B/op\t     131 allocs/op",
            "extra": "92996 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile - ns/op",
            "value": 13993,
            "unit": "ns/op",
            "extra": "92996 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile - B/op",
            "value": 10760,
            "unit": "B/op",
            "extra": "92996 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile - allocs/op",
            "value": 131,
            "unit": "allocs/op",
            "extra": "92996 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid",
            "value": 25.07,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "47747230 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid - ns/op",
            "value": 25.07,
            "unit": "ns/op",
            "extra": "47747230 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "47747230 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "47747230 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite",
            "value": 49607,
            "unit": "ns/op\t   34696 B/op\t     226 allocs/op",
            "extra": "23336 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite - ns/op",
            "value": 49607,
            "unit": "ns/op",
            "extra": "23336 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite - B/op",
            "value": 34696,
            "unit": "B/op",
            "extra": "23336 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite - allocs/op",
            "value": 226,
            "unit": "allocs/op",
            "extra": "23336 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite",
            "value": 206706,
            "unit": "ns/op\t   54680 B/op\t    2069 allocs/op",
            "extra": "5607 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite - ns/op",
            "value": 206706,
            "unit": "ns/op",
            "extra": "5607 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite - B/op",
            "value": 54680,
            "unit": "B/op",
            "extra": "5607 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite - allocs/op",
            "value": 2069,
            "unit": "allocs/op",
            "extra": "5607 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite",
            "value": 153362,
            "unit": "ns/op\t   61104 B/op\t     721 allocs/op",
            "extra": "7526 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite - ns/op",
            "value": 153362,
            "unit": "ns/op",
            "extra": "7526 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite - B/op",
            "value": 61104,
            "unit": "B/op",
            "extra": "7526 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite - allocs/op",
            "value": 721,
            "unit": "allocs/op",
            "extra": "7526 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDIATWrite",
            "value": 231380104,
            "unit": "1210428822609210245A094101Federal Reserve Bank   My Bank Name                   ",
            "extra": "101 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDIATWrite - 1210428822609210245A094101Federal",
            "value": 231380104,
            "unit": "1210428822609210245A094101Federal",
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
          }
        ]
      },
      {
        "commit": {
          "author": {
            "name": "renovate[bot]",
            "username": "renovate[bot]",
            "email": "29139614+renovate[bot]@users.noreply.github.com"
          },
          "committer": {
            "name": "GitHub",
            "username": "web-flow",
            "email": "noreply@github.com"
          },
          "id": "f36ebb7ef2e6d678ef95ca6278d2b1f2deae4299",
          "message": "chore(deps): update github/codeql-action action to v4.38.1 (#1863)\n\nCo-authored-by: renovate[bot] <29139614+renovate[bot]@users.noreply.github.com>",
          "timestamp": "2026-09-19T02:30:43Z",
          "url": "https://github.com/moov-io/ach/commit/f36ebb7ef2e6d678ef95ca6278d2b1f2deae4299"
        },
        "date": 1789958591176,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkAlphaFieldShort",
            "value": 989.7,
            "unit": "ns/op\t      80 B/op\t       3 allocs/op",
            "extra": "1204866 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldShort - ns/op",
            "value": 989.7,
            "unit": "ns/op",
            "extra": "1204866 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldShort - B/op",
            "value": 80,
            "unit": "B/op",
            "extra": "1204866 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldShort - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "1204866 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong",
            "value": 97.85,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "12155425 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong - ns/op",
            "value": 97.85,
            "unit": "ns/op",
            "extra": "12155425 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "12155425 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "12155425 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort",
            "value": 59.44,
            "unit": "ns/op\t      16 B/op\t       2 allocs/op",
            "extra": "19468945 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort - ns/op",
            "value": 59.44,
            "unit": "ns/op",
            "extra": "19468945 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "19468945 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "19468945 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong",
            "value": 25.27,
            "unit": "ns/op\t       8 B/op\t       1 allocs/op",
            "extra": "45837855 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong - ns/op",
            "value": 25.27,
            "unit": "ns/op",
            "extra": "45837855 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong - B/op",
            "value": 8,
            "unit": "B/op",
            "extra": "45837855 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "45837855 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField",
            "value": 14.98,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "81739598 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField - ns/op",
            "value": 14.98,
            "unit": "ns/op",
            "extra": "81739598 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "81739598 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "81739598 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField",
            "value": 5.615,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "213939886 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField - ns/op",
            "value": 5.615,
            "unit": "ns/op",
            "extra": "213939886 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "213939886 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "213939886 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles",
            "value": 261813,
            "unit": "ns/op\t   54132 B/op\t     306 allocs/op",
            "extra": "4983 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles - ns/op",
            "value": 261813,
            "unit": "ns/op",
            "extra": "4983 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles - B/op",
            "value": 54132,
            "unit": "B/op",
            "extra": "4983 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles - allocs/op",
            "value": 306,
            "unit": "allocs/op",
            "extra": "4983 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts",
            "value": 242779,
            "unit": "ns/op\t   54153 B/op\t     306 allocs/op",
            "extra": "4600 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts - ns/op",
            "value": 242779,
            "unit": "ns/op",
            "extra": "4600 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts - B/op",
            "value": 54153,
            "unit": "B/op",
            "extra": "4600 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts - allocs/op",
            "value": 306,
            "unit": "allocs/op",
            "extra": "4600 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir",
            "value": 146018,
            "unit": "ns/op\t   54541 B/op\t     310 allocs/op",
            "extra": "8253 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir - ns/op",
            "value": 146018,
            "unit": "ns/op",
            "extra": "8253 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir - B/op",
            "value": 54541,
            "unit": "B/op",
            "extra": "8253 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir - allocs/op",
            "value": 310,
            "unit": "allocs/op",
            "extra": "8253 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts",
            "value": 139491,
            "unit": "ns/op\t   54575 B/op\t     310 allocs/op",
            "extra": "7789 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts - ns/op",
            "value": 139491,
            "unit": "ns/op",
            "extra": "7789 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts - B/op",
            "value": 54575,
            "unit": "B/op",
            "extra": "7789 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts - allocs/op",
            "value": 310,
            "unit": "allocs/op",
            "extra": "7789 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups",
            "value": 308319,
            "unit": "ns/op\t   59903 B/op\t     366 allocs/op",
            "extra": "3943 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups - ns/op",
            "value": 308319,
            "unit": "ns/op",
            "extra": "3943 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups - B/op",
            "value": 59903,
            "unit": "B/op",
            "extra": "3943 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups - allocs/op",
            "value": 366,
            "unit": "allocs/op",
            "extra": "3943 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups",
            "value": 305690,
            "unit": "ns/op\t   59911 B/op\t     366 allocs/op",
            "extra": "4093 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups - ns/op",
            "value": 305690,
            "unit": "ns/op",
            "extra": "4093 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups - B/op",
            "value": 59911,
            "unit": "B/op",
            "extra": "4093 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups - allocs/op",
            "value": 366,
            "unit": "allocs/op",
            "extra": "4093 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups",
            "value": 296233,
            "unit": "ns/op\t   59909 B/op\t     366 allocs/op",
            "extra": "4104 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups - ns/op",
            "value": 296233,
            "unit": "ns/op",
            "extra": "4104 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups - B/op",
            "value": 59909,
            "unit": "B/op",
            "extra": "4104 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups - allocs/op",
            "value": 366,
            "unit": "allocs/op",
            "extra": "4104 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups",
            "value": 286429,
            "unit": "ns/op\t   59601 B/op\t     367 allocs/op",
            "extra": "4796 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups - ns/op",
            "value": 286429,
            "unit": "ns/op",
            "extra": "4796 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups - B/op",
            "value": 59601,
            "unit": "B/op",
            "extra": "4796 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups - allocs/op",
            "value": 367,
            "unit": "allocs/op",
            "extra": "4796 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader",
            "value": 244596961,
            "unit": "ns/op\t42449897 B/op\t  203652 allocs/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader - ns/op",
            "value": 244596961,
            "unit": "ns/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader - B/op",
            "value": 42449897,
            "unit": "B/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader - allocs/op",
            "value": 203652,
            "unit": "allocs/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator",
            "value": 78984470,
            "unit": "ns/op\t42424852 B/op\t  203602 allocs/op",
            "extra": "14 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator - ns/op",
            "value": 78984470,
            "unit": "ns/op",
            "extra": "14 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator - B/op",
            "value": 42424852,
            "unit": "B/op",
            "extra": "14 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator - allocs/op",
            "value": 203602,
            "unit": "allocs/op",
            "extra": "14 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader",
            "value": 446930011,
            "unit": "ns/op\t62131594 B/op\t  905792 allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader - ns/op",
            "value": 446930011,
            "unit": "ns/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader - B/op",
            "value": 62131594,
            "unit": "B/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader - allocs/op",
            "value": 905792,
            "unit": "allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator",
            "value": 149980272,
            "unit": "ns/op\t60508704 B/op\t  705838 allocs/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator - ns/op",
            "value": 149980272,
            "unit": "ns/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator - B/op",
            "value": 60508704,
            "unit": "B/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator - allocs/op",
            "value": 705838,
            "unit": "allocs/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader",
            "value": 1286118887,
            "unit": "ns/op\t215541208 B/op\t 1042865 allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader - ns/op",
            "value": 1286118887,
            "unit": "ns/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader - B/op",
            "value": 215541208,
            "unit": "B/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader - allocs/op",
            "value": 1042865,
            "unit": "allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator",
            "value": 396713221,
            "unit": "ns/op\t215424672 B/op\t 1042798 allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator - ns/op",
            "value": 396713221,
            "unit": "ns/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator - B/op",
            "value": 215424672,
            "unit": "B/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator - allocs/op",
            "value": 1042798,
            "unit": "allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader",
            "value": 2213504098,
            "unit": "ns/op\t313545400 B/op\t 4568038 allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader - ns/op",
            "value": 2213504098,
            "unit": "ns/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader - B/op",
            "value": 313545400,
            "unit": "B/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader - allocs/op",
            "value": 4568038,
            "unit": "allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator",
            "value": 752607028,
            "unit": "ns/op\t305430196 B/op\t 3568072 allocs/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator - ns/op",
            "value": 752607028,
            "unit": "ns/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator - B/op",
            "value": 305430196,
            "unit": "B/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator - allocs/op",
            "value": 3568072,
            "unit": "allocs/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead",
            "value": 41033,
            "unit": "ns/op\t   23312 B/op\t     114 allocs/op",
            "extra": "29236 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead - ns/op",
            "value": 41033,
            "unit": "ns/op",
            "extra": "29236 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead - B/op",
            "value": 23312,
            "unit": "B/op",
            "extra": "29236 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead - allocs/op",
            "value": 114,
            "unit": "allocs/op",
            "extra": "29236 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead",
            "value": 71164,
            "unit": "ns/op\t   26920 B/op\t     159 allocs/op",
            "extra": "16328 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead - ns/op",
            "value": 71164,
            "unit": "ns/op",
            "extra": "16328 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead - B/op",
            "value": 26920,
            "unit": "B/op",
            "extra": "16328 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead - allocs/op",
            "value": 159,
            "unit": "allocs/op",
            "extra": "16328 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead",
            "value": 30957,
            "unit": "ns/op\t   21848 B/op\t      77 allocs/op",
            "extra": "39312 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead - ns/op",
            "value": 30957,
            "unit": "ns/op",
            "extra": "39312 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead - B/op",
            "value": 21848,
            "unit": "B/op",
            "extra": "39312 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead - allocs/op",
            "value": 77,
            "unit": "allocs/op",
            "extra": "39312 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead",
            "value": 268854,
            "unit": "ns/op\t   56353 B/op\t     552 allocs/op",
            "extra": "4618 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead - ns/op",
            "value": 268854,
            "unit": "ns/op",
            "extra": "4618 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead - B/op",
            "value": 56353,
            "unit": "B/op",
            "extra": "4618 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead - allocs/op",
            "value": 552,
            "unit": "allocs/op",
            "extra": "4618 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2",
            "value": 271136,
            "unit": "ns/op\t   56353 B/op\t     552 allocs/op",
            "extra": "4557 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2 - ns/op",
            "value": 271136,
            "unit": "ns/op",
            "extra": "4557 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2 - B/op",
            "value": 56353,
            "unit": "B/op",
            "extra": "4557 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2 - allocs/op",
            "value": 552,
            "unit": "allocs/op",
            "extra": "4557 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3",
            "value": 102326,
            "unit": "ns/op\t   29752 B/op\t     264 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3 - ns/op",
            "value": 102326,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3 - B/op",
            "value": 29752,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3 - allocs/op",
            "value": 264,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile",
            "value": 13913,
            "unit": "ns/op\t   10760 B/op\t     131 allocs/op",
            "extra": "92692 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile - ns/op",
            "value": 13913,
            "unit": "ns/op",
            "extra": "92692 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile - B/op",
            "value": 10760,
            "unit": "B/op",
            "extra": "92692 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile - allocs/op",
            "value": 131,
            "unit": "allocs/op",
            "extra": "92692 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid",
            "value": 25.09,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "46411038 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid - ns/op",
            "value": 25.09,
            "unit": "ns/op",
            "extra": "46411038 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "46411038 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "46411038 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite",
            "value": 48675,
            "unit": "ns/op\t   34696 B/op\t     226 allocs/op",
            "extra": "23120 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite - ns/op",
            "value": 48675,
            "unit": "ns/op",
            "extra": "23120 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite - B/op",
            "value": 34696,
            "unit": "B/op",
            "extra": "23120 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite - allocs/op",
            "value": 226,
            "unit": "allocs/op",
            "extra": "23120 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite",
            "value": 207655,
            "unit": "ns/op\t   54680 B/op\t    2069 allocs/op",
            "extra": "6031 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite - ns/op",
            "value": 207655,
            "unit": "ns/op",
            "extra": "6031 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite - B/op",
            "value": 54680,
            "unit": "B/op",
            "extra": "6031 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite - allocs/op",
            "value": 2069,
            "unit": "allocs/op",
            "extra": "6031 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite",
            "value": 153799,
            "unit": "ns/op\t   61104 B/op\t     721 allocs/op",
            "extra": "7455 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite - ns/op",
            "value": 153799,
            "unit": "ns/op",
            "extra": "7455 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite - B/op",
            "value": 61104,
            "unit": "B/op",
            "extra": "7455 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite - allocs/op",
            "value": 721,
            "unit": "allocs/op",
            "extra": "7455 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDIATWrite",
            "value": 231380104,
            "unit": "1210428822609220243A094101Federal Reserve Bank   My Bank Name                   ",
            "extra": "101 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDIATWrite - 1210428822609220243A094101Federal",
            "value": 231380104,
            "unit": "1210428822609220243A094101Federal",
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
          }
        ]
      },
      {
        "commit": {
          "author": {
            "name": "renovate[bot]",
            "username": "renovate[bot]",
            "email": "29139614+renovate[bot]@users.noreply.github.com"
          },
          "committer": {
            "name": "GitHub",
            "username": "web-flow",
            "email": "noreply@github.com"
          },
          "id": "5f5046060c4be6fd99c19068a1d8e61809f1b172",
          "message": "fix(deps): update module github.com/aws/aws-lambda-go to v1.55.1 (#1867)\n\nCo-authored-by: renovate[bot] <29139614+renovate[bot]@users.noreply.github.com>",
          "timestamp": "2026-09-21T21:23:57Z",
          "url": "https://github.com/moov-io/ach/commit/5f5046060c4be6fd99c19068a1d8e61809f1b172"
        },
        "date": 1790045089528,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkAlphaFieldShort",
            "value": 968.3,
            "unit": "ns/op\t      80 B/op\t       3 allocs/op",
            "extra": "1252368 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldShort - ns/op",
            "value": 968.3,
            "unit": "ns/op",
            "extra": "1252368 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldShort - B/op",
            "value": 80,
            "unit": "B/op",
            "extra": "1252368 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldShort - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "1252368 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong",
            "value": 98.14,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "12061948 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong - ns/op",
            "value": 98.14,
            "unit": "ns/op",
            "extra": "12061948 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "12061948 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "12061948 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort",
            "value": 57.1,
            "unit": "ns/op\t      16 B/op\t       2 allocs/op",
            "extra": "20760729 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort - ns/op",
            "value": 57.1,
            "unit": "ns/op",
            "extra": "20760729 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "20760729 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "20760729 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong",
            "value": 25.41,
            "unit": "ns/op\t       8 B/op\t       1 allocs/op",
            "extra": "43315162 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong - ns/op",
            "value": 25.41,
            "unit": "ns/op",
            "extra": "43315162 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong - B/op",
            "value": 8,
            "unit": "B/op",
            "extra": "43315162 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "43315162 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField",
            "value": 14.93,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "83134358 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField - ns/op",
            "value": 14.93,
            "unit": "ns/op",
            "extra": "83134358 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "83134358 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "83134358 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField",
            "value": 5.617,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "213641396 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField - ns/op",
            "value": 5.617,
            "unit": "ns/op",
            "extra": "213641396 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "213641396 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "213641396 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles",
            "value": 260186,
            "unit": "ns/op\t   54144 B/op\t     306 allocs/op",
            "extra": "4960 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles - ns/op",
            "value": 260186,
            "unit": "ns/op",
            "extra": "4960 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles - B/op",
            "value": 54144,
            "unit": "B/op",
            "extra": "4960 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles - allocs/op",
            "value": 306,
            "unit": "allocs/op",
            "extra": "4960 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts",
            "value": 260192,
            "unit": "ns/op\t   54157 B/op\t     306 allocs/op",
            "extra": "4479 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts - ns/op",
            "value": 260192,
            "unit": "ns/op",
            "extra": "4479 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts - B/op",
            "value": 54157,
            "unit": "B/op",
            "extra": "4479 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts - allocs/op",
            "value": 306,
            "unit": "allocs/op",
            "extra": "4479 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir",
            "value": 148446,
            "unit": "ns/op\t   54550 B/op\t     310 allocs/op",
            "extra": "7632 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir - ns/op",
            "value": 148446,
            "unit": "ns/op",
            "extra": "7632 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir - B/op",
            "value": 54550,
            "unit": "B/op",
            "extra": "7632 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir - allocs/op",
            "value": 310,
            "unit": "allocs/op",
            "extra": "7632 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts",
            "value": 149770,
            "unit": "ns/op\t   54586 B/op\t     310 allocs/op",
            "extra": "6728 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts - ns/op",
            "value": 149770,
            "unit": "ns/op",
            "extra": "6728 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts - B/op",
            "value": 54586,
            "unit": "B/op",
            "extra": "6728 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts - allocs/op",
            "value": 310,
            "unit": "allocs/op",
            "extra": "6728 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups",
            "value": 312104,
            "unit": "ns/op\t   59875 B/op\t     366 allocs/op",
            "extra": "3782 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups - ns/op",
            "value": 312104,
            "unit": "ns/op",
            "extra": "3782 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups - B/op",
            "value": 59875,
            "unit": "B/op",
            "extra": "3782 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups - allocs/op",
            "value": 366,
            "unit": "allocs/op",
            "extra": "3782 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups",
            "value": 306189,
            "unit": "ns/op\t   59854 B/op\t     366 allocs/op",
            "extra": "4027 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups - ns/op",
            "value": 306189,
            "unit": "ns/op",
            "extra": "4027 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups - B/op",
            "value": 59854,
            "unit": "B/op",
            "extra": "4027 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups - allocs/op",
            "value": 366,
            "unit": "allocs/op",
            "extra": "4027 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups",
            "value": 299662,
            "unit": "ns/op\t   59837 B/op\t     366 allocs/op",
            "extra": "4242 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups - ns/op",
            "value": 299662,
            "unit": "ns/op",
            "extra": "4242 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups - B/op",
            "value": 59837,
            "unit": "B/op",
            "extra": "4242 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups - allocs/op",
            "value": 366,
            "unit": "allocs/op",
            "extra": "4242 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups",
            "value": 301105,
            "unit": "ns/op\t   59566 B/op\t     367 allocs/op",
            "extra": "4899 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups - ns/op",
            "value": 301105,
            "unit": "ns/op",
            "extra": "4899 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups - B/op",
            "value": 59566,
            "unit": "B/op",
            "extra": "4899 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups - allocs/op",
            "value": 367,
            "unit": "allocs/op",
            "extra": "4899 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader",
            "value": 247539203,
            "unit": "ns/op\t42450299 B/op\t  203656 allocs/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader - ns/op",
            "value": 247539203,
            "unit": "ns/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader - B/op",
            "value": 42450299,
            "unit": "B/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader - allocs/op",
            "value": 203656,
            "unit": "allocs/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator",
            "value": 83859695,
            "unit": "ns/op\t42424847 B/op\t  203602 allocs/op",
            "extra": "13 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator - ns/op",
            "value": 83859695,
            "unit": "ns/op",
            "extra": "13 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator - B/op",
            "value": 42424847,
            "unit": "B/op",
            "extra": "13 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator - allocs/op",
            "value": 203602,
            "unit": "allocs/op",
            "extra": "13 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader",
            "value": 457314721,
            "unit": "ns/op\t62131616 B/op\t  905793 allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader - ns/op",
            "value": 457314721,
            "unit": "ns/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader - B/op",
            "value": 62131616,
            "unit": "B/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader - allocs/op",
            "value": 905793,
            "unit": "allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator",
            "value": 155931133,
            "unit": "ns/op\t60508708 B/op\t  705838 allocs/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator - ns/op",
            "value": 155931133,
            "unit": "ns/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator - B/op",
            "value": 60508708,
            "unit": "B/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator - allocs/op",
            "value": 705838,
            "unit": "allocs/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader",
            "value": 1313171894,
            "unit": "ns/op\t215541432 B/op\t 1042868 allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader - ns/op",
            "value": 1313171894,
            "unit": "ns/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader - B/op",
            "value": 215541432,
            "unit": "B/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader - allocs/op",
            "value": 1042868,
            "unit": "allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator",
            "value": 420089780,
            "unit": "ns/op\t215424853 B/op\t 1042801 allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator - ns/op",
            "value": 420089780,
            "unit": "ns/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator - B/op",
            "value": 215424853,
            "unit": "B/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator - allocs/op",
            "value": 1042801,
            "unit": "allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader",
            "value": 2216982701,
            "unit": "ns/op\t313545336 B/op\t 4568037 allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader - ns/op",
            "value": 2216982701,
            "unit": "ns/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader - B/op",
            "value": 313545336,
            "unit": "B/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader - allocs/op",
            "value": 4568037,
            "unit": "allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator",
            "value": 774234034,
            "unit": "ns/op\t305430308 B/op\t 3568074 allocs/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator - ns/op",
            "value": 774234034,
            "unit": "ns/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator - B/op",
            "value": 305430308,
            "unit": "B/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator - allocs/op",
            "value": 3568074,
            "unit": "allocs/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead",
            "value": 41916,
            "unit": "ns/op\t   23312 B/op\t     114 allocs/op",
            "extra": "29011 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead - ns/op",
            "value": 41916,
            "unit": "ns/op",
            "extra": "29011 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead - B/op",
            "value": 23312,
            "unit": "B/op",
            "extra": "29011 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead - allocs/op",
            "value": 114,
            "unit": "allocs/op",
            "extra": "29011 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead",
            "value": 71961,
            "unit": "ns/op\t   26920 B/op\t     159 allocs/op",
            "extra": "16137 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead - ns/op",
            "value": 71961,
            "unit": "ns/op",
            "extra": "16137 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead - B/op",
            "value": 26920,
            "unit": "B/op",
            "extra": "16137 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead - allocs/op",
            "value": 159,
            "unit": "allocs/op",
            "extra": "16137 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead",
            "value": 31817,
            "unit": "ns/op\t   21848 B/op\t      77 allocs/op",
            "extra": "39369 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead - ns/op",
            "value": 31817,
            "unit": "ns/op",
            "extra": "39369 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead - B/op",
            "value": 21848,
            "unit": "B/op",
            "extra": "39369 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead - allocs/op",
            "value": 77,
            "unit": "allocs/op",
            "extra": "39369 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead",
            "value": 269332,
            "unit": "ns/op\t   56353 B/op\t     552 allocs/op",
            "extra": "4448 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead - ns/op",
            "value": 269332,
            "unit": "ns/op",
            "extra": "4448 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead - B/op",
            "value": 56353,
            "unit": "B/op",
            "extra": "4448 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead - allocs/op",
            "value": 552,
            "unit": "allocs/op",
            "extra": "4448 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2",
            "value": 272132,
            "unit": "ns/op\t   56353 B/op\t     552 allocs/op",
            "extra": "4406 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2 - ns/op",
            "value": 272132,
            "unit": "ns/op",
            "extra": "4406 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2 - B/op",
            "value": 56353,
            "unit": "B/op",
            "extra": "4406 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2 - allocs/op",
            "value": 552,
            "unit": "allocs/op",
            "extra": "4406 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3",
            "value": 103524,
            "unit": "ns/op\t   29752 B/op\t     264 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3 - ns/op",
            "value": 103524,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3 - B/op",
            "value": 29752,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3 - allocs/op",
            "value": 264,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile",
            "value": 14092,
            "unit": "ns/op\t   10760 B/op\t     131 allocs/op",
            "extra": "93123 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile - ns/op",
            "value": 14092,
            "unit": "ns/op",
            "extra": "93123 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile - B/op",
            "value": 10760,
            "unit": "B/op",
            "extra": "93123 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile - allocs/op",
            "value": 131,
            "unit": "allocs/op",
            "extra": "93123 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid",
            "value": 26.05,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "44854726 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid - ns/op",
            "value": 26.05,
            "unit": "ns/op",
            "extra": "44854726 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "44854726 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "44854726 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite",
            "value": 50111,
            "unit": "ns/op\t   34696 B/op\t     226 allocs/op",
            "extra": "22468 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite - ns/op",
            "value": 50111,
            "unit": "ns/op",
            "extra": "22468 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite - B/op",
            "value": 34696,
            "unit": "B/op",
            "extra": "22468 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite - allocs/op",
            "value": 226,
            "unit": "allocs/op",
            "extra": "22468 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite",
            "value": 208067,
            "unit": "ns/op\t   54680 B/op\t    2069 allocs/op",
            "extra": "5754 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite - ns/op",
            "value": 208067,
            "unit": "ns/op",
            "extra": "5754 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite - B/op",
            "value": 54680,
            "unit": "B/op",
            "extra": "5754 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite - allocs/op",
            "value": 2069,
            "unit": "allocs/op",
            "extra": "5754 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite",
            "value": 154513,
            "unit": "ns/op\t   61104 B/op\t     721 allocs/op",
            "extra": "7227 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite - ns/op",
            "value": 154513,
            "unit": "ns/op",
            "extra": "7227 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite - B/op",
            "value": 61104,
            "unit": "B/op",
            "extra": "7227 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite - allocs/op",
            "value": 721,
            "unit": "allocs/op",
            "extra": "7227 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDIATWrite",
            "value": 231380104,
            "unit": "1210428822609230244A094101Federal Reserve Bank   My Bank Name                   ",
            "extra": "101 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDIATWrite - 1210428822609230244A094101Federal",
            "value": 231380104,
            "unit": "1210428822609230244A094101Federal",
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
          }
        ]
      },
      {
        "commit": {
          "author": {
            "name": "Faizan Khan",
            "username": "jellyfishing2346",
            "email": "108194702+jellyfishing2346@users.noreply.github.com"
          },
          "committer": {
            "name": "GitHub",
            "username": "web-flow",
            "email": "noreply@github.com"
          },
          "id": "b2687e19212cedc7cdf33692200bbc9919c65caa",
          "message": "docs: add comprehensive OpenAPI Guide to README (#1868)\n\nAdded a detailed OpenAPI Guide section to the README to complement the API spec enhancement work. The guide includes:\n\n- Overview of the OpenAPI specification and its uses\n- Instructions for generating client SDKs using various tools\n- Interactive documentation options (Swagger UI, Redoc)\n- API testing workflows with Postman/Insomnia\n- Key API endpoints overview\n- Validation options documentation\n- Practical example of creating an ACH file via the API\n\nThis documentation helps users understand how to leverage the OpenAPI specification for SDK generation, API testing, and integration work.\n\nAddresses #1584",
          "timestamp": "2026-09-22T14:35:02Z",
          "url": "https://github.com/moov-io/ach/commit/b2687e19212cedc7cdf33692200bbc9919c65caa"
        },
        "date": 1790131478393,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkAlphaFieldShort",
            "value": 955.8,
            "unit": "ns/op\t      80 B/op\t       3 allocs/op",
            "extra": "1251704 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldShort - ns/op",
            "value": 955.8,
            "unit": "ns/op",
            "extra": "1251704 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldShort - B/op",
            "value": 80,
            "unit": "B/op",
            "extra": "1251704 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldShort - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "1251704 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong",
            "value": 100.1,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "12012879 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong - ns/op",
            "value": 100.1,
            "unit": "ns/op",
            "extra": "12012879 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "12012879 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "12012879 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort",
            "value": 56.38,
            "unit": "ns/op\t      16 B/op\t       2 allocs/op",
            "extra": "20590659 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort - ns/op",
            "value": 56.38,
            "unit": "ns/op",
            "extra": "20590659 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "20590659 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "20590659 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong",
            "value": 25.41,
            "unit": "ns/op\t       8 B/op\t       1 allocs/op",
            "extra": "44024731 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong - ns/op",
            "value": 25.41,
            "unit": "ns/op",
            "extra": "44024731 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong - B/op",
            "value": 8,
            "unit": "B/op",
            "extra": "44024731 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "44024731 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField",
            "value": 14.96,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "83507521 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField - ns/op",
            "value": 14.96,
            "unit": "ns/op",
            "extra": "83507521 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "83507521 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "83507521 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField",
            "value": 5.611,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "213937585 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField - ns/op",
            "value": 5.611,
            "unit": "ns/op",
            "extra": "213937585 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "213937585 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "213937585 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles",
            "value": 249592,
            "unit": "ns/op\t   54149 B/op\t     306 allocs/op",
            "extra": "5070 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles - ns/op",
            "value": 249592,
            "unit": "ns/op",
            "extra": "5070 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles - B/op",
            "value": 54149,
            "unit": "B/op",
            "extra": "5070 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles - allocs/op",
            "value": 306,
            "unit": "allocs/op",
            "extra": "5070 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts",
            "value": 259870,
            "unit": "ns/op\t   54157 B/op\t     306 allocs/op",
            "extra": "4777 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts - ns/op",
            "value": 259870,
            "unit": "ns/op",
            "extra": "4777 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts - B/op",
            "value": 54157,
            "unit": "B/op",
            "extra": "4777 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts - allocs/op",
            "value": 306,
            "unit": "allocs/op",
            "extra": "4777 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir",
            "value": 148826,
            "unit": "ns/op\t   54554 B/op\t     310 allocs/op",
            "extra": "7489 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir - ns/op",
            "value": 148826,
            "unit": "ns/op",
            "extra": "7489 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir - B/op",
            "value": 54554,
            "unit": "B/op",
            "extra": "7489 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir - allocs/op",
            "value": 310,
            "unit": "allocs/op",
            "extra": "7489 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts",
            "value": 154659,
            "unit": "ns/op\t   54597 B/op\t     310 allocs/op",
            "extra": "6693 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts - ns/op",
            "value": 154659,
            "unit": "ns/op",
            "extra": "6693 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts - B/op",
            "value": 54597,
            "unit": "B/op",
            "extra": "6693 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts - allocs/op",
            "value": 310,
            "unit": "allocs/op",
            "extra": "6693 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups",
            "value": 311956,
            "unit": "ns/op\t   59852 B/op\t     366 allocs/op",
            "extra": "4035 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups - ns/op",
            "value": 311956,
            "unit": "ns/op",
            "extra": "4035 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups - B/op",
            "value": 59852,
            "unit": "B/op",
            "extra": "4035 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups - allocs/op",
            "value": 366,
            "unit": "allocs/op",
            "extra": "4035 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups",
            "value": 310028,
            "unit": "ns/op\t   59900 B/op\t     366 allocs/op",
            "extra": "3894 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups - ns/op",
            "value": 310028,
            "unit": "ns/op",
            "extra": "3894 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups - B/op",
            "value": 59900,
            "unit": "B/op",
            "extra": "3894 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups - allocs/op",
            "value": 366,
            "unit": "allocs/op",
            "extra": "3894 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups",
            "value": 303621,
            "unit": "ns/op\t   59843 B/op\t     366 allocs/op",
            "extra": "4242 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups - ns/op",
            "value": 303621,
            "unit": "ns/op",
            "extra": "4242 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups - B/op",
            "value": 59843,
            "unit": "B/op",
            "extra": "4242 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups - allocs/op",
            "value": 366,
            "unit": "allocs/op",
            "extra": "4242 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups",
            "value": 313004,
            "unit": "ns/op\t   59587 B/op\t     367 allocs/op",
            "extra": "4710 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups - ns/op",
            "value": 313004,
            "unit": "ns/op",
            "extra": "4710 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups - B/op",
            "value": 59587,
            "unit": "B/op",
            "extra": "4710 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups - allocs/op",
            "value": 367,
            "unit": "allocs/op",
            "extra": "4710 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader",
            "value": 248463479,
            "unit": "ns/op\t42450206 B/op\t  203654 allocs/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader - ns/op",
            "value": 248463479,
            "unit": "ns/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader - B/op",
            "value": 42450206,
            "unit": "B/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader - allocs/op",
            "value": 203654,
            "unit": "allocs/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator",
            "value": 83497034,
            "unit": "ns/op\t42424753 B/op\t  203601 allocs/op",
            "extra": "13 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator - ns/op",
            "value": 83497034,
            "unit": "ns/op",
            "extra": "13 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator - B/op",
            "value": 42424753,
            "unit": "B/op",
            "extra": "13 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator - allocs/op",
            "value": 203601,
            "unit": "allocs/op",
            "extra": "13 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader",
            "value": 453171973,
            "unit": "ns/op\t62131488 B/op\t  905791 allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader - ns/op",
            "value": 453171973,
            "unit": "ns/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader - B/op",
            "value": 62131488,
            "unit": "B/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader - allocs/op",
            "value": 905791,
            "unit": "allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator",
            "value": 155486981,
            "unit": "ns/op\t60508854 B/op\t  705840 allocs/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator - ns/op",
            "value": 155486981,
            "unit": "ns/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator - B/op",
            "value": 60508854,
            "unit": "B/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator - allocs/op",
            "value": 705840,
            "unit": "allocs/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader",
            "value": 1331188016,
            "unit": "ns/op\t215541432 B/op\t 1042868 allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader - ns/op",
            "value": 1331188016,
            "unit": "ns/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader - B/op",
            "value": 215541432,
            "unit": "B/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader - allocs/op",
            "value": 1042868,
            "unit": "allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator",
            "value": 412378907,
            "unit": "ns/op\t215424805 B/op\t 1042800 allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator - ns/op",
            "value": 412378907,
            "unit": "ns/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator - B/op",
            "value": 215424805,
            "unit": "B/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator - allocs/op",
            "value": 1042800,
            "unit": "allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader",
            "value": 2229312925,
            "unit": "ns/op\t313545192 B/op\t 4568035 allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader - ns/op",
            "value": 2229312925,
            "unit": "ns/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader - B/op",
            "value": 313545192,
            "unit": "B/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader - allocs/op",
            "value": 4568035,
            "unit": "allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator",
            "value": 763631428,
            "unit": "ns/op\t305430300 B/op\t 3568074 allocs/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator - ns/op",
            "value": 763631428,
            "unit": "ns/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator - B/op",
            "value": 305430300,
            "unit": "B/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator - allocs/op",
            "value": 3568074,
            "unit": "allocs/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead",
            "value": 41296,
            "unit": "ns/op\t   23312 B/op\t     114 allocs/op",
            "extra": "29355 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead - ns/op",
            "value": 41296,
            "unit": "ns/op",
            "extra": "29355 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead - B/op",
            "value": 23312,
            "unit": "B/op",
            "extra": "29355 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead - allocs/op",
            "value": 114,
            "unit": "allocs/op",
            "extra": "29355 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead",
            "value": 71627,
            "unit": "ns/op\t   26920 B/op\t     159 allocs/op",
            "extra": "16227 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead - ns/op",
            "value": 71627,
            "unit": "ns/op",
            "extra": "16227 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead - B/op",
            "value": 26920,
            "unit": "B/op",
            "extra": "16227 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead - allocs/op",
            "value": 159,
            "unit": "allocs/op",
            "extra": "16227 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead",
            "value": 31232,
            "unit": "ns/op\t   21848 B/op\t      77 allocs/op",
            "extra": "39646 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead - ns/op",
            "value": 31232,
            "unit": "ns/op",
            "extra": "39646 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead - B/op",
            "value": 21848,
            "unit": "B/op",
            "extra": "39646 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead - allocs/op",
            "value": 77,
            "unit": "allocs/op",
            "extra": "39646 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead",
            "value": 272352,
            "unit": "ns/op\t   56353 B/op\t     552 allocs/op",
            "extra": "4440 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead - ns/op",
            "value": 272352,
            "unit": "ns/op",
            "extra": "4440 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead - B/op",
            "value": 56353,
            "unit": "B/op",
            "extra": "4440 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead - allocs/op",
            "value": 552,
            "unit": "allocs/op",
            "extra": "4440 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2",
            "value": 270732,
            "unit": "ns/op\t   56353 B/op\t     552 allocs/op",
            "extra": "4453 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2 - ns/op",
            "value": 270732,
            "unit": "ns/op",
            "extra": "4453 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2 - B/op",
            "value": 56353,
            "unit": "B/op",
            "extra": "4453 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2 - allocs/op",
            "value": 552,
            "unit": "allocs/op",
            "extra": "4453 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3",
            "value": 102202,
            "unit": "ns/op\t   29752 B/op\t     264 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3 - ns/op",
            "value": 102202,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3 - B/op",
            "value": 29752,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3 - allocs/op",
            "value": 264,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile",
            "value": 13979,
            "unit": "ns/op\t   10760 B/op\t     131 allocs/op",
            "extra": "93356 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile - ns/op",
            "value": 13979,
            "unit": "ns/op",
            "extra": "93356 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile - B/op",
            "value": 10760,
            "unit": "B/op",
            "extra": "93356 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile - allocs/op",
            "value": 131,
            "unit": "allocs/op",
            "extra": "93356 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid",
            "value": 27.64,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "41297487 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid - ns/op",
            "value": 27.64,
            "unit": "ns/op",
            "extra": "41297487 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "41297487 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "41297487 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite",
            "value": 49370,
            "unit": "ns/op\t   34696 B/op\t     226 allocs/op",
            "extra": "22890 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite - ns/op",
            "value": 49370,
            "unit": "ns/op",
            "extra": "22890 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite - B/op",
            "value": 34696,
            "unit": "B/op",
            "extra": "22890 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite - allocs/op",
            "value": 226,
            "unit": "allocs/op",
            "extra": "22890 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite",
            "value": 212009,
            "unit": "ns/op\t   54680 B/op\t    2069 allocs/op",
            "extra": "5774 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite - ns/op",
            "value": 212009,
            "unit": "ns/op",
            "extra": "5774 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite - B/op",
            "value": 54680,
            "unit": "B/op",
            "extra": "5774 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite - allocs/op",
            "value": 2069,
            "unit": "allocs/op",
            "extra": "5774 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite",
            "value": 153415,
            "unit": "ns/op\t   61104 B/op\t     721 allocs/op",
            "extra": "7258 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite - ns/op",
            "value": 153415,
            "unit": "ns/op",
            "extra": "7258 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite - B/op",
            "value": 61104,
            "unit": "B/op",
            "extra": "7258 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite - allocs/op",
            "value": 721,
            "unit": "allocs/op",
            "extra": "7258 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDIATWrite",
            "value": 231380104,
            "unit": "1210428822609240244A094101Federal Reserve Bank   My Bank Name                   ",
            "extra": "101 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDIATWrite - 1210428822609240244A094101Federal",
            "value": 231380104,
            "unit": "1210428822609240244A094101Federal",
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
          }
        ]
      },
      {
        "commit": {
          "author": {
            "name": "Sash",
            "username": "SashaMIT",
            "email": "sash@ela.city"
          },
          "committer": {
            "name": "GitHub",
            "username": "web-flow",
            "email": "noreply@github.com"
          },
          "id": "d9f39c464da465b75a856355d1022013ed37ae94",
          "message": "fix: keep checking IAT addenda traces after a correction entry (#1869)\n\nisAddendaSequence returned success at the first Addenda98 entry, so a later entry with a mismatched addenda trace still passed Validate.",
          "timestamp": "2026-09-23T18:19:43Z",
          "url": "https://github.com/moov-io/ach/commit/d9f39c464da465b75a856355d1022013ed37ae94"
        },
        "date": 1790217254325,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkAlphaFieldShort",
            "value": 972.5,
            "unit": "ns/op\t      80 B/op\t       3 allocs/op",
            "extra": "1210798 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldShort - ns/op",
            "value": 972.5,
            "unit": "ns/op",
            "extra": "1210798 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldShort - B/op",
            "value": 80,
            "unit": "B/op",
            "extra": "1210798 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldShort - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "1210798 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong",
            "value": 98.55,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "12102511 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong - ns/op",
            "value": 98.55,
            "unit": "ns/op",
            "extra": "12102511 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "12102511 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "12102511 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort",
            "value": 57.4,
            "unit": "ns/op\t      16 B/op\t       2 allocs/op",
            "extra": "20602755 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort - ns/op",
            "value": 57.4,
            "unit": "ns/op",
            "extra": "20602755 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "20602755 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "20602755 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong",
            "value": 25.48,
            "unit": "ns/op\t       8 B/op\t       1 allocs/op",
            "extra": "45453770 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong - ns/op",
            "value": 25.48,
            "unit": "ns/op",
            "extra": "45453770 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong - B/op",
            "value": 8,
            "unit": "B/op",
            "extra": "45453770 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "45453770 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField",
            "value": 14.87,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "83573106 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField - ns/op",
            "value": 14.87,
            "unit": "ns/op",
            "extra": "83573106 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "83573106 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "83573106 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField",
            "value": 5.623,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "213552604 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField - ns/op",
            "value": 5.623,
            "unit": "ns/op",
            "extra": "213552604 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "213552604 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "213552604 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles",
            "value": 253034,
            "unit": "ns/op\t   54145 B/op\t     306 allocs/op",
            "extra": "5137 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles - ns/op",
            "value": 253034,
            "unit": "ns/op",
            "extra": "5137 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles - B/op",
            "value": 54145,
            "unit": "B/op",
            "extra": "5137 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles - allocs/op",
            "value": 306,
            "unit": "allocs/op",
            "extra": "5137 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts",
            "value": 256092,
            "unit": "ns/op\t   54154 B/op\t     306 allocs/op",
            "extra": "4444 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts - ns/op",
            "value": 256092,
            "unit": "ns/op",
            "extra": "4444 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts - B/op",
            "value": 54154,
            "unit": "B/op",
            "extra": "4444 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts - allocs/op",
            "value": 306,
            "unit": "allocs/op",
            "extra": "4444 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir",
            "value": 160819,
            "unit": "ns/op\t   54557 B/op\t     310 allocs/op",
            "extra": "7510 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir - ns/op",
            "value": 160819,
            "unit": "ns/op",
            "extra": "7510 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir - B/op",
            "value": 54557,
            "unit": "B/op",
            "extra": "7510 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir - allocs/op",
            "value": 310,
            "unit": "allocs/op",
            "extra": "7510 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts",
            "value": 151812,
            "unit": "ns/op\t   54591 B/op\t     310 allocs/op",
            "extra": "7772 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts - ns/op",
            "value": 151812,
            "unit": "ns/op",
            "extra": "7772 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts - B/op",
            "value": 54591,
            "unit": "B/op",
            "extra": "7772 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts - allocs/op",
            "value": 310,
            "unit": "allocs/op",
            "extra": "7772 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups",
            "value": 319296,
            "unit": "ns/op\t   59854 B/op\t     366 allocs/op",
            "extra": "4177 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups - ns/op",
            "value": 319296,
            "unit": "ns/op",
            "extra": "4177 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups - B/op",
            "value": 59854,
            "unit": "B/op",
            "extra": "4177 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups - allocs/op",
            "value": 366,
            "unit": "allocs/op",
            "extra": "4177 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups",
            "value": 307214,
            "unit": "ns/op\t   59892 B/op\t     366 allocs/op",
            "extra": "3906 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups - ns/op",
            "value": 307214,
            "unit": "ns/op",
            "extra": "3906 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups - B/op",
            "value": 59892,
            "unit": "B/op",
            "extra": "3906 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups - allocs/op",
            "value": 366,
            "unit": "allocs/op",
            "extra": "3906 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups",
            "value": 301339,
            "unit": "ns/op\t   59873 B/op\t     366 allocs/op",
            "extra": "4172 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups - ns/op",
            "value": 301339,
            "unit": "ns/op",
            "extra": "4172 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups - B/op",
            "value": 59873,
            "unit": "B/op",
            "extra": "4172 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups - allocs/op",
            "value": 366,
            "unit": "allocs/op",
            "extra": "4172 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups",
            "value": 291603,
            "unit": "ns/op\t   59571 B/op\t     367 allocs/op",
            "extra": "4810 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups - ns/op",
            "value": 291603,
            "unit": "ns/op",
            "extra": "4810 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups - B/op",
            "value": 59571,
            "unit": "B/op",
            "extra": "4810 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups - allocs/op",
            "value": 367,
            "unit": "allocs/op",
            "extra": "4810 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader",
            "value": 244991072,
            "unit": "ns/op\t42450206 B/op\t  203654 allocs/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader - ns/op",
            "value": 244991072,
            "unit": "ns/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader - B/op",
            "value": 42450206,
            "unit": "B/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader - allocs/op",
            "value": 203654,
            "unit": "allocs/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator",
            "value": 82680558,
            "unit": "ns/op\t42424617 B/op\t  203601 allocs/op",
            "extra": "13 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator - ns/op",
            "value": 82680558,
            "unit": "ns/op",
            "extra": "13 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator - B/op",
            "value": 42424617,
            "unit": "B/op",
            "extra": "13 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator - allocs/op",
            "value": 203601,
            "unit": "allocs/op",
            "extra": "13 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader",
            "value": 451392620,
            "unit": "ns/op\t62131461 B/op\t  905791 allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader - ns/op",
            "value": 451392620,
            "unit": "ns/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader - B/op",
            "value": 62131461,
            "unit": "B/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader - allocs/op",
            "value": 905791,
            "unit": "allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator",
            "value": 156408716,
            "unit": "ns/op\t60508733 B/op\t  705838 allocs/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator - ns/op",
            "value": 156408716,
            "unit": "ns/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator - B/op",
            "value": 60508733,
            "unit": "B/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator - allocs/op",
            "value": 705838,
            "unit": "allocs/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader",
            "value": 1304744333,
            "unit": "ns/op\t215541208 B/op\t 1042865 allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader - ns/op",
            "value": 1304744333,
            "unit": "ns/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader - B/op",
            "value": 215541208,
            "unit": "B/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader - allocs/op",
            "value": 1042865,
            "unit": "allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator",
            "value": 413356401,
            "unit": "ns/op\t215424784 B/op\t 1042800 allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator - ns/op",
            "value": 413356401,
            "unit": "ns/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator - B/op",
            "value": 215424784,
            "unit": "B/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator - allocs/op",
            "value": 1042800,
            "unit": "allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader",
            "value": 2233365877,
            "unit": "ns/op\t313545480 B/op\t 4568039 allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader - ns/op",
            "value": 2233365877,
            "unit": "ns/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader - B/op",
            "value": 313545480,
            "unit": "B/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader - allocs/op",
            "value": 4568039,
            "unit": "allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator",
            "value": 780916276,
            "unit": "ns/op\t305430308 B/op\t 3568074 allocs/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator - ns/op",
            "value": 780916276,
            "unit": "ns/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator - B/op",
            "value": 305430308,
            "unit": "B/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator - allocs/op",
            "value": 3568074,
            "unit": "allocs/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead",
            "value": 41532,
            "unit": "ns/op\t   23312 B/op\t     114 allocs/op",
            "extra": "29275 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead - ns/op",
            "value": 41532,
            "unit": "ns/op",
            "extra": "29275 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead - B/op",
            "value": 23312,
            "unit": "B/op",
            "extra": "29275 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead - allocs/op",
            "value": 114,
            "unit": "allocs/op",
            "extra": "29275 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead",
            "value": 71001,
            "unit": "ns/op\t   26920 B/op\t     159 allocs/op",
            "extra": "15973 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead - ns/op",
            "value": 71001,
            "unit": "ns/op",
            "extra": "15973 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead - B/op",
            "value": 26920,
            "unit": "B/op",
            "extra": "15973 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead - allocs/op",
            "value": 159,
            "unit": "allocs/op",
            "extra": "15973 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead",
            "value": 31767,
            "unit": "ns/op\t   21848 B/op\t      77 allocs/op",
            "extra": "40298 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead - ns/op",
            "value": 31767,
            "unit": "ns/op",
            "extra": "40298 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead - B/op",
            "value": 21848,
            "unit": "B/op",
            "extra": "40298 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead - allocs/op",
            "value": 77,
            "unit": "allocs/op",
            "extra": "40298 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead",
            "value": 272781,
            "unit": "ns/op\t   56353 B/op\t     552 allocs/op",
            "extra": "4502 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead - ns/op",
            "value": 272781,
            "unit": "ns/op",
            "extra": "4502 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead - B/op",
            "value": 56353,
            "unit": "B/op",
            "extra": "4502 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead - allocs/op",
            "value": 552,
            "unit": "allocs/op",
            "extra": "4502 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2",
            "value": 271018,
            "unit": "ns/op\t   56353 B/op\t     552 allocs/op",
            "extra": "4616 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2 - ns/op",
            "value": 271018,
            "unit": "ns/op",
            "extra": "4616 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2 - B/op",
            "value": 56353,
            "unit": "B/op",
            "extra": "4616 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2 - allocs/op",
            "value": 552,
            "unit": "allocs/op",
            "extra": "4616 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3",
            "value": 103343,
            "unit": "ns/op\t   29752 B/op\t     264 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3 - ns/op",
            "value": 103343,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3 - B/op",
            "value": 29752,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3 - allocs/op",
            "value": 264,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile",
            "value": 14212,
            "unit": "ns/op\t   10760 B/op\t     131 allocs/op",
            "extra": "93368 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile - ns/op",
            "value": 14212,
            "unit": "ns/op",
            "extra": "93368 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile - B/op",
            "value": 10760,
            "unit": "B/op",
            "extra": "93368 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile - allocs/op",
            "value": 131,
            "unit": "allocs/op",
            "extra": "93368 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid",
            "value": 26.09,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "45296545 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid - ns/op",
            "value": 26.09,
            "unit": "ns/op",
            "extra": "45296545 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "45296545 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "45296545 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite",
            "value": 50155,
            "unit": "ns/op\t   34696 B/op\t     226 allocs/op",
            "extra": "22710 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite - ns/op",
            "value": 50155,
            "unit": "ns/op",
            "extra": "22710 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite - B/op",
            "value": 34696,
            "unit": "B/op",
            "extra": "22710 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite - allocs/op",
            "value": 226,
            "unit": "allocs/op",
            "extra": "22710 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite",
            "value": 217955,
            "unit": "ns/op\t   54680 B/op\t    2069 allocs/op",
            "extra": "6004 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite - ns/op",
            "value": 217955,
            "unit": "ns/op",
            "extra": "6004 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite - B/op",
            "value": 54680,
            "unit": "B/op",
            "extra": "6004 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite - allocs/op",
            "value": 2069,
            "unit": "allocs/op",
            "extra": "6004 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite",
            "value": 152852,
            "unit": "ns/op\t   61104 B/op\t     721 allocs/op",
            "extra": "7437 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite - ns/op",
            "value": 152852,
            "unit": "ns/op",
            "extra": "7437 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite - B/op",
            "value": 61104,
            "unit": "B/op",
            "extra": "7437 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite - allocs/op",
            "value": 721,
            "unit": "allocs/op",
            "extra": "7437 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDIATWrite",
            "value": 231380104,
            "unit": "1210428822609250234A094101Federal Reserve Bank   My Bank Name                   ",
            "extra": "101 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDIATWrite - 1210428822609250234A094101Federal",
            "value": 231380104,
            "unit": "1210428822609250234A094101Federal",
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
          }
        ]
      },
      {
        "commit": {
          "author": {
            "name": "moov-bot",
            "email": "oss@moov.io"
          },
          "committer": {
            "name": "moov-bot",
            "email": "oss@moov.io"
          },
          "id": "4f426a55165331dcecf392c2b6b9ceb0b0d585a8",
          "message": "chore: updating wasm webui [skip ci]",
          "timestamp": "2026-09-24T22:18:55Z",
          "url": "https://github.com/moov-io/ach/commit/4f426a55165331dcecf392c2b6b9ceb0b0d585a8"
        },
        "date": 1790304690766,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkAlphaFieldShort",
            "value": 960.7,
            "unit": "ns/op\t      80 B/op\t       3 allocs/op",
            "extra": "1248410 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldShort - ns/op",
            "value": 960.7,
            "unit": "ns/op",
            "extra": "1248410 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldShort - B/op",
            "value": 80,
            "unit": "B/op",
            "extra": "1248410 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldShort - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "1248410 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong",
            "value": 98.1,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "11909674 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong - ns/op",
            "value": 98.1,
            "unit": "ns/op",
            "extra": "11909674 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "11909674 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "11909674 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort",
            "value": 56.89,
            "unit": "ns/op\t      16 B/op\t       2 allocs/op",
            "extra": "20696292 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort - ns/op",
            "value": 56.89,
            "unit": "ns/op",
            "extra": "20696292 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "20696292 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "20696292 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong",
            "value": 26.29,
            "unit": "ns/op\t       8 B/op\t       1 allocs/op",
            "extra": "44353462 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong - ns/op",
            "value": 26.29,
            "unit": "ns/op",
            "extra": "44353462 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong - B/op",
            "value": 8,
            "unit": "B/op",
            "extra": "44353462 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "44353462 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField",
            "value": 14.73,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "83827642 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField - ns/op",
            "value": 14.73,
            "unit": "ns/op",
            "extra": "83827642 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "83827642 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "83827642 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField",
            "value": 5.68,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "213512671 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField - ns/op",
            "value": 5.68,
            "unit": "ns/op",
            "extra": "213512671 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "213512671 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "213512671 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles",
            "value": 278882,
            "unit": "ns/op\t   54148 B/op\t     306 allocs/op",
            "extra": "3836 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles - ns/op",
            "value": 278882,
            "unit": "ns/op",
            "extra": "3836 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles - B/op",
            "value": 54148,
            "unit": "B/op",
            "extra": "3836 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles - allocs/op",
            "value": 306,
            "unit": "allocs/op",
            "extra": "3836 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts",
            "value": 261807,
            "unit": "ns/op\t   54156 B/op\t     306 allocs/op",
            "extra": "4453 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts - ns/op",
            "value": 261807,
            "unit": "ns/op",
            "extra": "4453 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts - B/op",
            "value": 54156,
            "unit": "B/op",
            "extra": "4453 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts - allocs/op",
            "value": 306,
            "unit": "allocs/op",
            "extra": "4453 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir",
            "value": 143664,
            "unit": "ns/op\t   54546 B/op\t     310 allocs/op",
            "extra": "7264 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir - ns/op",
            "value": 143664,
            "unit": "ns/op",
            "extra": "7264 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir - B/op",
            "value": 54546,
            "unit": "B/op",
            "extra": "7264 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir - allocs/op",
            "value": 310,
            "unit": "allocs/op",
            "extra": "7264 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts",
            "value": 154334,
            "unit": "ns/op\t   54587 B/op\t     310 allocs/op",
            "extra": "7170 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts - ns/op",
            "value": 154334,
            "unit": "ns/op",
            "extra": "7170 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts - B/op",
            "value": 54587,
            "unit": "B/op",
            "extra": "7170 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts - allocs/op",
            "value": 310,
            "unit": "allocs/op",
            "extra": "7170 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups",
            "value": 322397,
            "unit": "ns/op\t   59927 B/op\t     366 allocs/op",
            "extra": "3943 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups - ns/op",
            "value": 322397,
            "unit": "ns/op",
            "extra": "3943 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups - B/op",
            "value": 59927,
            "unit": "B/op",
            "extra": "3943 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups - allocs/op",
            "value": 366,
            "unit": "allocs/op",
            "extra": "3943 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups",
            "value": 304991,
            "unit": "ns/op\t   59929 B/op\t     366 allocs/op",
            "extra": "4072 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups - ns/op",
            "value": 304991,
            "unit": "ns/op",
            "extra": "4072 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups - B/op",
            "value": 59929,
            "unit": "B/op",
            "extra": "4072 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups - allocs/op",
            "value": 366,
            "unit": "allocs/op",
            "extra": "4072 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups",
            "value": 311387,
            "unit": "ns/op\t   59879 B/op\t     366 allocs/op",
            "extra": "4024 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups - ns/op",
            "value": 311387,
            "unit": "ns/op",
            "extra": "4024 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups - B/op",
            "value": 59879,
            "unit": "B/op",
            "extra": "4024 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups - allocs/op",
            "value": 366,
            "unit": "allocs/op",
            "extra": "4024 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups",
            "value": 291712,
            "unit": "ns/op\t   59534 B/op\t     367 allocs/op",
            "extra": "4975 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups - ns/op",
            "value": 291712,
            "unit": "ns/op",
            "extra": "4975 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups - B/op",
            "value": 59534,
            "unit": "B/op",
            "extra": "4975 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups - allocs/op",
            "value": 367,
            "unit": "allocs/op",
            "extra": "4975 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader",
            "value": 248637174,
            "unit": "ns/op\t42450206 B/op\t  203654 allocs/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader - ns/op",
            "value": 248637174,
            "unit": "ns/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader - B/op",
            "value": 42450206,
            "unit": "B/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader - allocs/op",
            "value": 203654,
            "unit": "allocs/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator",
            "value": 82741056,
            "unit": "ns/op\t42424886 B/op\t  203602 allocs/op",
            "extra": "13 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator - ns/op",
            "value": 82741056,
            "unit": "ns/op",
            "extra": "13 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator - B/op",
            "value": 42424886,
            "unit": "B/op",
            "extra": "13 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator - allocs/op",
            "value": 203602,
            "unit": "allocs/op",
            "extra": "13 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader",
            "value": 453894131,
            "unit": "ns/op\t62131642 B/op\t  905793 allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader - ns/op",
            "value": 453894131,
            "unit": "ns/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader - B/op",
            "value": 62131642,
            "unit": "B/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader - allocs/op",
            "value": 905793,
            "unit": "allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator",
            "value": 155243732,
            "unit": "ns/op\t60508875 B/op\t  705840 allocs/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator - ns/op",
            "value": 155243732,
            "unit": "ns/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator - B/op",
            "value": 60508875,
            "unit": "B/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator - allocs/op",
            "value": 705840,
            "unit": "allocs/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader",
            "value": 1309103151,
            "unit": "ns/op\t215541208 B/op\t 1042865 allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader - ns/op",
            "value": 1309103151,
            "unit": "ns/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader - B/op",
            "value": 215541208,
            "unit": "B/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader - allocs/op",
            "value": 1042865,
            "unit": "allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator",
            "value": 410937451,
            "unit": "ns/op\t215424853 B/op\t 1042801 allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator - ns/op",
            "value": 410937451,
            "unit": "ns/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator - B/op",
            "value": 215424853,
            "unit": "B/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator - allocs/op",
            "value": 1042801,
            "unit": "allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader",
            "value": 2237498623,
            "unit": "ns/op\t313545336 B/op\t 4568037 allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader - ns/op",
            "value": 2237498623,
            "unit": "ns/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader - B/op",
            "value": 313545336,
            "unit": "B/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader - allocs/op",
            "value": 4568037,
            "unit": "allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator",
            "value": 779560157,
            "unit": "ns/op\t305430228 B/op\t 3568073 allocs/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator - ns/op",
            "value": 779560157,
            "unit": "ns/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator - B/op",
            "value": 305430228,
            "unit": "B/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator - allocs/op",
            "value": 3568073,
            "unit": "allocs/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead",
            "value": 41744,
            "unit": "ns/op\t   23312 B/op\t     114 allocs/op",
            "extra": "28834 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead - ns/op",
            "value": 41744,
            "unit": "ns/op",
            "extra": "28834 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead - B/op",
            "value": 23312,
            "unit": "B/op",
            "extra": "28834 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead - allocs/op",
            "value": 114,
            "unit": "allocs/op",
            "extra": "28834 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead",
            "value": 71243,
            "unit": "ns/op\t   26920 B/op\t     159 allocs/op",
            "extra": "15690 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead - ns/op",
            "value": 71243,
            "unit": "ns/op",
            "extra": "15690 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead - B/op",
            "value": 26920,
            "unit": "B/op",
            "extra": "15690 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead - allocs/op",
            "value": 159,
            "unit": "allocs/op",
            "extra": "15690 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead",
            "value": 31300,
            "unit": "ns/op\t   21848 B/op\t      77 allocs/op",
            "extra": "39705 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead - ns/op",
            "value": 31300,
            "unit": "ns/op",
            "extra": "39705 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead - B/op",
            "value": 21848,
            "unit": "B/op",
            "extra": "39705 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead - allocs/op",
            "value": 77,
            "unit": "allocs/op",
            "extra": "39705 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead",
            "value": 275164,
            "unit": "ns/op\t   56353 B/op\t     552 allocs/op",
            "extra": "4490 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead - ns/op",
            "value": 275164,
            "unit": "ns/op",
            "extra": "4490 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead - B/op",
            "value": 56353,
            "unit": "B/op",
            "extra": "4490 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead - allocs/op",
            "value": 552,
            "unit": "allocs/op",
            "extra": "4490 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2",
            "value": 273234,
            "unit": "ns/op\t   56353 B/op\t     552 allocs/op",
            "extra": "4468 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2 - ns/op",
            "value": 273234,
            "unit": "ns/op",
            "extra": "4468 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2 - B/op",
            "value": 56353,
            "unit": "B/op",
            "extra": "4468 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2 - allocs/op",
            "value": 552,
            "unit": "allocs/op",
            "extra": "4468 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3",
            "value": 103682,
            "unit": "ns/op\t   29752 B/op\t     264 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3 - ns/op",
            "value": 103682,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3 - B/op",
            "value": 29752,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3 - allocs/op",
            "value": 264,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile",
            "value": 14127,
            "unit": "ns/op\t   10760 B/op\t     131 allocs/op",
            "extra": "93146 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile - ns/op",
            "value": 14127,
            "unit": "ns/op",
            "extra": "93146 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile - B/op",
            "value": 10760,
            "unit": "B/op",
            "extra": "93146 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile - allocs/op",
            "value": 131,
            "unit": "allocs/op",
            "extra": "93146 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid",
            "value": 26.04,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "44433564 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid - ns/op",
            "value": 26.04,
            "unit": "ns/op",
            "extra": "44433564 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "44433564 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "44433564 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite",
            "value": 49728,
            "unit": "ns/op\t   34696 B/op\t     226 allocs/op",
            "extra": "22864 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite - ns/op",
            "value": 49728,
            "unit": "ns/op",
            "extra": "22864 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite - B/op",
            "value": 34696,
            "unit": "B/op",
            "extra": "22864 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite - allocs/op",
            "value": 226,
            "unit": "allocs/op",
            "extra": "22864 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite",
            "value": 215736,
            "unit": "ns/op\t   54680 B/op\t    2069 allocs/op",
            "extra": "5884 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite - ns/op",
            "value": 215736,
            "unit": "ns/op",
            "extra": "5884 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite - B/op",
            "value": 54680,
            "unit": "B/op",
            "extra": "5884 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite - allocs/op",
            "value": 2069,
            "unit": "allocs/op",
            "extra": "5884 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite",
            "value": 153964,
            "unit": "ns/op\t   61104 B/op\t     721 allocs/op",
            "extra": "7378 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite - ns/op",
            "value": 153964,
            "unit": "ns/op",
            "extra": "7378 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite - B/op",
            "value": 61104,
            "unit": "B/op",
            "extra": "7378 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite - allocs/op",
            "value": 721,
            "unit": "allocs/op",
            "extra": "7378 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDIATWrite",
            "value": 231380104,
            "unit": "1210428822609260251A094101Federal Reserve Bank   My Bank Name                   ",
            "extra": "101 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDIATWrite - 1210428822609260251A094101Federal",
            "value": 231380104,
            "unit": "1210428822609260251A094101Federal",
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
          }
        ]
      },
      {
        "commit": {
          "author": {
            "name": "moov-bot",
            "email": "oss@moov.io"
          },
          "committer": {
            "name": "moov-bot",
            "email": "oss@moov.io"
          },
          "id": "7ee7ad03d7342e1f651c32db22fc8168c2b97cce",
          "message": "chore: updating wasm webui [skip ci]",
          "timestamp": "2026-09-25T14:06:50Z",
          "url": "https://github.com/moov-io/ach/commit/7ee7ad03d7342e1f651c32db22fc8168c2b97cce"
        },
        "date": 1790391279883,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkAlphaFieldShort",
            "value": 894.7,
            "unit": "ns/op\t      80 B/op\t       3 allocs/op",
            "extra": "1352413 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldShort - ns/op",
            "value": 894.7,
            "unit": "ns/op",
            "extra": "1352413 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldShort - B/op",
            "value": 80,
            "unit": "B/op",
            "extra": "1352413 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldShort - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "1352413 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong",
            "value": 99.59,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "11621517 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong - ns/op",
            "value": 99.59,
            "unit": "ns/op",
            "extra": "11621517 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "11621517 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "11621517 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort",
            "value": 55.19,
            "unit": "ns/op\t      16 B/op\t       2 allocs/op",
            "extra": "21350217 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort - ns/op",
            "value": 55.19,
            "unit": "ns/op",
            "extra": "21350217 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "21350217 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "21350217 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong",
            "value": 25.33,
            "unit": "ns/op\t       8 B/op\t       1 allocs/op",
            "extra": "45204044 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong - ns/op",
            "value": 25.33,
            "unit": "ns/op",
            "extra": "45204044 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong - B/op",
            "value": 8,
            "unit": "B/op",
            "extra": "45204044 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "45204044 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField",
            "value": 13.44,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "85256936 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField - ns/op",
            "value": 13.44,
            "unit": "ns/op",
            "extra": "85256936 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "85256936 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "85256936 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField",
            "value": 5.639,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "202452296 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField - ns/op",
            "value": 5.639,
            "unit": "ns/op",
            "extra": "202452296 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "202452296 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "202452296 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles",
            "value": 259945,
            "unit": "ns/op\t   54138 B/op\t     306 allocs/op",
            "extra": "5326 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles - ns/op",
            "value": 259945,
            "unit": "ns/op",
            "extra": "5326 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles - B/op",
            "value": 54138,
            "unit": "B/op",
            "extra": "5326 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles - allocs/op",
            "value": 306,
            "unit": "allocs/op",
            "extra": "5326 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts",
            "value": 247241,
            "unit": "ns/op\t   54159 B/op\t     306 allocs/op",
            "extra": "4720 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts - ns/op",
            "value": 247241,
            "unit": "ns/op",
            "extra": "4720 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts - B/op",
            "value": 54159,
            "unit": "B/op",
            "extra": "4720 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts - allocs/op",
            "value": 306,
            "unit": "allocs/op",
            "extra": "4720 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir",
            "value": 139667,
            "unit": "ns/op\t   54542 B/op\t     310 allocs/op",
            "extra": "8418 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir - ns/op",
            "value": 139667,
            "unit": "ns/op",
            "extra": "8418 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir - B/op",
            "value": 54542,
            "unit": "B/op",
            "extra": "8418 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir - allocs/op",
            "value": 310,
            "unit": "allocs/op",
            "extra": "8418 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts",
            "value": 140582,
            "unit": "ns/op\t   54588 B/op\t     310 allocs/op",
            "extra": "7971 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts - ns/op",
            "value": 140582,
            "unit": "ns/op",
            "extra": "7971 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts - B/op",
            "value": 54588,
            "unit": "B/op",
            "extra": "7971 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts - allocs/op",
            "value": 310,
            "unit": "allocs/op",
            "extra": "7971 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups",
            "value": 303217,
            "unit": "ns/op\t   59880 B/op\t     366 allocs/op",
            "extra": "4254 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups - ns/op",
            "value": 303217,
            "unit": "ns/op",
            "extra": "4254 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups - B/op",
            "value": 59880,
            "unit": "B/op",
            "extra": "4254 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups - allocs/op",
            "value": 366,
            "unit": "allocs/op",
            "extra": "4254 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups",
            "value": 304132,
            "unit": "ns/op\t   59875 B/op\t     366 allocs/op",
            "extra": "4399 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups - ns/op",
            "value": 304132,
            "unit": "ns/op",
            "extra": "4399 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups - B/op",
            "value": 59875,
            "unit": "B/op",
            "extra": "4399 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups - allocs/op",
            "value": 366,
            "unit": "allocs/op",
            "extra": "4399 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups",
            "value": 298073,
            "unit": "ns/op\t   59927 B/op\t     366 allocs/op",
            "extra": "4436 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups - ns/op",
            "value": 298073,
            "unit": "ns/op",
            "extra": "4436 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups - B/op",
            "value": 59927,
            "unit": "B/op",
            "extra": "4436 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups - allocs/op",
            "value": 366,
            "unit": "allocs/op",
            "extra": "4436 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups",
            "value": 290735,
            "unit": "ns/op\t   59474 B/op\t     367 allocs/op",
            "extra": "5164 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups - ns/op",
            "value": 290735,
            "unit": "ns/op",
            "extra": "5164 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups - B/op",
            "value": 59474,
            "unit": "B/op",
            "extra": "5164 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups - allocs/op",
            "value": 367,
            "unit": "allocs/op",
            "extra": "5164 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader",
            "value": 236795458,
            "unit": "ns/op\t42450225 B/op\t  203655 allocs/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader - ns/op",
            "value": 236795458,
            "unit": "ns/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader - B/op",
            "value": 42450225,
            "unit": "B/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader - allocs/op",
            "value": 203655,
            "unit": "allocs/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator",
            "value": 71923858,
            "unit": "ns/op\t42424766 B/op\t  203602 allocs/op",
            "extra": "15 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator - ns/op",
            "value": 71923858,
            "unit": "ns/op",
            "extra": "15 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator - B/op",
            "value": 42424766,
            "unit": "B/op",
            "extra": "15 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator - allocs/op",
            "value": 203602,
            "unit": "allocs/op",
            "extra": "15 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader",
            "value": 435248227,
            "unit": "ns/op\t62131760 B/op\t  905795 allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader - ns/op",
            "value": 435248227,
            "unit": "ns/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader - B/op",
            "value": 62131760,
            "unit": "B/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader - allocs/op",
            "value": 905795,
            "unit": "allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator",
            "value": 140529506,
            "unit": "ns/op\t60508822 B/op\t  705839 allocs/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator - ns/op",
            "value": 140529506,
            "unit": "ns/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator - B/op",
            "value": 60508822,
            "unit": "B/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator - allocs/op",
            "value": 705839,
            "unit": "allocs/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader",
            "value": 1315310378,
            "unit": "ns/op\t215541208 B/op\t 1042865 allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader - ns/op",
            "value": 1315310378,
            "unit": "ns/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader - B/op",
            "value": 215541208,
            "unit": "B/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader - allocs/op",
            "value": 1042865,
            "unit": "allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator",
            "value": 376588287,
            "unit": "ns/op\t215424949 B/op\t 1042802 allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator - ns/op",
            "value": 376588287,
            "unit": "ns/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator - B/op",
            "value": 215424949,
            "unit": "B/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator - allocs/op",
            "value": 1042802,
            "unit": "allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader",
            "value": 2152286578,
            "unit": "ns/op\t313545480 B/op\t 4568039 allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader - ns/op",
            "value": 2152286578,
            "unit": "ns/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader - B/op",
            "value": 313545480,
            "unit": "B/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader - allocs/op",
            "value": 4568039,
            "unit": "allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator",
            "value": 711424788,
            "unit": "ns/op\t305430228 B/op\t 3568073 allocs/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator - ns/op",
            "value": 711424788,
            "unit": "ns/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator - B/op",
            "value": 305430228,
            "unit": "B/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator - allocs/op",
            "value": 3568073,
            "unit": "allocs/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead",
            "value": 42065,
            "unit": "ns/op\t   23312 B/op\t     114 allocs/op",
            "extra": "28628 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead - ns/op",
            "value": 42065,
            "unit": "ns/op",
            "extra": "28628 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead - B/op",
            "value": 23312,
            "unit": "B/op",
            "extra": "28628 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead - allocs/op",
            "value": 114,
            "unit": "allocs/op",
            "extra": "28628 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead",
            "value": 70071,
            "unit": "ns/op\t   26920 B/op\t     159 allocs/op",
            "extra": "16555 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead - ns/op",
            "value": 70071,
            "unit": "ns/op",
            "extra": "16555 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead - B/op",
            "value": 26920,
            "unit": "B/op",
            "extra": "16555 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead - allocs/op",
            "value": 159,
            "unit": "allocs/op",
            "extra": "16555 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead",
            "value": 32035,
            "unit": "ns/op\t   21848 B/op\t      77 allocs/op",
            "extra": "38330 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead - ns/op",
            "value": 32035,
            "unit": "ns/op",
            "extra": "38330 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead - B/op",
            "value": 21848,
            "unit": "B/op",
            "extra": "38330 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead - allocs/op",
            "value": 77,
            "unit": "allocs/op",
            "extra": "38330 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead",
            "value": 260046,
            "unit": "ns/op\t   56353 B/op\t     552 allocs/op",
            "extra": "4749 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead - ns/op",
            "value": 260046,
            "unit": "ns/op",
            "extra": "4749 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead - B/op",
            "value": 56353,
            "unit": "B/op",
            "extra": "4749 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead - allocs/op",
            "value": 552,
            "unit": "allocs/op",
            "extra": "4749 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2",
            "value": 259375,
            "unit": "ns/op\t   56353 B/op\t     552 allocs/op",
            "extra": "4750 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2 - ns/op",
            "value": 259375,
            "unit": "ns/op",
            "extra": "4750 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2 - B/op",
            "value": 56353,
            "unit": "B/op",
            "extra": "4750 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2 - allocs/op",
            "value": 552,
            "unit": "allocs/op",
            "extra": "4750 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3",
            "value": 98171,
            "unit": "ns/op\t   29752 B/op\t     264 allocs/op",
            "extra": "12082 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3 - ns/op",
            "value": 98171,
            "unit": "ns/op",
            "extra": "12082 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3 - B/op",
            "value": 29752,
            "unit": "B/op",
            "extra": "12082 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3 - allocs/op",
            "value": 264,
            "unit": "allocs/op",
            "extra": "12082 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile",
            "value": 13550,
            "unit": "ns/op\t   10760 B/op\t     131 allocs/op",
            "extra": "97520 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile - ns/op",
            "value": 13550,
            "unit": "ns/op",
            "extra": "97520 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile - B/op",
            "value": 10760,
            "unit": "B/op",
            "extra": "97520 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile - allocs/op",
            "value": 131,
            "unit": "allocs/op",
            "extra": "97520 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid",
            "value": 24.95,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "48097746 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid - ns/op",
            "value": 24.95,
            "unit": "ns/op",
            "extra": "48097746 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "48097746 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "48097746 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite",
            "value": 47458,
            "unit": "ns/op\t   34696 B/op\t     226 allocs/op",
            "extra": "23482 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite - ns/op",
            "value": 47458,
            "unit": "ns/op",
            "extra": "23482 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite - B/op",
            "value": 34696,
            "unit": "B/op",
            "extra": "23482 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite - allocs/op",
            "value": 226,
            "unit": "allocs/op",
            "extra": "23482 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite",
            "value": 186396,
            "unit": "ns/op\t   54680 B/op\t    2069 allocs/op",
            "extra": "6513 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite - ns/op",
            "value": 186396,
            "unit": "ns/op",
            "extra": "6513 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite - B/op",
            "value": 54680,
            "unit": "B/op",
            "extra": "6513 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite - allocs/op",
            "value": 2069,
            "unit": "allocs/op",
            "extra": "6513 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite",
            "value": 148401,
            "unit": "ns/op\t   61104 B/op\t     721 allocs/op",
            "extra": "7737 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite - ns/op",
            "value": 148401,
            "unit": "ns/op",
            "extra": "7737 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite - B/op",
            "value": 61104,
            "unit": "B/op",
            "extra": "7737 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite - allocs/op",
            "value": 721,
            "unit": "allocs/op",
            "extra": "7737 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDIATWrite",
            "value": 231380104,
            "unit": "1210428822609270254A094101Federal Reserve Bank   My Bank Name                   ",
            "extra": "101 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDIATWrite - 1210428822609270254A094101Federal",
            "value": 231380104,
            "unit": "1210428822609270254A094101Federal",
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
          }
        ]
      },
      {
        "commit": {
          "author": {
            "name": "moov-bot",
            "email": "oss@moov.io"
          },
          "committer": {
            "name": "moov-bot",
            "email": "oss@moov.io"
          },
          "id": "7ee7ad03d7342e1f651c32db22fc8168c2b97cce",
          "message": "chore: updating wasm webui [skip ci]",
          "timestamp": "2026-09-25T14:06:50Z",
          "url": "https://github.com/moov-io/ach/commit/7ee7ad03d7342e1f651c32db22fc8168c2b97cce"
        },
        "date": 1790477820474,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkAlphaFieldShort",
            "value": 964.3,
            "unit": "ns/op\t      80 B/op\t       3 allocs/op",
            "extra": "1252572 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldShort - ns/op",
            "value": 964.3,
            "unit": "ns/op",
            "extra": "1252572 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldShort - B/op",
            "value": 80,
            "unit": "B/op",
            "extra": "1252572 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldShort - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "1252572 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong",
            "value": 98.59,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "12011187 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong - ns/op",
            "value": 98.59,
            "unit": "ns/op",
            "extra": "12011187 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "12011187 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "12011187 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort",
            "value": 57.04,
            "unit": "ns/op\t      16 B/op\t       2 allocs/op",
            "extra": "20926694 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort - ns/op",
            "value": 57.04,
            "unit": "ns/op",
            "extra": "20926694 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "20926694 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "20926694 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong",
            "value": 25.72,
            "unit": "ns/op\t       8 B/op\t       1 allocs/op",
            "extra": "44391064 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong - ns/op",
            "value": 25.72,
            "unit": "ns/op",
            "extra": "44391064 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong - B/op",
            "value": 8,
            "unit": "B/op",
            "extra": "44391064 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "44391064 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField",
            "value": 14.94,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "83789619 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField - ns/op",
            "value": 14.94,
            "unit": "ns/op",
            "extra": "83789619 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "83789619 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "83789619 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField",
            "value": 5.64,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "213079939 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField - ns/op",
            "value": 5.64,
            "unit": "ns/op",
            "extra": "213079939 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "213079939 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "213079939 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles",
            "value": 257890,
            "unit": "ns/op\t   54139 B/op\t     306 allocs/op",
            "extra": "4932 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles - ns/op",
            "value": 257890,
            "unit": "ns/op",
            "extra": "4932 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles - B/op",
            "value": 54139,
            "unit": "B/op",
            "extra": "4932 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles - allocs/op",
            "value": 306,
            "unit": "allocs/op",
            "extra": "4932 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts",
            "value": 262300,
            "unit": "ns/op\t   54155 B/op\t     306 allocs/op",
            "extra": "4932 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts - ns/op",
            "value": 262300,
            "unit": "ns/op",
            "extra": "4932 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts - B/op",
            "value": 54155,
            "unit": "B/op",
            "extra": "4932 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts - allocs/op",
            "value": 306,
            "unit": "allocs/op",
            "extra": "4932 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir",
            "value": 146294,
            "unit": "ns/op\t   54555 B/op\t     310 allocs/op",
            "extra": "7652 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir - ns/op",
            "value": 146294,
            "unit": "ns/op",
            "extra": "7652 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir - B/op",
            "value": 54555,
            "unit": "B/op",
            "extra": "7652 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir - allocs/op",
            "value": 310,
            "unit": "allocs/op",
            "extra": "7652 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts",
            "value": 150077,
            "unit": "ns/op\t   54588 B/op\t     310 allocs/op",
            "extra": "7908 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts - ns/op",
            "value": 150077,
            "unit": "ns/op",
            "extra": "7908 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts - B/op",
            "value": 54588,
            "unit": "B/op",
            "extra": "7908 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts - allocs/op",
            "value": 310,
            "unit": "allocs/op",
            "extra": "7908 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups",
            "value": 319423,
            "unit": "ns/op\t   59876 B/op\t     366 allocs/op",
            "extra": "3561 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups - ns/op",
            "value": 319423,
            "unit": "ns/op",
            "extra": "3561 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups - B/op",
            "value": 59876,
            "unit": "B/op",
            "extra": "3561 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups - allocs/op",
            "value": 366,
            "unit": "allocs/op",
            "extra": "3561 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups",
            "value": 311386,
            "unit": "ns/op\t   59928 B/op\t     366 allocs/op",
            "extra": "4050 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups - ns/op",
            "value": 311386,
            "unit": "ns/op",
            "extra": "4050 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups - B/op",
            "value": 59928,
            "unit": "B/op",
            "extra": "4050 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups - allocs/op",
            "value": 366,
            "unit": "allocs/op",
            "extra": "4050 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups",
            "value": 310203,
            "unit": "ns/op\t   59886 B/op\t     366 allocs/op",
            "extra": "4155 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups - ns/op",
            "value": 310203,
            "unit": "ns/op",
            "extra": "4155 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups - B/op",
            "value": 59886,
            "unit": "B/op",
            "extra": "4155 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups - allocs/op",
            "value": 366,
            "unit": "allocs/op",
            "extra": "4155 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups",
            "value": 293587,
            "unit": "ns/op\t   59637 B/op\t     367 allocs/op",
            "extra": "4681 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups - ns/op",
            "value": 293587,
            "unit": "ns/op",
            "extra": "4681 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups - B/op",
            "value": 59637,
            "unit": "B/op",
            "extra": "4681 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups - allocs/op",
            "value": 367,
            "unit": "allocs/op",
            "extra": "4681 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader",
            "value": 251418749,
            "unit": "ns/op\t42450742 B/op\t  203656 allocs/op",
            "extra": "4 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader - ns/op",
            "value": 251418749,
            "unit": "ns/op",
            "extra": "4 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader - B/op",
            "value": 42450742,
            "unit": "B/op",
            "extra": "4 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader - allocs/op",
            "value": 203656,
            "unit": "allocs/op",
            "extra": "4 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator",
            "value": 83086454,
            "unit": "ns/op\t42424753 B/op\t  203601 allocs/op",
            "extra": "13 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator - ns/op",
            "value": 83086454,
            "unit": "ns/op",
            "extra": "13 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator - B/op",
            "value": 42424753,
            "unit": "B/op",
            "extra": "13 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator - allocs/op",
            "value": 203601,
            "unit": "allocs/op",
            "extra": "13 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader",
            "value": 454707010,
            "unit": "ns/op\t62131498 B/op\t  905791 allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader - ns/op",
            "value": 454707010,
            "unit": "ns/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader - B/op",
            "value": 62131498,
            "unit": "B/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader - allocs/op",
            "value": 905791,
            "unit": "allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator",
            "value": 156909312,
            "unit": "ns/op\t60508809 B/op\t  705839 allocs/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator - ns/op",
            "value": 156909312,
            "unit": "ns/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator - B/op",
            "value": 60508809,
            "unit": "B/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator - allocs/op",
            "value": 705839,
            "unit": "allocs/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader",
            "value": 1331290987,
            "unit": "ns/op\t215541432 B/op\t 1042868 allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader - ns/op",
            "value": 1331290987,
            "unit": "ns/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader - B/op",
            "value": 215541432,
            "unit": "B/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader - allocs/op",
            "value": 1042868,
            "unit": "allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator",
            "value": 411273661,
            "unit": "ns/op\t215424901 B/op\t 1042801 allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator - ns/op",
            "value": 411273661,
            "unit": "ns/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator - B/op",
            "value": 215424901,
            "unit": "B/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator - allocs/op",
            "value": 1042801,
            "unit": "allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader",
            "value": 2225124225,
            "unit": "ns/op\t313545048 B/op\t 4568033 allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader - ns/op",
            "value": 2225124225,
            "unit": "ns/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader - B/op",
            "value": 313545048,
            "unit": "B/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader - allocs/op",
            "value": 4568033,
            "unit": "allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator",
            "value": 788269602,
            "unit": "ns/op\t305429972 B/op\t 3568069 allocs/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator - ns/op",
            "value": 788269602,
            "unit": "ns/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator - B/op",
            "value": 305429972,
            "unit": "B/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator - allocs/op",
            "value": 3568069,
            "unit": "allocs/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead",
            "value": 41604,
            "unit": "ns/op\t   23312 B/op\t     114 allocs/op",
            "extra": "29227 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead - ns/op",
            "value": 41604,
            "unit": "ns/op",
            "extra": "29227 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead - B/op",
            "value": 23312,
            "unit": "B/op",
            "extra": "29227 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead - allocs/op",
            "value": 114,
            "unit": "allocs/op",
            "extra": "29227 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead",
            "value": 72131,
            "unit": "ns/op\t   26920 B/op\t     159 allocs/op",
            "extra": "16051 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead - ns/op",
            "value": 72131,
            "unit": "ns/op",
            "extra": "16051 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead - B/op",
            "value": 26920,
            "unit": "B/op",
            "extra": "16051 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead - allocs/op",
            "value": 159,
            "unit": "allocs/op",
            "extra": "16051 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead",
            "value": 31268,
            "unit": "ns/op\t   21848 B/op\t      77 allocs/op",
            "extra": "39517 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead - ns/op",
            "value": 31268,
            "unit": "ns/op",
            "extra": "39517 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead - B/op",
            "value": 21848,
            "unit": "B/op",
            "extra": "39517 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead - allocs/op",
            "value": 77,
            "unit": "allocs/op",
            "extra": "39517 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead",
            "value": 275671,
            "unit": "ns/op\t   56353 B/op\t     552 allocs/op",
            "extra": "4275 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead - ns/op",
            "value": 275671,
            "unit": "ns/op",
            "extra": "4275 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead - B/op",
            "value": 56353,
            "unit": "B/op",
            "extra": "4275 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead - allocs/op",
            "value": 552,
            "unit": "allocs/op",
            "extra": "4275 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2",
            "value": 273061,
            "unit": "ns/op\t   56352 B/op\t     552 allocs/op",
            "extra": "4518 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2 - ns/op",
            "value": 273061,
            "unit": "ns/op",
            "extra": "4518 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2 - B/op",
            "value": 56352,
            "unit": "B/op",
            "extra": "4518 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2 - allocs/op",
            "value": 552,
            "unit": "allocs/op",
            "extra": "4518 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3",
            "value": 102992,
            "unit": "ns/op\t   29752 B/op\t     264 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3 - ns/op",
            "value": 102992,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3 - B/op",
            "value": 29752,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3 - allocs/op",
            "value": 264,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile",
            "value": 14193,
            "unit": "ns/op\t   10760 B/op\t     131 allocs/op",
            "extra": "91338 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile - ns/op",
            "value": 14193,
            "unit": "ns/op",
            "extra": "91338 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile - B/op",
            "value": 10760,
            "unit": "B/op",
            "extra": "91338 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile - allocs/op",
            "value": 131,
            "unit": "allocs/op",
            "extra": "91338 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid",
            "value": 26.12,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "42837072 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid - ns/op",
            "value": 26.12,
            "unit": "ns/op",
            "extra": "42837072 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "42837072 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "42837072 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite",
            "value": 50753,
            "unit": "ns/op\t   34696 B/op\t     226 allocs/op",
            "extra": "22704 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite - ns/op",
            "value": 50753,
            "unit": "ns/op",
            "extra": "22704 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite - B/op",
            "value": 34696,
            "unit": "B/op",
            "extra": "22704 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite - allocs/op",
            "value": 226,
            "unit": "allocs/op",
            "extra": "22704 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite",
            "value": 214889,
            "unit": "ns/op\t   54680 B/op\t    2069 allocs/op",
            "extra": "5802 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite - ns/op",
            "value": 214889,
            "unit": "ns/op",
            "extra": "5802 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite - B/op",
            "value": 54680,
            "unit": "B/op",
            "extra": "5802 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite - allocs/op",
            "value": 2069,
            "unit": "allocs/op",
            "extra": "5802 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite",
            "value": 153988,
            "unit": "ns/op\t   61104 B/op\t     721 allocs/op",
            "extra": "7191 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite - ns/op",
            "value": 153988,
            "unit": "ns/op",
            "extra": "7191 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite - B/op",
            "value": 61104,
            "unit": "B/op",
            "extra": "7191 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite - allocs/op",
            "value": 721,
            "unit": "allocs/op",
            "extra": "7191 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDIATWrite",
            "value": 231380104,
            "unit": "1210428822609280256A094101Federal Reserve Bank   My Bank Name                   ",
            "extra": "101 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDIATWrite - 1210428822609280256A094101Federal",
            "value": 231380104,
            "unit": "1210428822609280256A094101Federal",
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
          }
        ]
      },
      {
        "commit": {
          "author": {
            "name": "moov-bot",
            "email": "oss@moov.io"
          },
          "committer": {
            "name": "moov-bot",
            "email": "oss@moov.io"
          },
          "id": "7ee7ad03d7342e1f651c32db22fc8168c2b97cce",
          "message": "chore: updating wasm webui [skip ci]",
          "timestamp": "2026-09-25T14:06:50Z",
          "url": "https://github.com/moov-io/ach/commit/7ee7ad03d7342e1f651c32db22fc8168c2b97cce"
        },
        "date": 1790564182102,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkAlphaFieldShort",
            "value": 1067,
            "unit": "ns/op\t      80 B/op\t       3 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldShort - ns/op",
            "value": 1067,
            "unit": "ns/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldShort - B/op",
            "value": 80,
            "unit": "B/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldShort - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong",
            "value": 98.13,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "12053548 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong - ns/op",
            "value": 98.13,
            "unit": "ns/op",
            "extra": "12053548 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "12053548 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "12053548 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort",
            "value": 56.88,
            "unit": "ns/op\t      16 B/op\t       2 allocs/op",
            "extra": "20488716 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort - ns/op",
            "value": 56.88,
            "unit": "ns/op",
            "extra": "20488716 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "20488716 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "20488716 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong",
            "value": 25.59,
            "unit": "ns/op\t       8 B/op\t       1 allocs/op",
            "extra": "44997210 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong - ns/op",
            "value": 25.59,
            "unit": "ns/op",
            "extra": "44997210 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong - B/op",
            "value": 8,
            "unit": "B/op",
            "extra": "44997210 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "44997210 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField",
            "value": 14.97,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "83049849 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField - ns/op",
            "value": 14.97,
            "unit": "ns/op",
            "extra": "83049849 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "83049849 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "83049849 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField",
            "value": 5.623,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "213536313 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField - ns/op",
            "value": 5.623,
            "unit": "ns/op",
            "extra": "213536313 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "213536313 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "213536313 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles",
            "value": 245109,
            "unit": "ns/op\t   54142 B/op\t     306 allocs/op",
            "extra": "5194 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles - ns/op",
            "value": 245109,
            "unit": "ns/op",
            "extra": "5194 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles - B/op",
            "value": 54142,
            "unit": "B/op",
            "extra": "5194 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles - allocs/op",
            "value": 306,
            "unit": "allocs/op",
            "extra": "5194 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts",
            "value": 246728,
            "unit": "ns/op\t   54150 B/op\t     306 allocs/op",
            "extra": "4779 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts - ns/op",
            "value": 246728,
            "unit": "ns/op",
            "extra": "4779 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts - B/op",
            "value": 54150,
            "unit": "B/op",
            "extra": "4779 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts - allocs/op",
            "value": 306,
            "unit": "allocs/op",
            "extra": "4779 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir",
            "value": 145879,
            "unit": "ns/op\t   54556 B/op\t     310 allocs/op",
            "extra": "7483 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir - ns/op",
            "value": 145879,
            "unit": "ns/op",
            "extra": "7483 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir - B/op",
            "value": 54556,
            "unit": "B/op",
            "extra": "7483 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir - allocs/op",
            "value": 310,
            "unit": "allocs/op",
            "extra": "7483 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts",
            "value": 144365,
            "unit": "ns/op\t   54576 B/op\t     310 allocs/op",
            "extra": "7526 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts - ns/op",
            "value": 144365,
            "unit": "ns/op",
            "extra": "7526 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts - B/op",
            "value": 54576,
            "unit": "B/op",
            "extra": "7526 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts - allocs/op",
            "value": 310,
            "unit": "allocs/op",
            "extra": "7526 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups",
            "value": 318572,
            "unit": "ns/op\t   59874 B/op\t     366 allocs/op",
            "extra": "3783 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups - ns/op",
            "value": 318572,
            "unit": "ns/op",
            "extra": "3783 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups - B/op",
            "value": 59874,
            "unit": "B/op",
            "extra": "3783 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups - allocs/op",
            "value": 366,
            "unit": "allocs/op",
            "extra": "3783 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups",
            "value": 314518,
            "unit": "ns/op\t   59877 B/op\t     366 allocs/op",
            "extra": "4190 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups - ns/op",
            "value": 314518,
            "unit": "ns/op",
            "extra": "4190 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups - B/op",
            "value": 59877,
            "unit": "B/op",
            "extra": "4190 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups - allocs/op",
            "value": 366,
            "unit": "allocs/op",
            "extra": "4190 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups",
            "value": 298845,
            "unit": "ns/op\t   59869 B/op\t     366 allocs/op",
            "extra": "3942 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups - ns/op",
            "value": 298845,
            "unit": "ns/op",
            "extra": "3942 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups - B/op",
            "value": 59869,
            "unit": "B/op",
            "extra": "3942 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups - allocs/op",
            "value": 366,
            "unit": "allocs/op",
            "extra": "3942 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups",
            "value": 287523,
            "unit": "ns/op\t   59602 B/op\t     367 allocs/op",
            "extra": "4771 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups - ns/op",
            "value": 287523,
            "unit": "ns/op",
            "extra": "4771 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups - B/op",
            "value": 59602,
            "unit": "B/op",
            "extra": "4771 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups - allocs/op",
            "value": 367,
            "unit": "allocs/op",
            "extra": "4771 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader",
            "value": 244139061,
            "unit": "ns/op\t42450299 B/op\t  203656 allocs/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader - ns/op",
            "value": 244139061,
            "unit": "ns/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader - B/op",
            "value": 42450299,
            "unit": "B/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader - allocs/op",
            "value": 203656,
            "unit": "allocs/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator",
            "value": 83606928,
            "unit": "ns/op\t42424880 B/op\t  203602 allocs/op",
            "extra": "12 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator - ns/op",
            "value": 83606928,
            "unit": "ns/op",
            "extra": "12 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator - B/op",
            "value": 42424880,
            "unit": "B/op",
            "extra": "12 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator - allocs/op",
            "value": 203602,
            "unit": "allocs/op",
            "extra": "12 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader",
            "value": 458232706,
            "unit": "ns/op\t62131642 B/op\t  905793 allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader - ns/op",
            "value": 458232706,
            "unit": "ns/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader - B/op",
            "value": 62131642,
            "unit": "B/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader - allocs/op",
            "value": 905793,
            "unit": "allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator",
            "value": 156341610,
            "unit": "ns/op\t60508635 B/op\t  705837 allocs/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator - ns/op",
            "value": 156341610,
            "unit": "ns/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator - B/op",
            "value": 60508635,
            "unit": "B/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator - allocs/op",
            "value": 705837,
            "unit": "allocs/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader",
            "value": 1302687325,
            "unit": "ns/op\t215541576 B/op\t 1042870 allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader - ns/op",
            "value": 1302687325,
            "unit": "ns/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader - B/op",
            "value": 215541576,
            "unit": "B/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader - allocs/op",
            "value": 1042870,
            "unit": "allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator",
            "value": 416117391,
            "unit": "ns/op\t215424826 B/op\t 1042800 allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator - ns/op",
            "value": 416117391,
            "unit": "ns/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator - B/op",
            "value": 215424826,
            "unit": "B/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator - allocs/op",
            "value": 1042800,
            "unit": "allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader",
            "value": 2243209539,
            "unit": "ns/op\t313544888 B/op\t 4568031 allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader - ns/op",
            "value": 2243209539,
            "unit": "ns/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader - B/op",
            "value": 313544888,
            "unit": "B/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader - allocs/op",
            "value": 4568031,
            "unit": "allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator",
            "value": 785183625,
            "unit": "ns/op\t305429996 B/op\t 3568070 allocs/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator - ns/op",
            "value": 785183625,
            "unit": "ns/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator - B/op",
            "value": 305429996,
            "unit": "B/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator - allocs/op",
            "value": 3568070,
            "unit": "allocs/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead",
            "value": 41444,
            "unit": "ns/op\t   23312 B/op\t     114 allocs/op",
            "extra": "29134 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead - ns/op",
            "value": 41444,
            "unit": "ns/op",
            "extra": "29134 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead - B/op",
            "value": 23312,
            "unit": "B/op",
            "extra": "29134 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead - allocs/op",
            "value": 114,
            "unit": "allocs/op",
            "extra": "29134 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead",
            "value": 71841,
            "unit": "ns/op\t   26920 B/op\t     159 allocs/op",
            "extra": "15961 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead - ns/op",
            "value": 71841,
            "unit": "ns/op",
            "extra": "15961 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead - B/op",
            "value": 26920,
            "unit": "B/op",
            "extra": "15961 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead - allocs/op",
            "value": 159,
            "unit": "allocs/op",
            "extra": "15961 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead",
            "value": 30759,
            "unit": "ns/op\t   21848 B/op\t      77 allocs/op",
            "extra": "39351 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead - ns/op",
            "value": 30759,
            "unit": "ns/op",
            "extra": "39351 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead - B/op",
            "value": 21848,
            "unit": "B/op",
            "extra": "39351 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead - allocs/op",
            "value": 77,
            "unit": "allocs/op",
            "extra": "39351 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead",
            "value": 270890,
            "unit": "ns/op\t   56353 B/op\t     552 allocs/op",
            "extra": "4598 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead - ns/op",
            "value": 270890,
            "unit": "ns/op",
            "extra": "4598 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead - B/op",
            "value": 56353,
            "unit": "B/op",
            "extra": "4598 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead - allocs/op",
            "value": 552,
            "unit": "allocs/op",
            "extra": "4598 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2",
            "value": 272710,
            "unit": "ns/op\t   56353 B/op\t     552 allocs/op",
            "extra": "4496 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2 - ns/op",
            "value": 272710,
            "unit": "ns/op",
            "extra": "4496 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2 - B/op",
            "value": 56353,
            "unit": "B/op",
            "extra": "4496 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2 - allocs/op",
            "value": 552,
            "unit": "allocs/op",
            "extra": "4496 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3",
            "value": 102870,
            "unit": "ns/op\t   29752 B/op\t     264 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3 - ns/op",
            "value": 102870,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3 - B/op",
            "value": 29752,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3 - allocs/op",
            "value": 264,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile",
            "value": 14058,
            "unit": "ns/op\t   10760 B/op\t     131 allocs/op",
            "extra": "93254 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile - ns/op",
            "value": 14058,
            "unit": "ns/op",
            "extra": "93254 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile - B/op",
            "value": 10760,
            "unit": "B/op",
            "extra": "93254 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile - allocs/op",
            "value": 131,
            "unit": "allocs/op",
            "extra": "93254 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid",
            "value": 26.02,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "45886926 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid - ns/op",
            "value": 26.02,
            "unit": "ns/op",
            "extra": "45886926 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "45886926 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "45886926 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite",
            "value": 49279,
            "unit": "ns/op\t   34696 B/op\t     226 allocs/op",
            "extra": "22930 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite - ns/op",
            "value": 49279,
            "unit": "ns/op",
            "extra": "22930 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite - B/op",
            "value": 34696,
            "unit": "B/op",
            "extra": "22930 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite - allocs/op",
            "value": 226,
            "unit": "allocs/op",
            "extra": "22930 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite",
            "value": 210017,
            "unit": "ns/op\t   54680 B/op\t    2069 allocs/op",
            "extra": "5754 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite - ns/op",
            "value": 210017,
            "unit": "ns/op",
            "extra": "5754 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite - B/op",
            "value": 54680,
            "unit": "B/op",
            "extra": "5754 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite - allocs/op",
            "value": 2069,
            "unit": "allocs/op",
            "extra": "5754 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite",
            "value": 153791,
            "unit": "ns/op\t   61104 B/op\t     721 allocs/op",
            "extra": "7266 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite - ns/op",
            "value": 153791,
            "unit": "ns/op",
            "extra": "7266 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite - B/op",
            "value": 61104,
            "unit": "B/op",
            "extra": "7266 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite - allocs/op",
            "value": 721,
            "unit": "allocs/op",
            "extra": "7266 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDIATWrite",
            "value": 231380104,
            "unit": "1210428822609290256A094101Federal Reserve Bank   My Bank Name                   ",
            "extra": "101 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDIATWrite - 1210428822609290256A094101Federal",
            "value": 231380104,
            "unit": "1210428822609290256A094101Federal",
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
          }
        ]
      },
      {
        "commit": {
          "author": {
            "name": "Sash",
            "username": "SashaMIT",
            "email": "sash.t.mitchell@gmail.com"
          },
          "committer": {
            "name": "GitHub",
            "username": "web-flow",
            "email": "noreply@github.com"
          },
          "id": "7cfa7dd4fb0b53376267b33f8004a6f60192e5ca",
          "message": "Reject an ACH file creation time after 23:59. (#1873)\n\nValidate accepted 2900 because the hour pattern allowed 24 through 29, and a four character time skipped the parse error.",
          "timestamp": "2026-09-28T20:11:29Z",
          "url": "https://github.com/moov-io/ach/commit/7cfa7dd4fb0b53376267b33f8004a6f60192e5ca"
        },
        "date": 1790653005146,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkAlphaFieldShort",
            "value": 676.7,
            "unit": "ns/op\t      80 B/op\t       3 allocs/op",
            "extra": "1768441 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldShort - ns/op",
            "value": 676.7,
            "unit": "ns/op",
            "extra": "1768441 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldShort - B/op",
            "value": 80,
            "unit": "B/op",
            "extra": "1768441 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldShort - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "1768441 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong",
            "value": 76.97,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "15365712 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong - ns/op",
            "value": 76.97,
            "unit": "ns/op",
            "extra": "15365712 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "15365712 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "15365712 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort",
            "value": 42.82,
            "unit": "ns/op\t      16 B/op\t       2 allocs/op",
            "extra": "27224754 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort - ns/op",
            "value": 42.82,
            "unit": "ns/op",
            "extra": "27224754 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "27224754 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "27224754 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong",
            "value": 19.71,
            "unit": "ns/op\t       8 B/op\t       1 allocs/op",
            "extra": "57395767 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong - ns/op",
            "value": 19.71,
            "unit": "ns/op",
            "extra": "57395767 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong - B/op",
            "value": 8,
            "unit": "B/op",
            "extra": "57395767 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "57395767 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField",
            "value": 10.43,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "100000000 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField - ns/op",
            "value": 10.43,
            "unit": "ns/op",
            "extra": "100000000 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "100000000 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "100000000 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField",
            "value": 4.383,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "262181206 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField - ns/op",
            "value": 4.383,
            "unit": "ns/op",
            "extra": "262181206 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "262181206 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "262181206 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles",
            "value": 200244,
            "unit": "ns/op\t   54145 B/op\t     306 allocs/op",
            "extra": "6746 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles - ns/op",
            "value": 200244,
            "unit": "ns/op",
            "extra": "6746 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles - B/op",
            "value": 54145,
            "unit": "B/op",
            "extra": "6746 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles - allocs/op",
            "value": 306,
            "unit": "allocs/op",
            "extra": "6746 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts",
            "value": 202871,
            "unit": "ns/op\t   54147 B/op\t     306 allocs/op",
            "extra": "5961 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts - ns/op",
            "value": 202871,
            "unit": "ns/op",
            "extra": "5961 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts - B/op",
            "value": 54147,
            "unit": "B/op",
            "extra": "5961 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts - allocs/op",
            "value": 306,
            "unit": "allocs/op",
            "extra": "5961 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir",
            "value": 122721,
            "unit": "ns/op\t   54556 B/op\t     310 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir - ns/op",
            "value": 122721,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir - B/op",
            "value": 54556,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir - allocs/op",
            "value": 310,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts",
            "value": 120616,
            "unit": "ns/op\t   54592 B/op\t     310 allocs/op",
            "extra": "9656 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts - ns/op",
            "value": 120616,
            "unit": "ns/op",
            "extra": "9656 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts - B/op",
            "value": 54592,
            "unit": "B/op",
            "extra": "9656 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts - allocs/op",
            "value": 310,
            "unit": "allocs/op",
            "extra": "9656 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups",
            "value": 281369,
            "unit": "ns/op\t   59849 B/op\t     366 allocs/op",
            "extra": "5097 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups - ns/op",
            "value": 281369,
            "unit": "ns/op",
            "extra": "5097 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups - B/op",
            "value": 59849,
            "unit": "B/op",
            "extra": "5097 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups - allocs/op",
            "value": 366,
            "unit": "allocs/op",
            "extra": "5097 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups",
            "value": 242833,
            "unit": "ns/op\t   59882 B/op\t     366 allocs/op",
            "extra": "5422 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups - ns/op",
            "value": 242833,
            "unit": "ns/op",
            "extra": "5422 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups - B/op",
            "value": 59882,
            "unit": "B/op",
            "extra": "5422 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups - allocs/op",
            "value": 366,
            "unit": "allocs/op",
            "extra": "5422 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups",
            "value": 235124,
            "unit": "ns/op\t   59877 B/op\t     366 allocs/op",
            "extra": "5604 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups - ns/op",
            "value": 235124,
            "unit": "ns/op",
            "extra": "5604 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups - B/op",
            "value": 59877,
            "unit": "B/op",
            "extra": "5604 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups - allocs/op",
            "value": 366,
            "unit": "allocs/op",
            "extra": "5604 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups",
            "value": 240591,
            "unit": "ns/op\t   59796 B/op\t     367 allocs/op",
            "extra": "6067 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups - ns/op",
            "value": 240591,
            "unit": "ns/op",
            "extra": "6067 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups - B/op",
            "value": 59796,
            "unit": "B/op",
            "extra": "6067 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups - allocs/op",
            "value": 367,
            "unit": "allocs/op",
            "extra": "6067 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader",
            "value": 192408262,
            "unit": "ns/op\t42449958 B/op\t  203654 allocs/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader - ns/op",
            "value": 192408262,
            "unit": "ns/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader - B/op",
            "value": 42449958,
            "unit": "B/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader - allocs/op",
            "value": 203654,
            "unit": "allocs/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator",
            "value": 55886946,
            "unit": "ns/op\t42424703 B/op\t  203601 allocs/op",
            "extra": "19 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator - ns/op",
            "value": 55886946,
            "unit": "ns/op",
            "extra": "19 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator - B/op",
            "value": 42424703,
            "unit": "B/op",
            "extra": "19 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator - allocs/op",
            "value": 203601,
            "unit": "allocs/op",
            "extra": "19 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader",
            "value": 343608126,
            "unit": "ns/op\t62131749 B/op\t  905793 allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader - ns/op",
            "value": 343608126,
            "unit": "ns/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader - B/op",
            "value": 62131749,
            "unit": "B/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader - allocs/op",
            "value": 905793,
            "unit": "allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator",
            "value": 107444476,
            "unit": "ns/op\t60508742 B/op\t  705838 allocs/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator - ns/op",
            "value": 107444476,
            "unit": "ns/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator - B/op",
            "value": 60508742,
            "unit": "B/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator - allocs/op",
            "value": 705838,
            "unit": "allocs/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader",
            "value": 1092351818,
            "unit": "ns/op\t215542216 B/op\t 1042873 allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader - ns/op",
            "value": 1092351818,
            "unit": "ns/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader - B/op",
            "value": 215542216,
            "unit": "B/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader - allocs/op",
            "value": 1042873,
            "unit": "allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator",
            "value": 287302400,
            "unit": "ns/op\t215424622 B/op\t 1042798 allocs/op",
            "extra": "4 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator - ns/op",
            "value": 287302400,
            "unit": "ns/op",
            "extra": "4 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator - B/op",
            "value": 215424622,
            "unit": "B/op",
            "extra": "4 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator - allocs/op",
            "value": 1042798,
            "unit": "allocs/op",
            "extra": "4 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader",
            "value": 1679620334,
            "unit": "ns/op\t313545976 B/op\t 4568040 allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader - ns/op",
            "value": 1679620334,
            "unit": "ns/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader - B/op",
            "value": 313545976,
            "unit": "B/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader - allocs/op",
            "value": 4568040,
            "unit": "allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator",
            "value": 539467371,
            "unit": "ns/op\t305430068 B/op\t 3568071 allocs/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator - ns/op",
            "value": 539467371,
            "unit": "ns/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator - B/op",
            "value": 305430068,
            "unit": "B/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator - allocs/op",
            "value": 3568071,
            "unit": "allocs/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead",
            "value": 33919,
            "unit": "ns/op\t   23312 B/op\t     114 allocs/op",
            "extra": "35950 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead - ns/op",
            "value": 33919,
            "unit": "ns/op",
            "extra": "35950 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead - B/op",
            "value": 23312,
            "unit": "B/op",
            "extra": "35950 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead - allocs/op",
            "value": 114,
            "unit": "allocs/op",
            "extra": "35950 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead",
            "value": 56527,
            "unit": "ns/op\t   26920 B/op\t     159 allocs/op",
            "extra": "20857 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead - ns/op",
            "value": 56527,
            "unit": "ns/op",
            "extra": "20857 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead - B/op",
            "value": 26920,
            "unit": "B/op",
            "extra": "20857 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead - allocs/op",
            "value": 159,
            "unit": "allocs/op",
            "extra": "20857 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead",
            "value": 25402,
            "unit": "ns/op\t   21848 B/op\t      77 allocs/op",
            "extra": "47602 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead - ns/op",
            "value": 25402,
            "unit": "ns/op",
            "extra": "47602 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead - B/op",
            "value": 21848,
            "unit": "B/op",
            "extra": "47602 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead - allocs/op",
            "value": 77,
            "unit": "allocs/op",
            "extra": "47602 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead",
            "value": 199947,
            "unit": "ns/op\t   56352 B/op\t     552 allocs/op",
            "extra": "6010 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead - ns/op",
            "value": 199947,
            "unit": "ns/op",
            "extra": "6010 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead - B/op",
            "value": 56352,
            "unit": "B/op",
            "extra": "6010 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead - allocs/op",
            "value": 552,
            "unit": "allocs/op",
            "extra": "6010 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2",
            "value": 201961,
            "unit": "ns/op\t   56352 B/op\t     552 allocs/op",
            "extra": "6078 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2 - ns/op",
            "value": 201961,
            "unit": "ns/op",
            "extra": "6078 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2 - B/op",
            "value": 56352,
            "unit": "B/op",
            "extra": "6078 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2 - allocs/op",
            "value": 552,
            "unit": "allocs/op",
            "extra": "6078 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3",
            "value": 76172,
            "unit": "ns/op\t   29752 B/op\t     264 allocs/op",
            "extra": "15238 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3 - ns/op",
            "value": 76172,
            "unit": "ns/op",
            "extra": "15238 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3 - B/op",
            "value": 29752,
            "unit": "B/op",
            "extra": "15238 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3 - allocs/op",
            "value": 264,
            "unit": "allocs/op",
            "extra": "15238 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile",
            "value": 10520,
            "unit": "ns/op\t   10760 B/op\t     131 allocs/op",
            "extra": "121449 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile - ns/op",
            "value": 10520,
            "unit": "ns/op",
            "extra": "121449 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile - B/op",
            "value": 10760,
            "unit": "B/op",
            "extra": "121449 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile - allocs/op",
            "value": 131,
            "unit": "allocs/op",
            "extra": "121449 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid",
            "value": 19.82,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "62602710 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid - ns/op",
            "value": 19.82,
            "unit": "ns/op",
            "extra": "62602710 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "62602710 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "62602710 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite",
            "value": 38007,
            "unit": "ns/op\t   34696 B/op\t     226 allocs/op",
            "extra": "29492 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite - ns/op",
            "value": 38007,
            "unit": "ns/op",
            "extra": "29492 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite - B/op",
            "value": 34696,
            "unit": "B/op",
            "extra": "29492 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite - allocs/op",
            "value": 226,
            "unit": "allocs/op",
            "extra": "29492 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite",
            "value": 145416,
            "unit": "ns/op\t   54680 B/op\t    2069 allocs/op",
            "extra": "8232 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite - ns/op",
            "value": 145416,
            "unit": "ns/op",
            "extra": "8232 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite - B/op",
            "value": 54680,
            "unit": "B/op",
            "extra": "8232 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite - allocs/op",
            "value": 2069,
            "unit": "allocs/op",
            "extra": "8232 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite",
            "value": 117980,
            "unit": "ns/op\t   61104 B/op\t     721 allocs/op",
            "extra": "9642 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite - ns/op",
            "value": 117980,
            "unit": "ns/op",
            "extra": "9642 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite - B/op",
            "value": 61104,
            "unit": "B/op",
            "extra": "9642 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite - allocs/op",
            "value": 721,
            "unit": "allocs/op",
            "extra": "9642 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDIATWrite",
            "value": 231380104,
            "unit": "1210428822609300336A094101Federal Reserve Bank   My Bank Name                   ",
            "extra": "101 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDIATWrite - 1210428822609300336A094101Federal",
            "value": 231380104,
            "unit": "1210428822609300336A094101Federal",
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
          }
        ]
      },
      {
        "commit": {
          "author": {
            "name": "Sash",
            "username": "SashaMIT",
            "email": "sash.t.mitchell@gmail.com"
          },
          "committer": {
            "name": "GitHub",
            "username": "web-flow",
            "email": "noreply@github.com"
          },
          "id": "7cfa7dd4fb0b53376267b33f8004a6f60192e5ca",
          "message": "Reject an ACH file creation time after 23:59. (#1873)\n\nValidate accepted 2900 because the hour pattern allowed 24 through 29, and a four character time skipped the parse error.",
          "timestamp": "2026-09-28T20:11:29Z",
          "url": "https://github.com/moov-io/ach/commit/7cfa7dd4fb0b53376267b33f8004a6f60192e5ca"
        },
        "date": 1790738570826,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkAlphaFieldShort",
            "value": 577.7,
            "unit": "ns/op\t      80 B/op\t       3 allocs/op",
            "extra": "2082715 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldShort - ns/op",
            "value": 577.7,
            "unit": "ns/op",
            "extra": "2082715 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldShort - B/op",
            "value": 80,
            "unit": "B/op",
            "extra": "2082715 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldShort - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "2082715 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong",
            "value": 63.17,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "18441002 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong - ns/op",
            "value": 63.17,
            "unit": "ns/op",
            "extra": "18441002 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "18441002 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "18441002 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort",
            "value": 34.63,
            "unit": "ns/op\t      16 B/op\t       2 allocs/op",
            "extra": "33768109 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort - ns/op",
            "value": 34.63,
            "unit": "ns/op",
            "extra": "33768109 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "33768109 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "33768109 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong",
            "value": 17.17,
            "unit": "ns/op\t       8 B/op\t       1 allocs/op",
            "extra": "67896798 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong - ns/op",
            "value": 17.17,
            "unit": "ns/op",
            "extra": "67896798 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong - B/op",
            "value": 8,
            "unit": "B/op",
            "extra": "67896798 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "67896798 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField",
            "value": 9.499,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "126226389 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField - ns/op",
            "value": 9.499,
            "unit": "ns/op",
            "extra": "126226389 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "126226389 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "126226389 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField",
            "value": 2.889,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "415342389 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField - ns/op",
            "value": 2.889,
            "unit": "ns/op",
            "extra": "415342389 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "415342389 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "415342389 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles",
            "value": 165954,
            "unit": "ns/op\t   54146 B/op\t     306 allocs/op",
            "extra": "8151 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles - ns/op",
            "value": 165954,
            "unit": "ns/op",
            "extra": "8151 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles - B/op",
            "value": 54146,
            "unit": "B/op",
            "extra": "8151 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles - allocs/op",
            "value": 306,
            "unit": "allocs/op",
            "extra": "8151 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts",
            "value": 158158,
            "unit": "ns/op\t   54164 B/op\t     306 allocs/op",
            "extra": "7305 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts - ns/op",
            "value": 158158,
            "unit": "ns/op",
            "extra": "7305 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts - B/op",
            "value": 54164,
            "unit": "B/op",
            "extra": "7305 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts - allocs/op",
            "value": 306,
            "unit": "allocs/op",
            "extra": "7305 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir",
            "value": 101324,
            "unit": "ns/op\t   54549 B/op\t     310 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir - ns/op",
            "value": 101324,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir - B/op",
            "value": 54549,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir - allocs/op",
            "value": 310,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts",
            "value": 108058,
            "unit": "ns/op\t   54586 B/op\t     310 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts - ns/op",
            "value": 108058,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts - B/op",
            "value": 54586,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts - allocs/op",
            "value": 310,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups",
            "value": 211350,
            "unit": "ns/op\t   59864 B/op\t     366 allocs/op",
            "extra": "5938 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups - ns/op",
            "value": 211350,
            "unit": "ns/op",
            "extra": "5938 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups - B/op",
            "value": 59864,
            "unit": "B/op",
            "extra": "5938 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups - allocs/op",
            "value": 366,
            "unit": "allocs/op",
            "extra": "5938 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups",
            "value": 206271,
            "unit": "ns/op\t   59864 B/op\t     366 allocs/op",
            "extra": "6332 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups - ns/op",
            "value": 206271,
            "unit": "ns/op",
            "extra": "6332 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups - B/op",
            "value": 59864,
            "unit": "B/op",
            "extra": "6332 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups - allocs/op",
            "value": 366,
            "unit": "allocs/op",
            "extra": "6332 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups",
            "value": 195874,
            "unit": "ns/op\t   59849 B/op\t     366 allocs/op",
            "extra": "5923 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups - ns/op",
            "value": 195874,
            "unit": "ns/op",
            "extra": "5923 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups - B/op",
            "value": 59849,
            "unit": "B/op",
            "extra": "5923 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups - allocs/op",
            "value": 366,
            "unit": "allocs/op",
            "extra": "5923 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups",
            "value": 188839,
            "unit": "ns/op\t   60065 B/op\t     367 allocs/op",
            "extra": "7660 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups - ns/op",
            "value": 188839,
            "unit": "ns/op",
            "extra": "7660 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups - B/op",
            "value": 60065,
            "unit": "B/op",
            "extra": "7660 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups - allocs/op",
            "value": 367,
            "unit": "allocs/op",
            "extra": "7660 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader",
            "value": 159957395,
            "unit": "ns/op\t42449820 B/op\t  203655 allocs/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader - ns/op",
            "value": 159957395,
            "unit": "ns/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader - B/op",
            "value": 42449820,
            "unit": "B/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader - allocs/op",
            "value": 203655,
            "unit": "allocs/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator",
            "value": 53036119,
            "unit": "ns/op\t42424649 B/op\t  203601 allocs/op",
            "extra": "21 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator - ns/op",
            "value": 53036119,
            "unit": "ns/op",
            "extra": "21 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator - B/op",
            "value": 42424649,
            "unit": "B/op",
            "extra": "21 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator - allocs/op",
            "value": 203601,
            "unit": "allocs/op",
            "extra": "21 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader",
            "value": 287297786,
            "unit": "ns/op\t62131184 B/op\t  905790 allocs/op",
            "extra": "4 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader - ns/op",
            "value": 287297786,
            "unit": "ns/op",
            "extra": "4 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader - B/op",
            "value": 62131184,
            "unit": "B/op",
            "extra": "4 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader - allocs/op",
            "value": 905790,
            "unit": "allocs/op",
            "extra": "4 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator",
            "value": 97617140,
            "unit": "ns/op\t60508764 B/op\t  705838 allocs/op",
            "extra": "12 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator - ns/op",
            "value": 97617140,
            "unit": "ns/op",
            "extra": "12 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator - B/op",
            "value": 60508764,
            "unit": "B/op",
            "extra": "12 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator - allocs/op",
            "value": 705838,
            "unit": "allocs/op",
            "extra": "12 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader",
            "value": 923521322,
            "unit": "ns/op\t215540544 B/op\t 1042861 allocs/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader - ns/op",
            "value": 923521322,
            "unit": "ns/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader - B/op",
            "value": 215540544,
            "unit": "B/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader - allocs/op",
            "value": 1042861,
            "unit": "allocs/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator",
            "value": 264232911,
            "unit": "ns/op\t215424470 B/op\t 1042796 allocs/op",
            "extra": "4 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator - ns/op",
            "value": 264232911,
            "unit": "ns/op",
            "extra": "4 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator - B/op",
            "value": 215424470,
            "unit": "B/op",
            "extra": "4 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator - allocs/op",
            "value": 1042796,
            "unit": "allocs/op",
            "extra": "4 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader",
            "value": 1419752075,
            "unit": "ns/op\t313545848 B/op\t 4568038 allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader - ns/op",
            "value": 1419752075,
            "unit": "ns/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader - B/op",
            "value": 313545848,
            "unit": "B/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader - allocs/op",
            "value": 4568038,
            "unit": "allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator",
            "value": 491452583,
            "unit": "ns/op\t305430090 B/op\t 3568072 allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator - ns/op",
            "value": 491452583,
            "unit": "ns/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator - B/op",
            "value": 305430090,
            "unit": "B/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator - allocs/op",
            "value": 3568072,
            "unit": "allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead",
            "value": 25144,
            "unit": "ns/op\t   23312 B/op\t     114 allocs/op",
            "extra": "50691 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead - ns/op",
            "value": 25144,
            "unit": "ns/op",
            "extra": "50691 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead - B/op",
            "value": 23312,
            "unit": "B/op",
            "extra": "50691 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead - allocs/op",
            "value": 114,
            "unit": "allocs/op",
            "extra": "50691 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead",
            "value": 41946,
            "unit": "ns/op\t   26920 B/op\t     159 allocs/op",
            "extra": "29896 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead - ns/op",
            "value": 41946,
            "unit": "ns/op",
            "extra": "29896 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead - B/op",
            "value": 26920,
            "unit": "B/op",
            "extra": "29896 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead - allocs/op",
            "value": 159,
            "unit": "allocs/op",
            "extra": "29896 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead",
            "value": 18423,
            "unit": "ns/op\t   21848 B/op\t      77 allocs/op",
            "extra": "67534 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead - ns/op",
            "value": 18423,
            "unit": "ns/op",
            "extra": "67534 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead - B/op",
            "value": 21848,
            "unit": "B/op",
            "extra": "67534 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead - allocs/op",
            "value": 77,
            "unit": "allocs/op",
            "extra": "67534 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead",
            "value": 157846,
            "unit": "ns/op\t   56352 B/op\t     552 allocs/op",
            "extra": "7275 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead - ns/op",
            "value": 157846,
            "unit": "ns/op",
            "extra": "7275 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead - B/op",
            "value": 56352,
            "unit": "B/op",
            "extra": "7275 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead - allocs/op",
            "value": 552,
            "unit": "allocs/op",
            "extra": "7275 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2",
            "value": 170638,
            "unit": "ns/op\t   56352 B/op\t     552 allocs/op",
            "extra": "7078 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2 - ns/op",
            "value": 170638,
            "unit": "ns/op",
            "extra": "7078 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2 - B/op",
            "value": 56352,
            "unit": "B/op",
            "extra": "7078 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2 - allocs/op",
            "value": 552,
            "unit": "allocs/op",
            "extra": "7078 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3",
            "value": 59803,
            "unit": "ns/op\t   29752 B/op\t     264 allocs/op",
            "extra": "19802 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3 - ns/op",
            "value": 59803,
            "unit": "ns/op",
            "extra": "19802 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3 - B/op",
            "value": 29752,
            "unit": "B/op",
            "extra": "19802 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3 - allocs/op",
            "value": 264,
            "unit": "allocs/op",
            "extra": "19802 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile",
            "value": 9301,
            "unit": "ns/op\t   10760 B/op\t     131 allocs/op",
            "extra": "139465 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile - ns/op",
            "value": 9301,
            "unit": "ns/op",
            "extra": "139465 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile - B/op",
            "value": 10760,
            "unit": "B/op",
            "extra": "139465 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile - allocs/op",
            "value": 131,
            "unit": "allocs/op",
            "extra": "139465 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid",
            "value": 18.33,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "65873892 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid - ns/op",
            "value": 18.33,
            "unit": "ns/op",
            "extra": "65873892 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "65873892 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "65873892 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite",
            "value": 31640,
            "unit": "ns/op\t   34696 B/op\t     226 allocs/op",
            "extra": "35570 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite - ns/op",
            "value": 31640,
            "unit": "ns/op",
            "extra": "35570 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite - B/op",
            "value": 34696,
            "unit": "B/op",
            "extra": "35570 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite - allocs/op",
            "value": 226,
            "unit": "allocs/op",
            "extra": "35570 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite",
            "value": 138367,
            "unit": "ns/op\t   54680 B/op\t    2069 allocs/op",
            "extra": "8520 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite - ns/op",
            "value": 138367,
            "unit": "ns/op",
            "extra": "8520 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite - B/op",
            "value": 54680,
            "unit": "B/op",
            "extra": "8520 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite - allocs/op",
            "value": 2069,
            "unit": "allocs/op",
            "extra": "8520 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite",
            "value": 95564,
            "unit": "ns/op\t   61104 B/op\t     721 allocs/op",
            "extra": "12398 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite - ns/op",
            "value": 95564,
            "unit": "ns/op",
            "extra": "12398 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite - B/op",
            "value": 61104,
            "unit": "B/op",
            "extra": "12398 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite - allocs/op",
            "value": 721,
            "unit": "allocs/op",
            "extra": "12398 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDIATWrite",
            "value": 231380104,
            "unit": "1210428822610010322A094101Federal Reserve Bank   My Bank Name                   ",
            "extra": "101 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDIATWrite - 1210428822610010322A094101Federal",
            "value": 231380104,
            "unit": "1210428822610010322A094101Federal",
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
          }
        ]
      },
      {
        "commit": {
          "author": {
            "name": "Sash",
            "username": "SashaMIT",
            "email": "sash.t.mitchell@gmail.com"
          },
          "committer": {
            "name": "GitHub",
            "username": "web-flow",
            "email": "noreply@github.com"
          },
          "id": "7cfa7dd4fb0b53376267b33f8004a6f60192e5ca",
          "message": "Reject an ACH file creation time after 23:59. (#1873)\n\nValidate accepted 2900 because the hour pattern allowed 24 through 29, and a four character time skipped the parse error.",
          "timestamp": "2026-09-28T20:11:29Z",
          "url": "https://github.com/moov-io/ach/commit/7cfa7dd4fb0b53376267b33f8004a6f60192e5ca"
        },
        "date": 1790825338964,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkAlphaFieldShort",
            "value": 638.1,
            "unit": "ns/op\t      80 B/op\t       3 allocs/op",
            "extra": "2096504 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldShort - ns/op",
            "value": 638.1,
            "unit": "ns/op",
            "extra": "2096504 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldShort - B/op",
            "value": 80,
            "unit": "B/op",
            "extra": "2096504 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldShort - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "2096504 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong",
            "value": 69.32,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "17030271 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong - ns/op",
            "value": 69.32,
            "unit": "ns/op",
            "extra": "17030271 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "17030271 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "17030271 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort",
            "value": 37.5,
            "unit": "ns/op\t      16 B/op\t       2 allocs/op",
            "extra": "30049614 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort - ns/op",
            "value": 37.5,
            "unit": "ns/op",
            "extra": "30049614 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "30049614 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "30049614 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong",
            "value": 18.76,
            "unit": "ns/op\t       8 B/op\t       1 allocs/op",
            "extra": "63719856 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong - ns/op",
            "value": 18.76,
            "unit": "ns/op",
            "extra": "63719856 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong - B/op",
            "value": 8,
            "unit": "B/op",
            "extra": "63719856 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "63719856 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField",
            "value": 10.11,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "100000000 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField - ns/op",
            "value": 10.11,
            "unit": "ns/op",
            "extra": "100000000 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "100000000 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "100000000 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField",
            "value": 3.149,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "378066973 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField - ns/op",
            "value": 3.149,
            "unit": "ns/op",
            "extra": "378066973 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "378066973 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "378066973 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles",
            "value": 159227,
            "unit": "ns/op\t   54137 B/op\t     306 allocs/op",
            "extra": "7467 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles - ns/op",
            "value": 159227,
            "unit": "ns/op",
            "extra": "7467 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles - B/op",
            "value": 54137,
            "unit": "B/op",
            "extra": "7467 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles - allocs/op",
            "value": 306,
            "unit": "allocs/op",
            "extra": "7467 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts",
            "value": 159850,
            "unit": "ns/op\t   54153 B/op\t     306 allocs/op",
            "extra": "7256 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts - ns/op",
            "value": 159850,
            "unit": "ns/op",
            "extra": "7256 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts - B/op",
            "value": 54153,
            "unit": "B/op",
            "extra": "7256 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts - allocs/op",
            "value": 306,
            "unit": "allocs/op",
            "extra": "7256 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir",
            "value": 98359,
            "unit": "ns/op\t   54558 B/op\t     310 allocs/op",
            "extra": "12396 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir - ns/op",
            "value": 98359,
            "unit": "ns/op",
            "extra": "12396 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir - B/op",
            "value": 54558,
            "unit": "B/op",
            "extra": "12396 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir - allocs/op",
            "value": 310,
            "unit": "allocs/op",
            "extra": "12396 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts",
            "value": 115277,
            "unit": "ns/op\t   54586 B/op\t     310 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts - ns/op",
            "value": 115277,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts - B/op",
            "value": 54586,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts - allocs/op",
            "value": 310,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups",
            "value": 204152,
            "unit": "ns/op\t   59863 B/op\t     366 allocs/op",
            "extra": "6248 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups - ns/op",
            "value": 204152,
            "unit": "ns/op",
            "extra": "6248 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups - B/op",
            "value": 59863,
            "unit": "B/op",
            "extra": "6248 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups - allocs/op",
            "value": 366,
            "unit": "allocs/op",
            "extra": "6248 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups",
            "value": 199674,
            "unit": "ns/op\t   59920 B/op\t     366 allocs/op",
            "extra": "5724 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups - ns/op",
            "value": 199674,
            "unit": "ns/op",
            "extra": "5724 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups - B/op",
            "value": 59920,
            "unit": "B/op",
            "extra": "5724 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups - allocs/op",
            "value": 366,
            "unit": "allocs/op",
            "extra": "5724 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups",
            "value": 194817,
            "unit": "ns/op\t   59844 B/op\t     366 allocs/op",
            "extra": "5923 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups - ns/op",
            "value": 194817,
            "unit": "ns/op",
            "extra": "5923 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups - B/op",
            "value": 59844,
            "unit": "B/op",
            "extra": "5923 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups - allocs/op",
            "value": 366,
            "unit": "allocs/op",
            "extra": "5923 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups",
            "value": 186839,
            "unit": "ns/op\t   60069 B/op\t     367 allocs/op",
            "extra": "7610 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups - ns/op",
            "value": 186839,
            "unit": "ns/op",
            "extra": "7610 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups - B/op",
            "value": 60069,
            "unit": "B/op",
            "extra": "7610 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups - allocs/op",
            "value": 367,
            "unit": "allocs/op",
            "extra": "7610 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader",
            "value": 157246861,
            "unit": "ns/op\t42449621 B/op\t  203653 allocs/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader - ns/op",
            "value": 157246861,
            "unit": "ns/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader - B/op",
            "value": 42449621,
            "unit": "B/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader - allocs/op",
            "value": 203653,
            "unit": "allocs/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator",
            "value": 51205331,
            "unit": "ns/op\t42424668 B/op\t  203601 allocs/op",
            "extra": "21 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator - ns/op",
            "value": 51205331,
            "unit": "ns/op",
            "extra": "21 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator - B/op",
            "value": 42424668,
            "unit": "B/op",
            "extra": "21 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator - allocs/op",
            "value": 203601,
            "unit": "allocs/op",
            "extra": "21 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader",
            "value": 284787090,
            "unit": "ns/op\t62130952 B/op\t  905787 allocs/op",
            "extra": "4 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader - ns/op",
            "value": 284787090,
            "unit": "ns/op",
            "extra": "4 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader - B/op",
            "value": 62130952,
            "unit": "B/op",
            "extra": "4 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader - allocs/op",
            "value": 905787,
            "unit": "allocs/op",
            "extra": "4 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator",
            "value": 96075339,
            "unit": "ns/op\t60508801 B/op\t  705839 allocs/op",
            "extra": "12 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator - ns/op",
            "value": 96075339,
            "unit": "ns/op",
            "extra": "12 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator - B/op",
            "value": 60508801,
            "unit": "B/op",
            "extra": "12 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator - allocs/op",
            "value": 705839,
            "unit": "allocs/op",
            "extra": "12 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader",
            "value": 882615344,
            "unit": "ns/op\t215540432 B/op\t 1042859 allocs/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader - ns/op",
            "value": 882615344,
            "unit": "ns/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader - B/op",
            "value": 215540432,
            "unit": "B/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader - allocs/op",
            "value": 1042859,
            "unit": "allocs/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator",
            "value": 259215517,
            "unit": "ns/op\t215424746 B/op\t 1042800 allocs/op",
            "extra": "4 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator - ns/op",
            "value": 259215517,
            "unit": "ns/op",
            "extra": "4 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator - B/op",
            "value": 215424746,
            "unit": "B/op",
            "extra": "4 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator - allocs/op",
            "value": 1042800,
            "unit": "allocs/op",
            "extra": "4 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader",
            "value": 1414304913,
            "unit": "ns/op\t313545976 B/op\t 4568040 allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader - ns/op",
            "value": 1414304913,
            "unit": "ns/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader - B/op",
            "value": 313545976,
            "unit": "B/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader - allocs/op",
            "value": 4568040,
            "unit": "allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator",
            "value": 527170684,
            "unit": "ns/op\t305430068 B/op\t 3568071 allocs/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator - ns/op",
            "value": 527170684,
            "unit": "ns/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator - B/op",
            "value": 305430068,
            "unit": "B/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator - allocs/op",
            "value": 3568071,
            "unit": "allocs/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead",
            "value": 26210,
            "unit": "ns/op\t   23312 B/op\t     114 allocs/op",
            "extra": "45889 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead - ns/op",
            "value": 26210,
            "unit": "ns/op",
            "extra": "45889 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead - B/op",
            "value": 23312,
            "unit": "B/op",
            "extra": "45889 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead - allocs/op",
            "value": 114,
            "unit": "allocs/op",
            "extra": "45889 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead",
            "value": 45148,
            "unit": "ns/op\t   26920 B/op\t     159 allocs/op",
            "extra": "27021 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead - ns/op",
            "value": 45148,
            "unit": "ns/op",
            "extra": "27021 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead - B/op",
            "value": 26920,
            "unit": "B/op",
            "extra": "27021 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead - allocs/op",
            "value": 159,
            "unit": "allocs/op",
            "extra": "27021 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead",
            "value": 18721,
            "unit": "ns/op\t   21848 B/op\t      77 allocs/op",
            "extra": "63170 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead - ns/op",
            "value": 18721,
            "unit": "ns/op",
            "extra": "63170 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead - B/op",
            "value": 21848,
            "unit": "B/op",
            "extra": "63170 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead - allocs/op",
            "value": 77,
            "unit": "allocs/op",
            "extra": "63170 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead",
            "value": 162115,
            "unit": "ns/op\t   56352 B/op\t     552 allocs/op",
            "extra": "7402 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead - ns/op",
            "value": 162115,
            "unit": "ns/op",
            "extra": "7402 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead - B/op",
            "value": 56352,
            "unit": "B/op",
            "extra": "7402 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead - allocs/op",
            "value": 552,
            "unit": "allocs/op",
            "extra": "7402 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2",
            "value": 160741,
            "unit": "ns/op\t   56352 B/op\t     552 allocs/op",
            "extra": "6798 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2 - ns/op",
            "value": 160741,
            "unit": "ns/op",
            "extra": "6798 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2 - B/op",
            "value": 56352,
            "unit": "B/op",
            "extra": "6798 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2 - allocs/op",
            "value": 552,
            "unit": "allocs/op",
            "extra": "6798 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3",
            "value": 60063,
            "unit": "ns/op\t   29752 B/op\t     264 allocs/op",
            "extra": "20324 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3 - ns/op",
            "value": 60063,
            "unit": "ns/op",
            "extra": "20324 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3 - B/op",
            "value": 29752,
            "unit": "B/op",
            "extra": "20324 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3 - allocs/op",
            "value": 264,
            "unit": "allocs/op",
            "extra": "20324 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile",
            "value": 9378,
            "unit": "ns/op\t   10760 B/op\t     131 allocs/op",
            "extra": "138248 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile - ns/op",
            "value": 9378,
            "unit": "ns/op",
            "extra": "138248 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile - B/op",
            "value": 10760,
            "unit": "B/op",
            "extra": "138248 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile - allocs/op",
            "value": 131,
            "unit": "allocs/op",
            "extra": "138248 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid",
            "value": 18.45,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "65694849 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid - ns/op",
            "value": 18.45,
            "unit": "ns/op",
            "extra": "65694849 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "65694849 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "65694849 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite",
            "value": 31488,
            "unit": "ns/op\t   34696 B/op\t     226 allocs/op",
            "extra": "35583 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite - ns/op",
            "value": 31488,
            "unit": "ns/op",
            "extra": "35583 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite - B/op",
            "value": 34696,
            "unit": "B/op",
            "extra": "35583 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite - allocs/op",
            "value": 226,
            "unit": "allocs/op",
            "extra": "35583 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite",
            "value": 141246,
            "unit": "ns/op\t   54680 B/op\t    2069 allocs/op",
            "extra": "8599 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite - ns/op",
            "value": 141246,
            "unit": "ns/op",
            "extra": "8599 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite - B/op",
            "value": 54680,
            "unit": "B/op",
            "extra": "8599 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite - allocs/op",
            "value": 2069,
            "unit": "allocs/op",
            "extra": "8599 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite",
            "value": 95092,
            "unit": "ns/op\t   61104 B/op\t     721 allocs/op",
            "extra": "12441 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite - ns/op",
            "value": 95092,
            "unit": "ns/op",
            "extra": "12441 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite - B/op",
            "value": 61104,
            "unit": "B/op",
            "extra": "12441 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite - allocs/op",
            "value": 721,
            "unit": "allocs/op",
            "extra": "12441 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDIATWrite",
            "value": 231380104,
            "unit": "1210428822610020328A094101Federal Reserve Bank   My Bank Name                   ",
            "extra": "101 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDIATWrite - 1210428822610020328A094101Federal",
            "value": 231380104,
            "unit": "1210428822610020328A094101Federal",
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
          }
        ]
      },
      {
        "commit": {
          "author": {
            "name": "Sash",
            "username": "SashaMIT",
            "email": "sash.t.mitchell@gmail.com"
          },
          "committer": {
            "name": "GitHub",
            "username": "web-flow",
            "email": "noreply@github.com"
          },
          "id": "7cfa7dd4fb0b53376267b33f8004a6f60192e5ca",
          "message": "Reject an ACH file creation time after 23:59. (#1873)\n\nValidate accepted 2900 because the hour pattern allowed 24 through 29, and a four character time skipped the parse error.",
          "timestamp": "2026-09-28T20:11:29Z",
          "url": "https://github.com/moov-io/ach/commit/7cfa7dd4fb0b53376267b33f8004a6f60192e5ca"
        },
        "date": 1790911761886,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkAlphaFieldShort",
            "value": 954.5,
            "unit": "ns/op\t      80 B/op\t       3 allocs/op",
            "extra": "1256106 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldShort - ns/op",
            "value": 954.5,
            "unit": "ns/op",
            "extra": "1256106 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldShort - B/op",
            "value": 80,
            "unit": "B/op",
            "extra": "1256106 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldShort - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "1256106 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong",
            "value": 98.41,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "12094274 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong - ns/op",
            "value": 98.41,
            "unit": "ns/op",
            "extra": "12094274 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "12094274 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "12094274 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort",
            "value": 56.7,
            "unit": "ns/op\t      16 B/op\t       2 allocs/op",
            "extra": "20687918 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort - ns/op",
            "value": 56.7,
            "unit": "ns/op",
            "extra": "20687918 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "20687918 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "20687918 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong",
            "value": 25.63,
            "unit": "ns/op\t       8 B/op\t       1 allocs/op",
            "extra": "44082306 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong - ns/op",
            "value": 25.63,
            "unit": "ns/op",
            "extra": "44082306 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong - B/op",
            "value": 8,
            "unit": "B/op",
            "extra": "44082306 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "44082306 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField",
            "value": 14.66,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "83538530 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField - ns/op",
            "value": 14.66,
            "unit": "ns/op",
            "extra": "83538530 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "83538530 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "83538530 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField",
            "value": 5.631,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "213425163 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField - ns/op",
            "value": 5.631,
            "unit": "ns/op",
            "extra": "213425163 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "213425163 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "213425163 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles",
            "value": 247451,
            "unit": "ns/op\t   54147 B/op\t     306 allocs/op",
            "extra": "5103 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles - ns/op",
            "value": 247451,
            "unit": "ns/op",
            "extra": "5103 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles - B/op",
            "value": 54147,
            "unit": "B/op",
            "extra": "5103 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles - allocs/op",
            "value": 306,
            "unit": "allocs/op",
            "extra": "5103 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts",
            "value": 248094,
            "unit": "ns/op\t   54158 B/op\t     306 allocs/op",
            "extra": "4665 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts - ns/op",
            "value": 248094,
            "unit": "ns/op",
            "extra": "4665 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts - B/op",
            "value": 54158,
            "unit": "B/op",
            "extra": "4665 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts - allocs/op",
            "value": 306,
            "unit": "allocs/op",
            "extra": "4665 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir",
            "value": 148902,
            "unit": "ns/op\t   54558 B/op\t     310 allocs/op",
            "extra": "7111 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir - ns/op",
            "value": 148902,
            "unit": "ns/op",
            "extra": "7111 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir - B/op",
            "value": 54558,
            "unit": "B/op",
            "extra": "7111 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir - allocs/op",
            "value": 310,
            "unit": "allocs/op",
            "extra": "7111 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts",
            "value": 148650,
            "unit": "ns/op\t   54583 B/op\t     310 allocs/op",
            "extra": "7825 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts - ns/op",
            "value": 148650,
            "unit": "ns/op",
            "extra": "7825 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts - B/op",
            "value": 54583,
            "unit": "B/op",
            "extra": "7825 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts - allocs/op",
            "value": 310,
            "unit": "allocs/op",
            "extra": "7825 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups",
            "value": 317162,
            "unit": "ns/op\t   59896 B/op\t     366 allocs/op",
            "extra": "3967 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups - ns/op",
            "value": 317162,
            "unit": "ns/op",
            "extra": "3967 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups - B/op",
            "value": 59896,
            "unit": "B/op",
            "extra": "3967 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups - allocs/op",
            "value": 366,
            "unit": "allocs/op",
            "extra": "3967 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups",
            "value": 309843,
            "unit": "ns/op\t   59846 B/op\t     366 allocs/op",
            "extra": "3981 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups - ns/op",
            "value": 309843,
            "unit": "ns/op",
            "extra": "3981 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups - B/op",
            "value": 59846,
            "unit": "B/op",
            "extra": "3981 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups - allocs/op",
            "value": 366,
            "unit": "allocs/op",
            "extra": "3981 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups",
            "value": 307487,
            "unit": "ns/op\t   59924 B/op\t     366 allocs/op",
            "extra": "4035 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups - ns/op",
            "value": 307487,
            "unit": "ns/op",
            "extra": "4035 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups - B/op",
            "value": 59924,
            "unit": "B/op",
            "extra": "4035 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups - allocs/op",
            "value": 366,
            "unit": "allocs/op",
            "extra": "4035 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups",
            "value": 289072,
            "unit": "ns/op\t   59561 B/op\t     367 allocs/op",
            "extra": "4884 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups - ns/op",
            "value": 289072,
            "unit": "ns/op",
            "extra": "4884 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups - B/op",
            "value": 59561,
            "unit": "B/op",
            "extra": "4884 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups - allocs/op",
            "value": 367,
            "unit": "allocs/op",
            "extra": "4884 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader",
            "value": 245033370,
            "unit": "ns/op\t42450427 B/op\t  203656 allocs/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader - ns/op",
            "value": 245033370,
            "unit": "ns/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader - B/op",
            "value": 42450427,
            "unit": "B/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader - allocs/op",
            "value": 203656,
            "unit": "allocs/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator",
            "value": 80795927,
            "unit": "ns/op\t42424716 B/op\t  203601 allocs/op",
            "extra": "14 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator - ns/op",
            "value": 80795927,
            "unit": "ns/op",
            "extra": "14 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator - B/op",
            "value": 42424716,
            "unit": "B/op",
            "extra": "14 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator - allocs/op",
            "value": 203601,
            "unit": "allocs/op",
            "extra": "14 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader",
            "value": 459930601,
            "unit": "ns/op\t62131877 B/op\t  905794 allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader - ns/op",
            "value": 459930601,
            "unit": "ns/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader - B/op",
            "value": 62131877,
            "unit": "B/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader - allocs/op",
            "value": 905794,
            "unit": "allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator",
            "value": 150227818,
            "unit": "ns/op\t60508800 B/op\t  705839 allocs/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator - ns/op",
            "value": 150227818,
            "unit": "ns/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator - B/op",
            "value": 60508800,
            "unit": "B/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator - allocs/op",
            "value": 705839,
            "unit": "allocs/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader",
            "value": 1299135808,
            "unit": "ns/op\t215542072 B/op\t 1042871 allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader - ns/op",
            "value": 1299135808,
            "unit": "ns/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader - B/op",
            "value": 215542072,
            "unit": "B/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader - allocs/op",
            "value": 1042871,
            "unit": "allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator",
            "value": 391109794,
            "unit": "ns/op\t215424853 B/op\t 1042801 allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator - ns/op",
            "value": 391109794,
            "unit": "ns/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator - B/op",
            "value": 215424853,
            "unit": "B/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator - allocs/op",
            "value": 1042801,
            "unit": "allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader",
            "value": 2214861733,
            "unit": "ns/op\t313545976 B/op\t 4568040 allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader - ns/op",
            "value": 2214861733,
            "unit": "ns/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader - B/op",
            "value": 313545976,
            "unit": "B/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader - allocs/op",
            "value": 4568040,
            "unit": "allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator",
            "value": 758580215,
            "unit": "ns/op\t305430308 B/op\t 3568074 allocs/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator - ns/op",
            "value": 758580215,
            "unit": "ns/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator - B/op",
            "value": 305430308,
            "unit": "B/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator - allocs/op",
            "value": 3568074,
            "unit": "allocs/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead",
            "value": 42429,
            "unit": "ns/op\t   23312 B/op\t     114 allocs/op",
            "extra": "28467 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead - ns/op",
            "value": 42429,
            "unit": "ns/op",
            "extra": "28467 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead - B/op",
            "value": 23312,
            "unit": "B/op",
            "extra": "28467 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead - allocs/op",
            "value": 114,
            "unit": "allocs/op",
            "extra": "28467 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead",
            "value": 72341,
            "unit": "ns/op\t   26920 B/op\t     159 allocs/op",
            "extra": "16051 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead - ns/op",
            "value": 72341,
            "unit": "ns/op",
            "extra": "16051 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead - B/op",
            "value": 26920,
            "unit": "B/op",
            "extra": "16051 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead - allocs/op",
            "value": 159,
            "unit": "allocs/op",
            "extra": "16051 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead",
            "value": 31674,
            "unit": "ns/op\t   21848 B/op\t      77 allocs/op",
            "extra": "38727 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead - ns/op",
            "value": 31674,
            "unit": "ns/op",
            "extra": "38727 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead - B/op",
            "value": 21848,
            "unit": "B/op",
            "extra": "38727 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead - allocs/op",
            "value": 77,
            "unit": "allocs/op",
            "extra": "38727 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead",
            "value": 271566,
            "unit": "ns/op\t   56353 B/op\t     552 allocs/op",
            "extra": "4380 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead - ns/op",
            "value": 271566,
            "unit": "ns/op",
            "extra": "4380 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead - B/op",
            "value": 56353,
            "unit": "B/op",
            "extra": "4380 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead - allocs/op",
            "value": 552,
            "unit": "allocs/op",
            "extra": "4380 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2",
            "value": 271174,
            "unit": "ns/op\t   56353 B/op\t     552 allocs/op",
            "extra": "4528 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2 - ns/op",
            "value": 271174,
            "unit": "ns/op",
            "extra": "4528 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2 - B/op",
            "value": 56353,
            "unit": "B/op",
            "extra": "4528 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2 - allocs/op",
            "value": 552,
            "unit": "allocs/op",
            "extra": "4528 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3",
            "value": 104284,
            "unit": "ns/op\t   29752 B/op\t     264 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3 - ns/op",
            "value": 104284,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3 - B/op",
            "value": 29752,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3 - allocs/op",
            "value": 264,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile",
            "value": 14248,
            "unit": "ns/op\t   10760 B/op\t     131 allocs/op",
            "extra": "92803 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile - ns/op",
            "value": 14248,
            "unit": "ns/op",
            "extra": "92803 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile - B/op",
            "value": 10760,
            "unit": "B/op",
            "extra": "92803 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile - allocs/op",
            "value": 131,
            "unit": "allocs/op",
            "extra": "92803 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid",
            "value": 28.13,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "44746899 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid - ns/op",
            "value": 28.13,
            "unit": "ns/op",
            "extra": "44746899 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "44746899 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "44746899 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite",
            "value": 51057,
            "unit": "ns/op\t   34696 B/op\t     226 allocs/op",
            "extra": "22292 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite - ns/op",
            "value": 51057,
            "unit": "ns/op",
            "extra": "22292 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite - B/op",
            "value": 34696,
            "unit": "B/op",
            "extra": "22292 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite - allocs/op",
            "value": 226,
            "unit": "allocs/op",
            "extra": "22292 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite",
            "value": 214354,
            "unit": "ns/op\t   54680 B/op\t    2069 allocs/op",
            "extra": "5852 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite - ns/op",
            "value": 214354,
            "unit": "ns/op",
            "extra": "5852 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite - B/op",
            "value": 54680,
            "unit": "B/op",
            "extra": "5852 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite - allocs/op",
            "value": 2069,
            "unit": "allocs/op",
            "extra": "5852 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite",
            "value": 152747,
            "unit": "ns/op\t   61104 B/op\t     721 allocs/op",
            "extra": "7611 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite - ns/op",
            "value": 152747,
            "unit": "ns/op",
            "extra": "7611 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite - B/op",
            "value": 61104,
            "unit": "B/op",
            "extra": "7611 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite - allocs/op",
            "value": 721,
            "unit": "allocs/op",
            "extra": "7611 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDIATWrite",
            "value": 231380104,
            "unit": "1210428822610030329A094101Federal Reserve Bank   My Bank Name                   ",
            "extra": "101 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDIATWrite - 1210428822610030329A094101Federal",
            "value": 231380104,
            "unit": "1210428822610030329A094101Federal",
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
          }
        ]
      },
      {
        "commit": {
          "author": {
            "name": "Sash",
            "username": "SashaMIT",
            "email": "sash.t.mitchell@gmail.com"
          },
          "committer": {
            "name": "GitHub",
            "username": "web-flow",
            "email": "noreply@github.com"
          },
          "id": "7cfa7dd4fb0b53376267b33f8004a6f60192e5ca",
          "message": "Reject an ACH file creation time after 23:59. (#1873)\n\nValidate accepted 2900 because the hour pattern allowed 24 through 29, and a four character time skipped the parse error.",
          "timestamp": "2026-09-28T20:11:29Z",
          "url": "https://github.com/moov-io/ach/commit/7cfa7dd4fb0b53376267b33f8004a6f60192e5ca"
        },
        "date": 1790997197163,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkAlphaFieldShort",
            "value": 875.9,
            "unit": "ns/op\t      80 B/op\t       3 allocs/op",
            "extra": "1378698 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldShort - ns/op",
            "value": 875.9,
            "unit": "ns/op",
            "extra": "1378698 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldShort - B/op",
            "value": 80,
            "unit": "B/op",
            "extra": "1378698 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldShort - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "1378698 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong",
            "value": 100.5,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "11866843 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong - ns/op",
            "value": 100.5,
            "unit": "ns/op",
            "extra": "11866843 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "11866843 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "11866843 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort",
            "value": 55.09,
            "unit": "ns/op\t      16 B/op\t       2 allocs/op",
            "extra": "21268132 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort - ns/op",
            "value": 55.09,
            "unit": "ns/op",
            "extra": "21268132 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "21268132 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "21268132 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong",
            "value": 25.25,
            "unit": "ns/op\t       8 B/op\t       1 allocs/op",
            "extra": "42475072 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong - ns/op",
            "value": 25.25,
            "unit": "ns/op",
            "extra": "42475072 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong - B/op",
            "value": 8,
            "unit": "B/op",
            "extra": "42475072 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "42475072 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField",
            "value": 14.55,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "88606110 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField - ns/op",
            "value": 14.55,
            "unit": "ns/op",
            "extra": "88606110 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "88606110 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "88606110 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField",
            "value": 5.732,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "210197102 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField - ns/op",
            "value": 5.732,
            "unit": "ns/op",
            "extra": "210197102 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "210197102 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "210197102 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles",
            "value": 243073,
            "unit": "ns/op\t   54140 B/op\t     306 allocs/op",
            "extra": "5413 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles - ns/op",
            "value": 243073,
            "unit": "ns/op",
            "extra": "5413 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles - B/op",
            "value": 54140,
            "unit": "B/op",
            "extra": "5413 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles - allocs/op",
            "value": 306,
            "unit": "allocs/op",
            "extra": "5413 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts",
            "value": 255189,
            "unit": "ns/op\t   54162 B/op\t     306 allocs/op",
            "extra": "4988 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts - ns/op",
            "value": 255189,
            "unit": "ns/op",
            "extra": "4988 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts - B/op",
            "value": 54162,
            "unit": "B/op",
            "extra": "4988 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts - allocs/op",
            "value": 306,
            "unit": "allocs/op",
            "extra": "4988 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir",
            "value": 140644,
            "unit": "ns/op\t   54546 B/op\t     310 allocs/op",
            "extra": "7899 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir - ns/op",
            "value": 140644,
            "unit": "ns/op",
            "extra": "7899 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir - B/op",
            "value": 54546,
            "unit": "B/op",
            "extra": "7899 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir - allocs/op",
            "value": 310,
            "unit": "allocs/op",
            "extra": "7899 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts",
            "value": 140927,
            "unit": "ns/op\t   54590 B/op\t     310 allocs/op",
            "extra": "7614 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts - ns/op",
            "value": 140927,
            "unit": "ns/op",
            "extra": "7614 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts - B/op",
            "value": 54590,
            "unit": "B/op",
            "extra": "7614 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts - allocs/op",
            "value": 310,
            "unit": "allocs/op",
            "extra": "7614 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups",
            "value": 304044,
            "unit": "ns/op\t   59855 B/op\t     366 allocs/op",
            "extra": "4095 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups - ns/op",
            "value": 304044,
            "unit": "ns/op",
            "extra": "4095 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups - B/op",
            "value": 59855,
            "unit": "B/op",
            "extra": "4095 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups - allocs/op",
            "value": 366,
            "unit": "allocs/op",
            "extra": "4095 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups",
            "value": 306900,
            "unit": "ns/op\t   59871 B/op\t     366 allocs/op",
            "extra": "4399 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups - ns/op",
            "value": 306900,
            "unit": "ns/op",
            "extra": "4399 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups - B/op",
            "value": 59871,
            "unit": "B/op",
            "extra": "4399 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups - allocs/op",
            "value": 366,
            "unit": "allocs/op",
            "extra": "4399 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups",
            "value": 294103,
            "unit": "ns/op\t   59835 B/op\t     366 allocs/op",
            "extra": "4284 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups - ns/op",
            "value": 294103,
            "unit": "ns/op",
            "extra": "4284 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups - B/op",
            "value": 59835,
            "unit": "B/op",
            "extra": "4284 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups - allocs/op",
            "value": 366,
            "unit": "allocs/op",
            "extra": "4284 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups",
            "value": 288804,
            "unit": "ns/op\t   59500 B/op\t     367 allocs/op",
            "extra": "5029 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups - ns/op",
            "value": 288804,
            "unit": "ns/op",
            "extra": "5029 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups - B/op",
            "value": 59500,
            "unit": "B/op",
            "extra": "5029 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups - allocs/op",
            "value": 367,
            "unit": "allocs/op",
            "extra": "5029 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader",
            "value": 239963289,
            "unit": "ns/op\t42450196 B/op\t  203653 allocs/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader - ns/op",
            "value": 239963289,
            "unit": "ns/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader - B/op",
            "value": 42450196,
            "unit": "B/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader - allocs/op",
            "value": 203653,
            "unit": "allocs/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator",
            "value": 71785509,
            "unit": "ns/op\t42424772 B/op\t  203601 allocs/op",
            "extra": "15 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator - ns/op",
            "value": 71785509,
            "unit": "ns/op",
            "extra": "15 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator - B/op",
            "value": 42424772,
            "unit": "B/op",
            "extra": "15 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator - allocs/op",
            "value": 203601,
            "unit": "allocs/op",
            "extra": "15 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader",
            "value": 441316903,
            "unit": "ns/op\t62131818 B/op\t  905794 allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader - ns/op",
            "value": 441316903,
            "unit": "ns/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader - B/op",
            "value": 62131818,
            "unit": "B/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader - allocs/op",
            "value": 905794,
            "unit": "allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator",
            "value": 137970432,
            "unit": "ns/op\t60508782 B/op\t  705838 allocs/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator - ns/op",
            "value": 137970432,
            "unit": "ns/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator - B/op",
            "value": 60508782,
            "unit": "B/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator - allocs/op",
            "value": 705838,
            "unit": "allocs/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader",
            "value": 1320685112,
            "unit": "ns/op\t215541928 B/op\t 1042869 allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader - ns/op",
            "value": 1320685112,
            "unit": "ns/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader - B/op",
            "value": 215541928,
            "unit": "B/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader - allocs/op",
            "value": 1042869,
            "unit": "allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator",
            "value": 361209029,
            "unit": "ns/op\t215424746 B/op\t 1042799 allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator - ns/op",
            "value": 361209029,
            "unit": "ns/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator - B/op",
            "value": 215424746,
            "unit": "B/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator - allocs/op",
            "value": 1042799,
            "unit": "allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader",
            "value": 2179582090,
            "unit": "ns/op\t313546056 B/op\t 4568041 allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader - ns/op",
            "value": 2179582090,
            "unit": "ns/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader - B/op",
            "value": 313546056,
            "unit": "B/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader - allocs/op",
            "value": 4568041,
            "unit": "allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator",
            "value": 695614107,
            "unit": "ns/op\t305429908 B/op\t 3568069 allocs/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator - ns/op",
            "value": 695614107,
            "unit": "ns/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator - B/op",
            "value": 305429908,
            "unit": "B/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator - allocs/op",
            "value": 3568069,
            "unit": "allocs/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead",
            "value": 42394,
            "unit": "ns/op\t   23312 B/op\t     114 allocs/op",
            "extra": "28203 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead - ns/op",
            "value": 42394,
            "unit": "ns/op",
            "extra": "28203 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead - B/op",
            "value": 23312,
            "unit": "B/op",
            "extra": "28203 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead - allocs/op",
            "value": 114,
            "unit": "allocs/op",
            "extra": "28203 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead",
            "value": 70736,
            "unit": "ns/op\t   26920 B/op\t     159 allocs/op",
            "extra": "16393 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead - ns/op",
            "value": 70736,
            "unit": "ns/op",
            "extra": "16393 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead - B/op",
            "value": 26920,
            "unit": "B/op",
            "extra": "16393 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead - allocs/op",
            "value": 159,
            "unit": "allocs/op",
            "extra": "16393 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead",
            "value": 32221,
            "unit": "ns/op\t   21848 B/op\t      77 allocs/op",
            "extra": "37911 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead - ns/op",
            "value": 32221,
            "unit": "ns/op",
            "extra": "37911 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead - B/op",
            "value": 21848,
            "unit": "B/op",
            "extra": "37911 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead - allocs/op",
            "value": 77,
            "unit": "allocs/op",
            "extra": "37911 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead",
            "value": 261292,
            "unit": "ns/op\t   56353 B/op\t     552 allocs/op",
            "extra": "4671 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead - ns/op",
            "value": 261292,
            "unit": "ns/op",
            "extra": "4671 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead - B/op",
            "value": 56353,
            "unit": "B/op",
            "extra": "4671 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead - allocs/op",
            "value": 552,
            "unit": "allocs/op",
            "extra": "4671 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2",
            "value": 261738,
            "unit": "ns/op\t   56353 B/op\t     552 allocs/op",
            "extra": "4675 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2 - ns/op",
            "value": 261738,
            "unit": "ns/op",
            "extra": "4675 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2 - B/op",
            "value": 56353,
            "unit": "B/op",
            "extra": "4675 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2 - allocs/op",
            "value": 552,
            "unit": "allocs/op",
            "extra": "4675 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3",
            "value": 100476,
            "unit": "ns/op\t   29752 B/op\t     264 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3 - ns/op",
            "value": 100476,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3 - B/op",
            "value": 29752,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3 - allocs/op",
            "value": 264,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile",
            "value": 13380,
            "unit": "ns/op\t   10760 B/op\t     131 allocs/op",
            "extra": "95070 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile - ns/op",
            "value": 13380,
            "unit": "ns/op",
            "extra": "95070 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile - B/op",
            "value": 10760,
            "unit": "B/op",
            "extra": "95070 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile - allocs/op",
            "value": 131,
            "unit": "allocs/op",
            "extra": "95070 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid",
            "value": 25.38,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "47327198 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid - ns/op",
            "value": 25.38,
            "unit": "ns/op",
            "extra": "47327198 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "47327198 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "47327198 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite",
            "value": 48513,
            "unit": "ns/op\t   34696 B/op\t     226 allocs/op",
            "extra": "23325 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite - ns/op",
            "value": 48513,
            "unit": "ns/op",
            "extra": "23325 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite - B/op",
            "value": 34696,
            "unit": "B/op",
            "extra": "23325 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite - allocs/op",
            "value": 226,
            "unit": "allocs/op",
            "extra": "23325 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite",
            "value": 187035,
            "unit": "ns/op\t   54680 B/op\t    2069 allocs/op",
            "extra": "6675 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite - ns/op",
            "value": 187035,
            "unit": "ns/op",
            "extra": "6675 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite - B/op",
            "value": 54680,
            "unit": "B/op",
            "extra": "6675 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite - allocs/op",
            "value": 2069,
            "unit": "allocs/op",
            "extra": "6675 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite",
            "value": 147992,
            "unit": "ns/op\t   61104 B/op\t     721 allocs/op",
            "extra": "7827 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite - ns/op",
            "value": 147992,
            "unit": "ns/op",
            "extra": "7827 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite - B/op",
            "value": 61104,
            "unit": "B/op",
            "extra": "7827 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite - allocs/op",
            "value": 721,
            "unit": "allocs/op",
            "extra": "7827 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDIATWrite",
            "value": 231380104,
            "unit": "1210428822610040313A094101Federal Reserve Bank   My Bank Name                   ",
            "extra": "101 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDIATWrite - 1210428822610040313A094101Federal",
            "value": 231380104,
            "unit": "1210428822610040313A094101Federal",
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
          }
        ]
      },
      {
        "commit": {
          "author": {
            "name": "Sash",
            "username": "SashaMIT",
            "email": "sash.t.mitchell@gmail.com"
          },
          "committer": {
            "name": "GitHub",
            "username": "web-flow",
            "email": "noreply@github.com"
          },
          "id": "7cfa7dd4fb0b53376267b33f8004a6f60192e5ca",
          "message": "Reject an ACH file creation time after 23:59. (#1873)\n\nValidate accepted 2900 because the hour pattern allowed 24 through 29, and a four character time skipped the parse error.",
          "timestamp": "2026-09-28T20:11:29Z",
          "url": "https://github.com/moov-io/ach/commit/7cfa7dd4fb0b53376267b33f8004a6f60192e5ca"
        },
        "date": 1791085283900,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkAlphaFieldShort",
            "value": 969.2,
            "unit": "ns/op\t      80 B/op\t       3 allocs/op",
            "extra": "1238544 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldShort - ns/op",
            "value": 969.2,
            "unit": "ns/op",
            "extra": "1238544 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldShort - B/op",
            "value": 80,
            "unit": "B/op",
            "extra": "1238544 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldShort - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "1238544 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong",
            "value": 99.24,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "11868074 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong - ns/op",
            "value": 99.24,
            "unit": "ns/op",
            "extra": "11868074 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "11868074 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "11868074 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort",
            "value": 57.04,
            "unit": "ns/op\t      16 B/op\t       2 allocs/op",
            "extra": "20330137 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort - ns/op",
            "value": 57.04,
            "unit": "ns/op",
            "extra": "20330137 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "20330137 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "20330137 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong",
            "value": 25.53,
            "unit": "ns/op\t       8 B/op\t       1 allocs/op",
            "extra": "44999827 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong - ns/op",
            "value": 25.53,
            "unit": "ns/op",
            "extra": "44999827 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong - B/op",
            "value": 8,
            "unit": "B/op",
            "extra": "44999827 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "44999827 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField",
            "value": 14.72,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "83465480 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField - ns/op",
            "value": 14.72,
            "unit": "ns/op",
            "extra": "83465480 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "83465480 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "83465480 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField",
            "value": 5.622,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "213735164 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField - ns/op",
            "value": 5.622,
            "unit": "ns/op",
            "extra": "213735164 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "213735164 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "213735164 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles",
            "value": 249031,
            "unit": "ns/op\t   54143 B/op\t     306 allocs/op",
            "extra": "4965 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles - ns/op",
            "value": 249031,
            "unit": "ns/op",
            "extra": "4965 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles - B/op",
            "value": 54143,
            "unit": "B/op",
            "extra": "4965 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles - allocs/op",
            "value": 306,
            "unit": "allocs/op",
            "extra": "4965 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts",
            "value": 256765,
            "unit": "ns/op\t   54147 B/op\t     306 allocs/op",
            "extra": "4902 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts - ns/op",
            "value": 256765,
            "unit": "ns/op",
            "extra": "4902 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts - B/op",
            "value": 54147,
            "unit": "B/op",
            "extra": "4902 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts - allocs/op",
            "value": 306,
            "unit": "allocs/op",
            "extra": "4902 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir",
            "value": 143534,
            "unit": "ns/op\t   54552 B/op\t     310 allocs/op",
            "extra": "7536 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir - ns/op",
            "value": 143534,
            "unit": "ns/op",
            "extra": "7536 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir - B/op",
            "value": 54552,
            "unit": "B/op",
            "extra": "7536 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir - allocs/op",
            "value": 310,
            "unit": "allocs/op",
            "extra": "7536 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts",
            "value": 149871,
            "unit": "ns/op\t   54584 B/op\t     310 allocs/op",
            "extra": "7663 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts - ns/op",
            "value": 149871,
            "unit": "ns/op",
            "extra": "7663 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts - B/op",
            "value": 54584,
            "unit": "B/op",
            "extra": "7663 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts - allocs/op",
            "value": 310,
            "unit": "allocs/op",
            "extra": "7663 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups",
            "value": 316380,
            "unit": "ns/op\t   59842 B/op\t     366 allocs/op",
            "extra": "3808 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups - ns/op",
            "value": 316380,
            "unit": "ns/op",
            "extra": "3808 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups - B/op",
            "value": 59842,
            "unit": "B/op",
            "extra": "3808 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups - allocs/op",
            "value": 366,
            "unit": "allocs/op",
            "extra": "3808 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups",
            "value": 313540,
            "unit": "ns/op\t   59899 B/op\t     366 allocs/op",
            "extra": "3902 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups - ns/op",
            "value": 313540,
            "unit": "ns/op",
            "extra": "3902 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups - B/op",
            "value": 59899,
            "unit": "B/op",
            "extra": "3902 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups - allocs/op",
            "value": 366,
            "unit": "allocs/op",
            "extra": "3902 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups",
            "value": 302643,
            "unit": "ns/op\t   59873 B/op\t     366 allocs/op",
            "extra": "3944 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups - ns/op",
            "value": 302643,
            "unit": "ns/op",
            "extra": "3944 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups - B/op",
            "value": 59873,
            "unit": "B/op",
            "extra": "3944 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups - allocs/op",
            "value": 366,
            "unit": "allocs/op",
            "extra": "3944 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups",
            "value": 294116,
            "unit": "ns/op\t   59599 B/op\t     367 allocs/op",
            "extra": "4748 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups - ns/op",
            "value": 294116,
            "unit": "ns/op",
            "extra": "4748 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups - B/op",
            "value": 59599,
            "unit": "B/op",
            "extra": "4748 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups - allocs/op",
            "value": 367,
            "unit": "allocs/op",
            "extra": "4748 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader",
            "value": 250662159,
            "unit": "ns/op\t42450842 B/op\t  203656 allocs/op",
            "extra": "4 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader - ns/op",
            "value": 250662159,
            "unit": "ns/op",
            "extra": "4 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader - B/op",
            "value": 42450842,
            "unit": "B/op",
            "extra": "4 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader - allocs/op",
            "value": 203656,
            "unit": "allocs/op",
            "extra": "4 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator",
            "value": 80183341,
            "unit": "ns/op\t42424767 B/op\t  203601 allocs/op",
            "extra": "14 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator - ns/op",
            "value": 80183341,
            "unit": "ns/op",
            "extra": "14 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator - B/op",
            "value": 42424767,
            "unit": "B/op",
            "extra": "14 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator - allocs/op",
            "value": 203601,
            "unit": "allocs/op",
            "extra": "14 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader",
            "value": 455167892,
            "unit": "ns/op\t62131781 B/op\t  905793 allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader - ns/op",
            "value": 455167892,
            "unit": "ns/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader - B/op",
            "value": 62131781,
            "unit": "B/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader - allocs/op",
            "value": 905793,
            "unit": "allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator",
            "value": 152446453,
            "unit": "ns/op\t60508834 B/op\t  705839 allocs/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator - ns/op",
            "value": 152446453,
            "unit": "ns/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator - B/op",
            "value": 60508834,
            "unit": "B/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator - allocs/op",
            "value": 705839,
            "unit": "allocs/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader",
            "value": 1336484233,
            "unit": "ns/op\t215542072 B/op\t 1042871 allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader - ns/op",
            "value": 1336484233,
            "unit": "ns/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader - B/op",
            "value": 215542072,
            "unit": "B/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader - allocs/op",
            "value": 1042871,
            "unit": "allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator",
            "value": 397989903,
            "unit": "ns/op\t215424805 B/op\t 1042800 allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator - ns/op",
            "value": 397989903,
            "unit": "ns/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator - B/op",
            "value": 215424805,
            "unit": "B/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator - allocs/op",
            "value": 1042800,
            "unit": "allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader",
            "value": 2216259897,
            "unit": "ns/op\t313546200 B/op\t 4568043 allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader - ns/op",
            "value": 2216259897,
            "unit": "ns/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader - B/op",
            "value": 313546200,
            "unit": "B/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader - allocs/op",
            "value": 4568043,
            "unit": "allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator",
            "value": 759966392,
            "unit": "ns/op\t305430300 B/op\t 3568074 allocs/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator - ns/op",
            "value": 759966392,
            "unit": "ns/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator - B/op",
            "value": 305430300,
            "unit": "B/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator - allocs/op",
            "value": 3568074,
            "unit": "allocs/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead",
            "value": 42394,
            "unit": "ns/op\t   23312 B/op\t     114 allocs/op",
            "extra": "28348 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead - ns/op",
            "value": 42394,
            "unit": "ns/op",
            "extra": "28348 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead - B/op",
            "value": 23312,
            "unit": "B/op",
            "extra": "28348 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead - allocs/op",
            "value": 114,
            "unit": "allocs/op",
            "extra": "28348 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead",
            "value": 74046,
            "unit": "ns/op\t   26920 B/op\t     159 allocs/op",
            "extra": "15555 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead - ns/op",
            "value": 74046,
            "unit": "ns/op",
            "extra": "15555 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead - B/op",
            "value": 26920,
            "unit": "B/op",
            "extra": "15555 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead - allocs/op",
            "value": 159,
            "unit": "allocs/op",
            "extra": "15555 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead",
            "value": 32122,
            "unit": "ns/op\t   21848 B/op\t      77 allocs/op",
            "extra": "38743 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead - ns/op",
            "value": 32122,
            "unit": "ns/op",
            "extra": "38743 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead - B/op",
            "value": 21848,
            "unit": "B/op",
            "extra": "38743 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead - allocs/op",
            "value": 77,
            "unit": "allocs/op",
            "extra": "38743 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead",
            "value": 280191,
            "unit": "ns/op\t   56353 B/op\t     552 allocs/op",
            "extra": "4542 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead - ns/op",
            "value": 280191,
            "unit": "ns/op",
            "extra": "4542 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead - B/op",
            "value": 56353,
            "unit": "B/op",
            "extra": "4542 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead - allocs/op",
            "value": 552,
            "unit": "allocs/op",
            "extra": "4542 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2",
            "value": 275957,
            "unit": "ns/op\t   56353 B/op\t     552 allocs/op",
            "extra": "4374 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2 - ns/op",
            "value": 275957,
            "unit": "ns/op",
            "extra": "4374 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2 - B/op",
            "value": 56353,
            "unit": "B/op",
            "extra": "4374 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2 - allocs/op",
            "value": 552,
            "unit": "allocs/op",
            "extra": "4374 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3",
            "value": 104249,
            "unit": "ns/op\t   29752 B/op\t     264 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3 - ns/op",
            "value": 104249,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3 - B/op",
            "value": 29752,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3 - allocs/op",
            "value": 264,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile",
            "value": 14380,
            "unit": "ns/op\t   10760 B/op\t     131 allocs/op",
            "extra": "91452 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile - ns/op",
            "value": 14380,
            "unit": "ns/op",
            "extra": "91452 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile - B/op",
            "value": 10760,
            "unit": "B/op",
            "extra": "91452 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile - allocs/op",
            "value": 131,
            "unit": "allocs/op",
            "extra": "91452 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid",
            "value": 28.82,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "44557042 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid - ns/op",
            "value": 28.82,
            "unit": "ns/op",
            "extra": "44557042 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "44557042 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "44557042 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite",
            "value": 50577,
            "unit": "ns/op\t   34696 B/op\t     226 allocs/op",
            "extra": "22310 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite - ns/op",
            "value": 50577,
            "unit": "ns/op",
            "extra": "22310 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite - B/op",
            "value": 34696,
            "unit": "B/op",
            "extra": "22310 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite - allocs/op",
            "value": 226,
            "unit": "allocs/op",
            "extra": "22310 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite",
            "value": 212762,
            "unit": "ns/op\t   54680 B/op\t    2069 allocs/op",
            "extra": "5746 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite - ns/op",
            "value": 212762,
            "unit": "ns/op",
            "extra": "5746 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite - B/op",
            "value": 54680,
            "unit": "B/op",
            "extra": "5746 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite - allocs/op",
            "value": 2069,
            "unit": "allocs/op",
            "extra": "5746 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite",
            "value": 154552,
            "unit": "ns/op\t   61104 B/op\t     721 allocs/op",
            "extra": "7528 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite - ns/op",
            "value": 154552,
            "unit": "ns/op",
            "extra": "7528 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite - B/op",
            "value": 61104,
            "unit": "B/op",
            "extra": "7528 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite - allocs/op",
            "value": 721,
            "unit": "allocs/op",
            "extra": "7528 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDIATWrite",
            "value": 231380104,
            "unit": "1210428822610050341A094101Federal Reserve Bank   My Bank Name                   ",
            "extra": "101 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDIATWrite - 1210428822610050341A094101Federal",
            "value": 231380104,
            "unit": "1210428822610050341A094101Federal",
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
          }
        ]
      },
      {
        "commit": {
          "author": {
            "name": "Sash",
            "username": "SashaMIT",
            "email": "sash.t.mitchell@gmail.com"
          },
          "committer": {
            "name": "GitHub",
            "username": "web-flow",
            "email": "noreply@github.com"
          },
          "id": "7cfa7dd4fb0b53376267b33f8004a6f60192e5ca",
          "message": "Reject an ACH file creation time after 23:59. (#1873)\n\nValidate accepted 2900 because the hour pattern allowed 24 through 29, and a four character time skipped the parse error.",
          "timestamp": "2026-09-28T20:11:29Z",
          "url": "https://github.com/moov-io/ach/commit/7cfa7dd4fb0b53376267b33f8004a6f60192e5ca"
        },
        "date": 1791170720574,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkAlphaFieldShort",
            "value": 961.1,
            "unit": "ns/op\t      80 B/op\t       3 allocs/op",
            "extra": "1227446 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldShort - ns/op",
            "value": 961.1,
            "unit": "ns/op",
            "extra": "1227446 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldShort - B/op",
            "value": 80,
            "unit": "B/op",
            "extra": "1227446 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldShort - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "1227446 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong",
            "value": 97.77,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "12097816 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong - ns/op",
            "value": 97.77,
            "unit": "ns/op",
            "extra": "12097816 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "12097816 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "12097816 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort",
            "value": 56.23,
            "unit": "ns/op\t      16 B/op\t       2 allocs/op",
            "extra": "20818440 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort - ns/op",
            "value": 56.23,
            "unit": "ns/op",
            "extra": "20818440 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "20818440 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "20818440 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong",
            "value": 25.49,
            "unit": "ns/op\t       8 B/op\t       1 allocs/op",
            "extra": "45260982 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong - ns/op",
            "value": 25.49,
            "unit": "ns/op",
            "extra": "45260982 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong - B/op",
            "value": 8,
            "unit": "B/op",
            "extra": "45260982 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "45260982 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField",
            "value": 14.79,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "83703879 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField - ns/op",
            "value": 14.79,
            "unit": "ns/op",
            "extra": "83703879 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "83703879 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "83703879 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField",
            "value": 5.617,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "213492078 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField - ns/op",
            "value": 5.617,
            "unit": "ns/op",
            "extra": "213492078 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "213492078 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "213492078 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles",
            "value": 248539,
            "unit": "ns/op\t   54142 B/op\t     306 allocs/op",
            "extra": "5131 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles - ns/op",
            "value": 248539,
            "unit": "ns/op",
            "extra": "5131 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles - B/op",
            "value": 54142,
            "unit": "B/op",
            "extra": "5131 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles - allocs/op",
            "value": 306,
            "unit": "allocs/op",
            "extra": "5131 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts",
            "value": 246710,
            "unit": "ns/op\t   54147 B/op\t     306 allocs/op",
            "extra": "4804 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts - ns/op",
            "value": 246710,
            "unit": "ns/op",
            "extra": "4804 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts - B/op",
            "value": 54147,
            "unit": "B/op",
            "extra": "4804 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts - allocs/op",
            "value": 306,
            "unit": "allocs/op",
            "extra": "4804 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir",
            "value": 139719,
            "unit": "ns/op\t   54557 B/op\t     310 allocs/op",
            "extra": "7568 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir - ns/op",
            "value": 139719,
            "unit": "ns/op",
            "extra": "7568 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir - B/op",
            "value": 54557,
            "unit": "B/op",
            "extra": "7568 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir - allocs/op",
            "value": 310,
            "unit": "allocs/op",
            "extra": "7568 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts",
            "value": 144630,
            "unit": "ns/op\t   54577 B/op\t     310 allocs/op",
            "extra": "7870 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts - ns/op",
            "value": 144630,
            "unit": "ns/op",
            "extra": "7870 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts - B/op",
            "value": 54577,
            "unit": "B/op",
            "extra": "7870 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts - allocs/op",
            "value": 310,
            "unit": "allocs/op",
            "extra": "7870 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups",
            "value": 308920,
            "unit": "ns/op\t   59909 B/op\t     366 allocs/op",
            "extra": "3954 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups - ns/op",
            "value": 308920,
            "unit": "ns/op",
            "extra": "3954 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups - B/op",
            "value": 59909,
            "unit": "B/op",
            "extra": "3954 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups - allocs/op",
            "value": 366,
            "unit": "allocs/op",
            "extra": "3954 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups",
            "value": 303066,
            "unit": "ns/op\t   59875 B/op\t     366 allocs/op",
            "extra": "3376 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups - ns/op",
            "value": 303066,
            "unit": "ns/op",
            "extra": "3376 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups - B/op",
            "value": 59875,
            "unit": "B/op",
            "extra": "3376 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups - allocs/op",
            "value": 366,
            "unit": "allocs/op",
            "extra": "3376 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups",
            "value": 304406,
            "unit": "ns/op\t   59864 B/op\t     366 allocs/op",
            "extra": "4234 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups - ns/op",
            "value": 304406,
            "unit": "ns/op",
            "extra": "4234 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups - B/op",
            "value": 59864,
            "unit": "B/op",
            "extra": "4234 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups - allocs/op",
            "value": 366,
            "unit": "allocs/op",
            "extra": "4234 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups",
            "value": 297789,
            "unit": "ns/op\t   59524 B/op\t     367 allocs/op",
            "extra": "4933 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups - ns/op",
            "value": 297789,
            "unit": "ns/op",
            "extra": "4933 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups - B/op",
            "value": 59524,
            "unit": "B/op",
            "extra": "4933 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups - allocs/op",
            "value": 367,
            "unit": "allocs/op",
            "extra": "4933 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader",
            "value": 243366218,
            "unit": "ns/op\t42450027 B/op\t  203653 allocs/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader - ns/op",
            "value": 243366218,
            "unit": "ns/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader - B/op",
            "value": 42450027,
            "unit": "B/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader - allocs/op",
            "value": 203653,
            "unit": "allocs/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator",
            "value": 79458488,
            "unit": "ns/op\t42424664 B/op\t  203600 allocs/op",
            "extra": "13 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator - ns/op",
            "value": 79458488,
            "unit": "ns/op",
            "extra": "13 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator - B/op",
            "value": 42424664,
            "unit": "B/op",
            "extra": "13 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator - allocs/op",
            "value": 203600,
            "unit": "allocs/op",
            "extra": "13 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader",
            "value": 454444763,
            "unit": "ns/op\t62131818 B/op\t  905794 allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader - ns/op",
            "value": 454444763,
            "unit": "ns/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader - B/op",
            "value": 62131818,
            "unit": "B/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader - allocs/op",
            "value": 905794,
            "unit": "allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator",
            "value": 151195279,
            "unit": "ns/op\t60508779 B/op\t  705839 allocs/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator - ns/op",
            "value": 151195279,
            "unit": "ns/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator - B/op",
            "value": 60508779,
            "unit": "B/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator - allocs/op",
            "value": 705839,
            "unit": "allocs/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader",
            "value": 1313926701,
            "unit": "ns/op\t215541784 B/op\t 1042867 allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader - ns/op",
            "value": 1313926701,
            "unit": "ns/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader - B/op",
            "value": 215541784,
            "unit": "B/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader - allocs/op",
            "value": 1042867,
            "unit": "allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator",
            "value": 398193689,
            "unit": "ns/op\t215424864 B/op\t 1042801 allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator - ns/op",
            "value": 398193689,
            "unit": "ns/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator - B/op",
            "value": 215424864,
            "unit": "B/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator - allocs/op",
            "value": 1042801,
            "unit": "allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader",
            "value": 2208611549,
            "unit": "ns/op\t313545880 B/op\t 4568039 allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader - ns/op",
            "value": 2208611549,
            "unit": "ns/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader - B/op",
            "value": 313545880,
            "unit": "B/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader - allocs/op",
            "value": 4568039,
            "unit": "allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator",
            "value": 757266066,
            "unit": "ns/op\t305430140 B/op\t 3568072 allocs/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator - ns/op",
            "value": 757266066,
            "unit": "ns/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator - B/op",
            "value": 305430140,
            "unit": "B/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator - allocs/op",
            "value": 3568072,
            "unit": "allocs/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead",
            "value": 42472,
            "unit": "ns/op\t   23312 B/op\t     114 allocs/op",
            "extra": "28494 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead - ns/op",
            "value": 42472,
            "unit": "ns/op",
            "extra": "28494 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead - B/op",
            "value": 23312,
            "unit": "B/op",
            "extra": "28494 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead - allocs/op",
            "value": 114,
            "unit": "allocs/op",
            "extra": "28494 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead",
            "value": 72369,
            "unit": "ns/op\t   26920 B/op\t     159 allocs/op",
            "extra": "16021 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead - ns/op",
            "value": 72369,
            "unit": "ns/op",
            "extra": "16021 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead - B/op",
            "value": 26920,
            "unit": "B/op",
            "extra": "16021 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead - allocs/op",
            "value": 159,
            "unit": "allocs/op",
            "extra": "16021 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead",
            "value": 31404,
            "unit": "ns/op\t   21848 B/op\t      77 allocs/op",
            "extra": "39302 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead - ns/op",
            "value": 31404,
            "unit": "ns/op",
            "extra": "39302 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead - B/op",
            "value": 21848,
            "unit": "B/op",
            "extra": "39302 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead - allocs/op",
            "value": 77,
            "unit": "allocs/op",
            "extra": "39302 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead",
            "value": 260338,
            "unit": "ns/op\t   56353 B/op\t     552 allocs/op",
            "extra": "4580 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead - ns/op",
            "value": 260338,
            "unit": "ns/op",
            "extra": "4580 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead - B/op",
            "value": 56353,
            "unit": "B/op",
            "extra": "4580 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead - allocs/op",
            "value": 552,
            "unit": "allocs/op",
            "extra": "4580 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2",
            "value": 273024,
            "unit": "ns/op\t   56353 B/op\t     552 allocs/op",
            "extra": "4437 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2 - ns/op",
            "value": 273024,
            "unit": "ns/op",
            "extra": "4437 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2 - B/op",
            "value": 56353,
            "unit": "B/op",
            "extra": "4437 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2 - allocs/op",
            "value": 552,
            "unit": "allocs/op",
            "extra": "4437 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3",
            "value": 103881,
            "unit": "ns/op\t   29752 B/op\t     264 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3 - ns/op",
            "value": 103881,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3 - B/op",
            "value": 29752,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3 - allocs/op",
            "value": 264,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile",
            "value": 14178,
            "unit": "ns/op\t   10760 B/op\t     131 allocs/op",
            "extra": "91531 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile - ns/op",
            "value": 14178,
            "unit": "ns/op",
            "extra": "91531 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile - B/op",
            "value": 10760,
            "unit": "B/op",
            "extra": "91531 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile - allocs/op",
            "value": 131,
            "unit": "allocs/op",
            "extra": "91531 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid",
            "value": 26.31,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "45411919 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid - ns/op",
            "value": 26.31,
            "unit": "ns/op",
            "extra": "45411919 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "45411919 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "45411919 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite",
            "value": 49959,
            "unit": "ns/op\t   34696 B/op\t     226 allocs/op",
            "extra": "22651 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite - ns/op",
            "value": 49959,
            "unit": "ns/op",
            "extra": "22651 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite - B/op",
            "value": 34696,
            "unit": "B/op",
            "extra": "22651 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite - allocs/op",
            "value": 226,
            "unit": "allocs/op",
            "extra": "22651 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite",
            "value": 213809,
            "unit": "ns/op\t   54680 B/op\t    2069 allocs/op",
            "extra": "5564 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite - ns/op",
            "value": 213809,
            "unit": "ns/op",
            "extra": "5564 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite - B/op",
            "value": 54680,
            "unit": "B/op",
            "extra": "5564 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite - allocs/op",
            "value": 2069,
            "unit": "allocs/op",
            "extra": "5564 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite",
            "value": 152448,
            "unit": "ns/op\t   61104 B/op\t     721 allocs/op",
            "extra": "7590 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite - ns/op",
            "value": 152448,
            "unit": "ns/op",
            "extra": "7590 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite - B/op",
            "value": 61104,
            "unit": "B/op",
            "extra": "7590 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite - allocs/op",
            "value": 721,
            "unit": "allocs/op",
            "extra": "7590 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDIATWrite",
            "value": 231380104,
            "unit": "1210428822610060325A094101Federal Reserve Bank   My Bank Name                   ",
            "extra": "101 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDIATWrite - 1210428822610060325A094101Federal",
            "value": 231380104,
            "unit": "1210428822610060325A094101Federal",
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
          }
        ]
      },
      {
        "commit": {
          "author": {
            "name": "Sash",
            "username": "SashaMIT",
            "email": "sash.t.mitchell@gmail.com"
          },
          "committer": {
            "name": "GitHub",
            "username": "web-flow",
            "email": "noreply@github.com"
          },
          "id": "7cfa7dd4fb0b53376267b33f8004a6f60192e5ca",
          "message": "Reject an ACH file creation time after 23:59. (#1873)\n\nValidate accepted 2900 because the hour pattern allowed 24 through 29, and a four character time skipped the parse error.",
          "timestamp": "2026-09-28T20:11:29Z",
          "url": "https://github.com/moov-io/ach/commit/7cfa7dd4fb0b53376267b33f8004a6f60192e5ca"
        },
        "date": 1791259949113,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkAlphaFieldShort",
            "value": 991.3,
            "unit": "ns/op\t      80 B/op\t       3 allocs/op",
            "extra": "1241214 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldShort - ns/op",
            "value": 991.3,
            "unit": "ns/op",
            "extra": "1241214 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldShort - B/op",
            "value": 80,
            "unit": "B/op",
            "extra": "1241214 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldShort - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "1241214 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong",
            "value": 98.25,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "12089359 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong - ns/op",
            "value": 98.25,
            "unit": "ns/op",
            "extra": "12089359 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "12089359 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "12089359 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort",
            "value": 56.49,
            "unit": "ns/op\t      16 B/op\t       2 allocs/op",
            "extra": "20640375 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort - ns/op",
            "value": 56.49,
            "unit": "ns/op",
            "extra": "20640375 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "20640375 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "20640375 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong",
            "value": 25.58,
            "unit": "ns/op\t       8 B/op\t       1 allocs/op",
            "extra": "44348332 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong - ns/op",
            "value": 25.58,
            "unit": "ns/op",
            "extra": "44348332 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong - B/op",
            "value": 8,
            "unit": "B/op",
            "extra": "44348332 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "44348332 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField",
            "value": 14.86,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "82830351 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField - ns/op",
            "value": 14.86,
            "unit": "ns/op",
            "extra": "82830351 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "82830351 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "82830351 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField",
            "value": 5.621,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "213782220 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField - ns/op",
            "value": 5.621,
            "unit": "ns/op",
            "extra": "213782220 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "213782220 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "213782220 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles",
            "value": 256605,
            "unit": "ns/op\t   54140 B/op\t     306 allocs/op",
            "extra": "5152 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles - ns/op",
            "value": 256605,
            "unit": "ns/op",
            "extra": "5152 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles - B/op",
            "value": 54140,
            "unit": "B/op",
            "extra": "5152 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles - allocs/op",
            "value": 306,
            "unit": "allocs/op",
            "extra": "5152 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts",
            "value": 258464,
            "unit": "ns/op\t   54159 B/op\t     306 allocs/op",
            "extra": "4508 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts - ns/op",
            "value": 258464,
            "unit": "ns/op",
            "extra": "4508 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts - B/op",
            "value": 54159,
            "unit": "B/op",
            "extra": "4508 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts - allocs/op",
            "value": 306,
            "unit": "allocs/op",
            "extra": "4508 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir",
            "value": 146787,
            "unit": "ns/op\t   54539 B/op\t     310 allocs/op",
            "extra": "7874 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir - ns/op",
            "value": 146787,
            "unit": "ns/op",
            "extra": "7874 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir - B/op",
            "value": 54539,
            "unit": "B/op",
            "extra": "7874 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir - allocs/op",
            "value": 310,
            "unit": "allocs/op",
            "extra": "7874 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts",
            "value": 146387,
            "unit": "ns/op\t   54584 B/op\t     310 allocs/op",
            "extra": "7561 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts - ns/op",
            "value": 146387,
            "unit": "ns/op",
            "extra": "7561 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts - B/op",
            "value": 54584,
            "unit": "B/op",
            "extra": "7561 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts - allocs/op",
            "value": 310,
            "unit": "allocs/op",
            "extra": "7561 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups",
            "value": 319271,
            "unit": "ns/op\t   59898 B/op\t     366 allocs/op",
            "extra": "3950 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups - ns/op",
            "value": 319271,
            "unit": "ns/op",
            "extra": "3950 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups - B/op",
            "value": 59898,
            "unit": "B/op",
            "extra": "3950 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups - allocs/op",
            "value": 366,
            "unit": "allocs/op",
            "extra": "3950 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups",
            "value": 312869,
            "unit": "ns/op\t   59876 B/op\t     366 allocs/op",
            "extra": "3976 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups - ns/op",
            "value": 312869,
            "unit": "ns/op",
            "extra": "3976 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups - B/op",
            "value": 59876,
            "unit": "B/op",
            "extra": "3976 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups - allocs/op",
            "value": 366,
            "unit": "allocs/op",
            "extra": "3976 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups",
            "value": 309678,
            "unit": "ns/op\t   59875 B/op\t     366 allocs/op",
            "extra": "4185 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups - ns/op",
            "value": 309678,
            "unit": "ns/op",
            "extra": "4185 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups - B/op",
            "value": 59875,
            "unit": "B/op",
            "extra": "4185 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups - allocs/op",
            "value": 366,
            "unit": "allocs/op",
            "extra": "4185 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups",
            "value": 293577,
            "unit": "ns/op\t   59574 B/op\t     367 allocs/op",
            "extra": "4824 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups - ns/op",
            "value": 293577,
            "unit": "ns/op",
            "extra": "4824 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups - B/op",
            "value": 59574,
            "unit": "B/op",
            "extra": "4824 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups - allocs/op",
            "value": 367,
            "unit": "allocs/op",
            "extra": "4824 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader",
            "value": 247113137,
            "unit": "ns/op\t42450289 B/op\t  203654 allocs/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader - ns/op",
            "value": 247113137,
            "unit": "ns/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader - B/op",
            "value": 42450289,
            "unit": "B/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader - allocs/op",
            "value": 203654,
            "unit": "allocs/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator",
            "value": 80362731,
            "unit": "ns/op\t42424739 B/op\t  203601 allocs/op",
            "extra": "13 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator - ns/op",
            "value": 80362731,
            "unit": "ns/op",
            "extra": "13 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator - B/op",
            "value": 42424739,
            "unit": "B/op",
            "extra": "13 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator - allocs/op",
            "value": 203601,
            "unit": "allocs/op",
            "extra": "13 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader",
            "value": 453444931,
            "unit": "ns/op\t62131722 B/op\t  905792 allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader - ns/op",
            "value": 453444931,
            "unit": "ns/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader - B/op",
            "value": 62131722,
            "unit": "B/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader - allocs/op",
            "value": 905792,
            "unit": "allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator",
            "value": 152047013,
            "unit": "ns/op\t60508841 B/op\t  705840 allocs/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator - ns/op",
            "value": 152047013,
            "unit": "ns/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator - B/op",
            "value": 60508841,
            "unit": "B/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator - allocs/op",
            "value": 705840,
            "unit": "allocs/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader",
            "value": 1312036659,
            "unit": "ns/op\t215541928 B/op\t 1042869 allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader - ns/op",
            "value": 1312036659,
            "unit": "ns/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader - B/op",
            "value": 215541928,
            "unit": "B/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader - allocs/op",
            "value": 1042869,
            "unit": "allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator",
            "value": 396496535,
            "unit": "ns/op\t215424949 B/op\t 1042802 allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator - ns/op",
            "value": 396496535,
            "unit": "ns/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator - B/op",
            "value": 215424949,
            "unit": "B/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator - allocs/op",
            "value": 1042802,
            "unit": "allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader",
            "value": 2238130048,
            "unit": "ns/op\t313545768 B/op\t 4568037 allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader - ns/op",
            "value": 2238130048,
            "unit": "ns/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader - B/op",
            "value": 313545768,
            "unit": "B/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader - allocs/op",
            "value": 4568037,
            "unit": "allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator",
            "value": 756754260,
            "unit": "ns/op\t305430308 B/op\t 3568074 allocs/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator - ns/op",
            "value": 756754260,
            "unit": "ns/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator - B/op",
            "value": 305430308,
            "unit": "B/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator - allocs/op",
            "value": 3568074,
            "unit": "allocs/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead",
            "value": 41796,
            "unit": "ns/op\t   23312 B/op\t     114 allocs/op",
            "extra": "28924 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead - ns/op",
            "value": 41796,
            "unit": "ns/op",
            "extra": "28924 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead - B/op",
            "value": 23312,
            "unit": "B/op",
            "extra": "28924 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead - allocs/op",
            "value": 114,
            "unit": "allocs/op",
            "extra": "28924 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead",
            "value": 71651,
            "unit": "ns/op\t   26920 B/op\t     159 allocs/op",
            "extra": "15916 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead - ns/op",
            "value": 71651,
            "unit": "ns/op",
            "extra": "15916 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead - B/op",
            "value": 26920,
            "unit": "B/op",
            "extra": "15916 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead - allocs/op",
            "value": 159,
            "unit": "allocs/op",
            "extra": "15916 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead",
            "value": 31183,
            "unit": "ns/op\t   21848 B/op\t      77 allocs/op",
            "extra": "39500 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead - ns/op",
            "value": 31183,
            "unit": "ns/op",
            "extra": "39500 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead - B/op",
            "value": 21848,
            "unit": "B/op",
            "extra": "39500 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead - allocs/op",
            "value": 77,
            "unit": "allocs/op",
            "extra": "39500 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead",
            "value": 273363,
            "unit": "ns/op\t   56353 B/op\t     552 allocs/op",
            "extra": "4624 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead - ns/op",
            "value": 273363,
            "unit": "ns/op",
            "extra": "4624 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead - B/op",
            "value": 56353,
            "unit": "B/op",
            "extra": "4624 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead - allocs/op",
            "value": 552,
            "unit": "allocs/op",
            "extra": "4624 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2",
            "value": 273737,
            "unit": "ns/op\t   56353 B/op\t     552 allocs/op",
            "extra": "4560 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2 - ns/op",
            "value": 273737,
            "unit": "ns/op",
            "extra": "4560 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2 - B/op",
            "value": 56353,
            "unit": "B/op",
            "extra": "4560 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2 - allocs/op",
            "value": 552,
            "unit": "allocs/op",
            "extra": "4560 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3",
            "value": 103793,
            "unit": "ns/op\t   29752 B/op\t     264 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3 - ns/op",
            "value": 103793,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3 - B/op",
            "value": 29752,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3 - allocs/op",
            "value": 264,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile",
            "value": 14311,
            "unit": "ns/op\t   10760 B/op\t     131 allocs/op",
            "extra": "92371 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile - ns/op",
            "value": 14311,
            "unit": "ns/op",
            "extra": "92371 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile - B/op",
            "value": 10760,
            "unit": "B/op",
            "extra": "92371 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile - allocs/op",
            "value": 131,
            "unit": "allocs/op",
            "extra": "92371 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid",
            "value": 26.49,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "42912240 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid - ns/op",
            "value": 26.49,
            "unit": "ns/op",
            "extra": "42912240 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "42912240 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "42912240 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite",
            "value": 49855,
            "unit": "ns/op\t   34696 B/op\t     226 allocs/op",
            "extra": "22731 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite - ns/op",
            "value": 49855,
            "unit": "ns/op",
            "extra": "22731 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite - B/op",
            "value": 34696,
            "unit": "B/op",
            "extra": "22731 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite - allocs/op",
            "value": 226,
            "unit": "allocs/op",
            "extra": "22731 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite",
            "value": 213540,
            "unit": "ns/op\t   54680 B/op\t    2069 allocs/op",
            "extra": "5550 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite - ns/op",
            "value": 213540,
            "unit": "ns/op",
            "extra": "5550 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite - B/op",
            "value": 54680,
            "unit": "B/op",
            "extra": "5550 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite - allocs/op",
            "value": 2069,
            "unit": "allocs/op",
            "extra": "5550 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite",
            "value": 153917,
            "unit": "ns/op\t   61104 B/op\t     721 allocs/op",
            "extra": "7572 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite - ns/op",
            "value": 153917,
            "unit": "ns/op",
            "extra": "7572 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite - B/op",
            "value": 61104,
            "unit": "B/op",
            "extra": "7572 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite - allocs/op",
            "value": 721,
            "unit": "allocs/op",
            "extra": "7572 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDIATWrite",
            "value": 231380104,
            "unit": "1210428822610070412A094101Federal Reserve Bank   My Bank Name                   ",
            "extra": "101 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDIATWrite - 1210428822610070412A094101Federal",
            "value": 231380104,
            "unit": "1210428822610070412A094101Federal",
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
          }
        ]
      },
      {
        "commit": {
          "author": {
            "name": "Sash",
            "username": "SashaMIT",
            "email": "sash.t.mitchell@gmail.com"
          },
          "committer": {
            "name": "GitHub",
            "username": "web-flow",
            "email": "noreply@github.com"
          },
          "id": "7cfa7dd4fb0b53376267b33f8004a6f60192e5ca",
          "message": "Reject an ACH file creation time after 23:59. (#1873)\n\nValidate accepted 2900 because the hour pattern allowed 24 through 29, and a four character time skipped the parse error.",
          "timestamp": "2026-09-28T20:11:29Z",
          "url": "https://github.com/moov-io/ach/commit/7cfa7dd4fb0b53376267b33f8004a6f60192e5ca"
        },
        "date": 1791344390134,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkAlphaFieldShort",
            "value": 961.5,
            "unit": "ns/op\t      80 B/op\t       3 allocs/op",
            "extra": "1255209 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldShort - ns/op",
            "value": 961.5,
            "unit": "ns/op",
            "extra": "1255209 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldShort - B/op",
            "value": 80,
            "unit": "B/op",
            "extra": "1255209 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldShort - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "1255209 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong",
            "value": 97.8,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "12240066 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong - ns/op",
            "value": 97.8,
            "unit": "ns/op",
            "extra": "12240066 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "12240066 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "12240066 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort",
            "value": 56.23,
            "unit": "ns/op\t      16 B/op\t       2 allocs/op",
            "extra": "20882254 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort - ns/op",
            "value": 56.23,
            "unit": "ns/op",
            "extra": "20882254 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "20882254 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "20882254 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong",
            "value": 25.47,
            "unit": "ns/op\t       8 B/op\t       1 allocs/op",
            "extra": "45810032 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong - ns/op",
            "value": 25.47,
            "unit": "ns/op",
            "extra": "45810032 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong - B/op",
            "value": 8,
            "unit": "B/op",
            "extra": "45810032 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "45810032 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField",
            "value": 15.02,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "81421345 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField - ns/op",
            "value": 15.02,
            "unit": "ns/op",
            "extra": "81421345 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "81421345 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "81421345 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField",
            "value": 5.654,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "195897272 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField - ns/op",
            "value": 5.654,
            "unit": "ns/op",
            "extra": "195897272 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "195897272 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "195897272 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles",
            "value": 250399,
            "unit": "ns/op\t   54148 B/op\t     306 allocs/op",
            "extra": "5080 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles - ns/op",
            "value": 250399,
            "unit": "ns/op",
            "extra": "5080 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles - B/op",
            "value": 54148,
            "unit": "B/op",
            "extra": "5080 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles - allocs/op",
            "value": 306,
            "unit": "allocs/op",
            "extra": "5080 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts",
            "value": 244351,
            "unit": "ns/op\t   54163 B/op\t     306 allocs/op",
            "extra": "4411 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts - ns/op",
            "value": 244351,
            "unit": "ns/op",
            "extra": "4411 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts - B/op",
            "value": 54163,
            "unit": "B/op",
            "extra": "4411 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts - allocs/op",
            "value": 306,
            "unit": "allocs/op",
            "extra": "4411 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir",
            "value": 148591,
            "unit": "ns/op\t   54549 B/op\t     310 allocs/op",
            "extra": "7598 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir - ns/op",
            "value": 148591,
            "unit": "ns/op",
            "extra": "7598 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir - B/op",
            "value": 54549,
            "unit": "B/op",
            "extra": "7598 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir - allocs/op",
            "value": 310,
            "unit": "allocs/op",
            "extra": "7598 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts",
            "value": 146170,
            "unit": "ns/op\t   54586 B/op\t     310 allocs/op",
            "extra": "7592 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts - ns/op",
            "value": 146170,
            "unit": "ns/op",
            "extra": "7592 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts - B/op",
            "value": 54586,
            "unit": "B/op",
            "extra": "7592 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts - allocs/op",
            "value": 310,
            "unit": "allocs/op",
            "extra": "7592 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups",
            "value": 316217,
            "unit": "ns/op\t   59853 B/op\t     366 allocs/op",
            "extra": "4044 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups - ns/op",
            "value": 316217,
            "unit": "ns/op",
            "extra": "4044 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups - B/op",
            "value": 59853,
            "unit": "B/op",
            "extra": "4044 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups - allocs/op",
            "value": 366,
            "unit": "allocs/op",
            "extra": "4044 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups",
            "value": 304962,
            "unit": "ns/op\t   59935 B/op\t     366 allocs/op",
            "extra": "4071 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups - ns/op",
            "value": 304962,
            "unit": "ns/op",
            "extra": "4071 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups - B/op",
            "value": 59935,
            "unit": "B/op",
            "extra": "4071 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups - allocs/op",
            "value": 366,
            "unit": "allocs/op",
            "extra": "4071 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups",
            "value": 301661,
            "unit": "ns/op\t   59925 B/op\t     366 allocs/op",
            "extra": "4080 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups - ns/op",
            "value": 301661,
            "unit": "ns/op",
            "extra": "4080 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups - B/op",
            "value": 59925,
            "unit": "B/op",
            "extra": "4080 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups - allocs/op",
            "value": 366,
            "unit": "allocs/op",
            "extra": "4080 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups",
            "value": 293099,
            "unit": "ns/op\t   59551 B/op\t     367 allocs/op",
            "extra": "4854 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups - ns/op",
            "value": 293099,
            "unit": "ns/op",
            "extra": "4854 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups - B/op",
            "value": 59551,
            "unit": "B/op",
            "extra": "4854 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups - allocs/op",
            "value": 367,
            "unit": "allocs/op",
            "extra": "4854 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader",
            "value": 241994520,
            "unit": "ns/op\t42450363 B/op\t  203655 allocs/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader - ns/op",
            "value": 241994520,
            "unit": "ns/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader - B/op",
            "value": 42450363,
            "unit": "B/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader - allocs/op",
            "value": 203655,
            "unit": "allocs/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator",
            "value": 79676788,
            "unit": "ns/op\t42424757 B/op\t  203601 allocs/op",
            "extra": "14 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator - ns/op",
            "value": 79676788,
            "unit": "ns/op",
            "extra": "14 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator - B/op",
            "value": 42424757,
            "unit": "B/op",
            "extra": "14 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator - allocs/op",
            "value": 203601,
            "unit": "allocs/op",
            "extra": "14 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader",
            "value": 449973218,
            "unit": "ns/op\t62131770 B/op\t  905793 allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader - ns/op",
            "value": 449973218,
            "unit": "ns/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader - B/op",
            "value": 62131770,
            "unit": "B/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader - allocs/op",
            "value": 905793,
            "unit": "allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator",
            "value": 149266691,
            "unit": "ns/op\t60508713 B/op\t  705838 allocs/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator - ns/op",
            "value": 149266691,
            "unit": "ns/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator - B/op",
            "value": 60508713,
            "unit": "B/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator - allocs/op",
            "value": 705838,
            "unit": "allocs/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader",
            "value": 1310402212,
            "unit": "ns/op\t215541720 B/op\t 1042866 allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader - ns/op",
            "value": 1310402212,
            "unit": "ns/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader - B/op",
            "value": 215541720,
            "unit": "B/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader - allocs/op",
            "value": 1042866,
            "unit": "allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator",
            "value": 392614798,
            "unit": "ns/op\t215424853 B/op\t 1042801 allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator - ns/op",
            "value": 392614798,
            "unit": "ns/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator - B/op",
            "value": 215424853,
            "unit": "B/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator - allocs/op",
            "value": 1042801,
            "unit": "allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader",
            "value": 2211792734,
            "unit": "ns/op\t313545976 B/op\t 4568040 allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader - ns/op",
            "value": 2211792734,
            "unit": "ns/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader - B/op",
            "value": 313545976,
            "unit": "B/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader - allocs/op",
            "value": 4568040,
            "unit": "allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator",
            "value": 754105776,
            "unit": "ns/op\t305430284 B/op\t 3568074 allocs/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator - ns/op",
            "value": 754105776,
            "unit": "ns/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator - B/op",
            "value": 305430284,
            "unit": "B/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator - allocs/op",
            "value": 3568074,
            "unit": "allocs/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead",
            "value": 41589,
            "unit": "ns/op\t   23312 B/op\t     114 allocs/op",
            "extra": "28791 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead - ns/op",
            "value": 41589,
            "unit": "ns/op",
            "extra": "28791 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead - B/op",
            "value": 23312,
            "unit": "B/op",
            "extra": "28791 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead - allocs/op",
            "value": 114,
            "unit": "allocs/op",
            "extra": "28791 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead",
            "value": 72108,
            "unit": "ns/op\t   26920 B/op\t     159 allocs/op",
            "extra": "16134 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead - ns/op",
            "value": 72108,
            "unit": "ns/op",
            "extra": "16134 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead - B/op",
            "value": 26920,
            "unit": "B/op",
            "extra": "16134 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead - allocs/op",
            "value": 159,
            "unit": "allocs/op",
            "extra": "16134 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead",
            "value": 30877,
            "unit": "ns/op\t   21848 B/op\t      77 allocs/op",
            "extra": "39476 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead - ns/op",
            "value": 30877,
            "unit": "ns/op",
            "extra": "39476 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead - B/op",
            "value": 21848,
            "unit": "B/op",
            "extra": "39476 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead - allocs/op",
            "value": 77,
            "unit": "allocs/op",
            "extra": "39476 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead",
            "value": 269962,
            "unit": "ns/op\t   56353 B/op\t     552 allocs/op",
            "extra": "4537 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead - ns/op",
            "value": 269962,
            "unit": "ns/op",
            "extra": "4537 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead - B/op",
            "value": 56353,
            "unit": "B/op",
            "extra": "4537 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead - allocs/op",
            "value": 552,
            "unit": "allocs/op",
            "extra": "4537 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2",
            "value": 271779,
            "unit": "ns/op\t   56352 B/op\t     552 allocs/op",
            "extra": "4549 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2 - ns/op",
            "value": 271779,
            "unit": "ns/op",
            "extra": "4549 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2 - B/op",
            "value": 56352,
            "unit": "B/op",
            "extra": "4549 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2 - allocs/op",
            "value": 552,
            "unit": "allocs/op",
            "extra": "4549 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3",
            "value": 102536,
            "unit": "ns/op\t   29752 B/op\t     264 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3 - ns/op",
            "value": 102536,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3 - B/op",
            "value": 29752,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3 - allocs/op",
            "value": 264,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile",
            "value": 14071,
            "unit": "ns/op\t   10760 B/op\t     131 allocs/op",
            "extra": "92350 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile - ns/op",
            "value": 14071,
            "unit": "ns/op",
            "extra": "92350 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile - B/op",
            "value": 10760,
            "unit": "B/op",
            "extra": "92350 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile - allocs/op",
            "value": 131,
            "unit": "allocs/op",
            "extra": "92350 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid",
            "value": 26.29,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "45281586 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid - ns/op",
            "value": 26.29,
            "unit": "ns/op",
            "extra": "45281586 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "45281586 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "45281586 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite",
            "value": 49399,
            "unit": "ns/op\t   34696 B/op\t     226 allocs/op",
            "extra": "22970 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite - ns/op",
            "value": 49399,
            "unit": "ns/op",
            "extra": "22970 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite - B/op",
            "value": 34696,
            "unit": "B/op",
            "extra": "22970 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite - allocs/op",
            "value": 226,
            "unit": "allocs/op",
            "extra": "22970 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite",
            "value": 208020,
            "unit": "ns/op\t   54680 B/op\t    2069 allocs/op",
            "extra": "6030 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite - ns/op",
            "value": 208020,
            "unit": "ns/op",
            "extra": "6030 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite - B/op",
            "value": 54680,
            "unit": "B/op",
            "extra": "6030 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite - allocs/op",
            "value": 2069,
            "unit": "allocs/op",
            "extra": "6030 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite",
            "value": 153010,
            "unit": "ns/op\t   61104 B/op\t     721 allocs/op",
            "extra": "7675 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite - ns/op",
            "value": 153010,
            "unit": "ns/op",
            "extra": "7675 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite - B/op",
            "value": 61104,
            "unit": "B/op",
            "extra": "7675 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite - allocs/op",
            "value": 721,
            "unit": "allocs/op",
            "extra": "7675 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDIATWrite",
            "value": 231380104,
            "unit": "1210428822610080339A094101Federal Reserve Bank   My Bank Name                   ",
            "extra": "101 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDIATWrite - 1210428822610080339A094101Federal",
            "value": 231380104,
            "unit": "1210428822610080339A094101Federal",
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
          }
        ]
      },
      {
        "commit": {
          "author": {
            "name": "Adam Shannon",
            "username": "adamdecaf",
            "email": "adamkshannon@gmail.com"
          },
          "committer": {
            "name": "moov-bot",
            "email": "oss@moov.io"
          },
          "id": "eef0a25a350c3c25eafb4e06d8c80124da40cbc0",
          "message": "ci: push release wasm to master again\n\nThe WebUI job should run make dist-webui and push docs/webui/assets/ach.wasm\nto master with [skip ci], not open a PR.",
          "timestamp": "2026-10-07T21:09:28Z",
          "url": "https://github.com/moov-io/ach/commit/eef0a25a350c3c25eafb4e06d8c80124da40cbc0"
        },
        "date": 1791431628360,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkAlphaFieldShort",
            "value": 994.7,
            "unit": "ns/op\t      80 B/op\t       3 allocs/op",
            "extra": "1203289 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldShort - ns/op",
            "value": 994.7,
            "unit": "ns/op",
            "extra": "1203289 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldShort - B/op",
            "value": 80,
            "unit": "B/op",
            "extra": "1203289 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldShort - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "1203289 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong",
            "value": 97.3,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "12104596 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong - ns/op",
            "value": 97.3,
            "unit": "ns/op",
            "extra": "12104596 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "12104596 times\n4 procs"
          },
          {
            "name": "BenchmarkAlphaFieldLong - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "12104596 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort",
            "value": 56.14,
            "unit": "ns/op\t      16 B/op\t       2 allocs/op",
            "extra": "20684589 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort - ns/op",
            "value": 56.14,
            "unit": "ns/op",
            "extra": "20684589 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "20684589 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldShort - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "20684589 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong",
            "value": 25.8,
            "unit": "ns/op\t       8 B/op\t       1 allocs/op",
            "extra": "44291118 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong - ns/op",
            "value": 25.8,
            "unit": "ns/op",
            "extra": "44291118 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong - B/op",
            "value": 8,
            "unit": "B/op",
            "extra": "44291118 times\n4 procs"
          },
          {
            "name": "BenchmarkNumericFieldLong - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "44291118 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField",
            "value": 14.96,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "80084178 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField - ns/op",
            "value": 14.96,
            "unit": "ns/op",
            "extra": "80084178 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "80084178 times\n4 procs"
          },
          {
            "name": "BenchmarkParseNumField - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "80084178 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField",
            "value": 5.915,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "208039592 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField - ns/op",
            "value": 5.915,
            "unit": "ns/op",
            "extra": "208039592 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "208039592 times\n4 procs"
          },
          {
            "name": "BenchmarkParseStringField - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "208039592 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles",
            "value": 265355,
            "unit": "ns/op\t   54144 B/op\t     306 allocs/op",
            "extra": "4971 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles - ns/op",
            "value": 265355,
            "unit": "ns/op",
            "extra": "4971 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles - B/op",
            "value": 54144,
            "unit": "B/op",
            "extra": "4971 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles - allocs/op",
            "value": 306,
            "unit": "allocs/op",
            "extra": "4971 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts",
            "value": 247099,
            "unit": "ns/op\t   54152 B/op\t     306 allocs/op",
            "extra": "4822 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts - ns/op",
            "value": 247099,
            "unit": "ns/op",
            "extra": "4822 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts - B/op",
            "value": 54152,
            "unit": "B/op",
            "extra": "4822 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_ValidateOpts - allocs/op",
            "value": 306,
            "unit": "allocs/op",
            "extra": "4822 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir",
            "value": 148664,
            "unit": "ns/op\t   54550 B/op\t     310 allocs/op",
            "extra": "8265 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir - ns/op",
            "value": 148664,
            "unit": "ns/op",
            "extra": "8265 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir - B/op",
            "value": 54550,
            "unit": "B/op",
            "extra": "8265 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir - allocs/op",
            "value": 310,
            "unit": "allocs/op",
            "extra": "8265 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts",
            "value": 143894,
            "unit": "ns/op\t   54587 B/op\t     310 allocs/op",
            "extra": "7213 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts - ns/op",
            "value": 143894,
            "unit": "ns/op",
            "extra": "7213 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts - B/op",
            "value": 54587,
            "unit": "B/op",
            "extra": "7213 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeDir_ValidateOpts - allocs/op",
            "value": 310,
            "unit": "allocs/op",
            "extra": "7213 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups",
            "value": 311312,
            "unit": "ns/op\t   59856 B/op\t     366 allocs/op",
            "extra": "4132 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups - ns/op",
            "value": 311312,
            "unit": "ns/op",
            "extra": "4132 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups - B/op",
            "value": 59856,
            "unit": "B/op",
            "extra": "4132 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_3Groups - allocs/op",
            "value": 366,
            "unit": "allocs/op",
            "extra": "4132 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups",
            "value": 302627,
            "unit": "ns/op\t   59926 B/op\t     366 allocs/op",
            "extra": "4063 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups - ns/op",
            "value": 302627,
            "unit": "ns/op",
            "extra": "4063 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups - B/op",
            "value": 59926,
            "unit": "B/op",
            "extra": "4063 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_5Groups - allocs/op",
            "value": 366,
            "unit": "allocs/op",
            "extra": "4063 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups",
            "value": 304171,
            "unit": "ns/op\t   59881 B/op\t     366 allocs/op",
            "extra": "4191 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups - ns/op",
            "value": 304171,
            "unit": "ns/op",
            "extra": "4191 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups - B/op",
            "value": 59881,
            "unit": "B/op",
            "extra": "4191 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_10Groups - allocs/op",
            "value": 366,
            "unit": "allocs/op",
            "extra": "4191 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups",
            "value": 289357,
            "unit": "ns/op\t   59758 B/op\t     367 allocs/op",
            "extra": "4347 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups - ns/op",
            "value": 289357,
            "unit": "ns/op",
            "extra": "4347 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups - B/op",
            "value": 59758,
            "unit": "B/op",
            "extra": "4347 times\n4 procs"
          },
          {
            "name": "BenchmarkMergeFiles/MergeFiles_100Groups - allocs/op",
            "value": 367,
            "unit": "allocs/op",
            "extra": "4347 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader",
            "value": 241358387,
            "unit": "ns/op\t42449980 B/op\t  203653 allocs/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader - ns/op",
            "value": 241358387,
            "unit": "ns/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader - B/op",
            "value": 42449980,
            "unit": "B/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/reader - allocs/op",
            "value": 203653,
            "unit": "allocs/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator",
            "value": 80081738,
            "unit": "ns/op\t42424662 B/op\t  203600 allocs/op",
            "extra": "14 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator - ns/op",
            "value": 80081738,
            "unit": "ns/op",
            "extra": "14 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator - B/op",
            "value": 42424662,
            "unit": "B/op",
            "extra": "14 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_200_batches_100000_entries/iterator - allocs/op",
            "value": 203600,
            "unit": "allocs/op",
            "extra": "14 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader",
            "value": 451426097,
            "unit": "ns/op\t62131781 B/op\t  905793 allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader - ns/op",
            "value": 451426097,
            "unit": "ns/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader - B/op",
            "value": 62131781,
            "unit": "B/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/reader - allocs/op",
            "value": 905793,
            "unit": "allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator",
            "value": 153711385,
            "unit": "ns/op\t60508829 B/op\t  705839 allocs/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator - ns/op",
            "value": 153711385,
            "unit": "ns/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator - B/op",
            "value": 60508829,
            "unit": "B/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_200_batches_100000_entries/iterator - allocs/op",
            "value": 705839,
            "unit": "allocs/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader",
            "value": 1292083726,
            "unit": "ns/op\t215542072 B/op\t 1042871 allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader - ns/op",
            "value": 1292083726,
            "unit": "ns/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader - B/op",
            "value": 215542072,
            "unit": "B/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/reader - allocs/op",
            "value": 1042871,
            "unit": "allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator",
            "value": 407758168,
            "unit": "ns/op\t215424853 B/op\t 1042801 allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator - ns/op",
            "value": 407758168,
            "unit": "ns/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator - B/op",
            "value": 215424853,
            "unit": "B/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/PPD_2500_batches_500000_entries/iterator - allocs/op",
            "value": 1042801,
            "unit": "allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader",
            "value": 2246030458,
            "unit": "ns/op\t313545976 B/op\t 4568040 allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader - ns/op",
            "value": 2246030458,
            "unit": "ns/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader - B/op",
            "value": 313545976,
            "unit": "B/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/reader - allocs/op",
            "value": 4568040,
            "unit": "allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator",
            "value": 768235757,
            "unit": "ns/op\t305430236 B/op\t 3568073 allocs/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator - ns/op",
            "value": 768235757,
            "unit": "ns/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator - B/op",
            "value": 305430236,
            "unit": "B/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "Benchmark_ReadLargeFile/CCD+addenda05_2500_batches_500000_entries/iterator - allocs/op",
            "value": 3568073,
            "unit": "allocs/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead",
            "value": 42110,
            "unit": "ns/op\t   23312 B/op\t     114 allocs/op",
            "extra": "28792 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead - ns/op",
            "value": 42110,
            "unit": "ns/op",
            "extra": "28792 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead - B/op",
            "value": 23312,
            "unit": "B/op",
            "extra": "28792 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitRead - allocs/op",
            "value": 114,
            "unit": "allocs/op",
            "extra": "28792 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead",
            "value": 71114,
            "unit": "ns/op\t   26920 B/op\t     159 allocs/op",
            "extra": "16129 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead - ns/op",
            "value": 71114,
            "unit": "ns/op",
            "extra": "16129 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead - B/op",
            "value": 26920,
            "unit": "B/op",
            "extra": "16129 times\n4 procs"
          },
          {
            "name": "BenchmarkWEBDebitRead - allocs/op",
            "value": 159,
            "unit": "allocs/op",
            "extra": "16129 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead",
            "value": 31103,
            "unit": "ns/op\t   21848 B/op\t      77 allocs/op",
            "extra": "39313 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead - ns/op",
            "value": 31103,
            "unit": "ns/op",
            "extra": "39313 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead - B/op",
            "value": 21848,
            "unit": "B/op",
            "extra": "39313 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDDebitFixedLengthRead - allocs/op",
            "value": 77,
            "unit": "allocs/op",
            "extra": "39313 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead",
            "value": 270866,
            "unit": "ns/op\t   56353 B/op\t     552 allocs/op",
            "extra": "4516 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead - ns/op",
            "value": 270866,
            "unit": "ns/op",
            "extra": "4516 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead - B/op",
            "value": 56353,
            "unit": "B/op",
            "extra": "4516 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead - allocs/op",
            "value": 552,
            "unit": "allocs/op",
            "extra": "4516 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2",
            "value": 270044,
            "unit": "ns/op\t   56353 B/op\t     552 allocs/op",
            "extra": "4514 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2 - ns/op",
            "value": 270044,
            "unit": "ns/op",
            "extra": "4514 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2 - B/op",
            "value": 56353,
            "unit": "B/op",
            "extra": "4514 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead2 - allocs/op",
            "value": 552,
            "unit": "allocs/op",
            "extra": "4514 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3",
            "value": 102236,
            "unit": "ns/op\t   29752 B/op\t     264 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3 - ns/op",
            "value": 102236,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3 - B/op",
            "value": 29752,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkACHFileRead3 - allocs/op",
            "value": 264,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile",
            "value": 14149,
            "unit": "ns/op\t   10760 B/op\t     131 allocs/op",
            "extra": "92554 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile - ns/op",
            "value": 14149,
            "unit": "ns/op",
            "extra": "92554 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile - B/op",
            "value": 10760,
            "unit": "B/op",
            "extra": "92554 times\n4 procs"
          },
          {
            "name": "BenchmarkBuildFile - allocs/op",
            "value": 131,
            "unit": "allocs/op",
            "extra": "92554 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid",
            "value": 25.35,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "46538966 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid - ns/op",
            "value": 25.35,
            "unit": "ns/op",
            "extra": "46538966 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "46538966 times\n4 procs"
          },
          {
            "name": "BenchmarkCalculateCheckDigit/valid - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "46538966 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite",
            "value": 49577,
            "unit": "ns/op\t   34696 B/op\t     226 allocs/op",
            "extra": "23034 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite - ns/op",
            "value": 49577,
            "unit": "ns/op",
            "extra": "23034 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite - B/op",
            "value": 34696,
            "unit": "B/op",
            "extra": "23034 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDWrite - allocs/op",
            "value": 226,
            "unit": "allocs/op",
            "extra": "23034 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite",
            "value": 217164,
            "unit": "ns/op\t   54680 B/op\t    2069 allocs/op",
            "extra": "5785 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite - ns/op",
            "value": 217164,
            "unit": "ns/op",
            "extra": "5785 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite - B/op",
            "value": 54680,
            "unit": "B/op",
            "extra": "5785 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWEBWrite - allocs/op",
            "value": 2069,
            "unit": "allocs/op",
            "extra": "5785 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite",
            "value": 153263,
            "unit": "ns/op\t   61104 B/op\t     721 allocs/op",
            "extra": "7508 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite - ns/op",
            "value": 153263,
            "unit": "ns/op",
            "extra": "7508 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite - B/op",
            "value": 61104,
            "unit": "B/op",
            "extra": "7508 times\n4 procs"
          },
          {
            "name": "BenchmarkIATWrite - allocs/op",
            "value": 721,
            "unit": "allocs/op",
            "extra": "7508 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDIATWrite",
            "value": 231380104,
            "unit": "1210428822610090353A094101Federal Reserve Bank   My Bank Name                   ",
            "extra": "101 times\n4 procs"
          },
          {
            "name": "BenchmarkPPDIATWrite - 1210428822610090353A094101Federal",
            "value": 231380104,
            "unit": "1210428822610090353A094101Federal",
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
          }
        ]
      }
    ]
  }
}