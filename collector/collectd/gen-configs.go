package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func init() {
	fmt.Println("Running docker compose down")
	c := exec.Command("docker", "compose", "down")
	err := c.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func newCollectdOutput(pgsqlHost, pgsqlPort, pgsqlUser string, pgsqlPass string) string {
	collectdConfig := `
LoadPlugin write_prometheus
LoadPlugin write_pgsql

<Plugin "write_prometheus">
  Port 9103
</Plugin>

<Plugin postgresql>
  <Query magic>
    Statement "SELECT magic FROM wizard WHERE host = $1;"
    Param hostname
    <Result>
      Type gauge
      InstancePrefix "magic"
      ValuesFrom magic
    </Result>
  </Query>

  <Query rt36_tickets>
    Statement "SELECT COUNT(type) AS count, type \
                      FROM (SELECT CASE \
                                   WHEN resolved = 'epoch' THEN 'open' \
                                   ELSE 'resolved' END AS type \
                                   FROM tickets) type \
                      GROUP BY type;"
    <Result>
      Type counter
      InstancePrefix "rt36_tickets"
      InstancesFrom "type"
      ValuesFrom "count"
    </Result>
  </Query>

  <Database foo>
    Host ` + pgsqlHost + `
	Port ` + pgsqlPort + `
    User ` + pgsqlUser + `
    Password ` + pgsqlPass + `
    SSLMode "prefer"
    KRBSrvName "kerberos_service_name"
    Query magic
  </Database>

  <Database bar>
    Service "service_name"
    Query backend # predefined
    Query rt36_tickets
  </Database>
</Plugin>`
	err := os.WriteFile("collectd.conf", []byte(collectdConfig), 0644)
	if err != nil {
		panic(err)
	}
	return collectdConfig
}

func newDockerCompose(pgsqlHost, pgsqlPort, pgsqlUser string, pgsqlPass string) string {
	dockerCompose := `
version: "3.2"

services:
	pgsql:
      container_name: pgsql-collectd
      image: postgres
      environment:
        POSTGRES_DB: ` + pgsqlHost + `
        POSTGRES_USER: ` + pgsqlUser + `
        POSTGRES_PASSWORD: ` + pgsqlPass + `
      networks:
          - integrations
    collectd:
      container_name: collectd
      build: .
	depends_on:
      - "pgsql"
    networks:
        - integrations
    otel-collector:
      container_name: otel-collect
      image: otel/opentelemetry-collector-contrib:0.50.0
      command: ["--config=/conf/config-prometheus.yaml"]
      environment:
        LS_ACCESS_TOKEN: "XXXX"
      networks:
          - integrations
      volumes:
          - ./config-prometheus.yaml:/conf/config-prometheus.yaml:rw

networks:
  integrations:`
	err := os.WriteFile("docker-compose.yaml", []byte(dockerCompose), 0644)
	if err != nil {
		panic(err)
	}
	return dockerCompose
}

func main() {
	var (
		pgsqlHost = "\"0.0.0.0\""
		pgsqlPort = "\"5030\""
		pgsqlUser = "\"otel\""
		pgsqlPass = "\"otel\""
	)
	newDockerCompose(pgsqlHost, pgsqlPort, pgsqlUser, pgsqlPass)
	newCollectdOutput(pgsqlHost, pgsqlPort, pgsqlUser, pgsqlPass)
	fmt.Println("Running docker compose up")
	c := exec.Command("docker", "compose", "up")
	err := c.Run()
	if err != nil {
		log.Fatal(err)
	}
}
