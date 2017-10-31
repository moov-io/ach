package ach

// Addendumer abstracts the different ACH addendum types that can be added to an EntryDetail record
type Addendumer interface {
	Parse(string)
	//TypeCode Defines the specific explanation and format for the addenda information
	TypeCode() string
	String() string
	Validate() error
}
