docker run -it --name=rp1 --net=host -p 6060:6060 -d rp 
docker run -it --name=lb1 --net=host -p 6061:6061 -d lb 
docker run -it --name=ct1 --net=host -p 3333:3333 -d ct 

# docker stop rp1 &
# docker stop lb1 &
# docker stop ct1 &


# xterm -T "<connTrace>" -e docker start -a ct1 &
# xterm -T "<reverseproxy>" -e docker start -a rp1 &
# xterm -T "<loadbalancer>" -e docker start -a lb1 &
