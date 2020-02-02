# intro

A InfluxDB console client, inspired by [mycli](https://github.com/dbcli/mycli).

## Usage:

	./ifcli -host https://<host>:3242 \ # Use http or https
		--user <user >\
		--pwd <password> \
		-db <db-name> \
		--disable-nil \                  # Do not show nil field value, use ENABLE_NIL to enable nil value print
		--prompt "influx-db-test-env"    # Prompt current connection name

## Demo

![gif](./tty.gif)
