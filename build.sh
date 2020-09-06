#!/bin/sh


make(){
    if [ "$1" == "all" ] ;then
       TARGETLIST=$(ls cmd)
       for i in $TARGETLIST
       do
            echo "build begin:"+$i
            verify_run "go build -o bin/$i cmd/$i/main.go " 
       done
    else
        verify_run "go build -o bin/$1 cmd/$1/main.go " 
    fi
}


function verify_run()
{
    echo "$@"
    cmd=$@
    $@
    if [ $? -ne 0 ]
    then
        echo $LOG_FATAL "run cmd error: $cmd"
        exit 255
    else
        echo $LOG_NOTICE "run cmd success: $cmd"
        return 0
    fi
}

print_help_local()
{
    echo "samples:"
    echo "----------------------------not rewriteable ---------------------------------------------"
    echo "输入  bash build.sh make all 编译全部模块"
    echo "-----------------------------------------------------------------------------------------"
}


Main()
{
    if [ "$1" == "-h" -o "$1" == "--help" -o "$1" == "-help" -o $# -eq 0 ]; then
        print_help_local
    elif [ "$1" == "make" ];then
        make $2
    elif [ "$1" == "test" ];then
        $1
    fi
    exit 0
}


Main $@
