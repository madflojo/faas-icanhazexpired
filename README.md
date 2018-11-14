# OpenFaaS ICanHazExpired

An OpenFaaS function that checks if the remote host has expired certificates.

## Deploy

To deploy this function it is assumed you already have an OpenFaaS instance setup. 

First, clone this git repository locally.

```sh
$ git clone https://github.com/madflojo/faas-icanhazexpired.git
```

Second, deploy the function to your OpenFaaS instance.

```sh
$ faas-cli deploy -f ./icanhazexpired.yml --gateway=http://<YOUR-GATEWAY> 
```

Third, hit the function endpoint provided.

```sh
$ curl -d "expired.badssl.com:443" http://<YOUR-GATEWAY>:8080/function/faas-icanhazexpired
EXPIRED
```
