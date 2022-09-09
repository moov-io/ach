package server

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/moov-io/ach"
	"github.com/moov-io/base/log"

	kitlog "github.com/go-kit/log"
)

const (
	numTransactions = 75000
	numBatches      = 1000
)

// Benchmark creating a large ACH file through the Go library
func BenchmarkCreateBigFile__Library(b *testing.B) {
	for i := 0; i < b.N; i++ {
		createBigFile(b, numTransactions, numBatches)
	}
}

// Benchmark creating a large ACH file through the HTTP server
func BenchmarkCreateBigFile__Server(b *testing.B) {
	logger := log.NewNopLogger()
	repo := NewRepositoryInMemory(24*time.Hour, logger)
	service := NewService(repo)
	// setup our HTTP handler
	handler := MakeHTTPHandler(service, repo, kitlog.NewNopLogger())
	// Spin up a local HTTP server
	server := httptest.NewServer(handler)
	defer server.Close()

	// Create a file with 75K entries and convert it to a NACHA file
	bb := &bytes.Buffer{}
	writer := ach.NewWriter(bb)
	f := createBigFile(b, numTransactions, numBatches)
	if err := writer.Write(f); err != nil {
		b.Fatal(err)
	}

	// Make the request
	for i := 0; i < b.N; i++ {
		req, err := http.NewRequest("POST", server.URL+"/files/create", bytes.NewReader(bb.Bytes()))
		if err != nil {
			b.Fatal(err)
		}
		req.Header.Set("Content-Type", "text/plain")
		resp, err := server.Client().Do(req)
		if err != nil {
			b.Fatal(err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != 200 {
			b.Fatalf("non-200 response code: %v", resp.StatusCode)
		}
	}
}

// Benchmark for flattening an ACH file
func Benchmark__FlattenBigFile(b *testing.B) {
	logger := log.NewNopLogger()
	repo := NewRepositoryInMemory(24*time.Hour, logger)
	service := NewService(repo)
	// setup our HTTP handler
	handler := MakeHTTPHandler(service, repo, kitlog.NewNopLogger())
	// Spin up a local HTTP server
	server := httptest.NewServer(handler)
	defer server.Close()

	// Create a file with 75K entries and convert it to a NACHA file
	bb := &bytes.Buffer{}
	writer := ach.NewWriter(bb)
	f := createBigFile(b, numTransactions, numBatches)
	if err := writer.Write(f); err != nil {
		b.Fatal(err)
	}

	if err := repo.StoreFile(f); err != nil {
		b.Fatal(err)
	}

	// Make the request
	url := fmt.Sprintf("%s/files/%s/flatten", server.URL, f.ID)
	for i := 0; i < b.N; i++ {
		req, err := http.NewRequest("POST", url, nil)
		if err != nil {
			b.Fatal(err)
		}
		req.Header.Set("Content-Type", "text/plain")
		resp, err := server.Client().Do(req)
		if err != nil {
			b.Fatal(err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != 200 {
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				b.Fatal(err)
			}

			b.Log(string(body))
			b.Fatalf("non-200 response code: %v", resp.StatusCode)

		}
	}
}

func createBigFile(b *testing.B, numTransactions int, numBatches int) *ach.File {
	fmt.Printf("Benchmarks for ACH file with %d transactions split up in %d batches\n", numTransactions, numBatches)

	f := ach.NewFile()
	f.ID = "foo"
	f.Header = *mockFileHeader()

	if numTransactions < numBatches {
		b.Fatal("numTransactions must be greater than numBatches")
	}

	if (numTransactions % numBatches) != 0 {
		b.Fatal("number of transactions must be evenly distributed across batches")
	}

	entriesPerBatch := numTransactions / numBatches

	traceNumSeq := 0
	// Create and add the batch
	for i := 0; i < numBatches; i++ {
		// Create entries and add to current batch
		batch := ach.NewBatchWEB(mockBatchHeaderWeb())

		for j := 0; j < entriesPerBatch; j++ {
			entry := ach.NewEntryDetail()
			entry.ID = "98765"
			entry.TransactionCode = ach.CheckingCredit
			entry.SetRDFI("231380104")
			entry.DFIAccountNumber = "123456789"
			entry.Amount = 1
			entry.IndividualName = "Wade Arnold"

			traceNumSeq += 1
			entry.SetTraceNumber(batch.Header.ODFIIdentification, traceNumSeq)
			//entry.TraceNumber = fmt.Sprintf("%d", i+j)
			entry.SetPaymentType("S")

			batch.AddEntry(entry)
		}

		if err := batch.Create(); err != nil {
			b.Fatal(err)
		}
		f.AddBatch(batch)
	}

	if err := f.Create(); err != nil {
		b.Fatal(err)
	}
	if err := f.Validate(); err != nil {
		b.Fatal(err)
	}

	return f
}
