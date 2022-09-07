# run-istio-demo sets up and follows the appropraite logs for debugging
#!/bin/zsh

set -xe

KUBEFOLLOWCMD="kubectl logs --follow"

tmux new-session -d -s istio-demo
tmux new-window -t istio-demo -n watch-machines
tmux new-window -t istio-demo -n watch-istiod
tmux new-window -t istio-demo -n watch-otel-collector
tmux new-window -t istio-demo -n watch-machine-istio-proxies
tmux new-window -t istio-demo -n run-machines

tmux send-keys -t watch-machines "${KUBEFOLLOWCMD} -l observability=true -n=machines" Enter
tmux send-keys -t watch-istiod "${KUBEFOLLOWCMD} -l app=istiod -n=istio-system" Enter
tmux send-keys -t watch-otel-collector "${KUBEFOLLOWCMD} -l app=otel-collector -n=istio-system " Enter
tmux send-keys -t watch-machine-istio-proxies "${KUBEFOLLOWCMD} -l observability -n=machines -c istio-proxy" Enter
tmux send-keys -t run-machines "./generator.sh" Enter
