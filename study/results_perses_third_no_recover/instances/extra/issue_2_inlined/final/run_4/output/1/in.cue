package redis

output: {
 _objects: {
  Deployment: DEPLOYMENT & deployment
  SERVICE
 }
 {}
 for objectSet in _generated
 for object in objectSet {
  object
 }
 DeploymentsWithRef: deployments: _objects.Deployment
 _generated: [(SERVICE_GENERATOR.#x & DeploymentsWithRef).out]
}
{}
deployment: redis: spec: template: spec: containers: {}
let DEPLOYMENT =
{
 ...
}
let SERVICE =
{
 ...
}
let SERVICE_GENERATOR = {
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
}
