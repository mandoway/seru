package redis

output: {
 {}
 for objectSet in _generated
 for object in objectSet {}
 DeploymentsWithRef: deployments:
 {}
 _generated: [{
  deployments
  out: {
   for deploymentName, deployment in deployments {
    _svc_AC73C747: (deploymentName): spec: ports:
     deployment
    for svcName, svc in _svc_AC73C747 {
     (svcName): svc
    }
   }
  }
  deployments: redis: spec: template: spec: containers: {}
 }.out]
}
{}
{}
{}
...
