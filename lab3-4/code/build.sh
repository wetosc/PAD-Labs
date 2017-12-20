#! /bin/sh
pm2 start dataNode/bin/www --name="node1" -- 3001
pm2 start dataNode/bin/www --name="node2" -- 3002
