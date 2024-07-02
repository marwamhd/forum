docker image build -f Dockerfile -t Fpro .
docker container run -p 8080:1010  --name Conp Fpro
docker stop Conp