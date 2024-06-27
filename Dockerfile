version: '3.9'

services:
  db:
    image: postgres
    restart: always
    ports:
        - 5431:5432
    environment:
      POSTGRES_PASSWORD: admin21
      POSTGRES_USER: postgres
      POSTGRES_DB: courseHub
  adminer:
    image: adminer
    restart: always
    ports:
      - 8081:8080

  minio:
    image: quay.io/minio/minio
    command: server /data --console-address ":9001"
    restart: always
    ports:
      - 9000:9000
      - 9001:9001
    volumes:
      -  ~/minio/data:/data
    environment:
        - MINIO_ROOT_USER=admin
        - MINIO_ROOT_PASSWORD=admin213



