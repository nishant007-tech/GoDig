package plugin

import "fmt"

// AnalyticsPlugin implements Plugin.
type AnalyticsPlugin struct{}

func NewAnalyticsPlugin() Plugin {
	return &AnalyticsPlugin{}
}

func (a *AnalyticsPlugin) Name() string { return "analytics" }

func (a *AnalyticsPlugin) Init() error {
	fmt.Println("Initializing Analytics Plugin")
	return nil
}
