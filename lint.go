package main

func (a *AlertmanagerConfig) Lint() []error {
	var errs []error

	errs = append(errs, CheckRouteHasReceiver(a)...)
	errs = append(errs, CheckRouteReceiverIsDefined(a)...)
	errs = append(errs, CheckEmptyReceivers(a)...)
	errs = append(errs, CheckSlackChannels(a)...)
	errs = append(errs, CheckSlackApiURL(a)...)
	errs = append(errs, CheckReceiverUniqueSlackChannel(a)...)
	errs = append(errs, CheckWebhookURLs(a)...)
	errs = append(errs, CheckReceiverUniqueWebhookURL(a)...)
	errs = append(errs, CheckEmailTo(a)...)
	errs = append(errs, CheckReceiverUniqueEmailTo(a)...)
	errs = append(errs, CheckSlackHttpConfigProxyURL(a)...)
	errs = append(errs, CheckWebhookHttpConfigProxyURL(a)...)
	errs = append(errs, CheckDefaultReceiver(a)...)

	return errs
}
