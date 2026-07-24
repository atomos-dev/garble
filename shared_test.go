// Copyright (c) 2026, The Garble Authors.
// See LICENSE for licensing information.

package main

import "testing"

func TestMatchGarblePatterns(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		patterns string
		target   string
		want     bool
	}{
		{name: "include", patterns: "dpncore*", target: "dpncore/platform/iosbind", want: true},
		{name: "unmatched", patterns: "dpncore*", target: "example.com/other", want: false},
		{
			name:     "excluded package",
			patterns: "dpncore*,!dpncore/utils/infra/log_store",
			target:   "dpncore/utils/infra/log_store",
			want:     false,
		},
		{
			name:     "excluded subpackage",
			patterns: "dpncore*,!dpncore/utils/infra/log_store",
			target:   "dpncore/utils/infra/log_store/internal",
			want:     false,
		},
		{
			name:     "sibling remains included",
			patterns: "dpncore*,!dpncore/utils/infra/log_store",
			target:   "dpncore/utils/infra/dpn_log",
			want:     true,
		},
		{
			name:     "exclusion wins regardless of order",
			patterns: "!dpncore/utils/infra/log_store,dpncore*",
			target:   "dpncore/utils/infra/log_store",
			want:     false,
		},
		{name: "exclusion without include", patterns: "!dpncore/private", target: "dpncore/public", want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			if got := matchGarblePatterns(test.patterns, test.target); got != test.want {
				t.Fatalf("matchGarblePatterns(%q, %q) = %v, want %v", test.patterns, test.target, got, test.want)
			}
		})
	}
}
