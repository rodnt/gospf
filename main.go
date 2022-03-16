package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/jamesog/iptoasn"
)

/*
Author: Rodolfo Tavares
*/

func all_records(domain string) string {
	addrs, err := net.LookupTXT(domain)
	if err != nil {
		return err.Error()
	}

	records := ""

	for _, val := range addrs {
		records += val + "\n"
	}

	if records == "" {
		fmt.Println("Cant find anything :/")
		os.Exit(1)
	}

	if strings.Contains(records, "spf1") {
		return records
	}

	return ""
}

func print_data(stg_array []string) {
	fmt.Println("")

	for _, value := range stg_array {
		split_value := strings.Split(value, ":")
		fmt.Println(split_value[1])
	}
	fmt.Println("")
}

func get_assets(records string) []string {

	mechanisms := []string{"ip4", "ip6", "ptr", "include", "a", "include", "mx", "exists"}
	data := []string{}

	splited_records := strings.Split(records, " ")

	for _, value := range splited_records {
		for _, value_mecha := range mechanisms {
			split_value := strings.Split(value, ":")
			if split_value[0] == value_mecha {
				data = append(data, value)
			}
		}
	}
	print_data(data)
	return data
}

func data_asn(assets []string) {
	for _, data := range assets {
		split_data := strings.Split(data, ":")
		if strings.Contains(split_data[0], "ip4") || strings.Contains(split_data[0], "ip6:") {
			get_asn(split_data[1])
		}
	}
}

func get_asn(cidr_ip string) {
	split_ip_cidr := strings.Split(cidr_ip, "/")
	asn_details(split_ip_cidr[0])
}

func asn_details(raw_ip string) {

	ip, err := iptoasn.LookupIP(raw_ip)
	if err != nil {
		fmt.Println("-> Failed to lookup ip")
	}

	fmt.Println()
	fmt.Println("-={ ASN Details }=-")
	fmt.Println("ANS: ", ip.ASName)
	fmt.Println("Ip: ", ip.IP)
	fmt.Println("BGP Prefix: ", ip.BGPPrefix)
	fmt.Println("Coutry: ", ip.Country)
	fmt.Println("Registry: ", ip.Registry)

}

func main() {

	var domain = ""
	flag.StringVar(&domain, "d", "", "domain name to search for: ex: example.com")
	flag.Parse()

	if domain == "" {
		fmt.Println("[+] Error, Missing domain name")
		flag.Usage()
		os.Exit(1)
	}

	records := all_records(domain)
	data_asn(get_assets(records))

}
