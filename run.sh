unset http_proxy
unset https_proxy
pkill -9 webconsole
nohup ./webconsole > /dev/null 2>&1 &
