# DNS Overrider

DNS Overrider is a tool that allows you to override DNS settings locally and set your own DNS records. It's particularly useful for development purposes when you need to test DNS configurations or simulate specific DNS-names with other addresses.
Or maybe you are just curious about the dns requests that your system is sending, you can use this program to log the requests.

Uses [miekg/dns](https://github.com/miekg/dns) dns library.

## Features

- Override DNS settings locally
- Log DNS queries
- Useful for development and testing purposes to simulate domain names

## Prerequisites
- Go

## Installation/usage

```bash
git clone https://github.com/Phillezi/dns-overrider.git
cd dns-overrider
go build .
sudo ./dns-overrider # needs to be run with root to bind port 53

```
Then just change your dns address setting to `127.0.0.1` and youre set!

## Configuration

DNS Overrider can be configured using the configuration file `override.conf`. 
You can use some pi-hole blocklists, like [this one](https://github.com/nickoppen/pihole-blocklists/blob/master/blocklist-advert_01.txt), just save it and write the relative path to it like in the example below.
Here's an example configuration file:

```conf
@externalDNS: 8.8.8.8
@blockLists: blocklists/blocklist-facebook.txt, blocklists/blocklist-advert_01.txt
example.com: 127.0.0.1
example2.com: 192.168.0.1
example3.com: 192.168.0.2
```

## Contributing

Contributions are welcome! If you have any ideas, suggestions, or bug fixes, feel free to open an issue or submit a pull request.
