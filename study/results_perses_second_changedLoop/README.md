These results were generated with an old version of SeRu.

Two main differences to the current version:

- Reduction loop used syntactic reduction first, then semantic reduction
    - Syntactic with no candidates -> semantic -> syntactic with new candidates
    - A new semantic candidate can be rejected when the syntactic reducer crashes for some reason.
    - Therefore, some candidates are lost.
    - in the new version, a run with the syntactic reducer is performed before the loop. Then semantic reducer generates
      candidates which are used in the syntactic reducer. After that, we check for a fixpoint.
- Constant propagation had unlimited recursion
  - If a field mentions itself in a nested struct, constant propagation would infinitely propagate this field on itself.
