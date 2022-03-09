# Lab 4 LAW: Implementasi CRUD
## Identitas
* Nama: Evando Wihalim
* NPM : 1806205445

## Endpoint
1. GET    /car/ \
    Example: \
    ![Get all cars](https://github.com/Evando2000/simplecarapi/blob/main/docs/1_get_all.PNG?raw=true)

2. POST   /car/
    * Content-type:
        * multipart/form-data
    * Request Body:
        * model (string)
        * color (string)
        * brand (string)

    Example: \
    ![Create a car](https://github.com/Evando2000/simplecarapi/blob/main/docs/2_create_car.PNG?raw=true)

3. GET    /car/[ID] \
    Example: \
    ![Get a car](https://github.com/Evando2000/simplecarapi/blob/main/docs/3_get_car_by_id.PNG?raw=true)

4. PUT    /car/[ID]
    * Content-type:
        * multipart/form-data
    * Request Body:
        * model (string)
        * color (string)
        * brand (string)

    Example: \
    ![Update a car](https://github.com/Evando2000/simplecarapi/blob/main/docs/4_update_car.PNG?raw=true)

5. DELETE /car/[ID] \
    Example: \
    ![Delete a car](https://github.com/Evando2000/simplecarapi/blob/main/docs/5_delete_car.PNG?raw=true)

6. POST   /file/ \
    Example: \
    ![Upload a file](https://github.com/Evando2000/simplecarapi/blob/main/docs/6_upload_file.PNG?raw=true)
    Uploaded file: \
    ![Uploaded file](https://github.com/Evando2000/simplecarapi/blob/main/docs/7_uploaded_file.PNG?raw=true)

## How to run (Windows)
1. Download file 'lab4.exe' and run it. Done.

## How to make binary executable file for linux from windows
1. Clone this repository
2. Run the 'makebin.bat' with desired output file name. Example:
    ```cmd
    makebin.bat lab4
    ```
3. On your linux terminal, execute the bin file. Example:
    ```cmd
    ./lab4
    ```
