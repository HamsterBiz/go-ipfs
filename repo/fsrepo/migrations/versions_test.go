package migrations

import (
	"context"
	"testing"

	"github.com/coreos/go-semver/semver"
)

const testDist = "go-ipfs"

func TestDistVersions(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	vers, err := DistVersions(ctx, testDist, true)
	if err != nil {
		t.Fatal(err)
	}
	if len(vers) == 0 {
		t.Fatal("no versions of", testDist)
	}
	t.Log("There are", len(vers), "versions of", testDist)
	t.Log("Latest 5 are:", vers[:5])
}

func TestLatestDistVersion(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	latest, err := LatestDistVersion(ctx, testDist)
	if err != nil {
		t.Fatal(err)
	}
	if len(latest) < 6 {
		t.Fatal("latest version string too short", latest)
	}
	_, err = semver.NewVersion(latest[1:])
	if err != nil {
		t.Fatal("latest version has invalid format:", latest)
	}
	t.Log("Latest version of", testDist, "is", latest)
}
