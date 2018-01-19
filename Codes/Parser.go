package main

import (
	"strconv"		//For number parsing
	"fmt"
	"net"
	"net/url"
)
func number_parser(){
	f, _ := strconv.ParseFloat("1.234", 64)			//Parse float 64 says how many bits of precision to parse
	fmt.Println(f)

	i, _ := strconv.ParseInt("123", 0, 64)		//Parse integer with base 0 and fit in 64 bits
	fmt.Println(i)

	d, _ := strconv.ParseInt("0x1c8", 0, 64)	//Can infer hexadecimal numbers as well
	fmt.Println(d)

	u, _ := strconv.ParseUint("789", 0, 64)		//Parse u int
	fmt.Println(u)

	k, _ := strconv.Atoi("135")								//'Atoi' convenient keyword for parsing integers with base 10
	fmt.Println(k)

	_, e := strconv.Atoi("What")								//return errors and bad input
	fmt.Println(e)
}

func url_parser(){
	s := "https://user:pass@host.com:5432/path?shirt_color=red#f"

	u, err := url.Parse(s)		// Parse the URL and ensure there are no errors.
	if err != nil {
		panic(err)
	}

	fmt.Println(u.Scheme)		//Access the scheme

	fmt.Println(u.User)				//call user
	fmt.Println(u.User.Username())	//user contains username and password
	p, _ := u.User.Password()
	fmt.Println(p)

	fmt.Println(u.Host)				//If the host is present it will contain bith the hostname and the port
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println(host)
	fmt.Println(port)

	fmt.Println(u.Path) 			//Extract path and the fragment after the #
	fmt.Println(u.Fragment)

	fmt.Println(u.RawQuery)			//To get all the parameters in a key value pair format use the raw query command
	m, _ := url.ParseQuery(u.RawQuery)	//You can also parse these into a map
	fmt.Println(m)
	fmt.Println(m["shirt_color"][0]) // As the values are available in the form of slices you can access it using its index position
}

func main() {

	number_parser()

	fmt.Println("________________________________________________")

	url_parser()

}


/*
1.234
123
456
789
135
strconv.ParseInt: parsing "What": invalid syntax
________________________________________________
https
user:pass
user
pass
host.com:5432
host.com
5432
/path
f
shirt_color=red
map[shirt_color:[red]]
red
 */