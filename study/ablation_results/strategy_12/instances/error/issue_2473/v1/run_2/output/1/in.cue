package main

#ClusterInfo: env_type: string
#ClusterInfoByEnvName: {
 #env:     "stage" | "production"
 env_type: #env
}
#ProductionEndpoints: {
 #cluster_info: #ClusterInfo
 #cluster_info.env_type
}
{
 #cluster_info: #ClusterInfoByEnvName & {
  #env: "production"
 }
 #ProductionEndpoints
}
