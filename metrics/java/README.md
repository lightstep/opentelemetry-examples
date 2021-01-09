# How to run these tests

## Start collector and build

1. Start collector
1. Build:

```shell script
make buid
```

## Java Metric Producer

1. Run

```
make run
```

It will produce 6 metrics:

* Counter
* UpDownCounter
* ValueObserver
* SumObserver
* UpDownSumObserver
* ValueRecorder

2. You can produce only individual metrics

* Counter

```
make counter
```

* UpDownCounter

```
make updown_counter
```

* SumObserver

```
make sum_observer
```

* UpDownSumObserver

```
make updown_sum_observer
```

* Value Observer

```
make value_observer
```

* Value Recorder

```
make value_recorder
```