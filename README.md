# Internet Outage Logger

This is a simple little program I threw together to track when my internet connection was cutting out. Works by pinging a url every N seconds (set with interval flag and defaults to 10 seconds).

The program has three flags you can pass:
- interval to check (in seconds): `-interval or -i`
- log file path `-logfile or -l`
- option to print to the terminal (stdout) `-print or -p`

```sh
./outage-logger -interval 5 -logfile outages.txt -print
```

You can also use the short versions of those flags:

```sh
./outage-logger -i 5 -l outages.txt -p
```

The log output will look like this:

```txt
2024/08/25 11:04:54 Internet connection restored
2024/08/25 11:20:22 Internet connection lost
2024/08/25 11:20:48 Internet connection restored
2024/08/25 11:31:53 Internet connection restored
2024/08/25 11:35:42 Internet connection lost
2024/08/25 11:36:02 Internet connection restored
2024/08/25 12:03:45 Internet connection lost
2024/08/25 12:04:14 Internet connection restored
2024/08/25 12:52:48 Internet connection lost
2024/08/25 12:53:03 Internet connection restored
```