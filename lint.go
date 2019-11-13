package main

func Lint(a *AlertmanagerConfig) []error {
	var errs []error

	errs = append(errs, CheckRouteHasReceiver(a)...)
	errs = append(errs, CheckRouteReceiverIsDefined(a)...)
	errs = append(errs, CheckEmptyReceivers(a)...)
	errs = append(errs, CheckReceiverSlackChannels(a)...)
	errs = append(errs, CheckReceiverHasSlackApiURL(a)...)
	errs = append(errs, CheckReceiverUniqueSlackChannel(a)...)
	errs = append(errs, CheckReceiverWebhookURLs(a)...)
	errs = append(errs, CheckReceiverUniqueWebhookURL(a)...)
	errs = append(errs, CheckReceiverEmailTo(a)...)
	errs = append(errs, CheckReceiverUniqueEmailTo(a)...)
	errs = append(errs, CheckReceiverSlackHttpConfigProxyURL(a)...)
	errs = append(errs, CheckReceiverWebhookHttpConfigProxyURL(a)...)
	errs = append(errs, CheckDefaultReceiver(a)...)

	return errs
}
