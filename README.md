
Lock and locker are working. Unlock.go and Unlocker.go should be nearly identical except unlocker.go should have the following steps
1. Validate subject (same as locker.go)

2. verify signature in keyfile.sig

3. read aeskey (ciphertext) from keyfile and store

4. delete keyfile and keyfile.sig

5. decrypt files and replace.

Can you please create a methods for 2 and 3? I can do the rest.



# Directory Encryption/Decryption Tool


## Prerequisite

+ golang version > 1.11.1



## `keygen` generation

```bash
make build
```

## `keygen` test

```bash
make test
```




## `keygen` utilization


### Elliptic Curve Public Key & Private Key generation

```bash
./keygen -t ec -s ATypicalSubjectName -pub testData/wECpub -priv testData/wECpriv
```

```bash
cat wECpriv
# ECC P256;15743654092892585608819065759089171092014055741737470039870784446426342155267%
```

```bash
cat wECpub
{
  "Subject":"ATypicalSubjectName",
  "PublicKeyAlgorithm":"ECC P256",
  "PublicKeyData":"18832934341109286343168220942514825886344142229070983551989341085355620201760;64764780779256393029477997977138577744871365114987836919565023199167732409469",
}
```

### RSA Public Key & Private Key generation


```bash
./keygen -t rsa -s ANonTypicalSubjectName -pub testData/kRSApub -priv testData/kRSApriv
```

```bash
cat kRSApriv
# RSA Encryption;10449807426307551027219093463542573263413149262641621031064508318073079725329868558685835415211835861392394702712573972376784159987008050362791731372585449593034975205583083100844375761537883605789685467787845922474102329804158314030744907401462374349342697833270408016887529963156796687308167421483995246349456316743065414976232901535332831135345433418015820925318344822539041166343194943212752810827233594917092991577309015871621848272279073362978666602814097464063416250500460266648472619989869498746695164238392383592936214475306572366737903436597271599345065526335823247762988763544559631810395632206018411057401
```

```bash
cat kRSApub
{
    "Subject":"ANonTypicalSubjectName",
    "PublicKeyAlgorithm":"RSA Encryption",
    "PublicKeyData":"256;29324699379032198838351362863757370213424148463892434594196911948340987666307424669461060015703437819991152335003595119793495910467951811108430320243432928405358104395319796059777248192339996483370669542794042057942332978863369162740127130100609729713662429900618426402447463012563457416122093358816331097885258455256829012597300621530135531996525075052327607474415584560716526979814182339973796650585274161231859299374468125353142391853850583371450024153197957982485885356944214920014355925617080270937010725560635376296510630599267091705378343646359694771599146069407885012398264814702029663196692413455687777481141;65537"
}
```













