/*
 * ACH API
 *
 * Moov ACH ([Automated Clearing House](https://en.wikipedia.org/wiki/Automated_Clearing_House)) implements an HTTP API for creating, parsing and validating ACH files. ACH is the primary method of electronic money movement throughout the United States.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client

// File struct for File
type File struct {
	// File ID
	ID          string      `json:"ID,omitempty"`
	FileHeader  FileHeader  `json:"fileHeader,omitempty"`
	Batches     []Batch     `json:"batches,omitempty"`
	IATBatches  []IatBatch  `json:"IATBatches,omitempty"`
	FileControl FileControl `json:"fileControl,omitempty"`
}