[[ "${ENV_NAME}" = "stg" || "${ENV_NAME}" = "prod" ]] && echo "latest" || echo "latest-${ENV_NAME}"
