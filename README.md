# explorer
Blockchain explorer for the IRIS network


# ENV
DB_URL: mongo url , 192.168.150.7:27017
DB_USER: mongo user, iris 
DB_PASSWORD: mongo password, irispassword

# Docker Run
docker run -p 8080:8080 -e ${ENV Variables} explorer