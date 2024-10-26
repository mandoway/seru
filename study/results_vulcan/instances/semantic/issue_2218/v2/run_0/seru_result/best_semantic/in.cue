_configs: {
	v: {}
	k: _configs: true
}
for k, _configs in {
	v: {}
	k: _configs: true
} {
	let v = _configs._configs
	"\(k)": v
}
