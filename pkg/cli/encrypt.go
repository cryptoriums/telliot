// Copyright (c) The Cryptorium Authors.
// Licensed under the MIT License.

package cli

import (
	"context"

	"github.com/cryptoriums/packages/private_file"
	"github.com/go-kit/log"
)

type EncryptCmd struct {
	File string `arg:"" required:"" type:"existingfile" help:"the file to encrypt"`
}

func (self *EncryptCmd) Run(cli *CLI, ctx context.Context, logger log.Logger) error {
	return private_file.EncryptWithPasswordLoop(self.File, self.File+".encrypted")
}
