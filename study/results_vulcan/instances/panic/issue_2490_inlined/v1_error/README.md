These runs resulted in crashes in the semantic strategy "Import".

Vulcan swaps identifiers randomly and changed the package declaration to an imported identifier.
When the strategy tries to remove an import, it also modified the package declaration to an invalid state.
This bug was fixed by ignoring package declaration for this specific strategy.