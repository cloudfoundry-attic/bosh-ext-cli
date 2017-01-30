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

- `set -e` in packaging?
- pre_packaging presence?
- job description presence
- `type: password` annotation
- extrapolate links usage
- discover_external_ip presence
- logrotate presence
- todo markers
- consolidate explanation
