// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package arm64

import (
	"internal/testenv"
	"regexp"
	"testing"
)

// TestSIMDEncoding verifies the correctness of SIMD floating-point instruction encoding.
func TestSIMDEncoding(t *testing.T) {
	testenv.MustHaveGoBuild(t)

	testCases := []struct {
		name string
		code string
		out  string
	}{
		{
			"VFADD encoding",
			"TEXT ·vfaddtest(SB),$0-0\nVFADD V0.S2, V1.S2, V2.S2\nRET\n",
			`0x0000\s00000\s\(.*\)\tVFADD\tV0\.S2,\sV1\.S2,\sV2\.S2`,
		},
		{
			"VFMUL encoding",
			"TEXT ·vfmultest(SB),$0-0\nVFMUL V0.S2, V1.S2, V2.S2\nRET\n",
			`0x0000\s00000\s\(.*\)\tVFMUL\tV0\.S2,\sV1\.S2,\sV2\.S2`,
		},
		{
			"VFMA encoding",
			"TEXT ·vfmatest(SB),$0-0\nVFMA V0.S2, V1.S2, V2.S2\nRET\n",
			`0x0000\s00000\s\(.*\)\tVFMA\tV0\.S2,\sV1\.S2,\sV2\.S2`,
		},
	}

	for _, test := range testCases {
		out := runAssembler(t, test.code)
		matched, err := regexp.MatchString(test.out, string(out))
		if err != nil {
			t.Fatal(err)
		}
		if !matched {
			t.Errorf("The %s testing failed!\ninput: %s\noutput: %s\n", test.name, test.code, out)
		}
	}
}