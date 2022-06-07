package urls

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.skia.org/infra/go/testutils/unittest"
)

const (
	version = "123456"
)

func TestDownloadURLs_UnknownOSOrArch_ReturnsError(t *testing.T) {
	unittest.SmallTest(t)
	_, err := NewDownloadURLs("unknown", "unknown")
	require.Error(t, err)
}

func TestDownloadURLs_ValieOSAndArch_ReturnsValidDownload(t *testing.T) {
	unittest.SmallTest(t)
	urls, err := NewDownloadURLs("darwin", "arm64")
	require.NoError(t, err)
	require.Equal(t, "https://chromedriver.storage.googleapis.com/LATEST_RELEASE", urls.LatestURL())
	require.Equal(t, "https://commondatastorage.googleapis.com/chromium-browser-snapshots/Mac_Arm/LAST_CHANGE", urls.LatestCanaryURL())
	require.Equal(t, "https://chromedriver.storage.googleapis.com/some-version-number/chromedriver_mac64_m1.zip", urls.DriverURL("some-version-number"))
	require.Equal(t, "https://commondatastorage.googleapis.com/chromium-browser-snapshots/Mac_Arm/some-version-number/chromedriver_mac64_m1.zip", urls.CanaryDriverURL("some-version-number"))
}

func TestGetVerionFromURL_GoodHTTPResponse_ReturnsVersionNumber(t *testing.T) {
	unittest.SmallTest(t)
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintln(w, version)
		assert.NoError(t, err)
	}))
	defer ts.Close()

	got, err := GetVersionFromURL(ts.URL, ts.Client())
	require.NoError(t, err)
	require.Equal(t, version, string(got))
}

func TestGetVerionFromURL_BadHTTPResponse_ReturnsError(t *testing.T) {
	unittest.SmallTest(t)
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "my fake error", http.StatusNotFound)
	}))
	defer ts.Close()

	_, err := GetVersionFromURL(ts.URL, ts.Client())
	require.Error(t, err)
}
