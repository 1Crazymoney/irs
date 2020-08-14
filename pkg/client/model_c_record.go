/*
 * IRS API
 *
 * Package github.com/moov-io/irs implements a file reader and writer written in Go along with a HTTP API and CLI for creating, parsing, validating, and transforming IRS electronic Filing Information Returns Electronically (FIRE). FIRE operates on a byte(ASCII) level making it difficult to interface with JSON and CSV/TEXT file formats. | Input      | Output     | |------------|------------| | JSON       | JSON       | | ASCII FIRE | ASCII FIRE | |            | PDF Form   | |            | SQL        |
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client

// CRecord struct for CRecord
type CRecord struct {
	RecordType           string `json:"record_type"`
	NumberOfPayees       int32  `json:"number_of_payees,omitempty"`
	ControlTotal1        int32  `json:"control_total_1,omitempty"`
	ControlTotal2        int32  `json:"control_total_2,omitempty"`
	ControlTotal3        int32  `json:"control_total_3,omitempty"`
	ControlTotal4        int32  `json:"control_total_4,omitempty"`
	ControlTotal5        int32  `json:"control_total_5,omitempty"`
	ControlTotal6        int32  `json:"control_total_6,omitempty"`
	ControlTotal7        int32  `json:"control_total_7,omitempty"`
	ControlTotal8        int32  `json:"control_total_8,omitempty"`
	ControlTotal9        int32  `json:"control_total_9,omitempty"`
	ControlTotalA        int32  `json:"control_total_A,omitempty"`
	ControlTotalB        int32  `json:"control_total_B,omitempty"`
	ControlTotalC        int32  `json:"control_total_C,omitempty"`
	ControlTotalD        int32  `json:"control_total_D,omitempty"`
	ControlTotalE        int32  `json:"control_total_E,omitempty"`
	ControlTotalF        int32  `json:"control_total_F,omitempty"`
	ControlTotalG        int32  `json:"control_total_G,omitempty"`
	RecordSequenceNumber int32  `json:"record_sequence_number"`
}
