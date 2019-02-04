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

import "flag"

var (
	flagPlist   = flag.String("plist", "Info.plist", "path to .plist file")
	flagLookup  = flag.String("lookup", "", "lookup the value for a key in the property list dictionary")
	flagVersion = flag.Bool("version", false, "print main Xcode project version in format 'marketingVersion (buildNumber)' and exit")

	// git
	flagGitCheck      = flag.Bool("git-check", true, "abort if input file has a dirty state in git")
	flagCommitChanges = flag.Bool("commit", false, "commit changes on plist file with git")

	// version bumping
	flagBumpVersion = flag.Bool("bump-version", false, "bump version and build number")
	flagNextVersion = flag.Bool("next-version", false, "only print next version value and dont bump version")
	flagBumpMajor   = flag.Bool("major", false, "bump major version")
	flagBumpMinor   = flag.Bool("minor", false, "bump minor version")
	flagBumpPatch   = flag.Bool("patch", true, "bump patch version")

	flagDir     = flag.String("dir", ".", "set directory")
	flagList    = flag.Bool("list", false, "recursively search for plists and list them all")
	flagListAll = flag.Bool("listall", false, "recursively search for plists and list them all")

	flagBuildNumber      = flag.Bool("build-number", false, "print only build number")
	flagMarketingVersion = flag.Bool("marketing-version", false, "print only marketing version")
	flagBumpInteractive  = flag.Bool("bump-interactive", false, "bump version interactive")
)
