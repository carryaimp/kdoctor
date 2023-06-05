// Copyright 2023 Authors of kdoctor-io
// SPDX-License-Identifier: Apache-2.0
package fileManager_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestIppoolCR(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "logger Suite")
}

var _ = BeforeSuite(func() {
	// nothing to do
})
