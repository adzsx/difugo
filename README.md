# Difugo
Difugo is a Directory Fuzzer focused on simplicity and speed.

# Usage:
difugo [options] 

## Options:

| Flag           | Argument    | Purpose                                             | Default Value |
| -------------- | ----------- | --------------------------------------------------- | ------------- |
| -h, --help     |             | Display help message                                |               |
| -u, --url      | URL/IP      | Set the hosts URL                                   |               |
| -w, --wordlist | File Path   | Set Wordlist for fuzzed parameter                   |               |
| -r, --robots   |             | Scan for directories in robots.txt                  |               |
| -s, --suffix   | string      | Set the suffix for the fuzzed parameter             |               |
| -f, --filter   | code1 code2 | Filter out specific status codes                    | 403, 404      |
| -c, --code     | code1 code2 | Only show specific status codes (empty to show all) |               |
| -a, --async    | int         | Use n seperate goroutines/threads                   | 32            |
| -v, --verbose  |             | Verbose mode                                        |               |

# Examples
fuzz google.com with wordlist.txt, use only 1 thread to decrease load for google<br>
`$ difugo -u google.com -w ./wordlist.txt -v -a 1`<br>
<br>
scan google.com for directories in google.com/robots.txt, use 100 threads, only show 200Ok responses<br>
`$ difugo -u google.com -r -a 100 -c 200`<br>
<br>
fuzz search parameter for google<br>
`$ difugo -u google.com -p /search?q= -w ./wordlist.txt`<br>
Note: This is completely useless, as every search will return a redirect or a 200, because it is a google search