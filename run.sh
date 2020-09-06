#!/bin/bash
source ~/.bash_profile
print_help()
{
    echo "Usage: "
    echo "    run.sh start|fstop"
}

if [ $# -ne 1 ]
then 
    print_help
    exit -1
fi

basepath=$(cd `dirname $0`; pwd)
module="main"
echo $basepath
case "$1" in
    start)
         echo "starting $module ..."
         nohup $basepath/$module -p 8088 > /work/log/myframe/stdout.log 2>&1 & ret=$?
         sleep 1
         echo "starting $module finished..."
         exit $ret
    ;;
    run)
         echo "starting $module ..."
         eval exec $basepath/$module -p 8088 > /work/log/gowar/stdout.log 2>&1
    ;;
    fstop)
         pid=`ps x | grep $module | grep infrastructure | grep -v grep | awk '{print $1}'`
         if [ x"$pid" == ""x ];then
             exit 0
         else
             kill -TERM $pid
         fi
         exit $?
    ;;

    *)
        print_help
    ;;

esac

exit -1

