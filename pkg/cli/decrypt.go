// Copyright (c) The Cryptorium Authors.
// Licensed under the MIT License.

package cli

import (
	"context"
	"io/ioutil"
	"os"

	"github.com/cryptoriums/telliot/pkg/private_file"
	"github.com/cryptoriums/telliot/pkg/prompt"
	"github.com/go-kit/kit/log"
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
	decrypted := private_file.DecryptWithPasswordLoop(in)

	filename := prompt.PromptFileName()

	return os.WriteFile(filename, decrypted, 0666)
}