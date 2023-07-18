#!/bin/bash

# 格式：\033[显示方式；前景色；背景色m
# 说明：
#
# 前景色       背景色       颜色
# -----------------------------
# 30           40          黑色
# 31           41          红色
# 32           42          绿色
# 33           43          黄色
# 34           44          蓝色
# 35           45          紫红色
# 36           46          青蓝色
# 37           47          白色
#
# 显示方式            意义
# --------------------------
#    0           终端默认设置
#    1           高亮显示
#    4           使用下划线
#    5           闪烁
#    7           反白显示
#    8           不可见
#
# 例子
# \033[1;31;40m   <！--1-高亮显示 31-前景色红色40-背景色黑色-->
# \033[0m         <！--采用终端默认设置，即取消颜色设置-->]]]

# 显示标题，用于大的模块的开始
show_title() {
  echo -e "\033[1;36m$1\033[0m"
}

# 显示成功提示
show_suc() {
  echo -e "\033[1;32m$1\033[0m"
}

# 显示错误信息
show_error() {
  echo -e "\033[1;31m$1\033[0m"
}

#显示普通提示信息
show_msg() {
  echo -e "\033[33m$1\033[0m"
}

#检测上一次执行的命令的返回值，如果不为0，则表示上次命令出现异常
check_error() {
  ERROR_NO=$1
  MESSAGE=$2
  if [ $ERROR_NO -ne 0 ]; then
    show_error "ERROR: $MESSAGE"
    exit $ERROR_NO
  fi
}

get_input() {
  tip=$1
  default_value=$2
  value=""
  while ([[ -z $value ]]); do
    read -p "$tip" value
    if [ -z $value ] && [ -n $default_value ]; then
      value=$default_value
    fi
  done
  echo $value
}

function create_static_folder() {
  path=$1
  #如果文件夹不存在，创建文件夹
  if [ ! -d "$path" ]; then
    mkdir $path
  fi
}

systemName=`uname  -a`
defaultStatic="./static"

#projectName="$(get_input "请输入您的项目名称(默认为 base-system-backend): " "base-system-backend")"
#staticPath="$(get_input "请输入项目静态文件地址[绝对路径](默认为项目根目录 static)：" $defaultStatic)"

if [[ $systemName =~ "Darwin" ]];then
    show_title "生成相关秘钥"
    show_msg "生成 secretKey"
    secretKey="$(< /dev/urandom LC_CTYPE=C tr -dc 'A-Za-z0-9!0$%=' | head -c64)"
    show_suc "$secretKey"

    show_msg "生成 aesKey"
    aesKey="$(< /dev/urandom LC_CTYPE=C tr -dc 'A-Za-z0-9!a$%S*()_+{}|:<>?=' | head -c16)"
    show_suc "$aesKey"

    show_title "创建静态文件目录"
    create_static_folder "$staticPath"
    check_error $? "创建文件夹失败"

    show_title "替换配置文件"
    show_msg "替换静态文件目录"
    sed -i "" -e "s#base_static#${staticPath}#g" ./config.yaml & wait

    show_msg "替换 secretKey"
    sed -i "" -e "s#base_secret_key#${secretKey}#g" ./config.yaml & wait

    show_msg "替换 aesKey"
    sed -i "" -e "s#base_aes_key#${aesKey}#g" ./config.yaml & wait
else
    show_title "生成相关秘钥"
    show_msg "生成 secretKey"
    secretKey="$(< /dev/urandom tr -dc 'A-Za-z0-9!0$%=' | head -c64)"
    show_suc "$secretKey"

    show_msg "生成 aesKey"
    aesKey="$(< /dev/urandom tr -dc 'A-Za-z0-9!a$%S*()_+{}|:<>?=' | head -c16)"
    show_suc "$aesKey"

    show_title "创建静态文件目录"
    create_static_folder "$staticPath"
    check_error $? "创建文件夹失败"

    show_title "替换配置文件"
    show_msg "替换静态文件目录"
    sed -i "s#base_static#${staticPath}#g" ./config.yaml & wait

    show_msg "替换 secretKey"
    sed -i "s#base_secret_key#${secretKey}#g" ./config.yaml & wait

    show_msg "替换 aesKey"
    sed -i "s#base_aes_key#${aesKey}#g" ./config.yaml & wait
fi


#show_msg "替换 base-system-backend 为 $projectName"
#find . -type f -not -name 'project_init.sh' -not -name 'README.md' -not -path './.git/*' -not -path './.idea/*' -exec sed -i 's/base-system-backend/'$projectName'/g' {} +
#
#show_msg "修改项目目录"
#cd ../ && mv ./base-system-backend ./$projectName
#
#show_titlle "删除 .git 文件"
#cd $projectName && rm -rf ./.git
