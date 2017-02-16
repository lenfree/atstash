## 0.0.1 (2017-02-02)

Initial release

## 0.1.0 (2017-02-06)

### Added

- Add Stash PR
- Add Git push

## 0.1.1 (2017-02-07)

Bugfix release

### Fixed

- Catch git push up-to-date error
- Catch PR exists error

## 0.1.2 (2017-02-07) 

Bugfix release 

### Fixed

- Fix hardcoded target project key for Stash PR
- Fix branch name from refs


## 0.1.3 (2017-02-08)

This is a working version that only make a Stash PR
and get rid of git push which is not the purpose of
this tool and makes a Slack Post request to a group

### Added

- Replace push with pr since this only does a PR
- Use git2go package to query git branch and remotes
- Add Webhook to Slack PR

### Removed

- Remove git push to keep it simple

## 0.1.4 (2017-02-16)

Fix bug where Stash PR description with commit message

### Fixed

- Empty description when creating PR