package redis

output: {
	_objects: {
		Deployment: DEPLOYMENT.#x & deployment
		Service:    SERVICE.#x & service
	}
	_localReferences: {}
	objects: [
		for objectSet in _generated
		for object in objectSet {
			object
		}]
	DeploymentsWithRef: deployments: _objects.Deployment & {
		{
			[string]: #config: {
				ref: _localReferences
			}
		}
	}
	_generated: [(SERVICE_GENERATOR.#x & DeploymentsWithRef).out]
}
service: {}
deployment: redis: spec: template: spec: containers: [{}, {
	ports: [{}]
}]

//cue:path: "example.com/mymod".#Deployment
let DEPLOYMENT = {
	#x: {
		[ID=string]: {
			...
		}
	}
}

//cue:path: "example.com/mymod".#Service
let SERVICE = {
	#x: {
		[ID=string]: {
			...
		}
	}
}

//cue:path: "example.com/mymod".#Service_Generator
let SERVICE_GENERATOR = {
	#x: {
		deployments: {
			...
		}
		out: {
			for deploymentName, deployment in deployments {
				_svc_AC73C747: SERVICE.#x & {
					(deploymentName): spec: ports: [
						for container in deployment.spec.template.spec.containers
						if container.ports != _|_ // explicit error (_|_ literal) in source
						for portObj in container.ports
						if portObj.#no_service == _|_ // explicit error (_|_ literal) in source
						{}]
				}
				for svcName, svc in _svc_AC73C747
				if len(svc.spec.ports) > 0 {
					(svcName): svc
				}
			}
		}
	}
}
