# PLIST WIZARD

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

## Future

- implement flagVersion
- bump multiple plist files in one project to the same version
- recursively find all plists in dir (listAll flag)
- write / add fields
- format XML

## Help

    $ plistwizard -h
    plistwizard [-plist <path/to/plist>] [-next-version] [-bump-version] [-lookup <key>]

    Usage of plistwizard:
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
    -lookup string
            lookup the value for a key in the property list dictionary
    -major
            bump major version
    -minor
            bump minor version
    -next-version
            only print next version value and dont bump version
    -patch
            bump patch version (default true)
    -plist string
            path to .plist file (default "Info.plist")

## Examples

Show a table with elements of the data dictionary:

    plistwizard -plist Project/Info.plist

Lookup a value for the specified key inside the dictionary inside the plist:

    plistwizard -plist Project/Info.plist -lookup LSApplicationCategoryType

Show the next version:

    plistwizard -plist Project/Info.plist -next-version

Bump the major version number:

    plistwizard -plist Project/Info.plist -bump-version -major

Bump the minor version number:

    plistwizard -plist Project/Info.plist -bump-version -minor

Bump the patch version number (default!):

    plistwizard -plist Project/Info.plist -bump-version -patch

## License

GPLv3