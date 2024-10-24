#ClusterInfo: env_type: env_type
#ClusterInfoByEnvName: {
 #env:     "dev" | "production"
 env_type: #env
}
#ProductionEndpoints: #ClusterInfoByEnvName: #ClusterInfo
env_type: "\(#cluster_info.env_type).example.net"
#cluster_info: #ClusterInfoByEnvName & {
 #env: "production"
}
if #cluster_info.env_type == "production" {
 #ProductionEndpoints
}
