# gopwgen

Password Generator written in Go

Running
-----------

```

-=[  PWGEN - The Go Password Generator  ]=-
Usage: 
pwgen     [-help] [-size N] [-pwn N] [-kn str] [-dict] [-A] [-a] [-@] [-0]
-help -h  : this message 
-size N: change the password size from 16 to N (default 16) 
-pwn  N: number of passwords to create 
-dict  : make random passwords in dictionary mode
-kn str: known pieces of the password (eg. alice).
		   (not implemented yet)  

-A      : add abcdefghijklmnopqrstuvwxy for the password char possibilities
-a      : add ABCDEFGHIJKLMNOPQRSTUVWXY for the password char possibilities
-0      : add 012345678 for the password char possibilities
-@      : add @#$%&.,?+=- for the password char possibilities
""      : all the combinations above within 
```
