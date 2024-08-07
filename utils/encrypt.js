
// For encryption and descryption for client side JS code use this exmple
// Need crypto-js package
//    npm install crypto-js

const CryptoJS = require('crypto-js');

const keyHex = 'b210e6ae7920d010dc7ea36fda78be24f3fc81adb81db87ffa3324b1e4ea1538'; // Key must be 16 bytes
const key = CryptoJS.enc.Hex.parse(keyHex);

function encrypt(text) {
    const iv = CryptoJS.lib.WordArray.random(16);
    const encrypted = CryptoJS.AES.encrypt(text, key, { iv: iv });
    const ciphertext = iv.concat(encrypted.ciphertext);
    return CryptoJS.enc.Base64.stringify(ciphertext);
}

function decrypt(ciphertext) {
    const ciphertextBytes = CryptoJS.enc.Base64.parse(ciphertext);
    const iv = CryptoJS.lib.WordArray.create(ciphertextBytes.words.slice(0, 4));
    const encrypted = CryptoJS.lib.WordArray.create(ciphertextBytes.words.slice(4));
    const decrypted = CryptoJS.AES.decrypt({ ciphertext: encrypted }, key, { iv: iv });
    return decrypted.toString(CryptoJS.enc.Utf8);
}

const text = "sk_test_51PhT0rLDOmWPE4uMOe95LkIZQxE3e7j0u3cWhTWLzWwwdqKIAvWohc64lYaVZLVjbr1NFUkW53JkQAW275B2uTKc00WLK5pozF";
const encrypted = encrypt(text);
console.log("Encrypted: ", encrypted);

const decrypted = decrypt(encrypted);
console.log("Decrypted: ", decrypted);
