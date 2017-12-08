package models

// All defined config types
var Models = map[string]interface{}{
	"child":            ChildModel{}, // testing only
	"Develop.mr_robot": DevelopMrRobotConfig{},
	"Test.vpn":         TestVPNConfig{},
}

// Simple model used only for testing purposes
type ChildModel struct {
	Model
	Magic int
}

// Config type 1 -- Develop.mr_robot
type DevelopMrRobotConfig struct {
	Model

	Host     string
	Port     int
	Database string
	User     string
	Password string
	Schema   string
}

// Config type 2 -- Test.vpn
type TestVPNConfig struct {
	Model

	Host        string
	Port        int
	Virtualhost string
	User        string
	Password    string
}
