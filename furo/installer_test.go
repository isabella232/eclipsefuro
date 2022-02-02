package main_test

import (
	"github.com/eclipse/eclipsefuro/furo/cmd"
	"github.com/eclipse/eclipsefuro/furo/test"
	"github.com/stretchr/testify/require"
	"os"
	"path"
	"testing"
)

func TestInstallCommand(t *testing.T) {
	dir, _ := test.CwdTestDir()
	defer test.RemoveTestDir(dir)

	os.Args = []string{"cmd", "init", "--repository=github.com/yourname/furo-test"}
	rco := cmd.RootCmd
	rco.Execute()

	require.Equal(t, true, test.FileExist(path.Join(dir, ".furo")))
	require.Equal(t, "3c90654659f1898586736e290c25dd5f", test.MustMd5Sum(path.Join(dir, ".furo")))

	os.Args = []string{"cmd", "install"}
	rco.Execute()

	require.Equal(t, true, test.FileExist(path.Join(dir, "dependencies")))

}
