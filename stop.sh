PID=$(ps -ef | grep service | grep -v grep | awk '{ print $2 }')
if [ -z "$PID" ]
then
    echo service is already stopped
else
    echo kill -9 $PID
    kill -9 $PID
fi
