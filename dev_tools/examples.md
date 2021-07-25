NANO: 
-> nano myfile.txt                                  ;create file
-> ctrl-X -> confirm name -> Y/n to save a file
-> nano myfile.txt                                  ;for edit

touch:
-> touch .gitignore                                 ; create empty file. Open with nano/vim for edit

cat | grep:
https://www.geeksforgeeks.org/grep-command-in-unixlinux/
-> cat some.log | grep auth -i                      ; read the content of some.log and show only lines with mention "auth"
-> cat some.log | grep auth -i > output.log         ; to route the output in a file
-> echo "hello world"                               ; write the line to stdout

ping | telnet
-> ping google.com                                  ; will send a requests for google. Usable to debug the routing and network issues
-> telnet google.com 80                             ; usually usable for checking opened ports
-> sudo tcpdump host example.com                    ; must show the http traffic with host
-> dig example.com                                  ; show the dns resolution details
-> curl example.com                                 ; show the http request to the host
