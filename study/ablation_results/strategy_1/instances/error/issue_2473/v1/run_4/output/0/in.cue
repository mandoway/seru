package main

#ClusterInfo: {
	env_type: string
}

#ClusterInfoByEnvName: #ClusterInfo & {
	#env:     "dev" | "stage" | "production"
	env_type: #env
}

#ProductionEndpoints: {
	#cluster_info: #ClusterInfo

	example_host: "\(#cluster_info.env_type).example.net"
}

endpoints: {
	#cluster_info: #ClusterInfoByEnvName & {
		#env: "production"
	}

	if #cluster_info.env_type == "production" {
		#ProductionEndpoints
	}
}
