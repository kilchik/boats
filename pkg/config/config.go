package config

type Desc struct {
	EnableDebugLogs bool   `toml:"enable_debug_logs"`
	DSN             string `toml:"dsn"`
	NausysAddr      string `toml:"nausys_addr"`
	NausysUser      string `toml:"nausys_user"`
	NausysPass      string `toml:"nausys_pass"`
}
