_parent_configs: PARENT: true
_configs: {
	v: property:        true
	property: _configs: "PARENT"
}
for k, _configs in {
	v: property:        true
	property: _configs: "PARENT"
} {
	let property = _parent_configs[_configs._configs]
	"\(k)":
		property
}
