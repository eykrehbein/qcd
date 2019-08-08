#!/bin/sh
rm /usr/local/bin/qcdscript
rm /usr/local/bin/qcdhelper

cp bin/qcdscript /usr/local/bin
cp bin/qcdhelper /usr/local/bin

chmod +x /usr/local/bin/qcdscript