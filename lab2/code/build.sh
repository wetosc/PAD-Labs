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


gnome-terminal --geometry 80x22+0+0 -x sh -c ' ./node/node -id 1 -n "2" -f data1.json  ; bash'
gnome-terminal --geometry 80x22+600+0 -x sh -c ' ./node/node -id 2 -n "1 4 3" -f data2.json  ; bash'
gnome-terminal --geometry 80x22-0+0 -x sh -c ' ./node/node -id 3 -n "2 5" -f data3.json  ; bash'
gnome-terminal --geometry 80x22+0+300 -x sh -c ' ./node/node -id 4 -n "2 5" -f data4.json  ; bash'
gnome-terminal --geometry 80x22+600+300 -x sh -c ' ./node/node -id 5 -n "3 4" -f data5.json  ; bash'
gnome-terminal --geometry 80x22-0+300 -x sh -c ' ./node/node -id 6 -f data6.json  ; bash'

gnome-terminal --geometry 80x20-0-0 -x sh -c ' ./mediator/mediator -n "2 6"  ; bash'

gnome-terminal --geometry 80x20+0-0 -x sh -c ' ./client/client -f json  ; bash'
gnome-terminal --geometry 80x20+0-0 -x sh -c ' ./client/client -f xml   ; bash'
