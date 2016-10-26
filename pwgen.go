package main

import(
    "runtime"
    "fmt" 
    "math/rand"
    "time"
    "sync"
    "flag"
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
    info.version     = 1.3
    info.title       = "-=[  PWGEN - The Go Password Generator  ]=- "
    info.contact     = "twitter: @danielfl"
    info.github      = "/danielfl"
}

func showChar(pass chan string,  ch string, kn string, size int, goGroup *sync.WaitGroup) { 

    password:=""
    for i:=1 ; i <= size ; i++ {
        tstamp := time.Now().UnixNano() 
        rand.Seed(tstamp)
        luckn:=rand.Intn(len(ch))
        password+=string(ch[luckn]) 
    }
    fmt.Println(password)
    goGroup.Done()
    pass <-"ok"
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
    dictmode:= 0
    kn      := ""

    //         0    '    10   '    20   '    30   '    40   '    50   '    60   '    70   '   
    possib := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789@#$%&.,?+=-" 

    //parameters
    dictPtr := flag.Bool("dict",false,"make random password dictionary mode")
    sizePtr := flag.Int("size", 16, "change the password size from 16 to N (default 16)")
    pwnPtr  := flag.Int("pwn", 1, "number of passwords to create")

    str:="add "+possib[0:25]+" %s for the password char possibilities"
    APtr :=        flag.Bool("A",false,str)

    str="add "+possib[26:51]+" %s for the password char possibilities"
    aPtr:= flag.Bool("a",false,str)

    str="add "+possib[52:61]+" %s for the password char possibilities"
    zPtr:= flag.Bool("0",false,str)

    str="add "+possib[62:len(possib)]+" %s for the password char possibilities"
    atPtr:= flag.Bool("@",false,str)

    flag.Parse()

    if *dictPtr {
        dictmode = 1
    }
    pwn = *pwnPtr
    size = *sizePtr

    if *aPtr  {
        chars += possib[0:25]
    }
    if *APtr  {
        chars += possib[26:51]
    }
    if *zPtr  {
        chars += possib[52:62]
    }
    if *atPtr {
        chars += possib[62:len(possib)]
    }

    if len(chars) == 0 {
        chars=possib
    }

    if dictmode != 1 {
        showHeader(size, pwn, chars)
    }

    runtime.GOMAXPROCS(2)
    pass := make(chan string)
    for loop := 0 ; loop < pwn ; loop++ {
        goGroup := new (sync.WaitGroup)

        goGroup.Add(1)
        go showChar(pass, chars, kn, size, goGroup) // create a proc
        runtime.Gosched()

        goGroup.Wait()
    }

    for  {
        <-pass
        break
    }

    if dictmode != 1 {
        fmt.Printf("\n\n%d password(s) created..\n", pwn)
    }
}
