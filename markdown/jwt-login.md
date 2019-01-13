###Login Flow
---

#### Login
```mermaid
graph LR
    A[Request] --> B[Server]
    B --> |Get User Data | C{DB}

    C --> |Has Data| F{Create New Token}
    C --> |No Data| Return


    F --> |Success| G[Set & Refresh Cookies]
    F --> |Failed| Return

    G --> Return
```


---

#### Create Token
```mermaid
graph LR
    A[Request] --> B[Server]

    B --> |Generate| C{CSRF Secret}

    C --> |Success| D[Create Refresh Token]
    C --> |Failed| Return

    D --> F[Create Auth Token] 
    F --> Return

    
```
<br/>
1. CSRF using Random number


---

#### Set & Refresh Cookies
```mermaid
graph LR
    A[Request] --> B[Set Auth Cookie]
    B --> C[Set Refresh Cookie]
    C --> Return
```

---

#### Create Refresh Token
````mermaid
graph LR
    a[request] --> |generte| b[expity time]
    b --> |store| c{refresh token}
    c --> |success| d[generate token claim]
    c --> |failed| return

    d --> f[sign token]
    f --> return
````


---

#### Create Auth Token
````mermaid
graph LR
    a[request] --> |generte| b[expity time]
    b --> |generte| d[token claim]

    d --> f[sign token]
    f --> g[return]
````

---
#### Store Refresh Token
````mermaid
graph LR
    a[request] --> |generte| b{random string}
    b --> |success| c{Is Valid}
    b --> |failed| return

    c --> |valid| d[store]
    c --> |invalid| b

    d --> return

````