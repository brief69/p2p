package main

import (
	"fmt"
	"os"

	"github.com/multiformats/go-multiaddr"
)

// HandleErr is a utility function to handle errors.
func HandleErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err)
		os.Exit(1)
	}
}

// StringToMultiAddr is a utility function to convert a string to a multiaddress.
func StringToMultiAddr(str string) (multiaddr.Multiaddr, error) {
	maddr, err := multiaddr.NewMultiaddr(str)
	if err != nil {
		return nil, fmt.Errorf("Failed to convert string to multiaddress: %s", err.Error())
	}
	return maddr, nil
}
