package upgrade

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/pkg/errors"
)

func StartUpgrade() error {
	hasNextVersion, release, err := NewVersionAvailable()
	if err != nil {
		return errors.Wrap(err, "fail to get new version")
	}

	if hasNextVersion {
		return fmt.Errorf("No new version available")
	}

	uid := os.Getuid()
	if uid != 0 {
		return fmt.Errorf("This script should be run as root")
	}

	archB, err := exec.Command("uname", "-m").Output()
	if err != nil {
		return errors.Wrap(err, "fail to find architecture")
	}

	arch := string(archB)

	if arch == "x86_64" {
		arch = "amd64"
	}

	if strings.HasPrefix(arch, "arm") {
		arch = "arm"
	}

	downloadURL := fmt.Sprintf("https://github.com/johnsudaar/acp/releases/download/%s/acp-%s-linux-%s.zip", release.TagName, release.TagName, arch)

	dir, err := ioutil.TempDir("", "acp-install")
	if err != nil {
		return errors.Wrap(err, "fail to create temp dir for install")
	}

	resp, err := http.Get(downloadURL)
	if err != nil {
		return errors.Wrap(err, "fail to download new release")
	}
	defer resp.Body.Close()

	if resp.StatusCode/100 != 2 {
		return fmt.Errorf("fail to download new release, status: %s", resp.Status)
	}

	out, err := os.Create(fmt.Sprintf("%s/archive.zip", dir))
	if err != nil {
		return errors.Wrap(err, "fail to create the archive file")
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return errors.Wrap(err, "fail to download file")
	}

	err = exec.Command("unzip", "archive.zip", "-d", dir).Run()
	if err != nil {
		return errors.Wrap(err, "fail to unzip file")
	}

	execBin := fmt.Sprintf("%s/acp-%s-linux-%s/acp", dir, release.TagName, arch)

	cmd := exec.Command(execBin, "upgrade", "continue", fmt.Sprintf("%s/acp-%s-linux-%s/", dir, release.TagName, arch))
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return errors.Wrap(err, "fail to get command output (stdout)")
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return errors.Wrap(err, "fail to get command output (stderr)")
	}

	go io.Copy(os.Stdout, stdout)
	go io.Copy(os.Stderr, stderr)

	err = cmd.Wait()
	if err != nil {
		return errors.Wrap(err, "fail to run upgrade")
	}

	return nil
}

func DoUpgrade(path string) error {
	fmt.Println("-> Creating directories")
	err := os.MkdirAll("/var/lib/acp", 0755)
	if err != nil {
		return errors.Wrap(err, "fail to create acp directory")
	}
	return nil
}
