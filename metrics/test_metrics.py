from subprocess import Popen
from json import load
from os.path import dirname, abspath, join
from os import getcwd
from time import sleep
from shlex import split

from requests import get
from pytest import fixture
from docker import APIClient


_client = APIClient()


def fix_expected(actual_root):

    with open(
        join(dirname(abspath(__file__)), "expected_template.json"), "r"
    ) as expected_template_file:

        expected_root = load(expected_template_file)
        expected = expected_root["resource_metrics"][0]

    actual = actual_root["resource_metrics"][0]

    for expected_attribute in expected["resource"]["attributes"]:

        expected_attribute_key = expected_attribute["key"]

        if expected_attribute_key == "telemetry.sdk.language":

            language = expected_attribute

        elif expected_attribute_key == "telemetry.sdk.version":

            version = expected_attribute

    for actual_attribute in actual["resource"]["attributes"]:

        actual_attribute_key = actual_attribute["key"]

        if actual_attribute_key == "telemetry.sdk.language":

            (
                language["value"]["Value"]["string_value"]
            ) = (
                actual_attribute["value"]["Value"]["string_value"]
            )

        elif actual_attribute_key == "telemetry.sdk.version":

            (
                version["value"]["Value"]["string_value"]
            ) = (
                actual_attribute["value"]["Value"]["string_value"]
            )

    expected_metrics = {}

    for expected_metric in (
        expected
        ["instrumentation_library_metrics"]
        [0]
        ["metrics"]
    ):

        if expected_metric["name"] == "counter":
            expected_counter = expected_metric
            expected_metrics["counter"] = expected_counter

        elif expected_metric["name"] == "updowncounter":
            expected_updowncounter = expected_metric
            expected_metrics["updowncounter"] = expected_updowncounter

        elif expected_metric["name"] == "sumobserver":
            expected_sumobserver = expected_metric
            expected_metrics["sumobserver"] = expected_sumobserver

        elif expected_metric["name"] == "updownsumobserver":
            expected_updownsumobserver = expected_metric
            expected_metrics["updownsumobserver"] = expected_updownsumobserver

        elif expected_metric["name"] == "valueobserver":
            expected_valueobserver = expected_metric
            expected_metrics["valueobserver"] = expected_valueobserver

    for actual_metric in (
        actual
        ["instrumentation_library_metrics"]
        [0]
        ["metrics"]
    ):

        if actual_metric["name"] == "counter":
            expected_counter_data_point = (
                expected_counter
                ["Data"]
                ["int_sum"]
                ["data_points"]
                [0]
            )
            actual_counter_data_point = (
                actual_metric
                ["Data"]
                ["int_sum"]
                ["data_points"]
                [0]
            )
            (
                expected_counter_data_point
                ["start_time_unix_nano"]
            ) = (
                actual_counter_data_point
                ["start_time_unix_nano"]
            )
            (
                expected_counter_data_point
                ["time_unix_nano"]
            ) = (
                actual_counter_data_point
                ["time_unix_nano"]
            )

        elif actual_metric["name"] == "updowncounter":
            expected_updowncounter_data_point = (
                expected_updowncounter
                ["Data"]
                ["int_sum"]
                ["data_points"]
                [0]
            )
            actual_updowncounter_data_point = (
                actual_metric
                ["Data"]
                ["int_sum"]
                ["data_points"]
                [0]
            )
            (
                expected_updowncounter_data_point
                ["start_time_unix_nano"]
            ) = (
                actual_updowncounter_data_point
                ["start_time_unix_nano"]
            )
            (
                expected_updowncounter_data_point
                ["time_unix_nano"]
            ) = (
                actual_updowncounter_data_point
                ["time_unix_nano"]
            )

        elif actual_metric["name"] == "sumobserver":
            expected_sumobserver_data_point = (
                expected_sumobserver
                ["Data"]
                ["int_sum"]
                ["data_points"]
                [0]
            )
            actual_sumobserver_data_point = (
                actual_metric
                ["Data"]
                ["int_sum"]
                ["data_points"]
                [0]
            )
            (
                expected_sumobserver_data_point
                ["start_time_unix_nano"]
            ) = (
                actual_sumobserver_data_point
                ["start_time_unix_nano"]
            )
            (
                expected_sumobserver_data_point
                ["time_unix_nano"]
            ) = (
                actual_sumobserver_data_point
                ["time_unix_nano"]
            )

        elif actual_metric["name"] == "updownsumobserver":
            expected_updownsumobserver_data_point = (
                expected_updownsumobserver
                ["Data"]
                ["int_sum"]
                ["data_points"]
                [0]
            )
            actual_updownsumobserver_data_point = (
                actual_metric
                ["Data"]
                ["int_sum"]
                ["data_points"]
                [0]
            )
            (
                expected_updownsumobserver_data_point
                ["start_time_unix_nano"]
            ) = (
                actual_updownsumobserver_data_point
                ["start_time_unix_nano"]
            )
            (
                expected_updownsumobserver_data_point
                ["time_unix_nano"]
            ) = (
                actual_updownsumobserver_data_point
                ["time_unix_nano"]
            )

        elif actual_metric["name"] == "valueobserver":
            expected_valueobserver_data_point = (
                expected_valueobserver
                ["Data"]
                ["int_gauge"]
                ["data_points"]
                [0]
            )
            actual_valueobserver_data_point = (
                actual_metric
                ["Data"]
                ["int_gauge"]
                ["data_points"]
                [0]
            )
            (
                expected_valueobserver_data_point
                ["start_time_unix_nano"]
            ) = (
                actual_valueobserver_data_point
                ["start_time_unix_nano"]
            )
            (
                expected_valueobserver_data_point
                ["time_unix_nano"]
            ) = (
                actual_valueobserver_data_point
                ["time_unix_nano"]
            )

    (
        expected
        ["instrumentation_library_metrics"]
        [0]
        ["metrics"]
    ) = []

    for actual_metric in (
        actual
        ["instrumentation_library_metrics"]
        [0]
        ["metrics"]
    ):

        (
            expected
            ["instrumentation_library_metrics"]
            [0]
            ["metrics"]
        ).append(expected_metrics[actual_metric["name"]])

    return expected_root


@fixture(autouse=True, scope="module")
def build_server():

    repository = "metric-test-collector"

    if not _client.images(name=repository):

        # Without the iteration through the returned generator the image does
        # not seem to be generated
        [
            _ for _ in _client.build(
                path=dirname(abspath(__file__)),
                dockerfile="Dockerfile.collector",
                tag=repository,
            )
        ]

    collector_containers = _client.containers(
        all=True, filters={"label": "collector=go"}
    )

    if collector_containers:
        container_id = collector_containers[0]["Id"]

    else:
        container_id = _client.create_container(
            image="{}:latest".format(repository),
            ports=[7001, 7002],
            volumes=["/app"],
            host_config=_client.create_host_config(
                port_bindings={
                    7001: ("0.0.0.0", 7001), 7002: ("0.0.0.0", 7002)
                },
                binds={getcwd(): {"bind": "/app", "mode": "rw"}}
            ),
            detach=True,
            tty=True,
            labels={"collector": "go"}
        )["Id"]

    _client.start(container_id)

    # Apparently, some time is needed here to make the service start properly
    sleep(5)

    yield

    _client.stop(container_id)


def test_python():

    Popen("tox", cwd="python")

    sleep(5)

    actual = get("http://127.0.0.1:7002").json()

    expected = fix_expected(actual)

    assert expected == actual


def test_nodejs():

    Popen(split("npm run all"), cwd="nodejs")

    sleep(5)

    actual = get("http://127.0.0.1:7002").json()

    assert fix_expected(actual) == actual


def test_java():

    Popen(split("make build"), cwd="java")
    Popen(split("make run"), cwd="java")

    sleep(10)

    actual = get("http://127.0.0.1:7002").json()

    assert fix_expected(actual) == actual


def test_go():

    Popen(split("go build"), cwd="go")
    Popen(split("go run main.go"), cwd="go")

    sleep(10)

    actual = get("http://127.0.0.1:7002").json()

    assert fix_expected(actual) == actual
