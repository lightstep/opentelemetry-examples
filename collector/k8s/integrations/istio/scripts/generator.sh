#!/bin/zsh

export MACHINEONE=$(kubectl get pod -l app=machine-1 -o jsonpath={.items..metadata.name} -n=machines)
export MACHINETWO=$(kubectl get pod -l app=machine-2 -o jsonpath={.items..metadata.name} -n=machines)
export MACHINETHREE=$(kubectl get pod -l app=machine-3 -o jsonpath={.items..metadata.name} -n=machines)

for (( ; ; ))
do
    kubectl exec $MACHINEONE -n machines -c machine-1 -- curl -sS -v machine-2.machines.svc.cluster.local:8081
    kubectl exec $MACHINETWO -n machines -c machine-2 -- curl -sS -v machine-3.machines.svc.cluster.local:8081
    kubectl exec $MACHINETHREE -n machines -c machine-3 -- curl -sS -v machine-1.machines.svc.cluster.local:8081
    sleep 2s
done
