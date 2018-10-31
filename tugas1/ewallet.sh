#!/bin/bash
if [ $1 == "getSaldo" ]; then
    if [ -z $2 ]; then
        echo "missing user_id"
    else
        $(curl -H "Content-Type: application/json" -X POST -d '{"user_id": '$2'}' localhost/ewallet/getSaldo)
    fi
elif [ $1 == "getTotalSaldo" ]; then
    if [ -z $2 ]; then
        echo "missing user_id"
    else
        $(curl -H "Content-Type: application/json" -X POST -d '{"user_id": '$2'}' localhost/ewallet/getTotalSaldo)
    fi
elif [ $1 == "transfer" ]; then
    if [[ -z $2 || -z $3 || -z $4 ]]; then
        echo "transfer destination, user_id, or transfer value missing"
    else
        $(curl -H "Content-Type: application/json" -X POST -d '{"user_id": "'$3'", "nilai": '$4'}' localhost/ewallet/transferMinus)
        $(curl -H "Content-Type: application/json" -X POST -d '{"user_id": "'$3'", "nilai": '$4'}' $2/ewallet/transfer)
    fi
elif [ $1 == "register" ]; then
    $(curl -H "Content-Type: application/json" -X POST -d '{"user_id": '$2', "nama": '$3'}' localhost/ewallet/register)
elif [ $1 == "ping" ]; then
    $(curl -H "Content-Type: application/json" -X POST localhost:8080/ewallet/ping)
else
    echo "command not recognized"
fi