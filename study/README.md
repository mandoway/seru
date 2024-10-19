# Study

This folder contains all files and scripts used to evaluate the tool under test: SeRu.


Issue 2490  
Not supported  
Multiple issues:
- Multiple files -> can be inlined
- Uses custom command
  - Only executable when using the _tool suffix
  - Also not really a problem
- -> Inlining DOES NOT WORK PROPERLY
  - Produces new bug
  - Fixing the bug results in different output than expected
- Hard to modify the instance while keeping the original intent
- Current version of reducer does not support additional files, only input and test file