# atstash
=========

A cli tool that help large team to create a PR and include all members
automatically. Unfortunately, Atlassian Stash/Bitbucket does not support
groups. Inspired by Stashify, however, this is suppose to just be simple
cli.

Build Status: [![Build Status](https://travis-ci.org/lenfree/atstash.svg?branch=master)](https://travis-ci.org/lenfree/atstash)

[Binary Releases](https://github.com/lenfree/atstash/releases)

## Usage

```bash
$ cat >>.env<<EOF
STASH_USERNAME="<change_me>"
STASH_PASSWORD="<change_me>"
STASH_URL="<stash_url>"
ORIGIN_REPO_NAME="remote_name"
PROJECT_KEY="project_key"
REPO_KEY="repo_name"
FORKED_REPO_NAME="forked_remote_name"
REVIEWERS="<username>"
EOF

$ atstash --help
```
## Install

To install, use `go get`:

```bash
$ go get -d github.com/lenfree/atstash
```

or 

```bash
$ curl -o <packagename> -L https://github.com/lenfree/atstash/releases/download/<version>/atstash-<darwin|linux>-amd64
$ <path_to_binary> --help
```

## Contribution

1. Fork ([https://github.com/lenfree/atstash/fork](https://github.com/lenfree/atstash/fork))
2. Create a feature branch
3. Commit your changes
4. Rebase your local changes against the master branch
5. TODO: Run test suite with the `go test ./...` command and confirm that it passes
6. Run `gofmt -s`
7. Create a new Pull Request

## Author

[lenfree](https://github.com/lenfree)
