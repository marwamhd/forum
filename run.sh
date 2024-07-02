docker image build -f Dockerfile -t manhello .
docker container run -p 9090:9090 --detach --name hello manhello
docker run -it --entrypoint /bin/bash manhello
docker stop hello