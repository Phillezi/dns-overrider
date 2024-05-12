# DNS Overrider

DNS Overrider is a tool that allows you to override DNS settings locally and set your own DNS records. It's particularly useful for development purposes when you need to test DNS configurations or simulate specific DNS responses.

Uses [miekg/dns](https://github.com/miekg/dns) dns library.

## Features

- Override DNS settings locally
- Set custom DNS records
- Useful for development and testing purposes

## Prerequisites
- Go

## Installation/usage

```bash
git clone https://github.com/Phillezi/dns-overrider.git
cd dns-overrider
go build .
sudo ./dns-overrider # needs to be run with root to bind port 53

```

## Configuration

DNS Overrider can be configured using a configuration file (e.g., `override.conf`). Here's an example configuration file:

```conf
@externalDNS: 8.8.8.8
example.com: 127.0.0.1
example2.com: 192.168.0.1
```

## Contributing

Contributions are welcome! If you have any ideas, suggestions, or bug fixes, feel free to open an issue or submit a pull request.
