/*
 * PLISTWIZARD - A magically simple tool for XML property lists from Xcode
 * Copyright (c) 2018 Philipp Mieden <dreadl0ck [at] protonmail [dot] ch>
 *
 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 */

package main

// the logo consists of characters that cannot be put into go source code
// so they will simply be embedded
var logo = []byte{
	0x0a, 0x20, 0x20, 0xe2, 0x96, 0x88, 0xe2, 0x96, 0x88, 0xe2, 0x96, 0x93, 0xe2, 0x96, 0x88, 0xe2,
	0x96, 0x88, 0xe2, 0x96, 0x88, 0x20, 0x20, 0x20, 0xe2, 0x96, 0x88, 0xe2, 0x96, 0x88, 0xe2, 0x96,
	0x93, 0x20, 0x20, 0x20, 0x20, 0x20, 0xe2, 0x96, 0x88, 0xe2, 0x96, 0x88, 0xe2, 0x96, 0x93, 0x20,
	0x20, 0xe2, 0x96, 0x88, 0xe2, 0x96, 0x88, 0xe2, 0x96, 0x88, 0xe2, 0x96, 0x88, 0xe2, 0x96, 0x88,
	0xe2, 0x96, 0x88, 0x20, 0xe2, 0x96, 0x84, 0xe2, 0x96, 0x84, 0xe2, 0x96, 0x84, 0xe2, 0x96, 0x88,
	0xe2, 0x96, 0x88, 0xe2, 0x96, 0x88, 0xe2, 0x96, 0x88, 0xe2, 0x96, 0x88, 0xe2, 0x96, 0x93, 0x20,
	0xe2, 0x96, 0x88, 0x20, 0x20, 0x20, 0x20, 0x20, 0xe2, 0x96, 0x88, 0xe2, 0x96, 0x91, 0x20, 0xe2,
	0x96, 0x88, 0xe2, 0x96, 0x88, 0xe2, 0x96, 0x93, 0xe2, 0x96, 0x92, 0xe2, 0x96, 0x88, 0xe2, 0x96,
	0x88, 0xe2, 0x96, 0x88, 0xe2, 0x96, 0x88, 0xe2, 0x96, 0x88, 0xe2, 0x96, 0x88, 0xe2, 0x96, 0x88,
	0xe2, 0x96, 0x92, 0x20, 0xe2, 0x96, 0x84, 0xe2, 0x96, 0x84, 0xe2, 0x96, 0x84, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0xe2, 0x96, 0x88, 0xe2, 0x96, 0x88, 0xe2, 0x96, 0x80, 0xe2, 0x96, 0x88,
	0xe2, 0x96, 0x88, 0xe2, 0x96, 0x88, 0x20, 0x20, 0xe2, 0x96, 0x93, 0xe2, 0x96, 0x88, 0xe2, 0x96,
	0x88, 0xe2, 0x96, 0x88, 0xe2, 0x96, 0x88, 0xe2, 0x96, 0x88, 0xe2, 0x96, 0x84, 0x20, 0x0a, 0x20,
	0xe2, 0x96, 0x93, 0xe2, 0x96, 0x88, 0xe2, 0x96, 0x88, 0xe2, 0x96, 0x91, 0x20, 0x20, 0xe2, 0x96,
	0x88, 0xe2, 0x96, 0x88, 0xe2, 0x96, 0x92, 0xe2, 0x96, 0x93, 0xe2, 0x96, 0x88, 0xe2, 0x96, 0x88,
	0xe2, 0x96, 0x92, 0x20, 0x20, 0x20, 0x20, 0xe2, 0x96, 0x93, 0xe2, 0x96, 0x88, 0xe2, 0x96, 0x88,
	0xe2, 0x96, 0x92, 0xe2, 0x96, 0x92, 0xe2, 0x96, 0x88, 0xe2, 0x96, 0x88, 0x20, 0x20, 0x20, 0x20,
	0xe2, 0x96, 0x92, 0x20, 0xe2, 0x96, 0x93, 0x20, 0x20, 0xe2, 0x96, 0x88, 0xe2, 0x96, 0x88, 0xe2,
	0x96, 0x92, 0x20, 0xe2, 0x96, 0x93, 0xe2, 0x96, 0x92, 0xe2, 0x96, 0x93, 0xe2, 0x96, 0x88, 0xe2,
	0x96, 0x91, 0x20, 0xe2, 0x96, 0x88, 0x20, 0xe2, 0x96, 0x91, 0xe2, 0x96, 0x88, 0xe2, 0x96, 0x91,
	0xe2, 0x96, 0x93, 0xe2, 0x96, 0x88, 0xe2, 0x96, 0x88, 0xe2, 0x96, 0x92, 0xe2, 0x96, 0x92, 0x20,
	0xe2, 0x96, 0x92, 0x20, 0xe2, 0x96, 0x92, 0x20, 0xe2, 0x96, 0x84, 0xe2, 0x96, 0x80, 0xe2, 0x96,
	0x91, 0xe2, 0x96, 0x92, 0xe2, 0x96, 0x88, 0xe2, 0x96, 0x88, 0xe2, 0x96, 0x88, 0xe2, 0x96, 0x88,
	0xe2, 0x96, 0x84, 0x20, 0x20, 0x20, 0x20, 0xe2, 0x96, 0x93, 0xe2, 0x96, 0x88, 0xe2, 0x96, 0x88,
	0x20, 0xe2, 0x96, 0x92, 0x20, 0xe2, 0x96, 0x88, 0xe2, 0x96, 0x88, 0xe2, 0x96, 0x92, 0xe2, 0x96,
	0x92, 0xe2, 0x96, 0x88, 0xe2, 0x96, 0x88, 0xe2, 0x96, 0x80, 0x20, 0xe2, 0x96, 0x88, 0xe2, 0x96,
	0x88, 0xe2, 0x96, 0x8c, 0x0a, 0x20, 0xe2, 0x96, 0x93, 0xe2, 0x96, 0x88, 0xe2, 0x96, 0x88, 0xe2,
	0x96, 0x91, 0x20, 0xe2, 0x96, 0x88, 0xe2, 0x96, 0x88, 0xe2, 0x96, 0x93, 0xe2, 0x96, 0x92, 0xe2,
	0x96, 0x92, 0xe2, 0x96, 0x88, 0xe2, 0x96, 0x88, 0xe2, 0x96, 0x91, 0x20, 0x20, 0x20, 0x20, 0xe2,
	0x96, 0x92, 0xe2, 0x96, 0x88, 0xe2, 0x96, 0x88, 0xe2, 0x96, 0x92, 0xe2, 0x96, 0x91, 0x20, 0xe2,
	0x96, 0x93, 0xe2, 0x96, 0x88, 0xe2, 0x96, 0x88, 0xe2, 0x96, 0x84, 0x20, 0x20, 0x20, 0xe2, 0x96,
	0x92, 0x20, 0xe2, 0x96, 0x93, 0xe2, 0x96, 0x88, 0xe2, 0x96, 0x88, 0xe2, 0x96, 0x91, 0x20, 0xe2,
	0x96, 0x92, 0xe2, 0x96, 0x91, 0xe2, 0x96, 0x92, 0xe2, 0x96, 0x88, 0xe2, 0x96, 0x91, 0x20, 0xe2,
	0x96, 0x88, 0x20, 0xe2, 0x96, 0x91, 0xe2, 0x96, 0x88, 0x20, 0xe2, 0x96, 0x92, 0xe2, 0x96, 0x88,
	0xe2, 0x96, 0x88, 0xe2, 0x96, 0x92, 0xe2, 0x96, 0x91, 0x20, 0xe2, 0x96, 0x92, 0x20, 0xe2, 0x96,
	0x84, 0xe2, 0x96, 0x80, 0xe2, 0x96, 0x92, 0xe2, 0x96, 0x91, 0x20, 0xe2, 0x96, 0x92, 0xe2, 0x96,
	0x88, 0xe2, 0x96, 0x88, 0x20, 0x20, 0xe2, 0x96, 0x80, 0xe2, 0x96, 0x88, 0xe2, 0x96, 0x84, 0x20,
	0x20, 0xe2, 0x96, 0x93, 0xe2, 0x96, 0x88, 0xe2, 0x96, 0x88, 0x20, 0xe2, 0x96, 0x91, 0xe2, 0x96,
	0x84, 0xe2, 0x96, 0x88, 0x20, 0xe2, 0x96, 0x92, 0xe2, 0x96, 0x91, 0xe2, 0x96, 0x88, 0xe2, 0x96,
	0x88, 0x20, 0x20, 0x20, 0xe2, 0x96, 0x88, 0xe2, 0x96, 0x8c, 0x0a, 0x20, 0xe2, 0x96, 0x92, 0xe2,
	0x96, 0x88, 0xe2, 0x96, 0x88, 0xe2, 0x96, 0x84, 0xe2, 0x96, 0x88, 0xe2, 0x96, 0x93, 0xe2, 0x96,
	0x92, 0x20, 0xe2, 0x96, 0x92, 0xe2, 0x96, 0x92, 0xe2, 0x96, 0x88, 0xe2, 0x96, 0x88, 0xe2, 0x96,
	0x91, 0x20, 0x20, 0x20, 0x20, 0xe2, 0x96, 0x91, 0xe2, 0x96, 0x88, 0xe2, 0x96, 0x88, 0xe2, 0x96,
	0x91, 0x20, 0x20, 0xe2, 0x96, 0x92, 0x20, 0x20, 0x20, 0xe2, 0x96, 0x88, 0xe2, 0x96, 0x88, 0xe2,
	0x96, 0x92, 0xe2, 0x96, 0x91, 0x20, 0xe2, 0x96, 0x93, 0xe2, 0x96, 0x88, 0xe2, 0x96, 0x88, 0xe2,
	0x96, 0x93, 0x20, 0xe2, 0x96, 0x91, 0x20, 0xe2, 0x96, 0x91, 0xe2, 0x96, 0x88, 0xe2, 0x96, 0x91,
	0x20, 0xe2, 0x96, 0x88, 0x20, 0xe2, 0x96, 0x91, 0xe2, 0x96, 0x88, 0x20, 0xe2, 0x96, 0x91, 0xe2,
	0x96, 0x88, 0xe2, 0x96, 0x88, 0xe2, 0x96, 0x91, 0x20, 0x20, 0xe2, 0x96, 0x84, 0xe2, 0x96, 0x80,
	0xe2, 0x96, 0x92, 0x20, 0x20, 0x20, 0xe2, 0x96, 0x91, 0xe2, 0x96, 0x91, 0xe2, 0x96, 0x88, 0xe2,
	0x96, 0x88, 0xe2, 0x96, 0x84, 0xe2, 0x96, 0x84, 0xe2, 0x96, 0x84, 0xe2, 0x96, 0x84, 0xe2, 0x96,
	0x88, 0xe2, 0x96, 0x88, 0x20, 0xe2, 0x96, 0x92, 0xe2, 0x96, 0x88, 0xe2, 0x96, 0x88, 0xe2, 0x96,
	0x80, 0xe2, 0x96, 0x80, 0xe2, 0x96, 0x88, 0xe2, 0x96, 0x84, 0x20, 0x20, 0xe2, 0x96, 0x91, 0xe2,
	0x96, 0x93, 0xe2, 0x96, 0x88, 0xe2, 0x96, 0x84, 0x20, 0x20, 0x20, 0xe2, 0x96, 0x8c, 0x0a, 0x20,
	0xe2, 0x96, 0x92, 0xe2, 0x96, 0x88, 0xe2, 0x96, 0x88, 0xe2, 0x96, 0x92, 0x20, 0xe2, 0x96, 0x91,
	0x20, 0x20, 0xe2, 0x96, 0x91, 0xe2, 0x96, 0x91, 0xe2, 0x96, 0x88, 0xe2, 0x96, 0x88, 0xe2, 0x96,
	0x88, 0xe2, 0x96, 0x88, 0xe2, 0x96, 0x88, 0xe2, 0x96, 0x88, 0xe2, 0x96, 0x92, 0xe2, 0x96, 0x91,
	0xe2, 0x96, 0x88, 0xe2, 0x96, 0x88, 0xe2, 0x96, 0x91, 0xe2, 0x96, 0x92, 0xe2, 0x96, 0x88, 0xe2,
	0x96, 0x88, 0xe2, 0x96, 0x88, 0xe2, 0x96, 0x88, 0xe2, 0x96, 0x88, 0xe2, 0x96, 0x88, 0xe2, 0x96,
	0x92, 0xe2, 0x96, 0x92, 0x20, 0x20, 0xe2, 0x96, 0x92, 0xe2, 0x96, 0x88, 0xe2, 0x96, 0x88, 0xe2,
	0x96, 0x92, 0x20, 0xe2, 0x96, 0x91, 0x20, 0xe2, 0x96, 0x91, 0xe2, 0x96, 0x91, 0xe2, 0x96, 0x88,
	0xe2, 0x96, 0x88, 0xe2, 0x96, 0x92, 0xe2, 0x96, 0x88, 0xe2, 0x96, 0x88, 0xe2, 0x96, 0x93, 0x20,
	0xe2, 0x96, 0x91, 0xe2, 0x96, 0x88, 0xe2, 0x96, 0x88, 0xe2, 0x96, 0x91, 0xe2, 0x96, 0x92, 0xe2,
	0x96, 0x88, 0xe2, 0x96, 0x88, 0xe2, 0x96, 0x88, 0xe2, 0x96, 0x88, 0xe2, 0x96, 0x88, 0xe2, 0x96,
	0x88, 0xe2, 0x96, 0x88, 0xe2, 0x96, 0x92, 0x20, 0xe2, 0x96, 0x93, 0xe2, 0x96, 0x88, 0x20, 0x20,
	0x20, 0xe2, 0x96, 0x93, 0xe2, 0x96, 0x88, 0xe2, 0x96, 0x88, 0xe2, 0x96, 0x92, 0xe2, 0x96, 0x91,
	0xe2, 0x96, 0x88, 0xe2, 0x96, 0x88, 0xe2, 0x96, 0x93, 0x20, 0xe2, 0x96, 0x92, 0xe2, 0x96, 0x88,
	0xe2, 0x96, 0x88, 0xe2, 0x96, 0x92, 0xe2, 0x96, 0x91, 0xe2, 0x96, 0x92, 0xe2, 0x96, 0x88, 0xe2,
	0x96, 0x88, 0xe2, 0x96, 0x88, 0xe2, 0x96, 0x88, 0xe2, 0x96, 0x93, 0x20, 0x0a, 0x20, 0xe2, 0x96,
	0x92, 0xe2, 0x96, 0x93, 0xe2, 0x96, 0x92, 0xe2, 0x96, 0x91, 0x20, 0xe2, 0x96, 0x91, 0x20, 0x20,
	0xe2, 0x96, 0x91, 0xe2, 0x96, 0x91, 0x20, 0xe2, 0x96, 0x92, 0xe2, 0x96, 0x91, 0xe2, 0x96, 0x93,
	0x20, 0x20, 0xe2, 0x96, 0x91, 0xe2, 0x96, 0x91, 0xe2, 0x96, 0x93, 0x20, 0x20, 0xe2, 0x96, 0x92,
	0x20, 0xe2, 0x96, 0x92, 0xe2, 0x96, 0x93, 0xe2, 0x96, 0x92, 0x20, 0xe2, 0x96, 0x92, 0x20, 0xe2,
	0x96, 0x91, 0x20, 0x20, 0xe2, 0x96, 0x92, 0x20, 0xe2, 0x96, 0x91, 0xe2, 0x96, 0x91, 0x20, 0x20,
	0x20, 0xe2, 0x96, 0x91, 0x20, 0xe2, 0x96, 0x93, 0xe2, 0x96, 0x91, 0xe2, 0x96, 0x92, 0x20, 0xe2,
	0x96, 0x92, 0x20, 0x20, 0xe2, 0x96, 0x91, 0xe2, 0x96, 0x93, 0x20, 0x20, 0xe2, 0x96, 0x91, 0xe2,
	0x96, 0x92, 0xe2, 0x96, 0x92, 0x20, 0xe2, 0x96, 0x93, 0xe2, 0x96, 0x91, 0xe2, 0x96, 0x92, 0xe2,
	0x96, 0x91, 0xe2, 0x96, 0x92, 0x20, 0xe2, 0x96, 0x92, 0xe2, 0x96, 0x92, 0x20, 0x20, 0x20, 0xe2,
	0x96, 0x93, 0xe2, 0x96, 0x92, 0xe2, 0x96, 0x88, 0xe2, 0x96, 0x91, 0xe2, 0x96, 0x91, 0x20, 0xe2,
	0x96, 0x92, 0xe2, 0x96, 0x93, 0x20, 0xe2, 0x96, 0x91, 0xe2, 0x96, 0x92, 0xe2, 0x96, 0x93, 0xe2,
	0x96, 0x91, 0x20, 0xe2, 0x96, 0x92, 0xe2, 0x96, 0x92, 0xe2, 0x96, 0x93, 0x20, 0x20, 0xe2, 0x96,
	0x92, 0x20, 0x0a, 0x20, 0xe2, 0x96, 0x91, 0xe2, 0x96, 0x92, 0x20, 0xe2, 0x96, 0x91, 0x20, 0x20,
	0x20, 0x20, 0x20, 0xe2, 0x96, 0x91, 0x20, 0xe2, 0x96, 0x91, 0x20, 0xe2, 0x96, 0x92, 0x20, 0x20,
	0xe2, 0x96, 0x91, 0x20, 0xe2, 0x96, 0x92, 0x20, 0xe2, 0x96, 0x91, 0xe2, 0x96, 0x91, 0x20, 0xe2,
	0x96, 0x91, 0xe2, 0x96, 0x92, 0x20, 0x20, 0xe2, 0x96, 0x91, 0x20, 0xe2, 0x96, 0x91, 0x20, 0x20,
	0x20, 0x20, 0xe2, 0x96, 0x91, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0xe2, 0x96, 0x92, 0x20, 0xe2,
	0x96, 0x91, 0x20, 0xe2, 0x96, 0x91, 0x20, 0x20, 0x20, 0xe2, 0x96, 0x92, 0x20, 0xe2, 0x96, 0x91,
	0xe2, 0x96, 0x91, 0xe2, 0x96, 0x91, 0xe2, 0x96, 0x92, 0x20, 0xe2, 0x96, 0x92, 0x20, 0xe2, 0x96,
	0x91, 0x20, 0xe2, 0x96, 0x92, 0x20, 0x20, 0xe2, 0x96, 0x92, 0x20, 0x20, 0x20, 0xe2, 0x96, 0x92,
	0xe2, 0x96, 0x92, 0x20, 0xe2, 0x96, 0x91, 0x20, 0x20, 0xe2, 0x96, 0x91, 0xe2, 0x96, 0x92, 0x20,
	0xe2, 0x96, 0x91, 0x20, 0xe2, 0x96, 0x92, 0xe2, 0x96, 0x91, 0x20, 0xe2, 0x96, 0x91, 0x20, 0xe2,
	0x96, 0x92, 0x20, 0x20, 0xe2, 0x96, 0x92, 0x20, 0x0a, 0x20, 0xe2, 0x96, 0x91, 0xe2, 0x96, 0x91,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0xe2, 0x96, 0x91, 0x20, 0xe2, 0x96, 0x91,
	0x20, 0x20, 0x20, 0x20, 0xe2, 0x96, 0x92, 0x20, 0xe2, 0x96, 0x91, 0xe2, 0x96, 0x91, 0x20, 0x20,
	0xe2, 0x96, 0x91, 0x20, 0x20, 0xe2, 0x96, 0x91, 0x20, 0x20, 0x20, 0x20, 0xe2, 0x96, 0x91, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0xe2, 0x96, 0x91, 0x20, 0x20, 0x20, 0xe2, 0x96, 0x91,
	0x20, 0x20, 0x20, 0xe2, 0x96, 0x92, 0x20, 0xe2, 0x96, 0x91, 0xe2, 0x96, 0x91, 0x20, 0xe2, 0x96,
	0x91, 0x20, 0xe2, 0x96, 0x91, 0x20, 0xe2, 0x96, 0x91, 0x20, 0xe2, 0x96, 0x91, 0x20, 0x20, 0xe2,
	0x96, 0x91, 0x20, 0x20, 0x20, 0xe2, 0x96, 0x92, 0x20, 0x20, 0x20, 0x20, 0x20, 0xe2, 0x96, 0x91,
	0xe2, 0x96, 0x91, 0x20, 0x20, 0x20, 0xe2, 0x96, 0x91, 0x20, 0x20, 0xe2, 0x96, 0x91, 0x20, 0xe2,
	0x96, 0x91, 0x20, 0x20, 0xe2, 0x96, 0x91,
}