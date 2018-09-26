package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric/protos/common"
	"github.com/syndtr/goleveldb/leveldb"
)

var (
	underConstructionLedgerKey = []byte("underConstructionLedgerKey")

	path = flag.String("path", "/var/hyperledger/production/ledgersData/ledgerProvider/", "leveldb path dir")
)

func main() {
	flag.Parse()

	db, err := leveldb.OpenFile(*path, nil)
	if err != nil {
		fmt.Printf("Open failed, err: %v", err)
		os.Exit(-1)
	}
	defer db.Close()

	iter := db.NewIterator(nil, nil)
	for iter.Next() {
		key := iter.Key()
		value := iter.Value()
		if err := showKeyValue(key, value); err != nil {
			return
		}
	}
	iter.Release()
	err = iter.Error()
	if err != nil {
		fmt.Printf("\n iter err: %v", err)
	}

	fmt.Println("\n game is over")
}

func showKeyValue(key, value []byte) error {
	fmt.Printf("\n\n ----- \n key: %v", string(key))

	if bytes.Equal(key, underConstructionLedgerKey) {
		fmt.Printf("\n value: %v", string(value))
		return nil
	}

	if bytes.HasPrefix(key, []byte("l")) {
		b := &common.Block{}

		if err := proto.Unmarshal(value, b); err != nil {
			fmt.Printf("\n proto.Unmarshal err: %v", err)
			return err
		}
		spew.Printf("\n common.Block:: %#+v  \n\n\n", b)

		return nil
	}

	fmt.Printf("\n unreachable code area")

	fmt.Printf("\n key: %v, value: %v", key, value)

	fmt.Printf("\n unreachable code area")

	return nil
}
