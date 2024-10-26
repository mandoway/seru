# Study

This folder contains all files and scripts used to evaluate the tool under test: SeRu.

Directories:
- `results_perses`
  - SeRu+Perses
  - All instances
- `results_vulcan`
  - SeRu+Vulcan
  - Timeout for instance (after 2h)
    - 2246/v1
    - 2209/v1
  - All other instances



## Notes

### Issue 2490  
Initially not supported  
#### Multiple issues
- Multiple files
- Uses custom command
  - Only executable when using the _tool suffix
- -> Cue Inlining (`cue def --inline-imports`) DOES NOT WORK PROPERLY when renaming to `in.cue`
  - Produces new bug
  - Fixing the bug results in different output than expected
- Hard to modify the instance while keeping the original intent
- Current version of reducer does not support additional files, only input and test file

#### Fixed by inlining
Steps to inline the original module (line numbers reference original `in_tool.cue` file):
- Move code from `config/all.cue` to `in_tool.cue`
- Delete `cue.mod` directory
- Inline `config directory`
  - Remove package import `"github.com/cue-examples/cue-terraform-github-config-experiment/config"`
  - Remove selector `config` on line 21
    - `config.target.terraform.github.org`
    - -> `target.terraform.github.org`
- Replace file creation with cli print task
  - Replace import `tool/file` with `tool/cli` on line 5
  - Replace task `file.Create` with `cli.Print` on line 24
  - Remove `filename` property on line 25
  - Replace property name `contents` with `text` on line 26