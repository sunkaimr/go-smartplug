#!/bin/bash

export SERVICE_NAME="smartplug"

function LogOut()
{
	echo "`date "+%Y-%m-%d %H:%M:%S"` " $@
}

function start()
{
    LogOut "service $service_name starting"
    ./smartplug
}

function stop()
{
    LogOut "service $service_name stopped"
    pkill $service_name
}

function restart()
{
    pkill $service_name
    sleep 1
    LogOut "service $service_name restarting"
     ./smartplug
}

function status()
{
    pid=`pgrep $service_name`
    LogOut "service $service_name is running, PID:$pid"
}


if [ $# == 1 ]; then
	case $1 in	
		"start")
            start            
			;;
		"stop")            
            stop
            exit 0
			;;	
        "status")
            status
            exit 0
            ;;
        "restart")
            restart
            ;;
		*)
			echo "Usage: $0 ( start | stop )";
			;;
	esac
else  
	echo "Usage: $0 ( status | start | stop | restart)";
	exit 1;
fi

