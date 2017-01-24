# Setup
```
./script/resetdb
```

# transactor

Writes transactions to the db

Currently only 'assert' works.

## trying it.

In one shell run

```
./script/trampoline server
```

And in another
 
```
./script/test_transactor
```

## Todo

* Accumulate support in API
* Ensure only one Backend running
* Broadcast updates

## Could do

* authorization
