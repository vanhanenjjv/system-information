package sysinfo

type NetworkDevice struct {
	Interface string `json:"interface"`
	Address   string `json:"address"`
}

type SystemInformation struct {
	SystemLoad         *float64        `json:"system_load"`
	MemoryUsage        *int64          `json:"memory_usage"`
	SwapUsage          *int64          `json:"swap_usage"`
	ProcessCount       *int64          `json:"processes"`
	UsersLoggedInCount *int64          `json:"users_logged_in"`
	NetworkDevices     []NetworkDevice `json:"network_devices"`
}
