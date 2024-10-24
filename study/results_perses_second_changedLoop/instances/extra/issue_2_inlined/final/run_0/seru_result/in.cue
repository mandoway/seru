package redis

output: {
 _objects: Deployment: redis: spec: template: spec: containers: {}
 for objectSet in _generated
 for object in objectSet {}
 DeploymentsWithRef: deployments: _objects.Deployment
 _generated: [({
  #x: {
   deployments: {
    ...
   }
   out: {
    for deploymentName, deployment in deployments {
     _svc_AC73C747: (deploymentName): spec: ports:
      deployment
     for svcName, svc in _svc_AC73C747 {
      (svcName): svc
     }
    }
   }
  }
 }.#x & DeploymentsWithRef).out]
}
...
