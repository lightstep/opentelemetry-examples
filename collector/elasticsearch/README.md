# Monitoring Elasticsearch

<<<<<<< HEAD
<<<<<<< HEAD
## About this Configuration

Elasticsearch is frequently run as part of the ELK stack with Logstash and Kibana. For clarity we limited this example to a single Elasticsearch node to cover just the information that you need to integrate with Lightstep.

## Setup

Make sure that you have set the environment variable `LS_ACCESS_TOKEN`. Then to prepare the environment you can run `make setup` which will create certificates. However, it isn't necessary to run `make setup` before running `make up`, because the `up` rule will check for the directory called `secrets` and generate certificates if it doesn't exist.

The file `.env` contains key configuration values which are also referenced in other files such as `collector.yml` and the docker compose configuration. Be sure to set the values here appropriately for your situation.

Other useful `make` rules are provided such as `down` and `prune` for cleaning up the Docker host after running the example. Consult the Makefile for more details.

## Running this Example

After you run `docker compose -f docker-compose.setup.yml`, you can run the example with `docker compose up`. For convenience there's also a make rule you can run with `make up`. This will detect whether setup has been run and will run it if needed.

Edit the .env file to adjust variables for your configuration.

## About the Configuration

The base `docker-commpose.yml` file includes the Elasticsearch node. The file `docker-compose.override.yml` includes the OTEL Collector. And `docker-compose.setup.yml` includes services that setup the requisite Elastic keystore and certificates.

Note that the file receiver of the OTEL receiver is also configured for this example to simplify inspection of the output.

## License

Some parts of this configuration are derived from the [elastdocker](https://github.com/sherifabdlnaby/elastdocker/) model configuration. It is used and provided under a commercially permissive MIT license.
<<<<<<< HEAD
=======
## The 
=======
## Setup

`make setup`

## 
>>>>>>> 43056f9 (simplify example)




<p align="center">
<img width="500px" src="https://user-images.githubusercontent.com/16992394/147855783-07b747f3-d033-476f-9e06-96a4a88a54c6.png">
</p>
<h2 align="center"><b>Elast</b>ic Stack on <b>Docker</b></h2>
<h3 align="center">Preconfigured Security, Tools, and Self-Monitoring</h3>
<h4 align="center">Configured to be ready to be used for Log, Metrics, APM, Alerting, Machine Learning, and Security (SIEM) usecases.</h4>
<p align="center">
   <a>
      <img src="https://img.shields.io/badge/Elastic%20Stack-8.2.0-blue?style=flat&logo=elasticsearch" alt="Elastic Stack Version 7^^">
   </a>
   <a>
      <img src="https://img.shields.io/github/v/tag/sherifabdlnaby/elastdocker?label=release&amp;sort=semver">
   </a>
   <a href="https://github.com/sherifabdlnaby/elastdocker/actions/workflows/build.yml">
      <img src="https://github.com/sherifabdlnaby/elastdocker/actions/workflows/build.yml/badge.svg">
   </a>
   <a>
      <img src="https://img.shields.io/badge/Log4Shell-mitigated-brightgreen?style=flat&logo=java">
   </a>
   <a>
      <img src="https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat" alt="contributions welcome">
   </a>
   <a href="https://github.com/sherifabdlnaby/elastdocker/network">
      <img src="https://img.shields.io/github/forks/sherifabdlnaby/elastdocker.svg" alt="GitHub forks">
   </a>
   <a href="https://github.com/sherifabdlnaby/elastdocker/issues">
        <img src="https://img.shields.io/github/issues/sherifabdlnaby/elastdocker.svg" alt="GitHub issues">
   </a>
   <a href="https://raw.githubusercontent.com/sherifabdlnaby/elastdocker/blob/master/LICENSE">
      <img src="https://img.shields.io/badge/license-MIT-blue.svg" alt="GitHub license">
   </a>
</p>

# Introduction
Elastic Stack (**ELK**) Docker Composition, preconfigured with **Security**, **Monitoring**, and **Tools**; Up with a Single Command.

Suitable for Demoing, MVPs and small production deployments.

Stack Version: [8.2.0](https://www.elastic.co/blog/whats-new-elastic-8-2-0) 🎉  - Based on [Official Elastic Docker Images](https://www.docker.elastic.co/)
> You can change Elastic Stack version by setting `ELK_VERSION` in `.env` file and rebuild your images. Any version >= 8.0.0 is compatible with this template.

### Main Features 📜

- Configured as a Production Single Node Cluster. (With a multi-node cluster option for experimenting).
- Security Enabled By Default.
- Configured to Enable:
  - Logging & Metrics Ingestion
  - APM
  - Alerting
  - Machine Learning
  - SIEM
  - Enabling Trial License
- Use Docker-Compose and `.env` to configure your entire stack parameters.
- Persist Elasticsearch's Keystore and SSL Certifications.
- Self-Monitoring Metrics Enabled.
- Prometheus Exporters for Stack Metrics.
- Collect Docker Host Logs to ELK via `make collect-docker-logs`.
- Embedded Container Healthchecks for Stack Images.
- [Rubban](https://github.com/sherifabdlnaby/rubban) for Kibana curating tasks.

#### More points
And comparing Elastdocker and the popular [deviantony/docker-elk](https://github.com/deviantony/docker-elk)

<details><summary>Expand...</summary>
<p>

One of the most popular ELK on Docker repositories is the awesome [deviantony/docker-elk](https://github.com/deviantony/docker-elk).
Elastdocker differs from `deviantony/docker-elk` in the following points.

- Security enabled by default using Basic license, not Trial.

- Persisting data by default in a volume.

- Run in Production Mode (by enabling SSL on Transport Layer, and add initial master node settings).

- Persisting Generated Keystore, and create an extendable script that makes it easier to recreate it every-time the container is created.

- Parameterize credentials in .env instead of hardcoding `elastich:changeme` in every component config.

- Parameterize all other Config like Heap Size.

- Add recommended environment configurations as Ulimits and Swap disable to the docker-compose.

- Make it ready to be extended into a multinode cluster.

- Configuring the Self-Monitoring and the Filebeat agent that ship ELK logs to ELK itself. (as a step to shipping it to a monitoring cluster in the future).

- Configured tools and Prometheus Exporters.

- The Makefile that simplifies everything into some simple commands.

</p>
</details>

-----

# Requirements

- [Docker 20.05 or higher](https://docs.docker.com/install/)
- [Docker-Compose 1.29 or higher](https://docs.docker.com/compose/install/)
- 4GB RAM (For Windows and MacOS make sure Docker's VM has more than 4GB+ memory.)

# Setup

1. Initialize Elasticsearch Keystore and TLS Self-Signed Certificates
    ```bash
    $ make setup
    ```
    > **For Linux's docker hosts only**. By default virtual memory [is not enough](https://www.elastic.co/guide/en/elasticsearch/reference/current/vm-max-map-count.html) so run the next command as root `sysctl -w vm.max_map_count=262144`

    > - Modify `.env` file for your needs, most importantly `ELASTIC_PASSWORD` that setup your superuser `elastic`'s password, `ELASTICSEARCH_HEAP` & `LOGSTASH_HEAP` for Elasticsearch & Logstash Heap Size.
    
#### To Rebuild Images
```shell
$ make build
```
#### Bring down the stack.
```shell
$ make down
```

#### Reset everything, Remove all containers, and delete **DATA**!
```shell
$ make prune
```

# Configuration

* Some Configuration are parameterized in the `.env` file.
  * `ELASTIC_PASSWORD`, user `elastic`'s password (default: `changeme` _pls_).
  * `ELK_VERSION` Elastic Stack Version (default: `8.2.0`)
  * `ELASTICSEARCH_HEAP`, how much Elasticsearch allocate from memory (default: 1GB -good for development only-)
  * `LOGSTASH_HEAP`, how much Logstash allocate from memory.
  * Other configurations which their such as cluster name, and node name, etc.
* Elasticsearch Configuration in `elasticsearch.yml` at `./elasticsearch/config`.
* Logstash Configuration in `logstash.yml` at `./elasticsearch/config/logstash.yml`.
* Logstash Pipeline in `main.conf` at `./elasticsearch/pipeline/main.conf`.
* Kibana Configuration in `kibana.yml` at `./kibana/config`.
* Rubban Configuration using Docker-Compose passed Environment Variables.

### Setting Up Keystore

You can extend the Keystore generation script by adding keys to `./setup/keystore.sh` script. (e.g Add S3 Snapshot Repository Credentials)

To Re-generate Keystore:
```
make keystore
```

# License
[MIT License](https://raw.githubusercontent.com/sherifabdlnaby/elastdocker/master/LICENSE)
<<<<<<< HEAD
Copyright (c) 2020 Sherif Abdel-Naby

# Contribution

PR(s) are Open and Welcomed.
>>>>>>> cbfd42f (Refactor docker-compose and .env)
=======
Copyright (c) 90% Sherif Abdel-Naby 2020-2022 + 10% Lightstep 2022
>>>>>>> 43056f9 (simplify example)
=======
>>>>>>> 93728d9 (update rules and names)
