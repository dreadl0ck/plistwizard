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
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	plist "howett.net/plist"
	prompt "github.com/c-bata/go-prompt"
)

func completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "major", Description: "Bump the major version"},
		{Text: "minor", Description: "Bump the minor version"},
		{Text: "patch", Description: "Bump the patch version (default)"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func parseProject() {

	var (
		files       = searchInfoPlists()
		oldVersion  string
		nextVersion string
	)

	if *flagBumpInteractive {
		fmt.Println("v{major.minor.patch}")
		fmt.Println("Please select what part of the version number to bump (Hit Tab for completion)")
		fmt.Println("Hit [Enter] to select the default (patch)")
		t := prompt.Input("> ", completer)
		switch t {
		case "major":
			*flagBumpMajor = true
		case "minor":
			*flagBumpMinor = true
		case "patch":
			*flagBumpPatch = true
		default:
		}
	}

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
		nextVersion = mv + " (" + v + ")"
		if *flagNextVersion {
			fmt.Println("would bump version from", oldVersion, "to", nextVersion)
			return
		}

		writeVersionUpdate(dict, format, v, mv, f, oldVersion)
	}

	// commit the changes to git if the flag is set
	if *flagCommitChanges {

		// commit
		out, err := exec.Command("git", "-m", "\"bumped version from "+oldVersion+" to "+nextVersion+"\"").CombinedOutput()
		if err != nil {
			log.Fatal("failed to commit changes with git:", err, out)
		}
	}
}

func writeVersionUpdate(dict map[string]interface{}, format int, v string, mv string, f string, oldVersion string) {

	dict[keyVersion] = v
	dict[keyMarketingVersion] = mv

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

// ranges the currenty directory recursively and searches for property list from Xcode
// prints each found path to stdout in the format: <path>: version (buildNumber)
func printAllVersions() {
	files := searchInfoPlists()
	for _, f := range files {
		dict, _ := readFile(f)
		v, mv := getVersionNumbers(dict)
		fmt.Println(f, mv+" ("+v+")")
	}
}

// read the current values of the version numbers from the dictionary
func getVersionNumbers(dict map[string]interface{}) (buildNumber string, versionString string) {
	return fmt.Sprint(dict[keyVersion]), fmt.Sprint(dict[keyMarketingVersion])
}

// This function parses the value from the version fields
// and bumps the requested part of the human readable version number (default: Patch -> v{Major,Minor,Patch})
// and the machine readable version number (aka build number) by 1
// returning the updated values
//
// From the Apple Documentation:
// For every new build you submit, you will need to invent a new build number whose value is greater than the last build number you used (for that same version).
// For iOS apps, you may re-use build numbers when submitting different versions.
// For macOS apps, you must chose a new build number for every submission that is unique
// and has never been used before in any submission you have provided to the App Store (including build numbers used in previous versions of your app).
func getNextVersionNumbers(dict map[string]interface{}) (buildNumber string, versionString string) {

	var (
		v, mv   = getVersionNumbers(dict)
		mvSlice = strings.Split(mv, ".")
		length  = len(mvSlice)
	)

	// increment and update build number
	i, err := strconv.Atoi(v)
	if err != nil {
		log.Fatal(err)
	}
	i++

	// increment patch value
	// {major.minor.patch}
	switch {
	case *flagBumpMajor:
		if length < 1 {
			log.Fatal("invalid version string:", mvSlice)
		}
		i, err := strconv.Atoi(mvSlice[0])
		if err != nil {
			log.Fatal(err)
		}
		i++
		mvSlice[0] = strconv.Itoa(i)

		// set minor to 0
		if length > 1 {
			mvSlice[1] = "0"
		}
		// set patch to 0
		if length == 3 {
			mvSlice[2] = "0"
		}
	case *flagBumpMinor:
		if length < 2 {
			log.Fatal("invalid version string:", mvSlice)
		}
		i, err := strconv.Atoi(mvSlice[1])
		if err != nil {
			log.Fatal(err)
		}
		i++
		mvSlice[1] = strconv.Itoa(i)

		// set patch to 0
		if length == 3 {
			mvSlice[2] = "0"
		}
	case *flagBumpPatch:
		if length < 3 {
			log.Fatal("invalid version string:", mvSlice)
		}
		i, err := strconv.Atoi(mvSlice[2])
		if err != nil {
			log.Fatal(err)
		}
		i++
		mvSlice[2] = strconv.Itoa(i)
	default:
		log.Fatal("unexpected length for marketingVersion slice", len(mvSlice))
	}

	return strconv.Itoa(i), strings.Join(mvSlice, ".")
}

func searchInfoPlists() (files []string) {

	err := filepath.Walk(*flagDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if filepath.Ext(path) == plistFileExtension {
			if *flagListAll {
				fmt.Println(path)
			}
			if info.Name() == infoPlistFileName || info.Name() == rootPlistFileName {
				// exclude files from Pods directory
				// and those generated by Xcode by setting the derived data dir to the project directory
				if !strings.Contains(path, "Pods") && !strings.Contains(path, "Build") {

					// ensure the file belongs to the current project and not some dependecy or embedded framework
					if len(strings.Split(path, "/")) == 2 {
						files = append(files, path)
					}
				}
			}
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	if *flagListAll {
		os.Exit(0)
	}

	return
}

func readFile(path string) (dict map[string]interface{}, format int) {

	dict = map[string]interface{}{}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("failed to read file:", err)
		os.Exit(1)
	}

	format, err = plist.Unmarshal(data, &dict)
	if err != nil {
		panic(err)
	}

	return dict, format
}
