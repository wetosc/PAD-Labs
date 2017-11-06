#! /bin/sh

cd server/
go build
cd ../client/
go build
cd ../

chmod +x server/server
chmod +x client/client

gnome-terminal -x sh -c ' ./server/server -p 1 -n 3   ; bash'
gnome-terminal -x sh -c ' ./server/server -p 2 -n 4   ; bash'
gnome-terminal -x sh -c ' ./server/server -p 3 -n 3   ; bash'

gnome-terminal -x sh -c ' ./client/client   ; bash'