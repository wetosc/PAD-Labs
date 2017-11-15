#! /bin/sh

cd client/
go build
cd ../mediator/
go build
cd ../node/
go build
cd ../


chmod +x client/client
chmod +x mediator/mediator
chmod +x node/node


gnome-terminal -x sh -c ' ./node/node -id 1 -n 2 -f data1.json  ; bash'
gnome-terminal -x sh -c ' ./node/node -id 2 -n 1 -f data2.json  ; bash'

gnome-terminal -x sh -c ' ./mediator/mediator -n "1"  ; bash'

gnome-terminal -x sh -c ' ./client/client  ; bash'
