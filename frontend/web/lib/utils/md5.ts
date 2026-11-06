/**
 * 计算字符串的 MD5 值（UTF-8 编码，返回 32 位小写十六进制），
 * 与后端 Go 的 fmt.Sprintf("%x", md5.Sum(...)) 结果一致。
 */
export function md5(input: string): string {
  const S = [
    7, 12, 17, 22, 7, 12, 17, 22, 7, 12, 17, 22, 7, 12, 17, 22, 5, 9, 14, 20, 5,
    9, 14, 20, 5, 9, 14, 20, 5, 9, 14, 20, 4, 11, 16, 23, 4, 11, 16, 23, 4, 11,
    16, 23, 4, 11, 16, 23, 6, 10, 15, 21, 6, 10, 15, 21, 6, 10, 15, 21, 6, 10,
    15, 21,
  ];

  // K[i] = floor(abs(sin(i + 1)) * 2^32)
  const K = new Uint32Array(64);
  for (let i = 0; i < 64; i++) {
    K[i] = Math.floor(Math.abs(Math.sin(i + 1)) * 0x100000000);
  }

  const data = new TextEncoder().encode(input);
  const byteLen = data.length;

  // 补位：追加 0x80，补 0 至 (len % 64 == 56)，最后 8 字节写入原始比特长度（小端）
  const paddedLen = (((byteLen + 8) >>> 6) << 6) + 64;
  const buffer = new Uint8Array(paddedLen);
  buffer.set(data);
  buffer[byteLen] = 0x80;

  const view = new DataView(buffer.buffer);
  view.setUint32(paddedLen - 8, (byteLen << 3) >>> 0, true);
  view.setUint32(paddedLen - 4, byteLen >>> 29, true);

  let a = 0x67452301;
  let b = 0xefcdab89;
  let c = 0x98badcfe;
  let d = 0x10325476;

  const M = new Uint32Array(16);
  for (let offset = 0; offset < paddedLen; offset += 64) {
    for (let j = 0; j < 16; j++) {
      M[j] = view.getUint32(offset + j * 4, true);
    }

    let A = a;
    let B = b;
    let C = c;
    let D = d;

    for (let i = 0; i < 64; i++) {
      let F: number;
      let g: number;
      if (i < 16) {
        F = (B & C) | (~B & D);
        g = i;
      } else if (i < 32) {
        F = (D & B) | (~D & C);
        g = (5 * i + 1) % 16;
      } else if (i < 48) {
        F = B ^ C ^ D;
        g = (3 * i + 5) % 16;
      } else {
        F = C ^ (B | ~D);
        g = (7 * i) % 16;
      }

      F = (F + A + K[i] + M[g]) >>> 0;
      A = D;
      D = C;
      C = B;
      const shift = S[i];
      B = (B + ((F << shift) | (F >>> (32 - shift)))) >>> 0;
    }

    a = (a + A) >>> 0;
    b = (b + B) >>> 0;
    c = (c + C) >>> 0;
    d = (d + D) >>> 0;
  }

  const toHexLE = (n: number) => {
    let hex = "";
    for (let i = 0; i < 4; i++) {
      hex += ((n >>> (i * 8)) & 0xff).toString(16).padStart(2, "0");
    }
    return hex;
  };

  return toHexLE(a) + toHexLE(b) + toHexLE(c) + toHexLE(d);
}
