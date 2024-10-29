c: {
 env:  "prod"
 env2: "dev" | "prod"
 env2: env
}
if c.env2 == "prod" {
 {
  c: env2: env2
  c
 }
}
