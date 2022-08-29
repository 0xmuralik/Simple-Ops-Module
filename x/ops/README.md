# Build chain

```bash
# it will create binary in build folder with `ethermintd`
$ make build
```

# Setup Chain

```bash
./build/chibaclonkd keys add root
./build/chibaclonkd init test-moniker --chain-id ethermint_9000-1
./build/chibaclonkd add-genesis-account $(./build/chibaclonkd keys show root -a) 1000000000000000000aphoton,1000000000000000000stake
./build/chibaclonkd gentx root 1000000000000000000stake --chain-id ethermint_9000-1
./build/chibaclonkd collect-gentxs
./build/chibaclonkd start
```

# Create Name record

```
$ ./build/chibaclonkd tx ops create Murali 22 --from root --chain-id $(./build/chibaclonkd status | jq .NodeInfo.network -r) --keyring-backend test

auth_info:
  fee:
    amount: []
    gas_limit: "200000"
    granter: ""
    payer: ""
  signer_infos: []
  tip: null
body:
  extension_options: []
  memo: ""
  messages:
  - '@type': /cosmos.ops.v1beta1.MsgCreateNameRecord
    age: "22"
    name: Murali
    signer: ethm1j5n2rgkjg3xrch0t0q25gh6vwjmnetj9tddel6
  non_critical_extension_options: []
  timeout_height: "0"
signatures: []
confirm transaction before signing and broadcasting [y/N]: y
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: '[]'
timestamp: ""
tx: null
txhash: F435A29AEF325EE250B978BB0564430A6C160E76F3D23DE04C0558CCD9252282
```

# Update name record

```
 ./build/chibaclonkd tx ops update Record-0 Sai 26 --from root --chain-id $(./build/chibaclonkd status | jq .NodeInfo.network -r) --keyring-backend test -y
auth_info:
  fee:
    amount: []
    gas_limit: "200000"
    granter: ""
    payer: ""
  signer_infos: []
  tip: null
body:
  extension_options: []
  memo: ""
  messages:
  - '@type': /cosmos.ops.v1beta1.MsgUpdateNameRecord
    age: "26"
    id: Record-0
    name: Sai
    signer: ethm1q400pvzczfc5hdxkke3t8pa6gvalt9ah653c0q
  non_critical_extension_options: []
  timeout_height: "0"
signatures: []
confirm transaction before signing and broadcasting [y/N]: y
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: '[]'
timestamp: ""
tx: null
txhash: 0574E795E4C6624087C2CDD90013497B6C6A099413C553350A02456E51351F3C
```

# Delete name record

```
./build/chibaclonkd tx ops delete Record-0 --from root --chain-id $(./build/chibaclonkd status | jq .NodeInfo.network -r) --keyring-backend test
auth_info:
  fee:
    amount: []
    gas_limit: "200000"
    granter: ""
    payer: ""
  signer_infos: []
  tip: null
body:
  extension_options: []
  memo: ""
  messages:
  - '@type': /cosmos.ops.v1beta1.MsgDeleteNameRecord
    id: Record-0
    signer: ethm1q400pvzczfc5hdxkke3t8pa6gvalt9ah653c0q
  non_critical_extension_options: []
  timeout_height: "0"
signatures: []
confirm transaction before signing and broadcasting [y/N]: y
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: '[]'
timestamp: ""
tx: null
txhash: 102067B95CA068F2CDACC3932DD0E9C6C1C4608BBA3B746D8F6EC742E281372C
```

# List Records

```
 $ ./build/chibaclonkd q ops list
pagination: null
records:
- age: "26"
  id: Record-0
  name: Sai
```

# Get Record by Id

```
$  ./build/chibaclonkd q ops record Record-0
record:
  age: "26"
  id: Record-0
  name: Sai
```
