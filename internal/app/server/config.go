package server

type Config struct {
	DbUrl string `toml:"database_url"`

	NatsURL   string `toml:"nats_url"`
	ClusterID string `toml:"cluster_id"`
	ClientID  string `toml:"client_id"`

	RedisIp string `toml:"redis_ip"`
}
