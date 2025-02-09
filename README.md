## peek

a **better looking**, **colorful** Curl alternative, perfect for testing REST APIs


### Install

to install, run the following command

```sh
curl -sSL https://raw.githubusercontent.com/0verread/peek/main/install.sh | bash
```

or, to install manually

```sh
git clone https://github.com/0verread/peek
cd peek
go install
```

### Usage

```sh
peek <url>
```
example:

```sh
peek jsonplaceholder.typicode.com/todos/1
```

### Why do I need this

I test all the rest apis using curl. lately, I found curl is not being enough. Make no mistake, Curl is still super simple yet powerful tool and peek is not a replacement of Curl. Peek is for people like me who don't use Postman like heavy bloated software to test simple rest apis, yet want something simple and powerful like Curl with all the feature Curl is missing for testing rest apis. I thought of building peek for myself but later, decided to turn it into fully CLI tool.


### LICENSE

This project is under [MIT License](./LICENSE)