package redis

import (
	"example.com/mymod"
)

output: {

	_objects: {
		Deployment: mymod.#Deployment & deployment
		Service:    mymod.#Service & service
	}

	_localReferences: {}

	objects: [
		for objectSet in _generated
		for object in objectSet {object},
	]

	DeploymentsWithRef: {
		deployments: {
			[string]: #config: ref: _localReferences
		} & _objects.Deployment
	}

	_generated: [
		( mymod.#Service_Generator & DeploymentsWithRef).out,
	]

}

service: {}

deployment: redis: spec: template: spec: containers: [{}, {ports: [{}]}]

