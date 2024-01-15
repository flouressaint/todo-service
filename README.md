# todo-service

### Run
    docker-compose up -d   
    go run cmd/app/main.go

### get access token
![Alt text](assets/README/image-11.png)

### Auth  
![Alt text](assets/README/image.png)
![Alt text](assets/README/image-7.png)
token expired
    ![Alt text](assets/README/image-8.png)
invalid token
    ![Alt text](assets/README/image-10.png)

### Get todos
![Alt text](assets/README/image-9.png)

### Create todo
![Alt text](assets/README/image-2.png)
![Alt text](assets/README/image-3.png)

### Change todo
![Alt text](assets/README/image-4.png)
![Alt text](assets/README/image-5.png)

### Delete todo
![Alt text](assets/README/image-6.png)

### Пользователь может получить доступ только к своим задачам
![Alt text](assets/README/image-12.png)

## SAST
### goSec
![Alt text](assets/README/image-13.png)
Issues: 0

### Bearer
![Alt text](assets/README/image-14.png)
42 checks were run and no failures were detected.
