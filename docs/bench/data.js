window.BENCHMARK_DATA = {
  "lastUpdate": 1789613314046,
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
      }
    ]
  }
}