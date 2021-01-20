# How to run these tests

## Install needed packages

1. Install needed packages 
```
    npm install
``` 

## Nodejs Metric Producer

1. Run 
```
npm run produce
```
It will produce 6 metrics:
* Counter
* UpdDownCounter
* ValueObserver 
* SumObserver
* UpDownSumObserver
* ValueRecorder
2. You can produce only individual metrics
* Counter
```
npm run counter
```
* UpDownCounter
```
npm run updown_counter
```
* SumObserver
```
npm run sum_observer
```
* UpDownSumObserver
```
npm run updown_sum_observer
```
* Value Observer
```
npm run value_observer
```
* Value Recorder
```
npm run value_recorder
```

## Run server via Docker
Make sure you are in main folder
```
docker build -f Dockerfile.collector -t metric-test-collector .
docker run -v $PWD:/app -p 0.0.0.0:7001:7001 metric-test-collector
```
