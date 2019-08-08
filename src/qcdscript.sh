#!/bin/sh

if [ "$#" -gt 0 ] 
then
  if [ $1 == 'add' ] || [ $1 == 'list' ] || [ $1 == 'remove' ] || [ $1 == 'rm' ] || [ $1 == 'help' ] || [ $1 == 'h' ] || [ $1 == 'version' ]
  then
    qcdhelper $@
  else
    TARGET=$(qcdhelper get $1)
    if [ $TARGET != 'UNDEFINED' ]
    then
      cd $TARGET
    else
      echo ""
      echo "This QuickLink couldn't be found"
      echo "Use: qcd add $1 <path>"
      echo ""
    fi
  fi
else
  qcdhelper h
fi