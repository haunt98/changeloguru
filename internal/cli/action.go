package cli

import (
	"cmp"
	"context"
	"log"

	"github.com/urfave/cli/v3"
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
		output          string
		fromRef         string
		toRef           string
		version         string
		repository      string
		filename        string
		filetype        string
		verbose         bool
		dryRun          bool
		interactive     bool
		interactiveFrom bool
		interactiveTo   bool
		autoGitCommit   bool
		autoGitTag      bool
		autoGitPush     bool
	}
}

func (a *action) RunHelp(ctx context.Context, c *cli.Command) error {
	return cli.ShowAppHelp(c)
}

func (a *action) getFlags(c *cli.Command) {
	a.flags.verbose = c.Bool(flagVerboseName)
	a.flags.version = c.String(flagVersionName)
	a.flags.fromRef = cmp.Or(c.String(flagFromName), c.String(flagFromReferenceName))
	a.flags.toRef = cmp.Or(c.String(flagToName), c.String(flagToReferenceName))

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
	a.flags.interactiveFrom = c.Bool(flagInteractiveFromName)
	a.flags.interactiveTo = c.Bool(flagInteractiveToName)
	a.flags.autoGitCommit = c.Bool(flagAutoGitCommitName)
	a.flags.autoGitTag = c.Bool(flagAutoGitTagName)
	a.flags.autoGitPush = c.Bool(flagAutoGitPushName)

	a.log("Flags %+v\n", a.flags)
}

func (a *action) log(format string, v ...any) {
	if a.flags.verbose {
		log.Printf(format, v...)
	}
}
