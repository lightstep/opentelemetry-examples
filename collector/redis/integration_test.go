package redis__test

import (
	"context"
	"io"
	"net/http"
	"regexp"
	"testing"
	"time"

	"github.com/testcontainers/testcontainers-go"

	"github.com/stretchr/testify/assert"
	tc "github.com/testcontainers/testcontainers-go/modules/compose"
)

// use maps as sets
type void struct{}

const path = "docker-compose.yml"
const otelService = "otel-collector"
const tfURL = "https://raw.githubusercontent.com/lightstep/terraform-opentelemetry-dashboards/main/collector-dashboards/otel-collector-redisreceiver-dashboard/main.tf"

func TestDockerCompose(t *testing.T) {
	compose, err := tc.NewDockerCompose(path)
	assert.NoError(t, err, "NewDockerComposeAPI()")

	t.Cleanup(func() {
		assert.NoError(t, compose.Down(context.Background(), tc.RemoveOrphans(true), tc.RemoveImagesLocal), "compose.Down()")
	})

	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)

	assert.NoError(
		t,
		compose.WithOsEnv().Up(ctx, tc.Wait(true)),
		"compose.Up()",
	)

	t.Logf("Services: %+v", compose.Services())

	container, err := compose.ServiceContainer(ctx, otelService)
	assert.NoError(t, err, "ServiceContainer(otel-collector)")

	logsMetrics, err := parseContainerLogs(ctx, container)
	assert.NoError(t, err, "ServiceContainer(otel-collector).Logs")
	t.Log(logsMetrics)

	// get tf metrics collection
	tfMetrics, err := parseTF(tfURL)
	assert.NoError(t, err, "tf.Parse")

	for metric := range logsMetrics {
		if _, ok := tfMetrics[metric]; !ok {
			assert.Failf(t, "Missing %s from terraform metrics.", metric)
		}
	}

	for metric := range tfMetrics {
		if _, ok := logsMetrics[metric]; !ok {
			t.Logf("WARN: Missing %s from logs metrics. "+
				"Probably you didn't generate enough logs or enough load for your service", metric)
		}
	}
}

func parseContainerLogs(ctx context.Context, c *testcontainers.DockerContainer) (map[string]void, error) {
	time.Sleep(20 * time.Second)

	logsReader, err := c.Logs(ctx)
	if err != nil {
		return nil, err
	}

	logByte, err := io.ReadAll(logsReader)
	if err != nil {
		return nil, err
	}

	logs := string(logByte)

	logsMetricsPattern := regexp.MustCompile(`-> Name: (.*)`)

	matches := logsMetricsPattern.FindAllStringSubmatch(logs, -1)
	logsMetrics := make(map[string]void, len(matches))
	for i := range matches {
		logsMetrics[matches[i][1]] = void{}
	}
	return logsMetrics, nil
}

func parseTF(link string) (map[string]void, error) {
	tfResponse, err := http.Get(link)
	if err != nil {
		return nil, err
	}

	tfByte, err := io.ReadAll(tfResponse.Body)
	if err != nil {
		return nil, err
	}

	tf := string(tfByte)

	tfMetricsPattern := regexp.MustCompile(`metric\s*=\s*"(.*)"`)

	matches := tfMetricsPattern.FindAllStringSubmatch(tf, -1)
	tfMetrics := make(map[string]void, len(matches))
	for i := range matches {
		tfMetrics[matches[i][1]] = void{}
	}

	return tfMetrics, nil
}
