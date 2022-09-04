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
		scopes        map[string]struct{}
		output        string
		from          string
		to            string
		version       string
		repository    string
		filename      string
		filetype      string
		verbose       bool
		dryRun        bool
		interactive   bool
		autoGitCommit bool
		autoGitTag    bool
		autoGitPush   bool
	}
}

func (a *action) RunHelp(c *cli.Context) error {
	return cli.ShowAppHelp(c)
}

func (a *action) getFlags(c *cli.Context) {
	a.flags.verbose = c.Bool(flagVerboseName)
	a.flags.version = c.String(flagVersionName)
	a.flags.from = c.String(flagFromName)
	a.flags.to = c.String(flagToName)

	a.flags.scopes = make(map[string]struct{})
	for _, scope := range c.StringSlice(flagScopeName) {
		a.flags.scopes[scope] = struct{}{}
	}

	a.flags.repository = c.String(flagRepositoryName)
	if a.flags.repository == "" {
		a.log("Fallback to default repository [%s]", defaultRepository)
		a.flags.repository = defaultRepository
	}

	a.flags.output = c.String(flagOutputName)
	if a.flags.output == "" {
		a.log("Fallback to default output [%s]\n", defaultOutput)
		a.flags.output = defaultOutput
	}

	a.flags.filename = c.String(flagFilenameName)
	if a.flags.filename == "" {
		a.log("Fallback to default filename [%s]\n", defaultFilename)
		a.flags.filename = defaultFilename
	}

	a.flags.filetype = c.String(flagFiletypeName)
	if a.flags.filetype == "" {
		a.log("Fallback to default filetype [%s]\n", defaultFiletype)
		a.flags.filetype = defaultFiletype
	}

	a.flags.dryRun = c.Bool(flagDryRunName)
	a.flags.interactive = c.Bool(flagInteractiveName)
	a.flags.autoGitCommit = c.Bool(flagAutoGitCommitName)
	a.flags.autoGitTag = c.Bool(flagAutoGitTagName)
	a.flags.autoGitPush = c.Bool(flagAutoGitPushName)

	a.log("Flags %+v\n", a.flags)
}

func (a *action) log(format string, v ...interface{}) {
	if a.flags.verbose {
		log.Printf(format, v...)
	}
}
