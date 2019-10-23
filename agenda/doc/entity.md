# entity功能函数介绍：
## userInfoOp.go
-   ReadUserInfoFromFile()
    ```bash
    /**
     * @arguments: nil
     * @return: []models.User
     */
    ```
    此函数用于从users.json文件中读取所有用户信息。
    通过利用文件读操作，包括os、bufio、json-iterator/go等库的使用，我们遍历整个用户信息文件，获取所有用户的models.Meeting切片然后返回。
-   WriteUserInfoToFile()
    ```bash
    /**
     * @arguments: []models.User
     * @return: nil
     */
    ```
    此函数用于将当前列表中的更新后的所有用户信息重新写入users.json文件中。
    通过利用文件写操作，包括os、bufio、json-iterator/go等库的使用，我们可以将所有用户信息编码为json格式的字符串并存储到users.json文件中。

-   SaveCurUserInfo()
    ```bash
    /**
     * @arguments: loginUser models.User
     * @return: nil
     */
    ```
    此函数用于将当前登陆的用户信息存储到curUser.txt文件中，方便登陆用户信息的存储。

-   ClearCurUserInfo()
    ```bash
    /**
     * @arguments: nil
     * @return: nil
     */
    ```
    当登陆用户登出的时候，我们利用os库Truncate函数来将登录用户信息从curUser.txt文件中删除。

-   IsLoggedIn()
    ```bash
    /**
     * @arguments: nil
     * @return: bool, models.User
     */
    ```
    此函数判断当前是否已经已经有用户登录，并且返回登录用户信息。
    我们可以利用此函数来加一些限定，因为未登录的用户不能进行cm、mtcancel等操作。

-   IsUser()
    ```bash
    /**
     * @arguments: name string
     * @return: bool
     */
    ```
    此函数用于判端当前用户名是否为已注册的用户，调用ReadUserInfoFromFile并加以判断即可。可以用于在创建、删除会议时判断用户是否存在；或者注册用户时判断该用户名是否已经被注册等。

-   RemoveUser()
    ```bash
    /**
     * @arguments: name string
     * @return: nil
     */
    ```
    此函数用于移除用处，主要是方便ru操作。调用ReadUserInfoFromFile获取用户信息，加以处理后再调用WriteUserInfoToFile更新用户信息即可。