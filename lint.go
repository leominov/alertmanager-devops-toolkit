package main

var (
	checks              = map[string]Check{}
	defaultCheckOptions = &CheckOptions{
		Active: true,
	}
)

type (
	Check func(a *AlertmanagerConfig, opt *CheckOptions) []error
)

type CheckOptions struct {
	Active bool   `yaml:"active"`
	Level  string `yaml:"level"`
}

func RegisterCheck(name string, fn Check) {
	checks[name] = fn
}

func Lint(a *AlertmanagerConfig, config *Config) []error {
	var errs []error
	for name, fn := range checks {
		opts, ok := config.Checks[name]
		if !ok {
			opts = defaultCheckOptions
		}
		if !opts.Active {
			continue
		}
		errs = append(errs, fn(a, opts)...)
	}
	return errs
}
