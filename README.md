# Setup
```
createdb datagol
creatuser datagol
psql -c "grant all on database datagol to datagol"
```

# transactor.Connection

Writes updates to the underlying db
Only one should run

## Todo

* Ensure only one Connection running
* Broadcast updates
* Actual http api (underway)
* Configure db, creds, port and interface bindings from command line
* move to transact subcommand

## Could do

* authorization
* user activity logging