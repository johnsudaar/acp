package upgrade

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/johnsudaar/acp/config"
	"github.com/pkg/errors"
)

const ReleaseURL = "https://api.github.com/repos/johnsudaar/acp/releases/latest"

type Release struct {
	ID      uint64  `json:"id"`
	Name    string  `json:"name"`
	Body    string  `json:"body"`
	TagName string  `json:"tag_name"`
	Assets  []Asset `json:"assets"`
}

type Asset struct {
	ID                 int64  `json:"id"`
	Name               string `json:"name"`
	BrowserDownloadURL string `json:"browser_download_url"`
}

func GetLatestRelease() (Release, error) {
	var release Release

	resp, err := http.Get(ReleaseURL)
	if err != nil {
		return release, errors.Wrap(err, "fail to get releases")
	}
	defer resp.Body.Close()

	if resp.StatusCode/100 != 2 {
		return release, fmt.Errorf("invalid response code: %s", resp.Status)
	}

	err = json.NewDecoder(resp.Body).Decode(&release)
	if err != nil {
		return release, errors.Wrap(err, "fail to parse response body")
	}
	return release, nil
}

func NewVersionAvailable() (bool, Release, error) {
	r, err := GetLatestRelease()
	if err != nil {
		return false, r, errors.Wrap(err, "fail to list versions")
	}

	return r.TagName != config.Get().Version, r, nil
}
