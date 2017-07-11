package main

import (
	"os"

	"github.com/consensus/cmd/consensus/commands"
	"github.com/tendermint/tmlibs/cli"
)

func main() {
	cmd := cli.PrepareBaseCmd(commands.RootCmd, "TM", os.ExpandEnv("$HOME/.consensus"))
	cmd.Execute()
}
