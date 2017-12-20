#! /bin/sh
pm2 delete all
pm2 start /opt/redis/redis-server 
pm2 start dataNode/bin/www --name="node1" -- 3001
pm2 start dataNode/bin/www --name="node2" -- 3002
pm2 start proxy/proxy
