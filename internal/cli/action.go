package cli

import (
	"log"

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
		scopes      map[string]struct{}
		output      string
		from        string
		to          string
		version     string
		repository  string
		filename    string
		filetype    string
		verbose     bool
		dryRun      bool
		interactive bool
		autoCommit  bool
	}
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
	a.flags.autoCommit = c.Bool(flagAutoCommit)

	a.log("flags %+v", a.flags)
}

func (a *action) log(format string, v ...interface{}) {
	if a.flags.verbose {
		log.Printf(format, v...)
	}
}
