# intro

A InfluxDB console client, inspired by [mycli](https://github.com/dbcli/mycli).

# Download

Only 64-bit available, if your have 32-bit CPU, you can build your own binary

- [Mac](./bin/mac/ifcli)
- [Linux](./bin/linux/ifcli)
- [Windows](./bin/windows/ifcli.exe)

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

	["influx::https://host-to-influxdb:8086"] # InfluxDB intance key: <user-name:host>
	  last_conn = 2020-02-04T10:28:52Z        # last connection time
	  last_exit = 2020-02-04T10:29:22Z        # last exit time
	  created = 2020-02-04T10:21:06Z          # connection create date
	  host = "https://host-to-influxdb:3242"  # influxdb instance host
	  user = "influx"                         # influxdb instance user name
	  # password has been encrypted 
	  password = "7d0bd5e7ea3500605b32845b41b1"
	  default_db = ""

	# another instance
	["datafllux_influx::https://host-to-influxdb:8086"]
	  last_conn = 2020-02-04T16:41:42Z
	  last_exit = 2020-02-04T16:41:43Z
	  created = 2020-02-04T16:41:42Z
	  host = "https://host-to-influxdb:3242"
	  user = "datafllux_influx"
	  password = "b144649f35776818c0ccde0055fa"
	  default_db = ""

We can use `CONN <influxdb-instance-key>` to switch between different InfluxDB instances

## Additional Commands

Additional commands/tips to make your happy:

- `DISABLE_NIL/ENABLE_NIL`: disalbe/enable `nil` value in console ouput
- `RESET_SUG`: It will add completer on measurement-name/db-name/field-name/tag-name, this may make the completer table too long and cause performance problem, use the command to clean these real-time-added completers, only keep InfluxDB reserved keyworkds
- Use `↑` and `↓` to select history. There is no such `~/ifcli-history` file, after `ifcli` exit, all history disappear
- Use `CONN ...` to switch between different InfluxDB instances
- Support Emacs-style line operation, such as `ctrl+w` to delete the word before the cursor, `ctrl+e` to move cursor to line end, `ctrl+l` to clean screen, and so on
- Windows support(recommand `cmd.exe` or `powershell.exe`, other terminal not tested)
- Use `tee output.file` to redirect output to file
- Use `TSCNT <db-name> <since-when>` to show DB's time series count. If `db-name` not sepcified, use current DB (if set); if `since-when` not set, default to `5m` (latest 5 minutes)
- Use `MOVE <db1>.<rp1>.<measurement> <db2>.<rp2>.<measurement2>` to backup data to another databases, or within same DB, with different RP or measurement name
- Use `BENCH <sql>` to benchmark SQL with runing 10 times
- Use `BENCHN <n> <sql>` to benchmark SQL with running `n` times

## Demo

![gif](./tty.gif)
