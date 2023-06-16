#!/bin/bash

script_dir=$(cd $(dirname $0);pwd)
echo "project : $script_dir"
pid_path="${script_dir}/pid"
echo "pid : $pid_path"
is_exist="0"
if [ -f "$pid_path" ]; then
  echo "pid file is exists"
  pid=$(cat ${script_dir}/pid)
  echo "pid is $pid"
  test=$(ps -ef | grep $pid | grep -v grep | awk '{print $2}')
  while read pid2
  do
    echo "find pid $pid2"
    if [ "$pid" -eq "$pid2" ]; then
      echo "find process"
      is_exist="1"
    fi
  done <<< $test
  echo "is_exist is $is_exist"
  if [ "$is_exist" -eq "1" ]; then
     echo kill -1 $pid
    kill -1 $pid
  else
     echo "pid is not exists"
      echo "$script_dir/service >> $script_dir/nohup.log 2>&1 &"
      nohup $script_dir/service >> $script_dir/nohup.log 2>&1 &
  fi
else
  echo "pid is not exists"
  echo "$script_dir/service >> $script_dir/nohup.log 2>&1 &"
  nohup $script_dir/service >> $script_dir/nohup.log 2>&1 &
fi



