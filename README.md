# Fake-Request-Gen
A very simple program that creates random HTTPS traffic

### Requirements
- Golang (Tested with 1.20, Should work with any relatively recent version)
- A browser (Default: Firefox)
- xvfb (Not needed if you are running it in a display server)

### Usage
```
git clone https://github.com/VMpc/Fake-Request-Gen
cd Fake-Request-Gen
go build -a -gcflags=all="-l -B -C" -ldflags "-s -w"
```

### Running
This will run the program with all defaults, see below to change them 
```
./Fake-Request-Gen
```
To run without a display server (ex usage: running in a tty)
```
xvfb-run ./Fake-Request-Gen 
```
To run without a browser (Uses the default http client, good for low memory usage but less realistic looking)
```
./Fake-Request-Gen --browser none
```

Command line args
```
Usage of ./Fake-Request-Gen:
  -browser string
        Sets the browser to use (default "firefox")
  -browserargs string
        Sets the browser args to use (default "--headless --private-window")
  -url string
        Sets the specified URL/File path (default "https://moz.com/top-500/download/?table=top500Domains")
  -breaktime int
        Sets the maximum amount of time the program will view pages for in seconds (default 60)
  -viewtime int
        Sets the maximum amount of time it views a page for in seconds (default 60)
  -b string
        Sets the browser to use (default "firefox")
  -ba string
        Sets the browser args to use (default "--headless --private-window")
  -bt int
        Sets the maximum amount of time the program will view pages for in seconds (default 60)
  -u string
        Sets the specified URL/File path (default "https://moz.com/top-500/download/?table=top500Domains")
  -vt int
        Sets the maximum amount of time it views a page for in seconds (default 60)
```
