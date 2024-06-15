// SPDX-FileCopyrightText: 2023-2024 Steffen Vogel <post@steffenvogel.de>
// SPDX-License-Identifier: Apache-2.0

package skeleton_test

import (
	"testing"

	"cunicu.li/skeleton"
)

func TestHello(t *testing.T) {
	if skeleton.Hello() != "hello" {
		t.Error()
	}
}
