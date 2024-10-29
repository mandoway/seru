package #x

output: {
 _objects: Deployment: DEPLOYMENT & #x
 SERVICE
 for #x in _generated
 for #x in #x {
  #x
 }
 DeploymentsWithRef: deployments: _objects.Deployment
 _generated: [(SERVICE_GENERATOR.#x & DeploymentsWithRef).#x]
}
#x: redis:
{}
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
  #x: {
   for deploymentName, #x in deployments {
    _svc_AC73C747: deploymentName:
     #x
    for #x , #x in _svc_AC73C747 {
     (deploymentName): #x
    }
   }
  }
 }
}
