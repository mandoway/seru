#env: "dev"
#ProductionEndpoints: {
 #cluster_info: env_type: string
 #cluster_info.env_type
}
{
 #cluster_info: {
  #env:     "dev" | "production"
  env_type: #env
 } & {
  #env: "production"
 }
 #ProductionEndpoints
}
