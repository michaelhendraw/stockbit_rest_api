package config

// AppConfig struct
type AppConfig struct {
	Server ServerConfig
	GRPC   GRPCConfig
}

// ServerConfig struct
type ServerConfig struct {
	Port  string
	Env   string
	Debug int
}

// GRPCConfig struct
type GRPCConfig struct {
	StockbitEndpoint string
}
