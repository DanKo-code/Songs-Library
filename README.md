## 🐳 Using Docker Compose

You can find [docker-compose.example.yml](./docker-compose.example.yml) in this repo and adapt it to your own needs. 

1. Get [Docker](https://www.docker.com/get-started) and [Docker Compose](https://www.digitalocean.com/community/tutorial-collections/how-to-install-docker-compose)

2. Download the Docker Compose example file and save it as `docker-compose.yml` on your local machine

```bash
curl https://raw.githubusercontent.com/vas3k/pepic/master/docker-compose.example.yml -o docker-compose.yml
```

3. Now run it

```bash
docker-compose up
```

4. Go to [http://localhost:8118](http://localhost:8118) and try uploading something. You should see uploaded images or videos in the local directory (`./uploads`) after that.
