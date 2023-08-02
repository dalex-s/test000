package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	//	"strings"
	//	"io"
	//	"log"
	//	"os"
	"github.com/jackc/pgx/v4/pgxpool"
)

func main() {

	f, err := os.OpenFile("current.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	wrt := io.MultiWriter(os.Stdout, f)
	log.SetOutput(wrt)
	log.Println(" START ")

	/* fmt.Print("Password: ")
	    bytepw, err := term.ReadPassword(int(syscall.Stdin))
	    if err != nil {
	        os.Exit(1)
	    }
	   pass := string(bytepw)
	   fmt.Println("")
	   fmt.Printf("\nYou've entered: %q\n", passs)
	*/

	/*fmt.Println("Continue?")
	  xx,err := fmt.Println(yesNo())
	  //log.Println(xx)
	  if xx == 6  {
	  fmt.Println("Вы выбрали NO. Выход!")
	  os.Exit(0)
	  } else {
	  fmt.Println("Вы выбрали YES. Продолжаем!")
	  }*/
	//fmt.Println(yesNo())

	//fmt.Println(len(os.Args),os.Args)
	db := flag.String("db", "template1", "database")
	user := flag.String("user", "postgres", "user")
	//    pass := flag.String("pass", passs, "pass")
	ip := flag.String("ip", "172.1.10.88", "ip")
	port := flag.String("port", "5432", "port")
	// hinfo := flag.Int("hinfo", 0, " host information level, levels 0,1,2")
	// dbinfo := flag.Int("dbinfo", 0, "db information level,levels: 0,1,2")
	// info := flag.Bool("info", true, "short information")
	// dselect := flag.String("dselect", "all", "select one database")
	flag.Parse()
	// setdselect := *dselect
	// sethinfo := *hinfo
	// setinfo := *info
	// levelDbInfo := *dbinfo
	conString0 := fmt.Sprintf("postgres://%s:%s@%s:%s", *user, "", *ip, *port)
	conString := fmt.Sprintf("%s/%s", conString0, *db)
	dbPool, err := pgxpool.Connect(context.Background(), conString) //databaseUrl)

	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer dbPool.Close()

} //main
