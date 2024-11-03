_parent_configs: {
	PARENT: true
}

_configs: {
	CHILD1: {
		property: true
	}
	CHILD2: {
		property: true
		parent:      "PARENT"
	}
}


disabled_parent_test: {for k, v in _configs {
	let parent_config = (*_parent_configs[v.parent] | false)
    "\(k)": {
        "parent_config": parent_config
    }
}}
