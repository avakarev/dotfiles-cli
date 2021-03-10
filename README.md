# dotfiles-cli

CLI utility to manage dotfiles symlinks

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

```
{
  "vim",
  "vimrc"
}
```

Example of configuration file with groups:

```
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

Both `source` and `target` path can include `~` or `$HOME`, which interpreted as home directory path.
