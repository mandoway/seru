#ClusterInfo: env_type: string
#env:          "dev"
#cluster_info: #ClusterInfo
{
 #cluster_info: {
  #env:     "dev" | "production"
  env_type: #env
 } & {
  #env: "production"
 }
 if #cluster_info.env_type == "production" {
  {
   #cluster_info: #ClusterInfo
   example_host:  "\(#cluster_info.env_type).example.net"
  }
 }
}
