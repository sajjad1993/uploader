# Hackethon Challenge

With this document you should have received three executable files:

+ [`ascii_linux_x64`](ascii_linux_x64): executable binary compiled for linux x64 platforms
+ [`ascii_darwin_x64`](ascii_darwin_x64): executable binary compiled for MacOS x64 platforms
+ [`ascii_windows_x64`](ascii_windows_x64.exe): executable binary compiled for Windows x64 platforms

## The Challenge

The proposed challenge is to implement an ASCII art hosting web service. The web service should expose a REST HTTP API interface.

The provided executable will simulate a client's interactions with the web service. They will connect through port `4444` on `localhost`.  Once connected, clients will be sending three kinds of requests in the following sequence:

1. **Image registration**: The client registers an image for upload. In order to register an image, the client provides its SHA256 hash for further reference. Registering an already existing image should result in an error (`409 Conflict`).
2. **Image chunks upload**: The client splits the image content in a sequence of chunks and uploads them. It sends each chunk separately as a JSON payload. Each chunk has an ID indicating its position in the sequence
3. **Downloading the complete image**: The client downloads the image from the web service. It then computes the downloaded image hash and compares it to the registered image. It is expected that an image could be downloaded multiple times.

The executable's output for a single image upload sequence will look like the following:

```bash
./ascii_linux_x64 -files 1
time="2019-02-28T11:55:15+01:00" level=info msg="registering image with server" image_sha256=8a99030199b315fe8e4cf93d93478facdf1801a0ddb0d9bc1325961597a42a3f
time="2019-02-28T11:55:15+01:00" level=info msg="uploading image chunks" chunks_count=3 image_sha256=8a99030199b315fe8e4cf93d93478facdf1801a0ddb0d9bc1325961597a42a3f
time="2019-02-28T11:55:15+01:00" level=info msg="chunk upload: OK" chunk_id=1 chunk_size=256 image_sha256=8a99030199b315fe8e4cf93d93478facdf1801a0ddb0d9bc1325961597a42a3f
time="2019-02-28T11:55:16+01:00" level=info msg="chunk upload: OK" chunk_id=0 chunk_size=256 image_sha256=8a99030199b315fe8e4cf93d93478facdf1801a0ddb0d9bc1325961597a42a3f
time="2019-02-28T11:55:16+01:00" level=info msg="chunk upload: OK" chunk_id=2 chunk_size=187 image_sha256=8a99030199b315fe8e4cf93d93478facdf1801a0ddb0d9bc1325961597a42a3f
time="2019-02-28T11:55:16+01:00" level=info msg="succesfully retrieved image" image_sha256=8a99030199b315fe8e4cf93d93478facdf1801a0ddb0d9bc1325961597a42a3f
```

## Notes

1. You should create two services: one for uploading the chunkcs and one for merging them.
2. these services should communicate via gRPC and use protobuf as data format.
3. the chunks must be stored in a shared repository.
4. Your services needs to be resilient to failures and to be able to restart or retry if an error occurred.

## The API

Our executable expects your HTTP API to implement the following endpoints:

+ **Registering an image**:
  + **method**: `POST`
  + **URI**: `/image`
  + **Content-Type**: `application/json`
  + **Request Body**:

      ```json
      {
        "sha256": "abc123easyasdoremi...",
        "size": 123456,
        "chunk_size": 256
      }
      ```

  + **Responses**:
    | Code                       |              Description           |
    |----------------------------|------------------------------------|
    | 201 Created                | Image successfully registered       |
    | 409 Conflict               | Image already exists               |
    | 400 Bad Request            | Malformed request                  |
    | 415 Unsupported Media Type | Unsupported payload format         |

+ **Uploading an image chunk**:
  + **method**: `POST`
  + **URI**: `/image/<sha256>/chunks`
  + **Content-Type**: `application/json`
  + **Request Body**:

      ```json
      {
        "id": 1,
        "size": 256,
        "data": "8   888   , 888    Y888 888 888    ,ee 888 888 888 888 ...",
      }
      ```

  + **Responses**:
    | Code          |              Description           |
    |---------------|------------------------------------|
    | 201 Created   | Chunk successfully uploaded         |
    | 409 Conflict  | Chunk already exists               |
    | 404 Not Found | Image not found                    |

+ **Downloading an image**:
  + **method**: `GET`
  + **URI**: `/image/<sha256>`
  + **Accept**: `text/plain`
  + **Responses**:
    | Code          |              Description           |
    |---------------|------------------------------------|
    | 200 OK        | Image successfully downloaded       |
    | 404 Not Found | Image not found                    |

  + **Note**: This endpoint returns plain text, not JSON. It should return the whole image instead of separate chunks.

+ **Errors**:
  + **Accept**: `application/json`
  + **Response body**:

    ```json
    {
      "code": "400",
      "message": "Chunk ID field is missing."
    }
    ```

## The client executable

 You can use this executable to test your API. We will use the same to test your solution. You can configure the executable's behavior using the following command-line options:

```bash
  -chunksize int
        size of chunks used (default 256)
  -files int
        Amount of files to generate and send to the host (default 100)
  -host string
        host to send the requests to (default "localhost")
  -port int
        port to use when sending requests to the host (default 4444)
  -seed int
        set the seed used to produce randomness; providing a value will allow reproducible runs (default -1)
```

## Your solution

Your solution should run using Docker. We will use it to build and run your server using the provided `Dockerfile`. It should expose port `4444`. To test your solution we will run the following commands:

```bash
docker build -t recruitment/<candidate> .
docker run -d -p 4444:4444 recruitment/<candidate>
./ascii_<platform>_x64
```

If your solution uses `docker-compose`, we will execute the following commands:

```bash
docker-compose up -d --build
./ascii_<platform>_x64
```

To test your solution, first make sure you have your server
running and listening on `http://localhost:4444`. Ensure our executable has execution rights (`chmod +x`) and run it:

```bash
./ascii_<platform>_x64
```
```text
Success!
```