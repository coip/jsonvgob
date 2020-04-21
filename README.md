take a look at the `diff` file for a ~patch of changes/overhead necessary for shifting from json encoding to gob encoding.
that overhead in vim looks like `:%s/json/gob/g`... so one command :) 


# results

### json: first 100B
``` bash
ethan@vanitas:~/xmljson$ ./jsonencoded | head -c 100 && echo ""
{"employees":[{"firstName":"John","lastName":"Doe the 1th"},{"firstName":"John","lastName":"Doe the
```

### gob: first 100B
``` bash
ethan@vanitas:~/xmljson$ ./gobencoded | head -c 100 && echo ""
    Employees E   []main.Employee   1Employee    FirstName L
```

seen here, JSON may be better for small/discrete payloads given some of the upfront cost of gobs, but the gob deduplication pays off at scale. amortization anybody?

### gob: first 220B
``` bash
ethan@vanitas:~/xmljson$ ./gobencoded | head -c 220 && echo ""
    Employees E   []main.Employee   1Employee    FirstName LastName   UuJohnDoe the 1th JohnDoe the 2th JohnDoe the 3th JohnDoe the 4th JohnDoe the 5th
```


## net bytes for 1000 Employees below, for each of `encoding/{json,gob}`:

``` bash
ethan@vanitas:~/xmljson$ ./gobencoded | wc -c
21992
ethan@vanitas:~/xmljson$ ./jsonencoded | wc -c
47860
```

-----
ps.

gophers speak bits/bytes... 0's and 1's. `encoding/gob` is a *binary* encoding.

sending gobs over the *wire(boundary really, local FD even) leverages the gopher language.

humans speak ~ASCII to an extent, and there is of-course value in human-readable -ness.

but **TANSTAAFL**, and human-readable character encoding is perhaps unnecessary for the gophers.

(id like to acknowledge regarding `encoding`/`transport`, i'm a fan of `protofbuf`/~`grpc`, and that would be a likely iteration candidate/better direction in lieu of gobs. but this is also a few minutes of thoughts to merely log the `json` v `gob` thing I think about from time to time.)
