## CORS Proxy
A transparent proxy which passes on requests to other web servers and adds CORS headers to the response.

## Usage

Prepend the proxy url to the actual url you want to call `proxy_url?real_url`.

## Example

Your cors proxy is running on `http://localhost:9999/` and you want to call `https://dynamic.lunanode.com/vm/list`

 
replace the call to `https://dynamic.lunanode.com/vm/list` with `http://localhost:9999/?https://dynamic.lunanode.com/vm/list`

By default this proxy only whitelists https://dynamic.lunanode.com. In case you want to use
your own whitelist please mount a custom tsconf.json onto /tsconf.json in the container

`docker run -p 9999:9999 -v yourTsConf.json:/tsconf.json schulterklopfer/cors_proxy` 