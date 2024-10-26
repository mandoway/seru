package redis

output: {
 for objectSet in _generated
 for object in objectSet {}
 DeploymentsWithRef: deployments:
 {}
 _generated: [
  {
   for deploymentName, deployment in {
    redis: spec: template: spec: containers: {}
   } {
    _svc_AC73C747: (deploymentName): spec: ports:
     deployment
    for svcName, svc in _svc_AC73C747 {
     (svcName): svc
    }
   }
  },
 ]
}
...
