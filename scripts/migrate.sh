#!/bin/sh

set -e

if [ -f ".env" ]
then
      source .env
fi

if [ -z "$APP_NAME" ]
then
      echo "\$APP_NAME is empty"; exit 1;
else
      echo "\$APP_NAME = $APP_NAME";
fi

if [ -z "$DB_USER" ]
then
      echo "\$DB_USER is empty"; exit 1;
else
      echo "\$DB_USER = $DB_USER";
fi

if [ -z "$DB_PASS" ]
then
      echo "\$DB_PASS is empty"; exit 1;
fi

if [ -z "$DB_MASTER_HOST=lbc-db
" ]
then
      echo "\$DB_MASTER_HOST is empty"; exit 1;
else
      echo "\$DB_MASTER_HOST = $DB_MASTER_HOST";
fi

if [ -z "$DB_MASTER_PORT" ]
then
      echo "\$DB_MASTER_PORT is empty"; exit 1;
else
      echo "\$DB_MASTER_PORT = $DB_MASTER_PORT";
fi

if [ -z "$DB_NAME" ]
then
      echo "\$DB_NAME is empty"; exit 1;
else
      echo "\$DB_NAME = $DB_NAME";
fi

DB_URL="$DB_DRIVER://$DB_USER:$DB_PASS@tcp($DB_MASTER_HOST:$DB_MASTER_PORT)/$DB_NAME"

MIGRATIONS_DIR="/tmp/$APP_NAME/migrations"
echo "MIGRATIONS_DIR = $MIGRATIONS_DIR"
rm -rf "$MIGRATIONS_DIR"
mkdir -p "$MIGRATIONS_DIR"

cp ./migrations/*.sql "$MIGRATIONS_DIR"

MIGRATE_ARGS="up"
if [ -z "$MIGRATE_ARGS" ]
then
  echo "Migrating up..."
else
    echo "Migrating $MIGRATE_ARGS..."
    MIGRATE_ARGS="$MIGRATE_ARGS"
fi

make MIGRATIONS_DIR="$MIGRATIONS_DIR" MIGRATE_ARGS="$MIGRATE_ARGS" MIGRATE_URL="$DB_URL" .migrate
