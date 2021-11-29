// Copyright (c) The Cryptorium Authors.
// Licensed under the MIT License.

package cli

import (
	"context"
	"io/ioutil"
	"os"

	"github.com/cryptoriums/packages/private_file"
	"github.com/cryptoriums/packages/prompt"
	"github.com/go-kit/log"
	"github.com/pkg/errors"
)

type DecryptCmd struct {
	File string `arg:"" required:"" type:"existingfile" help:"the file to encrypt"`
}

func (self *DecryptCmd) Run(cli *CLI, ctx context.Context, logger log.Logger) error {
	in, err := ioutil.ReadFile(self.File)
	if err != nil {
		return errors.Wrap(err, "reading input file")
	}

	decrypted, err := private_file.DecryptWithPasswordLoop(in)
	if err != nil {
		return errors.Wrap(err, "decrypt input file")
	}

	filename := prompt.PromptFileName()

	return os.WriteFile(filename, decrypted, 0666)
}
