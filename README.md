## ChaCha20 encryption/decryption tool  

*Usage:* 
  $ chacha20 <keyfile> <noncefile> < infile > outfile  

*Arguments:*  
  keyfile    Path to the file  
  noncefile  Path to the file  

The program reads from stdin and writes to stdout.  
Use input/output redirection for files.  

*Examples:*  
  Encryption: chacha20 key.hex nonce.hex < plaintext.txt > encrypted.bin  
  Decryption: chacha20 key.hex nonce.hex < encrypted.bin > decrypted.txt  

*Note:*  
 The key must be 32 hex bytes (64 characters) long.  
 The nonce must be 12 hex bytes (24 characters) long.  

For nonce creation, I recommend my nora program.  

