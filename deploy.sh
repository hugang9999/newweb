#! /bin/bash
kill -9 $(pgrep webserver)
cd ~/newweb/
git pull https://172.20.23.19:5000/newweb.git
cd webserver/
./webserver &