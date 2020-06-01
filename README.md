**String encode / decode**

Translate string based on Pig Latin rules.

Encode vowels to integers according to rules (and vice versa):
```
a -> 1
e -> 2
i -> 3
o -> 4
u -> 5
```


**Building**

Build with `./scripts/build.sh`


**Running**

Run ./bin/stringCryptor with passing one of the flags: `"toPigLatin"`, `"encodeVowelsToIntegers"`, `"decodeIntegersToVowels"` and text

E.g.
```
./bin/stringCryptor toPigLatin The weather is fine
```
The output should be:
```
eThay eatherway isay inefay
```
