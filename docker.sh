docker image build -t ascii-art-web-dockerize .
docker container run -d -p 8080:8080 ascii-art-web-dockerize
docker images
docker ps -a