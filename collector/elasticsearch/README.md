
To get started I grab the files from the docs.

```sh
wget https://raw.githubusercontent.com/elastic/elasticsearch/8.2/docs/reference/setup/install/docker/docker-compose.yml
wget https://raw.githubusercontent.com/elastic/elasticsearch/8.2/docs/reference/setup/install/docker/.env
```

Then I update the `.env` file with passwords (fakes for testing) and a semver dotted triple. 

Now it's time to run. But before firing this one we should consider that the images are HUGE. So you should run `docker compose build` before you try `docker compose up`.Otherwise you're likely to face timeouts while waiting for resources.
