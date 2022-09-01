# otten-Go-
Repo for Otten Coffee entry test

# Command
We can run it by using ```go run cmd/main.go``` or we can build the binary first <br />
To build binary run the command ```go build -o [name.exe] cmd/main.go``` if Windows or ```go build -o [name] cmd/main.go``` if macOS <br />
Then run the binary by using command ```./name.exe``` or ```./name``` <br />

# Endpoint
To get tracking status use the GET method and call ```baseUrl:portAddress/otten``` <br />
Example: GET ```http://localhost:8099/otten``` <br />


# Config
Config is read from config.yaml file <br />
The default are: <br /><br />

```server:``` <br />
 ```address: ":8099"``` <br />
```data:``` <br />
  ```url: "https://gist.githubusercontent.com/nubors/eecf5b8dc838d4e6cc9de9f7b5db236f/raw/d34e1823906d3ab36ccc2e687fcafedf3eacfac9/jne-awb.html"``` <br />
```message:``` <br />
  ```errorCode: "060102"``` <br />
  ```successCode: "060101"``` <br />
  ```errorText: "Failed to fetch delivery tracking detail"``` <br />
  ```successText: "Delivery tracking detail fetched successfully"``` <br />

  # Collection
  Postman API collection = ```https://www.getpostman.com/collections/1d2fcd25bdf7100a1507``` <br />