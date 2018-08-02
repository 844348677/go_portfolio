package third



// AES-CTR and AES-CBC with an HMAC.
// data is first encrypted with AES in the appropriate mode, then an HMAC is appended.
//  the nonce must be only ever used once with the same key. Reusing a nonce can be catastrophic

// AES-CTR
// 
