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

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/evilsocket/islazy/tui"
)

const (
	// Xcode keys for project versions
	// version string
	keyMarketingVersion = "CFBundleShortVersionString"

	// build number
	keyVersion = "CFBundleVersion"

	// file extension of property lists
	plistFileExtension = ".plist"

	// Xcode project file names
	infoPlistFileName = "Info.plist"
	rootPlistFileName = "Root.plist"
)

func main() {

	// modify usage to show usage synopsis
	originalUsage := flag.Usage
	flag.Usage = func() {
		fmt.Println("plistwizard [-plist <path/to/plist>] [-next-version] [-bump-version] [-lookup <key>]")
		fmt.Println()
		originalUsage()
	}

	// parse commandline flags
	flag.Parse()

	// if the version flag is set
	// check the specified directory recursively
	// for Xcode project files
	if *flagList || *flagListAll {
		printAllVersions()
		os.Exit(0)
	}

	if *flagVersion || *flagBuildNumber || *flagMarketingVersion {

		var (
			files   = searchInfoPlists()
			version string
			mv      string
			v       string
		)

		if len(files) == 0 {
			fmt.Println("no plists found!")
			os.Exit(1)
		}

		for i, f := range files {
			dict, _ := readFile(f)
			v, mv = getVersionNumbers(dict)
			versionString := mv + " (" + v + ")"
			if version != "" && version != versionString {
				printAllVersions()
				log.Fatal("different versions in multiple files: ", files[i-1], " and ", f)
			} else {
				version = versionString
			}
		}

		if *flagBuildNumber {
			// print full version and exit
			fmt.Println(v)
			os.Exit(0)
		}

		if *flagMarketingVersion {
			// print full version and exit
			fmt.Println(mv)
			os.Exit(0)
		}

		// print full version and exit
		fmt.Println(version)
		os.Exit(0)
	}

	// either bump marketing version and build number in all target in the directory flag
	// or show what would be the next version number
	if *flagBumpVersion || *flagNextVersion || *flagBumpInteractive {
		parseProject()
		return
	}

	// read default file
	dict, _ := readFile(*flagPlist)

	// lookup and print value from named argument in dict
	if *flagLookup != "" {
		fmt.Println(dict[*flagLookup])
		return
	}

	// dump a table
	// collect keys and sort alphabetically
	var sortedKeys []string
	for k := range dict {
		sortedKeys = append(sortedKeys, k)
	}
	sort.Strings(sortedKeys)

	// iterate over sorted keys and collect them for table
	var rows [][]string
	for _, key := range sortedKeys {
		if v, ok := dict[key]; ok {
			rows = append(rows, []string{key, fmt.Sprint(v)})
		}
	}

	// print table to stdout
	tui.Table(os.Stdout, []string{"Key", "Value"}, rows)
	return
}
