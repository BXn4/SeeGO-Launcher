export function stripMarkup(raw: string): string {
  return raw
    .replace(/\[\[href\]\].*?\[\[href\]\]/gs, "")
    .replace(/\[\[.*?\]\]/g, "")
    .replace(/#+ c/g, "")
    .replace(/#+/g, "")
    .replace(/\n+-/g, "")
    .replace(/\r?\n+/g, " ")
    .trim();
}
