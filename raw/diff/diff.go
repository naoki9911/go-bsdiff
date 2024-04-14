// Copyright (c) 2017-2022 Carl Kittelberger.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE.txt file.

package diff

import (
	"io"

	"github.com/icedream/go-bsdiff/internal/native"
)

/*
Diff generates a raw patch from old content that will be read in completely from
oldReader and new content that will be read in completely from newReader and
saves that patch to patchWriter.

It may be helpful to save away the new content size along with the actual
patch as it will be needed in order to reuse the patch.
*/
func Diff(oldBytes, newBytes []byte, patchWriter io.Writer) (err error) {
	return native.Diff(oldBytes, newBytes, patchWriter)
}
