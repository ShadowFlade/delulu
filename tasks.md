### Tasks to Beta version
* security (verifying that request was made from the site) - check from postman - now checking only with referer, later mb need to finish cookie verification somehow 
* ci/cd (make prod branch, when push to prod it builds and reboots service)
* recalc stats (учесть инфляцию чтобы было актуально к 2024 году) + news stats about salary in 2024
### Tasks to 1.0.0
* make everything local (scripts, fonts etc)
* yandex metrika ?
* crontask for calcing site stats (make table for this)
* make local stats block
* caching
* indexing


# CURRENT PROBLEM

nginx proxies request and because of that golang app thinks that all requests are coming from inside (remote ip is always my ip - need to give original remote_ip to the server)