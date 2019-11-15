#!/bin/bash
# this script is used to start/stop mailsvr
app="mailsvr"
currentDir=`pwd`
pidfile=${currentDir}/pidfile.txt

usage() {
	echo "Usage:"
	echo "    $0 {start|stop|restart}"
	exit
}

getLocalPid() {
        if [ -f "$pidfile" ]; then
                localPid=`cat $pidfile`
		echo $localPid
	else
		echo 0
        fi
}

getRunningPid() {
	localPid=$(getLocalPid)
	if [ $localPid -gt 0 ]; then
		runningPid=`ps aux | grep "$localPid" | grep "$app" | grep -v grep | awk '{print $2}'`
		if [ "$localPid" = "$runningPid" ]; then
                	echo $localPid
        	else
                	echo 0
        	fi
	else
		echo 0
	fi
}

start() {
	currentPid=$(getRunningPid)
	if [ $currentPid -gt 0 ];then
		echo "program already running !"
		exit
	fi

	nohup ./${app} > error.log 2>&1 & echo $! > pidfile.txt
	#need to sleep, in order to wait for io task finish
    sleep 0.1 

    if [ -s error.log ]; then
        echo "something is wrong, please check! detail in error.log..."
    else
        echo "start succ."
    fi
}

stop() {
	currentPid=$(getRunningPid)
	if [ $currentPid -le 0 ]; then
		echo "program not running !"
		exit
	fi
	kill $currentPid
	echo "stop succ. "
}

case "$1" in
	'start')
		start
		;;
	'stop')
		stop
		;;
	'restart')
		stop
		start
		;;
	*)
		echo "Usage: $0 {start|stop|restart}"
		exit 1
	;;
esac
