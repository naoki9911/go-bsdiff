// Copyright (c) 2017-2022 Carl Kittelberger.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE.txt file.

package patch

import (
	"io"

	"github.com/icedream/go-bsdiff/internal/native"
)

/*
Patch reads a raw patch from patchReader and applies it on top of the old
content which will be read from oldReader, saving the resulting new content to
newWriter.

newSize needs to be exactly the size of the new file that should be generated
from the patch.
*/
func Patch(oldBytes []byte, patchReader io.Reader, newSize uint64) ([]byte, error) {
	newBytes := make([]byte, newSize)
	err := native.Patch(oldBytes, newBytes, patchReader)
	if err != nil {
		return nil, err
	}

	return newBytes, nil
}
