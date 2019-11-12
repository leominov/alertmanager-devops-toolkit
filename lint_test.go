package main

import "testing"

func TestCheckRouteReceiverIsDefined(t *testing.T) {
	a := &AlertmanagerConfig{
		RouteRoot: &RouteRoot{
			Routes: []*Route{
				{
					Receiver: "foobar",
				},
			},
		},
		Receivers: []*Receiver{
			{
				Name: "foobar",
			},
		},
	}
	errs := a.CheckRouteReceiverIsDefined()
	if len(errs) != 0 {
		t.Error("CheckRouteReceiverIsDefined() != 0")
	}
	a = &AlertmanagerConfig{
		RouteRoot: &RouteRoot{
			Routes: []*Route{
				{
					Receiver: "foobar",
				},
			},
		},
		Receivers: []*Receiver{
			{
				Name: "barfoo",
			},
		},
	}
	errs = a.CheckRouteReceiverIsDefined()
	if len(errs) != 1 {
		t.Error("CheckRouteReceiverIsDefined() != 1")
	}
}

func TestCheckSlackChannels(t *testing.T) {
	a := &AlertmanagerConfig{
		Receivers: []*Receiver{
			{
				Name: "foobar",
			},
		},
	}
	errs := a.CheckSlackChannels()
	if len(errs) != 0 {
		t.Error("CheckSlackChannels() != 0")
	}
	a = &AlertmanagerConfig{
		Receivers: []*Receiver{
			{
				Name: "foobar",
				SlackConfigs: []*SlackConfig{
					{
						Channel: "",
					},
				},
			},
		},
	}
	errs = a.CheckSlackChannels()
	if len(errs) != 0 {
		t.Error("CheckSlackChannels() != 0")
	}
	a = &AlertmanagerConfig{
		Receivers: []*Receiver{
			{
				Name: "foobar",
				SlackConfigs: []*SlackConfig{
					{
						Channel: "@l.aminov",
					},
				},
			},
		},
	}
	errs = a.CheckSlackChannels()
	if len(errs) != 0 {
		t.Error("CheckSlackChannels() != 0")
	}
	a = &AlertmanagerConfig{
		Receivers: []*Receiver{
			{
				Name: "foobar",
				SlackConfigs: []*SlackConfig{
					{
						Channel: "+l.aminov",
					},
				},
			},
		},
	}
	errs = a.CheckSlackChannels()
	if len(errs) != 1 {
		t.Error("CheckSlackChannels() != 1")
	}
}

func TestCheckEmptyReceivers(t *testing.T) {
	a := &AlertmanagerConfig{
		Receivers: []*Receiver{
			{
				Name: "foobar",
				SlackConfigs: []*SlackConfig{
					{
						Channel: "@l.aminov",
					},
				},
			},
		},
	}
	errs := a.CheckEmptyReceivers()
	if len(errs) != 0 {
		t.Error("CheckEmptyReceivers() != 0")
	}
	a = &AlertmanagerConfig{
		Receivers: []*Receiver{
			{
				Name: "foobar",
			},
		},
	}
	errs = a.CheckEmptyReceivers()
	if len(errs) != 1 {
		t.Error("CheckEmptyReceivers() != 1")
	}
}

func TestCheckRouteHasReceiver(t *testing.T) {
	a := &AlertmanagerConfig{
		RouteRoot: &RouteRoot{
			Routes: []*Route{
				{
					Receiver: "foobar",
				},
			},
		},
		Receivers: []*Receiver{
			{
				Name: "foobar",
			},
		},
	}
	errs := a.CheckRouteHasReceiver()
	if len(errs) != 0 {
		t.Error("CheckRouteHasReceiver() != 0")
	}
	a = &AlertmanagerConfig{
		RouteRoot: &RouteRoot{
			Routes: []*Route{
				{
					Receiver: "",
				},
			},
		},
		Receivers: []*Receiver{
			{
				Name: "foobar",
			},
		},
	}
	errs = a.CheckRouteHasReceiver()
	if len(errs) != 1 {
		t.Error("CheckRouteHasReceiver() != 1")
	}
}

func TestCheckSlackApiURL(t *testing.T) {
	a := &AlertmanagerConfig{
		Receivers: []*Receiver{
			{
				Name: "foobar",
				SlackConfigs: []*SlackConfig{
					{
						Channel: "@l.aminov",
					},
				},
			},
		},
	}
	errs := a.CheckSlackApiURL()
	if len(errs) != 0 {
		t.Error("CheckSlackApiURL() != 0")
	}
	a = &AlertmanagerConfig{
		Receivers: []*Receiver{
			{
				Name: "foobar",
				SlackConfigs: []*SlackConfig{
					{
						ApiURL: "https://google.com",
					},
				},
			},
		},
	}
	errs = a.CheckSlackApiURL()
	if len(errs) != 1 {
		t.Error("CheckSlackApiURL() != 1")
	}
}

func TestCheckDefaultReceiver(t *testing.T) {
	a := &AlertmanagerConfig{
		RouteRoot: &RouteRoot{
			Receiver: "foobar",
		},
		Receivers: []*Receiver{
			{
				Name: "foobar",
			},
		},
	}
	errs := a.CheckDefaultReceiver()
	if len(errs) != 0 {
		t.Error("CheckDefaultReceiver() != 0")
	}
	a = &AlertmanagerConfig{
		RouteRoot: &RouteRoot{
			Receiver: "barfoo",
		},
		Receivers: []*Receiver{
			{
				Name: "foobar",
			},
		},
	}
	errs = a.CheckDefaultReceiver()
	if len(errs) != 1 {
		t.Error("CheckDefaultReceiver() != 1")
	}
}

func TestCheckReceiverUniqueWebhookURL(t *testing.T) {
	a := &AlertmanagerConfig{
		Receivers: []*Receiver{
			{
				Name: "foobar",
				WebhookConfigs: []*WebhookConfig{
					{
						URL: "https://google.com",
					},
					{
						URL: "https://google.ru",
					},
				},
			},
		},
	}
	errs := a.CheckReceiverUniqueWebhookURL()
	if len(errs) != 0 {
		t.Error("CheckReceiverUniqueWebhookURL() != 0")
	}
	a = &AlertmanagerConfig{
		Receivers: []*Receiver{
			{
				Name: "foobar",
			},
		},
	}
	errs = a.CheckReceiverUniqueWebhookURL()
	if len(errs) != 0 {
		t.Error("CheckReceiverUniqueWebhookURL() != 0")
	}
	a = &AlertmanagerConfig{
		Receivers: []*Receiver{
			{
				Name: "foobar",
				WebhookConfigs: []*WebhookConfig{
					{
						URL: "https://google.com",
					},
					{
						URL: "https://google.com",
					},
				},
			},
		},
	}
	errs = a.CheckReceiverUniqueWebhookURL()
	if len(errs) != 1 {
		t.Error("CheckReceiverUniqueWebhookURL() != 1")
	}
}

func TestCheckReceiverUniqueSlackChannel(t *testing.T) {
	a := &AlertmanagerConfig{
		Receivers: []*Receiver{
			{
				Name: "foobar",
				SlackConfigs: []*SlackConfig{
					{
						Channel: "@l.aminov",
					},
					{
						Channel: "@m.aminova",
					},
				},
			},
		},
	}
	errs := a.CheckReceiverUniqueSlackChannel()
	if len(errs) != 0 {
		t.Error("CheckReceiverUniqueSlackChannel() != 0")
	}
	a = &AlertmanagerConfig{
		Receivers: []*Receiver{
			{
				Name: "foobar",
			},
		},
	}
	errs = a.CheckReceiverUniqueSlackChannel()
	if len(errs) != 0 {
		t.Error("CheckReceiverUniqueSlackChannel() != 0")
	}
	a = &AlertmanagerConfig{
		Receivers: []*Receiver{
			{
				Name: "foobar",
				SlackConfigs: []*SlackConfig{
					{
						Channel: "@l.aminov",
					},
					{
						Channel: "@l.aminov",
					},
				},
			},
		},
	}
	errs = a.CheckReceiverUniqueSlackChannel()
	if len(errs) != 1 {
		t.Error("CheckReceiverUniqueSlackChannel() != 1")
	}
}

func TestCheckReceiverUniqueEmailTo(t *testing.T) {
	a := &AlertmanagerConfig{
		Receivers: []*Receiver{
			{
				Name: "foobar",
				EmailConfigs: []*EmailConfig{
					{
						To: "foobar@gmail.com",
					},
					{
						To: "barfoo@gmail.com",
					},
				},
			},
		},
	}
	errs := a.CheckReceiverUniqueEmailTo()
	if len(errs) != 0 {
		t.Error("CheckReceiverUniqueEmailTo() != 0")
	}
	a = &AlertmanagerConfig{
		Receivers: []*Receiver{
			{
				Name: "foobar",
			},
		},
	}
	errs = a.CheckReceiverUniqueEmailTo()
	if len(errs) != 0 {
		t.Error("CheckReceiverUniqueEmailTo() != 0")
	}
	a = &AlertmanagerConfig{
		Receivers: []*Receiver{
			{
				Name: "foobar",
				EmailConfigs: []*EmailConfig{
					{
						To: "foobar@gmail.com",
					},
					{
						To: "foobar@gmail.com",
					},
				},
			},
		},
	}
	errs = a.CheckReceiverUniqueEmailTo()
	if len(errs) != 1 {
		t.Error("CheckReceiverUniqueEmailTo() != 1")
	}
}

func TestList(t *testing.T) {
	a := &AlertmanagerConfig{
		RouteRoot: &RouteRoot{},
	}
	a.Lint()
}
