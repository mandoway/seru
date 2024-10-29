c: {
	env:  "prod"
	env:  "dev" | "prod"
	env2: env
}

if c.env2 == "prod" {
	{
		c: env2: string
		example: c.env2
	}
}
