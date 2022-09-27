#!/bin/bash

#####   name:Linux 服务注册 #####

create_service(){
    #获取当前目录
    gopower_path=`echo $(pwd) | sed 's/script//g'`
    #二进制文件
    gopower_bin=${gopower_path}/GoPower
    #创建服务文件
    touch /etc/systemd/system/gopower.service
    # 写入服务文件
    cat >>/etc/systemd/system/gopower.service << EOF
[Unit]
Description=GoPower
After=network.target

[Service]
ExecStart=${gopower_bin} start

[Install]
WantedBy=multi-user.target
EOF
    #重载服务
    systemctl daemon-reload
}

create_service