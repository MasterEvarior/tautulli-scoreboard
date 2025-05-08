# Gize

![quality workflow](https://github.com/MasterEvarior/tautulli-scoreboard/actions/workflows/quality.yaml/badge.svg) ![release workflow](https://github.com/MasterEvarior/tautulli-scoreboard/actions/workflows/publish.yaml/badge.svg)

TODO

## Development

### Linting

Linting is done with [golangci-lint](https://golangci-lint.run/), which can be run like so:

```shell
golangci-lint run
```

Run all other linters with the treefmt command. Note that the command does not install the required formatters.

```shell
treefmt
```

### Git Hooks

There are some hooks for formatting and the like. To use those, execute the following command:

```shell
git config --local core.hooksPath .githooks/
```

### Nix

If you are using [NixOS or the Nix package manager](https://nixos.org/), there is a dev shell available for your convenience. This will install Go, everything needed for formatting, set the Git hooks and some default environment variables. Start it with this command:

```shell
nix develop
```

If you happen to use [nix-direnv](https://github.com/nix-community/nix-direnv), this is also supported.

## Improvements, issues and more

Pull requests, improvements and issues are always welcome.
