#ClusterInfo: env_type: string
#env: "dev"
#ProductionEndpoints: {
 #cluster_info: #ClusterInfo
 example_host:  "\(#cluster_info.env_type).example.net"
}
{
 #cluster_info: {
  #env:     "dev" | "production"
  env_type: #env
 } & {
  #env: "production"
 }
 if #cluster_info.env_type == "production" {
  #ProductionEndpoints
 }
}
