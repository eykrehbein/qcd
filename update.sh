#!/bin/sh
rm /usr/local/bin/qmscript
rm /usr/local/bin/qmhelper

cp bin/qmscript /usr/local/bin
cp bin/qmhelper /usr/local/bin

chmod +x /usr/local/bin/qmscript