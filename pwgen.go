package main

import(
    "runtime"
    "fmt" 
    "strings"
    "math/rand"
    "time"
    "sync"
    "os"
    "strconv"
) 

var chars string

func showChar(num int,  ch string, goGroup *sync.WaitGroup) { 
    tstamp := time.Now().UnixNano() 
    rand.Seed(tstamp)
    luckn:=rand.Intn(len(ch))
    fmt.Printf("%c", ch[luckn]) 

    goGroup.Done()
}

 func stringInSlice(str string, list []string) bool {
    for _, v := range list {
        if v == str {
            return true
        }
    }
    return false
 }

func main() { 

    chars:="" 
    size:=16
    pwn :=1
    nargs := len(os.Args)
    //         0    '    10   '    20   '    30   '    40   '    50   '    60   '    70   '   
    possib := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789@#$%&.,?+=-" 

    //read args and add the corresponding chars to the password
    for i, arg := range os.Args {
        if arg == "a" {
            chars += possib[0:25]
        }else if arg == "A" {
            chars += possib[26:51]
        }else if arg == "0" {
            chars += possib[52:62]
        }else if arg == "@" {
            chars += possib[62:len(possib)]
        }else if arg == "-pwn" { 
            p,err := strconv.Atoi(os.Args[i+1])
            _ = err
            pwn = p
            nargs -= 2
        }else if arg == "-size" { 
            s,err := strconv.Atoi(os.Args[i+1])
            _ = err
            size = s
            nargs -= 2
        }else if arg == "-help" {
            fmt.Println("PWGEN - The Go password creator: ")
            fmt.Println("Usage: ")
            fmt.Printf("%s     [-help] [-size N] [A] [a] [@] [0]\n", os.Args[0])
            fmt.Printf("-help  : this message \n")
            fmt.Printf("-size N: change the password size from 16 to N \n")
            fmt.Printf("-pwn  N: number of passwords to create \n")
            fmt.Printf("A      : add %s for the password char possibilities\n", possib[0:25])
            fmt.Printf("a      : add %s for the password char possibilities\n", possib[26:51])
            fmt.Printf("0      : add %s for the password char possibilities\n", possib[52:61])
            fmt.Printf("@      : add %s for the password char possibilities\n", possib[62:len(possib)])
            fmt.Printf("null   : all the combinations above within a 16 char password")

            os.Exit(0)
        } 
    }

    if nargs == 1 {
          chars=possib
    }

    fmt.Println("Size :", size)
    fmt.Println("Chars:", chars)

    fmt.Println("Starting...\n")

    //header
    fmt.Println("N 0    '    10   '    20   '    30   '    40   '    50")
    fmt.Printf("N %s\n", strings.Repeat(possib[52:62], 6))

    //playing with paralelism just for fun 
    for loop := 0 ; loop < pwn ; loop++ {
        fmt.Print  ("P  ")

        //do the magic
        goGroup := new (sync.WaitGroup) 
        goGroup.Add(size)

        runtime.GOMAXPROCS(4)
        iterations := size 
        for i := 0; i<iterations; i++ {
           go  showChar(i, chars, goGroup) // create a proc
        } 
        runtime.Gosched()

        //Wait for the password creation being finished
        goGroup.Wait()

        fmt.Println("")
    }
    fmt.Println("\n\nPassword(s) created..")
} 
