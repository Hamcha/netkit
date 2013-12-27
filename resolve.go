package main;

import (
	"fmt"
	"net"
	"flag"
)

func main(){
	resolveCNAME := flag.Bool("c", false, "Resolve CNAME host")
	resolveMX := flag.Bool("m", false, "Resolve MX hosts")
	resolveNS := flag.Bool("n", false, "Resolve Nameservers")
	resolveRev := flag.Bool("r", false, "Perform Reverse Lookup on IP Address")

	flag.Parse()
	host := flag.Args()
	if len(host) < 1 {
		fmt.Println("Usage: resolve [-cnmrh] HOSTNAME")
		return 
	}
	var addrs []string
	var err error

	if *resolveCNAME {
		var name string
		name, err = net.LookupCNAME(host[0])
		addrs = []string{name}
	} else if *resolveNS {
		var ns []*net.NS
		ns, err = net.LookupNS(host[0])
		addrs = make([]string,len(ns))
		for i := range ns {
			addrs[i] = ns[i].Host
		}
	} else if *resolveMX {
		var mx []*net.MX
		mx, err = net.LookupMX(host[0])
		addrs = make([]string,len(mx))
		for i := range mx {
			addrs[i] = fmt.Sprintf("%d %s",mx[i].Pref,mx[i].Host)
		} 
	} else if *resolveRev {
		addrs, err = net.LookupAddr(host[0])
	} else {
		addrs, err = net.LookupHost(host[0])
	}

	if err != nil {
		fmt.Println(err)
		return
	}
	for i := range addrs {
		fmt.Println(addrs[i]);
	}
}