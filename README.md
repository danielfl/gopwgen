# gopwgen

Password Generator written in Go

Running
-----------

```

PWGEN - The Go password creator: 
Usage: 
pwgen     [-help] [-size N] [A] [a] [@] [0]
-help  : this message 
-size N: change the password size from 16 to N 
-pwn  N: number of passwords to create 
A      : add abcdefghijklmnopqrstuvwxy for the password char possibilities
a      : add ABCDEFGHIJKLMNOPQRSTUVWXY for the password char possibilities
0      : add 012345678 for the password char possibilities
@      : add @#$%&.,?+=- for the password char possibilities
null   : all the combinations above within a 16 char password
```
