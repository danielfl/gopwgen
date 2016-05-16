package main

import(
    "runtime"
    "fmt" 
    "math/rand"
    "time"
    "sync"
    "os"
    "strconv"
) 

var chars string
var info struct {
    author string
    version float32
    title string 
    contact string
    github string
}

func defineInfo() {
    info.author      = "Daniel Ferreira de Lima"
    info.version     = 1.2
    info.title       = "-=[  PWGEN - The Go Password Generator  ]=-\n"
    info.contact     = "twitter: @danielfl"
    info.github      = "/danielfl"
}

func showChar(num int,  ch string, kn string, goGroup *sync.WaitGroup) { 
    tstamp := time.Now().UnixNano() 
    rand.Seed(tstamp)
    luckn:=rand.Intn(len(ch))
    fmt.Printf("%c", ch[luckn]) 

    goGroup.Done()
}
func showHeader(s int, n int, ch string){ 
    fmt.Println(info.title)
    //header
    fmt.Printf("conf len[%d] pwn[%d] chars[%s]\n\n", s,n, ch)

    for i := 1 ; i <= int(s / 10)+1 ; i++ {
        fmt.Print("    '   ",i*10)
    }
    fmt.Printf("\n")

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
    defineInfo() 

    chars   :=""
    size    :=16
    pwn     :=1
    nargs   := len(os.Args)
    dictmode:= 0
    kn      := ""

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
        }else if arg == "-dict" {
            dictmode = 1
        }else if arg == "-kn" {
            kn = os.Args[i+1]
            nargs -= 2 
        }else if arg == "-help" {
            fmt.Println(info.title,  info.version)
            fmt.Println("Usage: ")
            fmt.Printf("%s     [-help] [-size N] [-pwn N] [-kn str] [-dict] [A] [a] [@] [0]\n", os.Args[0])
            fmt.Printf("-help  : this message \n")
            fmt.Printf("-size N: change the password size from 16 to N (default 16) \n")
            fmt.Printf("-pwn  N: number of passwords to create \n")
            fmt.Printf("-dict  : make random password dictionary mode\n")
            fmt.Printf("-kn str: known pieces of the password (eg. alice).\n\t\t   (not implemented yet)  \n\n") 
            fmt.Printf("A      : add %s for the password char possibilities\n", possib[0:25])
            fmt.Printf("a      : add %s for the password char possibilities\n", possib[26:51])
            fmt.Printf("0      : add %s for the password char possibilities\n", possib[52:61])
            fmt.Printf("@      : add %s for the password char possibilities\n", possib[62:len(possib)])
            fmt.Printf("null   : all the combinations above within \n")

            os.Exit(0)
        }
    }

    if nargs == 1 {
        chars=possib
    }

    if dictmode != 1 {
        showHeader(size, pwn, chars) 
    } 

    //playing with paralelism just for fun 
    for loop := 0 ; loop < pwn ; loop++ { 
        //do the magic
        goGroup := new (sync.WaitGroup) 
        goGroup.Add(size)

        runtime.GOMAXPROCS(4)
        iterations := size 
        for i := 0; i<iterations; i++ {
           go  showChar(i, chars, kn, goGroup) // create a proc
        } 
        runtime.Gosched()

        //Wait for the password creation being finished
        goGroup.Wait()

        fmt.Println("")
    }
    if dictmode != 1 {
        fmt.Printf("\n\n%d password(s) created..\n", pwn)
    }
} 
