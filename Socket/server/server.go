package main

import(
	"fmt"
	"net"
	"os"
	"time"
	"strings"	
)

func main(){


	var service = ":9500"

	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)	
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	fmt.Println("server listen: " + service)

	for{

		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		
		fmt.Println("new connection: " + conn.RemoteAddr().String())
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

		cmdText := strings.TrimSpace(string(request[:readLen]))

		if( readLen == 0) {
			break

		} else if ( cmdText == "exit" || cmdText == "quit" || cmdText == "q" ) {	

			fmt.Println("server close!")
			os.Exit(0)
			
		} else if ( cmdText == "disconnect" || cmdText == "close") {
			
			fmt.Println("disconnected: " + conn.RemoteAddr().String())
			break

		} else if ( cmdText == "time" ){
		
			conn.Write([]byte(time.Now().String()))	

		} else {

			fmt.Println(conn.RemoteAddr().String() + "> " + cmdText + "\r\n")
			conn.Write([]byte(cmdText))	

		}

    	request = make([]byte, 128)
	}	
}



func checkError(e error){
	if e != nil {
		fmt.Fprintln(os.Stderr, "Fatal Error: " + e.Error())
		os.Exit(1)
	}
}
