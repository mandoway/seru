string
#ClusterInfoByEnvName: {
 #env:     "dev" | "production"
 env_type: #env
}
#ProductionEndpoints: {
 #cluster_info: #ClusterInfoByEnvName
 #cluster_info.env_type
}
#cluster_info: #ClusterInfoByEnvName & {
 #env: "production"
}
#ProductionEndpoints
