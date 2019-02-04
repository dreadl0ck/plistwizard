# PLIST WIZARD

         ██▓███   ██▓     ██▓  ██████ ▄▄▄█████▓ █     █░ ██▓▒███████▒ ▄▄▄       ██▀███  ▓█████▄
        ▓██░  ██▒▓██▒    ▓██▒▒██    ▒ ▓  ██▒ ▓▒▓█░ █ ░█░▓██▒▒ ▒ ▒ ▄▀░▒████▄    ▓██ ▒ ██▒▒██▀ ██▌
        ▓██░ ██▓▒▒██░    ▒██▒░ ▓██▄   ▒ ▓██░ ▒░▒█░ █ ░█ ▒██▒░ ▒ ▄▀▒░ ▒██  ▀█▄  ▓██ ░▄█ ▒░██   █▌
        ▒██▄█▓▒ ▒▒██░    ░██░  ▒   ██▒░ ▓██▓ ░ ░█░ █ ░█ ░██░  ▄▀▒   ░░██▄▄▄▄██ ▒██▀▀█▄  ░▓█▄   ▌
        ▒██▒ ░  ░░██████▒░██░▒██████▒▒  ▒██▒ ░ ░░██▒██▓ ░██░▒███████▒ ▓█   ▓██▒░██▓ ▒██▒░▒████▓
        ▒▓▒░ ░  ░░ ▒░▓  ░░▓  ▒ ▒▓▒ ▒ ░  ▒ ░░   ░ ▓░▒ ▒  ░▓  ░▒▒ ▓░▒░▒ ▒▒   ▓▒█░░ ▒▓ ░▒▓░ ▒▒▓  ▒
        ░▒ ░     ░ ░ ▒  ░ ▒ ░░ ░▒  ░ ░    ░      ▒ ░ ░   ▒ ░░░▒ ▒ ░ ▒  ▒   ▒▒ ░  ░▒ ░ ▒░ ░ ▒  ▒
        ░░         ░ ░    ▒ ░░  ░  ░    ░        ░   ░   ▒ ░░ ░ ░ ░ ░  ░   ▒     ░░   ░  ░ ░  ░

[![Go Report Card](https://goreportcard.com/badge/github.com/dreadl0ck/plistwizard)](https://goreportcard.com/report/github.com/dreadl0ck/plistwizard)
[![License](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://raw.githubusercontent.com/dreadl0ck/plistwizard/master/docs/LICENSE)
[![Golang](https://img.shields.io/badge/Go-1.11-blue.svg)](https://golang.org)
![Linux](https://img.shields.io/badge/Supports-Linux-green.svg)
![macOS](https://img.shields.io/badge/Supports-macOS-green.svg)
![windows](https://img.shields.io/badge/Supports-windows-green.svg)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/dreadl0ck/plistwizard)

This magically simple tool was created to parse *property list* XML files from Xcode,
more specifically the *Info.plist* files that contain various project related information.

It can be used to query for specific values or bump the version and build number automatically.

Apple provides the [agvtool](http://www.manpagez.com/man/1/agvtool/) for this purpose, however the tool has various annoying limitations:

- bugs when working on projects with multiple targets
- no option to specify a path to a plist file of interest
- no option to increment the marketing version automatically
- no git integration
- source code of the tool is not available

Since parsing XML is not witchcraft, this is a simple and extensible implementation of the desired functionality in Go.

## Future Plans

- add support for adding / updating entries in the dictionary
- format XML

## Examples

Show version inside an Xcode project directory:

    $ plistwizard -version
    0.4.12 (117)

Show a table with elements of the data dictionary:

    $ plistwizard -plist MyProject/Info.plist
    +--------------------------------+-----------------------------------------------------------+
    |              Key               |                          Value                            |
    +--------------------------------+-----------------------------------------------------------+
    | CFBundleDevelopmentRegion      | $(DEVELOPMENT_LANGUAGE)                                   |
    | CFBundleExecutable             | $(EXECUTABLE_NAME)                                        |
    | CFBundleIconFile               |                                                           |
    | CFBundleIdentifier             | $(PRODUCT_BUNDLE_IDENTIFIER)                              |
    | CFBundleInfoDictionaryVersion  | 6.0                                                       |
    | CFBundleName                   | $(PRODUCT_NAME)                                           |
    | CFBundlePackageType            | APPL                                                      |
    | CFBundleShortVersionString     | 0.4.12                                                    |
    | CFBundleVersion                | 117                                                       |
    | LSApplicationCategoryType      | public.app-category.utilities                             |
    | LSMinimumSystemVersion         | $(MACOSX_DEPLOYMENT_TARGET)                               |
    | LSUIElement                    | true                                                      |
    | NSHumanReadableCopyright       | Copyright © 2018 Me. All rights reserved.                 |
    | NSMainStoryboardFile           | Main                                                      |
    | NSPrincipalClass               | NSApplication                                             |
    | NSUserNotificationAlertStyle   | alert                                                     |
    +--------------------------------+-----------------------------------------------------------+

Lookup a value for the specified key inside the dictionary inside the plist:

    $ plistwizard -plist MyProject/Info.plist -lookup LSApplicationCategoryType
    public.app-category.utilities

Show the next version:

    $ plistwizard -next-version
    would bump version from 0.4.12 (117) to 0.4.13 (118)

Bump the major version number (works also when there are multiple targets inside the project):

    $ plistwizard -bump-version -major
    bumped version in MyProject/Info.plist 0.4.12 (117) to 1.0.0 (118)
    bumped version in MyProject Helper/Info.plist 0.4.12 (117) to 1.0.0 (118)

This will set all targets to the same version and fatal if the versions are different.

Bump the minor version number:

    $ plistwizard -bump-version -minor
    bumped version in MyProject/Info.plist 0.4.12 (117) to 0.5.0 (118)
    bumped version in MyProject Helper/Info.plist 0.4.12 (117) to 0.5.0 (118)

Bump the patch version number (default!):

    $ plistwizard -bump-version
    bumped version in MyProject/Info.plist 0.4.12 (117) to 0.4.13 (118)
    bumped version in MyProject Helper/Info.plist 0.4.12 (117) to 0.4.13 (118)

For usage in scripts, checkout the interactive bumping:

    $ plistwizard -bump-version-interactive

This will ask which part of the version to bump in an interactive shell with tab completion!

## Help

    $ plistwizard -h
    plistwizard [-plist <path/to/plist>] [-next-version] [-bump-version] [-lookup <key>]
    
    Usage of plistwizard:
    -build-number
            print only build number
    -bump-interactive
            bump version interactive
    -bump-version
            bump version and build number
    -commit
            commit changes on plist file with git
    -dir string
            set directory (default ".")
    -git-check
            abort if input file has a dirty state in git (default true)
    -list
            recursively search for plists and list them all
    -listall
            recursively search for plists and list them all
    -lookup string
            lookup the value for a key in the property list dictionary
    -major
            bump major version
    -marketing-version
            print only marketing version
    -minor
            bump minor version
    -next-version
            only print next version value and dont bump version
    -patch
            bump patch version (default true)
    -plist string
            path to .plist file (default "Info.plist")
    -version
            print main Xcode project version in format 'marketingVersion (buildNumber)' and exit

## License

GPLv3