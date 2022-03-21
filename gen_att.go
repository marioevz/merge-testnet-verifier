// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package main

import (
	"encoding/json"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

var _ = (*AttestationMarshaling)(nil)

// MarshalJSON marshals as JSON.
func (a Attestation) MarshalJSON() ([]byte, error) {
	type Attestation struct {
		AggregationBits hexutil.Uint64  `json:"aggregation_bits"`
		Data            AttestationData `json:"data"`
		Signature       string          `json:"signature"`
	}
	var enc Attestation
	enc.AggregationBits = hexutil.Uint64(a.AggregationBits)
	enc.Data = a.Data
	enc.Signature = a.Signature
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (a *Attestation) UnmarshalJSON(input []byte) error {
	type Attestation struct {
		AggregationBits *hexutil.Uint64  `json:"aggregation_bits"`
		Data            *AttestationData `json:"data"`
		Signature       *string          `json:"signature"`
	}
	var dec Attestation
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.AggregationBits != nil {
		a.AggregationBits = uint64(*dec.AggregationBits)
	}
	if dec.Data != nil {
		a.Data = *dec.Data
	}
	if dec.Signature != nil {
		a.Signature = *dec.Signature
	}
	return nil
}
