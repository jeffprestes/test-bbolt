# test bbolt

This is a MVC CRUD in REST Application built on top of Mercurius framework ( <https://github.com/novatrixtech/mercurius> ) to test BBolt database and Storm Framework.

It has very good performance and very simple and flexible sintaxe.

## Structure of the Project

```
/conf 
Application configuration including environment-specific configs

/conf/app
Middlewares and routes configuration

/handler
HTTP handlers

/locale
Language specific content bundles

/lib
Common libraries to be used across your app

/model
Models

/public
Web resources that are publicly available

/public/templates
Jade templates

/repository
Database comunication following repository pattern

main.go
Application entry

user.db
BBolt Database for Users

## Build 

To build it uses env GO111MODULE=on go build for Mac or GO111MODULE=on go build for Linux or define GO111MODULE on your Windows
```