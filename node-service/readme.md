# Docker with NodeJS + expres

Simple rest api with NodeJS and Express. I used typescript for unknown reason but you can do it in regular JS.

## How to get started

Have docker installed on your system. 

Then simple run: `docker build -t docker-expressjs .` (docker-expressjs) is just an example, you can use anything. `docker build -t potato .`. But keep in mind to have meaningful docker names, so it makes sense once name is given.

After docker is built, we can run: `docker run -d 8080:8080 docker-express` 
We are running `docker run -d ...` the `-d`
`--detach , -d		Run container in background and print container ID`


we just running that container in the background process with the container port of 8080 and opening it to 8080 with the given container name: `docker-express` in this example, again it can be anything.

After this, you're done! your first simple Docker contrainer that is powered with NodeJS+Express.

More work to be done...