resource "cloudflare_zone" "this" {
  account_id = "bd39ee208bf196f83cb046bbdfd4cbe6"
  zone       = "seancheng.space"
}

resource "cloudflare_record" "prod" {
  name    = "ekko"
  type    = "A"
  zone_id = cloudflare_zone.this.id
  value   = var.ip
  proxied = true
}
