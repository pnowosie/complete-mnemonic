# Complete a BIP-39 Mnemonic with a correct checksum word

This little fun function helps you to create valid BIP-39 menemonic phrases that contains only a few words repeated. It might be useful if easy to remember wallet is needed and security is not a concern.

:warning: **Never use such generated wallets to store real money!**

## How to run

- Clone this repo into your file system
- In [DigitOcean web console](https://cloud.digitalocean.com/functions) create a function namespace `lambda`
- Make sure you have `doctl` installed and configured
- **Deploy:** Run command to deploy the function
    ```bash
    make deploy
    ```
- **Run:**
    ```bash
    make run WORD=alien_alert | jq -r '.body.mnemonic'
    # echo: alien alert alien alert alien alert alien alert alien alert alien alley
    ```
- **Test:** If you have Go installed you can run tests with command
    ```bash
    make test
    ```

## Phrase length to entropy table

| words length | entropy bits | checksum bits | entropy bits of last word |
|-------------:|-------------:|--------------:|--------------------------:|
|           12 |          128 |             4 |                         7 | 
|           15 |          160 |             5 |                         6 |
|           18 |          192 |             6 |                         5 |
|           21 |          224 |             7 |                         4 |
|           24 |          256 |             8 |                         3 |

In general each word encodes 11-bits of information, where the last word contains 4-8 bits of checksum.
When you choose 3-words phrase, you actually have 33-bits of entropy, no more.

## Single word samples

In the samples folder you can find single word phrases with a corresponding checksum word.
E.g. `abandon about` line from [single-12](samples/single12.txt) file means 12-words phrase: `11x abandon + about`.

To easy generate full phrase from the sample, use the following bash command:

```bash
make single WORD=abandon LEN=12
```
You can test its correctness on Ian Coleman's [BIP-39 tool](https://iancoleman.io/bip39/).

If you played with the following command multiple times, you might spotted a lucky words that are checksum for itself.

<details>
  <summary><b>Lucky words from single-12 file</b></summary>

<pre>
action agent aim all ankle announce audit awesome \
beef believe blue border brand breeze bus business \
cannon canyon carry cave century cereal chronic coast convince cute \
dawn dilemma divorce dry \
elevator else embrace enroll escape evolve exclude excuse exercise expire \
fetch fever forward fury \
garment gauge gym \
half harsh hole hybrid \
illegal include index into invest involve \
jeans \
kick kite \
later layer legend life lyrics \
margin melody mom more morning \
nation neck neglect never noble novel \
obvious ocean oil orphan oxygen \
pause peasant permit piano proof pumpkin \
question \
real report rough rude \
salad scale screen sea seat sell seminar seven sheriff siege silver soldier spell split spray stadium sugar sunny sure \
tobacco tongue track tree trouble twelve twice type \
uniform useless \
valid very vibrant virtual vocal \
warrior word world \
yellow
</pre>
</details>

<details>
  <summary><b>Lucky words from single-18 file</b></summary>

<pre>
ahead desert dove dumb egg episode express fiction glad glass gorilla \
kiss leader misery mobile mother quiz rally response school sense spend stock \
upper usage wonder
</pre>
</details>

<details>
  <summary><b>Lucky words from single-24 file</b></summary>

<pre>
bacon flag gas great slice solution summer they trade trap zebra
</pre>
</details>


Knowing a little bash magic can be very helpful to generate such 👆 lists
```bash
egrep "(\b[a-zA-Z]+) \1\b" samples/single12.txt | cut -f1 -d' ' | xargs
```


## Acknowledgements

This project would not be possible without the following work of others.

- [Bitcoin's BIP-39]() - the original BIP-39 specification
- [tyler-smith/go-bip39](https://github.com/tyler-smith/go-bip39/blob/master/bip39_test.go) - the golang implementation of BIP-39 by @tyler-smith
- [Serverless](https://www.digitalocean.com/products/functions) - DigitalOcean awesome services especially serverless functions 

This is OSS project sample of my vast collection of little robots which makes my life easier. Many of them I host for free in the cloud (kudos to Digital Ocean).