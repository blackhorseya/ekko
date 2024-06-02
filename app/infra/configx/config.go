package configx

// Configuration defines the configuration struct.
type Configuration struct {
	PlatformRest    Application `json:"platform_rest" yaml:"platformRest"`
	PlatformLinebot Application `json:"platform_linebot" yaml:"platformLinebot"`
}
