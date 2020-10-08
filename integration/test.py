#!/usr/bin/env python3
#
# This integration test does the following:
#   - creates a new span called "integration_test_requests"
#   - sends a request to each configured server
#   - use the Lightstep API to query for the trace describing that span
#   - ensures it contains information from all of the servers
#
# If the trace is incomplete, then the context did not correctly propagate
# across all the servers.

import os
import time

import requests
import yaml
from opentelemetry import trace

tracer = trace.get_tracer(__name__)


# the integration test is instrumented to propagate
# context across to all servers via the opentelemetry-launcher
INTEGRATION_TEST_APP = os.environ.get("LS_SERVICE_NAME")
API_URL = "https://api.lightstep.com/public/v0.2"
TEST_ORG = os.environ.get("ORG_NAME")  # maybe this should be a different one
PROJECT = os.environ.get("PROJECT_NAME")
TOKEN = os.environ.get("API_KEY")


def _get_headers():
    return {
        "Authorization": "{}".format(TOKEN),
        "Content-Type": "application/json",
    }


def test_auth():
    url = "{}/{}/test".format(API_URL, TEST_ORG)
    response = requests.get(url, headers=_get_headers())
    assert response.status_code == 200


def _get_destinations():
    config_file = os.environ.get("INTEGRATION_CONFIG_FILE")
    if not config_file:
        raise Exception("Config file not specified!!")

    config_data = {}
    with open(config_file) as f:
        config_data = yaml.load(f, Loader=yaml.FullLoader)

    return config_data.get("endpoints")


def _get_services():
    config_file = os.environ.get("INTEGRATION_CONFIG_FILE")
    if not config_file:
        raise Exception("Config file not specified!!")

    config_data = {}
    with open(config_file) as f:
        config_data = yaml.load(f, Loader=yaml.FullLoader)

    return config_data.get("services")


def create_trace():
    span_id = None
    with tracer.start_as_current_span("integration_test_requests") as span:
        for url in _get_destinations():
            try:
                if "/order" in url:
                    res = requests.post(
                        url, data='{"donuts":[{"flavor":"cinnamon","quantity":1}]}'
                    )
                else:
                    res = requests.get(url)
                print(f"Request to {url}, got {len(res.content)} bytes")
            except Exception as e:
                print(f"Request to {url} failed {e}")
        span_id = span.get_context().span_id
    return span_id


def test_traces():
    # give time for services to spin up
    # note this should probably be using requests retries instead
    time.sleep(300)

    # send a trace
    span_id = create_trace()
    assert span_id is not None

    # give time for services to report traces
    time.sleep(60)

    url = "{}/{}/projects/{}/snapshots".format(API_URL, TEST_ORG, PROJECT)
    querystring = 'service IN ("{}")'.format(INTEGRATION_TEST_APP)
    payload = {"data": {"attributes": {"query": querystring}}}

    # create a snapshot to make the trace we generated available
    response = requests.post(url, headers=_get_headers(), json=payload)
    assert response.status_code == 200

    url = "{}/{}/projects/{}/stored-traces".format(API_URL, TEST_ORG, PROJECT)
    querystring = {"span-id": format(span_id, "x")}

    # search the snapshot for our trace
    response = requests.get(url, headers=_get_headers(), params=querystring)
    retries = 5
    while retries > 0 and response.status_code != 200:
        retries -= 1
        time.sleep(8)
        response = requests.get(url, headers=_get_headers(), params=querystring)

    assert response.status_code == 200
    results = response.json()
    reporters = (
        results.get("data", [{}])[0].get("relationships", {}).get("reporters", {})
    )
    services = _get_services()

    # the integration test will report as well
    services.append(INTEGRATION_TEST_APP)
    expected_services_count = len(services)

    # each server will be listed as a reporter in the trace being retrieved
    # we're inspecting the list of reporters rather than individual span
    # to prevent having to update this test every time an example application
    # updates spans being generated
    for reporter in reporters:
        service_name = reporter.get("attributes", {}).get("lightstep.component_name")
        if service_name in services:
            services.remove(service_name)

    # assert number of reporters are the the same as expected
    assert len(reporters) == expected_services_count, "Services not found: {}".format(
        services
    )
    assert services == []

