package repo

import (
	"errors"
	"os"
	"sync"

	"github.com/lukaszglowacki/repo/pkg/errutil"
	"github.com/spf13/cobra"
)

type SyncOption struct {
	FileName string
}

func NewSyncOptions() SyncOption {
	return SyncOption{
		FileName: "default.xml",
	}
}
func NewCmdSync() *cobra.Command {
	options := NewSyncOptions()

	cmd := &cobra.Command{
		Use: "sync -f FILENAME",
		DisableFlagsInUseLine: true,
		Short: "Initialize repozitories from manifest file",
		Long: `
This command clones (or pulls) all projects from manifest file.
Please visit https://gerrit.googlesource.com/git-repo/+/master/docs/manifest-format.txt to read more about manifest file.`,
		Example: "repo sync -f manifest.xml",
		Run: func(cmd *cobra.Command, args []string) {
			errutil.CheckErrFatal(validateSyncOptions(options))
			errutil.CheckErrFatal(runSync(options))
		},
	}

	cmd.Flags().StringVarP(&options.FileName, "filename", "f", options.FileName, "Selector that contains path to the manifest file to apply. Default value: default.xml")
	return cmd
}

func runSync(options SyncOption) error {
	m, err := NewManifestFromFile(options.FileName)
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	wg.Add(len(m.Project))

	for _, p := range m.Project {
		r := m.Remote.ByName(p.Remote)

		o := r.Fetch + p.Name
		l := "./.repo/" + p.Path

		go func(origin, local string) {
			defer wg.Done()
			downloadRepo(origin, local)
		}(o, l)
	}

	wg.Wait()
	return nil
}

func downloadRepo(origin, path string) {
	git := NewGitCommand(path)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		git.Clone(origin)
		return
	}

	git.Pull()
}

func validateSyncOptions(o SyncOption) error {
	if o.FileName == "" {
		return errors.New("File name can not be empty")
	}
	return nil
}
