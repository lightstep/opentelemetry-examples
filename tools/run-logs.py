import os
import subprocess
# import threading
import time
import argparse

import docker
# from fire import Fire

def ping_docker() -> bool:
    try:
        client = docker.from_env()
        client.ping()
        return True
    except Exception as e:
        print(f"Docker daemon is not accessible: {e}")
        return False

def docker_compose_up():
    subprocess.run(["docker-compose", "up", "-d", "--wait"])

def docker_compose_down():
    subprocess.run(["docker-compose", "down", "--remove-orphans"])

def get_logs(service_name: str):
    logs = subprocess.check_output(["docker-compose", "logs", service_name])
    with open("logs.txt", "wb") as log_file:
        log_file.write(logs)

# run_docker_compose()
def run_example(app: str):

    app_dir = "{}/../collector/{}".format(os.path.dirname(os.path.realpath(__file__)), app)

    if os.path.exists(f"{app_dir}/docker-compose.yml"):
        if not ping_docker():
            print(f"can't find docker")
            return

        try:
            os.chdir(app_dir)
            docker_compose_up()
        except Exception as e:
            print(f"failed to start example: {e}")
            return # short circuit

        # let the service run so we capture metrics
        time.sleep(60)

        try:
            get_logs("otel-collector")
        except Exception as e:
            print("failed to get logs")
        finally:
            docker_compose_down()
    elif os.path.exists(f"{app_dir}/Makefile"):
        # requires Makefile for examples with "run", "get-logs" and "stop" targets
        try:
            subprocess.run(["make", "run", "-C", app_dir])
            time.sleep(60)
            logs = subprocess.check_output(["make", "get-logs", "-C", app_dir])
            with open("{}/logs.txt".format(app_dir), "wb") as log_file:
                log_file.write(logs)
        except Exception as e:
            print("failed to start example or get logs")
        finally:
            subprocess.run(["make", "stop", "-C", app_dir])

# Fire(run_example)

parser = argparse.ArgumentParser(
        prog='demo environment runner',
        description='Runs docker compose examples',
        )

parser.add_argument('dir')

args = parser.parse_args()

run_example(args.dir)
