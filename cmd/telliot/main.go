// Copyright (c) The Cryptorium Authors.
// Licensed under the MIT License.

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/alecthomas/kong"
	"github.com/cryptoriums/telliot/pkg/cli"
	"github.com/cryptoriums/telliot/pkg/github"
	"github.com/cryptoriums/telliot/pkg/logging"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
)

var GitTag string
var GitHash string

func main() {
	l := logging.NewLogger()
	c := context.Background()
	// Don't show the version message when it's an help command.
	shouldShowVersionMessage := true
	for _, arg := range os.Args {
		if arg == "--help" || arg == "-h" {
			shouldShowVersionMessage = false
			break
		}
	}
	if shouldShowVersionMessage {
		//lint:ignore faillint looks cleaner with print instead of logs
		fmt.Printf(cli.VersionMessage, GitTag, GitHash)

		newRelease, err := github.CheckNewVersion("cryptoriums", "telliot", GitTag)
		if err != nil {
			level.Error(l).Log("msg", "checking for new release", "err", err)
		}
		if newRelease != "" {
			level.Error(l).Log("msg", "THERE IS A NEW RELEASE", "version", newRelease)
		}
	}

	ctx := kong.Parse(&cli.CLIDefault, kong.BindTo(c, (*context.Context)(nil)), kong.BindTo(l, (*log.Logger)(nil)), kong.Name("telliot"),
		kong.Description("The unofficial Tellor cli tool"),
		kong.UsageOnError())

	ctx.FatalIfErrorf(ctx.Run(*ctx))
}
