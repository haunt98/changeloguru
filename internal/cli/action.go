package cli

import (
	"fmt"
	"log"

	"github.com/make-go-great/buildinfo-go"
	"github.com/urfave/cli/v2"
)

const (
	currentDir       = "."
	markdownFiletype = "md"
	rstFiletype      = "rst"

	defaultRepository = currentDir
	defaultOutput     = currentDir
	defaultFilename   = "CHANGELOG"
	defaultFiletype   = markdownFiletype
)

type action struct {
	flags struct {
		verbose     bool
		from        string
		to          string
		version     string
		scopes      map[string]struct{}
		repository  string
		output      string
		filename    string
		filetype    string
		dryRun      bool
		interactive bool
	}
}

func (a *action) RunVersion(c *cli.Context) error {
	info, ok := buildinfo.Read()
	if !ok {
		return nil
	}

	fmt.Printf("%s-%s\n", info.MainModuleVersion, info.GitCommit)
	return nil
}

func (a *action) RunHelp(c *cli.Context) error {
	return cli.ShowAppHelp(c)
}

func (a *action) getFlags(c *cli.Context) {
	a.flags.verbose = c.Bool(flagVerbose)
	a.flags.version = c.String(flagVersion)
	a.flags.from = c.String(flagFrom)
	a.flags.to = c.String(flagTo)

	a.flags.scopes = make(map[string]struct{})
	for _, scope := range c.StringSlice(flagScope) {
		a.flags.scopes[scope] = struct{}{}
	}

	a.flags.repository = c.String(flagRepository)
	a.flags.output = c.String(flagOutput)
	a.flags.filename = c.String(flagFilename)
	a.flags.filetype = c.String(flagFiletype)
	a.flags.dryRun = c.Bool(flagDryRun)
	a.flags.interactive = c.Bool(flagInteractive)

	a.log("flags %+v", a.flags)
}

func (a *action) log(format string, v ...interface{}) {
	if a.flags.verbose {
		log.Printf(format, v...)
	}
}
