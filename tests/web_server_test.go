package tests

import (
	"testing"
	"time"

	"github.com/gruntwork-io/terratest/modules/http-helper"
	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestWebServer(t *testing.T) {
	terraformOptions := &terraform.Options{
		TerraformDir: "../web",
	}

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)
	url := terraform.Output(t, terraformOptions, "url")
	status := 200
	text := "Hello, World"
	retries := 15
	sleep := 5 * time.Second
	http_helper.HttpGetWithRetry(t, url, status, text, retries, sleep)
}
