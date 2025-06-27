package mask

import (
	"bytes"
	"strings"
	"unicode/utf8"

	"github.com/moov-io/ach"
)

func Number(s string) string {
	length := utf8.RuneCountInString(s)
	if length < 5 {
		return strings.Repeat("*", 5) // too short, we can't show anything
	}

	out := bytes.Repeat([]byte("*"), length)
	var unmaskedDigits int
	// Since we want the right-most digits unmasked start from the end of our string
	for i := length - 1; i >= 2; i-- {
		r := rune(s[i])
		if r == ' ' {
			// If the char to our right is masked then mask this left-aligned space as well.
			if i+1 < length && out[i+1] == '*' {
				out[i] = byte('*')
			} else {
				out[i] = byte(' ')
			}
		} else {
			if unmaskedDigits < 4 {
				unmaskedDigits += 1
				out[i] = byte(r)
			}
		}
	}
	return string(out)
}

func Name(s string) string {
	words := strings.Fields(s)

	var out []string
	for i := range words {
		length := utf8.RuneCountInString(words[i])
		if length > 3 {
			out = append(out, words[i][0:2]+strings.Repeat("*", length-2))
		} else {
			out = append(out, strings.Repeat("*", length))
		}
	}
	return strings.Join(out, " ")
}

type Options struct {
	MaskNames          bool
	MaskAccountNumbers bool
	MaskCorrectedData  bool
	MaskIdentification bool
}

func File(file *ach.File, options Options) *ach.File {
	out := ach.NewFile()
	out.Header = file.Header
	out.Control = file.Control

	for b := range file.Batches {
		batch, _ := ach.NewBatch(file.Batches[b].GetHeader())
		if batch == nil {
			continue
		}
		batch.SetControl(file.Batches[b].GetControl())

		entries := file.Batches[b].GetEntries()
		for e := range entries {
			if options.MaskAccountNumbers {
				entries[e].DFIAccountNumber = Number(entries[e].DFIAccountNumberField())
			}
			if options.MaskNames {
				entries[e].IndividualName = Name(entries[e].IndividualNameField())
			}
			if options.MaskIdentification {
				entries[e].IdentificationNumber = Number(entries[e].IdentificationNumberField())
			}

			// Mask some addenda records
			for a := range entries[e].Addenda05 {
				switch file.Batches[b].(type) {
				case *ach.BatchENR:
					paymentInfo, _ := ach.ParseENRPaymentInformation(entries[e].Addenda05[a])
					if paymentInfo != nil {
						if options.MaskNames {
							paymentInfo.IndividualName = Name(paymentInfo.IndividualName)
						}
						if options.MaskAccountNumbers {
							paymentInfo.IndividualIdentification = Number(paymentInfo.IndividualIdentification)
							paymentInfo.DFIAccountNumber = Number(paymentInfo.DFIAccountNumber)
						}

						entries[e].Addenda05[a].PaymentRelatedInformation = paymentInfo.String()
					}

				case *ach.BatchDNE:
					paymentInfo, _ := ach.ParseDNEPaymentInformation(entries[e].Addenda05[a])
					if paymentInfo != nil {
						if options.MaskNames || options.MaskAccountNumbers {
							paymentInfo.CustomerSSN = Number(paymentInfo.CustomerSSN)
						}

						entries[e].Addenda05[a].PaymentRelatedInformation = paymentInfo.String()
					}
				}
			}

			batch.AddEntry(entries[e])
		}

		out.AddBatch(batch)
	}

	for b := range file.IATBatches {
		batch := ach.NewIATBatch(file.IATBatches[b].Header)
		batch.Control = file.IATBatches[b].Control

		entries := file.IATBatches[b].GetEntries()

		for e := range entries {
			if options.MaskAccountNumbers {
				entries[e].DFIAccountNumber = Number(entries[e].DFIAccountNumberField())
			}

			batch.AddEntry(entries[e])
		}

		out.AddIATBatch(batch)
	}

	return out
}
