# lunch-manager-v2

##### Clone this repository
```
git clone git@github.com:CitizenF1/lunch-manager-v2.git 
```

##### Build docker image

```
sudo docker build -t bot -f Dockerfile .
```

##### Run docker container and past bot token 

```
sudo docker run --rm -d -e BOTTOKEN={TOKEN} bot
```