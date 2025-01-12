#!/bin/bash

log_error() {
    echo -e "$1" >&2
}

log_info() {
    echo -e "$1"
}

prompt_numeric_input() {
    local prompt="$1"
    local default_value="$2"
    local input

    while true; do
        read -p "$prompt" input

        if [[ -z "$input" ]]; then
            echo "$default_value"
            return 0
        fi

        if [[ "$input" =~ ^[0-9]+$ ]]; then
            echo "$input"
            return 0
        fi

        log_error "Invalid input. Please enter a numeric value."
    done
}

configure_environment() {
    local DEFAULT_HOST="localhost"
    local DEFAULT_POSTGRES_PORT="5432"
    local DEFAULT_POSTGRES_USER="postgres"
    local DEFAULT_POSTGRES_PASSWORD="postgres"
    local DEFAULT_POSTGRES_DB="crm_rules"
    local DEFAULT_REDIS_PORT="6379"
    local DEFAULT_REDIS_CHANNEL="crm_rules"

    log_info "Configuring environment variables..."
    log_info "Tip: Press Enter to use default values or skip configuration"

    read -p "Enter Postgres host [$DEFAULT_HOST]: " POSTGRES_HOST
    POSTGRES_HOST=${POSTGRES_HOST:-$DEFAULT_HOST}

    POSTGRES_PORT=$(prompt_numeric_input "Enter Postgres port [$DEFAULT_POSTGRES_PORT]: " "$DEFAULT_POSTGRES_PORT")

    read -p "Enter Postgres user [$DEFAULT_POSTGRES_USER]: " POSTGRES_USER
    POSTGRES_USER=${POSTGRES_USER:-$DEFAULT_POSTGRES_USER}

    read -p "Enter Postgres password [$DEFAULT_POSTGRES_PASSWORD]: " POSTGRES_PASSWORD
    POSTGRES_PASSWORD=${POSTGRES_PASSWORD:-$DEFAULT_POSTGRES_PASSWORD}

    read -p "Enter Postgres database [$DEFAULT_POSTGRES_DB]: " POSTGRES_DB
    POSTGRES_DB=${POSTGRES_DB:-$DEFAULT_POSTGRES_DB}

    REDIS_PORT=$(prompt_numeric_input "Enter Redis port [$DEFAULT_REDIS_PORT]: " "$DEFAULT_REDIS_PORT")

    read -p "Enter Redis channel [$DEFAULT_REDIS_CHANNEL]: " REDIS_CHANNEL
    REDIS_CHANNEL=${REDIS_CHANNEL:-$DEFAULT_REDIS_CHANNEL}

    cat > .env << EOF
POSTGRES_HOST=$POSTGRES_HOST
POSTGRES_PORT=$POSTGRES_PORT
POSTGRES_USER=$POSTGRES_USER
POSTGRES_PASSWORD=$POSTGRES_PASSWORD
POSTGRES_DB=$POSTGRES_DB
REDIS_PORT=$REDIS_PORT
REDIS_CHANNEL=$REDIS_CHANNEL
EOF

    log_info "Environment configuration complete."
}

main() {
    if [[ ! -f .env ]] || [[ ! -s .env ]]; then
        configure_environment
    else
        log_info ".env file already exists and is not empty. Skipping configuration."
    fi
}

main