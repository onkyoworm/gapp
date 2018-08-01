# Applications Packager

G and Chain files packager for import into the Ecosystems created in the Gachain blockchain.
Utilite that can convert import json bundle from/to files of sim, ptl, csv, json.


### struct.dot

Is created in the process of packing or unpacking. Shows the structure of an application. Can be opened using [graphviz](http://graphviz.org/download/) or [webgraphviz](http://webgraphviz.com/)


## Examples

### Unpack file from "basic.sim" to "basic/"

>./gapp basic.sim

### Pack files from "basic/" to basic.json

>./gapp basic/


## build

Windows tip. If you do not plan to work in the console, add "-ldflags -H=windowsgui"

>go build  

Binary files can be found in the current directory 

### on linux for windows

 >env GOARCH=amd64 GOOS=windows CGO_ENABLED=1 CC=/usr/bin/x86_64-w64-mingw32-gcc CXX=/usr/bin/x86_64-w64-mingw32-g++  go build
