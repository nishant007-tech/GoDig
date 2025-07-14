package plugin

import "fmt"

// PaymentPlugin implements Plugin.
type PaymentPlugin struct{}

func NewPaymentPlugin() Plugin {
	return &PaymentPlugin{}
}

func (p *PaymentPlugin) Name() string { return "payment" }
func (p *PaymentPlugin) Init() error {
	fmt.Println("Initializing Payment Plugin")
	return nil
}
