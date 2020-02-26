#!/bin/bash


cd /root/go/src/req_resp/

startTime=100000
program=./send

#Section promgram (程序执行部分)
perDate=$(date "+%Y%m%d")
isNewDay=1
isFirstTime=1

while true ; do
    curTime=$(date "+%H%M%S")
    curDate=$(date "+%Y%m%d")

    #Check week day(周末不执行)
    week=`date +%w`
    if [ $week -eq 6 ] || [ $week -eq 0 ];then
        isNewDay=0
        sleep 1
        continue

    else
        #check and run script(工作日执行)
        if [ "$isNewDay" -eq "1" ];then
            if [ "$curTime" -gt "$startTime" ];then
                if [ "$isFirstTime" -eq "0" ];then
		            : > nohup.out
                    $program
                fi
                isNewDay=0
            else
                if [ "$isFirstTime" -eq "1" ];then
                    isFirstTime=0
                fi

            fi
        else
            #new day start(开始新的一天)
            if [ "$curDate" -gt "$perDate" ];then
                isNewDay=1
                perDate=$curDate
            fi
        fi
        sleep 1
    fi
done
