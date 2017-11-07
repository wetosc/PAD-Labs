#! /bin/sh

cd server/
go build
cd ../client/
go build
cd ../

chmod +x server/server
chmod +x client/client

gnome-terminal -x sh -c ' ./server/server -id 1 -n 3 -f data1.json  ; bash'
gnome-terminal -x sh -c ' ./server/server -id 2 -n 2 -f data2.json  ; bash'
gnome-terminal -x sh -c ' ./server/server -id 3 -n 1 2 -f data3.json  ; bash'

gnome-terminal -x sh -c ' ./client/client   ; bash'