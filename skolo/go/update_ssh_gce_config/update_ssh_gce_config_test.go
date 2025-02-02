package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.skia.org/infra/go/executil"
	"go.skia.org/infra/go/util"
)

func TestUpdateSshCfg_Success(t *testing.T) {
	ctx := executil.FakeTestsContext("Test_FakeExe_Gcloud")

	// Write a ssh.cfg file with stale contents.
	sshCfgFile := filepath.Join(t.TempDir(), "ssh.cfg")
	originalSshCfgContents := `Stuff before the autogenerated block.
# BEGIN GCE MACHINES. DO NOT EDIT! This block is automatically generated.
Host should-go-away
	Hostname 1.2.3.4
# END GCE MACHINES.
Stuff after the autogenerated block.
`
	require.NoError(t, util.WithWriteFile(sshCfgFile, func(w io.Writer) error {
		_, err := w.Write([]byte(originalSshCfgContents))
		return err
	}))

	// Update the ssh.cfg file.
	require.NoError(t, updateSshCfg(ctx, sshCfgFile))

	// Read ssh.cfg file and assert that it was correctly updated.
	updatedSshCfgFileBytes, err := ioutil.ReadFile(sshCfgFile)
	require.NoError(t, err)
	expectedSshCfgContents := `Stuff before the autogenerated block.
# BEGIN GCE MACHINES. DO NOT EDIT! This block is automatically generated.
Host skia-e-gce-100
  Hostname 1.1.1.1
Host skia-e-gce-101
  Hostname 2.2.2.2
Host skia-e-gce-102
  Hostname 3.3.3.3
Host skia-e-gce-300
  Hostname 7.7.7.7
Host skia-e-gce-301
  Hostname 8.8.8.8
Host skia-e-gce-302
  Hostname 9.9.9.9
# END GCE MACHINES.
Stuff after the autogenerated block.
`
	assert.Equal(t, expectedSshCfgContents, string(updatedSshCfgFileBytes))
}

func Test_FakeExe_Gcloud(t *testing.T) {
	if !executil.IsCallingFakeCommand() {
		return
	}
	require.Equal(t, []string{
		"gcloud",
		"compute",
		"instances",
		"list",
		"--project=skia-swarming-bots",
		"--format=csv(name, networkInterfaces[0].accessConfigs[0].natIP, disks[0].licenses[0])",
		"--filter=name~skia-e-*",
		"--sort-by=name",
	}, executil.OriginalArgs())
	fmt.Printf(`name,nat_ip,licenses
skia-e-gce-100,1.1.1.1,https://www.googleapis.com/compute/v1/projects/debian-cloud/global/licenses/debian-10-buster
skia-e-gce-101,2.2.2.2,https://www.googleapis.com/compute/v1/projects/debian-cloud/global/licenses/debian-10-buster
skia-e-gce-102,3.3.3.3,https://www.googleapis.com/compute/v1/projects/debian-cloud/global/licenses/debian-10-buster
skia-e-gce-200,4.4.4.4,https://www.googleapis.com/compute/v1/projects/windows-cloud/global/licenses/windows-server-2019-dc
skia-e-gce-201,5.5.5.5,https://www.googleapis.com/compute/v1/projects/windows-cloud/global/licenses/windows-server-2019-dc
skia-e-gce-202,6.6.6.6,https://www.googleapis.com/compute/v1/projects/windows-cloud/global/licenses/windows-server-2019-dc
skia-e-gce-300,7.7.7.7,https://www.googleapis.com/compute/v1/projects/debian-cloud/global/licenses/debian-10-buster
skia-e-gce-301,8.8.8.8,https://www.googleapis.com/compute/v1/projects/debian-cloud/global/licenses/debian-10-buster
skia-e-gce-302,9.9.9.9,https://www.googleapis.com/compute/v1/projects/debian-cloud/global/licenses/debian-10-buster
`)
	os.Exit(0)
}
