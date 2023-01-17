# ReSTful API (Respresentational State Transfer)

CRUD {Create, Read, Update, Delete}

HTTP Methods => {GET, POST, DELETE, PUT, OPTIONS, ...}

## Employee Management Server (JSON)

```
CRUD            Action                HTTP Method             URI                   Req Body                    Res Body
----------------------------------------------------------------------------------------------------------------------------
Read            Index                   GET                 /employees                -                         [{...}, ...]
Read            Show                    GET                 /employees/{id}           -                           {...}
Create          Create                  POST                /employees               {...}                        {id: , ...}
Update          Update                  PUT                 /employees/{id}          {...}                        {...}
Update          Update                 PATCH                /employees/{id}          {some attrs}                 {...}
Destroy         Destroy                DELETE               /employees/{id}           -                           {...} / -
```
