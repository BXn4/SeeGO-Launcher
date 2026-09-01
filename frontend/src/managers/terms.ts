import { writable } from "svelte/store";
import { GetCachedTerms } from "../../bindings/seegolauncher/internal/services/cacheservice";
import { parseTerms } from "../utils/string";

export let terms = writable<string>("");
export let title = writable<string>("");
export let modified = writable<string>("");

export async function initTerms() {
  const raw = await GetCachedTerms();
  let termsData;
  let titleData;
  let modifiedData;
  [termsData, titleData, modifiedData] = parseTerms(raw);

  terms.set(termsData)
  title.set(titleData)
  modified.set(modifiedData)
}
