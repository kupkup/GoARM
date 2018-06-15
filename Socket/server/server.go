package main

import(
	"fmt"
	"net"
	"os"
	"time"
	"strings"	
)

const (  
	datetime   = "2006-01-02 15:04:05"   
    date       = "2006-01-02"  
    longtime   = "15:04:05"  
    shorttime  = "15:04"  
) 

func main(){


	var service = ":9500"

	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)	
	handleError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	handleError(err)

	fmt.Printf("[%s] server listen: %s \n", time.Now().Format(longtime), service)

	for{

		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		
		fmt.Printf("[%s] new connection: %s \n", time.Now().Format(longtime), conn.RemoteAddr().String())
		go handleClient(conn)

	}
}

func handleClient(conn net.Conn){

	conn.SetReadDeadline(time.Now().Add(2 * time.Minute)) //设置超时	
	request := make([]byte, 128)
	defer conn.Close() 

	for {
		readLen, err := conn.Read(request)

		if err != nil {
			fmt.Println(err)
			break
		}

		if( readLen == 0) {
			break
		}

		cmdText := strings.TrimSpace(string(request[:readLen]))
		
		fmt.Printf("[%s] %s> %s \n", time.Now().Format(longtime), conn.RemoteAddr().String(), cmdText)
		
		if ( cmdText == "exit" || cmdText == "quit" || cmdText == "q" ) {	

			fmt.Printf("[%s] server close! \n", time.Now().Format(longtime))
			os.Exit(0)
			
		} else if ( cmdText == "disconnect" || cmdText == "close") {
			
			fmt.Printf("[%s] disconnect: %s \n", time.Now().Format(longtime), conn.RemoteAddr().String())
			break

		} else if ( cmdText == "time" ){
		
			conn.Write([]byte(time.Now().Format(datetime)))	

		} else {

			conn.Write([]byte(cmdText))	

		}

    	request = make([]byte, 128)
	}	
}



func handleError(e error){
	if e != nil {
		fmt.Fprintln(os.Stderr, "Fatal Error: " + e.Error())
		os.Exit(1)
	}
}
