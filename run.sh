docker image build -f Dockerfile -t plimage .
docker run -p 4343:4343 plimage