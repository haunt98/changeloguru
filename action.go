package main

import (
	"log"

	"github.com/urfave/cli/v2"
)

const (
	currentDir       = "."
	markdownFiletype = "md"

	defaultRepository = currentDir
	defaultOutput     = currentDir
	defaultFilename   = "CHANGELOG"
	defaultFiletype   = markdownFiletype
)

type action struct {
	flags struct {
		verbose    bool
		from       string
		to         string
		version    string
		scopes     map[string]struct{}
		repository string
		output     string
		filename   string
		filetype   string
	}
}

func (a *action) runHelp(c *cli.Context) error {
	return cli.ShowAppHelp(c)
}

func (a *action) getFlags(c *cli.Context) {
	a.flags.verbose = c.Bool(verboseFlag)
	a.flags.from = c.String(fromFlag)
	a.flags.to = c.String(toFlag)
	a.flags.version = c.String(versionFlag)

	a.flags.scopes = make(map[string]struct{})
	for _, scope := range c.StringSlice(scopeFlag) {
		a.flags.scopes[scope] = struct{}{}
	}

	a.flags.repository = a.getFlagValue(c, repositoryFlag, defaultRepository)
	a.flags.output = a.getFlagValue(c, outputFlag, defaultOutput)
	a.flags.filename = a.getFlagValue(c, filenameFlag, defaultFilename)
	a.flags.filetype = a.getFlagValue(c, filetypeFlag, defaultFiletype)

	if a.flags.verbose {
		a.log("flags %+v", a.flags)
	}
}

func (a *action) getFlagValue(c *cli.Context, flag, fallback string) string {
	value := c.String(flag)
	if value == "" {
		value = fallback
	}

	return value
}

func (a *action) log(format string, v ...interface{}) {
	if a.flags.verbose {
		log.Printf(format, v...)
	}
}
