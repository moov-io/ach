package fix

import (
	"github.com/moov-io/ach"
)

func updateEED(conf Config) batchHeaderFixer {
	return func(bh *ach.BatchHeader) error {
		bh.EffectiveEntryDate = conf.UpdateEED
		return nil
	}
}
