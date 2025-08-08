Docker-compose for postgresql helps you setup postresql database and PGadmin using a single docker-compose file.
Makesure you have docker and docker-composed install in your host machine. You can find a guide to install docker in the offical site [Docker-offical](https://docs.docker.com/get-docker/). 

__NOTE__: 
>It is important to note that this file is not intended for production use.

</br>

Once you have installed docker and docker-compose on your local machine. Follow the following guide. 
---

__Clone the project__ 
```git
git https://github.com/Cavertech/OrgHub.git
cd OrgHub/Docker-Compose/dev-env-postgresql
```
__Start Docker-compose__
```shell
$ docker-compose up -d
```

__Check if the containers are working__
```shell
$ docker-compose ps
```
![Docker-ps](../../Docker-Compose/dev-env-postgresql/img/docker-ps.png)

__Stop Docker-compose__

```shell
$ docker-compose down 
```

You can modify following values if you wish to:
--- 
```docker
services:
  db:
    ports:
      - "5432:5432"
    environment:
      POSTGRES_USER: user-name
      POSTGRES_PASSWORD: strong-password
  pgadmin:
    ports:
      - "8888:80"
    environment:
      PGADMIN_DEFAULT_EMAIL: user-name@domain-name.com
      PGADMIN_DEFAULT_PASSWORD: strong-password
```
Now the use PGAdmin tool, open the browser and access ``http://localhost:8888/``. Enter the username and password for PGAdmin.

Configuring PGAdmin to server:
---
1. Login to the admin page via the url provide above. 
2. Enter the user name password that you have provide in the ``docker-compose.yml`` file (_lines 21-22_). 
3. Click on ``Add New Server`` under Quick links .
    ![psadmin](docker/pg-db/img/add-server.png)
4. Please enter database server name.

    __General__ tab
    ```
    Name : OrgHub
    Server Group : Server 
    ``` 
    __Connection__ tab
    ```
    Host Name/address : local_pgdb
    Port : 5432
    Username : developer 
    Password : dev123
    ```
5. right click on the ___OrgHub___ to create a database within the Server.
    ![OrgHub-DB](docker/pg-db/img/create-db.png)
    
    add ___OrgHub___ as DB
    ![Config-db](docker/pg-db/img/config-db.png)

6. Setup complete.
    ![Overview-OrgHub-DB](docker/pg-db/img/Overview-Org-Hub-DB.png)