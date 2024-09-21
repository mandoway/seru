# SeRu

SeRu is a tool for automatic program reduction.  
It uses a combination of syntactic and semantic reduction strategies to combine speed and effectiveness.

SeRu is designed to support any semantic reduction command and any language.

# Usage

SeRu currently supports:
- Syntactic reducers
  - Perses
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
4. Run SeRu using Go or standalone
   1. `go run .`
   2. `go build && ./seru`
   3. Command line options:
      1. `-i <input>` _Required_  
      Target file to reduce
      2. `-t <test>` _Required_  
      Test script checking if the reduced file still kept the required property  
      A test script **must return 0** when the property was kept and 1 (or any code) if the property was lost
      3. `-l <language>`  
      Programming language of the input file. Will be inferred from the file extension if omitted.
      4. `-s`  
      Use strategy isolation.  
      This mode will apply only one semantic strategy and try to reduce all returned candidates using the syntactic reducer.  
      Default mode: strategy combination  
      In strategy combination, all strategies are applied and combined to one "best candidate". Then this one candidate will be reduced by the syntactic reducer.

# Future plans

- Tooling to help create semantic reducers
- More languages
