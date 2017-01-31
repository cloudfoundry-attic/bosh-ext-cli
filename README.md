# bosh-lint

bosh-lint is a tool that gives suggestions for BOSH releases and other assets.

## Development

```
$ source .envrc
$ ginkgo -r src/github.com/cppforlife/bosh-lint/
$ bin/build
$ out/bosh-lint lint-release --dir ~/workspace/whatever-release
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
