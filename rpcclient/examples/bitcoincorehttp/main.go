// Copyright (c) 2014-2017 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package main

import (
	"log"

	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/rpcclient"
)

func main() {
	// Connect to local bitcoin core RPC server using HTTP POST mode.
	connCfg := &rpcclient.ConnConfig{
		Host:         "localhost:9332",
		User:         "lekai",
		Pass:         "btccom",
		HTTPPostMode: true, // Bitcoin core only supports HTTP POST mode
		DisableTLS:   true, // Bitcoin core does not provide TLS by default
	}
	// Notice the notification parameter is nil since notifications are
	// not supported in HTTP POST mode.
	client, err := rpcclient.New(connCfg, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Shutdown()

	// Get the current block count.
	blockHash, _ := chainhash.NewHashFromStr("e8f29c11145bab2729fe4ea3cf2b44a64adef6815b8fb15c4795ae6b331eeed2")
	rest, err := client.GetBlockVerboseTx(blockHash)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Block count: %v", rest)
}
