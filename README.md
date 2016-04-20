# hamcode
Encode, decode and correct (single bit) a binary code using Hamming code.

## Install
```Bash
go install github.com/mahdavipanah/hamcode
```

## Usage
```Bash
$ hamcode help
Encode, decode and correct (single bit) a binary code using Hamming code.

Usage: hamcode [command] [binary code]

Available Commands:
  correct            Print the corrected binary code
  encode             Print the encoded data binary using Hamming code
  decode             Print the data binary code inside the input Hamming code
  help, -h, --help   Print the help

Available Options:
  --rtl		     Interpret the code from right to left

Author: Hamidreza Mahdavipanah
Repository: http://github.com/mahdavipanah/hamcode
```
### Example
```Bash
$ hamcode encode 1011 --rtl
1010101
$ hamcode decode 1010101 --rtl
1011
```
