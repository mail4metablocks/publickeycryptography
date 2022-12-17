import * as crypto from 'crypto';

// Generate a new pair of keys
const { publicKey, privateKey } = crypto.generateKeyPairSync('rsa', {
  modulusLength: 4096,
  publicKeyEncoding: {
    type: 'spki',
    format: 'pem'
  },
  privateKeyEncoding: {
    type: 'pkcs8',
    format: 'pem'
  }
});

// Alice wants to send a message to Bob
const message = 'Hello, Bob!';

// Alice encrypts the message using Bob's public key
const encryptedMessage = crypto.publicEncrypt(publicKey, Buffer.from(message));

// Bob receives the encrypted message and decrypts it using his private key
const decryptedMessage = crypto.privateDecrypt(privateKey, encryptedMessage);

console.log(decryptedMessage.toString()); // "Hello, Bob!"
