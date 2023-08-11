<h1 align="center">Orion</h1>
<img src="https://cdna.artstation.com/p/assets/images/images/025/789/352/original/pixel-jeff-galaxy-far-far-away.gif?1586928273" />
<img src="https://user-images.githubusercontent.com/73097560/115834477-dbab4500-a447-11eb-908a-139a6edaec5c.gif"></p>


## Description
Orion is a simple event planning app, in this case there are lots of relationships between tables to interact with. For example, to retrieve user data when creating an event or joining an event, and this is an example project that you can test if you want to learn about database relations. This project uses golang and the mysql database, I originally wanted to use mongoDB but it seems that in this case it's better to use a database that supports and is efficient in making relationships.

This project still uses the rest api, I plan to make a gRPC version, for the database and endpoint schemas, you can see in this image: 

<img src="https://media.discordapp.net/attachments/1125397316971528272/1139484058351575090/Blank_diagram_1.png?width=993&height=611" />

## Usage
### 1.1 Usage with Golang
Run this command to running download all module
```go
go mod tidy
```
then running file `server.go` to run the server
```go
go run server.go
```
and the output will be like this
```
2023/08/08 14:26:20 Server running on :8080
```

### 1.2 Usage with Docker 
Run this command to download docker-compose (this command for linux)
```bash
sudo apt install docker-compose -y
```
then running this command
```
docker-compose up -d
```
and service will be running on port `8080`

<img src="https://user-images.githubusercontent.com/73097560/115834477-dbab4500-a447-11eb-908a-139a6edaec5c.gif"></p>

> For example schema to input request and see response can see in [`example`](https://github.com/rulanugrh/orion/tree/master/example) folder