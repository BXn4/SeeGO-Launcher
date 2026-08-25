export function base64ToBlob(base64: string, mime = "image/png") {
  const binary = atob(base64);
  const bytes = new Uint8Array(binary.length);

  for (let i = 0; i < binary.length; i++) {
    bytes[i] = binary.charCodeAt(i);
  }

  const blob = new Blob([bytes], { type: mime });

  return URL.createObjectURL(blob);
}

export function sleep(ms: number) {
  return new Promise(resolve => setTimeout(resolve, ms));
}
