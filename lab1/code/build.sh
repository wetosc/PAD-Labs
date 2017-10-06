#! /bin/sh

cd broker/
go build
cd ../client/
go build
cd ../

chmod +x broker/broker
chmod +x client/client

gnome-terminal -x sh -c '  ./broker/broker -v debug  ; bash'

gnome-terminal -x sh -c '  ./client/client -info "Publisher ALFA" -type sender -queue "a.a" -v debug  ; bash'
# gnome-terminal -x sh -c '  ./client/client -info "Publisher BETA" -type sender -queue "a.b" -v debug  ; bash'

gnome-terminal -x sh -c '  ./client/client -type receiver -queue "a.a" -v debug  ; bash'
# gnome-terminal -x sh -c '  ./client/client -type receiver -queue "a.b" -v debug  ; bash'
# gnome-terminal -x sh -c '  ./client/client -type receiver -queue "a.c" -v debug  ; bash'



