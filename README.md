# openfaas-certinfo

[![Build Status](https://travis-ci.org/stefanprodan/openfaas-certinfo.svg?branch=master)](https://travis-ci.org/stefanprodan/openfaas-certinfo)

OpenFaaS function that returns SSL/TLS certificate information for a given URL

### Usage

Deploy:

```bash
$ faas-cli deploy -f ./certinfo.yml --gateway=http://<GATEWAY-IP> 
```

Invoke:

```bash
$ echo -n "www.openfaas.com" | faas-cli invoke certinfo --gateway=<GATEWAY-IP>

Host 147.75.74.69
Port 443
Issuer Let's Encrypt Authority X3
CommonName www.openfaas.com
NotBefore 2017-10-06 23:54:56 +0000 UTC
NotAfter 2018-01-04 23:54:56 +0000 UTC
SANs [www.openfaas.com]
```

Local build:

```bash
$ git clone https://github.com/stefanprodan/openfaas-certinfo
$ cd openfaas-certinfo
$ faas-cli build -f ./certinfo.yml
```

Test local build:

```bash
$ docker run -dp 8080:8080 --name certinfo stefanprodan/certinfo
$ curl -d "cli.openfaas.com" localhost:8080

Host 147.75.74.69
Port 443
Issuer Let's Encrypt Authority X3
CommonName cli.openfaas.com
NotBefore 2017-10-07 11:55:06 +0000 UTC
NotAfter 2018-01-05 11:55:06 +0000 UTC
SANs [cli.openfaas.com]
```
