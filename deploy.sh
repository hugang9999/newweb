#! /bin/bash
kill -9 $(pgrep webserver)
cd ~/newweb/newweb
git pull
cd webserver/
./webserver &