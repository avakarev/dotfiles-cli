# dotfiles-cli

[![Latest Release](https://img.shields.io/github/release/avakarev/dotfiles-cli.svg)](https://github.com/avakarev/dotfiles-cli/releases)
[![CI Status](https://github.com/avakarev/dotfiles-cli/actions/workflows/go.yml/badge.svg)](https://github.com/avakarev/dotfiles-cli/actions)
[![Go ReportCard](http://goreportcard.com/badge/avakarev/dotfiles-cli)](http://goreportcard.com/report/avakarev/dotfiles-cli)

CLI utility to manage dotfiles symlinks

## Installation

### Packages

#### Linux

[Packages](https://github.com/avakarev/dotfiles/releases) in Alpine, Debian & RPM formats

There is a one-liner to install `dotfiles` package using remote script:

    /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/avakarev/dotfiles-cli/master/scripts/install.sh)"

Running command from above will downloads a script and runs it.
Script might ask you for sudo password cause package managers like `apt` and `yum` require it to install the package.
It's a good idea to review everything you run on your machine, so please feel free to review the
[install script](./scripts/install.sh).

#### macOS

With [Homebrew](https://brew.sh/): `brew tap avakarev/tap && brew install dotfiles`

### Binaries

[Binaries](https://github.com/avakarev/dotfiles/releases) for Linux and macOS

### From source

Make sure you have a working Go environment (Go 1.12 or higher is required).
See the [install instructions](http://golang.org/doc/install.html).

Compiling `dotfiles` is easy, simply run:

    git clone https://github.com/avakarev/dotfiles-cli.git
    cd dotfiles-cli
    make build

## Usage

You can simply start with checking status of symlinks from configuration file:

    dotfiles status

To link dotfiles use `link` command:

    dotfiles link

To delete symlinks use `unlink` command:

    dotfiles unlink

Each command supports optional group(s) filter argument:

    dotfiles status group1 group2

## Configuration

`dotfiles` needs configuration file describing symlink groups.

File path could be specified either with flag `--config=/path/to/file` or using environment variable `$DOTFILES_CONFIG=/path/to/file`.
If none from options above is set, `dotfiles` will try to load `.dotfilesrc` from working directory.

`dotfiles` uses the following precedence order. Each item takes precedence over the item below it:

* `--config` / `-c` flag
* `DOTFILES_CONFIG` environment variable
* `.dotfilesrc` file in working directory

Configuration file format is `JSON` and could be either a flat list of symlinks (the group "default" is assigned implicitly)
or an object of groups each listing its own list.

Example of configuration file with no groups:

```json
[
  "vim",
  "vimrc"
]
```

Example of configuration file with groups:

```json
{
  "zsh": [
    "zsh",
    "zshrc",
  ],
  "git": [
    "gitconfig",
    "gitignore-global",
    "gitattributes-global"
  ]
}
```

Each symlink entry consists of source (mandatory) and target path (optional) separated by colon: `"source:target"`.
If source path just a filename, it's expected to be located in the same directory where the configuration file is.
If target part is missing, it's assumed as `"~/.<source>"`.

Both `source` and `target` path can include `~` or `$HOME`, which is interpreted as home directory path.

## License

`dotfiles` is licensed under MIT license. (see [LICENSE](./LICENSE))
