package mymod

#Service: [ID=string]: {...}

#Deployment: [ID=string]: {...}

#Service_Generator: {
	deployments: {...}
	out: {
		for deploymentName, deployment in deployments {
			_svc: #Service & {
				(deploymentName): {
					spec: {
						ports: [
							for container in deployment.spec.template.spec.containers
							if container.ports != _|_
							for portObj in container.ports
							if portObj.#no_service == _|_ {},
						]
					}
				}
			}
			for svcName, svc in _svc
			if len(svc.spec.ports) > 0 {
				(svcName): svc
			}
		}
	}
}
