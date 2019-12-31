# gregif

## Overview
Compress images(jpg, png, gif).
Set the reduction ratio and resize the image.

## Usage

```
NAME:
   grimg - Compress images(jpg, png, gif)

USAGE:
   main [global options] command [command options] [arguments...]

VERSION:
   0.0.1

AUTHOR:
   hmarf

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --input value, -i value        [required] string
                                  input file path (default: "None")
   --output value, -o value       string
                                  output file path [default] output.(file type) (default: "output")
   --compression value, -c value  [required] float(0.1 ~ 0.9)
                                  Specify compression ratio (default: 0)
   --help, -h                     show help
   --version, -v                  print the version
```

## important information
**This program is not completed.**  
Some gifs contain noise.The cause is being analyzed.  

<img src="https://github.com/hmarf/gregif/blob/master/img/input_d.gif?raw=true" width="500px">  
<img src="https://github.com/hmarf/gregif/blob/master/img/output_d.gif?raw=true" width="250px">

## Example
png and jpeg
- Resize './img/input.png' to './output.png'
```
go run cmd/main.go -i ./img/input.png -c 0.5
```
<img src="https://github.com/hmarf/gregif/blob/master/img/input.png?raw=true" width="150px">
<img src="https://github.com/hmarf/gregif/blob/master/img/output.png?raw=true" width="75px">

- Resize './img/input.gif' to './output.gif'
```
go run cmd/main.go -i ./img/input.gif -c 0.5
```
<img src="https://github.com/hmarf/gregif/blob/master/img/input.gif?raw=true" width="150px">
<img src="https://github.com/hmarf/gregif/blob/master/img/output.gif?raw=true" width="75px">

- Resize './img/input_d.gif' to './output_d.gif'
```
go run cmd/main.go -i ./img/input_d.gif -c 0.5
```
<img src="https://github.com/hmarf/gregif/blob/master/img/input_d.gif?raw=true" width="500px">  
<img src="https://github.com/hmarf/gregif/blob/master/img/output_d.gif?raw=true" width="250px">
