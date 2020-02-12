
docker build -f ./docker/rp.dockerfile -t rp . 
docker build -f ./docker/lb.dockerfile -t lb . 
docker build -f ./docker/ct.dockerfile -t ct .


