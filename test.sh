IP="127.0.0.1"
PORT="8080"
CERT="certs/client.pem"
KEY="certs/client.key"

while true
do 
    clear

    DATA="$(openssl rand -hex 64)"

    go run ./main.go --ip $IP --port $PORT --mode client --data $DATA --key $KEY --cert $CERT
    sleep 0.3
done