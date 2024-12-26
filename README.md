# Convertir imágenes

Convertir imágenes a diferentes formatos a webp, jpg o png y subir a appwrite

## Requisitos

- Go 1.16
- Libwebp 
- Gcc

## Instalación de libwebp

### Windows

Instale msys2 e instale el gcc puede seguir esta guía [https://sajidifti.medium.com/how-to-install-gcc-and-gdb-on-windows-using-msys2-tutorial-0fceb7e66454]tutorial-gcc. Actualice su msys2

```bash
pacman -Syu
pacman -Su
  
```
Luego instale segun 32 o 64 bits
    
```bash
// Install GCC for C and C++
// 64bits
pacman -S mingw-w64-x86_64-gcc 

// 32bits
pacman -S mingw-w64-i686-gcc

// Install GDB for C and C++ (Optional)
// 64bits
pacman -S mingw-w64-x86_64-gdb

// 32bits
pacman -S mingw-w64-i686-gdb
```

En las variables de entorno coloque

```bash
C:\msys64\ucrt64\bin
C:\msys64\mingw64\bin
```

Abra una terminal y instale el libwebp con pacman

```bash
pacman -S mingw-w64-x86_64-libwebp
pacman -S mingw-w64-ucrt-x86_64-libwebp
pacman -S mingw-w64-i686-libwebp
pacman -S mingw-w64-clang-x86_64-libwebp
```

### Ubuntu
```bash
sudo apt-get update
sudo apt install gcc
sudo apt-get install libwebp-dev
```

## Instalación

```bash
go get github.com/nelsonp17/gosquoosh
```


## Ejecución

Renombrar el archivo `main.go.example` a `main.go` y configurar según sus necesidades
```bash
go run github.com/nelsonp17/gosquoosh
```


### POST /convert

Todas las conversiones son por método POST

# Los formatos aceptados son

| **Formato de entrada** | **Formato de salida** | Tipo |
| --- | --- | --- |
| image/png | image/webp | Conversión |
| image/jpg | image/webp | Conversión |
| image/webp | image/jpg | Conversión |
| image/webp | image/png | Conversión |
| image/png | image/png | Compresión |

# Los parámetros al subir una imagen:

| Campo | Tipo de dato | Descripción |
| --- | --- | --- |
| image | File | El archivo a subir, debe usar form data multipart |
| quality | Integer | la calidad de compresión  |
| from | String | Formato de la imagen que subio, sí es jpg, png o webp. Consulte los formatos aceptados |
| to | String | Formato de salida de la imagen, sí es jpg, png o webp. Consulte los formatos aceptados |

# Uso básico

El proceso de conversión de la imagen guarda la imagen en `frontend/public/convert` mientras que comprimir las guarda en `frontend/public/compress` para limpiar las imágenes el contenido de estas carpetas llame por método `post` a `/api/v1/image/cache/clear` Cuando se sube directamente la imagen a appwrite o se descarga por los endpoints se eliminar la imagen usada

- Endpoint `[POST] /api/v1/image/convert/`
    
    Consulte los parámetros de la solicitud para enviar un archivo
    
    ### Request
    
    ```json
    {
      "image": "path/bestmenu.png",
      "quality": 75,
      "from": "image/png",
      "to": "image/webp"
    }
    ```
    
    ### Response
    
    ```json
    {
        "data": {
            "original": {
                "url": "http://127.0.0.1:3000/api/v1/image/visor/input/bestmenu.png",
                "download": "http://127.0.0.1:3000/api/v1/image/download/output/bestmenu.png",
                "size": 24181,
                "size_kb": 23.6142578125,
                "width": 0,
                "height": 0,
                "mimetype": "image/png",
                "filename": "bestmenu.png"
            },
            "converted": {
    		        "url": "http://127.0.0.1:3000/api/v1/image/visor/input/bestmenu.webp",
                "download": "http://127.0.0.1:3000/api/v1/image/download/output/bestmenu.webp",
                "size": 10124,
                "size_kb": 9.88671875,
                "width": 0,
                "height": 0,
                "mimetype": "image/webp",
                "filename": "bestmenu.webp"
            }
        }
    }
    ```
    
- Endpoint `[POST] /api/v1/image/compress`
    
    Los parámetros de la solicitud son:
    
    | Campo | Tipo de dato | Descripción |
    | --- | --- | --- |
    | image | File | El archivo a subir, debe usar form data multipart |
    | quality | Integer | la calidad de compresión  |
    | from | String | Formato de la imagen que subió, para compresión por ahora solo es png el único aceptado Consulte los formatos aceptados |
    | to | String | Formato de salida de la imagen, para compresión por ahora solo es png el único aceptado Consulte los formatos aceptados |
    
    ### Request
    
    ```json
    {
      "image": "path/bestmenu.png",
      "quality": 75,
      "from": "image/png",
      "to": "image/png"
    }
    ```
    
    ### Response
    
    ```json
    {
        "data": {
            "original": {
                "url": "http://127.0.0.1:3000/api/v1/image/visor/input/bestmenu.png",
                "download": "http://127.0.0.1:3000/api/v1/image/download/output/bestmenu.png",
                "size": 24181,
                "size_kb": 23.6142578125,
                "width": 0,
                "height": 0,
                "mimetype": "image/png",
                "filename": "bestmenu.png"
            },
            "compressed": {
                "url": "http://127.0.0.1:3000/api/v1/image/visor/input/bestmenu.png",
                "download": "http://127.0.0.1:3000/api/v1/image/download/output/bestmenu.png",
                "size": 23485,
                "size_kb": 22.9345703125,
                "width": 0,
                "height": 0,
                "mimetype": "image/png",
                "filename": "bestmenu.png"
            }
        }
    }
    ```
    
- Endpoint `[POST] /api/v1/image/upload/appwrite`
    
    Sube la imagen a appwrite. Por defecto sube el archivo a la instalación de appwrite que tenemos pero si se necesita subir a otra instancia proporcione los datos completos con el api_secret, endpoint, project_id y bucket_id. Recuerde que el filename es el nombre que le dio el Response del paso 1
    
    ### Request
    
    ```json
    {
      "project_id": "PROJECT_ID_APPWRITE", // REQUIRED
      "filename": "FILE_NAME", // REQUIRED, NOMBRE PROPORCIONADO POR EL PASO 1, EJ: bestmenu.webp
      "bucket": "BUCKET_ID", // REQUIRED
      "dir": "convert", // REQUIRED, convert o compress segun el caso
      "endpoint": "ENDPOINT_APPWRITE", // OPCIONAL
      "api_secret": "API_SECRET", // OPCIONAL
      "id": "FILE_ID" // OPCIONAL
    }
    ```
    
    ### Response
    
    ```json
    {
        "data": {
            "file": {
                "$id": "62a194c658d2e25bf7ce",
                "bucketId": "674b64ea001949ba11f2",
                "$createdAt": "2024-12-25T14:46:46.892+00:00",
                "$updatedAt": "2024-12-25T14:46:46.892+00:00",
                "$permissions": [
                    "read(\"any\")"
                ],
                "name": "bestmenu.webp",
                "signature": "1d0c8b80444acb309f15e9bf28dd8861",
                "mimeType": "image/webp",
                "sizeOriginal": 10124,
                "chunksTotal": 1,
                "chunksUploaded": 1
            },
            "preview": "https://cloud.appwrite.io/v1/storage/buckets/674b64ea001949ba11f2/files/62a194c658d2e25bf7ce/view?project=67487f540033dd62d3af"
        }
    }
    ```
    
- Endpoint `[POST] /api/v1/image/upload/appwrite/convert`
    
    Sube la imagen a appwrite y convierte directamente
    
    ### Request
    
    ```json
    {
      "project_id": "PROJECT_ID_APPWRITE", // REQUIRED
      "filename": "FILE_NAME", // REQUIRED, NOMBRE PROPORCIONADO POR EL PASO 1, EJ: bestmenu.webp
      "bucket": "BUCKET_ID", // REQUIRED
      "dir": "convert", // REQUIRED, convert o compress segun el caso
      "endpoint": "ENDPOINT_APPWRITE", // OPCIONAL
      "api_secret": "API_SECRET", // OPCIONAL
      "id": "FILE_ID", // OPCIONAL
      "image": "path/bestmenu.png",
      "quality": 75,
      "from": "image/png",
      "to": "image/webp"
    }
    ```
    
    ### Response
    
    ```json
    {
        "data": {
            "appwrite": {
                "file": {
                    "$id": "62a197dca2d856a911a0",
                    "bucketId": "674b64ea001949ba11f2",
                    "$createdAt": "2024-12-25T15:00:36.636+00:00",
                    "$updatedAt": "2024-12-25T15:00:36.636+00:00",
                    "$permissions": [
                        "read(\"any\")"
                    ],
                    "name": "testing-1200x900.webp",
                    "signature": "8ae86b287654835ea724a06c3495011c",
                    "mimeType": "image/webp",
                    "sizeOriginal": 32118,
                    "chunksTotal": 1,
                    "chunksUploaded": 1
                },
                "preview": "https://cloud.appwrite.io/v1/storage/buckets/674b64ea001949ba11f2/files/62a197dca2d856a911a0/view?project=67487f540033dd62d3af"
            },
            "converted": {
                "url": "https://cloud.appwrite.io/v1/storage/buckets/674b64ea001949ba11f2/files/62a197dca2d856a911a0/view?project=67487f540033dd62d3af",
                "size": 32118,
                "size_kb": 31.365234375,
                "width": 0,
                "height": 0,
                "mimetype": "image/webp",
                "filename": "testing-1200x900.webp"
            },
            "original": {
                "url": "",
                "size": 96116,
                "size_kb": 93.86328125,
                "width": 0,
                "height": 0,
                "mimetype": "image/jpeg",
                "filename": "testing-1200x900.jpg"
            }
        }
    }
    ```
    

> Sino va a subir la imagen a appwrite y simplemente necesita descargar asegúrese de limpiar el directorio después de descargar el recurso llamando a `[POST] /api/v1/image/cache/clear`
> 

### Error Responses

- 400 Bad Request - Invalid input parameters
- 415 Unsupported Media Type - Unsupported image format
- 500 Internal Server Error - Server processing error

## Integration with Appwrite

The service integrates with Appwrite for:

- Authentication using Appwrite API keys
- File storage in Appwrite Storage
- Event processing through Appwrite Functions