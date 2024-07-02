docker image build -f Dockerfile -t mumu .
docker container run -p 1010:1010 --detach --name mummmm mumu
docker run -it --entrypoint /bin/bash mumu
docker stop mummmm