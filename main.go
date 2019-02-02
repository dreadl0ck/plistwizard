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
	"os/exec"
	"sort"

	"github.com/evilsocket/islazy/tui"

	plist "github.com/DHowett/go-plist"
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

	if *flagVersion {

		var (
			files   = searchInfoPlists()
			version string
		)

		if len(files) == 0 {
			fmt.Println("no plists found!")
			os.Exit(1)
		}

		for i, f := range files {
			var (
				dict, _       = readFile(f)
				v, mv         = getVersionNumbers(dict)
				versionString = mv + " (" + v + ")"
			)
			if version != "" && version != versionString {
				printAllVersions()
				log.Fatal("different versions in multiple files: ", files[i-1], " and ", f)
			} else {
				version = versionString
			}
		}

		// print version and exit
		fmt.Println(version)
		os.Exit(0)
	}

	// what would be the next version number?
	if *flagNextVersion {
		dict, _ := readFile(*flagPlist)
		v, mv := getNextVersionNumbers(dict)
		fmt.Println("would bump version from", fmt.Sprint(dict[keyMarketingVersion])+" ("+fmt.Sprint(dict[keyVersion])+")", "to", mv+" ("+v+")")
		return
	}

	// bump marketing version and build number in all target in the directory flag
	if *flagBumpVersion {

		var (
			files       = searchInfoPlists()
			oldVersion  string
			nextVersion string
		)

		for i, f := range files {

			// check if the specified file has a dirty state in git
			// before we even think about modifying it
			if *flagGitCheck {
				err := exec.Command("git", "diff", "--quiet", f).Run()
				if err != nil {
					log.Fatal("file ", f, " is in a dirty state: ", err)
				}
			}

			// read property list file from flag
			dict, format := readFile(f)

			currentVersion := fmt.Sprint(dict[keyMarketingVersion]) + " (" + fmt.Sprint(dict[keyVersion]) + ")"
			if oldVersion != "" && oldVersion != currentVersion {
				log.Fatal("different versions in multiple files: ", files[i-1], " and ", f)
			}
			oldVersion = currentVersion

			v, mv := getNextVersionNumbers(dict)
			dict[keyVersion] = v
			dict[keyMarketingVersion] = mv

			nextVersion = mv + " (" + v + ")"

			// update plist
			data, err := plist.Marshal(dict, format)
			if err != nil {
				log.Fatal(err)
			}

			// formatted := xmlfmt.FormatXML(string(data), "", "")
			formatted := string(data)
			// fmt.Println(formatted)

			// get a file handle and overwrite file contents
			fh, err := os.Create(f)
			if err != nil {
				log.Fatal(err)
			}

			_, err = fh.Write([]byte(formatted))
			if err != nil {
				log.Fatal(err)
			}

			err = fh.Close()
			if err != nil {
				log.Fatal("failed to close file handle:", err)
			}

			fmt.Println("bumped version in", f, "from", oldVersion, "to", mv+" ("+v+")")

			// add changes if they shall be committed via git
			if *flagCommitChanges {
				out, err := exec.Command("git", "add", f).CombinedOutput()
				if err != nil {
					log.Fatal("failed to add file with git:", err, out)
				}
			}
		}

		// commit the changes to git if the flag is set
		if *flagCommitChanges {

			// commit
			out, err := exec.Command("git", "-m", "\"bumped version from "+oldVersion+" to "+nextVersion+"\"").CombinedOutput()
			if err != nil {
				log.Fatal("failed to commit changes with git:", err, out)
			}
		}

		return
	}

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
