# intro

A InfluxDB console client, inspired by [mycli](https://github.com/dbcli/mycli).

## Usage:

	./ifcli -host https://<host>:3242 \ # Use http or https
		--user <user> \
		--pwd <password> \
		--db <db-name> \
		--disable-nil \        # Do not show nil field value, use ENABLE_NIL to enable nil value print
		--prompt "test-env"    # Prompt current connection name

Or

	./ifcli

If we have connect some InfluxDB before, `ifcli` will record them in `~/.ifclirc`, seems like this:

	["influx::https://yyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyy:3242"] # InfluxDB intance key: <user-name:host>
	  last_conn = 2020-02-04T10:28:52Z      # last connection time
	  last_exit = 2020-02-04T10:29:22Z      # last exit time
	  created = 2020-02-04T10:21:06Z        # connection create date
	  host = "https://yyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyy:3242" # influxdb instance host
	  user = "influx"                                                        # influxdb instance user name
	  # password has been encrypted 
	  password = "7d0bd5e7ea3500605b32845b41b14a40074017397687d088e1efa653c896decb83054db27552142b7fa0623f70"
	  default_db = ""

	# another instance
	["datafllux_influx::https://xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx:3242"]
	  last_conn = 2020-02-04T16:41:42Z
	  last_exit = 2020-02-04T16:41:43Z
	  created = 2020-02-04T16:41:42Z
	  host = "https://xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx:3242"
	  user = "datafllux_influx"
	  password = "b144649f35776818c0ccde0055fa1b5e411a2b0719a96b6939cbedac4902f991c0b970f8dc8f44b53b8dc0e5"
	  default_db = ""

We can use `CONN <influxdb-instance-key>` to switch between different InfluxDB instances


## Additional Commands

Additional commands are supported to make your happy:

- `DISABLE_NIL/ENABLE_NIL`: disalbe/enable `nil` value in console ouput
- `RESET_SUG`: It will add completer on measurement-name/db-name/field-name/tag-name, this may make the completer table too long and cause performance problem, use the command to clean these real-time-added completers, only keep InfluxDB reserved keyworkds
- Use `↑` and `↓` to select history. There is no such `~/ifcli-history` file, after `ifcli` exit, all history disappear
- Use `CONN ...` to switch between different InfluxDB instances

## Demo

![gif](./tty.gif)
