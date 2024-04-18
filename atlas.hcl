variable "dsn" {
  type    = string
  default = getenv("DB_DSN")
}

env "local" {
  url     = "${var.dsn}"
  exclude = ["_litestream*", "stories_fts*"]
}
