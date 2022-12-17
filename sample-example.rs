extern crate ring;

use ring::{
    rand::{SecureRandom, SystemRandom},
    signature::{Ed25519KeyPair, KeyPair, SigningKey, VerificationKey},
};

fn main() {
    // Generate a new pair of keys
    let rng = SystemRandom::new();

    // Generate a new Ed25519 key pair using the system's secure random number generator
    let pkcs8_bytes = Ed25519KeyPair::generate_pkcs8(&rng).unwrap();

    // Create an Ed25519 key pair from the PKCS#8 encoding of the private key
    let key_pair = Ed25519KeyPair::from_pkcs8(pkcs8_bytes.as_ref()).unwrap();

    // Extract the public and private keys from the key pair
    let public_key = key_pair.public_key();
    let private_key = key_pair.private_key();

    // Alice wants to send a message to Bob
    let message = b"Hello, Bob!";

    // Alice signs the message using her private key
    let signature = key_pair.sign(message);

    // Bob receives the message and signature and verifies them using Alice's public key
    assert!(public_key.verify(message, &signature).is_ok());
}
