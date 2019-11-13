package main

import "testing"

func TestRender(t *testing.T) {
	_, err := Render(".alertmanager.tmpl.yml", map[string]interface{}{
		"Env": map[string]string{
			"SLACK_API_URL":      "http://slack.com/blablah",
			"SMTP_HOST":          "localhost",
			"SMTP_FROM":          "no-reply@localhost.com",
			"SMTP_AUTH_USERNAME": "user",
			"SMTP_AUTH_PASSWORD": "pass",
		},
		"Values": map[string]interface{}{
			"Channel": "#channel",
			"Phone":   911,
		},
	})
	if err != nil {
		t.Error(err)
	}
}

func TestInLineRender(t *testing.T) {
	tests := map[string]string{
		"foobar@mail.com": "foobar@mail.com",
		"{{ .GroupLabels.label_mail_user }}@mail.com": "empty@mail.com",
	}
	for in, out := range tests {
		rout, err := InLineRender(in)
		if err != nil {
			t.Error(err)
			continue
		}
		if rout != out {
			t.Errorf("InLineRender(%s) != %s", in, out)
		}
	}
}
