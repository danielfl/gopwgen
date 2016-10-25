package main

import(
    "runtime"
    "fmt" 
    "math/rand"
    "time"
    "sync"
    "os"
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
           go showChar(i, chars, kn, goGroup) // create a proc
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
