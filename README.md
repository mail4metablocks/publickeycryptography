# Public key cryptography


Public key cryptography is a type of cryptography that uses a pair of keys to encrypt and decrypt messages. One of the keys is called the public key, and the other is called the private key. The public key is intended to be shared with others, while the private key must be kept secret.

Here's how it works:

Alice wants to send a message to Bob. She uses Bob's public key to encrypt the message.
The encrypted message is sent to Bob.
Bob receives the message and uses his private key to decrypt it.
The key feature of public key cryptography is that the public key can be used to encrypt the message, but only the corresponding private key can be used to decrypt it. This means that anyone can send an encrypted message to Bob, but only Bob will be able to read it.

Public key cryptography is used in a variety of applications, including email, online banking, and secure communication over the internet. It is a very important part of modern information security.


# Sequence diagram

Process of sending an encrypted message using public key cryptography:

Alice generates a message and wants to send it to Bob.
Alice looks up Bob's public key and uses it to encrypt the message.
The encrypted message is sent to Bob.
Bob receives the encrypted message and uses his private key to decrypt it.
Bob reads the decrypted message.

  +----------+     Generate message     +------------+
  |  Alice   |------------------------>|   Bob      |
  +----------+                         +------------+
                                               |
  +----------+   Encrypt message      +------------+
  |  Alice   |------------------------>|   Bob      |
  +----------+      using Bob's       +------------+
                  public key (Kp)            |
                                               |
  +----------+                         +------------+
  |  Alice   |                         |   Bob      |
  +----------+                         +------------+
                                               |
                                     Receive encrypted message
                                               |
  +----------+                         +------------+
  |  Alice   |                         |   Bob      |
  +----------+                         +------------+
                                               |
                                    Decrypt message
                                               |
  +----------+                         +------------+
  |  Alice   |                         |   Bob      |
  +----------+      using Bob's       +------------+
                  private key (Ks)            |
                                               |
                                      Read decrypted message
                                               |
