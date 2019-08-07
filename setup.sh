#!/bin/sh

cp bin/qmscript /usr/local/bin
cp bin/qmhelper /usr/local/bin

chmod +x /usr/local/bin/qmscript

printf "\nalias qm='source /usr/local/bin/qmscript'" >> ~/.bash_profile