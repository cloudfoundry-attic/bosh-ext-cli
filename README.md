# bosh-lint

bosh-lint is a tool that gives suggestions for BOSH releases and other assets.

## Commands

```
# Lint release directory and show suggestions
$ bosh-lint lint-release

# Show detailed Director task debug information
$ bosh task X --debug | bosh-lint debug-task -
$ bosh task X --debug | bosh-lint debug-task - -a
$ bosh task X --debug | bosh-lint debug-task - -a -s duration
$ bosh task X --debug | bosh-lint debug-task - -l
```

## Development

```
$ source .envrc
$ bin/build
$ out/bosh-lint lint-release --dir ~/workspace/whatever-release
```

Run tests:

```
$ bin/test
```

## Todo

- release: notice common props between release jobs
- release: multiple jobs?
- release: camelCase vs snake case
- `set -e` in packaging?
- pre_packaging presence?
- job description presence
- `type: password` annotation
- extrapolate links usage
- discover_external_ip presence
- logrotate presence
- todo markers
- consolidate explanation
- greedy blob inclusion?
