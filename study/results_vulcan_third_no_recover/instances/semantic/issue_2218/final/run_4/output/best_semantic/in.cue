_input: {
	v: {}
	k: _input: true
}
for k, _input in {
	v: {}
	k: _input: true
} {
	let v = _input._input
	(k): v
}
