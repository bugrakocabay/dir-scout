Scan URL directories with a wordlist.


**Available flags:**
```
Flags:
  -h, --help              help for endpoint-discovery
  -o, --output string     output file name
  -s, --success           filters out status codes 400 and above
  -v, --verbosity int16   Provide a value between 0-10 (default 5)
      --version           version for endpoint-discovery
```

------------


**Example:**

`./dir-scout facebook.com data/common.txt -v 10 -s`
```
██████╗ ██╗██████╗       ███████╗ ██████╗ ██████╗ ██╗   ██╗████████╗
██╔══██╗██║██╔══██╗      ██╔════╝██╔════╝██╔═══██╗██║   ██║╚══██╔══╝
██║  ██║██║██████╔╝█████╗███████╗██║     ██║   ██║██║   ██║   ██║   
██║  ██║██║██╔══██╗╚════╝╚════██║██║     ██║   ██║██║   ██║   ██║   
██████╔╝██║██║  ██║      ███████║╚██████╗╚██████╔╝╚██████╔╝   ██║   
╚═════╝ ╚═╝╚═╝  ╚═╝      ╚══════╝ ╚═════╝ ╚═════╝  ╚═════╝    ╚═╝          v0.0.1
======================================================
[+] Target:     facebook.com
[+] Wordlist:   data/common.txt
======================================================
+ /.bash_history 200
+ / 200
+ /.bashrc 200
+ /.config 200
+ /.cache 200
+ /.cvsignore 200
+ /.cvs 200

```
