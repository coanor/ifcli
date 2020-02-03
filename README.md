# intro

A InfluxDB console client, inspired by [mycli](https://github.com/dbcli/mycli).

## Usage:

	./ifcli -host https://<host>:3242 \ # Use http or https
		--user <user> \
		--pwd <password> \
		--db <db-name> \
		--disable-nil \        # Do not show nil field value, use ENABLE_NIL to enable nil value print
		--prompt "test-env"    # Prompt current connection name

## Additional Commands

Additional commands are supported to make your happy:

- `DISABLE_NIL/ENABLE_NIL`: disalbe/enable `nil` avalue in console ouput
- `RESET_SUG`: It will add completer on measurement-name/db-name/field-name/tag-name, this may make the completer table too long and cause performance problem, use the command to clean these real-time-added completer, just keep InfluxDB reserved keyworkds
- Use ↑ and ↓ to select history. There is no such `~/ifcli-history` file, after `ifcli` exit, all history disappear

## Demo

![gif](./tty.gif)
