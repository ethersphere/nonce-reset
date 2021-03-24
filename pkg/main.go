package main

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/ethersphere/bee/pkg/logging"
	"github.com/ethersphere/bee/pkg/statestore/leveldb"
)

func main() {
	logger := logging.New(os.Stdout, 6)
	if len(os.Args) != 2 {
		logger.Error("datadir needs to be specified as first argument")
		return
	}

	if err := fix(logger, os.Args[1]); err != nil {
		logger.Errorf("%v", err)
	}
}

func fix(logger logging.Logger, datadir string) error {
	if _, err := os.Stat(filepath.Join(datadir, "statestore")); os.IsNotExist(err) {
		return errors.New("not a bee data directory")
	}

	store, err := leveldb.NewStateStore(filepath.Join(datadir, "statestore"), logger)
	if err != nil {
		return err
	}
	defer store.Close()
	return store.Iterate("transaction_nonce_", func(key []byte, value []byte) (bool, error) {
		logger.Infof("removing key %s", string(key))
		return false, store.Delete(string(key))
	})
}
