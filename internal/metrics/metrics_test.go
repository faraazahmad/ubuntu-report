package metrics_test

import (
	"testing"

	"github.com/ubuntu/ubuntu-report/internal/metrics"
)

func TestGetIDS(t *testing.T) {

	testCases := []struct {
		name string
		root string

		wantDistro  string
		wantVersion string
		wantErr     bool
	}{
		{"regular", "testdata/good", "ubuntu", "18.04", false},
		{"doesn't exist", "testdata/none", "", "", true},
		{"empty file", "testdata/empty", "", "", true},
		{"missing distro", "testdata/missing/ids/distro", "", "", true},
		{"missing version", "testdata/missing/ids/version", "", "", true},
		{"missing both", "testdata/missing/ids/both", "", "", true},
		{"empty distro", "testdata/empty-fields/ids/distro", "", "", true},
		{"empty version", "testdata/empty-fields/ids/version", "", "", true},
		{"empty both", "testdata/empty-fields/ids/both", "", "", true},
		{"garbage content", "testdata/garbage", "", "", true},
	}
	for _, tc := range testCases {
		tc := tc // capture range variable for parallel execution
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			m, err := metrics.New(metrics.WithRootAt(tc.root))
			if err != nil {
				t.Fatal("can't create metrics object", err)
			}
			d, v, err := m.GetIDS()

			if err != nil && !tc.wantErr {
				t.Fatal("got an unexpected err:", err)
			}
			if err == nil && tc.wantErr {
				t.Error("expected an error and got none")
			}

			if d != tc.wantDistro {
				t.Errorf("got for distro: %s; want %s", d, tc.wantDistro)
			}
			if v != tc.wantVersion {
				t.Errorf("got for version: %s; want %s", v, tc.wantVersion)
			}
		})
	}
}
