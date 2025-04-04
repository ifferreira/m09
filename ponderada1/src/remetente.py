'''
referencia
https://gist.github.com/joeladdison/5244877
'''

import sys

def hamming_encode_nibble(data):
    def extract_bit(byte, pos):
        return (byte >> pos) & 0x01

    d = [0, 0, 0, 0]
    for i in range(4):
        d[i] = extract_bit(data, i)

    h = [0, 0, 0]
    h[0] = (d[1] + d[2] + d[3]) % 2
    h[1] = (d[0] + d[2] + d[3]) % 2
    h[2] = (d[0] + d[1] + d[3]) % 2

    p = 0 ^ d[0] ^ d[1] ^ d[2] ^ d[3] ^ h[0] ^ h[1] ^ h[2]

    encoded = (data & 0x0f)
    encoded |= (p << 7) | (h[2] << 6) | (h[1] << 5) | (h[0] << 4)
    
    return encoded

payload = sys.argv[1]

cabeçalho = "1111"
terminador = "0000"

encoded_payload = ""
for i in range(0, len(payload), 4):
    nibble = int(payload[i:i+4], 2)
    encoded_byte = hamming_encode_nibble(nibble)
    encoded_payload += format(encoded_byte, '07b')

frame = cabeçalho + encoded_payload + terminador

sys.stdout.write(frame + "\n")
