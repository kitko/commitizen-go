# Commitizen-go
> Command line utility to standardize git commit messages, golang version.

Support git hooks

![demo](docs/images/demo.gif)

Contents
- Getting Started
- Usage
- Configure

## Getting Started
### installation with source code:
```
$ make && make install
```
or 
```
$ make && ./commitizen-go install
```
### commit with commitizen
```
$ git cz
```

## Install and run as git hook

### Checkout repo

```
$ git clone ....
```

### Install

```
$ go install
```

### Setup as git hook, and push in repo

Copy folder `.githooks` in your project folder

```
$ cp .githooks /path/to/project
```
Switch to `/path/to/project`

```
$ cd /path/to/project
```
Make file `.githooks/prepare-commit-msg` executable if it's not

```
$ chmod +x .githooks/prepare-commit-msg
```

Setup git config to use `.githooks` folder

```
$ git config core.hooksPath .githooks
```
Turn off default editor (witch is replaced by this tool)

```
git config core.editor :
```
Now you can use it like this :)

```
$ git commit
```

### Setup as git hook, local in .git folder

Copy file `.githooks/prepare-commit-msg` in your project `.git/hooks` folder

```
$ cp .githooks/prepare-commit-msg /path/to/project/.git/hooks
```
Switch to `/path/to/project`

```
$ cd /path/to/project
```
Make file `.git/hooks/prepare-commit-msg` executable if it's not

```
$ chmod +x .git/hooks/prepare-commit-msg
```
Turn off default editor (witch is replaced by this tool)

```
git config core.editor :
```
Now you can use it like this :)

```
$ git commit
```


## Usage
```

Usage:
  commitizen-go [flags]
  commitizen-go [command]

Available Commands:
  help        Help about any command
  install     Install this tool to git-core as git-cz
  version     Print version information and quit

Flags:
  -a, --all     tell the command to automatically stage files that have been modified and deleted, but new files you have not told Git about are not affected
  -d, --debug   debug mode, output debug info to debug.log
  -h, --help    help for commitizen-go
      --hook    tell the command to save message to .git/COMMIT_EDITMSG file

Use "commitizen-go [command] --help" for more information about a command.
```

## Configure
You can customize your commit type, items, and how to assemble these items through the configuration file.
Configuration file path is `~/.git-czrc`, and the format is the same as the defaultConfig string in the file commit/defaultConfig.go.
Type item like that:
```
...
{
	"name": "type",
	"desc": "Select the type of change that you're committing:",
	"form": "select",
	"options": [
		{ "name": "feat", "desc": "feat: A new feature" },
		{ "name": "fix", "desc": "fix: A bug fix" },
		{ "name": "docs", "desc": "docs: Documentation only changes" },
		{
		  "name": "style",
		  "desc":
			"style: Changes that do not affect the meaning of the code\n            (white-space, formatting, missing semi-colons, etc)"
		},
		{
		  "name": "refactor",
		  "desc": "refactor: A code change that neither fixes a bug nor adds a feature"
		},
		{
		  "name": "perf",
		  "desc": "perf: A code change that improves performance"
		},
		{ "name": "test", "desc": "test: Adding missing tests" },
		{
		  "name": "chore",
		  "desc":
			"chore: Changes to the build process or auxiliary tools\n            and libraries such as documentation generation"
		},
		{ "name": "revert", "desc": "revert: Revert to a commit" },
		{ "name": "WIP", "desc": "WIP: Work in progress" }
	],
	"required": true
},
...
```
Template like that:
```
...
"template": "{{.type}}{{with .scope}}({{.}}){{end}}: {{.subject}}{{with .body}}\n\n{{.}}{{end}}{{with .footer}}\n\n{{.}}{{end}}"
...
```
