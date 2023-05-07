SERVICE_NAME=$1
POD=$2
PORT_APP=$3
PORT_METRICS_APP=$4
SIZE_CLUSTER=$5

for ((i = 0; i <= ${SIZE_CLUSTER}; i++)); do
  docker-compose run --rm -d \
    --name ${SERVICE_NAME}_${POD}_${PORT_APP}_${PORT_METRICS_APP} \
    -e POD_UUID=${POD} \
    -e PORT_APP=:${PORT_APP} \
    -e PORT_METRICS_APP=:${PORT_METRICS_APP} \
    -p ${PORT_APP}:${PORT_APP} \
    -p ${PORT_METRICS_APP}:${PORT_METRICS_APP} \
    ${SERVICE_NAME}

  let "PORT_APP += 1"
  let "PORT_METRICS_APP += 1"
done