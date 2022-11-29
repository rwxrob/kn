# KEG Node Utility

***💥 DEPRECATED: This code and utility has been replace by [`keg`](https://github.com/rwxrob/keg). It is kept around for historical purposes and to ensure everything in it has been implemented in its replacement, which is a composable, stand-alone stateful command branch as opposed to a Perl script with obvious problematic dependencies.***

## Action Summary

```
kn epoch [<mod>] - print seconds since Jan 1, 1970
kn config        - config crud
kn html          - convert stdin to markdown

kn cd            - change into node directory
kn jax (change to snip)
kn snip          - snippets crud
kn ren[der]      - render the entire root node

kn day  [<sep>] [<mod>] - 20210321
kn time [<sep>] [<mod>] - 165603
kn date [<fmt>] [<mod>] - date unix command + go date formatting
kn tstamp        - print RFC3339 time stamp

kn find 

kn log           - log crud
kn link          - save and manage links, URLs (bookmarks)

kn new           - create a new root local node
```

## Implementation Plan

Getting a good prototype created and available quickly for alpha testing
with a broad group is therefore the initial goal. This will allow rapid,
organic enhancements ensuring maximum usability before implementing the
main release.

Prototype development on `kn` is in Shell and Perl. Final version will
be in Go.

## Extension Support 

Extensions use the UNIX philosophy to allow them to be written in any
language that will execute on the target operating system.

## Environment Variables

The `kn` tool does not depend on any environment variables itself
(although extensions might have their own dependency). Instead, it uses
the [`conf-go`](https://github.com/rwxrob/conf-go) package with defaults
(which creates matching configuration files under `~/.local/config`).

*[Note that `KN` and `KNPATH` were once used, but no longer.]*
