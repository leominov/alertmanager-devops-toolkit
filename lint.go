package main

var (
	checks = map[string]Check{}
)

type (
	Check func(a *AlertmanagerConfig) []error
)

func RegisterCheck(name string, fn Check) {
	checks[name] = fn
}

func Lint(a *AlertmanagerConfig) []error {
	var errs []error
	for _, fn := range checks {
		errs = append(errs, fn(a)...)
	}
	return errs
}
