variable "token" {
  type    = string
  default = getenv("TURSO_TOKEN")
}

env "local" {
  url     = "libsql+wss://gigachad-sahidrahman404.turso.io?authToken=${var.token}"
  exclude = ["_litestream*", "stories_fts*"]
}
