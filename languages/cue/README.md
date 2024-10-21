# CUE Semantic Reducer

## Pre-processing

The reducer SeRu can only work with one input file.  
To generate a single CUE file for a whole module, follow these steps:
> Will not work for custom commands! If the module contains *_tool.cue files, the pre-processing must be done manually.

1. Fix any issues with the module  
    ```shell
    cue mod fix
    ```
2. Inline all files into a new file (do not use the "-s" flag to simplify the output)
    ```shell
   cue def --inline-imports > in.cue
   ```
   
