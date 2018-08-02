package fourth



//  AES operates on blocks of data, and a full block is needed to encrypt or decrypt.
// CBC requires messages be padded to the appropriate length.
// CBC also does not use nonces in the same way.

// In CBC mode, each block of ciphertext is XOR’d with the previous block.  This leads to the question of what the first block is XOR’d with.
// In CBC, we use a sort of dummy block called an initialisation vector. It may be randomly generated, which is often the right choice.

// The standard padding scheme used is the PKCS #7 padding scheme.
// ????
// We pad the remaining bytes with a byte containing the number of bytes of padding:
// if we have to add three bytes of padding, we’ll append 0x03 0x03 0x03 to the end of our plaintext.



