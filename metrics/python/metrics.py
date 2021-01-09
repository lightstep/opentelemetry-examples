from unittest.mock import Mock

from opentelemetry.metrics import get_meter_provider, set_meter_provider
from opentelemetry.exporter.otlp.metrics_exporter import OTLPMetricsExporter
from opentelemetry.sdk.metrics import MeterProvider


meter_provider = MeterProvider()
set_meter_provider(meter_provider)
meter = get_meter_provider().get_meter(
    "name", instrumenting_library_version="version"
)
meter_provider.start_pipeline(
    meter,
    OTLPMetricsExporter(endpoint="127.0.0.1:7001", insecure=True),
    interval=1
)

counter = meter.create_counter(
    name="counter",
    description="description",
    unit="1",
    value_type=int,
)

updowncounter = meter.create_updowncounter(
    name="updowncounter",
    description="description",
    unit="1",
    value_type=int,
)

sumobserver = meter.register_sumobserver(
    callback=Mock(),
    name="sumobserver",
    description="description",
    unit="1",
    value_type=int,
)

updownsumobserver = meter.register_updownsumobserver(
    callback=Mock(),
    name="updownsumobserver",
    description="description",
    unit="1",
    value_type=int,
)

valueobserver = meter.register_valueobserver(
    callback=Mock(),
    name="valueobserver",
    description="description",
    unit="1",
    value_type=int,
)

labels = {"A": "B"}

test_values = [-1, 4, 3, 6, -5]

for test_case in test_values:
    counter.add(test_case, labels=labels)
    updowncounter.add(test_case, labels=labels)
    sumobserver.observe(test_case, labels=labels)
    updownsumobserver.observe(test_case, labels=labels)
    valueobserver.observe(test_case, labels=labels)

print("preshutdown")
get_meter_provider().shutdown()
print("postshutdown")
