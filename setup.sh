#!/bin/sh

cp bin/qcdscript /usr/local/bin
cp bin/qcdhelper /usr/local/bin

chmod +x /usr/local/bin/qcdscript

printf "\nalias qcd='source /usr/local/bin/qcdscript'" >> ~/.bash_profile