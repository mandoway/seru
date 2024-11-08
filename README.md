# SeRu

SeRu is a framework and tool for automatic program reduction.  
It uses a combination of syntactic and semantic reduction strategies to combine speed and effectiveness.

SeRu is designed to support any syntactic reduction command and any language.

# Usage

SeRu currently supports:

- Syntactic reducers
    - Perses
    - Vulcan
- Semantic reducers for languages
    - CUE

## Pre-requisites

> Steps 1 & 2 are only necessary while the repository are private  
> SeRu will download such dependencies automatically

1. Download Perses & CUE-specific extension (latest from the release page)
    1. [perses_deploy.jar](https://github.com/mandoway/seru/releases/download/v0.0.1-alpha/perses_deploy.jar)
    2. [cue.jar](https://github.com/mandoway/seru/releases/download/v0.0.1-alpha/cue.jar)
2. Compile CUE-specific semantic reduction strategies
    1. Run
   ```bash
   go generate ./...
   ```
3. Make sure you have Java installed (for Perses)

## Run seru
Run SeRu using Go or standalone  
- `go run .`
- `go build && ./seru`
- Install from the release page

## Command line options

| Option                       | Shorthand | Required | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              |
|------------------------------|-----------|----------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `--input <path>`             | `-i`      | ✅        | path to input file                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       |
| `--test <path>`              | `-t`      | ✅        | path to property testscript, a test script **must return 0** when the property holds and 1 (or any code) if the property is invalid                                                                                                                                                                                                                                                                                                                                                                                                      |
| `--lang <string>`            | `-l`      |          | language of file, e.g. cue. Will be inferred from the file extension if omitted.                                                                                                                                                                                                                                                                                                                                                                                                                                                         |
| `--reducer <perses\|vulcan>` | `-r`      |          | either 'perses' OR 'vulcan' (default "perses"). Perses is faster, Vulcan can be more effective.                                                                                                                                                                                                                                                                                                                                                                                                                                          |
| `--enable-metrics`           | `-m`      |          | store metrics as a json file                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             |
| `--strategy-isolation`       | `-s`      |          | (*untested*) Activates strategy isolation.<br/> Modes: <br/>- Isolation (many candidates, slow): only one strategy will be applied before next iteration. This mode will apply only one semantic strategy and try to reduce all returned candidates using the syntactic reducer. <br/>- Strategy combination (**default**, much faster, less calls to syntactic reducer). In strategy combination, all strategies are applied and combined to one "best candidate". Then this single candidate will be reduced by the syntactic reducer. |
| `--active-strategies <ints>` |           |          | list of indices of active strategies, enter -1 to disable semantic reduction                                                                                                                                                                                                                                                                                                                                                                                                                                                             |

# Future plans

- Tooling to help create semantic reducers
- More languages
