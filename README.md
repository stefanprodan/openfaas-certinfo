# openfaas-certinfo

[![Build Status](https://travis-ci.org/stefanprodan/openfaas-certinfo.svg?branch=master)](https://travis-ci.org/stefanprodan/openfaas-certinfo)

OpenFaaS function that returns SSL/TLS certificate information for a given URL

### Usage

Deploy example for OpenFaaS on Kubernetes:

```bash
faas-cli deploy --name=certinfo \
    --image=stefanprodan/certinfo:latest \		
    --fprocess="./certinfo" \		
    --network=openfaas-fn \		
    --gateway=http://<GATEWAY-IP> 
```

Invoke example:

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

