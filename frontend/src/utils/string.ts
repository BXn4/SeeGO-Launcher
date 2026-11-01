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

export function formatNewsMarkdownToHTML(v: string): string {
  const esc = (s: string) =>
    s
      .replace(/&/g, "&amp;")
      .replace(/</g, "&lt;")
      .replace(/>/g, "&gt;")
      .replace(/"/g, "&quot;")
      .replace(/'/g, "&#39;");

  const linkCard = (url: string) => {
    const yt = url.match(/(?:youtu\.be\/|[?&;]v=)([\w-]{11})/);
    const domain = url.replace(/^https?:\/\//, "").split("/")[0];
    const thumb = yt
      ? `<img src="https://img.youtube.com/vi/${yt[1]}/hqdefault.jpg" alt="">`
      : "";

    return (
      `<a class="link-card interactive" onclick="window._openURL('${url}')">` +
      thumb +
      `<span class="link-card-info">` +
      `<span class="link-card-url highlighted link interactive">${url}</span>` +
      `</span></a>`
    );
  };

  return v
    .replace(/\r\n/g, "\n")
    .split("\n")
    .map((raw) => {
      const m = raw.match(/^\s*(c#)?\s*(#{1,3})?\s*(c#)?/)!;
      const size = m[2]?.length ?? 0;

      let text = raw.slice(m[0].length).replace(/\[\[(?!href\]\])[^\]]*\]\]/g, "");
      text = esc(text).replace(/\[\[href\]\](.*?)\[\[href\]\]/g, (_, url) => linkCard(url));

      if (size) return `<h${size}>${text}</h${size}>`;
      return `<p>${text}</p>`;
    })
    .join("");
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
