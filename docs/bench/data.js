window.BENCHMARK_DATA = {
  "lastUpdate": 1789399613443,
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
      }
    ]
  }
}