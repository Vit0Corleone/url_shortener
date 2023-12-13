Simple Url Shortener

Before starting you should set actual config data in .env file

To start run in terminal: make && ./app

Endpoints:
- [GET] host/{alias} => redirects you to the url connected with this alias
- [POST] host/url => create new alias for url using json post data
- [DELETE] host/url/{alias} => deletes everything from database connected with alias

