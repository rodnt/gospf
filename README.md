## gospf

### Install

```bash
go get github.com/incogbyte/gospf

# get from releases

https://github.com/incogbyte/gospf/releases/tag/v1.0

```

### Usage
```bash
gospf -d target.com
```

### What is SPF ?

* SPF stands for Sender Policy Framework.
It is a globally supported system that helps to protect domain names and their owners from being spoofed by spammers.
It does this by allowing a domain name owner to publish a policy which declares where emails can originate from.
It doesn't prevent spoofed emails from being sent, but it allows incoming email providers to verify the source of an email and take appropriate action.

### How does it work ?

* The domain name owner adds an SPF record to their domain name's DNS records;
* The SPF record declares all the IP addresses that are permitted to send email for that domain name;
* If the domain name does not have an SPF record, the message will be delivered as normal;
* If the domain name does have an SPF record, they will check if the IP address is listed in the SPF record;
* If the IP address is not listed in the SPF record, the message will be rejected or filtered;
* If the IP address is listed in the SPF record, the message will be delivered as normal


### Is SPF mandatory ?

* No, there is no mandatory requirement that forces you to add an SPF record to your domain name but there are many advantages:
    - It protects your domain name from spoofing and phishing;
    - Having an SPF record improves the deliverability of your emails;
    - Some email providers require you to have an SPF record before they will accept any email from your domain name;
    - You need to have an SPF record setup if you want to use similar systems such as DKIM signing and DMARC

### What is ASN ?

* An Autonomous System (AS) is a set of Internet routable IP prefixes belonging to a network or a collection of networks that are all managed, controlled and supervised by a single entity or organization. An AS utilizes a common routing policy controlled by the entity. The AS is assigned a globally unique 16 digit identification number一known as the autonomous system number or ASN一by the Internet Assigned Numbers Authority (IANA).
* Autonomous systems numbered one to 64511 are available by IANA for global use. The 64512 to 65535 series is reserved for private and reserved purposes. Autonomous Systems were introduced to regulate networking organizations such as Internet Service Providers (ISP), educational institutions and government agencies.

* Border Gateway Protocol (BGP) is the protocol that manages the routed peerings, prefix advertisement and routing of packets between different autonomous systems across the Internet. BGP uses the ASN to uniquely identify each system. In effect, BGP is the routing protocol for AS paths across the Internet.

### F.A.Q

* The tool can be used to get more information about your target.
* If your target is set well, you might not get anything interesting.
* If you get ip ranges ( cidr ) use probe tools such as httprobe, httpx or nmap to try to resolve these ips.
* This is not an exploit tool, this tool to aid your recon process or your infrastructure security

---


---
![s2Edwin](foo.jpeg)
