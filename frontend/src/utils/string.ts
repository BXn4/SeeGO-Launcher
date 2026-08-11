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


export function parseTerms(raw: string): [string, string, string] {
  let title = "";
  let modified = "";

  const lines = raw.split("\n");
  let parsed = "";
  let headerParsed = false;
  for (const line of lines) {
    const trimmed = line.trim();
    if (!headerParsed) {
      if (trimmed === "ind#") {
        headerParsed = true;
        continue;
      }
      if (!title && trimmed) {
        title = trimmed;
        continue;
      }
      if (!modified && trimmed) {
        modified = trimmed;
        continue;
      }
      continue;
    }
    if (!trimmed) continue;
    if (trimmed.startsWith("ncl# ###") || trimmed.startsWith("###")) {
      const heading = trimmed.replace(/^ncl# ###|^###/, "").trim();
      parsed += `
<p class="highlighted terms-category-sub">${heading}</p>`;
    } else if (
      trimmed.startsWith("ncl# ##") ||
      trimmed.startsWith("##")
    ) {
      const heading = trimmed.replace(/^ncl# ##|^##/, "").trim();
      parsed += `
	<p class="highlighted terms-category-sub">${heading}</p>`;
    } else if (trimmed.startsWith("#")) {
      const heading = trimmed.replace(/^#/, "").trim();
      parsed += `
	<p class="highlighted terms-category">${heading}<p>`;
    } else if (trimmed.startsWith("c#")) {
      const text = trimmed.replace(/^c#/, "").trim();
      parsed += `
	<p class="comment terms-contact">${text}</p>`;
    } else {
      const text = trimmed.replace(
        /\[\[href\]\](.*?)\[\[href\]\]/g,
        '<a class="highlighted link interactive" onclick="window._openURL(\'$1\')">$1</a>',
      );
      parsed += `<p class="text terms-text">${text}</p>`;
    }
  }
  return [parsed, title, modified];
}
