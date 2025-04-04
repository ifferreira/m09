'''
referencia
https://gist.github.com/joeladdison/5244877
'''

import sys

def extract_bit(byte, pos):
    return (byte >> pos) & 0x01

def hamming_decode_byte(byte):
    error = 0
    corrected = 0

    s = [0, 0, 0]
    s[0] = (extract_bit(byte, 1) + extract_bit(byte, 2) + extract_bit(byte, 3) + extract_bit(byte, 4)) % 2
    s[1] = (extract_bit(byte, 0) + extract_bit(byte, 2) + extract_bit(byte, 3) + extract_bit(byte, 5)) % 2
    s[2] = (extract_bit(byte, 0) + extract_bit(byte, 1) + extract_bit(byte, 3) + extract_bit(byte, 6)) % 2

    syndrome = (s[0] << 2) | (s[1] << 1) | s[2]

    if syndrome:
        byte ^= (1 << syndrome)
        corrected = 1
        error = 1

    p = 0
    for i in range(7):
        p ^= extract_bit(byte, i)

    if p != extract_bit(byte, 7):
        corrected += 1

    return ((byte & 0x0f), error, corrected)

cabeçalho = "1111"
terminador = "0000"

frame = sys.stdin.read().strip()

if frame.startswith(cabeçalho) and frame.endswith(terminador):
    payload_completo = frame[len(cabeçalho):-len(terminador)]
else:
    print("Erro de sincronização", file=sys.stderr)
    sys.exit(1)

decoded_payload = ""
for i in range(0, len(payload_completo), 7):
    byte = int(payload_completo[i:i+7], 2)
    decoded_byte, error, corrected = hamming_decode_byte(byte)
    decoded_payload += format(decoded_byte, '04b')

sys.stdout.write(decoded_payload + "\n")
