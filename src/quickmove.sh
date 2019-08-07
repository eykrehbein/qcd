#!/bin/sh

if [ "$#" -gt 0 ] 
then
  if [ $1 == 'add' ] || [ $1 == 'list' ]
  then
    qmhelper $@
  else
    TARGET=$(qmhelper get $1)
    if [ $TARGET != 'UNDEFINED' ]
    then
      cd $TARGET
    else
      echo ""
      echo "This QuickLink couldn't be found"
      echo "Use: qm add $1 <path>"
      echo ""
    fi
  fi
else
  qmhelper h
fi