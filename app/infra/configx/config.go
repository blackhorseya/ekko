package configx

// Configuration defines the configuration struct.
type Configuration struct {
	PlatformRest Application `json:"platform_rest" yaml:"platformRest"`
	Linebot      Application `json:"linebot" yaml:"linebot"`
}
