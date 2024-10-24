## Vulcan issues

- Issue 2584 v2 timeout after 2h for one run  
Seems to generate the same output but does not terminate anyway

- Issue 2490 inlined strategy 8 too many candidates

Cause: Constant propagation and recursive definitions
Fix: limit recursive constant propagation

- Issue 2246 timeout after 2 hours
- 2209 timeout