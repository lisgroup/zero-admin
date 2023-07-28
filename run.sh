#!/bin/bash

# 定义运行的命令
# shellcheck disable=SC2034
UMS_CMD="go run rpc/ums/ums.go -f rpc/ums/etc/ums.yaml > ums.log 2>&1 &"
OMS_CMD="go run rpc/oms/oms.go -f rpc/oms/etc/oms.yaml > oms.log 2>&1 &"
PMS_CMD="go run rpc/pms/pms.go -f rpc/pms/etc/pms.yaml > pms.log 2>&1 &"
SMS_CMD="go run rpc/sms/sms.go -f rpc/sms/etc/sms.yaml > sms.log 2>&1 &"
CMS_CMD="go run rpc/cms/cms.go -f rpc/cms/etc/cms.yaml > cms.log 2>&1 &"
SYS_CMD="go run rpc/sys/sys.go -f rpc/sys/etc/sys.yaml > sys.log 2>&1 &"
API_CMD="go run api/admin.go -f api/etc/admin-api.yaml > admin-api.log 2>&1 &"

# kill

# shellcheck disable=SC2034
commands=("$UMS_CMD" "$OMS_CMD" "$PMS_CMD" "$SMS_CMD" "$CMS_CMD" "$SYS_CMD" "$API_CMD")

for cmd in "${commands[@]}"; do
    # shellcheck disable=SC2009
    pid=$(ps aux | grep "$cmd" | grep -v grep | awk '{print $2}')
    # 如果找到了PID,就结束该进程
    if [ -n "$pid" ]
    then
        echo "Killing process $pid"
        kill -9 "$pid"
    else
        echo "Process not found for command $cmd"
    fi
done

# 定义工作目录变量为当前目录
work_dir=$(pwd)
echo "$work_dir"
# 切换到工作目录
cd "$work_dir"||return

## 执行命令
for cmd in "${commands[@]}"; do
    echo "Running command: $cmd"
    eval "$cmd"
done

#echo "ums"
#go run rpc/ums/ums.go -f rpc/ums/etc/ums.yaml > ums.log 2>&1 &
#echo "oms"
#go run rpc/oms/oms.go -f rpc/oms/etc/oms.yaml > oms.log 2>&1 &
#echo "pms"
#go run rpc/pms/pms.go -f rpc/pms/etc/pms.yaml > pms.log 2>&1 &
#echo "sms"
#go run rpc/sms/sms.go -f rpc/sms/etc/sms.yaml > sms.log 2>&1 &
#echo "cms"
#go run rpc/cms/cms.go -f rpc/cms/etc/cms.yaml > cms.log 2>&1 &
#go run rpc/sys/sys.go -f rpc/sys/etc/sys.yaml > sys.log 2>&1 &
#echo "api"
#go run api/admin.go -f api/etc/admin-api.yaml > admin-api.log 2>&1
