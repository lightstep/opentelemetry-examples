enable: 
	docker-compose exec mon1 ceph mgr module enable dashboard
	docker-compose exec mon1 ceph mgr module enable prometheus
	docker-compose exec mon1 ceph config set mgr mgr/prometheus/server_addr 0.0.0.0
	docker-compose exec mon1 ceph config set mgr mgr/prometheus/server_port 9283