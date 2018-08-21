#!/bin/bash

set -e

# mongoexport --port=6002 -d suizhou -c col_user -o col_user.dat
# mongoimport --port 6002 -d test -c col_robot col_user.dat

# mongoimport -d test -c students --type csv --headerline --file students_csv.dat

# mongodump --collection COLLECTION --db DB_NAME
# mongorestore -h dbhost -d dbname --directoryperdb dbdirectory

usage() {
    echo " ./ctrl cmd "
    echo " cmd : mongodb (start|stop) "
    echo " cmd : mongo-console "
    echo " cmd : mongodump 备份整个数据库"
    echo " cmd : mongorestore (DBNAME PATH) 恢复指定位置备份"
    echo " cmd : coldump 备份bson格式集合"
    echo " cmd : colimport (DBNAME PATH) 导入bson格式集合 "
}

# server conf
workDir=$(cd `dirname $0`; pwd)

cd $workDir

# mongodb conf
db_host=127.0.0.1:6002
#db_name=huiyin-dev
#collections_array=(col_env col_game col_game_record col_id_gen col_log_chip col_log_chip_today col_log_pay_today col_pkten col_stat_record col_trends col_user col_user_record col_user_stat_record t_action t_last_id t_perm t_role t_role_perm t_user t_user_role)
db_name=happy-test
#collections_array=(col_apply_cash col_env col_game col_game_record col_id_gen col_log_accounting_log col_log_agent_fee col_log_agent_fee_log col_log_chip col_log_chip_today col_log_give_record col_log_pay_today col_pkten col_stat_record col_trends col_user col_user_record col_user_stat_record t_action t_last_id t_perm t_role t_role_perm t_user t_user_role)
#去掉col_log_chip
#collections_array=(col_apply_cash col_env col_game col_game_record col_id_gen col_log_accounting_log col_log_agent_fee col_log_agent_fee_log col_log_chip_today col_log_give_record col_log_pay_today col_pkten col_stat_record col_trends col_user col_user_record col_user_stat_record t_action t_last_id t_perm t_role t_role_perm t_user t_user_role)
#happy-test
collections_array=(col_activity col_game col_id_gen col_lucky col_notice col_shop col_task col_login_prize col_trade_record col_user col_user_info col_log_day_profit col_log_pay_today col_log_bank col_log_profit_order)
collections_log_array=(col_role_record col_room_record col_round_record col_log_task col_log_coin col_log_diamond col_log_login col_log_regist col_log_profit col_log_sys_profit)
collections_admin_array=(t_action t_user tt_label tt_onlink tt_operation_log tt_perm tt_role tt_role_perm tt_sequence tt_smsMessage tt_user tt_user_role)
remote_host=xx.xx.xx.xx
remote_addr=/home/database/backup/
db_conf=mongodb.conf
db_path=/home/database
db_data=${db_path}/data
db_logpath=${db_path}/mongodb.log
db_pidfile=${db_path}/mongodb.pid

MONGODB_DIR="/home/mongodb/bin"
BACKUP_DIR="${db_path}/backup"
TAR="/bin/tar"

get_curr_last_month() {
    #echo `date -d last-month +"%Y-%m-%d"`
    echo `date +"%Y-%m-%d"`
}

get_curr_day() {
    echo `date +"%Y-%m-%d"`
}

get_curr_time() {
    echo `date +"%Y-%m-%d %H:%M:%S"`
}

get_curr_time_str() {
    echo `date +"%Y%m%d%H%M%S"`
}

MONGO="${MONGODB_DIR}/mongo"
MONGOD="${MONGODB_DIR}/mongod"
BSONDUMP="${MONGODB_DIR}/bsondump"
MONGODUMP="${MONGODB_DIR}/mongodump"
MONGORESTORE="${MONGODB_DIR}/mongorestore"
MONGOEXPORT="${MONGODB_DIR}/mongoexport"
MONGOIMPORT="${MONGODB_DIR}/mongoimport"
LOGFILE="${BACKUP_DIR}/backup_database.log"

CURRDAY=`get_curr_day`
CURRDAY_LASTMONTH=`get_curr_last_month`
BACKUPPARENTDIR="${BACKUP_DIR}/${db_name}"

BACKUPDIR="${BACKUPPARENTDIR}/${CURRDAY}"
BACKUPFILENAME="${BACKUPPARENTDIR}/${db_name}_${CURRDAY}.tar.gz"

BACKUPDIR_LASTMONTH="${BACKUPPARENTDIR}/${CURRDAY_LASTMONTH}"
BACKUPFILENAME_LASTMONTH="${BACKUPPARENTDIR}/${db_name}_${CURRDAY_LASTMONTH}.tar.gz"

CURR_TIME_STR=`get_curr_time_str`
#BACKUPDIR_STR="${BACKUPPARENTDIR}/${CURR_TIME_STR}"
BACKUPDIR_STR="${BACKUPPARENTDIR}/${db_name}_coldump_${CURR_TIME_STR}"
#BACKUPFILENAME_STR="${BACKUPPARENTDIR}/${db_name}_${CURR_TIME_STR}.tar.gz"
BACKUPFILENAME_STR="${BACKUPPARENTDIR}/${db_name}_coldump_${CURR_TIME_STR}.tar.gz"
BACKUPDIR_STR2="${BACKUPPARENTDIR}/${db_name}_coldumplog_${CURR_TIME_STR}"
BACKUPFILENAME_STR2="${BACKUPPARENTDIR}/${db_name}_coldumplog_${CURR_TIME_STR}.tar.gz"
BACKUPDIR_STR3="${BACKUPPARENTDIR}/${db_name}_coldumpadmin_${CURR_TIME_STR}"
BACKUPFILENAME_STR3="${BACKUPPARENTDIR}/${db_name}_coldumpadmin_${CURR_TIME_STR}.tar.gz"

save_log() {
    echo "" >> ${LOGFILE}
    echo "Time: "`get_curr_time`"   $1 " >> ${LOGFILE}
}

# test mysql connection
mysql_conn() {
    mysql --host=127.0.0.1 --port=8081 -uroot -p`cat ~/.mysql` -e "show databases;"
}

# mongodb – this script starts and stops the mongodb daemon
# chkconfig: - 85 15
# description: MongoDB is a non-relational database storage system.
# processname: mongodb
mongodb_() {
    test -x $DAEMON || exit 0
    #set -e
    case "$1" in
        start)
            echo -n "Starting MongoDB... "
            ${MONGOD} -f ${db_conf}
            ;;
        stop)
            echo -n "Stopping MongoDB... "
            # pid=`ps -o pid,command ax | grep mongod | awk '!/awk/ && !/grep/ {print $1}'`;
            pid=`ps aux | grep mongod | awk '!/awk/ && !/grep/ {print $2}'`;
            if [ "${pid}" != "" ]; then
                kill -2 ${pid};
            fi
            ;;
        *)
            echo "Usage: ./mongodb {start|stop}" >&2
            exit 1
            ;;
    esac
    exit 0
}

remote_backup_sync() {
    if [ ! -e "${1}" ]
    then
        echo "${1} file not exist! "
        exit 1
    fi
    echo "${1} sync to ${remote_host}"
    save_log "${1} sync to ${remote_host}"
    scp $1 root@${remote_host}:${remote_addr}
    save_log "${1} sync done"
}

mongorestore_() {
    name=$1
    if [ "${name}" == "" ]
    then
        echo "${$name} db not exist! "
        exit 1
    fi
    path=$2
    # 目录或文件已经存在，则程序退出
    if [ ! -d "${path}" ]
    then
        echo "${path} directory not exist! "
        exit 1
    fi
    echo "Start restore ${path} to db ${name}! "
    save_log "Start restore ${path} to db ${name}! "
    ${MONGORESTORE} -h ${db_host} -d ${name} ${path}
    save_log "RESTORE DATABASE ${name} SUCCEED"
}

mongodump_() {
    # 目录或文件已经存在，则程序退出
    if [ -d "${BACKUPDIR}" ]
    then
        echo "${BACKUPDIR} directory had exist! "
        exit 1
    fi
    if [ -e "${BACKUPFILENAME}" ]
    then
        echo "${BACKUPFILENAME} file had exist! "
        exit 1
    fi

    if [ -d "${BACKUPDIR_LASTMONTH}" ]
    then
        echo "${BACKUPDIR_LASTMONTH} directory had exist! "
        #rm -rf ${BACKUPDIR_LASTMONTH}
        save_log "${BACKUPDIR_LASTMONTH} directory already remove! "
    fi

    if [ -e "${BACKUPFILENAME_LASTMONTH}" ]
    then
        echo "${BACKUPFILENAME_LASTMONTH} file had exist! "
        #rm -rf ${BACKUPFILENAME_LASTMONTH}
        save_log "${BACKUPFILENAME_LASTMONTH} file already remove! "
    fi

    mkdir -p ${BACKUPDIR}

    echo "Start dump ${db_name} all data to DIR ${BACKUPDIR}"
    save_log "Start dump ${db_name} all data to DIR ${BACKUPDIR}"
    cd ${BACKUPPARENTDIR}
    ${MONGODUMP} --host=${db_host} -d ${db_name} -o ${BACKUPDIR}
    echo "Start compress to ${BACKUPFILENAME}"
    save_log "Start compress to ${BACKUPFILENAME}"
    cd ${BACKUPPARENTDIR}
    ${TAR} cvf ${BACKUPFILENAME} ${CURRDAY}

    save_log "BACKUP DATABASE ${db_name} SUCCEED"

    #remote_backup_sync ${BACKUPFILENAME}
}

collections_import_by_json() {
    name=$1
    if [ "${name}" == "" ]
    then
        echo "${$name} db not exist! "
        exit 1
    fi
    path=$2
    # 目录或文件已经存在，则程序退出
    if [ ! -d "${path}" ]
    then
        echo "${path} directory not exist! "
        exit 1
    fi

    for col in ${collections_array[@]};
    do
        file=${path}/${col}.json
        if [ ! -e "${file}" ]
        then
            echo "${file} file not exist! "
            exit 1
        fi
        echo "Start import collection ${file} to db ${name}! "
        save_log "Start import collection ${file} to db ${name}! "
        ${MONGOIMPORT} -h ${db_host} -d ${name} -c ${col} --type json --file ${file}
    done

    save_log "RESTORE DATABASE ${name} SUCCEED"
}

collections_dump_to_json() {
    # 目录或文件已经存在，则程序退出
    if [ -d "${BACKUPDIR_STR}" ]
    then
        echo "${BACKUPDIR_STR} directory had exist! "
        exit 1
    fi
    if [ -e "${BACKUPFILENAME_STR}" ]
    then
        echo "${BACKUPFILENAME_STR} file had exist! "
        exit 1
    fi

    mkdir -p ${BACKUPDIR_STR}

    cd ${BACKUPPARENTDIR}
    for col in ${collections_array[@]};
    do
        echo "Start dump collection ${col} to DIR ${BACKUPDIR_STR}"
        save_log "Start dump collection ${col} to DIR ${BACKUPDIR_STR}"
        ${MONGODUMP} --host=${db_host} -d ${db_name} --collection ${col} -o ${BACKUPDIR_STR}
        #echo "bson to json ${col} to DIR ${BACKUPDIR_STR}"
        #save_log "bson to json ${col} to DIR ${BACKUPDIR_STR}"
        #${BSONDUMP} --outFile ${BACKUPDIR_STR}/${db_name}/${col}.json ${BACKUPDIR_STR}/${db_name}/${col}.bson
    done

    echo "Start compress to ${BACKUPFILENAME_STR}"
    save_log "Start compress to ${BACKUPFILENAME_STR}"
    cd ${BACKUPPARENTDIR}
    #${TAR} cvf ${BACKUPFILENAME_STR} ${CURR_TIME_STR}/${db_name}/*.json
    ${TAR} cvf ${BACKUPFILENAME_STR} ${BACKUPDIR_STR}

    save_log "BACKUP DATABASE ${db_name} SUCCEED"

    #remote_backup_sync ${BACKUPFILENAME_STR}
}

collections_log_dump_to_json() {
    # 目录或文件已经存在，则程序退出
    if [ -d "${BACKUPDIR_STR2}" ]
    then
        echo "${BACKUPDIR_STR2} directory had exist! "
        exit 1
    fi
    if [ -e "${BACKUPFILENAME_STR2}" ]
    then
        echo "${BACKUPFILENAME_STR2} file had exist! "
        exit 1
    fi

    mkdir -p ${BACKUPDIR_STR2}

    cd ${BACKUPPARENTDIR}
    for col in ${collections_log_array[@]};
    do
        echo "Start dump collection ${col} to DIR ${BACKUPDIR_STR2}"
        save_log "Start dump collection ${col} to DIR ${BACKUPDIR_STR2}"
        ${MONGODUMP} --host=${db_host} -d ${db_name} --collection ${col} -o ${BACKUPDIR_STR2}
        #echo "bson to json ${col} to DIR ${BACKUPDIR_STR2}"
        #save_log "bson to json ${col} to DIR ${BACKUPDIR_STR2}"
        #${BSONDUMP} --outFile ${BACKUPDIR_STR2}/${db_name}/${col}.json ${BACKUPDIR_STR2}/${db_name}/${col}.bson
    done

    echo "Start compress to ${BACKUPFILENAME_STR2}"
    save_log "Start compress to ${BACKUPFILENAME_STR2}"
    cd ${BACKUPPARENTDIR}
    #${TAR} cvf ${BACKUPFILENAME_STR2} ${CURR_TIME_STR}/${db_name}/*.json
    ${TAR} cvf ${BACKUPFILENAME_STR2} ${BACKUPDIR_STR2}

    save_log "BACKUP DATABASE ${db_name} SUCCEED"

    #remote_backup_sync ${BACKUPFILENAME_STR2}
}

collections_admin_dump_to_json() {
    # 目录或文件已经存在，则程序退出
    if [ -d "${BACKUPDIR_STR3}" ]
    then
        echo "${BACKUPDIR_STR3} directory had exist! "
        exit 1
    fi
    if [ -e "${BACKUPFILENAME_STR3}" ]
    then
        echo "${BACKUPFILENAME_STR3} file had exist! "
        exit 1
    fi

    mkdir -p ${BACKUPDIR_STR3}

    cd ${BACKUPPARENTDIR}
    for col in ${collections_admin_array[@]};
    do
        echo "Start dump collection ${col} to DIR ${BACKUPDIR_STR3}"
        save_log "Start dump collection ${col} to DIR ${BACKUPDIR_STR3}"
        ${MONGODUMP} --host=${db_host} -d ${db_name} --collection ${col} -o ${BACKUPDIR_STR3}
        #echo "bson to json ${col} to DIR ${BACKUPDIR_STR3}"
        #save_log "bson to json ${col} to DIR ${BACKUPDIR_STR3}"
        #${BSONDUMP} --outFile ${BACKUPDIR_STR3}/${db_name}/${col}.json ${BACKUPDIR_STR3}/${db_name}/${col}.bson
    done

    echo "Start compress to ${BACKUPFILENAME_STR3}"
    save_log "Start compress to ${BACKUPFILENAME_STR3}"
    cd ${BACKUPPARENTDIR}
    #${TAR} cvf ${BACKUPFILENAME_STR3} ${BACKUPDIR_STR3}/${db_name}/*.json
    ${TAR} cvf ${BACKUPFILENAME_STR3} ${BACKUPDIR_STR3}

    save_log "BACKUP DATABASE ${db_name} SUCCEED"

    #remote_backup_sync ${BACKUPFILENAME_STR3}
}

mongo_console() {
    ${MONGO} --host=${db_host}
}

case $1 in
    mongodb)
        mongodb_ $2;;
    mongodump)
        mongodump_;;
    mongorestore)
        mongorestore_ $2 $3;;
    coldump)
        collections_dump_to_json;;
    coldumplog)
        collections_log_dump_to_json;;
    coldumpadmin)
        collections_admin_dump_to_json;;
    colimport)
        collections_import_by_json $2 $3;;
    mongo-console)
        mongo_console;;
    mysql)
        mysql_conn;;
    *)
        usage;;
esac
