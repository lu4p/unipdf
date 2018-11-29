/*
 * This file is subject to the terms and conditions defined in
 * file 'LICENSE.md', which is part of this source code package.
 */

package textencoding

// NewWinAnsiTextEncoder returns a SimpleEncoder that implements WinAnsiEncoding.
func NewWinAnsiTextEncoder() SimpleEncoder {
	const baseName = "WinAnsiEncoding"
	enc := newSimpleTextEncoder(winAnsiEncoding, baseName, nil)
	return *enc
}
