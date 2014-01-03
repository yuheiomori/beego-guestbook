# Beego Guestbook

This is sample guestbook app with beego.


## upload to heroku

```
$ heroku create -b https://github.com/kr/heroku-buildpack-go.git
$ heroku addons:add cleardb:ignite
$ git push heroku master
$ heroku open
```