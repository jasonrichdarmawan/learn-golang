If you can't run compiled Go code in `scratch` image. Try to:
1. do `$ ldd /main` to list the dynamic dependencies, assuming `/main` is the compiled Go code. Then, copy the libraries to the `scratch` image.
2. or you can use the `CGO_ENABLED=0` flag. [See SO's question](https://stackoverflow.com/questions/52640304/standard-init-linux-go190-exec-user-process-caused-no-such-file-or-directory)

Note:
1. Not sure why, but running compiled Go code in `alpine:3.16` and `scratch` image will have different behavior:
   1. In `scratch` image, if a user send a request larger than the `maxMemory`, the `ParseMultipartForm` will return `non-nil error`.
   2. In `alpine:3.16` image, `ParseMultipartForm` will not return `non-nil error`. This is because `ParseMultipartForm` will store the remainder in temporary filels.
   
   The solution is to create `/tmp` folder. Therefore, make sure to create unit tests. The conclusion is different environment, different behavior.