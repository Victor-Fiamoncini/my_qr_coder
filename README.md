# My QR Coder

An API rest to generate a qr-code based on the provided text from request body. The qr-code will be stored in a S3 public Bucket.

## Tools 🪓

- Golang v1.24.1
- Fiber-V2 v2.52.6
- AWS-SDK-V2 v1.36.3

## How to start

### Environment Variables

Create a `.env` file from `.env.example` file in project root folder and fill the variables with your AWS S3 settings:

```env
AWS_ACCESS_KEY_ID=your_access_key
AWS_SECRET_ACCESS_KEY=your_secret_key
AWS_REGION=your_region
AWS_BUCKET_NAME=your_bucket_name
```

### Start the Docker container

- First build the Docker image:

```bash
docker build -t my-qr-code-api .
```

- Run the container:

```bash
docker run --env-file .env -p 8080:8080 my-qr-code-api
```

### About the  API

- All the supported API routes were written in `text.http` file.

----------
Released in 2025. This project is under the MIT license.

By [Victor B. Fiamoncini](https://github.com/Victor-Fiamoncini) ☕️
