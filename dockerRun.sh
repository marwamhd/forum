docker image build -f Dockerfile -t fpro .
docker container run -p 8080:5050  --name conp fpro
docker stop conp