// Copyright (c) The Cryptorium Authors.
// Licensed under the MIT License.

package cli

import (
	"github.com/cryptoriums/telliot/pkg/private_file"
)

type EncryptCmd struct {
	File string `arg:"" required:"" type:"existingfile" help:"the file to encrypt"`
}

func (self *EncryptCmd) Run() error {
	return private_file.EncryptWithPasswordLoop(self.File, self.File+".encrypted")
}
