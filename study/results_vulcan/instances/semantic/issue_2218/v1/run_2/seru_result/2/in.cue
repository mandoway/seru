_parent_configs: PARENT: true
for k, _configs in {
 k: true
 PARENT: _configs: "PARENT"
} {
 let PARENT = _parent_configs[_configs._configs]
 "\(k)":
  PARENT
}
