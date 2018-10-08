package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/moov-io/ach"
	"log"
	"os"
	"path/filepath"
	"runtime/pprof"
	"time"
)

var (
	fPath      = flag.String("fPath", "", "File Path")
	cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

	// output formats
	flagJson = flag.Bool("json", false, "Output file in json")
)

// main creates an ACH File with 4 batches of SEC Code PPD.
// Each batch contains an EntryAddendaCount of 2500.
func main() {
	flag.Parse()

	filename := time.Now().UTC().Format("200601021504")
	if *flagJson {
		filename += ".json"
	} else {
		filename += ".ach"
	}

	path := filepath.Join(*fPath, filename)
	write(path)
}

func write(path string) {
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	f, err := os.Create(path)
	if err != nil {
		fmt.Printf("%T: %s", err, err)
	}

	// To create a file
	fh := ach.NewFileHeader()
	fh.ImmediateDestination = "231380104"
	fh.ImmediateOrigin = "121042882"
	fh.FileCreationDate = time.Now()
	fh.ImmediateDestinationName = "Citadel"
	fh.ImmediateOriginName = "Wells Fargo"
	file := ach.NewFile()
	file.SetHeader(fh)

	// Create 4 Batches of SEC Code PPD
	for i := 0; i < 4; i++ {
		bh := ach.NewBatchHeader()
		bh.ServiceClassCode = 200
		bh.CompanyName = "Wells Fargo"
		bh.CompanyIdentification = "121042882"
		bh.StandardEntryClassCode = "PPD"
		bh.CompanyEntryDescription = "Trans. Description"
		bh.EffectiveEntryDate = time.Now().AddDate(0, 0, 1)
		bh.ODFIIdentification = "121042882"

		batch, _ := ach.NewBatch(bh)

		// Create Entry
		entrySeq := 0
		for i := 0; i < 1250; i++ {
			entrySeq = entrySeq + 1

			entryEntrySeq := ach.NewEntryDetail()
			entryEntrySeq.TransactionCode = 22
			entryEntrySeq.SetRDFI("231380104")
			entryEntrySeq.DFIAccountNumber = "81967038518"
			entryEntrySeq.Amount = 100000
			//entryEntrySeq.IndividualName = randomdata.FullName(randomdata.RandomGender)
			entryEntrySeq.IndividualName = "Steven Tander"
			entryEntrySeq.SetTraceNumber(bh.ODFIIdentification, entrySeq)
			//entryEntrySeq.IdentificationNumber = "#" + randomdata.RandStringRunes(13) + "#"
			entryEntrySeq.IdentificationNumber = "#83738AB#"
			entryEntrySeq.Category = ach.CategoryForward

			// Add addenda record for an entry
			addendaEntrySeq := ach.NewAddenda05()
			addendaEntrySeq.PaymentRelatedInformation = "bonus pay for amazing work on #OSS"
			entryEntrySeq.AddAddenda(addendaEntrySeq)

			// Add entries
			batch.AddEntry(entryEntrySeq)

		}

		// Create the batch.
		if err := batch.Create(); err != nil {
			fmt.Printf("%T: %s", err, err)
		}

		// Add batch to the file
		file.AddBatch(batch)
	}

	// Create the file
	if err := file.Create(); err != nil {
		fmt.Printf("%T: %s", err, err)
	}

	// Write to a file
	if *flagJson {
		// Write in JSON format
		if err := json.NewEncoder(f).Encode(file); err != nil {
			fmt.Printf("%T: %s", err, err)
		}
	} else {
		// Write in ACH plain text format
		w := ach.NewWriter(f)
		if err := w.Write(file); err != nil {
			fmt.Printf("%T: %s", err, err)
		}
		w.Flush()
	}

	if err := f.Close(); err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("Wrote %s\n", path)
}
